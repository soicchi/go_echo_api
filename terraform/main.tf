module "go_api" {
  # https://registry.terraform.io/modules/terraform-aws-modules/lambda/aws/latest
  source = "terraform-aws-modules/lambda/aws"

  function_name = "go-echo-api"
  description   = "Go Echo API"

  create_package = false
  image_uri      = "${data.aws_ecr_repository.go_echo_api.repository_url}:latest"
  package_type   = "Image"

  create_current_version_allowed_triggers = false

  allowed_triggers = {
    APIGatewayAny = {
      service = "apigateway"
      source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*"
    }
  }

  vpc_subnet_ids = [
    local.subnet_id_web_app_a,
  ]
}

module "api_gateway" {
  # https://registry.terraform.io/modules/terraform-aws-modules/apigateway-v2/aws/latest
  source = "terraform-aws-modules/apigateway-v2/aws"

  name                   = "go-echo-api"
  description            = "Go Echo API"
  protocol_type          = "HTTP"
  create_api_domain_name = false

  default_stage_access_log_destination_arn = module.logs.cloudwatch_log_group_arn
  default_stage_access_log_format = "$context.identity.sourceIp - - [$context.requestTime] \"$context.httpMethod $context.routeKey $context.protocol\" $context.status $context.responseLength $context.requestId $context.integrationErrorMessage"

  # Temporary all allow
  cors_configuration = {
    allow_headers = ["*"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }

  integrations = {
    "ANY /" = {
      lambda_arn = module.go_api.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds  = 30000
    }
  }
}

module "logs" {
  # https://registry.terraform.io/modules/terraform-aws-modules/cloudwatch/aws/latest
  source = "terraform-aws-modules/cloudwatch/aws//modules/log-group"
  version = "~> 3.0"

  name = "api-gateway/go-echo-api"
  # temporary set to 7 days
  retention_in_days = 7
}

module "rds_proxy" {
  # https://registry.terraform.io/modules/terraform-aws-modules/rds-proxy/aws/latest
  source = "terraform-aws-modules/rds-proxy/aws"

  name = "go-echo-api"
  iam_role_name = "go-echo-api-rds-proxy-role"
  vpc_subnet_ids = [
    local.subnet_id_web_app_a,
    local.subnet_id_web_app_c
  ]
  vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]

  engine_family = "POSTGRESQL"
  debug_logging = true

  auth = {
    "postgres" = {
      secret_arn = module.secret_manager.secret_arn
    }
  }

  target_db_instance = true
  db_instance_identifier = module.rds.db_instance_identifier
}

module "rds_proxy_sg" {
  # https://registry.terraform.io/modules/terraform-aws-modules/security-group/aws/latest
  source = "terraform-aws-modules/security-group/aws"

  name = "go-echo-api"
  description = "Security group for RDS Proxy"
  vpc_id = local.vpc_id

  ingress_cidr_blocks = [local.subnet_cidr_block_web_app_a]
  
  egress_cidr_blocks = ["0.0.0.0/0"]
  egress_ipv6_cidr_blocks = ["::/0"]
  egress_rules = ["all-all"]
}

module "rds" {
  # https://registry.terraform.io/modules/terraform-aws-modules/rds/aws/latest
  source = "terraform-aws-modules/rds/aws"

  identifier = "go-echo-api"

  engine = "postgres"
  engine_version = "14"
  family = "postgres14"
  major_engine_version = "14"
  instance_class = "db.t3.micro"
  allocated_storage = 20

  db_name = "postgres"
  username = "postgres"
  password = "password"
  port = 5432

  max_allocated_storage           = 100
  enabled_cloudwatch_logs_exports = ["postgresql"]
  storage_encrypted               = false

  # backup_window           = "17:00-17:30"
  # backup_retention_period = 5

  iam_database_authentication_enabled = true

  vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]

  create_db_subnet_group = true
  subnet_ids = [local.subnet_id_web_app_a, local.subnet_id_web_app_c]
}

module "secret_manager" {
  # https://registry.terraform.io/modules/terraform-aws-modules/secrets-manager/aws/latest
  source = "terraform-aws-modules/secrets-manager/aws"

  name_prefix = "go-echo-api"
  description = "Secrets for Go Echo API"
  recovery_window_in_days = 30

  create_policy = true
  block_public_policy = false
  policy_statements = {
    read = {
      principals = [{
        type = "AWS"
        identifiers = ["*"]
      }]
      actions = ["secretsmanager:GetSecretValue"]
      resources = ["*"]
    }
  }

  secret_string = jsonencode({
    username = "postgres"
    password = "password"
  })
}
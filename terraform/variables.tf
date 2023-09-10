data "aws_ecr_repository" "go_echo_api" {
  name = "go-echo-api"
}

locals {
  prefix = "go-echo-api"
}
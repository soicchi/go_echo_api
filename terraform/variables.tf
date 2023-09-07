data "aws_ecr_repository" "go_echo_api" {
  name = "go-echo-api"
}

data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    region = "ap-northeast-1"
    bucket = "miitel-tfstate"
    key    = "dev/network.tfstate"
  }
}

locals {
  prefix = "d02-tyo-scan-to-call"

  # Remote resources
  vpc_id = data.terraform_remote_state.network.outputs.network.vpc.id
  subnet_cidr_block_web_app_a = data.terraform_remote_state.network.outputs.subnets.web-application-a.cidr_block
  subnet_id_web_app_a = data.terraform_remote_state.network.outputs.subnets.web-application-a.id
  subnet_id_web_app_c = data.terraform_remote_state.network.outputs.subnets.web-application-c.id
}
terraform {
  required_version = ">= 1.0.0"
  required_providers {
    aws = {
        source = "hashicorp/aws"
        version = "~> 5.0.0"
    }
  }

  backend "s3" {
    region = "ap-northeast-1"
    bucket = "terraform-nishizono"
    key = "go-echo-api.tfstate"
  }
}

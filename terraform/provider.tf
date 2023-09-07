provider "aws" {
  shared_credentials_files = ["$HOME/.aws/credentials"]
  profile                  = "miitel-terraform-dev"
  region                   = "ap-northeast-1"

  default_tags {
    tags = {
      Creator = "nishizono"
      Purpose = "API test for Echo"
    }
  }
}
terraform {
# TODO: setting backend(s3) for tfstate file
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.64"
    }
  }
  required_version = ">= 1.5.7"
}

provider "aws" {
  profile = local.aws_profile
  region  = "ap-northeast-1"
  default_tags {
    tags = {
      Product = "hybird"
    }
  }
}

# local-exec for build and push of docker image
resource "null_resource" "build_push_dkr_img" {
  triggers = {
#     detect_docker_source_changes = local.dkr_img_src_sha256
    detect_docker_source_changes = timestamp()
  }
  provisioner "local-exec" {
    command = local.dkr_build_cmd
  }
}

module "vpc" {
  source = "./modules/vpc"

  cidr_block = "10.0.0.0/16"
  availability_zones = ["ap-northeast-1a", "ap-northeast-1c"]
  vpc_name = "hybird-vpc"
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
}

module "ecr" {
  source = "./modules/ecr"
  repository_name = local.ecr_repo
}

data "aws_ecr_image" "latest" {
  repository_name = local.ecr_repo
  most_recent     = true
}

module "iam" {
  source = "./modules/iam"
}

module mysql {
  source = "./modules/rds/mysql"
  vpc_id =  module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnet_ids
  vpc_cidr_block = module.vpc.vpc_cidr_block
}

module "lambda" {
  source = "./modules/lambda"
  vpc_id = module.vpc.vpc_id
  vpc_cidr_block = module.vpc.vpc_cidr_block
  image_uri   = data.aws_ecr_image.latest.image_uri
  subnet_id   = module.vpc.private_subnet_ids[0]
  lambda_exec_role_arn = module.iam.lambda_exec_role_arn
  db_host = module.mysql.db_host
  db_name = module.mysql.db_name
  db_pass = module.mysql.db_pass
  db_user = module.mysql.db_user
  cors_allowed_origin = local.cors_allowed_origin
}
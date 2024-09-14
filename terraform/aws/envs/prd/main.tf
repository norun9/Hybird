terraform {
  backend "s3" {
    bucket         = "hybird-terraform-tfstate"
    key            = "./terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "Connections" # DynamoDB table for lock management (optional)
    encrypt        = true
  }
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
  region  = local.aws_region
  default_tags {
    tags = {
      Product = "hybird"
    }
  }
}

# local-exec for build and push of docker image
resource "null_resource" "build_push_dkr_img" {
  triggers = {
    detect_docker_source_changes = timestamp()
  }
  provisioner "local-exec" {
    command = local.api_dkr_build_cmd
  }
}

resource "null_resource" "build_push_ws_connect_dkr_img" {
  triggers = {
    detect_docker_source_changes = timestamp()
  }
  provisioner "local-exec" {
    command = local.ws_connect_dkr_build_cmd
  }
}

resource "null_resource" "build_push_ws_disconnect_dkr_img" {
  triggers = {
    detect_docker_source_changes = timestamp()
  }
  provisioner "local-exec" {
    command = local.ws_disconnect_dkr_build_cmd
  }
}

resource "null_resource" "build_push_ws_default_dkr_img" {
  triggers = {
    detect_docker_source_changes = timestamp()
  }
  provisioner "local-exec" {
    command = local.ws_default_dkr_build_cmd
  }
}

# NOTE: Maybe I don't need to modularize it, but I'd like to use 'module'.

module "vpc" {
  source = "./modules/vpc"

  cidr_block                            = "10.0.0.0/16"
  availability_zones_for_private_subnet = ["ap-northeast-1a", "ap-northeast-1c"]
  availability_zone_for_public_subnet   = "ap-northeast-1a"
  public_subnet                         = "10.0.3.0/24"
  vpc_name                              = "hybird-vpc"
  private_subnets                       = ["10.0.1.0/24", "10.0.2.0/24"]
}

module "ecr" {
  source          = "./modules/ecr"
  repository_name = local.ecr_repo
}

data "aws_ecr_image" "api" {
  repository_name = local.ecr_repo
  image_tag       = "api"
  depends_on      = [null_resource.build_push_dkr_img]
}

data "aws_ecr_image" "ws_connect" {
  repository_name = local.ecr_repo
  image_tag       = "ws_connect"
  depends_on      = [null_resource.build_push_ws_connect_dkr_img]
}

data "aws_ecr_image" "ws_disconnect" {
  repository_name = local.ecr_repo
  image_tag       = "ws_disconnect"
  depends_on      = [null_resource.build_push_ws_disconnect_dkr_img]
}

data "aws_ecr_image" "ws_default" {
  repository_name = local.ecr_repo
  image_tag       = "ws_default"
  depends_on      = [null_resource.build_push_ws_default_dkr_img]
}

module "iam" {
  source = "./modules/iam"
}

module "mysql" {
  source         = "./modules/rds/mysql"
  vpc_id         = module.vpc.vpc_id
  subnet_ids     = module.vpc.private_subnet_ids
  vpc_cidr_block = module.vpc.vpc_cidr_block
}

module "lambda" {
  source               = "./modules/lambda"
  vpc_id               = module.vpc.vpc_id
  vpc_cidr_block       = module.vpc.vpc_cidr_block
  image_uri            = data.aws_ecr_image.api.image_uri
  subnet_id            = module.vpc.private_subnet_ids[0]
  lambda_exec_role_arn = module.iam.lambda_exec_role_arn
  db_host              = module.mysql.db_host
  db_name              = module.mysql.db_name
  db_pass              = module.mysql.db_pass
  db_user              = module.mysql.db_user
  cors_allowed_origin  = local.cors_allowed_origin
}

module "api_gw" {
  source               = "./modules/api_gw"
  lambda_invoke_arn    = module.lambda.lambda_invoke_arn
  lambda_function_name = module.lambda.lambda_function_name
}

module "ec2" {
  source                = "./modules/ec2"
  vpc_id                = module.vpc.vpc_id
  vpc_cidr_block        = module.vpc.vpc_cidr_block
  subnet_id             = module.vpc.public_subnet_id
  ec2_bastion_role_name = module.iam.ec2_bastion_role_name
}
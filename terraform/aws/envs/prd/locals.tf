locals {
  aws_account = var.aws_account # AWS account
  aws_region = var.aws_region   # AWS region
  aws_profile = var.aws_profile # AWS profile

  cors_allowed_origin = var.cors_allowed_origin

  ecr_reg   = "${local.aws_account}.dkr.ecr.${local.aws_region}.amazonaws.com"
  ecr_repo  = "hybird_repo"

  dkr_build_context_path = "${path.module}/../../../../backend"
  dkr_img_src_path = "${path.module}/../../../../backend/api"
  ws_dkr_img_src_path = "${path.module}/../../../../backend/lambda/websocket"
  dkr_img_src_sha256 = filesha256("${local.dkr_img_src_path}/Dockerfile")

  api_dkr_build_cmd = <<-EOT
        docker build --provenance=false --no-cache -t ${local.ecr_reg}/${local.ecr_repo}:api \
            -f ${local.dkr_img_src_path}/Dockerfile ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:api
  EOT

  ws_connect_dkr_build_cmd = <<-EOT
        docker build --provenance=false --no-cache --build-arg LAMBDA_SOURCE_DIR=connect -t ${local.ecr_reg}/${local.ecr_repo}:ws_connect \
            -f ${local.ws_dkr_img_src_path}/Dockerfile ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:ws_connect
  EOT

  ws_disconnect_dkr_build_cmd = <<-EOT
        docker build --provenance=false --no-cache --build-arg LAMBDA_SOURCE_DIR=disconnect -t ${local.ecr_reg}/${local.ecr_repo}:ws_disconnect \
            -f ${local.ws_dkr_img_src_path}/Dockerfile ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:ws_disconnect
  EOT

  ws_default_dkr_build_cmd = <<-EOT
        docker build --provenance=false --no-cache --build-arg LAMBDA_SOURCE_DIR=_default -t ${local.ecr_reg}/${local.ecr_repo}:ws_default \
            -f ${local.ws_dkr_img_src_path}/Dockerfile ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:ws_default
  EOT
}
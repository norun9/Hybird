locals {
  aws_account = var.aws_account # AWS account
  aws_region  = var.aws_region  # AWS region
  aws_profile = var.aws_profile # AWS profile

  apigateway_deploy_stage = "prd"

  cors_allowed_origin = var.cors_allowed_origin

  ecr_reg  = "${local.aws_account}.dkr.ecr.${local.aws_region}.amazonaws.com"
  ecr_repo = "hybird_repo"

  dkr_build_context_path = "${path.module}/../../../../backend"
  dkr_img_src_path       = "${local.dkr_build_context_path}/api"
  ws_dkr_img_src_path    = "${local.dkr_build_context_path}/lambda/websocket"
  dkr_img_src_sha256     = filesha256("${local.dkr_img_src_path}/Dockerfile")

  # common command template for building and pushing Docker image to ECR
  dkr_img_build_push_cmd = <<-EOT
        docker build --provenance=false --no-cache {build_arg} -t ${local.ecr_reg}/${local.ecr_repo}:{tag} \
            -f {dkrfile_path} ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:{tag}
  EOT

  # command for building "REST API" Docker image
  api_dkr_build_cmd = replace(
    replace(
  replace(local.dkr_img_build_push_cmd, "{build_arg}", ""), "{tag}", "api"), "{dkrfile_path}", "${local.dkr_img_src_path}/Dockerfile")

  # command for building WebSocket "connect" Docker image
  ws_connect_dkr_build_cmd = replace(
    replace(
      replace(
        local.dkr_img_build_push_cmd, "{build_arg}", "--build-arg LAMBDA_SOURCE_DIR=connect"
      ), "{tag}", "ws_connect"
    ), "{dkrfile_path}", "${local.ws_dkr_img_src_path}/Dockerfile"
  )

  # command for building WebSocket "disconnect" Docker image
  ws_disconnect_dkr_build_cmd = replace(
    replace(
      replace(
        local.dkr_img_build_push_cmd, "{build_arg}", "--build-arg LAMBDA_SOURCE_DIR=disconnect"
      ), "{tag}", "ws_disconnect"
    ), "{dkrfile_path}", "${local.ws_dkr_img_src_path}/Dockerfile"
  )

  # command for building WebSocket "default" Docker image
  ws_default_dkr_build_cmd = replace(
    replace(
      replace(
        local.dkr_img_build_push_cmd, "{build_arg}", "--build-arg LAMBDA_SOURCE_DIR=_default"
      ), "{tag}", "ws_default"
    ), "{dkrfile_path}", "${local.ws_dkr_img_src_path}/Dockerfile"
  )
}
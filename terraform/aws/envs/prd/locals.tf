locals {
  aws_account = var.aws_account # AWS account
  aws_region = var.aws_region   # AWS region
  aws_profile = var.aws_profile # AWS profile

  cors_allowed_origin = var.cors_allowed_origin

  ecr_reg   = "${local.aws_account}.dkr.ecr.${local.aws_region}.amazonaws.com"
  ecr_repo  = "hybird_repo"
  image_tag = "latest"

  dkr_build_context_path = "${path.module}/../../../../backend"
  dkr_img_src_path = "${path.module}/../../../../backend/api"
  dkr_img_src_sha256 = filesha256("${local.dkr_img_src_path}/Dockerfile")

  dkr_build_cmd = <<-EOT
        docker build --no-cache -t ${local.ecr_reg}/${local.ecr_repo}:${local.image_tag} \
            -f ${local.dkr_img_src_path}/Dockerfile ${local.dkr_build_context_path}

        aws --profile ${local.aws_profile} ecr get-login-password --region ${local.aws_region} | \
            docker login --username AWS --password-stdin ${local.ecr_reg}

        docker push ${local.ecr_reg}/${local.ecr_repo}:${local.image_tag}
  EOT
}
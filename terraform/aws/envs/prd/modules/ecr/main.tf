resource "aws_ecr_repository" "this" {
  name = var.repository_name
  image_tag_mutability = "MUTABLE"
}

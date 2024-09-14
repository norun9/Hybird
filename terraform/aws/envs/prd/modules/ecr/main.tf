resource "aws_ecr_repository" "this" {
  name                 = var.repository_name
  image_tag_mutability = "MUTABLE"

  lifecycle {
    prevent_destroy = true
    ignore_changes  = [image_tag_mutability]
  }
}
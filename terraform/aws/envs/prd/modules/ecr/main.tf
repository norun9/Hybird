resource "aws_ecr_repository" "this" {
  name                 = var.repository_name
  image_tag_mutability = "MUTABLE"
  force_delete         = true

  #   lifecycle {
  #     prevent_destroy = true
  #     ignore_changes  = [image_tag_mutability]
  #   }
  lifecycle {
    prevent_destroy = false # Ensure this is not preventing the deletion
  }
}
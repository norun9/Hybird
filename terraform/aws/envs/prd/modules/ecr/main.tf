resource "aws_ecr_repository" "this" {
  name = var.repository_name
  image_tag_mutability = "MUTABLE"

  lifecycle {
    prevent_destroy = true
    ignore_changes  = [image_tag_mutability]
  }
}

resource "aws_ecr_lifecycle_policy" "this" {
  repository = aws_ecr_repository.this.name

  policy = <<EOF
  {
    "rules": [
      {
        "rulePriority": 1,
        "description": "Retain only the latest tagged image",
        "selection": {
          "tagStatus": "tagged",
          "countType": "imageCountMoreThan",
          "countNumber": 1
        },
        "action": {
          "type": "expire"
        }
      },
      {
        "rulePriority": 2,
        "description": "Delete untagged images",
        "selection": {
          "tagStatus": "untagged",
          "countType": "imageCountMoreThan",
          "countNumber": 0
        },
        "action": {
          "type": "expire"
        }
      }
    ]
  }
  EOF
}


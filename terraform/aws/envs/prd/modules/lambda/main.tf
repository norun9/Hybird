resource "aws_security_group" "lambda_sg" {
  name        = "lambda_security_group"
  description = "Security group for Lambda function"
  vpc_id      = var.vpc_id

  # Inbound rules for the Go Gin server running inside the Lambda function.
  # This rule allows traffic on port 8080, which is the port where the Go Gin server listens.
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]  # Adjust as per your VPC's CIDR
  }

  # Inbound rules (example: allow outbound to RDS on port 3306)
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]  # Adjust as per your VPC's CIDR
  }

  # Outbound rules (allow all by default)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lambda_function" "hybird_lambda" {
  function_name = "hybird_lambda_function"
  role          = var.lambda_exec_role_arn
  package_type  = "Image"
  image_uri     = var.image_uri
  timeout = 10
  memory_size   = 256

  environment {
    variables = {
      GIN_MODE           = "release"
      HYBIRD_ENV_NAME    = "prd"
      HYBIRD_PRD_DB_HOST = var.db_host
      HYBIRD_PRD_DB_NAME = var.db_name
      HYBIRD_PRD_DB_PASS = var.db_pass
      HYBIRD_PRD_DB_USER = var.db_user
      HYBIRD_PRD_HTTP_CORS_ALLOWED_ORIGIN = var.cors_allowed_origin
    }
  }

  vpc_config {
    subnet_ids         = [var.subnet_id]
    security_group_ids = [aws_security_group.lambda_sg.id]
  }
}
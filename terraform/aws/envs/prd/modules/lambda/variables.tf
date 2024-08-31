variable "subnet_id" {
  description = "The ID of the subnet to associate with the Lambda function"
  type        = string
}

variable "vpc_id" {
  description = "The ID of the vpc to associated with the security group"
}

variable "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
  type = string
}

variable "image_uri" {
  description = "The URI of the Docker image to use for the Lambda function"
  type        = string
}

variable "lambda_exec_role_arn" {
  description = "The ARN of the IAM role that the Lambda function will assume"
  type        = string
}

variable "cors_allowed_origin" {
  type        = string
}

variable "db_host" {
  description = "The hostname of the production database"
  type        = string
}

variable "db_name" {
  description = "The name of the production database"
  type        = string
}

variable "db_pass" {
  description = "The password for the production database"
  type        = string
  sensitive   = true
}

variable "db_user" {
  description = "The username for the production database"
  type        = string
}

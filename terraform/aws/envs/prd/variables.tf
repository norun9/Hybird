variable "project_name" {
  description = "Project name"
  type        = string
  default     = "hybird"
}

variable "site_domain" {
  description = "Web site domain name"
  type        = string
}

variable "aws_account" {
  description = "AWS account ID"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-northeast-1" # Optional: Set a default region
}

variable "aws_profile" {
  description = "AWS CLI profile name"
  type        = string
  default     = "default" # Optional: Set a default profile
}

variable "cors_allowed_origin" {
  type = string
}
variable "vpc_id" {
  description = "The ID of the vpc to associated with the security group"
}

variable "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
  type = string
}

variable "subnet_id" {
  description = "The ID of the subnet to associate with the Lambda function"
  type        = string
}

variable "ec2_bastion_role_name" {
  type = string
}

variable "ami_image_for_bastion" {
  default = "al2023-ami-2023.1.*-kernel-6.*-x86_64"
}
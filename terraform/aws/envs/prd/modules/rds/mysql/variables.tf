variable "subnet_ids" {
  description = "The ID of the subnet to associate with the RDS function"
  type        = list(string)
}

variable "vpc_id" {
  description = "The ID of the vpc to associated with the security group"
  type = string
}

variable "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
  type = string
}
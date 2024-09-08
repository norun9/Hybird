variable "cidr_block" {
  description = "CIDR block for the VPC"
  type        = string
}

variable "private_subnets" {
  description = "List of private subnet CIDR blocks"
  type        = list(string)
}

variable "public_subnet" {
  description = "Public subnet CIDR blocks"
  type        = string
}

variable "availability_zone_for_public_subnet" {
  description = "Availability zone for the public subnet"
  type = string
}


variable "vpc_name" {
  description = "Name of the VPC"
  type        = string
}

variable "availability_zones_for_private_subnet" {
  description = "List of availability zones for the subnets"
  type        = list(string)
}
resource "aws_vpc" "this" {
  cidr_block                       = var.cidr_block
  instance_tenancy                 = "default"
  enable_dns_support               = true
  enable_dns_hostnames = true
  assign_generated_ipv6_cidr_block = false
  tags = {
    Name = "hybird-vpc"
  }
}

resource "aws_subnet" "private" {
  count = length(var.private_subnets)
  vpc_id                  = aws_vpc.this.id
  availability_zone       = element(var.availability_zones, count.index)
  cidr_block              = var.private_subnets[count.index]
  map_public_ip_on_launch = false
  tags = {
    Name = "private-subnet-${count.index + 1}"
  }
}
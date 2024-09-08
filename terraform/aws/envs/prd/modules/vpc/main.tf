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
  availability_zone       = element(var.availability_zones_for_private_subnet, count.index)
  cidr_block              = var.private_subnets[count.index]
  map_public_ip_on_launch = false
  tags = {
    Name = "private-subnet-${count.index + 1}"
  }
}

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = {
    Name = "hybird-igw"
  }
}


resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.this.id
  cidr_block              = var.public_subnet
  map_public_ip_on_launch = true
  availability_zone       = var.availability_zone_for_public_subnet
}

resource "aws_route_table" "hybird_public_rt" {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }

  tags = {
    Name = "hybird-public-rt"
  }
}

resource "aws_route_table_association" "my_public_route_table_assoc" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.hybird_public_rt.id
}
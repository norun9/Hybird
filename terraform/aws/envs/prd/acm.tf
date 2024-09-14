provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

data "aws_acm_certificate" "cert" {
  provider    = aws.us_east_1
  domain      = "hybird.click"
  statuses    = ["ISSUED"]
  most_recent = true
}
resource "aws_cloudfront_distribution" "hybird_distribution" {
  aliases = [
    "hybird.click",
  ]
  default_root_object = "index.html"
  enabled             = true
  is_ipv6_enabled     = true
  price_class         = "PriceClass_All"
  wait_for_deployment = true

  default_cache_behavior {
    allowed_methods = [
      "GET",
      "HEAD",
    ]
    cache_policy_id = "658327ea-f89d-4fab-a63d-7e88639e58f6"
    cached_methods = [
      "GET",
      "HEAD",
    ]
    compress               = true
    default_ttl            = 0
    max_ttl                = 0
    min_ttl                = 0
    smooth_streaming       = false
    target_origin_id       = "hybird-ssg.s3-website-ap-northeast-1.amazonaws.com"
    viewer_protocol_policy = "https-only"
  }

  origin {
    connection_attempts = 3
    connection_timeout  = 10
    domain_name         = "hybird-ssg.s3-website-ap-northeast-1.amazonaws.com"
    origin_id           = "hybird-ssg.s3-website-ap-northeast-1.amazonaws.com"

    custom_header {
      name  = "Referer"
      value = "https://d2o614elmpk6gf.cloudfront.net"
    }

    custom_origin_config {
      http_port                = 80
      https_port               = 443
      origin_keepalive_timeout = 5
      origin_protocol_policy   = "http-only"
      origin_read_timeout      = 30
      origin_ssl_protocols = [
        "TLSv1",
        "TLSv1.1",
        "TLSv1.2",
      ]
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn            = data.aws_acm_certificate.cert.arn
    cloudfront_default_certificate = false
    minimum_protocol_version       = "TLSv1.2_2021"
    ssl_support_method             = "sni-only"
  }
}

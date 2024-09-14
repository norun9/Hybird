resource "aws_route53_zone" "hybird_click" {
  name = "hybird.click"
}

resource "aws_route53_record" "hybird_click_a" {
  zone_id = aws_route53_zone.hybird_click.zone_id
  name    = "hybird.click"
  type    = "A"

  alias {
    name                   = "d2o614elmpk6gf.cloudfront.net"
    zone_id                = "Z2FDTNDATAQYW2"
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "hybird_click_ns" {
  zone_id = aws_route53_zone.hybird_click.zone_id
  name    = "hybird.click"
  type    = "NS"
  ttl     = 172800
  records = ["ns-89.awsdns-11.com.", "ns-1160.awsdns-17.org.", "ns-788.awsdns-34.net.", "ns-1992.awsdns-57.co.uk."]
}

resource "aws_route53_record" "hybird_click_soa" {
  zone_id = aws_route53_zone.hybird_click.zone_id
  name    = "hybird.click"
  type    = "SOA"
  ttl     = 900
  records = ["ns-89.awsdns-11.com. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400"]
}

resource "aws_route53_record" "hybird_click_cname" {
  name = "_1450c9721804eef38b340f9fe8293d31.hybird.click"
  records = [
    "_6e87faf46317782636c65b5bb585d59d.sdgjtdhdhz.acm-validations.aws.",
  ]
  ttl     = 300
  type    = "CNAME"
  zone_id = aws_route53_zone.hybird_click.zone_id
}

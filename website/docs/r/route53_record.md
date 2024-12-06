---
subcategory: "Route 53"
layout: "aws"
page_title: "aws_route53_record"
description: |-
  Provides a Route53 record resource.
---

# Resource: aws_route53_record

Provides a Route53 record resource.

## Example Usage

### Simple routing policy

```terraform
resource "aws_route53_record" "www" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "www.example.com"
  type    = "A"
  ttl     = "300"
  records = [aws_eip.lb.public_ip]
}
```

### NS Record Management

When creating Route 53 zones, the `NS` records for the zone are automatically created. Enabling the `allow_overwrite` argument will allow managing these records in a single Terraform run without the requirement for `terraform import`.

```terraform
resource "aws_route53_zone" "example" {
  name = "test.example.com"
}

resource "aws_route53_record" "example" {
  allow_overwrite = true
  name            = "test.example.com"
  ttl             = 172800
  type            = "NS"
  zone_id         = aws_route53_zone.example.zone_id

  records = [
    aws_route53_zone.example.name_servers[0],
    aws_route53_zone.example.name_servers[1],
  ]
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The ID of the hosted zone to contain this record.
* `name` - (Required) The name of the record.
* `type` - (Required) The record type. Valid values are `A`, `AAAA`, `CNAME`, `MX`, `NS`, `PTR`, `SRV` and `TXT`.
* `ttl` - (Required) The TTL of the record.
* `records` - (Required) A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\" \"` inside the Terraform configuration string to split characters into multiple text strings (e.g., `"first255characters\" \"next255characters"`).
* `allow_overwrite` - (Optional) Allow creation of this record in Terraform to overwrite an existing record, if any. This does not affect the ability to update the record in Terraform and does not prevent other resources within Terraform or manual Route 53 changes outside Terraform from overwriting this record. `false` by default. This configuration is not recommended for most environments.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - The name of the record.
* `fqdn` - [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `set_identifier` - Unique identifier to differentiate records with routing policies from one another. Always empty.
* `health_check_id` - The health check the record should be associated with. Always empty.
* `alias` - An alias block. Always empty.
* `failover_routing_policy` - A block indicating the routing behavior when associated health check fails. Always empty.
* `geolocation_routing_policy` - A block indicating a routing policy based on the geolocation of the requestor. Always empty.
* `latency_routing_policy` - A block indicating a routing policy based on the latency between the requestor and an AWS region. Always empty.
* `weighted_routing_policy` - A block indicating a weighted routing policy. Always empty.
* `multivalue_answer_routing_policy` - Set to `true` to indicate a multivalue answer routing policy. Always empty.

## Import

Route53 Records can be imported using ID of the record, which is the zone identifier, record name, and record type, separated by underscores (`_`)E.g.,

```
$ terraform import aws_route53_record.myrecord z-xxxxxxxx_dev.example.com_NS
```

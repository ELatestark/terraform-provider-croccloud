---
subcategory: "Transit Gateway"
layout: "aws"
page_title: "aws_ec2_transit_gateway"
description: |-
  Provides information about a transit gateway.
---

[describe-tgw]: https://docs.cloud.croc.ru/en/api/ec2/transit_gateways/DescribeTransitGateways.html

# Data Source: aws_ec2_transit_gateway

Provides information about a transit gateway.

## Example Usage

### By Filter

```terraform
data "aws_ec2_transit_gateway" "selected" {
  filter {
    name   = "owner-id"
    values = ["project@customer"]
  }
}
```

### By Identifier

```terraform
data "aws_ec2_transit_gateway" "selected" {
  id = "tgw-12345678"
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) One or more configuration blocks containing name-values filters.
  The structure of this block is [described below](#filter).
* `id` - (Optional) The ID of the transit gateway.

### filter

* `name` - (Required) The name of the field to filter by it.
  Valid values can be found in the [EC2 API documentation][describe-tgw].
* `values` - (Required) List of one or more values for the filter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `association_default_route_table_id` - The ID of the default association route table.
* `default_route_table_association` - Indicates whether the association with default association route table is created automatically.
* `default_route_table_propagation` - Indicates whether the routes are automatically propagated to the default propagation route table.
* `description` - The description of the transit gateway
* `owner_id` - The ID of the project that owns the transit gateway.
* `propagation_default_route_table_id` - The ID of the default propagation route table.
* `shared_owners` - List of project IDs that are granted access to the transit gateway.
* `tags` - Map of tags assigned to the transit gateway.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `amazon_side_asn` - Private Autonomous System Number (ASN) for the Amazon side of a BGP session. Always `0`.
* `arn` - The ARN of the transit gateway. Always `""`.
* `auto_accept_shared_attachments` - Whether resource attachment requests are automatically accepted. Always `""`.
* `dns_support` - Whether DNS support is enabled. Always `""`.
* `multicast_support` - Whether Multicast support is enabled. Always `""`.
* `transit_gateway_cidr_blocks` - One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Always empty.
* `vpn_ecmp_support` - Whether VPN Equal Cost Multipath Protocol support is enabled. Always `""`.


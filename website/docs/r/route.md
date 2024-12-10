---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_route"
description: |-
  Provides a resource to create a routing entry in a VPC routing table.
---

# Resource: aws_route

Provides a resource to create a routing table entry (a route) in a VPC routing table.

~> **NOTE on Route Tables and Routes:** Terraform currently provides both a standalone Route resource and a [`aws_route_table`][tf-route-table] resource with routes defined in-line. At this time you cannot use a Route Table with in-line routes in conjunction with any Route resources. Doing so will cause a conflict of rule settings and will overwrite rules.

## Example Usage

```terraform
variable route_table_id {}
variable instance_id {}

resource "aws_route" "example" {
  route_table_id         = var.route_table_id
  destination_cidr_block = "10.0.0.0/22"
  instance_id            = var.instance_id
}
```

## Argument Reference

The following arguments are supported:

* `route_table_id` - (Required) ID of the routing table.

One of the following destination arguments must be supplied:

* `destination_cidr_block` - (Required) The destination CIDR block.

One of the following target arguments must be supplied:

* `gateway_id` - (Optional) ID of an internet gateway or virtual private gateway.
* `instance_id` - (Optional) ID of an EC2 instance.
* `network_interface_id` - (Optional) ID of an EC2 network interface.
* `transit_gateway_id` - (Optional) The ID of the transit gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

~> **NOTE:** Only the arguments that are configured (one of the above) will be exported as an attribute once the resource is created.

* `id` - Route identifier computed from the routing table identifier and route destination.
* `instance_owner_id` - The project ID that owns the EC2 instance.
* `origin` - How the route was created - `CreateRouteTable`, `CreateRoute` or `EnableVgwRoutePropagation`.
* `state` - The state of the route - `active` or `blackhole`.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `carrier_gateway_id` - ID of a carrier gateway. Always `""`.
* `core_network_arn` - ARN of a core network. Always `""`.
* `destination_ipv6_cidr_block` - The destination IPv6 CIDR block. Always `""`.
* `destination_prefix_list_id` - ID of a managed prefix list destination of the route. Always `""`.
* `egress_only_gateway_id` - ID of a VPC Egress Only Internet Gateway. Always `""`.
* `local_gateway_id` - ID of an Outpost local gateway. Always `""`.
* `nat_gateway_id` - ID of a VPC NAT gateway. Always `""`.
* `vpc_endpoint_id` - ID of a VPC Endpoint. Always `""`.
* `vpc_peering_connection_id` - ID of a VPC peering connection. Always `""`.

## Timeouts

`aws_route` provides the following [Timeouts](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts) configuration options:

- `create` - (Default `5 minutes`) Used for route creation
- `update` - (Default `2 minutes`) Used for route creation
- `delete` - (Default `5 minutes`) Used for route deletion

## Import

Individual routes can be imported using `ROUTETABLEID_DESTINATION`.

For example, import a route in route table `rtb-12345678` with an IPv4 destination CIDR of `10.1.0.0/16` like this:

```console
$ terraform import aws_route.my_route rtb-12345678_10.1.0.0/16
```

[tf-route-table]: route_table.html

---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_route_table"
description: |-
    Provides details about a specific Route Table
---

# Data Source: aws_route_table

`aws_route_table` provides details about a specific Route Table.

This resource can prove useful when a module accepts a Subnet ID as an input variable and needs to, for example, add a route in the Route Table.

## Example Usage

The following example shows how one might accept a Route Table ID as a variable and use this data source to obtain the data necessary to create a route.

```terraform
variable "vpc_id" {}
variable "network_interface_id" {}

data "aws_route_table" "selected" {
  vpc_id = var.vpc_id
}

resource "aws_route" "example" {
  route_table_id         = data.aws_route_table.selected.id
  destination_cidr_block = "10.0.0.0/22"
  network_interface_id   = var.network_interface_id
}
```

## Argument Reference

The arguments of this data source act as filters for querying the available Route Table in the current region. The given filters must match exactly one Route Table whose data will be exported as attributes.

The following arguments are optional:

* `filter` - (Optional) Configuration block. Detailed below.
* `subnet_id` - (Optional) ID of a Subnet which is associated with the Route Table (not exported if not passed as a parameter).
* `tags` - (Optional) Map of tags, each pair of which must exactly match a pair on the desired Route Table.
* `vpc_id` - (Optional) ID of the VPC that the desired Route Table belongs to.

### filter

Complex filters can be expressed using one or more `filter` blocks.

The following arguments are required:

* `name` - (Required) Name of the field to filter by it.
* `values` - (Required) Set of values that are accepted for the given field. A Route Table will be selected if any one of the given values matches.

For more information about filtering, see the [EC2 API documentation][describe-route-tables].

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `arn` - ARN of the route table.
* `associations` - List of associations with attributes detailed below.
* `routes` - List of routes with attributes detailed below.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `owner_id` - The ID of the project that owns the route table. Always `""`.

### routes

When relevant, routes are also exported with the following attributes:

For destinations:

* `cidr_block` - CIDR block of the route.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `destination_prefix_list_id` - ID of a managed prefix list destination of the route. Always `""`.
* `ipv6_cidr_block` - The IPv6 CIDR block of the route. Always `""`.

For targets:

* `gateway_id` - ID of the Internet Gateway or Virtual Private Gateway.
* `instance_id` - ID of the EC2 instance.
* `network_interface_id` - ID of the EC2 network interface.
* `transit_gateway_id` - The ID of the transit gateway.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `carrier_gateway_id` - ID of the Carrier Gateway. Always `""`.
* `core_network_arn` - ARN of the core network. Always `""`.
* `egress_only_gateway_id` - ID of the Egress Only Internet Gateway. Always `""`.
* `local_gateway_id` - Local Gateway ID. Always `""`.
* `nat_gateway_id` - NAT Gateway ID. Always `""`.
* `vpc_endpoint_id` - VPC Endpoint ID. Always `""`.
* `vpc_peering_connection_id` - VPC Peering ID. Always `""`.

### associations

Associations are also exported with the following attributes:

* `gateway_id` - ID of the Internet Gateway or Virtual Private Gateway.
* `main` - Whether the association is due to the main route table.
* `route_table_association_id` - Association ID.
* `route_table_id` - Route Table ID.
* `subnet_id` - Subnet ID. Only set when associated with a subnet.

[describe-route-tables]: https://docs.cloud.croc.ru/en/api/ec2/routes/DescribeRouteTables.html

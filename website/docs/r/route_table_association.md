---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_route_table_association"
description: |-
  Provides a resource to create an association between a route table and a subnet.
---

# Resource: aws_route_table_association

Provides a resource to create an association between a route table and a subnet.

## Example Usage

```terraform
resource "aws_vpc" "example" {
  cidr_block = "10.1.0.0/16"
}

resource "aws_route_table" "example" {
  vpc_id = aws_vpc.example.id
}

resource "aws_subnet" "example" {
  availability_zone = "ru-msk-vol52"
  vpc_id            = aws_vpc.example.id

  cidr_block = cidrsubnet(aws_vpc.example.cidr_block, 1, 0)
}

resource "aws_route_table_association" "example" {
  subnet_id      = aws_subnet.example.id
  route_table_id = aws_route_table.example.id
}
```

## Argument Reference

The following arguments are supported:

* `subnet_id` - (Required) Subnet ID to create an association.
* `route_table_id` - (Required) ID of the routing table to associate with.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the association

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `gateway_id` - Gateway ID to create an association. Always `""`.

## Import

EC2 Route Table Associations can be imported using the associated resource ID and route table ID
separated by a forward slash (`/`).

For example with EC2 subnets:

```
$ terraform import aws_route_table_association.assoc subnet-12345678/rtb-6F78E00
```

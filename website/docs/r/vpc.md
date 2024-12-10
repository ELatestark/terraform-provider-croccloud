---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_vpc"
description: |-
  Provides a VPC resource.
---

# Resource: aws_vpc

Provides a VPC resource.

For more information, see the documentation on [VPC][vpc].

## Example Usage

Basic usage:

```terraform
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}
```

Basic usage with tags:

```terraform
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "main"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Optional) The IPv4 CIDR block for the VPC.
* `enable_dns_support` - (Optional) A boolean flag to enable/disable DNS support in the VPC. Defaults true.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of VPC
* `id` - ID of the VPC
* `main_route_table_id` - ID of the main route table associated with
     this VPC. Note that you can change a VPC's main route table by using an
     [`aws_main_route_table_association`][tf-main-route-table-association].
* `default_network_acl_id` - ID of the network ACL created by default on VPC creation
* `default_security_group_id` - ID of the security group created by default on VPC creation
* `default_route_table_id` - ID of the route table created by default on VPC creation
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `assign_generated_ipv6_cidr_block` - (Optional) Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. Always `false`.
* `enable_classiclink` - Whether the VPC has Classiclink enabled. Always `false`.
* `enable_classiclink_dns_support` - Whether the ClassicLink DNS Support for the VPC is enabled. Always `false`.
* `enable_dns_hostnames` - Whether the VPC has DNS hostname support. Always `false`.
* `instance_tenancy` - The allowed tenancy of instances launched into the selected VPC. Always `default`.
* `ipv4_ipam_pool_id` -  ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. Always empty.
* `ipv4_netmask_length` - The netmask length of the IPv4 CIDR you want to allocate to this VPC. Always empty.
* `ipv6_association_id` - The association ID for the IPv6 CIDR block. Always `""`.
* `ipv6_cidr_block` - The IPv6 CIDR block. Always `""`.
* `ipv6_cidr_block_network_border_group` - The Network Border Group Zone name. Always `""`.
* `ipv6_ipam_pool_id` - IPAM Pool ID for a IPv6 pool. Always `""`.
* `ipv6_netmask_length` - Netmask length to request from IPAM Pool. Always `0`.
* `owner_id` - The ID of the project that owns the VPC. Always `""`.

## Import

VPCs can be imported using the `vpc id`, e.g.,

```
$ terraform import aws_vpc.test_vpc vpc-12345678
```

[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block
[tf-main-route-table-association]: main_route_table_association.html
[vpc]: https://docs.cloud.croc.ru/en/services/networks/privatecloud.html

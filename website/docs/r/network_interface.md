---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "aws_network_interface"
description: |-
  Provides an elastic network interface (ENI) resource.
---

# Resource: aws_network_interface

Provides an elastic network interface (ENI) resource.

For more information, see the documentation on [Network interfaces][network-interfaces].

[network-interfaces]: https://docs.cloud.croc.ru/en/services/networks/interfaces/operations.html

## Example Usage

```terraform
resource "aws_network_interface" "example" {
  subnet_id   = "subnet-12345678"
  private_ips = ["10.0.31.50"]

  attachment {
    instance     = "i-12345678"
    device_index = 1
  }
}
```

## Argument Reference

The following arguments are required:

* `subnet_id` - (Required) Subnet ID to create the ENI in.

The following arguments are optional:

* `attachment` - (Optional) Configuration block to define the attachment of the ENI. See [Attachment](#attachment) below for more details!
* `description` - (Optional) Description for the network interface.
* `private_ip_list` - (Optional) List of private IPs to assign to the ENI in sequential order. One value only.
* `private_ips` - (Optional) List of private IPs to assign to the ENI without regard to order. One value only.
* `security_groups` - (Optional) List of security group IDs to assign to the ENI.
* `source_dest_check` - (Optional) Whether to enable source destination checking for the ENI. Default true.
* `tags` - (Optional) Map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.

### Attachment

The `attachment` block supports the following:

* `instance` - (Required) ID of the instance to attach to.
* `device_index` - (Required) Integer to define the devices index.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - ARN of the network interface.
* `id` - ID of the network interface.
* `mac_address` - MAC address of the network interface.
* `owner_id` - The project ID.
* `private_dns_name` - Private DNS name of the network interface (IPv4).
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `interface_type` - Type of network interface to create. Always `"interface"`.
* `ipv4_prefix_count` - Number of IPv4 prefixes that AWS automatically assigns to the network interface. Always `0`.
* `ipv4_prefixes` - One or more IPv4 prefixes assigned to the network interface. Always empty.
* `ipv6_address_count` - Number of IPv6 addresses to assign to a network interface. Always `0`.
* `ipv6_address_list_enable` - Whether `ipv6_addreses_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_addresses_count` become read-only. Always `false`.
* `ipv6_address_list` - List of private IPs to assign to the ENI in sequential order. Always empty.
* `ipv6_addresses` - One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Always empty.
* `ipv6_prefix_count` - Number of IPv6 prefixes that AWS automatically assigns to the network interface. Always `0`.
* `ipv6_prefixes` - One or more IPv6 prefixes assigned to the network interface. Always empty.
* `private_ip_list_enable` - Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Always `false`.
* `private_ips_count` - Number of secondary private IPs to assign to the ENI. Always `0`.

## Import

Network Interfaces can be imported using the `id`, e.g.,

```
$ terraform import aws_network_interface.test eni-12345678
```

[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block

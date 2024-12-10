---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_instance"
description: |-
  Provides information about an EC2 instance.
---

[base64decode-function]: https://www.terraform.io/docs/configuration/functions/base64decode.html
[describe-instances]: https://docs.cloud.croc.ru/en/api/ec2/instances/DescribeInstances.html

# Data Source: aws_instance

Provides information about an EC2 instance. This data source can be used to get the ID of an EC2 instance for use in other resources.

## Example Usage

```terraform
data "aws_instance" "selected" {
  instance_id = "i-12345678"

  filter {
    name   = "image-id"
    values = ["cmi-12345678"]
  }

  instance_tags = {
    type = "test"
  }

  filter {
    name   = "tag:Name"
    values = ["example"]
  }
}
```

## Argument Reference

* `get_user_data` - (Optional) Retrieve Base64 encoded user data contents into the `user_data_base64` attribute.
  A SHA-1 hash of the user data contents will always be present in the `user_data` attribute. Defaults to `false`.
* `instance_id` - (Optional) Specify the exact instance ID with which to populate the data source.
* `instance_tags` - (Optional) A map of tags, each pair of which must exactly match a pair on the desired instance.

* `filter` - (Optional) One or more name/value pairs to use as filters.

For more information about filtering, see the [EC2 API documentation][describe-instances].

~> **NOTE:** At least one of `filter`, `instance_tags`, or `instance_id` must be specified.

~> **NOTE:** If anything other than a single match is returned by the search,
Terraform will fail. Ensure that your search is specific enough to return
a single instance ID only.

## Attributes Reference

`id` is set to the ID of the found instance. In addition, the following attributes
are exported:

~> **NOTE:** Some values are not always set and may not be available for
interpolation.

* `affinity` - The affinity setting for an instance on a dedicated host.
* `ami` - The ID of the image used to launch the instance.
* `arn` - The ARN of the instance.
* `associate_public_ip_address` - Whether the instance is associated with a public IP address or not.
* `availability_zone` - The availability zone of the instance.
* `ebs_block_device` - The EBS block device mappings of the instance.
    * `delete_on_termination` - If the EBS volume will be deleted on termination.
    * `device_name` - The physical name of the device.
    * `iops` - `0` If the EBS volume is not a provisioned IOPS image, otherwise the supported IOPS count.
    * `snapshot_id` - The ID of the snapshot.
    * `volume_size` - The size of the volume, in GiB.
    * `volume_type` - The volume type.
* `ephemeral_block_device` - The ephemeral block device mappings of the instance.
    * `device_name` - The physical name of the device.
    * `no_device` - Whether the specified device included in the device mapping was suppressed or not.
    * `virtual_name` - The virtual device name.
* `host_id` - The ID of the dedicated host that the instance will be assigned to.
* `instance_state` - The state of the instance. One of: `pending`, `running`, `shutting-down`, `terminated`, `stopping`, `stopped`.
* `instance_type` - The type of the instance.
* `key_name` - The key name of the instance.
* `monitoring` - Whether detailed monitoring is enabled or disabled for the instance.
* `network_interface_id` - The ID of the network interface that was created with the instance.
* `placement_group` - The placement group of the instance.
* `private_dns` - The private DNS name assigned to the instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.
* `private_ip` - The private IP address assigned to the instance.
* `secondary_private_ips` - The secondary private IPv4 addresses assigned to the instance's primary network interface in a VPC.
* `public_dns` - The public DNS name assigned to the instance. For EC2-VPC, this is only available if you've enabled DNS hostnames for your VPC.
* `public_ip` - The public IP address assigned to the instance, if applicable. **NOTE**: If you are using an [`aws_eip`](../resources/eip.md) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
* `root_block_device` - The root block device mappings of the instance
    * `device_name` - The physical name of the device.
    * `delete_on_termination` - If the root block device will be deleted on termination.
    * `iops` - `0` If the volume is not a provisioned IOPS image, otherwise the supported IOPS count.
    * `volume_size` - The size of the volume, in GiB.
    * `volume_type` - The type of the volume.
* `security_groups` - The associated security groups.
* `source_dest_check` - Whether the network interface performs source/destination checking.
* `subnet_id` - The VPC subnet ID.
* `user_data` - SHA-1 hash of user data supplied to the instance.
* `user_data_base64` - Base64 encoded contents of User Data supplied to the instance. Valid UTF-8 contents can be decoded with the [`base64decode` function][base64decode-function]. This attribute is only exported if `get_user_data` is true.
* `tags` - A map of tags assigned to the instance.
* `tenancy` - The placement type.
* `vpc_security_group_ids` - The associated security groups in a non-default VPC.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `credit_specification` - Configuration block for customizing the credit specification of the instance. Always empty.
* `ebs_block_device`:
    * `encrypted` - Whether to enable volume encryption. Always `false`.
    * `kms_key_id` - The ARN of the KMS Key to use when encrypting the volume. Always `""`.
    * `throughput` - Throughput to provision for a volume in mebibytes per second (MiB/s). Always `0`.
* `ebs_optimized` - If true, the launched EC2 instance will be EBS-optimized. Always `false`.
* `enclave_options` - Enable Nitro Enclaves on launched instances. Always empty.
* `get_password_data` - (Optional) If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
* `iam_instance_profile` - IAM Instance Profile to launch the instance with. Always `""`.
* `ipv6_addresses` - Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface. Always empty.
* `maintenance_options` - The maintenance and recovery options for the instance. Always empty.
    * `auto_recovery` - The automatic recovery behavior of the Instance.
* `metadata_options` - Customize the metadata options of the instance. Always empty.
    * `http_endpoint` - Whether the metadata service is available.
    * `http_tokens` - Whether the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_.
    * `http_put_response_hop_limit` - Desired HTTP PUT response hop limit for instance metadata requests.
    * `instance_metadata_tags` - Enables or disables access to instance tags from the instance metadata service.
* `network_interface`:
    * `network_card_index` - Integer index of the network card. Limited by instance type. Always `0`.
* `outpost_arn` - The ARN of the Outpost the instance is assigned to. Always `""`.
* `password_data` - Base-64 encoded encrypted password data for the instance. Always `""`.
* `placement_partition_number` - The number of the partition the instance is in. Always empty.
* `root_block_device`:
    * `encrypted` - Whether to enable volume encryption. Always `false`.
    * `kms_key_id` - The ARN of the KMS Key to use when encrypting the volume. Always `""`.
    * `throughput` - Throughput to provision for a volume in mebibytes per second (MiB/s). Always `0`.

---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_launch_template"
description: |-
  Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.
---

[asg-create]: https://docs.cloud.croc.ru/en/services/autoscaling.html#createautoscalinggroup
[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block
[describe-images]: https://docs.cloud.croc.ru/en/api/ec2/images/DescribeImages.html

# Resource: aws_launch_template

Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.

## Example Usage

```terraform
resource "aws_launch_template" "example" {
  name = "tf-lt"

  block_device_mappings {
    device_name = "disk1"

    ebs {
      volume_size = 20
    }
  }

  disable_api_termination = true

  instance_initiated_shutdown_behavior = "terminate"

  image_id      = "cmi-12345678" # add image id, change instance type if needed
  instance_type = "m1.micro"

  monitoring {
    enabled = true
  }

  placement {
    availability_zone = "ru-msk-vol52"
  }

  tag_specifications {
    resource_type = "instance"

    tags = {
      Name = "tf-lt"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `block_device_mappings` - (Optional) Specify volumes to attach to the instance besides the volumes specified by the image.
  See [Block Devices](#block-devices) below for details.
* `default_version` - (Optional, Conflicts with `update_default_version`) Default version of the launch template.
* `description` - (Optional) Description of the launch template version.
* `disable_api_termination` - (Optional) If `true`, disables the possibility to terminate an instance via API.
* `image_id` - (Required) The ID of the image from which to launch the instance.
* `instance_initiated_shutdown_behavior` - (Optional) Shutdown behavior for the instance. Valid values are `stop`, `terminate`.
* `instance_type` - (Optional) The type of the instance.
* `key_name` - (Optional) The key name to use for the instance.
* `monitoring` - (Optional) The monitoring option for the instance. See [Monitoring](#monitoring) below for more details.
* `name` - (Optional, Conflicts with `name_prefix`) The name of the launch template. If you leave this blank, Terraform will auto-generate a unique name.
* `name_prefix` - (Optional, Conflicts with `name`) Creates a unique name beginning with the specified prefix.
* `network_interfaces` - (Optional) Customize network interfaces to be attached at instance boot time.
  See [Network Interfaces](#network-interfaces) below for more details.
* `placement` - (Optional) The placement of the instance. See [Placement](#placement) below for more details.
* `tag_specifications` - (Optional) The tags to apply to the resources during launch. See [Tag Specifications](#tag-specifications) below for more details.
* `tags` - (Optional) A map of tags to assign to the launch template. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.
* `update_default_version` - (Optional, Conflicts with `default_version`) Whether to update default version each update.
* `user_data` - (Optional) The base64-encoded user data to provide when launching the instance. The text length must not exceed 16 KB.
* `vpc_security_group_ids` - (Optional) A list of security group IDs to associate with.

### Block devices

Configure additional volumes of the instance besides specified by the image.

To find out more information for an existing image to override the configuration, such as `device_name`, use the [EC2 API][describe-images].

Each `block_device_mappings` supports the following:

* `device_name` - (Optional) The name of the device to mount.
* `ebs` - (Optional) Configure EBS volume properties.
* `no_device` - (Optional) Suppresses the specified device included in the block device mapping.

The `ebs` block supports the following:

* `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination.
* `iops` - (Optional) The amount of provisioned IOPS. This must be set with a volume_type of `io2`.
* `snapshot_id` - (Optional) The snapshot ID to mount.
* `volume_size` - (Optional) The size of the volume in gigabytes.
* `volume_type` - (Optional) Type of volume. Valid values are `st2`, `gp2`, `io2`.

### Monitoring

The `monitoring` block supports the following:

* `enabled` - If `true`, the launched EC2 instance will have detailed monitoring enabled.

### Network Interfaces

Attaches one or more network interfaces to the instance.

For the details about configuring network interfaces when creating an auto scaling group, see the [user documentation][asg-create].

Each `network_interfaces` block supports the following:

* `associate_public_ip_address` - (Optional) Whether a public IP address should be associated with the network interface.
  The address will be assigned to the `eth0` interface if there are free allocated external addresses.
  This operation is available only for instances running in the VPC and for new network interfaces.
* `delete_on_termination` - (Optional) Whether the network interface should be destroyed on instance termination.
* `description` - (Optional) Description of the network interface.
* `device_index` - (Optional) The integer index of the network interface attachment.
* `network_interface_id` - (Optional) The ID of the network interface to attach.
* `private_ip_address` - (Optional) The primary private IPv4 address.
* `security_groups` - (Optional) A list of security group IDs to associate.
* `subnet_id` - (Optional) The VPC subnet ID to associate.

### Placement

The placement group of the instance.

The `placement` block supports the following:

* `affinity` - (Optional) The affinity setting for an instance on a dedicated host. The parameter could be set to `host` only if `tenancy` is `host`.
* `availability_zone` - (Optional) The availability zone for the instance.
* `group_name` - (Optional) The name of the placement group for the instance.
* `host_id` - (Optional) The ID of the dedicated host for the instance.
* `tenancy` - (Optional) The tenancy of the instance (if the instance is running in a VPC). Valid values are `default`, `host`. Defaults to `default`.

~> **NOTE:** If you use the `host` value, you may encounter the `NotEnoughResourcesForInstanceType` error when running an instance. To avoid this, it is recommended to provide either the `subnet_id` argument or the `availability_zone` argument.

### Tag Specifications

The tags to apply to the resources during launch. You can tag instances and volumes.

Each `tag_specifications` block supports the following:

* `resource_type` - (Optional) The type of resource to tag. Valid values are `instance`, `volume`.
* `tags` - (Optional) A map of tags to assign to the resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - The ARN of the launch template.
* `id` - The ID of the launch template.
* `latest_version` - The latest version of the launch template.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `block_device_mappings`:
    * `ebs`:
        * `encrypted` - Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume. Always `""`.
        * `kms_key_id` - The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use when creating the encrypted volume. Always `""`.
        * `throughput` - The throughput to provision for a `gp3` volume in MiB/s (specified as an integer, e.g., 500), with a maximum of 1,000 MiB/s. Always `0`.
    * `virtual_name` - The [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames). Always `""`.
* `capacity_reservation_specification` - Targeting for EC2 capacity reservations. Always empty.
    * `capacity_reservation_preference` - Indicates the instance's Capacity Reservation preferences.
    * `capacity_reservation_target` - Used to target a specific Capacity Reservation.
        * `capacity_reservation_id` - The ID of the Capacity Reservation in which to run the instance.
        * `capacity_reservation_resource_group_arn` - The ARN of the Capacity Reservation resource group in which to run the instance.
* `cpu_options` - The CPU options for the instance. Always empty.
    * `core_count` - The number of CPU cores for the instance.
    * `threads_per_core` - The number of threads per CPU core.
* `credit_specification` - Customize the credit specification of the instance. Always empty.
    * `cpu_credits` - The credit option for CPU usage.
* `ebs_optimized` - If `true`, the launched EC2 instance will be EBS-optimized.
* `elastic_gpu_specifications` - The elastic GPU to attach to the instance. Always empty.
    * `type` - The [Elastic GPU Type](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-gpus.html#elastic-gpus-basics)
* `elastic_inference_accelerator` - Configuration block containing an Elastic Inference Accelerator to attach to the instance. Always empty.
    * `type` - Accelerator type.
* `enclave_options` - Enable Nitro Enclaves on launched instances. Always empty.
    * `enabled` - If set to `true`, Nitro Enclaves will be enabled on the instance.
* `hibernation_options` - The hibernation options for the instance. Always empty.
    * `configured` - If set to `true`, the launched EC2 instance will hibernation enabled.
* `iam_instance_profile` - The IAM Instance Profile to launch the instance with. Always empty.
    * `arn` - The ARN of the instance profile.
    * `name` - The name of the instance profile.
* `instance_market_options` - The market (purchasing) option for the instance. Always empty.
    * `market_type` - The market type. Can be `spot`.
    * `spot_options` - The options for [Spot Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
* `instance_requirements` - The attribute requirements for the type of instance. Always empty.
    * `accelerator_count` - Block describing the minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips).
    * `accelerator_manufacturers` - List of accelerator manufacturer names.
    * `accelerator_names` - List of accelerator names.
    * `accelerator_total_memory_mib` - Block describing the minimum and maximum total memory of the accelerators.
    * `accelerator_types` - List of accelerator types.
    * `bare_metal` - Indicate whether bare metal instace types should be `included`, `excluded`, or `required`.
    * `baseline_ebs_bandwidth_mbps` - Block describing the minimum and maximum baseline EBS bandwidth, in Mbps.
    * `burstable_performance` - Indicate whether burstable performance instance types should be `included`, `excluded`, or `required`.
    * `cpu_manufacturers` List of CPU manufacturer names.
    * `excluded_instance_types` - List of instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (\*).
    * `instance_generations` - List of instance generation names.
    * `local_storage` - Indicate whether instance types with local storage volumes are `included`, `excluded`, or `required`.
    * `local_storage_types` - List of local storage type names. Default any storage type.
    * `memory_gib_per_vcpu` - Block describing the minimum and maximum amount of memory (GiB) per vCPU.
    * `memory_mib` - Block describing the minimum and maximum amount of memory (MiB).
    * `network_interface_count` - Block describing the minimum and maximum number of network interfaces.
    * `on_demand_max_price_percentage_over_lowest_price` - The price protection threshold for On-Demand Instances.
    * `require_hibernate_support` - Indicate whether instance types must support On-Demand Instance Hibernation, either `true` or `false`.
    * `spot_max_price_percentage_over_lowest_price` - The price protection threshold for Spot Instances.
    * `total_local_storage_gb` - Block describing the minimum and maximum total local storage (GB).
    * `vcpu_count` - Block describing the minimum and maximum number of vCPUs.
* `kernel_id` - The kernel ID. Always `""`.
* `license_specification` - A list of license specifications to associate with. Always empty.
    * `license_configuration_arn` - The ARN of the license configuration.
* `maintenance_options` - The maintenance options for the instance. Always empty.
    * `auto_recovery` - Disables the automatic recovery behavior of your instance or sets it to default.
* `metadata_options` - Customize the metadata options for the instance. Always empty.
    * `http_endpoint` - Whether the metadata service is available.
    * `http_tokens` - Whether the metadata service requires session tokens, also referred to as _Instance Metadata Service Version 2 (IMDSv2)_.
    * `http_put_response_hop_limit` - The desired HTTP PUT response hop limit for instance metadata requests.
    * `http_protocol_ipv6` - Enables or disables the IPv6 endpoint for the instance metadata service.
    * `instance_metadata_tags` - Enables or disables access to instance tags from the instance metadata service.
* `network_interfaces`:
    * `associate_carrier_ip_address` - Associate a Carrier IP address with `eth0` for a new network interface. Always `""`.
    * `interface_type` - The type of network interface. Always `""`.
    * `ipv4_address_count` - The number of secondary private IPv4 addresses to assign to a network interface. Always `0`.
    * `ipv4_addresses` - One or more private IPv4 addresses to associate. Always empty.
    * `ipv4_prefix_count` - The number of IPv4 prefixes to be automatically assigned to the network interface. Always `0`.
    * `ipv4_prefixes` - One or more IPv4 prefixes to be assigned to the network interface. Always empty.
    * `ipv6_address_count` - The number of IPv6 addresses to assign to a network interface. Always `0`.
    * `ipv6_addresses` - One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Always empty.
    * `ipv6_prefix_count` - The number of IPv6 prefixes to be automatically assigned to the network interface. Always `0`.
    * `ipv6_prefixes` - One or more IPv6 prefixes to be assigned to the network interface. Always empty.
    * `network_card_index` - The index of the network card. Some instance types support multiple network cards. Always `0`.
* `placement`:
    * `host_resource_group_arn` - The ARN of the Host Resource Group in which to launch instances. Always `""`.
    * `spread_domain` - Reserved for future use. Always `""`.
    * `partition_number` - The number of the partition the instance should launch in. Always `0`.
* `private_dns_name_options` - The options for the instance hostname. The default values are inherited from the subnet. Always empty.
    * `enable_resource_name_dns_aaaa_record` - Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
    * `enable_resource_name_dns_a_record` - Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
    * `hostname_type` - The type of hostname for Amazon EC2 instances.
* `ram_disk_id` - The ID of the RAM disk. Always `""`.
* `security_group_names` - A list of security group names to associate with. Always empty.

## Import

Launch Templates can be imported using the `id`, e.g.,

```
$ terraform import aws_launch_template.web lt-12345678
```

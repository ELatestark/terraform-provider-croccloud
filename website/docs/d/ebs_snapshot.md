---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "aws_ebs_snapshot"
description: |-
  Get information on an EBS snapshot.
---

# Data Source: aws_ebs_snapshot

Use this data source to get information about an EBS snapshot for use when provisioning EBS volumes

## Example Usage

```terraform
data "aws_ebs_snapshot" "ebs_snapshot" {
  most_recent = true
  owners      = ["self"]

  filter {
    name   = "volume-size"
    values = ["40"]
  }

  filter {
    name   = "tag:Name"
    values = ["Example"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `most_recent` - (Optional) If more than one result is returned, use the most recent snapshot.
* `owners` - (Optional) List of the snapshot owners. Valid items are the project ID (`project@customer`) or `self`.  
* `snapshot_ids` - (Optional) Returns information on a specific snapshot ID.
* `restorable_by_user_ids` - (Optional) List of the project IDs (`project@customer`).
  that can create volumes from the snapshot.
* `filter` - (Optional) One or more name/value pairs to filter.

For more information about filtering, see the [EC2 API documentation][describe-snapshots].

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the EBS snapshot.
* `id` - The snapshot ID (e.g., snap-12345678).
* `snapshot_id` - The snapshot ID (e.g., snap-12345678).
* `description` - A description for the snapshot
* `owner_id` - The project ID.
* `owner_alias` - The alias of the EBS snapshot owner.
* `volume_id` - The volume ID (e.g., vol-12345678).
* `volume_size` - The size of the drive in GiB.
* `state` - The snapshot state.
* `tags` - A map of tags for the resource.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `data_encryption_key_id` - The data encryption key identifier for the snapshot. Always `""`.
* `encrypted` - Whether the snapshot is encrypted. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always `""`.
* `outpost_arn` - The ARN of the Outpost on which the snapshot is stored. Always `""`.
* `storage_tier` - The storage tier in which the snapshot is stored. Always `""`.

[describe-snapshots]: https://docs.cloud.croc.ru/en/api/ec2/snapshots/DescribeSnapshots.html

---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "aws_ebs_snapshot"
description: |-
  Provides an elastic block storage snapshot resource.
---

# Resource: aws_ebs_snapshot

Creates a snapshot of an EBS volume.

## Example Usage

```terraform
resource "aws_ebs_volume" "example" {
  availability_zone = "ru-msk-vol52"
  size              = 40

  tags = {
    Name = "HelloWorld"
  }
}

resource "aws_ebs_snapshot" "example_snapshot" {
  volume_id = aws_ebs_volume.example.id

  tags = {
    Name = "HelloWorld_snap"
  }
}
```

## Argument Reference

The following arguments are supported:

* `volume_id` - (Required) The volume ID of which to make a snapshot.
* `description` - (Optional) A description of what the snapshot is.
* `tags` - (Optional) A map of tags to assign to the snapshot. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.

### Timeouts

`aws_ebs_snapshot` provides the following
[Timeouts](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts) configuration options:

- `create` - (Default `10 minutes`) Used for creating the ebs snapshot
- `delete` - (Default `10 minutes`) Used for deleting the ebs snapshot

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the EBS snapshot.
* `id` - The snapshot ID (e.g., snap-12345678).
* `owner_id` - The project ID.
* `owner_alias` - The alias of the EBS snapshot owner.
* `volume_size` - The size of the drive in GiB.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `data_encryption_key_id` - The data encryption key identifier for the snapshot. Always `""`.
* `encrypted` - Whether the snapshot is encrypted. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always `""`.
* `outpost_arn` - The ARN of the Outpost on which the snapshot is stored. Always `""`.
* `permanent_restore` - Indicates whether to permanently restore an archived snapshot. Always empty.
* `storage_tier` - The storage tier in which the snapshot is stored. Always `""`.
* `temporary_restore_days` - The number of days for which to temporarily restore an archived snapshot. Always empty.

## Import

EBS Snapshot can be imported using the `id`, e.g.,

```
$ terraform import aws_ebs_snapshot.id snap-12345678
```

[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block

---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "aws_ebs_snapshot_import"
description: |-
  Provides an elastic block storage snapshot import resource.
---

# Resource: aws_ebs_snapshot_import

Imports a disk image from S3 as a snapshot.

## Example Usage

```terraform
resource "aws_ebs_snapshot_import" "example" {
  disk_container {
    format = "VHD"
    user_bucket {
      s3_bucket = "disk-images"
      s3_key    = "source.vhd"
    }
  }

  tags = {
    Name = "HelloWorld"
  }
}
```

## Argument Reference


The following arguments are supported:

* `description` - (Optional) The description string for the import snapshot task.
* `disk_container` - (Required) Information about the disk container. Detailed below.
* `tags` - (Optional) A map of tags to assign to the snapshot.

### disk_container Configuration Block

* `description` - (Optional) The description of the disk image being imported.
* `format` - (Required) The format of the disk image being imported. One of `VHD`, `VMDK` or `RAW`.
* `user_bucket` - (Required) The S3 bucket for the disk image.

### user_bucket Configuration Block

* `s3_bucket` - The name of the S3 bucket where the disk image is located.
* `s3_key` - The file name of the disk image.

### Timeouts

`aws_ebs_snapshot_import` provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - (Default `60 minutes`) Used for importing the EBS snapshot
- `delete` - (Default `10 minutes`) Used for deleting the EBS snapshot

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the EBS snapshot.
* `id` - The snapshot ID (e.g., snap-12345678).
* `owner_id` - The project ID.
* `owner_alias` - The alias of the EBS snapshot owner.
* `volume_size` - The size of the drive in GiB.
* `tags_all` - A map of tags assigned to the resource.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `client_data` - The client-specific data. Always empty.
* `data_encryption_key_id` - The data encryption key identifier for the snapshot. Always `""`.
* `disk_container`
    * `url` - The URL to the Amazon S3-based disk image being imported. Always `""`.
* `encrypted` - Whether the snapshot is encrypted. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always `""`.
* `outpost_arn` - The ARN of the Outpost on which the snapshot is stored. Always `""`.
* `permanent_restore` - Indicates whether to permanently restore an archived snapshot. Always empty.
* `role_name` - The name of the IAM Role the VM Import/Export service will assume. Always `vmimport`.
* `storage_tier` - The storage tier in which the snapshot is stored. Always `""`.
* `temporary_restore_days` - The number of days for which to temporarily restore an archived snapshot. Always empty.

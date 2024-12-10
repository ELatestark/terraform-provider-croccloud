---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_ami_launch_permission"
description: |-
  Adds a launch permission to an Amazon Machine Image (AMI).
---

# Resource: aws_ami_launch_permission

Adds a launch permission to an image.

## Example Usage

### AWS Account ID

```terraform
resource "aws_ami_launch_permission" "example" {
  image_id   = "cmi-12345678"
  account_id = "123456789012"
}
```

### Public Access

```terraform
# The cloud currently restricts adding public access permissions to images.
# Applying the resource must throw an error.
resource "aws_ami_launch_permission" "example" {
  image_id = "cmi-12345678"
  group    = "all"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) The project ID (`project@customer`) for the launch permission.
* `group` - (Optional) The name of the group for the launch permission. Valid values: `"all"`.
* `image_id` - (Required) The ID of the image.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Launch permission ID.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `organization_arn` - The ARN of an organization for the launch permission. Always `""`.
* `organizational_unit_arn` - The ARN of an organizational unit for the launch permission. Always `""`.

## Import

-> **Unsupported operation**
Import image launch permission is currently unsupported.

Image launch permissions can be imported using `[ACCOUNT-ID|GROUP-NAME]/IMAGE-ID`, e.g.,

```sh
$ terraform import aws_ami_launch_permission.example 123456789012/cmi-12345678
```

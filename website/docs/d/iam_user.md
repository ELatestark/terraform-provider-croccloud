---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_user"
description: |-
  Provides information about an IAM user.
---

[RFC3339 format]: https://datatracker.ietf.org/doc/html/rfc3339#section-5.8

# Data Source: aws_iam_user

Provides information about an IAM user.

## Example Usage

```terraform
data "aws_iam_user" "selected" {
  name = "user-name"
}
```

## Argument Reference

* `name` - (Required) The name of the user.

## Attribute Reference

* `arn` - The Amazon Resource Name (ARN) of the user.
* `display_name` - The displayed name of the user.
* `email` - The email of the user.
* `enabled` - Indicates whether the user is **not** locked.
* `id` - The name of the user.
* `last_login_date` - The time in [RFC3339 format] when the user last logged in to the web interface.
* `login` - The login of the user.
* `otp_required` -  Indicates whether the user is required to use two-factor authentication to log in to the web interface.
* `phone` - The phone number of the user.
* `update_date` - The time in [RFC3339 format] when the user was last updated.
* `user_id` - The ID of the user.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path` - The path to the user. Always `""`.
* `permissions_boundary` - The ARN of the policy that is used to limit permissions for the user. Always `""`.
* `tags` - Map of tags assigned to the user. Always empty.

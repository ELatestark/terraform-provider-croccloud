---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_users"
description: |-
  Provides ARNs and names of selected IAM users.
---

# Data Source: aws_iam_users

Provides ARNs (Amazon Resource Names) and the names of selected IAM users.

## Example Usage

### All Users

```terraform
data "aws_iam_users" "selected" {}
```

### Users Filtered By Name Regex

```terraform
data "aws_iam_users" "selected" {
  name_regex = "user.*"
}
```

## Argument Reference

The following arguments are supported:

* `name_regex` - (Optional) A regex string to apply to the list of users returned by IAM API.

~> This filtering is performed locally and can affect performance when the list of users is large.

## Attribute Reference

* `arns` - List of ARNs of the users.
* `id` - The region (e.g., `region-1`).
* `names` - List of user names.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path_prefix` - The path prefix to filter results. Always empty.

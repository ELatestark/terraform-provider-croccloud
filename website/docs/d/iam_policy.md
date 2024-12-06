---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "aws_iam_policy"
description: |-
  Provides information about an IAM policy.
---

[RFC3339 format]: https://datatracker.ietf.org/doc/html/rfc3339#section-5.8

# Data Source: aws_iam_policy

Provides information about an IAM policy.

## Example Usage

### By ARN

```terraform
data "aws_iam_policy" "selected" {
  arn = "arn:c2:iam:::policy/EC2ReadOnlyAccess"
}
```

### By Name

```terraform
data "aws_iam_policy" "selected" {
  name = "EC2ReadOnlyAccess"
}
```

## Argument Reference

* `arn` - (Required if `name` is not specified) The Amazon Resource Name (ARN) of the policy
  (e.g. `arn:c2:iam::<customer-name>:policy/<policy-name>`).
* `name` - (Required if `arn` is not specified) The name of the policy.

~> Filtering by `name` is performed locally and can affect performance when the list of policies is large.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `create_date` - The time in [RFC3339 format] when the policy was created.
* `description` - The description of the policy.
* `id` - The ARN of the policy.
* `owner` - The owner of the policy.
* `policy` - Policy-defined access rules in JSON format.
* `policy_id` - The ID of the policy.
* `type` - The type of the policy.
* `update_date` - The time in [RFC3339 format] when the policy was last updated.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `path` - The path to the policy. Always `""`.
* `path_prefix` - The prefix of the path to the IAM policy. Always empty.
* `tags` - Key-value mapping of tags for the IAM policy. Always empty.

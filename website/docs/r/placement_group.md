---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_placement_group"
description: |-
  Provides an EC2 placement group.
---

# Resource: aws_placement_group

Provides an EC2 placement group.

For more information, see the documentation on [Placement groups][placement-groups].

[placement-groups]: https://docs.cloud.croc.ru/en/services/compute/placementgroups.html

## Example Usage

```terraform
resource "aws_placement_group" "example" {
  name     = "test-pg"
  strategy = "cluster"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the placement group.
* `strategy` - (Required) The placement strategy. Possible values: `"spread"`.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the placement group.
* `id` - The name of the placement group.
* `placement_group_id` - The ID of the placement group.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `partition_count` - The number of partitions to create in the placement group. Always `0`.

## Import

Placement groups can be imported using the `name`, e.g.,

```
$ terraform import aws_placement_group.prod_pg production-placement-group
```

[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block

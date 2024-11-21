---
subcategory: "Auto Scaling"
layout: "aws"
page_title: "aws_autoscaling_group"
description: |-
  Get information on an Auto Scaling Group.
---

# Data Source: aws_autoscaling_group

Use this data source to get information on an existing Auto Scaling Group.

## Example Usage

```terraform
data "aws_autoscaling_group" "example" {
  name = "example-asg"
}
```

## Argument Reference

* `name` - Specify the exact name of the desired Auto Scaling Group.

## Attributes Reference

~> **NOTE:** Some values are not always set and may not be available for interpolation.

* `arn` - The Amazon Resource Name (ARN) of the Auto Scaling Group.
* `availability_zones` - One or more Availability Zones for the group.
* `default_cool_down` - The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
* `desired_capacity` - The desired size of the group.
* `health_check_grace_period` - The amount of time, in seconds, after which Auto Scaling Group can perform a health check on its instances.
* `id` - Name of the Auto Scaling Group.
* `max_size` - The maximum size of the group.
* `min_size` - The minimum size of the group.
* `name` - Name of the Auto Scaling Group.
* `new_instances_protected_from_scale_in` - Indicates whether new instances are protected from deletion when Auto Scaling Group is scaled in.
* `status` -  The status of the Auto Scaling Group when it is deleted.
* `vpc_zone_identifier` - The IDs of the subnets in which instances are created.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `enabled_metrics` - The list of metrics enabled for collection. Always empty.
* `health_check_type` - The service to use for the health checks. Always `""`.
* `launch_configuration` - The name of the associated launch configuration. Always `""`.
* `load_balancers` - One or more load balancers associated with the group. Always empty.
* `placement_group` - The name of the placement group into which to launch your instances, if any. Always `""`.
* `service_linked_role_arn` - The ARN of the service-linked role that the Auto Scaling Group uses to call other services on your behalf. Always `""`.
* `target_group_arns` - The ARN of the target groups for your load balancer. Always empty.
* `termination_policies` - The termination policies for the group. Always empty.

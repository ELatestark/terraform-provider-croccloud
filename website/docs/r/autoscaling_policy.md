---
subcategory: "Auto Scaling"
layout: "aws"
page_title: "aws_autoscaling_policy"
description: |-
  Provides an Auto Scaling Group resource.
---

# Resource: aws_autoscaling_policy

Provides an Auto Scaling Policy resource.

~> **NOTE:** You may want to omit `desired_capacity` attribute from attached `aws_autoscaling_group`
when using autoscaling policies. It's good practice to pick either manual or dynamic (policy-based) scaling.

## Example Usage

```terraform
resource "aws_autoscaling_policy" "example" {
  name                   = "terraform-test"
  scaling_adjustment     = 4
  adjustment_type        = "ChangeInCapacity"
  cooldown               = 300
  autoscaling_group_name = "example-asg" # asg is created manually
}
```

## Argument Reference

* `name` - (Required) The name of the policy.
* `autoscaling_group_name` - (Required) The name of the autoscaling group.
* `adjustment_type` - (Optional) Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
* `policy_type` - (Optional) The policy type. Valid value is `SimpleScaling`.
* `min_adjustment_magnitude` - (Optional) Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.
* `cooldown` - (Optional) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
* `scaling_adjustment` - (Optional) The amount by which the Auto Scaling Group is scaled when the scaling policy is executed

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the scaling policy.
* `id` - The scaling policy's name.
* `name` - The scaling policy's name.
* `autoscaling_group_name` - The scaling policy's assigned autoscaling group.
* `adjustment_type` - The scaling policy's adjustment type.
* `policy_type` - The scaling policy's type.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `estimated_instance_warmup` - The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Always `0`.
* `metric_aggregation_type` - The aggregation type for the policy's metrics. Always `""`.
* `predictive_scaling_configuration` - The predictive scaling policy configuration. Always empty.
* `step_adjustment` - A set of adjustments that manage group scaling. Always empty.
* `target_tracking_configuration` - The target tracking policy configuration. Always empty.

## Import

AutoScaling scaling policy can be imported using the `autoscaling_group_name` and `name` separated by `/`.

```
$ terraform import aws_autoscaling_policy.test-policy asg-name/policy-name
```

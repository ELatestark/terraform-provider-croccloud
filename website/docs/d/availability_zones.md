---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_availability_zones"
description: |-
    Provides a list of availability zones.
---

# Data Source: aws_availability_zones

Provides a list of availability zones.

This is different from the [`aws_availability_zone`][tf-availability-zone] (singular) data source,
which provides some details about a specific availability zone.

[tf-availability-zone]: availability_zone.html

## Example Usage

### By State

```terraform
data "aws_availability_zones" "available" {
  state = "available"
}
```


## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.
* `state` - (Optional) Allows to filter list of availability zones based on their
current state. Can be either `"available"`, `"information"`, `"impaired"` or
`"unavailable"`.

### filter Configuration Block

The following arguments are supported by the `filter` configuration block:

* `name` - (Required) The name of the filter field.
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

For more information about filtering, see the [EC2 API documentation][describe-azs].

[describe-azs]: https://docs.cloud.croc.ru/en/api/ec2/placements/DescribeAvailabilityZones.html

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Region of the availability zones.
* `names` - A list of the availability zone names available to the account.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `all_availability_zones` - Whether all availability zones and local zones are included regardless of your opt in status. Always empty.
* `exclude_names` - List of availability zone names to exclude. Always empty.
* `exclude_zone_ids` - List of availability zone IDs to exclude. Always empty.
* `group_names` A set of the availability zone Group names. Always empty.
* `zone_ids` - A list of the availability zone IDs available to the account. Always empty.

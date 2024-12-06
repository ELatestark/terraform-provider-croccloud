---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "aws_availability_zone"
description: |-
    Provides details about a specific availability zone
---

# Data Source: aws_availability_zone

`aws_availability_zone` provides details about a specific availability zone (AZ).

This is different from the [`aws_availability_zones`][tf-availability-zones] (plural) data source,
which provides a list of the available zones.

[tf-availability-zones]: availability_zones.html

## Example Usage

```terraform
data "aws_availability_zone" "example" {
  name = "ru-msk-vol52"
}

output "availability_zone_to_region" {
  value = data.aws_availability_zone.example.id
}
```

## Argument Reference

The arguments of this data source act as filters for querying the available
availability zones. The given filters must match exactly one availability
zone whose data will be exported as attributes.

* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.
* `name` - (Optional) The full name of the availability zone to select.
* `state` - (Optional) A specific availability zone state to require. May be any of `"available"`, `"information"` or `"impaired"`.

### filter Configuration Block

The following arguments are supported by the `filter` configuration block:

* `name` - (Required) The name of the filter field.
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

For more information about filtering, see the [EC2 API documentation][describe-azs].

[describe-azs]: https://docs.cloud.croc.ru/en/api/ec2/placements/DescribeAvailabilityZones.html

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `region` - The region where the selected availability zone resides.
* `state` - A specific availability zone state to require. Possible values: `"available"`, `"information"`, `"impaired"`, `"unavailable"`.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `all_availability_zones` - Whether all availability zones and local zones are included regardless of your opt in status. Always empty.
* `group_name` - Group name. Always `""`.
* `name_suffix` - The part of the AZ name that appears after the region name, uniquely identifying the AZ within its region. Always `""`.
* `network_border_group` - The name of the location from which the address is advertised. Always `""`.
* `opt_in_status` - For Availability Zones, this always has the value of `opt-in-not-required`. Always `""`.
* `parent_zone_id` - The ID of the zone that handles some Local Zone or Wavelength Zone control plane operations, such as API calls. Always `""`.
* `parent_zone_name` - The name of the zone that handles some Local Zone or Wavelength Zone control plane operations, such as API calls.  Always `""`.
* `zone_id` - The zone ID of the availability zone to select.  Always `""`.
* `zone_type` - The type of zone. Values are `availability-zone`, `local-zone`, and `wavelength-zone`. Always `""`.

---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "aws_s3_bucket"
description: |-
    Provides details about a specific S3 bucket.
---

# Data Source: aws_s3_bucket

Provides details about a specific S3 bucket.

## Example Usage

```terraform
data "aws_s3_bucket" "selected" {
  bucket = "tf-example"
}

output "bucket_arn" {
  value = data.aws_s3_bucket.selected.arn
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The name of the bucket.
* `arn` - The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
* `region` - The region this bucket resides in.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `bucket_domain_name` - The bucket domain name. Contains domain name of format `bucketname.s3.amazonaws.com`.
* `bucket_regional_domain_name` - The bucket region-specific domain name. Contains domain name based on AWS region.
* `hosted_zone_id` - The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region. Contains Zone ID based on AWS region.
* `website_domain` - The domain of the website endpoint. Contains domain based on AWS region if the bucket is configured with a website or `""`.
* `website_endpoint` - The website endpoint. Contains endpoint based on AWS region if the bucket is configured with a website or `""`.

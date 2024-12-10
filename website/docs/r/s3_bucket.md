---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "aws_s3_bucket"
description: |-
  Provides a S3 bucket resource.
---

[bucket-naming]: https://docs.cloud.croc.ru/en/services/object_storage/operations.html#s3bucketnaming
[canned-acl]: https://docs.cloud.croc.ru/en/api/s3/acl.html#cannedacl
[cors]: https://docs.cloud.croc.ru/en/services/object_storage/operations.html#cors
[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block
[hosting-website]: https://docs.cloud.croc.ru/en/services/object_storage/operations.html#objectstoragestaticwebsitesmanual
[lifecycle-management]: https://docs.cloud.croc.ru/en/services/object_storage/operations.html#id24
[s3-versioning]: https://docs.cloud.croc.ru/en/services/object_storage/operations.html#s3versioningmanual
[website-redirect-rules]: https://docs.cloud.croc.ru/en/services/object_storage/instructions.html#s3setredirectiontowebsite

# Resource: aws_s3_bucket

Provides a S3 bucket resource.

~> **NOTE on S3 Bucket canned ACL Configuration:** S3 Bucket canned ACL can be configured in either the standalone resource [`aws_s3_bucket_acl`](s3_bucket_acl.md)
or with the deprecated parameter `acl` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket ACL Grants Configuration:** S3 Bucket grants can be configured in either the standalone resource [`aws_s3_bucket_acl`](s3_bucket_acl.md)
or with the deprecated parameter `grant` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket CORS Configuration:** S3 Bucket CORS can be configured in either the standalone resource [`aws_s3_bucket_cors_configuration`](s3_bucket_cors_configuration.md)
or with the deprecated parameter `cors_rule` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket Lifecycle Configuration:** S3 Bucket Lifecycle can be configured in either the standalone resource [`aws_s3_bucket_lifecycle_configuration`](s3_bucket_lifecycle_configuration.md)
or with the deprecated parameter `lifecycle_rule` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket Policy Configuration:** S3 Bucket Policy can be configured in either the standalone resource [`aws_s3_bucket_policy`](s3_bucket_policy.md)
or with the deprecated parameter `policy` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket Versioning Configuration:** S3 Bucket versioning can be configured in either the standalone resource [`aws_s3_bucket_versioning`](s3_bucket_versioning.md)
or with the deprecated parameter `versioning` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

~> **NOTE on S3 Bucket Website Configuration:** S3 Bucket Website can be configured in either the standalone resource [`aws_s3_bucket_website_configuration`](s3_bucket_website_configuration.md)
or with the deprecated parameter `website` in the resource `aws_s3_bucket`.
Configuring with both will cause inconsistencies and may overwrite configuration.

## Example Usage

### Private Bucket w/ Tags

```terraform
resource "aws_s3_bucket" "example" {
  bucket = "tf-example"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  tags = {
    Name        = "tf-example"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.example.id
  acl    = "private"
}
```

### Static Website Hosting

-> **NOTE:** The parameter `website` is deprecated.
Use the resource [`aws_s3_bucket_website_configuration`](s3_bucket_website_configuration.md) instead.

```terraform
resource "aws_s3_bucket" "example" {
  bucket = "tf-example"
  acl    = "public-read"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  website {
    index_document = "index.html"
    error_document = "error.html"

    routing_rules = <<EOF
[{
    "Condition": {
        "KeyPrefixEquals": "docs/"
    },
    "Redirect": {
        "ReplaceKeyPrefixWith": "documents/"
    }
}]
EOF
  }
}
```

### Using CORS

-> **NOTE:** The parameter `cors_rule` is deprecated.
Use the resource [`aws_s3_bucket_cors_configuration`](s3_bucket_cors_configuration.md) instead.

```terraform
resource "aws_s3_bucket" "example" {
  bucket = "tf-example"
  acl    = "public-read"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "POST"]
    allowed_origins = ["https://s3-website-test.hashicorp.com"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}
```

### Using versioning

-> **NOTE:** The parameter `versioning` is deprecated.
Use the resource [`aws_s3_bucket_versioning`](s3_bucket_versioning.md) instead.

```terraform
resource "aws_s3_bucket" "example" {
  bucket = "tf-example"
  acl    = "private"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  versioning {
    enabled = true
  }
}
```

### Using object lifecycle

-> **NOTE:** The parameter `lifecycle_rule` is deprecated.
Use the resource [`aws_s3_bucket_lifecycle_configuration`](s3_bucket_lifecycle_configuration.md) instead.

```terraform
resource "aws_s3_bucket" "bucket" {
  bucket = "tf-example"
  acl    = "private"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  lifecycle_rule {
    id      = "log"
    enabled = true

    prefix = "log/"

    tags = {
      rule      = "log"
      autoclean = "true"
    }

    expiration {
      days = 90
    }
  }

  lifecycle_rule {
    id      = "tmp"
    prefix  = "tmp/"
    enabled = true

    expiration {
      date = "2016-01-12"
    }
  }
}

resource "aws_s3_bucket" "versioning_bucket" {
  bucket = "tf-example"
  acl    = "private"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  versioning {
    enabled = true
  }

  lifecycle_rule {
    prefix  = "config/"
    enabled = true

    noncurrent_version_expiration {
      days = 90
    }
  }
}
```

### Using ACL policy grants

-> **NOTE:** The parameters `acl` and `grant` are deprecated.
Use the resource [`aws_s3_bucket_acl`](s3_bucket_acl.md) instead.

```terraform
data "aws_canonical_user_id" "current_user" {}

resource "aws_s3_bucket" "example" {
  bucket = "tf-example"

  # Use the predefined provider configuration to connect to object storage
  # https://docs.cloud.croc.ru/en/api/tools/terraform.html#providers-tf
  provider = aws.noregion

  grant {
    id          = data.aws_canonical_user_id.current_user.id
    type        = "CanonicalUser"
    permissions = ["FULL_CONTROL"]
  }

  grant {
    type        = "Group"
    permissions = ["READ_ACP", "WRITE"]
    uri         = "http://acs.amazonaws.com/groups/global/AuthenticatedUsers"
  }
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules may be found in [user documentation][bucket-naming].
* `bucket_prefix` - (Optional, Conflicts with `bucket`, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Must be lowercase and less than or equal to 37 characters in length. A full list of bucket naming rules may be found in [user documentation][bucket-naming].
* `acl` - (Optional, **Deprecated**, Conflicts with `grant`) The [canned ACL][canned-acl] to apply. Valid values are `private`, `public-read`, `public-read-write`, `authenticated-read`. Defaults to `private`. Terraform will only perform drift detection if a configuration value is provided. Use the resource [`aws_s3_bucket_acl`](s3_bucket_acl.md) instead.
* `grant` - (Optional, **Deprecated**, Conflicts with `acl`) An ACL policy grant. See [Grant](#grant) below for details. Terraform will only perform drift detection if a configuration value is provided. Use the resource [`aws_s3_bucket_acl`](s3_bucket_acl.md) instead.
* `cors_rule` - (Optional, **Deprecated**) A rule of [Cross-Origin Resource Sharing][cors]. See [CORS rule](#cors-rule) below for details. Terraform will only perform drift detection if a configuration value is provided. Use the resource [`aws_s3_bucket_cors_configuration`](s3_bucket_cors_configuration.md) instead.
* `force_destroy` - (Optional) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. Defaults to `false`.
* `lifecycle_rule` - (Optional, **Deprecated**) A configuration of [object lifecycle management][lifecycle-management]. See [Lifecycle Rule](#lifecycle-rule) below for details. Terraform will only perform drift detection if a configuration value is provided.
  Use the resource [`aws_s3_bucket_lifecycle_configuration`](s3_bucket_lifecycle_configuration.md) instead.
* `policy` - (Optional, **Deprecated**) A valid bucket policy JSON document. Note that if the policy document is not specific enough (but still valid), Terraform may view the policy as constantly changing in a `terraform plan`. In this case, please make sure you use the verbose/specific version of the policy.
  Terraform will only perform drift detection if a configuration value is provided.
  Use the resource [`aws_s3_bucket_policy`](s3_bucket_policy.md) instead.
* `versioning` - (Optional, **Deprecated**) A configuration of the [S3 bucket versioning state][s3-versioning]. See [Versioning](#versioning) below for details. Terraform will only perform drift detection if a configuration value is provided. Use the resource [`aws_s3_bucket_versioning`](s3_bucket_versioning.md) instead.
* `website` - (Optional, **Deprecated**) A configuration of the [S3 bucket website][hosting-website]. See [Website](#website) below for details. Terraform will only perform drift detection if a configuration value is provided.
  Use the resource [`aws_s3_bucket_website_configuration`](s3_bucket_website_configuration.md) instead.
* `tags` - (Optional) A map of tags to assign to the bucket. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level.

### CORS Rule

~> **NOTE:** Currently, changes to the `cors_rule` configuration of _existing_ resources cannot be automatically detected by Terraform. To manage changes of CORS rules to an S3 bucket, use the `aws_s3_bucket_cors_configuration` resource instead. If you use `cors_rule` on an `aws_s3_bucket`, Terraform will assume management over the full set of CORS rules for the S3 bucket, treating additional CORS rules as drift. For this reason, `cors_rule` cannot be mixed with the external `aws_s3_bucket_cors_configuration` resource for a given S3 bucket.

The `cors_rule` configuration block supports the following arguments:

* `allowed_headers` - (Optional) List of headers allowed.
* `allowed_methods` - (Required) One or more HTTP methods that you allow the origin to execute. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
* `allowed_origins` - (Required) One or more origins you want customers to be able to access the bucket from.
* `expose_headers` - (Optional) One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
* `max_age_seconds` - (Optional) Specifies time in seconds that browser can cache the response for a preflight request.

### Grant

~> **NOTE:** Currently, changes to the `grant` configuration of _existing_ resources cannot be automatically detected by Terraform. To manage changes of ACL grants to an S3 bucket, use the `aws_s3_bucket_acl` resource instead. If you use `grant` on an `aws_s3_bucket`, Terraform will assume management over the full set of ACL grants for the S3 bucket, treating additional ACL grants as drift. For this reason, `grant` cannot be mixed with the external `aws_s3_bucket_acl` resource for a given S3 bucket.

The `grant` configuration block supports the following arguments:

* `id` - (Optional) Canonical user ID to grant for (S3 User ID). Used only when `type` is `CanonicalUser`.
* `type` - (Required) Type of grantee to apply for. Valid values are `CanonicalUser` and `Group`. `AmazonCustomerByEmail` is not supported.
* `permissions` - (Required) List of permissions to apply for grantee. Valid values are `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`.
* `uri` - (Optional) Uri address to grant for. Supported groups are `http://acs.amazonaws.com/groups/global/AllUsers` and `http://acs.amazonaws.com/groups/global/AuthenticatedUsers`. Used only when `type` is `Group`.

### Lifecycle Rule

~> **NOTE:** Currently, changes to the `lifecycle_rule` configuration of _existing_ resources cannot be automatically detected by Terraform. To manage changes of Lifecycle rules to an S3 bucket, use the `aws_s3_bucket_lifecycle_configuration` resource instead. If you use `lifecycle_rule` on an `aws_s3_bucket`, Terraform will assume management over the full set of Lifecycle rules for the S3 bucket, treating additional Lifecycle rules as drift. For this reason, `lifecycle_rule` cannot be mixed with the external `aws_s3_bucket_lifecycle_configuration` resource for a given S3 bucket.

~> **NOTE:** At least one of `abort_incomplete_multipart_upload_days`, `expiration`, `noncurrent_version_expiration`, must be specified.

The `lifecycle_rule` configuration block supports the following arguments:

* `id` - (Optional) Unique identifier for the rule. Must be less than or equal to 255 characters in length.
* `tags` - (Optional) Specifies object tags key and value.
* `enabled` - (Required) Specifies lifecycle rule status.
* `expiration` - (Optional) Specifies a period in the object's expire. See [Expiration](#expiration) below for details.
* `noncurrent_version_expiration` - (Optional) Specifies when noncurrent object versions expire. See [Noncurrent Version Expiration](#noncurrent-version-expiration) below for details.

#### Expiration

The `expiration` configuration block supports the following arguments:

* `date` - (Optional) Specifies the date after which you want the corresponding action to take effect.
* `days` - (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
* `expired_object_delete_marker` - (Optional) On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct S3 to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.

#### Noncurrent Version Expiration

The `noncurrent_version_expiration` configuration block supports the following arguments:

* `days` - (Required) Specifies the number of days noncurrent object versions expire.

### Versioning

~> **NOTE:** Currently, changes to the `versioning` configuration of _existing_ resources cannot be automatically detected by Terraform. To manage changes of versioning state to an S3 bucket, use the `aws_s3_bucket_versioning` resource instead. If you use `versioning` on an `aws_s3_bucket`, Terraform will assume management over the versioning state of the S3 bucket, treating additional versioning state changes as drift. For this reason, `versioning` cannot be mixed with the external `aws_s3_bucket_versioning` resource for a given S3 bucket.

The `versioning` configuration block supports the following arguments:

* `enabled` - (Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.

### Website

~> **NOTE:** Currently, changes to the `website` configuration of _existing_ resources cannot be automatically detected by Terraform. To manage changes to the website configuration of an S3 bucket, use the `aws_s3_bucket_website_configuration` resource instead. If you use `website` on an `aws_s3_bucket`, Terraform will assume management over the configuration of the website of the S3 bucket, treating additional website configuration changes as drift. For this reason, `website` cannot be mixed with the external `aws_s3_bucket_website_configuration` resource for a given S3 bucket.

The `website` configuration block supports the following arguments:

* `index_document` - (Required, unless using `redirect_all_requests_to`) S3 returns this index document when requests are made to the root domain or any of the subfolders.
* `error_document` - (Optional) An absolute path to the document to return in case of a 4XX error.
* `redirect_all_requests_to` - (Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The default is the protocol that is used in the original request.
* `routing_rules` - (Optional) A json array containing [routing rules][website-redirect-rules] describing redirect behavior and when redirects are applied.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The name of the bucket.
* `arn` - The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
* `region` - The region this bucket resides in.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags].

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `acceleration_status` - Sets the accelerate configuration of an existing bucket. Always `""`.
* `bucket_domain_name` - The bucket domain name. Contains domain name of format `bucketname.s3.amazonaws.com`.
* `bucket_regional_domain_name` - The bucket region-specific domain name. Contains domain name based on AWS region.
* `hosted_zone_id` - The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region. Contains Zone ID based on AWS region.
* `lifecycle_rule`:
    * `abort_incomplete_multipart_upload_days` - Specifies the number of days after initiating a multipart upload when the multipart upload must be completed. Always `0`.
    * `prefix` - Prefix identifying one or more objects to which the rule applies. Always `""`.
    * `noncurrent_version_transition` - Set of configuration blocks that specify the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class. Always empty.
        * `newer_noncurrent_versions` - The number of noncurrent versions Amazon S3 will retain.
        * `noncurrent_days` - The number of days an object is noncurrent before Amazon S3 can perform the associated action.
        * `storage_class` - The class of storage used to store the object.
    * `transition` - Set of configuration blocks that specify when an Amazon S3 object transitions to a specified storage class. Always empty.
        * `date` - The date objects are transitioned to the specified storage class.
        * `days` - The number of days after creation when objects are transitioned to the specified storage class.
        * `storage_class` - The class of storage used to store the object.
* `logging` - A configuration of [S3 bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) parameters. Always empty.
    * `target_bucket` - The name of the bucket that will receive the log objects.
    * `target_prefix` - To specify a key prefix for log objects.
* `object_lock_configuration` - A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). Always empty.
    * `object_lock_enabled` - Indicates whether this bucket has an Object Lock configuration enabled.
    * `rule` - The Object Lock rule in place for this bucket.
        * `default_retention` - The default retention period that you want to apply to new objects placed in this bucket.
            * `mode` - The default Object Lock retention mode you want to apply to new objects placed in this bucket.
            * `days` - The number of days that you want to specify for the default retention period.
            * `years` - The number of years that you want to specify for the default retention period.
* `object_lock_enabled` - Indicates whether this bucket has an Object Lock configuration enabled. Always `false`.
* `replication_configuration` - A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html).
    * `role` - The ARN of the IAM role for Amazon S3 to assume when replicating the objects. Always `""`.
    * `rules` - Specifies the rules managing the replication. Always empty.
        * `delete_marker_replication_status` - Whether delete markers are replicated.
        * `destination` - Specifies the destination for the rule.
        * `filter` - Filter that identifies subset of objects to which the replication rule applies.
        * `id` - Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        * `prefix` - Object keyname prefix identifying one or more objects to which the rule applies.
        * `priority` - The priority associated with the rule.
        * `source_selection_criteria` - Specifies special object selection criteria .
        * `status` - The status of the rule.
* `request_payer` - Specifies who should bear the cost of Amazon S3 data transfer. Always `BucketOwner`.
* `server_side_encryption_configuration` - A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html).
    * `rule` - A single object for server-side encryption by default configuration. Always empty.
        * `apply_server_side_encryption_by_default` - A single object for setting server-side encryption by default. (documented below)
            * `sse_algorithm` - The server-side encryption algorithm to use. Valid values are `AES256` and `aws:kms`
            * `kms_master_key_id` - The AWS KMS master key ID used for the SSE-KMS encryption.
        * `bucket_key_enabled` - Whether to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
* `versioning` - Specifies who should bear the cost of Amazon S3 data transfer.
    * `mfa_delete` - Enable MFA delete for either `Change the versioning state of your bucket` or `Permanently delete an object version`. Always false.
* `website_domain` - The domain of the website endpoint. Contains domain based on AWS region if the bucket is configured with a website or `""`.
* `website_endpoint` - The website endpoint. Contains endpoint based on AWS region if the bucket is configured with a website or `""`.

## Import

S3 bucket can be imported using the `bucket`, e.g.,

```
$ terraform import aws_s3_bucket.bucket bucket-name
```

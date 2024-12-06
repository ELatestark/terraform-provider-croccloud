---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "aws_s3_object"
description: |-
    Provides metadata and optionally content of an S3 object
---

[ETag]: https://en.wikipedia.org/wiki/HTTP_ETag
[set-lifecycle]: https://docs.cloud.croc.ru/en/services/object_storage/instructions.html#s3setlifecycle

# Data Source: aws_s3_object

The S3 object data source allows access to the metadata and
_optionally_ (see below) content of an object stored inside S3 bucket.

~> **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.

## Example Usage

```terraform
data "aws_s3_object" "example" {
  bucket = "tf-example"
  key    = "test.txt"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket to read the object from.
* `key` - (Required) The full path to the object inside the bucket
* `version_id` - (Optional) Specific version ID of the object returned (defaults to latest version).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `body` - Object data (see **limitations above** to understand cases in which this field is actually available)
* `cache_control` - Specifies caching behavior along the request/reply chain.
* `content_disposition` - Specifies presentational information for the object.
* `content_encoding` - Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
* `content_language` - The language the content is in.
* `content_length` - Size of the body in bytes.
* `content_type` - A standard MIME type describing the format of the object data.
* `etag` - [ETag] generated for the object (an MD5 sum of the object content in case it's not encrypted)
* `expiration` - If the object expiration is configured (see [how to configure lifecycle][set-lifecycle]), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
* `id` - The full path to the object inside the bucket.
* `last_modified` - Last modified date of the object in RFC1123 format (e.g., `Mon, 02 Jan 2006 15:04:05 MST`)
* `metadata` - A map of metadata stored with the object in S3.
* `version_id` - The latest version ID of the object returned.
* `website_redirect_location` - If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. S3 stores the value of this header in the object metadata.
* `tags` - A map of tags assigned to the object.

-> **Note:** Terraform ignores all leading `/`s in the object's `key` and treats multiple `/`s in the rest of the object's `key` as a single `/`, so values of `/index.html` and `index.html` correspond to the same S3 object as do `first//second///third//` and `first/second/third/`.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `bucket_key_enabled` - Whether to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS. Always `false`.
* `expires` - The date and time at which the object is no longer cacheable. Always `""`.
* `force_destroy` - Whether to allow the object to be deleted by removing any legal hold on any object version. This value should be set to `true` only if the bucket has S3 object lock enabled. Always `false`.
* `kms_key_id` - The ARN for the KMS encryption key. Always empty.
* `object_lock_legal_hold_status` - [Legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Always `""`.
* `object_lock_mode` - Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Always `""`.
* `object_lock_retain_until_date` - Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods). Always `""`.
* `server_side_encryption` - Server-side encryption of the object in S3. Always `""`.
* `sse_kms_key_id` - If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object. Always `""`.
* `storage_class` - [Storage Class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html#AmazonS3-PutObject-request-header-StorageClass) for the object. Always `STANDARD`.


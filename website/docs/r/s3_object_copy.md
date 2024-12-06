---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "aws_s3_object_copy"
description: |-
  Provides a resource for copying an S3 object.
---

[canned-acl]: https://docs.cloud.croc.ru/en/api/s3/acl.html#cannedacl
[default-tags]: https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block
[RFC3339 format]: https://tools.ietf.org/html/rfc3339#section-5.8
[w3c cache_control]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
[w3c content_disposition]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1
[w3c content_encoding]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11

# Resource: aws_s3_object_copy

Provides a resource for copying an S3 object.

## Example Usage

```terraform
resource "aws_s3_object_copy" "example" {
  bucket = "tf-example"
  key    = "destination_key"
  source = "source_bucket/source_key"

  acl = "authenticated-read"
}
```

## Argument Reference

The following arguments are required:

* `bucket` - (Required) Name of the bucket to put the file in.
* `key` - (Required) Name of the object once it is in the bucket.
* `source` - (Required) Specifies the source object for the copy operation. The source object consists of the name of the source bucket and the key of the source object, separated by a slash (`/`). For example, `testbucket/test1.json`.

The following arguments are optional:

* `acl` - (Optional, Conflicts with `grant`) [Canned ACL][canned-acl] to apply. Valid values are `private`, `public-read`, `public-read-write`, `authenticated-read`. Defaults to `private`.
* `content_type` - (Optional) Standard MIME type describing the format of the object data, e.g., `application/octet-stream`. All Valid MIME Types are valid for this input.
* `grant` - (Optional, Conflicts with `acl`) Configuration block for header grants [documented below](#grant).

### grant

This configuration block has the following required arguments:

* `permissions` - (Required) List of permissions to grant to grantee. Valid values are `READ`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`.
* `type` - (Required) - Type of grantee. Valid values are `CanonicalUser`, `Group`, and `AmazonCustomerByEmail`.

This configuration block has the following optional arguments (one of the three is required):

* `email` - (Optional) Email address of the grantee (S3 Project email). Used only when `type` is `AmazonCustomerByEmail`.
* `id` - (Optional) The canonical user ID of the grantee (S3 User ID). Used only when `type` is `CanonicalUser`.
* `uri` - (Optional) URI of the grantee group. Supported groups are `http://acs.amazonaws.com/groups/global/AllUsers` and `http://acs.amazonaws.com/groups/global/AuthenticatedUsers`. Used only when `type` is `Group`.

-> **Note:** Terraform ignores all leading `/`s in the object's `key` and treats multiple `/`s in the rest of the object's `key` as a single `/`, so values of `/index.html` and `index.html` correspond to the same S3 object as do `first//second///third//` and `first/second/third/`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `etag` - ETag generated for the object (an MD5 sum of the object content). For plaintext objects the hash is an MD5 digest of the object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest, regardless of the method of encryption.
* `id` - The `key` of the resource supplied above.
* `last_modified` - Returns the date that the object was last modified, in [RFC3339 format].
* `version_id` - Version ID of the newly created copy.

->  **Unsupported attributes**
These exported attributes are currently unsupported:

* `cache_control` - Specifies caching behavior along the request/reply chain. Read [w3c cache_control] for further details. Always `""`.
* `content_disposition` - Specifies presentational information for the object. Read [w3c content_disposition] for further information. Always `""`.
* `content_encoding` - Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content_encoding] for further information. Always `""`.
* `content_language` - Language the content is in e.g., en-US or en-GB. Always `""`.
* `copy_if_match` - Copies the object if its entity tag (ETag) matches the specified tag. Always empty.
* `copy_if_modified_since` - Copies the object if it has been modified since the specified time, in [RFC3339 format]. Always empty.
* `copy_if_none_match` - Copies the object if its entity tag (ETag) is different from the specified ETag. Always empty.
* `copy_if_unmodified_since` - Copies the object if it hasn't been modified since the specified time, in [RFC3339 format]. Always empty.
* `customer_algorithm` - Specifies the algorithm to use to when encrypting the object (for example, AES256). Always `""`.
* `customer_key` - Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.  Always `""`.
* `customer_key_md5` - The 128-bit MD5 digest of the encryption key according to RFC 1321. Always `""`.
* `expiration` - If the object expiration is configured, this attribute will be set. Always `""`.
* `expected_bucket_owner` - Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error. Always empty.
* `expected_source_bucket_owner` - Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error. Always empty.
* `expires` - Date and time at which the object is no longer cacheable, in [RFC3339 format]. Always `""`.
* `force_destroy` - Allow the object to be deleted by removing any legal hold on any object version. Always `false`.
* `kms_encryption_context` - Specifies the AWS KMS Encryption Context to use for object encryption. Always `""`.
* `kms_key_id` - Specifies the AWS KMS Key ARN to use for object encryption. Always `""`.
* `metadata` - A map of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API). Always empty.
* `metadata_directive` - Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Always empty.
* `object_lock_legal_hold_status` - The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Always `""`.
* `object_lock_mode` - The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Always `""`.
* `object_lock_retain_until_date` - The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods). Always `""`.
* `request_charged` - If present, indicates that the requester was successfully charged for the request.  Always `false`.
* `request_payer` - Confirms that the requester knows that they will be charged for the request. Always empty.
* `server_side_encryption` - Specifies server-side encryption of the object in S3. Always `""`.
* `source_customer_algorithm` - Specifies the algorithm to use when decrypting the source object (for example, AES256). Always empty.
* `source_customer_key` - Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. Always `""`.
* `source_customer_key_md5` - Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Always empty.
* `source_version_id` - Version of the copied object in the source bucket. Always `""`.
* `storage_class` - Specifies the desired [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html#AmazonS3-CopyObject-request-header-StorageClass) for the object. Always `STANDARD`.
* `tagging_directive` - Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Always empty.
* `tags` - A map of tags to assign to the object. If configured with a provider [`default_tags` configuration block][default-tags] present, tags with matching keys will overwrite those defined at the provider-level. Always empty.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block][default-tags]. Always empty.
* `website_redirect` - Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html). Always `""`.

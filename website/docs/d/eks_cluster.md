---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "aws_eks_cluster"
description: |-
  Retrieves information about an EKS cluster.
---

# Data Source: aws_eks_cluster

Retrieves information about an EKS cluster.

## Example Usage

```terraform
data "aws_eks_cluster" "example" {
  name = "example"
}
```

## Argument Reference

* `name` - (Required) The name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).

## Attributes Reference

* `arn` - Cluster ID.
* `certificate_authority` - Nested attribute containing `certificate-authority-data` for your cluster.
    * `data` - The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
* `created_at` - The Unix epoch time stamp in seconds for when the cluster was created.
* `id` - The name of the cluster.
* `kubernetes_network_config` - Nested list containing Kubernetes Network Configuration.
    * `ip_family` - The IP family used to assign Kubernetes pod and service addresses.
    * `service_ipv4_cidr` - The CIDR block to assign Kubernetes service IP addresses from.
* `platform_version` - The platform version for the cluster.
* `status` - The status of the EKS cluster. One of `CLAIMED`, `CREATING`, `DELETED`, `DELETING`, `ERROR`, `MODIFYING`, `PENDING`, `PROVISIONING`, `READY`, `REPAIRING`.
* `version` - The Kubernetes server version for the cluster.
* `vpc_config` - Nested list containing VPC configuration for the cluster.
    * `cluster_security_group_id` - The cluster security group that was created by the cloud for the cluster.
    * `security_group_ids` – List of security group IDs.
    * `subnet_ids` – List of subnet IDs.
    * `vpc_id` – The VPC associated with your cluster.
* `tags` - Key-value map of resource tags.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `enabled_cluster_log_types` - The enabled control plane logs. Always empty.
* `encryption_config` - Configuration block with encryption configuration for the cluster. Always empty.
* `endpoint` - The endpoint for your Kubernetes API server. Always `""`.
* `identity` - Nested attribute containing identity provider information for your cluster. Always empty.
* `role_arn` - The ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to API operations on your behalf. Always `""`.
* `vpc_config` - Nested list containing VPC configuration for the cluster.
    * `endpoint_private_access` - Indicates whether or not the EKS private API server endpoint is enabled. Always `false`.
    * `endpoint_public_access` - Indicates whether or not the EKS public API server endpoint is enabled. Always `false`.
    * `public_access_cidrs` - List of CIDR blocks. Indicates which CIDR blocks can access the EKS public API server endpoint. Always empty.

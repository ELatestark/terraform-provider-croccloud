---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "aws_eks_node_group"
description: |-
  Retrieves information about an EKS node group.
---

# Data Source: aws_eks_node_group

Retrieves information about an EKS node group.

## Example Usage

```terraform
data "aws_eks_node_group" "example" {
  cluster_name    = "example"
  node_group_name = "example"
}
```

## Argument Reference

* `cluster_name` - (Required) The name of the cluster.
* `node_group_name` - (Required) The name of the node group.

## Attributes Reference

* `arn` - EKS node group ID.
* `disk_size` - Disk size in GiB for worker nodes.
* `id` - EKS cluster name and EKS node group name separated by a colon (`:`).
* `instance_types` - Set of instance types associated with the EKS node group.
* `labels` - Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS node group will not be managed.
* `remote_access` - Configuration block with remote access settings.
    * `ec2_ssh_key` - EC2 key pair name that provides access for SSH communication with the worker nodes in the EKS node group.
* `resources` - List of objects containing information about underlying resources.
    * `autoscaling_groups` - List of objects containing information about autoscaling groups.
        * `name` - Name of the autoscaling group.
* `scaling_config` - Configuration block with scaling settings.
    * `desired_size` - Desired number of worker nodes.
    * `max_size` - Maximum number of worker nodes.
    * `min_size` - Minimum number of worker nodes.
* `status` - Status of the EKS node group. One of `CREATING`, `ACTIVE`, `PENDING`, `UPDATING`, `DELETING`, `CREATE_FAILED`, `DELETE_FAILED`, `DEGRADED`.
* `subnet_ids` – Identifiers of EC2 Subnets to associate with the EKS node group.
* `tags` - Key-value map of resource tags.
* `taints` - List of objects containing information about taints applied to the nodes in the EKS node group.
    * `key` - The key of the taint.
    * `value` - The value of the taint.
    * `effect` - The effect of the taint.
* `version` – Kubernetes version.

->  **Unsupported attributes**
These attributes are currently unsupported:

* `ami_type` - Type of image associated with the EKS node group. Always `""`.
* `node_role_arn` – The ARN of the IAM Role that provides permissions for the EKS node group. Always `""`.
* `release_version` – Image version of the EKS node group. Always `""`.
* `remote_access` - Configuration block with remote access settings.
    * `source_security_group_ids` - Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. Always empty.
* `resources` - List of objects containing information about underlying resources.
    * `remote_access_security_group_id` - ID of the remote access EC2 security group. Always `""`.


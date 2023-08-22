# Cluster Resource

This resource is used to manage clusters in Cloudability.

## Example Usage

```hcl
resource "cloudability_cluster" "example" {
  cluster_name       = "example-cluster"
  kubernetes_version = "1.22"
}
```

## Argument Reference

* `cluster_name` - (Required) The name of the cluster
* `kubernetes_version` - (Optional) The version of Kubernetes that is running on the cluster. Either `kubernetes_version` or `cluster_version` must be specified.
* `cluster_version` - (Optional) The version the cluster is running on. Either `kubernetes_version` or `cluster_version` must be specified.

## Attribute Reference

* `cluster_name` - The name of the cluster
* `kubernetes_version` - The version of Kubernetes that is running on the cluster. 
* `cluster_version` - The version the cluster is running on. 

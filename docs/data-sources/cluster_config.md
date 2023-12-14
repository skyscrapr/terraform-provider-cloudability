# Cluster Config Datasource

This datasource tracks a Cluster Config in Cloudability.

## Example Usage

```hcl
data "cloudability_cluster_config" "test" {
  cluster_id = "12345"
}
```

## Argument Reference

* `cluster_id` - (Required) The ID of the cluster

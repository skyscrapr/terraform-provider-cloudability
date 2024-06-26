package cloudability

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"strconv"
	"time"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Create"))

	newCluster := &cloudability.Cluster{
		ClusterName:       d.Get("cluster_name").(string),
		KubernetesVersion: d.Get("kubernetes_version").(string),
		ClusterVersion:    d.Get("cluster_version").(string),
	}

	cluster, err := client.Containers().NewCluster(newCluster)
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("New cluster provisioned with ID: %d", cluster.ID))

	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(cluster.ID))
	d.Set("cluster_name", cluster.ClusterName)
	d.Set("kubernetes_version", cluster.KubernetesVersion)
	d.Set("cluster_version", cluster.ClusterVersion)

	return nil
}

func resourceClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Read"))

	cluster, err := client.Containers().GetCluster(d.Id())
	if err != nil {
		return err
	}
	if cluster == nil {
		tflog.Warn(context.Background(), "Cluster not found", map[string]interface{}{"id": d.Id()})
		d.SetId("")
		return nil
	}

	d.Set("cluster_name", cluster.ClusterName)
	d.Set("cluster_version", cluster.ClusterVersion)
	d.Set("kubernetes_version", cluster.KubernetesVersion)

	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Update"))

	idInt, _ := strconv.Atoi(d.Id())
	cluster := &cloudability.Cluster{
		ID:                idInt,
		ClusterName:       d.Get("cluster_name").(string),
		KubernetesVersion: d.Get("kubernetes_version").(string),
		ClusterVersion:    d.Get("cluster_version").(string),
	}

	err := client.Containers().UpdateCluster(cluster)
	if err != nil {
		return err
	}

	return resourceClusterRead(d, meta)
}

func resourceClusterDelete(d *schema.ResourceData, meta interface{}) error {
	tflog.Warn(context.Background(), "Deletion is not supported for Cluster resources", map[string]interface{}{"id": d.Id()})
	d.SetId("")

	return nil
}

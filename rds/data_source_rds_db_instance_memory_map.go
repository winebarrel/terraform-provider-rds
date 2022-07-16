package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRdsDbInstanceMemoryMap() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRdsDbInstanceMemoryMapRead,
		Schema: map[string]*schema.Schema{
			"memory_by_instance_class": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	}
}

func dataSourceRdsDbInstanceMemoryMapRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	memByClass := map[string]interface{}{}

	for k, v := range memoryByDBInstanceClass {
		memByClass[k] = v
	}

	if err := d.Set("memory_by_instance_class", memByClass); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())

	return diags
}

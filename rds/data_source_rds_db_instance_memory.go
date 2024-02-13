package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRdsDbInstanceMemory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRdsDbInstanceMemoryRead,
		Schema: map[string]*schema.Schema{
			"instance_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourceRdsDbInstanceMemoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	clazz := d.Get("instance_class").(string)
	mem, ok := memoryByDBInstanceClass[clazz]

	if !ok {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Instance class info not found",
			},
		}
	}

	if err := d.Set("memory", mem); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.UniqueId())

	return diags
}

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRdsDbInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRdsDbInstancesRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"values": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_class": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRdsDbInstancesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	svc := m.(*rds.Client)
	input := &rds.DescribeDBInstancesInput{}
	instances := []map[string]interface{}{}

	if filtersI, filtersOk := d.GetOk("filter"); filtersOk {
		input.Filters = buildFilters(filtersI.(*schema.Set))
	}

	for {
		output, err := svc.DescribeDBInstances(context.Background(), input)

		if err != nil {
			return diag.FromErr(err)
		}

		for _, i := range output.DBInstances {
			instance := map[string]interface{}{}
			instance["name"] = aws.ToString(i.DBInstanceIdentifier)
			instance["instance_class"] = aws.ToString(i.DBInstanceClass)
			instances = append(instances, instance)
		}

		if output.Marker == nil {
			break
		}

		input.Marker = output.Marker
	}

	if err := d.Set("instances", instances); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resource.UniqueId())

	return diags
}

func buildFilters(filterSet *schema.Set) []types.Filter {
	filters := []types.Filter{}

	for _, filterI := range filterSet.List() {
		filterMapI := filterI.(map[string]interface{})
		name := filterMapI["name"].(string)
		values := []string{}

		for _, v := range filterMapI["values"].(*schema.Set).List() {
			values = append(values, v.(string))
		}

		filters = append(filters, types.Filter{
			Name:   aws.String(name),
			Values: values,
		})
	}

	return filters
}

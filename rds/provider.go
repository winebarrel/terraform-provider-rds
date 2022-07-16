package rds

import (
	"context"
	_ "embed"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//go:embed rds.json
var rdsJson []byte

var memoryByDBInstanceClass map[string]float64

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"rds_instance_memory": dataSourceRdsDbInstanceMemory(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	err := json.Unmarshal(rdsJson, &memoryByDBInstanceClass)

	if err != nil {
		// must not happen
		panic(err)
	}

	return nil, diags
}

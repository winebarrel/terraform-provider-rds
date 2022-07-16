package rds

import (
	"context"
	_ "embed"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	sdkrds "github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//go:embed rds.json
var rdsJson []byte

var memoryByDBInstanceClass map[string]float64

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"rds_db_instance_memory":     dataSourceRdsDbInstanceMemory(),
			"rds_db_instance_memory_map": dataSourceRdsDbInstanceMemoryMap(),
			"rds_db_instances":           dataSourceRdsDbInstances(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	if err := json.Unmarshal(rdsJson, &memoryByDBInstanceClass); err != nil {
		// must not happen
		panic(err)
	}

	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	svc := sdkrds.NewFromConfig(cfg)

	return svc, diags
}

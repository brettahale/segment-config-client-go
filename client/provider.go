package segment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   false,
				DefaultFunc: schema.EnvDefaultFunc("SEGMENT_HOST", "https://platform.segmentapis.com"),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SEGMENT_TOKEN", nil),
			},
			"workspace": &schema.Schema{
            	Type:        schema.TypeString,
            	Required:    true,
            	Sensitive:   false,
            	DefaultFunc: schema.EnvDefaultFunc("SEGMENT_WORKSPACE", "default"),
            },
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"segment_destination":      dataSourceDestination(),
            "segment_source":           dataSourceSource(),
            "segment_tracking_plan":    dataSourceTrackingPlan(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
    host        := d.Get("host").(string)
    token       := d.Get("token").(string)
    workspace   := d.Get("workspace").(string)
	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics

	c, err := NewClient(&host, &workspace, &token)
    	if err != nil {
    		return nil, diag.FromErr(err)
    	}

    return c, diags

}

package helmoci

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"helm.sh/helm/v3/pkg/chartutil"
)

func dataOci() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataOciRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Release name.",
			},
			"version_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tag",
			},
			"chart_url": {
				Type:        schema.TypeString,
				Optional:    false,
				Description: "full chart url",
			},
			"registry_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Username for HTTP basic authentication",
			},
			"registry_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password for HTTP basic authentication",
			},
		},
	}
}

func dataOciRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	url := d.Get("chart_url").(string)
	user := d.Get("registry_username").(string)
	password := d.Get("registry_password").(string)
	tag := d.Get("version_tag").(string)

	if tag == "" {
		tag = "latest"
	}

	base := url[:strings.IndexByte(url, '/')]

	client, _ := NewClient()

	if user != "" || password != "" {
		err := client.Login(base, user, password, true)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	chartRef := &Reference{
		Repo: url,
		Tag:  tag,
	}

	_, err := client.PullChart(chartRef)

	if err != nil {
		return diag.FromErr(err)
	}

	chrt, err := client.LoadChart(chartRef)

	if err != nil {
		return diag.FromErr(err)
	}

	err = chartutil.SaveDir(chrt, "charts")

	if err != nil {
		return diag.FromErr(err)
	}

	return nil

}

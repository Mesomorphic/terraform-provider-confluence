package confluence

import (
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSpaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSpacesRead,
		Schema: map[string]*schema.Schema{
			"keys": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"spaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"author_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"homepage_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSpacesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	spaceKeys := interfaceSetToStringSlice(d.Get("keys"))
	spaceResponse, err := client.SearchSpaces(spaceKeys)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(fmt.Sprintf("%d", crc32.ChecksumIEEE([]byte(strings.Join(spaceKeys, ",")))))
	spaces := make([]map[string]interface{}, len(spaceResponse))
	for i := range spaceResponse {
		space := map[string]interface{}{
			"id":          spaceResponse[i].Id,
			"key":         spaceResponse[i].Key,
			"name":        spaceResponse[i].Name,
			"description": spaceResponse[i].Description,
			"type":        spaceResponse[i].Type,
			"author_id":   spaceResponse[i].AuthorId,
			"homepage_id": spaceResponse[i].HomepageId,
			"url":         client.URL(spaceResponse[i].Links.WebUI),
		}

		spaces[i] = space
	}
	_ = d.Set("spaces", spaces)
	return nil
}

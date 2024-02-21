package confluence

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSpace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSpaceRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
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
	}
}

func dataSourceSpaceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	spaceKeys := []string{d.Get("key").(string)}
	spaceResponse, err := client.SearchSpaces(spaceKeys)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("failed to find space: %v", err)
	}

	spaceCount := len(spaceResponse)
	if spaceCount < 1 {
		return fmt.Errorf("space with key '%s' does not exist", spaceKeys[0])
	}
	if spaceCount > 1 {
		return fmt.Errorf("found multiple spaces with key '%s', provide a unique key or use the plural data source", spaceKeys[0])
	}

	space := spaceResponse[0]
	d.SetId(space.Id)
	spaceMap := map[string]interface{}{
		"key":         space.Key,
		"name":        space.Name,
		"description": space.Description,
		"type":        space.Type,
		"author_id":   space.AuthorId,
		"homepage_id": space.HomepageId,
		"url":         client.URL(space.Links.WebUI),
	}

	for k, v := range spaceMap {
		err := d.Set(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

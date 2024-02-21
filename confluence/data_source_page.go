package confluence

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePageRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"space_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"author_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"view_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"edit_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	params := PageSearchParams{}
	if pageId, ok := d.Get("id").(string); ok {
		params.Id = pageId
	}
	if spaceId, ok := d.Get("space_id").(string); ok {
		params.SpaceID = spaceId
	}
	if title, ok := d.Get("title").(string); ok {
		params.Title = title
	}

	pageResponse, err := client.SearchPages(params, true)
	if err != nil {
		d.SetId("")
		return err
	}

	pageCount := len(pageResponse)
	if pageCount < 1 {
		return fmt.Errorf("unable to find page")
	}
	if pageCount > 1 {
		return fmt.Errorf("found multiple pages, provide a unique arguments or use the plural data source")
	}

	page := pageResponse[0]
	d.SetId(page.Id)
	pageMap := map[string]interface{}{
		"id":          page.Id,
		"space_id":    page.SpaceId,
		"title":       page.Title,
		"parent_id":   page.ParentId,
		"parent_type": page.ParentType,
		"author_id":   page.AuthorId,
		"owner_id":    page.OwnerId,
		"body":        page.Body.Storage.Value,
		"version":     page.Version.Number,
		"view_url":    client.URL(page.Links.WebUI),
		"edit_url":    client.URL(page.Links.EditUI),
	}

	for k, v := range pageMap {
		err := d.Set(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

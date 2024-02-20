package confluence

import (
	"fmt"
	"hash/crc32"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePagesRead,
		Schema: map[string]*schema.Schema{
			"space_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"space_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourcePagesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	params := PageSearchParams{}
	var idStr string
	if spaceId, ok := d.Get("space_id").(string); ok {
		params.SpaceID = spaceId
		idStr += spaceId
	}
	if title, ok := d.Get("title").(string); ok {
		params.Title = title
		idStr += title
	}

	pageResponse, err := client.SearchPages(params, false)
	if err != nil {
		d.SetId("")
		return err
	}

	pages := make([]map[string]interface{}, len(pageResponse))
	for i := range pageResponse {
		page := map[string]interface{}{
			"id":          pageResponse[i].Id,
			"space_id":    pageResponse[i].SpaceId,
			"title":       pageResponse[i].Title,
			"parent_id":   pageResponse[i].ParentId,
			"parent_type": pageResponse[i].ParentType,
			"author_id":   pageResponse[i].AuthorId,
			"owner_id":    pageResponse[i].OwnerId,
			"version":     pageResponse[i].Version.Number,
			"view_url":    client.URL(pageResponse[i].Links.WebUI),
			"edit_url":    client.URL(pageResponse[i].Links.EditUI),
		}

		pages[i] = page
	}

	d.SetId(fmt.Sprintf("%d", crc32.ChecksumIEEE([]byte(idStr))))
	_ = d.Set("pages", pages)

	return nil
}

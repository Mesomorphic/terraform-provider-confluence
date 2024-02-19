package confluence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Content is a primary resource in Confluence
type Space struct {
	Id    int         `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Key   string      `json:"key,omitempty"`
	Links *SpaceLinks `json:"_links,omitempty"`
}

// ContentLinks is part of Content
type SpaceLinks struct {
	Base  string `json:"base,omitempty"`
	WebUI string `json:"webui,omitempty"`
}

// The space schema from API v2, note the id datatype differs from the v1 API.
type SpaceV2 struct {
	Id          string        `json:"id,omitempty"`
	AuthorId    string        `json:"authorId,omitempty"`
	HomepageId  string        `json:"homepageId,omitempty"`
	Type        string        `json:"type,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Key         string        `json:"key,omitempty"`
	Links       *SpaceV2Links `json:"_links,omitempty"`
}

type SpaceV2Links struct {
	WebUI string `json:"webui,omitempty"`
}

type SpaceSearchResults struct {
	Results []SpaceV2                `json:"results,omitempty"`
	Links   SpaceSearchResponseLinks `json:"_links,omitempty"`
}

type SpaceSearchResponseLinks struct {
	Next string `json:"next,omitempty"`
}

func (c *Client) CreateSpace(space *Space) (*Space, error) {
	var response Space
	if err := c.Post("/rest/api/space", space, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSpace(id string) (*Space, error) {
	var response Space
	path := fmt.Sprintf("/rest/api/space/%s", id)
	if err := c.Get(path, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SearchSpaces(keys []string) ([]SpaceV2, error) {
	var response SpaceSearchResults
	var result []SpaceV2
	path := fmt.Sprintf("/api/v2/spaces?limit=250&keys=%s", strings.Join(keys, ","))
	if err := c.Get(path, &response); err != nil {
		return nil, err
	}

	for i := 0; i < 500; i++ {
		result = append(result, response.Results...)
		if response.Links.Next == "" {
			break
		}
		// Strip "/wiki" prefix from the provided URL.
		nextPageUrl := response.Links.Next[5:len(response.Links.Next)]
		if err := c.Get(nextPageUrl, &response); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Client) UpdateSpace(space *Space) (*Space, error) {
	var response Space

	path := fmt.Sprintf("/rest/api/space/%s", space.Key)
	if err := c.Put(path, space, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteSpace(id string) error {
	path := fmt.Sprintf("/rest/api/space/%s", id)
	if err := c.Delete(path); err != nil {
		if strings.HasPrefix(err.Error(), "202 ") {
			//202 is the delete API success response
			//Other APIs return 204. Because, reasons.
			return nil
		}
		return err
	}
	return nil
}

// Convert a resource string set to a native slice.
func interfaceSetToStringSlice(ifSet interface{}) []string {
	var strSlice []string
	ifList := ifSet.(*schema.Set).List()
	strSlice = make([]string, len(ifList))
	for i, rawVal := range ifList {
		if strVal, ok := rawVal.(string); ok {
			strSlice[i] = strVal
		}
	}
	return strSlice
}

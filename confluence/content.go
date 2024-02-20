package confluence

import (
	"fmt"
)

// Body is part of Content
type Body struct {
	Storage *Storage `json:"storage,omitempty"`
}

// Content is a primary resource in Confluence
type Content struct {
	Id        string        `json:"id,omitempty"`
	Type      string        `json:"type,omitempty"`
	Title     string        `json:"title,omitempty"`
	Space     *SpaceKey     `json:"space,omitempty"`
	Version   *Version      `json:"version,omitempty"`
	Body      *Body         `json:"body,omitempty"`
	Links     *ContentLinks `json:"_links,omitempty"`
	Ancestors []*Content    `json:"ancestors,omitempty"`
}

// ContentLinks is part of Content
type ContentLinks struct {
	Context string `json:"context,omitempty"`
	WebUI   string `json:"webui,omitempty"`
}

type Page struct {
	Id         string     `json:"id,omitempty"`
	SpaceId    string     `json:"spaceId,omitempty"`
	AuthorId   string     `json:"authorId,omitempty"`
	OwnerId    string     `json:"ownerId,omitempty"`
	ParentId   string     `json:"parentId,omitempty"`
	ParentType string     `json:"parentType,omitempty"`
	Title      string     `json:"title,omitempty"`
	Version    *Version   `json:"version,omitempty"`
	Body       *Body      `json:"body,omitempty"`
	Links      *PageLinks `json:"_links,omitempty"`
}

type PageLinks struct {
	EditUI string `json:"editui,omitempty"`
	TinyUI string `json:"tinyui,omitempty"`
	WebUI  string `json:"webui,omitempty"`
}

type PageSearchParams struct {
	Id      string
	SpaceID string
	Title   string
}

type PageSearchResponse struct {
	Results []Page                  `json:"results,omitempty"`
	Links   PageSearchResponseLinks `json:"_links,omitempty"`
}

type PageSearchResponseLinks struct {
	Next string `json:"next,omitempty"`
}

// SpaceKey is part of Content
type SpaceKey struct {
	Key string `json:"key,omitempty"`
}

// Storage is part of Body
type Storage struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

// Version is part of Content
type Version struct {
	Number int `json:"number,omitempty"`
}

func (c *Client) CreateContent(content *Content) (*Content, error) {
	var response Content
	if err := c.Post("/rest/api/content", content, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetContent(id string) (*Content, error) {
	var response Content
	path := fmt.Sprintf("/rest/api/content/%s?expand=space,body.storage,version,ancestors", id)
	if err := c.Get(path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetPage(id string) (*Page, error) {
	url := fmt.Sprintf("/api/v2/pages/%s?body-format=storage", id)

	var response Page
	if err := c.Get(url, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SearchPages(searchParams PageSearchParams, includeBody bool) ([]Page, error) {

	if searchParams.Id != "" {
		response, err := c.GetPage(searchParams.Id)
		if err != nil {
			return nil, err
		}
		return []Page{*response}, nil
	}

	url := "/api/v2/pages?limit=250"
	if includeBody {
		url += "&body-format=storage"
	}
	if searchParams.SpaceID != "" {
		url += fmt.Sprintf("&space-id=%s", searchParams.SpaceID)
	}
	if searchParams.Title != "" {
		url += fmt.Sprintf("&title=%s", searchParams.Title)
	}

	var response PageSearchResponse
	var result []Page
	if err := c.Get(url, &response); err != nil {
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

func (c *Client) UpdateContent(content *Content) (*Content, error) {
	var response Content
	content.Version.Number++
	path := fmt.Sprintf("/rest/api/content/%s", content.Id)
	if err := c.Put(path, content, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteContent(id string) error {
	path := fmt.Sprintf("/rest/api/content/%s", id)
	if err := c.Delete(path); err != nil {
		return err
	}
	return nil
}

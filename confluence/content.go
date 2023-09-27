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
	Id        string           `json:"id,omitempty"`
	Type      string           `json:"type,omitempty"`
	Title     string           `json:"title,omitempty"`
	Space     *SpaceKey        `json:"space,omitempty"`
	Version   *Version         `json:"version,omitempty"`
	Body      *Body            `json:"body,omitempty"`
	Links     *ContentLinks    `json:"_links,omitempty"`
	Ancestors []*Content       `json:"ancestors,omitempty"`
	Metadata  *ContentMetadata `json:"metadata,omitempty"`
}

// ContentLinks is part of Content
type ContentLinks struct {
	Context string `json:"context,omitempty"`
	WebUI   string `json:"webui,omitempty"`
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

// ContentMetadata is part of Content
type ContentMetadata struct {
	Labels []*Label `json:"labels,omitempty"`
}

// Label is part of Metadata
type Label struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
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

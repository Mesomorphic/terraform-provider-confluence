---
layout: "confluence"
page_title: "Data Source: confluence_pages"
sidebar_current: "docs-confluence-data-pages"
description: |-
  Get zero or more Confluence pages
---

# Data Source: confluence_pages

Retrieve details about zero or more existing Confluence pages.

## Example Usage

```hcl
data confluence_space "my_space" {
  key  = "MYSPACE"
}

data confluence_pages "my_space_pages" {
  space_id = data.confluence_space.my_space.id
}
```

## Argument Reference

The following arguments are supported:

- `space_id` - (Optional) The ID of the space containing the target Confluence page.
- `title` - (Optional) The target Confluence page's title.

## Attributes Reference

This resource exports the following attributes:

- `pages` - The list of pages that match the provided arguments, each with the following attributes:
  
  - `id` - (Optional) This page's ID.
  
  - `space_id` - (Optional) The ID of the space containing this page.
  
  - `title` - (Optional) This page's title.

  - `parent_id` - The content id of the page's parent.

  - `parent_type` - The content type of the page's parent.

  - `author_id` - The ID of the page's most recent author.

  - `owner_id` - The ID of the page's owner.

  - `view_url` - The URL to view this page.

  - `edit_url` - The URL to edit this page.

  - `version` - The version of this page.

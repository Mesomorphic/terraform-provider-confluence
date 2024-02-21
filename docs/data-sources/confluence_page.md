---
layout: "confluence"
page_title: "Data Source: confluence_page"
sidebar_current: "docs-confluence-data-page"
description: |-
  Get a Confluence page
---

# Data Source: confluence_page

Retrieve details about an existing Confluence page.

## Example Usage

```hcl
data confluence_space "my_space" {
  key  = "MYSPACE"
}

data confluence_page "my_page" {
  space_id = data.confluence_space.my_space.id
  title    = "my page"
}
```

## Argument Reference

The following arguments are supported:

- `id` - (Optional) The target Confluence page's ID.

- `space_id` - (Optional) The ID of the space containing the target Confluence page.

- `title` - (Optional) The target Confluence page's title.

## Attributes Reference

This resource exports the following attributes:

- `id` - (Optional) This page's ID.

- `space_id` - (Optional) The ID of the space containing this page.

- `title` - (Optional) This page's title.

- `body` - The actual content of the page in Confluence Storage Format.

- `parent_id` - The content id of the page's parent.

- `parent_type` - The content type of the page's parent.

- `author_id` - The ID of the page's most recent author.

- `owner_id` - The ID of the page's owner.

- `view_url` - The URL to view this page.

- `edit_url` - The URL to edit this page.

- `version` - The version of this page.

---
layout: "confluence"
page_title: "Data Source: confluence_spaces"
sidebar_current: "docs-confluence-data-spaces"
description: |-
  Get zero or more Confluence spaces
---

# Data Source: confluence_spaces

Retrieve details about zero or more existing Confluence spaces.

## Example Usage

```hcl
data confluence_spaces "all_spaces" {
}
```

## Argument Reference

The following arguments are supported:

- `keys` - (Optional) A list of space keys to retrieve, when omitted all spaces will be returned.

## Attribute Reference

The following attributes are provided:

- `spaces` - The list of spaces that match the provided arguments, each with the following attributes:
  
  - `id` - The space's id.
  
  - `key` - The space's key.
  
  - `name` - The space's name.
  
  - `type` - The space's type.
  
  - `description` - The space's description.
  
  - `author_id` - The ID of the space author.
  
  - `url` - The space's URL.

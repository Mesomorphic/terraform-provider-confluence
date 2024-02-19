---
layout: "confluence"
page_title: "Data Source: confluence_space"
sidebar_current: "docs-confluence-data-space"
description: |-
  Get a Confluence space
---

# Data Source: confluence_space

Retrieve details about an existing Confluence space.

## Example Usage

```hcl
data confluence_space "my_space" {
  key  = "MYSPACE"
}
```

## Argument Reference

The following arguments are supported:

- `key` - (Required) The key of the target space.

## Attribute Reference

The following attributes are provided:

- `id` - The space's id.

- `key` - The space's key.

- `name` - The space's name.

- `type` - The space's type.

- `description` - The space's description.

- `author_id` - The ID of the space author.

- `url` - The space's URL.

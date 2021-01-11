---
page_title: "domeneshop_forward Resource - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this resource to create and manage HTTP forwards ("WWW forwarding").
---

# Resource `domeneshop_forward`

Use this resource to create and manage HTTP forwards ("WWW forwarding").

## Example Usage

```terraform
resource "domeneshop_forward" "example" {
  domain_id = 1234
  host      = "www"
  url       = "https://example.com/"
}
```

## Schema

### Required

- **domain_id** (Number) ID of the domain.
- **host** (String) Subdomain of the forward, `@` for the root domain.
- **url** (String) The URL to forward to. Must include scheme, e.g. `https://` or `ftp://`.

### Optional

- **frame** (Boolean) Whether to enable frame forwarding using an iframe embed. NOT recommended for a variety of reasons.
- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import domeneshop_forward.example <domain_id>/<host>
```
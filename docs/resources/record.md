---
page_title: "domeneshop_record Resource - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this resource to create and manage DNS records.
---

# Resource `domeneshop_record`

Use this resource to create and manage DNS records.

## Example Usage

```terraform
resource "domeneshop_record" "example" {
  domain_id = 1234
  host      = "example"
  type      = "A"
  data      = "192.0.2.56"
  ttl       = 300
}
```

## Schema

### Required

- **data** (String) The value of the record.
- **domain_id** (Number) ID of the domain.
- **host** (String) The host/subdomain the DNS record applies to.
- **type** (String) The type of the record. Possible values are: `A`, `AAAA`, `ANAME`, `CNAME`, `DS`, `MX`, `NS`, `SRV`, `TXT`, `TLSA`.

### Optional

- **alg** (Number) DS record algorithm.
- **digest** (Number) DS record digest type.
- **dtype** (Number) TLSA record matching type.
- **flags** (Number) CAA record flags.
- **id** (String) The ID of this resource.
- **port** (Number) SRV record port. The port where the service is found.
- **priority** (Number) MX/SRV record priority, also known as preference. Lower values are usually preferred first, but this is not guaranteed
- **selector** (Number) TLSA record selector.
- **tag** (Number) CAA/DS record tag.
- **ttl** (Number) TTL of DNS record in seconds.
- **usage** (Number) TLSA record certificate usage.
- **weight** (Number) SRV record weight. Relevant if multiple records have same preference.



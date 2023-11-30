---
title: "Steampipe Table: crtsh_ca_issuer - Query crt.sh CA Issuers using SQL"
description: "Allows users to query crt.sh Certificate Authority (CA) Issuers, providing insights into the entities that issue SSL/TLS certificates."
---

# Table: crtsh_ca_issuer - Query crt.sh CA Issuers using SQL

crt.sh is a free Certificate Transparency Log (CTL) Search tool. It provides a useful mechanism to monitor SSL/TLS certificates issued for a particular domain. This tool can aid in identifying misissued and rogue certificates, thereby enhancing the security posture of an organization.

## Table Usage Guide

The `crtsh_ca_issuer` table provides insights into the Certificate Authorities (CAs) that issue SSL/TLS certificates. As a security analyst or IT administrator, explore details of CA issuers through this table, including their names, keys and associated metadata. Utilize it to uncover information about CA issuers, such as the number of certificates issued by each CA, and the verification of issuer keys.

## Examples

### Issuers with their CA detail
Determine the areas in which issuers and their corresponding Certificate Authorities (CA) interact. This can provide insights into the relationship between issuers and CAs, helping to understand the trust hierarchy in a digital certificate infrastructure.

```sql
with ca_issuers as (
  select * from crtsh_ca_issuer order by ca_id
),
cas as (
  select * from crtsh_ca order by id
)
select
  *
from
  ca_issuers,
  cas
where
  ca_issuers.ca_id = cas.id;
```

### CA Issuers by Content Type
Uncover the details of different content types and their corresponding counts to understand which types are most prevalent. This can be useful in assessing the distribution and dominance of specific content types.

```sql
select
  content_type,
  count(*)
from
  crtsh_ca_issuer
group by
  content_type
order by
  count desc;
```

### Inactive CA Issuers
Discover the segments that consist of Certificate Authority (CA) issuers that are currently inactive. This is useful in maintaining network security by identifying and managing inactive elements.

```sql
select
  ca_id,
  url,
  result,
  is_active
from
  crtsh_ca_issuer
where
  not is_active;
```

### Get all CA's issued by the CA Issuer with ID 12
This example allows you to identify all the Certificate Authorities (CAs) that have been issued by a specific CA issuer. This is particularly useful in managing digital certificates and ensuring secure communication within your network.

```sql
with ca_certs as (
  select
    ca_cert_id::bigint as id
  from
    crtsh_ca_issuer as cai,
    jsonb_array_elements_text(cai.ca_certificate_ids) as ca_cert_id
  where
    cai.ca_id = 12
  order by
    id
)
select
  *
from
  crtsh_ca
where
  id in (select id from ca_certs);
```

### Check URL of CA Issuers that reported text/plain content
Explore which CA issuers have reported 'text/plain' content and assess the corresponding HTTP response details to identify any potential issues or errors. This can be useful in monitoring and maintaining the integrity and reliability of certificate authorities within your network.

```sql
select
  cai.ca_id,
  cai.url,
  req.response_status_code,
  req.response_error,
  jsonb_pretty(req.response_headers)
from
  crtsh_ca_issuer as cai,
  net_http_request as req
where
  cai.content_type = 'text/plain'
  and req.url = cai.url;
```
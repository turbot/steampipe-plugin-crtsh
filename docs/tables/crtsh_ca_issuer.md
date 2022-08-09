# Table: crtsh_ca_issuer

Certificate Authority Issuers known to crt.sh, including the status of their last check.

## Examples

### Issuers with their CA detail

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

# Table: crtsh_certificate

Find certificates from certificate transparency log records.

## Examples

### All certificates for a given domain and its subdomains

```sql
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io'
```

### Get a specific certificate by crt.sh ID

```sql
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  id = 7203584052
```

### Current certificates for a specific domain

```sql
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'cloud.steampipe.io'
  and dns_names ? 'cloud.steampipe.io'
  and not_before < now()
  and not_after > now()
```

### Certificates valid at the current time

```sql
select
  dns_names,
  not_before,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before < now()
  and not_after > now()
```

### Certificates expiring in the next 30 days

```sql
select
  dns_names,
  not_before,
  not_after,
  -age(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_after between now() and now() + interval '30 days'
```

### Certificates issued in the last 30 days

```sql
select
  dns_names,
  not_before,
  not_after,
  -age(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before between now() and now() - interval '30 days'
```

### Certificate Authorities that have issued certificates for my domain

```sql
select
  issuer -> 'Organization' ->> 0 as issuer_org,
  count(*)
from
  crtsh_certificate
where
  query = 'steampipe.io'
group by
  issuer_org
order by
  count desc
```

### Certificates by public key algorithm

```sql
select
  public_key_algorithm,
  count(*)
from
  crtsh_certificate
where
  query = 'steampipe.io'
group by
  public_key_algorithm
order by
  count desc
```

### Get certificate log entries for all current certificates of a domain

```sql
-- Use a CTE with order by to force the Postgres planning sequence
with certs as (
  select
    *
  from
    crtsh_certificate
  where
    query = 'cloud.steampipe.io'
    and dns_names ? 'cloud.steampipe.io'
    and not_before < now()
    and not_after > now()
  order by id
)
select
  le.entry_id,
  le.ct_log_id,
  le.certificate_id,
  c.dns_names
from
  certs as c,
  crtsh_log_entry as le
where
  c.id = le.certificate_id
order by
  le.entry_id
```

---
title: "Steampipe Table: crtsh_certificate - Query crt.sh Certificate Transparency Logs using SQL"
description: "Allows users to query Certificate Transparency Logs from crt.sh, providing insights into the SSL/TLS certificates for a given domain."
---

# Table: crtsh_certificate - Query crt.sh Certificate Transparency Logs using SQL

crt.sh is a Certificate Transparency Log (CTL) monitor and search engine developed by Sectigo. It allows users to search for SSL/TLS certificates issued for a specific domain by various certificate authorities. This tool is useful for identifying misissued certificates or discovering certificates issued for your domains by unauthorized CAs.

## Table Usage Guide

The `crtsh_certificate` table provides insights into the SSL/TLS certificates for a specific domain. As a security engineer or a site administrator, explore certificate-specific details through this table, including issuer name, validity period, and associated metadata. Utilize it to uncover information about certificates, such as those that are near their expiration date, issued by unauthorized certificate authorities, and the verification of certificate details.

## Examples

### All certificates for a given domain and its subdomains
Determine the validity period of all security certificates associated with a particular domain and its subdomains. This is useful for ensuring ongoing website security and preventing unexpected certificate expirations.

```sql+postgres
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io';
```

```sql+sqlite
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io';
```

### Enumerate and discover subdomains for a domain via certificate transparency
Explore the subdomains associated with a specific domain to understand its structure and relationships. This can be useful for identifying potential security vulnerabilities or for mapping out the digital footprint of a domain.

```sql+postgres
with raw_domains as (
  select distinct
    jsonb_array_elements_text(dns_names) as domain
  from
    crtsh_certificate
  where
    query = 'steampipe.io'
)
select
  *
from
  raw_domains
where
  -- filter out mixed domains (e.g. from shared status page services)
  domain like '%steampipe.io'
order by
  domain;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```

### Get a specific certificate by crt.sh ID
Identify instances where a specific certificate, based on its crt.sh ID, is about to expire. This allows for proactive renewal and avoids potential service disruptions.

```sql+postgres
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  id = 7203584052;
```

```sql+sqlite
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  id = 7203584052;
```

### Certificates valid at the current time
Explore which certificates are currently valid for a specific domain. This can help ensure the security and authenticity of the domain, making it a useful tool for maintaining online safety standards.

```sql+postgres
select
  dns_names,
  not_before,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before < now()
  and not_after > now();
```

```sql+sqlite
select
  dns_names,
  not_before,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before < datetime('now')
  and not_after > datetime('now');
```

### Current certificates for a specific domain
Explore which current certificates are valid for a specific domain to ensure secure and encrypted connections. This is beneficial in identifying any potential security risks or lapses in your domain's SSL/TLS setup.

```sql+postgres
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'cloud.steampipe.io'
  and dns_names ? 'cloud.steampipe.io'
  and not_before < now()
  and not_after > now();
```

```sql+sqlite
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'cloud.steampipe.io'
  and json_extract(dns_names, '$.cloud.steampipe.io') is not null
  and not_before < datetime('now')
  and not_after > datetime('now');
```

### Certificates expiring in the next 30 days
Assess the elements within your domain's SSL certificates that are set to expire within the next 30 days. This is useful for maintaining website security and avoiding service interruptions due to expired certificates.

```sql+postgres
select
  dns_names,
  not_before,
  not_after,
  -age(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_after between now() and now() + interval '30 days';
```

```sql+sqlite
select
  dns_names,
  not_before,
  not_after,
  julianday('now') - julianday(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_after between datetime('now') and datetime('now', '+30 days');
```

### Certificates issued in the last 30 days
Determine the areas in which certificates have been issued in the last 30 days for a specific domain. This allows for an understanding of the certificate's lifespan and helps in tracking their expiry dates.

```sql+postgres
select
  dns_names,
  not_before,
  not_after,
  -age(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before between now() and now() - interval '30 days';
```

```sql+sqlite
select
  dns_names,
  not_before,
  not_after,
  julianday('now') - julianday(not_after) as expiration_countdown
from
  crtsh_certificate
where
  query = 'steampipe.io'
  and not_before between datetime('now') and datetime('now', '-30 days');
```

### Certificate Authorities that have issued certificates for my domain
Explore which certificate authorities have issued certificates for your domain, enabling you to assess the security and credibility of your website's SSL certificates. This query is beneficial in identifying potential security risks and ensuring only trusted authorities are used.

```sql+postgres
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
  count desc;
```

```sql+sqlite
select
  json_extract(json_extract(issuer, '$.Organization'), '$[0]') as issuer_org,
  count(*)
from
  crtsh_certificate
where
  query = 'steampipe.io'
group by
  issuer_org
order by
  count(*) desc;
```

### Certificates by public key algorithm
Determine the prevalence of different public key algorithms used in certificates related to a specific domain. This can help you understand the security measures in place and identify potential vulnerabilities.

```sql+postgres
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
  count desc;
```

```sql+sqlite
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
  count(*) desc;
```

### Get certificate log entries for all current certificates of a domain
Determine the areas in which current domain certificates have logged entries. This is useful to understand the activity and validity of your domain's certificates, helping you maintain secure and active certificates.

```sql+postgres
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
  le.entry_id;
```

```sql+sqlite
with certs as (
  select
    *
  from
    crtsh_certificate
  where
    query = 'cloud.steampipe.io'
    and json_extract(dns_names, '$."cloud.steampipe.io"') is not null
    and not_before < datetime('now')
    and not_after > datetime('now')
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
  le.entry_id;
```

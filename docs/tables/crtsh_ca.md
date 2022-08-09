# Table: crtsh_ca

[Certificate Authorities](https://en.wikipedia.org/wiki/Certificate_authority) found in certificate logs.

## Examples

### Top 10 CA's by number of certificates issued

```sql
select
  *
from
  crtsh_ca
order by
  num_certs_issued desc nulls last
limit 10;
```

### Details of top 10 CA's by number of certificates issued

```sql
select
  id,
  name,
  num_certs_issued - num_certs_expired as num_certs_current,
  num_certs_issued
from
  crtsh_ca
order by
  num_certs_current desc nulls last
limit 10;
```

### CA's based in Australia

```sql
select
  *
from
  crtsh_ca
where
  name ilike 'C=AU%'
order by
  name;
```

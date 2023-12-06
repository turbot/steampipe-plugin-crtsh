---
title: "Steampipe Table: crtsh_ca - Query crt.sh Certificate Authorities using SQL"
description: "Allows users to query crt.sh Certificate Authorities, specifically enabling the retrieval of detailed information about various Certificate Authorities recorded by crt.sh."
---

# Table: crtsh_ca - Query crt.sh Certificate Authorities using SQL

crt.sh is a service provided by Sectigo that monitors the issuance of SSL/TLS certificates by various Certificate Authorities. It allows users to search for certificates by several criteria, including domain name, Subject Public Key Info, and Certificate Authority. It provides a comprehensive view of the SSL/TLS certificate ecosystem, helping identify misissued certificates and potential security risks.

## Table Usage Guide

The `crtsh_ca` table provides insights into Certificate Authorities within crt.sh. As a security analyst, explore details about each Certificate Authority through this table, including their name, key, and associated metadata. Utilize it to uncover information about Certificate Authorities, such as their key details, the issuance of certificates, and the verification of their status.

## Examples

### Top 10 CA's by number of certificates issued
Determine the areas in which Certificate Authorities have issued the most certificates. This can assist in identifying which Certificate Authorities are the most active or popular, providing valuable insights into the digital certificate landscape.

```sql+postgres
select
  *
from
  crtsh_ca
order by
  num_certs_issued desc nulls last
limit 10;
```

```sql+sqlite
select
  *
from
  crtsh_ca
order by
  num_certs_issued desc
limit 10;
```

### Details of top 10 CA's by number of certificates issued
Explore which Certificate Authorities have issued the most certificates, providing a ranking of the top 10 based on the number of current, non-expired certificates they've issued. This helps in understanding the distribution and influence of various Certificate Authorities in the industry.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  num_certs_issued - num_certs_expired as num_certs_current,
  num_certs_issued
from
  crtsh_ca
order by
  (case when num_certs_current is null then 1 else 0 end), num_certs_current desc
limit 10;
```

### CA's based in Australia
Explore which Certificate Authorities are based in Australia. This can be beneficial when you want to identify and assess the elements within the Australian digital security landscape.

```sql+postgres
select
  *
from
  crtsh_ca
where
  name ilike 'C=AU%'
order by
  name;
```

```sql+sqlite
select
  *
from
  crtsh_ca
where
  name like 'C=AU%'
order by
  name;
```
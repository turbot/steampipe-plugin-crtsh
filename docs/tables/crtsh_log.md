---
title: "Steampipe Table: crtsh_log - Query crt.sh Certificate Logs using SQL"
description: "Allows users to query Certificate Logs from crt.sh, specifically providing details about issued SSL certificates, including their validity, issuer details, and associated domains."
---

# Table: crtsh_log - Query crt.sh Certificate Logs using SQL

crt.sh is a free Certificate Transparency Log (CTL) search engine provided by Sectigo. It allows users to search for, and retrieve, details about SSL certificates that have been issued by all Certificate Authorities participating in the CTL. This includes information about the validity of certificates, issuer details, and the domains associated with each certificate.

## Table Usage Guide

The `crtsh_log` table provides insights into SSL certificates logged in the crt.sh Certificate Transparency Log. As a security analyst or systems administrator, explore certificate-specific details through this table, including issuer data, associated domains, and certificate validity. Utilize it to uncover information about certificates, such as those issued by specific Certificate Authorities, the domains they cover, and their validity periods.

## Examples

### Active log operators
Explore which log operators are currently active, helping you understand the latest updates and inclusion status in both Apple and Chrome. This can assist in identifying potential issues or changes in their status that may require attention.

```sql
select
  id,
  operator,
  name,
  url,
  is_active,
  latest_update,
  latest_sth_timestamp,
  apple_inclusion_status,
  chrome_inclusion_status
from
  crtsh_log
where
  is_active
order by
  operator,
  name;
```

### Logs run by Google
Explore the logs operated by Google to gain insights into their activity status, update frequency, and inclusion status in Apple and Chrome. This helps in monitoring and assessing the performance and reach of these logs.

```sql
select
  id,
  operator,
  name,
  url,
  is_active,
  latest_update,
  apple_inclusion_status,
  chrome_inclusion_status
from
  crtsh_log
where
  operator = 'Google'
order by
  operator,
  name;
```

### Log operators included in Chrome
Explore which log operators are currently usable in Chrome to ensure you're working with updated and active resources. This can be particularly useful in maintaining secure and efficient operations.

```sql
select
  id,
  operator,
  name,
  url,
  is_active,
  latest_update,
  apple_inclusion_status,
  chrome_inclusion_status
from
  crtsh_log
where
  chrome_inclusion_status = 'Usable';
```

### Log operators with a different inclusion status in Apple and Chrome
Analyze the settings to understand discrepancies between the inclusion statuses of log operators in Apple and Chrome. This query is useful for identifying inconsistencies in the status of the same operator across different platforms.

```sql
select
  id,
  operator,
  name,
  url,
  is_active,
  latest_update,
  apple_inclusion_status,
  chrome_inclusion_status
from
  crtsh_log
where
  chrome_inclusion_status <> apple_inclusion_status;
```

### Log operators by Chrome inclusion status
Assess the distribution of log operators based on their inclusion status in Chrome. This is useful for understanding the prevalence of different inclusion statuses within your logs, which can inform security and compliance efforts.

```sql
select
  chrome_inclusion_status,
  count(*)
from
  crtsh_log
group by
  chrome_inclusion_status
order by
  count desc;
```
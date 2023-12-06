---
title: "Steampipe Table: crtsh_log_entry - Query crt.sh Log Entries using SQL"
description: "Allows users to query crt.sh Log Entries, specifically the certificate transparency log entries, providing insights into SSL/TLS certificates and their transparency."
---

# Table: crtsh_log_entry - Query crt.sh Log Entries using SQL

crt.sh is a free public Certificate Transparency Log (CT Log) search engine provided by Sectigo. It allows users to search for SSL/TLS certificates by various criteria, including domain name, organization name, and many others. The service helps to enhance transparency and security in the use of SSL/TLS certificates.

## Table Usage Guide

The `crtsh_log_entry` table provides insights into the Certificate Transparency Log (CT Log) entries in crt.sh. As a security analyst, explore entry-specific details through this table, including certificate details, log operator, and associated metadata. Utilize it to uncover information about the certificates, such as those issued by specific organizations, the CT logs they are included in, and the verification of the certificates' transparency.

## Examples

### Log entries for a particular certificate
Determine the log entries associated with a specific certificate to analyze its activity and troubleshoot potential issues.

```sql+postgres
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046;
```

```sql+sqlite
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046;
```

### Most recent entries for a given log
Analyze the most recent entries in a given log to monitor changes or unusual activity over the past hour. This can be particularly useful for identifying potential security issues or troubleshooting ongoing problems.

```sql+postgres
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046
  and ct_log_id = 91
  and entry_timestamp > now() - interval '1 hr';
```

```sql+sqlite
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046
  and ct_log_id = 91
  and entry_timestamp > datetime('now', '-1 hour');
```
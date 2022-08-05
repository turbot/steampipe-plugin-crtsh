# Table: crtsh_log_entry

Certificate transparency log operators that track and record log entries.

## Examples

### Log entries for a certificate

```sql
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046
```

### Log entries for a certificate

```sql
select
  *
from
  crtsh_certificate2
where
  query = 'steampipe.io'
  and dns_names ? 'steampipe.io'
```

### Most recent entries for a given log

```sql
select
  *
from
  crtsh_log_entry
where
  ct_log_id = 100
  and entry_timestamp > now() - interval '1 hr'
```

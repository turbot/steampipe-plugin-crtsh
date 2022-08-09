# Table: crtsh_log_entry

Certificate transparency log operators that track and record log entries.

## Examples

### Log entries for a particular certificate

```sql
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046;
```

### Most recent entries for a given log

```sql
select
  *
from
  crtsh_log_entry
where
  certificate_id = 6760944046
  and ct_log_id = 91
  and entry_timestamp > now() - interval '1 hr';
```

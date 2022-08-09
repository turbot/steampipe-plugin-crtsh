# Table: crtsh_log

Certificate transparency log operators that track and record log entries.

## Examples

### Active log operators

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

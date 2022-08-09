![image](https://hub.steampipe.io/images/plugins/turbot/crtsh-social-graphic.png)

# crt.sh Plugin for Steampipe

Use SQL to query certificates, log entries and more from the crt.sh certificate transparency database.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/crtsh)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/crtsh/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-crtsh/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install crtsh
```

Configure the server address in `~/.steampipe/config/crtsh.spc`:

```hcl
connection "crtsh" {
  plugin = "crtsh"
}
```

Run steampipe:

```shell
steampipe query
```

Query certificates:

```sql
select
  dns_names,
  not_after
from
  crtsh_certificate
where
  query = 'steampipe.io';
```

```
+------------------------+---------------------------+
| dns_names              | not_after                 |
+------------------------+---------------------------+
| ["steampipe.io"]       | 2022-10-24T08:48:52-04:00 |
| ["cloud.steampipe.io"] | 2022-10-20T22:56:08-04:00 |
+------------------------+---------------------------+
```

Enumerate and discover subdomains for a given domain:

```sql
with raw_domains as (
  -- Search for any certificates matching steampipe.io
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

```
+--------------------+
| domain             |
+--------------------+
| cloud.steampipe.io |
| hub.steampipe.io   |
| steampipe.io       |
| www.steampipe.io   |
+--------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-crtsh.git
cd steampipe-plugin-crtsh
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/crtsh.spc
```

Try it!

```
steampipe query
> .inspect crtsh
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-crtsh/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [crt.sh Plugin](https://github.com/turbot/steampipe-plugin-crtsh/labels/help%20wanted)

---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/crtsh.svg"
brand_color: "#00B373"
display_name: "crt.sh"
short_name: "crtsh"
description: "Steampipe plugin to query certificates, logs and more from the crt.sh certificate transparency database."
og_description: "Query crt.sh with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/crtsh-social-graphic.png"
---

# crt.sh + Steampipe

[crt.sh](https://crt.sh) provides a searchable database of certificate transparency logs.

[Certificate Transparency](https://en.wikipedia.org/wiki/Certificate_Transparency) is an
Internet security standard and open source framework for monitoring and
auditing digital certificates. The standard creates a system of public logs
that seek to eventually record all certificates issued by publicly trusted
certificate authorities, allowing efficient identification of mistakenly or
maliciously issued certificates.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Query certificates for a domain:

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/crtsh/tables)**

## Get started

### Install

Download and install the latest crt.sh plugin:

```bash
steampipe plugin install crtsh
```

### Configuration

Installing the latest crtsh plugin will create a config file (`~/.steampipe/config/crtsh.spc`) with a single connection named `crtsh`:

```hcl
connection "crtsh" {
  plugin = "crtsh"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-crtsh
- Community: [Slack Channel](https://steampipe.io/community/join)

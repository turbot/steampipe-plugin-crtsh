## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#12](https://github.com/turbot/steampipe-plugin-crtsh/pull/12))
- Recompiled plugin with Go version `1.21`. ([#12](https://github.com/turbot/steampipe-plugin-crtsh/pull/12))

## v0.2.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#6](https://github.com/turbot/steampipe-plugin-crtsh/pull/6))

## v0.1.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#3](https://github.com/turbot/steampipe-plugin-crtsh/pull/3))

## v0.0.1 [2022-08-10]

_What's new?_

- New tables added
  - [crtsh_ca](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_ca)
  - [crtsh_ca_issuer](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_ca_issuer)
  - [crtsh_certificate](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_certificate)
  - [crtsh_log](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_log)
  - [crtsh_log_entry](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_log_entry)
  - [crtsh_log_operator](https://https://hub.steampipe.io/plugins/turbot/crtsh/tables/crtsh_log_operator)

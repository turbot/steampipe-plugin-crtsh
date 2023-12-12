## v0.4.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-crtsh/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#23](https://github.com/turbot/steampipe-plugin-crtsh/pull/23))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-crtsh/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-crtsh/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to  column, and fixing connection and potential divide-by-zero bugs. ([#](https://github.com/turbot/steampipe-plugin-crtsh/pull/))

## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#15](https://github.com/turbot/steampipe-plugin-crtsh/pull/15))

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

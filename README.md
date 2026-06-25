[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nmollerup/sensu-check-load)
![Go Test](https://github.com/nmollerup/sensu-check-load/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nmollerup/sensu-check-load/workflows/goreleaser/badge.svg)

# Sensu load average check

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

The Sensu load average check is a [Sensu Check][1] that alerts on system load
average per CPU core. It is a Go port of the Ruby
[sensu-plugins-load-checks][6] plugin. Metrics are provided in
[nagios_perfdata][5] format.

Load averages are divided by the number of logical CPU cores so that thresholds
remain consistent across machines with different core counts.

## Usage examples

```
Check system load average per CPU core

Usage:
  sensu-check-load [flags]
  sensu-check-load [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
      --warn-load1 float    Warning threshold for 1-minute per-core load average (default 2.75)
      --warn-load5 float    Warning threshold for 5-minute per-core load average (default 2.5)
      --warn-load15 float   Warning threshold for 15-minute per-core load average (default 2)
      --crit-load1 float    Critical threshold for 1-minute per-core load average (default 3.5)
      --crit-load5 float    Critical threshold for 5-minute per-core load average (default 3.25)
      --crit-load15 float   Critical threshold for 15-minute per-core load average (default 3)
  -h, --help                help for sensu-check-load

Use "sensu-check-load [command] --help" for more information about a command.
```

## Configuration

### Asset registration

[Sensu Assets][2] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add nmollerup/sensu-check-load
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][3].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-load
  namespace: default
spec:
  command: >-
    sensu-check-load
    --warn-load1 2.75
    --warn-load5 2.5
    --warn-load15 2.0
    --crit-load1 3.5
    --crit-load5 3.25
    --crit-load15 3.0
  output_metric_format: nagios_perfdata
  output_metric_handlers:
    - influxdb
  subscriptions:
  - system
  runtime_assets:
  - nmollerup/sensu-check-load
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable from this
source.

From the local path of the sensu-check-load repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][4].

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[3]: https://bonsai.sensu.io/assets/nmollerup/sensu-check-load
[4]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[5]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/collect-metrics-with-checks/#supported-output-metric-formats
[6]: https://github.com/sensu-plugins/sensu-plugins-load-checks

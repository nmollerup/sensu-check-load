# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][1] and this project adheres to [Semantic Versioning][2].

## [Unreleased]

## [0.1.0] - 2026-06-25

### Added
- Initial release of `check-load` — a Go port of [sensu-plugins-load-checks][3]
- Per-core load average check for 1, 5, and 15 minute intervals
- Configurable warning and critical thresholds for each interval
- Nagios performance data output

[1]: https://keepachangelog.com/en/1.0.0/
[2]: https://semver.org/spec/v2.0.0.html
[3]: https://github.com/sensu-plugins/sensu-plugins-load-checks

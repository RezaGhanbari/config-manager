
![Carpino config manager](http://www.myiconfinder.com/uploads/iconsets/256-256-5fbc60a4335d01cd9c35dcf8fae02410.png)

# Config Manager

[![Go Report Card](https://goreportcard.com/badge/github.com/CarpinoTaxi/config-manager)](https://goreportcard.com/report/github.com/CarpinoTaxi/config-manager)


This is a lightweight, open source config manager and discovery management platform tha enables you to control your microservice environments from a single file. It will watch over changes and reloads the config every n seconds.


Go version 1.7 is required to build `master`, the current
development version. This project is officially supported on `linux/amd64`,
`linux/i386` and `linux/arm64`.

Tests are run against both Go versions 1.9 & 1.10, however at present, only Go 1.9 is officially supported.

## What is a Config Manager?

Config server can be treated as a remote property source. It can be backed by numerous storage providers, like Git (which is the default setting), SVN, or distributed file system. You can even have more than one storage at the same time. With Git we can easily manage configuration files for any application and any environment. Moreover, we gain access control, which means we can decide who can modify particular files (for example restricting access to production properties). We also have full traceability - we know who changed what, and when, and we can easily revert those changes.




### License
please see [LICENSE.md](LICENSE.md) for a full version of the license.

### Contributing

For more information about contributing PRs and issues, see [CONTRIBUTING.md](CONTRIBUTING.md).


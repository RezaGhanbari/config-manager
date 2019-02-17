
# Contributing

1. Sign one of the contributor license agreements below.
1. [Install Go](https://golang.org/doc/install)
1. Get the package:

    `go get -d github.com/CarpinoTaxi/config-manager`
1. Change into the checked out source:

    `cd $(go env GOPATH)/src/github.com/CarpinoTaxi/config-manager`
1. Fork the repo.
1. Set your fork as a remote:

    `git remote add fork git@github.com:GITHUB_USERNAME/config-manage.git`
1. Make changes, commit to your fork.
1. Send a pull request with your changes.

# Formatting

All code must be formatted with `gofmt` (with the latest Go version) and pass
`go vet`.

# Testing

Tests are required for all samples. When writing a pull request, be sure to
write and run the tests in any modified directories.



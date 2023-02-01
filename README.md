<div align="center">
<p>
    <img
        style="width: 200px"
        width="200"
        src="https://avatars.githubusercontent.com/u/4426989?s=200&v=4"
    >
</p>
<h1>Golang Reference</h1>

[Changelog](#) |
[Contributing](./CONTRIBUTING.md)

</div>
<div align="center">

[![Golang CI][gocibadge]][gociurl]
[![Shell Check][shellcheckbadge]][shellcheckurl]
[![License: MIT][licensebadge]][licenseurl]

</div>

This repository contains different GO apps and tools that are used in different projects here at [NaN Labs](https://www.nanlabs.com/). We also provide reusable GO packages.

## Applications

Collection of apps and examples that were created as a composition of different examples that
can be found separately in the [examples](./examples/) directory.
Read more about the examples in the [examples](#examples) section.

## Packages

We provide a collection of different packages that can be used in different projects. Each package has its own README file with more details about the package. Check out the [pkg](./pkg/) directory.

## Examples

Collection of examples that solve specific problems using small pieces of code.

We have a collection of examples that solve specific problems using small pieces of code. Each example has its own README file with more details about the example. Check out the [examples](./examples/) directory for more details!

- [Todo REST API](./examples/golang-todo-rest-crud/README.md): REST API to create, update and retrieve ToDos, including grateful shutdown, rate limiting, structured logging, unit tests, integration tests, environment variables, health check and API documentation with swagger. Technologies: Golang 1.19, MongoDB (with Docker Compose), Gorilla Mux, Go Swagger, Tollbooth (rate limiting), Zap (logging), Viper, Mockery, Makefile, Pre-commit, and DockerTest (integration tests).

## Contributing

Contributions are welcome!

## Contributors

<a href="https://github.com/nanlabs/nancy.go/contributors">
  <img src="https://contrib.rocks/image?repo=nanlabs/nancy.go"/>
</a>

Made with [contributors-img](https://contrib.rocks).

[gocibadge]: https://github.com/nanlabs/nancy.go/actions/workflows/go-ci.yml/badge.svg
[shellcheckbadge]: https://github.com/nanlabs/nancy.go/actions/workflows/shellcheck.yml/badge.svg
[licensebadge]: https://img.shields.io/badge/License-MIT-blue.svg
[gociurl]: https://github.com/nanlabs/nancy.go/actions/workflows/go-ci.yml
[shellcheckurl]: https://github.com/nanlabs/nancy.go/actions/workflows/shellcheck.yml
[licenseurl]: https://github.com/nanlabs/nancy.go/blob/main/LICENSE

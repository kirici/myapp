# myapp

Placeholder app to demo metrics instrumentation.

## Requirements

A few tools are assumed to be available:

- Go - check the `go.mod` file for versions
- Podman, or alternatively Docker
- [direnv.net](https://github.com/direnv/direnv)
- [git-cliff.org](https://github.com/orhun/git-cliff)
- [just.systems](https://github.com/casey/just)
- [vegeta](https://github.com/tsenart/vegeta) - optional, for load testing

## Usage

Copy `.env.sample` to `.env` and adjust the values as needed. If `direnv` is
not available, add `export ` in front of all environment variables and source
the `.env` file in your shell instead.

Then, run:

`just up`

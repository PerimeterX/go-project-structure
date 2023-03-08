# Go Project Structure Example

<img align="right" width="320" alt="go-project-structure-logo" src="https://raw.githubusercontent.com/PerimeterX/go-project-structure/assets/logo.svg">

[![Licence](https://img.shields.io/github/license/perimeterx/go-project-structure)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/perimeterx/go-project-structure)](https://github.com/perimeterx/go-project-structure/releases)
[![Issues](https://img.shields.io/github/issues/perimeterx/go-project-structure?logo=github)](https://github.com/perimeterx/go-project-structure/issues)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

This is an example branch for the Go Project Structure repo.
For information about this repository visit the [main branch](https://github.com/PerimeterX/go-project-structure).

## How To Run This Example Project

First, build it:
```shell
go build ./cmd/todoapp
```
Then, run it:
```shell
./todoapp
```
Then, insert tasks:
```shell
# insert expired task
curl -X POST localhost:8080/task -H 'Content-Type: application/json' -d '{"value":"some value...","date":"1000-01-01T00:00:00Z"}'

# insert future task
curl -X POST localhost:8080/task -H 'Content-Type: application/json' -d '{"value":"some other value...","date":"3000-01-01T00:00:00Z"}'
```
Finally, get all future tasks:
```shell
curl localhost:8080/future_tasks
```

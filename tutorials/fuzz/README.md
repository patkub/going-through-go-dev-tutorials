# fuzz

Followed [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)

Fuzzing automates test inputs in attempt to find inputs that cause crashes, bugs, and vulnerabilities.

Resolve and download dependencies

```sh
go get .
```

Run code in current directory

```sh
go run .
```

Run unit tests

```sh
go test
```

Run tests with fuzzing

```sh
go test -fuzz=Fuzz -fuzztime 30s
```

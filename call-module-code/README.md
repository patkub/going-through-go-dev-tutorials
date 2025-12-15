# call module code

Followed [Tutorial: Call your code from another module](https://go.dev/doc/tutorial/call-module-code)

### Setup

Run the following commands inside `hello` directory. (i.e. `cd hello`)

Resolve and download dependencies

```sh
go get .
```

Edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

```sh
go mod edit -replace example.com/greetings=../greetings
```

Run code in current directory

```sh
go run .
```

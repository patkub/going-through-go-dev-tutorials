# fuzz

Followed [Tutorial: Find and fix vulnerable dependencies with govulncheck](https://go.dev/doc/tutorial/govulncheck)

Resolve and download dependencies

```sh
go get .
```

Example:

```sh
go run . en pl ro
```

Check vulnerabilities:

Install and run govulncheck:
```sh
go install golang.org/x/vuln/cmd/govulncheck@latest
```

```sh
govulncheck ./
```

Downgrade the version of golang.org/x/text to v0.3.5, which contains known vulnerabilities by running:
```sh
go get golang.org/x/text@v0.3.5
```

This should show a vulnerability:

```sh
$ govulncheck ./
=== Symbol Results ===

Vulnerability #1: GO-2021-0113
    Out-of-bounds read in golang.org/x/text/language
  More info: https://pkg.go.dev/vuln/GO-2021-0113
  Module: golang.org/x/text
    Found in: golang.org/x/text@v0.3.5
    Fixed in: golang.org/x/text@v0.3.7
    Example traces found:
      #1: main.go:12:29: vuln.main calls language.Parse

Your code is affected by 1 vulnerability from 1 module.
This scan also found 1 vulnerability in packages you import and 0
vulnerabilities in modules you require, but your code doesn't appear to call
these vulnerabilities.
Use '-show verbose' for more details.
```

Upgrade `golang.org/x/text` to latest to resolve vulnerabilities.
```sh
$ go get golang.org/x/text@latest
go: upgraded golang.org/x/text v0.3.5 => v0.32.0
$ govulncheck ./
No vulnerabilities found.
```

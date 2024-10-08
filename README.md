# Goker

A Texas Hold'em Poker implementation written in Go while following [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests).

## Setup

1. `git clone <url>`
2. `asdf install`
3. `go test ./...`
4. `go run ./cmd/webserver/main.go` OR `go run ./cmd/cli/main.go`

## Useful Commands

```sh
go run <package-path>
go build [<package-path>]

go test [<package-path>][/...] [-v] [-cover] [-race] [-parallel <number>]
go test -bench=. [<package-path>] [-count <number>] [-benchmem] [-benchtime 2s] [-memprofile <name>]

go test -coverprofile <name> [<package-path>]
go tool cover -html <name>
go tool cover -func <name>

go doc [<package-path>]
go fmt [<package-path>]
go vet [<package-path>]

go mod init [<module-path>]
go mod tidy
go mod vendor
go mod download

go work init [<module-path-1> [<module-path-2>] [...]]
go work use [<module-path-1> [<module-path-2>] [...]]
go work sync

# Adjust dependencies in `go.mod`.
go get <package-path>[@<version>]

# Build and install commands.
go install <package-path>[@<version>]

go list -m [all]
```

## Useful Resources

- [Go - Learn](https://go.dev/learn)
- [Go - Documentation](https://go.dev/doc)
- [Go - A Tour of Go](https://go.dev/tour)
- [Go - Effective Go](https://go.dev/doc/effective_go)
- [Go - Playground](https://go.dev/play)
- [Go by Example](https://gobyexample.com)
- [100 Go Mistakes and How to Avoid Them](https://100go.co)

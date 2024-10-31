# Go Study Progress

## Update Go

Download Go from [^1] (see also [^2]), and then:
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
```

For updating the binaries installed via `go`, use `gup` [^3]:
```
go install github.com/nao1215/gup@latest
$GOPATH/bin/gup update --dry-run
$GOPATH/bin/gup help
```

[^1]: https://go.dev/dl/
[^2]: https://go.dev/doc/install
[^3]: https://github.com/nao1215/gup

## A Tour of Go (2024/07/11)

See [tour directory](tour/).

## Documentation/Tutorials/Blog Posts

- [o] The Go Blog
  - [X] [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)
  - [X] [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
  - [X] [How to Write Go Code](how_to_write_go_code/)
  - [X] [Writing Web Applications](gowiki/)
  - [ ] [Error handling and Go](https://go.dev/blog/error-handling-and-go)
  - [ ] [Coverage profiling support for integration tests](https://go.dev/doc/build-cover)
- [ ] Effective Go
  - [ ] [Interfaces and other types](https://go.dev/doc/effective_go#interfaces_and_types)
- [X] Codewalk
  - [X] See [Codewalk: Share Memory By Communicating](share_memory_by_communicating/) directory.

It's also always good to look at the annotated code pieces at [Go By Example](https://gobyexample.com/).

## Learn Go With Tests

See [tests directory](tests/).

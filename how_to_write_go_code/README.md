# How to Write Go Code

[Source](https://go.dev/doc/code)

## Code organization

```
1 Repository consists of Module(s)
|
 - 1 Module := Collection of Packages that are released together
   |
    - 1 Package : = Collection of source files in the same directory that are compiled together
```

Typically:
* 1 Repository consists of 1 Module at the repository's root
* `go.mod` declares the *module path*: the import path prefix for all packages within the module.


## `hello`

```
$ cd ~/src/go_study_progress/how_to_write_go_code/hello

$ cat go.mod
module github.com/stepga/go_study_progress/how_to_write_go_code/hello

go 1.21.4

$ go install github.com/stepga/go_study_progress/how_to_write_go_code/hello

$ cd /tmp

$ hello
hello

$ which hello
/home/stepga/go/bin/hello
```

## `morestrings` package

* create `~/src/go_study_progress/how_to_write_go_code/hello/morestrings`
* verify successful build via `go build` within this new package
* re-install hello as above

This repository contains a small and self-contained Go program called
`build.go`. It used to compile a binary (package `main`) from either a checkout
of the repository, or from an extracted release tar file. This enables
end-users to compile the program without having to setup a `GOPATH`.

For it to function correctly, all dependencies need to be vendored, e.g. with
`dep` or `go mod vendor`. Then, your build does not depend on any third-party
resources on the Internet. For Go >= 1.11, modules are used for building and
`GOPROXY` is set to `off` so that no Internet is needed when building.

The program has a build tag that is not set normally (`ignore_build_go`) so it
is not considered when compiling the other Go code in a repository.

Usage
=====

In order to use it, copy `build.go` to the root level into your repository and
edit the configuration section at the top. You can see an example in the
[restic repository](https://github.com/restic/restic/blob/master/build.go).

Instruct your users to call `go run build.go` and it will produce a binary from
either a checkout of the repository or from an extracted release tar file. For
Go 1.11, it needs to be called as `go run -mod=vendor build.go` so that no
network access is needed.

For cross-compilation, the options `--goos` and `--goarch` can be used, e.g.
like this:
```
$ go run build.go --goos windows --goarch 386
```

The tests can be run by specifying `--test`. By default, `cgo` is explicitly
disabled by passing `CGO_ENABLED=0` to `go build`, it can be re-enabled
manually by running `go run build.go --enable-cgo`.

The program will set the string variable `version` in package `main` to a
version string consisting of the contents of the file `VERSION` (if present)
and the output of `git describe` (if available).

The version string can then be used e.g. in a `version` command, like with
restic:
```
$ ./restic version
restic 0.8.1 (v0.8.1-154-g74665a22)
```

The version string consists of:
 * The contents of the `VERSION` file: `0.8.1`
 * The nearest tag (`v0.8.1`), the number of commits (`154`) and the Git commit hash (`74665a22`)

Background
==========

The program `build.go` constructs a temporary `GOPATH` (for Go < 1.11, or when
no `go.mod` exists) in a temporary directory as configured at the beginning of
the program, then calls `go build` using the temporary `GOPATH`. This means
that end-users do not have to setup a `GOPATH` of their own.

Testing
=======

The tests need to be run with the build tag `ignore_build_go` set:
```
$ go test -tags ignore_build_go
PASS
ok  	github.com/fd0/build-go	0.001s
```

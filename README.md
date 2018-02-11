This repository contains a small and self-contained Go program called
`build.go`. It used to compile a binary (package `main`) from either a checkout
of the repository, or from an extracted release tar file.

The program has a build tag that is not set normally (`ignore_build_go`) so it
is not considered when compiling the other Go code in a repository.

Usage
=====

In order to use it, copy `build.go` at the root level into your repository and
edit the configuration section at the top. You can see an example in the
[restic repository](https://github.com/restic/restic).

Instruct your users to call `go run build.go` and it will produce a binary.

For cross-compilation, the options `--goos` and `--goarch` can be used, e.g.
like this: `go run build.go --goos windows --goarch 386`

The program will set the string variable `version` in package `main` to a
version string consisting of the contents of the file `VERSION` (if present)
and the output of `git describe` (if availabel).

The version string can then be used e.g. in a `version` command, like with
restic:
```
$ ./restic version
restic 0.8.1 (v0.8.1-154-g74665a22)
```

The version string consists of:
 * The contents of the `VERSION` file: `0.8.1`
 * The nearest tag (`v0.8.1`), the number of commits (`154`) and the Git commit hash (`74665a22`)

Testing
=======

The tests need to be run with the build tag `ignore_build_go` set:
```
$ go test -tags ignore_build_go
PASS
ok  	github.com/fd0/build-go	0.001s
```

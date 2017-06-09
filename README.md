# goldflags

Canonical version information injection target for Go's -ldflags.

Use this snippet to inject Git-based version information into binaries:

```Makefile
COMMIT_ID   = $(shell git describe --tags --always --dirty=-dev)
COMMIT_UNIX = $(shell git show -s --format=%ct HEAD)
BUILD_COUNT = $(shell git rev-list --count HEAD)
BUILD_UNIX  = $(shell date +%s)

LDFLAGS = \
	-X 'github.com/digineo/goldflags.commitID=$(COMMIT_ID)' \
	-X 'github.com/digineo/goldflags.commitUnix=$(COMMIT_UNIX)' \
	-X 'github.com/digineo/goldflags.buildCount=$(BUILD_COUNT)' \
	-X 'github.com/digineo/goldflags.buildUnix=$(BUILD_UNIX)'

binary: main.go
	go build -ldflags "$(LDFLAGS)"
```

You can include the `goldflags.mk` to keep your Makefile tidy (cf.
`_examples/Makefile`).

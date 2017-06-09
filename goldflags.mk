# makefile include

COMMIT_ID   ?= $(shell git describe --tags --always --dirty=-dev)
COMMIT_UNIX ?= $(shell git show -s --format=%ct HEAD)
BUILD_COUNT ?= $(shell git rev-list --count HEAD)
BUILD_UNIX  ?= $(shell date +%s)

GOLDFLAGS = \
	-X 'github.com/digineo/goldflags.commitID=$(COMMIT_ID)' \
	-X 'github.com/digineo/goldflags.commitUnix=$(COMMIT_UNIX)' \
	-X 'github.com/digineo/goldflags.buildCount=$(BUILD_COUNT)' \
	-X 'github.com/digineo/goldflags.buildUnix=$(BUILD_UNIX)'

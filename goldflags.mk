# makefile include

COMMIT_ID   ?= $(shell git describe --tags --always --dirty=-dev)
COMMIT_UNIX ?= $(shell git show -s --format=%ct HEAD)
BUILD_COUNT ?= $(shell git rev-list --count HEAD)
BUILD_UNIX  ?= $(shell date +%s)

GOLDFLAGS = -s -w \
	-X 'github.com/digineo/goldflags.commitID=$(COMMIT_ID)' \
	-X 'github.com/digineo/goldflags.commitUnix=$(COMMIT_UNIX)' \
	-X 'github.com/digineo/goldflags.buildCount=$(BUILD_COUNT)' \
	-X 'github.com/digineo/goldflags.buildUnix=$(BUILD_UNIX)'

# Build a platform specific version of $target. It creates a
# "$(basename $target)" (or "$(basename).exe" for Windows) in the
# directory "$(dirname $target)/$(goos)-$(goarch)", e.g.
# "bin/linux-amd64/main" for target=bin/main.
#
# call-seq:
#	$(call goldbuild,target,goos,goarch)
define goldbuild
	@mkdir -p "$(dir $(1))/$(2)-$(3)"
	GOOS="$(2)" GOARCH="$(3)" go build \
		-ldflags "$(GOLDFLAGS)" \
		-o "$(dir $(1))/$(2)-$(3)/$(notdir $(1))$(if $(filter windows,$(2)),.exe)"
endef

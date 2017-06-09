include $(GOPATH)/src/github.com/digineo/goldflags/goldflags.mk

.PHONY: run
run: bin/main
	./bin/main

bin/main:
	go build -ldflags "-s -w $(GOLDFLAGS)" -o bin/main

.PHONY: multi
multi:
	$(call goldbuild,bin/main,linux,amd64)
	$(call goldbuild,bin/main,linux,arm64)
	$(call goldbuild,bin/main,darwin,amd64)
	$(call goldbuild,bin/main,windows,amd64)
	$(call goldbuild,bin/main,freebsd,amd64)

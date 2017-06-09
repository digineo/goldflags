include $(GOPATH)/src/github.com/digineo/goldflags/goldflags.mk

.PHONY: run
run: main
	./main

main:
	go build -ldflags "-s -w $(GOLDFLAGS)" -o main

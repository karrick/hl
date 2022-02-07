PROJECT = hl
MAN = $(PROJECT).1
MANGZ = $(MAN).gz

GO ?= go

build: $(PROJECT) $(MANGZ)

clean:
	rm -f $(PROJECT) $(MANGZ)

lint: $(MAN)
	$(GO) vet ./...
	mandoc -T lint $(MAN)

$(PROJECT):
	$(GO) build -o $@

$(MANGZ): $(MAN)
	gzip -9c $(MAN) > $@

.PHONY: build clean lint

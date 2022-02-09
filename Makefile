PROJECT = hl
MDOC = $(PROJECT).1
MAN = $(MDOC).gz
MARKDOWN = README.md

GO ?= go

build: $(PROJECT) $(MAN) $(MARKDOWN)

clean:
	rm -f $(PROJECT) $(MAN) $(MARKDOWN)

lint: $(MDOC)
	$(GO) vet ./...
	mandoc -T lint $(MDOC)

$(PROJECT):
	$(GO) build -o $@

$(MAN): $(MDOC)
	gzip -9c $(MDOC) > $@

# NOTE: The following multi-command pipeline produces slightly better Markdown
# when using a mdoc(7) source file. It uses mandoc(1) to convert mdoc(7) text
# to man(7) text, and uses pandoc(1) to convert that to Markdown.
$(MARKDOWN): $(MDOC)
	mandoc -T man $(MDOC) | pandoc --from man --to gfm --fail-if-warnings -s -o $@

.PHONY: build clean lint

SOL_BINS ?= $(shell find ./sols -name *.sol 2> /dev/null | sed s/\\.sol/.bin/)

.PHONY: all clean
all: gucumber $(SOL_BINS)
	@gucumber

%.bin: %.sol
	solc --hashes --bin -o sols/ $<

gucumber:
	@go get github.com/gucumber/gucumber/cmd/gucumber

clean:
	@find sols -regex '.*bin\|.*signatures' -exec rm -f '{}' \;

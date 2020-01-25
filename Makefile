outfile = step-generator

.PHONY: cli
cli:
	@sh ./scripts/cli.sh

.PHONY: build
build:
	@sh ./scripts/build.sh


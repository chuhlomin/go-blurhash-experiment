.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: vet
## vet: runs the go vet command
vet:
	@go vet ./...

.PHONY: build-encoder
## build-encoder: builds an Encoder binary
build-encoder:
	@go build -o encoder cmd/encoder/encoder.go

.PHONY: build-decoder
## build-decoder: builds a Decoder binary
build-decoder:
	@go build -o decoder cmd/decoder/decoder.go

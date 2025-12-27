CI_RAW_URL = https://raw.githubusercontent.com/yeencloud/dpl-ci/refs/heads/main

update:
	mkdir -p .github/scripts
	curl -O $(CI_RAW_URL)/makefile \
         -O $(CI_RAW_URL)/.golangci.yml \
		 -O $(CI_RAW_URL)/.github/scripts/openapi-generate.sh \
		 -O $(CI_RAW_URL)/.github/config/oapi-codegen.yml
lint:
	golangci-lint run ./...

test:
	go test -race -v ./...

openapi:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I contract/proto \
      --go_out=contract/proto/generated --go_opt=paths=source_relative \
      --go-grpc_out=contract/proto/generated --go-grpc_opt=paths=source_relative \
      contract/proto/*.proto
	find contract/proto/generated -name '*.pb.go' -not -name '*_grpc.pb.go' \
	  -exec sh -c 'f="$$1"; mv "$$f" "$${f%.pb.go}.gen.go"' _ {} \;
	find contract/proto/generated -name '*_grpc.pb.go' \
	  -exec sh -c 'f="$$1"; mv "$$f" "$${f%_grpc.pb.go}_grpc.gen.go"' _ {} \;

air:
	@printf "\033]0;%s\007" "$(notdir $(CURDIR))"
	@go tool air --build.cmd "go build -o ./tmp/$(notdir $(CURDIR))exe cmd/main.go" \
	  --build.entrypoint "./tmp/$(notdir $(CURDIR))exe" \
	  --build.include_file ".env" \
	  --build.include_dir "../../libs"

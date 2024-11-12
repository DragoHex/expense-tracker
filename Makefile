BINARY_NAME=extr
SQLC_BINARY_NAME=extr-sqlc

.PHONY: bin
bin:
	CGO_ENABLED=1 go build -o artifacts/${BINARY_NAME} cmd/main.go

.PHONY: bin-sqlc
bin-sqlc:
	CGO_ENABLED=1 go build -o artifacts/${SQLC_BINARY_NAME} cmd/extr-sqlc/main.go

.PHONY: run
run:
	go run cmd/main.go


IMAGE_NAME=go-expense-tracker-builder
CONTAINER_NAME=builder-container

.PHONY: builds
builds:
	echo $(shell pwd)
	make bin
	docker build -t $(IMAGE_NAME) -f docker/Build-Container .
	-docker run --name $(CONTAINER_NAME) -v $(CURDIR)/artifacts:/app/artifacts $(IMAGE_NAME)
	echo "Builds created successfully"
	echo "Cleaning up docker images..."
	make clean-container
	# docker rmi -f go-expense-tracker-builder
	# docker image prune -f

# to be used only inside the build container
.PHONY: bins
bins:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags '-static'" -o artifacts/extr-linux-arm64 cmd/main.go
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags '-static'" -o artifacts/extr-linux-amd64 cmd/main.go
	# GOOS=windows GOARCH=arm64 CGO_ENABLED=1 go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags '-static'" -o artifacts/extr-windows-arm64 cmd/main.go
	# GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -tags "sqlite_omit_load_extension" -ldflags "-extldflags '-static'" -o artifacts/extr-windows-amd64 cmd/main.go

# Clean up the container after the build is complete
.PHONY: clean-container
clean-container:
	@echo "Cleaning up container and image"
	@if docker ps -q -f name=$(CONTAINER_NAME); then \
		docker rm $(CONTAINER_NAME); \
	fi
	@docker rmi $(IMAGE_NAME)

.PHONY: clean
clean:
	go clean
	@rm -r artifacts 2> /dev/null

${V}.SILENT:

BINARY_NAME=GoldWatcher
SOURCE_DIR=./cmd/goldwatcher
FYNE_PACKAGE=github.com/fyne-io/fyne/v2/cmd/fyne

build:
	@echo "Building $(BINARY_NAME) with Fyne..."
	@rm -f ${SOURCE_DIR}/*.tar.xz
	@cd ${SOURCE_DIR} && fyne package -os linux --release

run:
	echo "Running $(BINARY_NAME)..."
	@env DB_PATH="./sql.db" go run $(SOURCE_DIR)

install:
	@echo "Building $(BINARY_NAME) with Fyne..."
	@cd ${SOURCE_DIR} && sudo fyne install

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f ${SOURCE_DIR}/*.tar.xz
	@echo "Cleaned!"

test:
	go test -v ./...

.PHONY: build run clean test

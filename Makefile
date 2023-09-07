# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_DIR=bin
BUILD_FLAGS=-v
CLEAN_CMD = del /Q $(BINARY_DIR)\*.* && rmdir /S /Q $(BINARY_DIR)

# Targets
all: build

build:
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIR)/ssm.exe ./cmd/ssm

install:
	$(GOCMD) install $(BUILD_FLAGS) ./cmd/ssm

clean:
	$(GOCLEAN)
	$(CLEAN_CMD)

.PHONY: clean
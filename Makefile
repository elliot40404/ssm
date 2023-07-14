# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_DIR=bin
BUILD_FLAGS=-v

ifeq ($(OS),Windows_NT)
    CLEAN_CMD = del /Q $(BINARY_DIR)\*.* && rmdir /S /Q $(BINARY_DIR)
else
    CLEAN_CMD = rm -rf $(BINARY_DIR)
endif

# Targets
all: windows

windows:
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIR)/ssm.exe main.go

clean:
	$(CLEAN_CMD)

.PHONY: clean
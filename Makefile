PROJECT_PATH := $(abspath .)

TOOLS_DIR := $(abspath ./.tools)
TOOLS_MOD_DIR := $(abspath ./tools)

$(TOOLS_DIR)/gofumpt: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum
	cd $(TOOLS_MOD_DIR) && go build -o $(TOOLS_DIR)/gofumpt.exe mvdan.cc/gofumpt

$(TOOLS_DIR)/gci: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum
	cd $(TOOLS_MOD_DIR) && go build -o $(TOOLS_DIR)/gci.exe github.com/daixiang0/gci

$(TOOLS_DIR)/golang-lint: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum
	cd $(TOOLS_MOD_DIR) && go build -o $(TOOLS_DIR)/golang-lint.exe github.com/golangci/golangci-lint/v2/cmd/golangci-lint

fumpt: $(TOOLS_DIR)/gofumpt
	$(TOOLS_DIR)/gofumpt.exe -w .

imports: $(TOOLS_DIR)/gci
	$(TOOLS_DIR)/gci.exe write -s standard -s default -s "prefix(github.com)" .

lint: $(TOOLS_DIR)/golang-lint
	$(TOOLS_DIR)/golang-lint.exe run --config $(PROJECT_PATH)/.golangci.yaml


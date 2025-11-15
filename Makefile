BINARY := basic_server.exe
SRC_DIR := src 
SRC := $(shell find $(SRC_DIR) -type f -name '*.go')

DEBUG ?=0

ifeq ($(DEBUG), 1)
	BUILD_FLAGS := -gcflags="all=-N -l"
	BUILD_MODE := debug
else
	BUILD_FLAGS :=
	BUILD_MODE := release
endif

all:build 

build: $(SRC)
	echo "Building $(BINARY) ($(BUILD_MODE));..."
	cd $(SRC_DIR) && go build $(BUILD_FLAGS) -o ../$(BINARY)

clean:
	rm -f $(BINARY)
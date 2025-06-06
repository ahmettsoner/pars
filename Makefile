CONFIG_FILE := build/$(if $(VERSION),$(VERSION),current)/config.mk
ifneq ("$(wildcard $(CONFIG_FILE))","")
    -include $(CONFIG_FILE)
endif

define build_cmake
	@cmake -B $(1) -S . -DVERSION=$(2) -DRELEASE_DATE=$(RELEASE_DATE) $(4) $(3)
	@echo CONFIG_VERSION=$(2) > $(1)/config.mk
	@echo CONFIG_RELEASE_DATE=$(RELEASE_DATE) >> $(1)/config.mk
endef


KEY_LIST := OUTPUT
@$(foreach key,$(KEY_LIST),echo "$(key): $($(key))";)
COMMAND_LIST := $(foreach key,$(KEY_LIST),"-D$(key)=$($(key))")

build.cmake.linux:
	$(call build_cmake,build/current,$(VERSION),-G "Unix Makefiles",$(COMMAND_LIST))

build.cmake.linux.%:
	$(call build_cmake,build/$*,$*,-G "Unix Makefiles",$(COMMAND_LIST))

build.cmake.macos:
	$(call build_cmake,build/current,$(VERSION),-G "Unix Makefiles",$(COMMAND_LIST))

build.cmake.macos.%:
	$(call build_cmake,build/$*,$*,-G "Unix Makefiles",$(COMMAND_LIST))

build.cmake.windows:
	$(call build_cmake,build/current,$(VERSION),-G "MinGW Makefiles",$(COMMAND_LIST))

build.cmake.windows.%:
	$(call build_cmake,build/$*,$*,-G "MinGW Makefiles",$(COMMAND_LIST))


ifneq ($(OS),Windows_NT)
UNAME_S := $(shell uname -s)
endif

build.cmake:
ifeq ($(OS),Windows_NT)
	$(MAKE) build.cmake.windows $(MAKEOVERRIDES)
else ifeq ($(UNAME_S),Linux)
	$(MAKE) build.cmake.linux $(MAKEOVERRIDES)
else ifeq ($(UNAME_S),Darwin)
	$(MAKE) build.cmake.macos $(MAKEOVERRIDES)
else
	$(error "Unsupported OS")
endif

build.cmake.%:
ifeq ($(OS),Windows_NT)
	$(MAKE) build.cmake.windows.$* $(MAKEOVERRIDES)
else ifeq ($(UNAME_S),Linux)
	$(MAKE) build.cmake.linux.$* $(MAKEOVERRIDES)
else ifeq ($(UNAME_S),Darwin)
	$(MAKE) build.cmake.macos.$* $(MAKEOVERRIDES)
else
	$(error "Unsupported OS")
endif


%:
	@echo "Building for version: $(if $(VERSION),$(VERSION),current)"
ifeq ($(VERSION),)
	$(MAKE) build.cmake VERSION=$(if $(VERSION),$(VERSION),$(CONFIG_VERSION)) $(MAKEOVERRIDES)
else
	$(MAKE) build.cmake.$(if $(CONFIG_VERSION),$(CONFIG_VERSION),$(VERSION)) $(MAKEOVERRIDES)
endif
	$(MAKE) -C build/$(if $(VERSION),$(VERSION),current) $@

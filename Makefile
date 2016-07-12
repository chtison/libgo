PACKAGES_ROOT := github.com/chtison/libgo
PACKAGES := $(addprefix $(PACKAGES_ROOT), \
			/shellcolors/cmd/sc \
			/baseconverter \
		)

GO_INSTALL := go install
GO_CLEAN   := go clean -i
GO_DOC     := godoc -http=':6060'
GO_DOC_URL := http://127.0.0.1:6060/pkg/
GO_TEST    := go test -v ./...

main:
	$(info $(usage))
	@true
define usage
usage: make [command]
make list                # print all go packages managed by this makefile
make doc                 # start a godoc server on ':6060'
make test                # run tests on all go packages
make install             # install all go packages
make clean               # remove all go packages
make re                  # make clean && make install
endef

list:
	$(call list)
define list
@for package in $(PACKAGES); do echo $$package ; done
endef

doc:
	@echo "$(GO_DOC_URL)$(PACKAGES_ROOT)"
	$(GO_DOC)

test:
	$(GO_TEST)

install:
	$(call install)
define install
@for package in $(PACKAGES) ; do $(GO_INSTALL) $$package ; done
endef

clean:
	$(call clean)
define clean
@for package in $(PACKAGES) ; do $(GO_CLEAN) $$package ; done
endef

re: clean install

.PHONY: main list doc install clean re

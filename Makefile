PACKAGES_ROOT := github.com/chtison/libgo
PACKAGES := $(addprefix $(PACKAGES_ROOT), \
			/shellcolors/cmd/sc \
		)

GO_INSTALL := go install
GO_CLEAN   := go clean -i
GO_DOC     := godoc -http=':6060'
GO_DOC_URL := http://127.0.0.1:6060/pkg/

define usage
usage: make [command]
make list                # print all go packages managed by this makefile
make doc                 # starts a godoc server on ':6060'
make install             # install all go packages
make clean               # remove all go packages
make re                  # make clean && make install
endef

main:
	$(info $(usage))
	@true

list:
	$(foreach package, $(PACKAGES), @echo $(package))

doc:
	@echo "$(GO_DOC_URL)$(PACKAGES_ROOT)"
	$(GO_DOC)


install:
	$(foreach package, $(PACKAGES), $(GO_INSTALL) $(package))

clean:
	$(foreach package, $(PACKAGES), $(GO_CLEAN) $(package))

re: clean install

.PHONY: main list doc install clean re

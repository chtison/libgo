.PHONY: generate

GENERATE  := fmt github.com/go-yaml/yaml
GENERATED := $(addsuffix /generated.go, $(notdir $(GENERATE)))

generate: $(GENERATED)
	$(MAKE) -C tmpl generate

$(GENERATED): generate.go
	$(eval PKGPATH := $(filter %$(subst /,,$(dir $@)), $(GENERATE)))
	go get $(PKGPATH)
	go run generate.go $(PKGPATH)
	gofmt -w $@

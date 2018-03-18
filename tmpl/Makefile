.PHONY: generate

GENERATE  := fmt strings time encoding/json github.com/go-yaml/yaml
GENERATED := $(addprefix generated/, $(addsuffix .go, $(notdir $(GENERATE))))

generate: $(GENERATED)

$(GENERATED): generate.go
	$(eval PKGPATH := $(filter %$(basename $(notdir $@)), $(GENERATE)))
	go run generate.go $(PKGPATH)
	goimports -w $@

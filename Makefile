.PHONY: gen

gen: schema.json webgpu.yml
	go run ./gen -gen-c-header -schema schema.json -yaml webgpu.yml -header webgpu.h

gen-docs: schema.json webgpu.yml
	go run ./gen -gen-docs -schema schema.json -yaml webgpu.yml -docs doc/src

gen-check: gen
	@git diff --quiet -- webgpu.h || { git diff -- webgpu.h; exit 1; }

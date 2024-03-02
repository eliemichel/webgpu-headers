package main

import (
	"flag"
	"os"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"

	"embed"

	"github.com/goccy/go-yaml"
)

//go:embed templates/c-header/webgpu.h.tmpl
var cheaderTmpl string

//go:embed templates/doc/*.md.tmpl
var docsTmpl embed.FS
var docsTmplRoot string = "templates/doc"

var (
	schemaPath string
	headerPath string
	yamlPath   string
	docsPath   string
	genDocs    bool
	genCAPI    bool
)

func main() {
	flag.StringVar(&yamlPath, "yaml", "", "path of the yaml spec")
	flag.StringVar(&schemaPath, "schema", "", "path of the json schema")
	flag.StringVar(&headerPath, "header", "", "output path of the header")
	flag.StringVar(&docsPath, "docs", "", "output directory where to write documentation source")
	flag.BoolVar(&genCAPI, "gen-c-header", false, "generate C API header")
	flag.BoolVar(&genDocs, "gen-docs", false, "generate documentation source")
	flag.Parse()
	if schemaPath == "" || yamlPath == "" || (genCAPI && headerPath == "") || (genDocs && docsPath == "") {
		flag.Usage()
		os.Exit(1)
	}

	if err := ValidateYaml(schemaPath, yamlPath); err != nil {
		panic(err)
	}

	src, err := os.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}

	var yml Yml
	if err := yaml.Unmarshal(src, &yml); err != nil {
		panic(err)
	}

	SortAndTransform(&yml)

	fileName := filepath.Base(yamlPath)
	fileNameSplit := strings.Split(fileName, ".")
	if len(fileNameSplit) != 2 {
		panic("got invalid file name: " + fileName)
	}

	var data Data
	data.Yml = &yml
	data.Name = fileNameSplit[0]

	if genCAPI {
		dst, err := os.Create(headerPath)
		if err != nil {
			panic(err)
		}

		if err := GenCHeader(&data, dst); err != nil {
			panic(err)
		}
	}

	if genDocs {
		if err := GenAllDocSources(&data, docsPath); err != nil {
			panic(err)
		}
	}
}

func SortAndTransform(yml *Yml) {
	// Sort structs
	SortStructs(yml.Structs)
	// Sort constants
	slices.SortStableFunc(yml.Constants, func(a, b Constant) int {
		return strings.Compare(PascalCase(a.Name), PascalCase(b.Name))
	})
	// Sort enums
	slices.SortStableFunc(yml.Enums, func(a, b Enum) int {
		return strings.Compare(PascalCase(a.Name), PascalCase(b.Name))
	})
	// Sort bitflags
	slices.SortStableFunc(yml.Bitflags, func(a, b Bitflag) int {
		return strings.Compare(PascalCase(a.Name), PascalCase(b.Name))
	})

	for _, s := range yml.Structs {
		if s.FreeMembers {
			yml.Objects = append(yml.Objects, Object{
				IsStruct: true,
				Name:     s.Name,
				Methods:  []Function{{Name: "free_members"}},
			})
		}
	}

	// Sort objects
	slices.SortStableFunc(yml.Objects, func(a, b Object) int {
		return strings.Compare(PascalCase(a.Name), PascalCase(b.Name))
	})
	// Sort methods
	for _, obj := range yml.Objects {
		slices.SortStableFunc(obj.Methods, func(a, b Function) int {
			return strings.Compare(PascalCase(a.Name), PascalCase(b.Name))
		})
	}
}

func GenAllDocSources(data *Data, rootPath string) error {
	return fs.WalkDir(docsTmpl, docsTmplRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}

		tmpl, err := docsTmpl.ReadFile(path)
		if err != nil {
			panic(err)
		}

		dstName := path[len(docsTmplRoot):len(path)-len(".tmpl")]
		dst, err := os.Create(rootPath + dstName)
		if err != nil {
			panic(err)
		}

		return GenDocSource(data, string(tmpl), dst)
	})
}

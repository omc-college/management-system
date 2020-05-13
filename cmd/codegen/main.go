package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/sirupsen/logrus"
)

func main() {
	var buffer = bytes.NewBufferString("")
	var err error
	var prefix, suffix string
	suffix = "/"
	prefix = ".go"

	specPath := flag.String("spec-path", "", "please use --spec-path /path/to/files/located")
	outputPath := flag.String("output-path", "", "please use --output-path show/indicate/point/go")

	flag.Parse()

	if specPath == nil || *specPath == "" {
		logrus.Fatalf("Please define OpenApi location  --spec-path")

	}
	if outputPath == nil || *outputPath == "" {
		logrus.Fatalf("Please define directory --output-path")

	}

	absPath, err := filepath.Abs(*specPath)
	if err != nil {
		logrus.Fatalf("cannot get absolute path of spec-path directory: %s", err.Error())
	}

	outputDirSlice := strings.Split(filepath.Dir(absPath), "/")

	outputDirName := outputDirSlice[len(outputDirSlice)-1]
	fmt.Fprintf(buffer, "package %s", outputDirName)

	specBytes, err := ioutil.ReadFile(*specPath)
	if err != nil {
		logrus.Fatalf("cannot read file: %s", err.Error())
	}

	loader := openapi3.NewSwaggerLoader()

	doc, err := loader.LoadSwaggerFromData(specBytes)
	if err != nil {
		logrus.Fatalf("cannot parse OpenApi spec: %s", err.Error())
	}

	generateNewTypes(doc, buffer)

	formatedTypesSource, err := format.Source(buffer.Bytes())
	if err != nil {
		logrus.Fatalf("cannot  format source types: %s", err.Error())
	}

	if _, err := os.Stat(*outputPath); os.IsNotExist(err) {
		os.MkdirAll(*outputPath, 0700)
	}
	if err != nil {
		logrus.Fatalf("cannot  create new folder: %s", err.Error())
	}

	fileName := filepath.Base(*specPath)
	if err != nil {
		logrus.Fatalf("cannot returns the last element of path: %s", err.Error())
	}

	outputName := strings.Split(strings.TrimSuffix(fileName, filepath.Ext(fileName)), "-")

	outputFileName := suffix + outputName[0] + prefix

	newTypesFile, err := os.Create(*outputPath + outputFileName )
	if err != nil {
		logrus.Fatalf("cannot create the named file: %s", err.Error())
	}
	defer newTypesFile.Close()

	newTypesFile.WriteString(string(formatedTypesSource))
}

func generateNewTypes(doc *openapi3.Swagger, buffer *bytes.Buffer) {
	for name, schema := range doc.Components.Schemas {
		generateStruct(name, schema, buffer)
	}
}

func generateStruct(name string, schema *openapi3.SchemaRef, buffer *bytes.Buffer) {
	fmt.Fprintf(buffer, "\ntype %s struct {\n", name)
	if schema.Value == nil {
		return
	}

	for fieldName, field := range schema.Value.Properties {
		generateField(fieldName, field, buffer)
	}
	fmt.Fprintf(buffer, "}\n")
}

func generateField(name string, field *openapi3.SchemaRef, buffer *bytes.Buffer) {
	goType := resolveGoType(field.Value, buffer)
	name = strings.ToUpper(name[:1]) + name[1:]
	fmt.Fprintf(buffer, "%s %s %s\n", name, goType, generateJsonTag(name))
}

func resolveGoType(v *openapi3.Schema, buffer *bytes.Buffer) string {
	switch v.Type {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "int32":
		return "int"
	case "boolean":
		return "bool"
	case "int64":
		return "int"
	case "object":
		return "interface{}"
	case "array":
		items := strings.Split(v.Items.Ref, "/")
		return fmt.Sprintf("[]%s", items[len(items)-1])
	default:
		panic("unsupported type")
	}
}

func generateJsonTag(name string) string {
	name = strings.ToLower(name[:1]) + name[1:]
	return fmt.Sprintf("`json:\"%s\"`", name)
}

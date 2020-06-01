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
)

var buffer = bytes.NewBufferString("")

func main() {
	var specPath = flag.String("spec-path", "/home/masha/go/src/github.com/omc-college/management-system/api/rbac/rbac-api.yaml", "please use --spec-path /path/to/files/located")
	var outputPath = flag.String("output-path", "example.go", "please use --output-path /path/to/files/located")

	flag.Parse()

	absPath, err := filepath.Abs(*outputPath)
	if err != nil {
		panic(err)
	}

	outputDirSlice := strings.Split(filepath.Dir(absPath), "/")

	outputDirName := outputDirSlice[len(outputDirSlice)-1]
	fmt.Fprintf(buffer, "package %s", outputDirName)

	specBytes, err := ioutil.ReadFile(*specPath)
	if err != nil {
		panic(err)
	}
	loader := openapi3.NewSwaggerLoader()
	doc, err := loader.LoadSwaggerFromData(specBytes)
	if err != nil {
		panic(err)
	}

	generatenewypes(doc)

	p, err := format.Source(buffer.Bytes())
	if err != nil {
		panic(err)
	}

	file, err := os.Create(*outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(string(p))

}

func generatenewypes(doc *openapi3.Swagger) {
	for name, schema := range doc.Components.Schemas {
		generatestruct(name, schema)
	}

}

func generatestruct(name string, schema *openapi3.SchemaRef) {
	fmt.Fprintf(buffer, "\ntype %s struct {\n", name)
	if schema.Value == nil {
		return
	}

	for fieldName, field := range schema.Value.Properties {
		generatefield(fieldName, field)
	}

	fmt.Fprintf(buffer, "}\n")
}

func generatefield(name string, field *openapi3.SchemaRef) {
	gotype := resolvegotype(field.Value)
	name = strings.ToUpper(name[:1]) + name[1:]
	fmt.Fprintf(buffer, "%s %s %s\n", name, gotype, generatejsontag(name))
}

func resolvegotype(v *openapi3.Schema) string {
	switch v.Type {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "array":
		items := strings.Split(v.Items.Ref, "/")
		return fmt.Sprintf("[]%s", items[len(items)-1])
	default:
		panic("unsupported type")
	}

}

func generatejsontag(name string) string {
	name = strings.ToLower(name)
	return fmt.Sprintf("`json:\"%s\"`", name)
}

//	for path, pathItem := range doc.Paths {
//		//if pathItem.Parameters != nil {
//		//	fmt.Printf("ref: %v\n", pathItem.Parameters)
//		//}
//		if pathItem == nil {
//			continue
//		}
//
//		printOperation(path, pathItem.Get, "GET")
//		printOperation(path, pathItem.Post, "POST")
//		printOperation(path, pathItem.Delete, "DELETE")
//		printOperation(path, pathItem.Put, "PUT")
//		printOperation(path, pathItem.Connect, "CONNECT")
//		printOperation(path, pathItem.Head, "HEAD")
//		printOperation(path, pathItem.Options, "OPTIONS")
//		printOperation(path, pathItem.Patch, "PATCH")
//		printOperation(path, pathItem.Trace, "TRACE")
//	}
//
//
//}
//func printOperation(path string, operation *openapi3.Operation, httpMethod string) {
//	if operation == nil{
//		return
//	}
//
//	var result = fmt.Sprintf( "router.HandleFunc(\"%s%s%v\n",path +"\", rolesHandler.",operation.OperationID+").Methods(http.Method",httpMethod+")")
//	fmt.Println(result)
//}

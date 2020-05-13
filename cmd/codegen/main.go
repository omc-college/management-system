package main

import (
	"flag"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"io/ioutil"
	"strings"
)


func main() {
	//// String defines a string flag with specified name, default value, and usage string.
	//// The return value is the address of a string variable that stores the value of the flag.
	var specPath = flag.String("spec-path", "", "please use --spec-path /path/to/files/located")
	flag.Parse()
	//// ReadFile reads the file named by filename and returns the contents.
	specBytes, err := ioutil.ReadFile(*specPath)
	if err != nil {
		panic(err)
	}
	loader := openapi3.NewSwaggerLoader()
	doc, err := loader.LoadSwaggerFromData(specBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println( " type Endpoint struct {")
	for k,v:= range doc.Components.Schemas{
	  if k != "Endpoint" {
	  	continue
	  }
	  if v.Ref != "" {
		   fmt.Printf("ref: %v\n\t", v.Ref)
	  }
	  if v.Value != nil {
		   for k, v := range v.Value.Properties {
			   fmt.Printf(" %s%s\t", strings.ToUpper(k[:1]),k[1:])
			     if v.Value != nil {
					 }
			   fmt.Printf(" %+v\t", v.Value.Type)
			   fmt.Printf(" %+v\n\t", v.Value.Description)

				 }

		   }
	  }

	fmt.Println( "}")


	fmt.Println( "type FeatureEntry struct {")
	for k, v:= range doc.Components.Schemas{

		if k != "FeatureEntry" {
			  	continue
		}

		if v.Ref != "" {
			   fmt.Printf("ref: %v\n\t", v.Ref)
		}

		if v.Value != nil {
				   for k, v := range v.Value.Properties {
					   fmt.Printf(" %s%s\t", strings.ToUpper(k[:1]),k[1:])
					     if v.Value != nil {
							 }
							  if v.Value.Type == "integer" {
								  fmt.Printf("int\t")
								  fmt.Printf(" %+v\n\t", v.Value.Description)
							  } else {
								  fmt.Printf(" %+v\t", v.Value.Type)
								  fmt.Printf(" %+v\n\t", v.Value.Description)
							  }

						 }

				   }
			   }
	fmt.Println( "}")


	fmt.Println( "type Role struct {")
	for k, v := range doc.Components.Schemas {

		if k != "Role" {
			continue
		}

		if v.Ref != "" {
			fmt.Printf("ref: %v", v.Ref)
		}

		if v.Value != nil {
			for k, v := range v.Value.Properties {
				fmt.Printf("%s%s\t", strings.ToUpper(k[:1]), k[1:])
				if v.Value != nil {
				}
				fmt.Printf("%v\t", v.Value.Type)
				fmt.Printf(" %+v\n", v.Value.Description)

			}

		}
	}
	fmt.Println( "}")
}
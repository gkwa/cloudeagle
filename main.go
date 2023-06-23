package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

const tfTemplate = `provider "aws" {
  region = "us-west-2"  # Replace with your desired AWS region
}

resource "aws_key_pair" "example" {
  key_name   = "{{.KeyName}}"
  public_key = file("{{.PublicKeyPath}}")
}

output "key_pair_name" {
  value = aws_key_pair.example.key_name
}
`

type TemplateData struct {
	PublicKeyPath string
	KeyName       string
}

func main() {
	var publicKeyPath string
	var keyName string
	flag.StringVar(&publicKeyPath, "publicKeyPath", "", "Path to the public key file")
	flag.StringVar(&keyName, "keyName", "", "Key name")
	flag.Parse()

	if publicKeyPath == "" {
		fmt.Println("Error: publicKeyPath flag is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read the template file
	tmpl := template.Must(template.New("testTfTemplate").Parse(tfTemplate))

	// Create the output file
	outputFile, err := os.Create("test.tf")
	if err != nil {
		fmt.Printf("Error creating test.tf: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Fill in the template with the provided data
	data := TemplateData{PublicKeyPath: publicKeyPath, KeyName: keyName}
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("test.tf file created successfully.")
}

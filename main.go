package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

type InputData struct {
	Maintainer           string                 `yaml:"maintainer"`
	BaseImage            string                 `yaml:"base_image"`
	LocalSourceDirectory string                 `yaml:"local_source_directory"`
	ServicePorts         []ServicePorts         `yaml:"service_ports"`
	EnvironmentVariables []EnvironmentVariables `yaml:"environment_variables"`
	ExecCommand          string                 `yaml:"exec_command"`
	WorkingDirectory     string                 `yaml:"working_directory"`
	AdditionalSteps      string                 `yaml:"additional_steps"`
	BinaryName           string                 `yaml:"binary_name"`
}
type ServicePorts struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type EnvironmentVariables struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func main() {
	// flag for type of build
	applicationLanguage := flag.String("L", "go", "language of application")
	// flag for the input yaml file
	inputYamlFile := flag.String("I", "input.yaml", "input yaml file")
	flag.Parse()

	// read the input yaml file
	input_data, err := ioutil.ReadFile(*inputYamlFile)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var data InputData
	err = yaml.Unmarshal(input_data, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal yaml: %s", err)
	}

	// Define a template.FuncMap object.
	//This is a map of functions that can be called from within the template.

	customFunctions := template.FuncMap{
		"cmdSmush": commandLineToDockerFormattedCmd,
	}

	dockerfileTemplateFile := fmt.Sprintf("templates/%s.go.tmpl", *applicationLanguage)
	template_name := fmt.Sprintf("%s.go.tmpl", *applicationLanguage)
	dockerfile_template, err := template.New(template_name).Funcs(customFunctions).ParseFiles(dockerfileTemplateFile)

	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
	}

	// create an Output file for the created Dockerfile
	// generate a filename based on the language and current datetime
	datetime := time.Now().Format("2006-01-02-1504") // Gosh aren't go's date formatting strings fun!
	output_file_path := fmt.Sprintf("generated/%s.%s.Dockerfile", *applicationLanguage, datetime)
	log.Println("Writing to file: ", output_file_path)

	output_file, err := os.Create(output_file_path)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer output_file.Close()

	// Execute the template, passing in the data and writing to stdout
	err = dockerfile_template.Execute(output_file, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
}

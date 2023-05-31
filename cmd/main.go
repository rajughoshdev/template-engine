package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rajughoshdev/template-engine/utils"
)

func main() {
	filePath := flag.String("config", "", "Path to the YAML file")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please provide the path to the YAML file using the -file flag.")
	}

	configData, err := utils.ReadYAMLFile(*filePath)
	utils.HandleError(err)

	fmt.Println("cloning repo.. ")
	projectName, _ := configData["PROJECT_NAME"].(string)

	projectDir := "./" + projectName

	templateName, _ := configData["TEMPLATE_NAME"].(string)

	repo := utils.CloneRepository(projectDir, templateName)
	utils.HandleError(repo)

	if err = utils.RemoveGitDirectory(projectDir); err != nil {
		panic(err)
	}

	fmt.Println("Generating file as per config")

	envs, ok := configData["ENVS"].(map[interface{}]interface{})
	if !ok {
		log.Fatalf("Invalid 'ENVS' section in YAML")
	}
	vendorsName := []string{"aws", "oci", "gcloud", "alicloud", "azure"}

	activeVendor, ok := envs["ENV_CLOUD_PROVIDER"].(string)

	if !ok {
		log.Fatalf("ENV_CLOUD_PROVIDER not found in YAML")
	}

	for _, vendor := range vendorsName {
		if vendor != activeVendor {
			dirName := projectDir + "/" + vendor

			if err := utils.DeleteDirectory(dirName); err != nil {
				panic(err)
			}
		}
	}

	for key, value := range envs {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		if err = utils.ReplaceInFiles(projectDir, strKey, strValue); err != nil {
			panic(err)
		}
	}

}

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

func ReplaceInFiles(directory, searchValue, replaceValue string) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			updatedContent := strings.ReplaceAll(string(content), searchValue, replaceValue)

			err = ioutil.WriteFile(path, []byte(updatedContent), info.Mode())
			if err != nil {
				return err
			}

			fmt.Printf("Updated file: %s\n", path)
		}

		return nil
	})

	return err
}

func ReadYAMLFile(filePath string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func RemoveGitDirectory(directoryPath string) error {
	gitDirectory := filepath.Join(directoryPath, ".git")

	err := os.RemoveAll(gitDirectory)
	if err != nil {
		return err
	}

	return nil
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func CloneRepository(projectDir, templateName string) error {
	_, err := git.PlainClone(projectDir, false, &git.CloneOptions{
		URL: templateName,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteDirectory(folderPath string) error {
	err := os.RemoveAll(folderPath)
	if err != nil {
		return fmt.Errorf("error deleting folder %s: %s", folderPath, err)
	}
	return nil
}

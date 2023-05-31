package utils

import (
	"reflect"
	"testing"
)

func TestReplaceInFiles(t *testing.T) {
	type args struct {
		directory    string
		searchValue  string
		replaceValue string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReplaceInFiles(tt.args.directory, tt.args.searchValue, tt.args.replaceValue); (err != nil) != tt.wantErr {
				t.Errorf("ReplaceInFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadYAMLFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadYAMLFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYAMLFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadYAMLFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveGitDirectory(t *testing.T) {
	type args struct {
		directoryPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveGitDirectory(tt.args.directoryPath); (err != nil) != tt.wantErr {
				t.Errorf("RemoveGitDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCloneRepository(t *testing.T) {
	type args struct {
		projectDir   string
		templateName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloneRepository(tt.args.projectDir, tt.args.templateName); (err != nil) != tt.wantErr {
				t.Errorf("CloneRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteDirectory(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteDirectory(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

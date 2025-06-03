package main

import (
	"boilercli/cmd"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := copyEmbeddedTemplate("template")
	if err != nil {
		fmt.Println("❌ Failed to generate project:", err)
	} else {
		fmt.Println("✅ Project created successfully.")
	}
	cmd.Execute()
}

func copyEmbeddedTemplate(outputDir string) error {
	return fs.WalkDir(TemplateFiles, "template", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		relPath := strings.TrimPrefix(path, "template")
		targetPath := filepath.Join(outputDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		data, err := TemplateFiles.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath, data, 0644)
	})
}

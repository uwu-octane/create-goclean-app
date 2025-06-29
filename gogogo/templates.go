package main

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed templates/*.tmpl templates/**/*
var templatesFS embed.FS

// templateFiles walks the embedded templates directory and returns a map where
// the key is the target file path (without the trailing .tmpl suffix) and the
// value is the file contents. This replaces the big literal map previously
// kept inside main.go, making templates easier to maintain.
func templateFiles() map[string]string {
	out := make(map[string]string)

	// WalkDir gives us every file under templates/ (recursively).
	_ = fs.WalkDir(templatesFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}
		// Strip the leading "templates/" segment.
		rel := strings.TrimPrefix(path, "templates/")
		// Remove the optional .tmpl suffix so that on-disk paths match the
		// generated file paths exactly.
		rel = strings.TrimSuffix(rel, ".tmpl")

		out[rel] = string(b)
		return nil
	})

	return out
}

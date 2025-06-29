// main.go
// create-goclean-app: minimal Go backend scaffold generator that keeps the full
// go-clean-template-style hierarchy and adds generic Docker deployment files.
// Usage:
//   go run ./main.go
// ...then follow the interactive prompts.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	survey "github.com/AlecAivazis/survey/v2"
)

type Options struct {
	ProjectName string
	ModuleName  string
	DB          string // "postgres", "mysql", "none"
	UseMQ       bool
}

func main() {
	var o Options

	// ----------------------------------------------
	// Interactive questions (Vite-style CLI wizard)
	// ----------------------------------------------
	_ = survey.AskOne(&survey.Input{Message: "Project name:", Default: "my-service"}, &o.ProjectName)
	_ = survey.AskOne(&survey.Input{Message: "Go module path:", Default: fmt.Sprintf("github.com/your-org/%s", o.ProjectName)}, &o.ModuleName)
	_ = survey.AskOne(&survey.Select{Message: "Choose database:", Options: []string{"postgres", "mysql", "none"}, Default: "postgres"}, &o.DB)
	_ = survey.AskOne(&survey.Confirm{Message: "Include RabbitMQ support?", Default: false}, &o.UseMQ)

	// sanitize inputs (remove accidental spaces/newlines)
	o.ProjectName = strings.TrimSpace(o.ProjectName)
	o.ModuleName = strings.TrimSpace(o.ModuleName)

	// ----------------------------------------------
	if err := scaffold(o); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("✅  Project %s generated. Next steps:\n", o.ProjectName)
	fmt.Printf("   cd %s && docker compose up --build\n", o.ProjectName)
}

// -----------------------------------------------------------------------------
// scaffold creates the directory tree + files according to the chosen options.
// -----------------------------------------------------------------------------
func scaffold(o Options) error {
	root := o.ProjectName

	// 1. create directory hierarchy mirroring go-clean-template.
	for _, d := range baseDirs() {
		dir := filepath.Join(root, d)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	// 2. create templated files.
	for path, body := range templateFiles() {
		// Conditional inclusion based on CLI answers
		if strings.Contains(path, "rabbitmq") && !o.UseMQ {
			continue
		}
		if strings.Contains(path, "postgres") && o.DB != "postgres" {
			continue
		}
		if strings.Contains(path, "mysql") && o.DB != "mysql" {
			continue
		}

		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			return err
		}
		f, err := os.Create(full)
		if err != nil {
			return err
		}
		if err := template.Must(template.New(path).Parse(body)).Execute(f, o); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}

// The `baseDirs` and `templateFiles` helpers have been moved to their own files
// to improve maintainability.

package engine

import "fmt"

// Executor will handle YAML template execution (Nuclei style)
type Executor struct {
	TemplatesDir string
}

func NewExecutor(templatesDir string) *Executor {
	return &Executor{TemplatesDir: templatesDir}
}

func (e *Executor) ExecuteTemplate(target string, templateName string) {
	// TODO: Full YAML template parsing & execution logic
	fmt.Printf("🔧 Executing template: %s on %s (coming soon...)\n", templateName, target)
}

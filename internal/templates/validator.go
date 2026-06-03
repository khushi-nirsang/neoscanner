package templates

import "fmt"

func ValidateTemplate(t Template) error {
	if t.ID == "" {
		return fmt.Errorf("template missing ID")
	}
	if t.Info.Name == "" {
		return fmt.Errorf("template missing name")
	}
	if len(t.Requests) == 0 {
		return fmt.Errorf("template has no requests")
	}
	fmt.Printf("✅ Template validated: %s\n", t.Info.Name)
	return nil
}

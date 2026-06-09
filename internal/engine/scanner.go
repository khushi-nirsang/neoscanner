package engine

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/khushi-nirsang/neoscanner/internal/templates"
	"github.com/khushi-nirsang/neoscanner/internal/utils"
)

// Scanner is the main scanning engine
type Scanner struct {
	Threads    int
	Results    *Results
	httpClient *utils.HTTPClient
	templates  []*templates.Template
}

// NewScanner creates a new scanner instance
func NewScanner(threads int, timeout int) *Scanner {
	return &Scanner{
		Threads:    threads,
		Results:    NewResults(),
		httpClient: utils.NewHTTPClient(timeout),
		templates:  make([]*templates.Template, 0),
	}
}

// LoadTemplates loads all YAML templates from directory (including subfolders)
func (s *Scanner) LoadTemplates(templateDir string) {
	os.MkdirAll(templateDir, 0755)

	count := 0

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		name := info.Name()
		if len(name) >= 5 && (strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")) {
			fmt.Printf("📄 Trying to load: %s\n", path)
			tmpl, err := templates.LoadTemplate(path)
			if err == nil {
				s.templates = append(s.templates, tmpl)
				count++
			} else {
				fmt.Printf("❌ Failed to load %s: %v\n", name, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("⚠️ Error walking templates: %v\n", err)
	}

	fmt.Printf("🔧 Loaded %d templates\n", count)
}

// StartScan starts scanning the target
func (s *Scanner) StartScan(target string) {
	fmt.Printf("🔍 Scanning target: %s\n", target)
	os.MkdirAll("reports", 0755)

	resp, err := s.httpClient.Get(target)
	if err != nil {
		fmt.Printf("❌ Failed to connect: %v\n", err)
		return
	}
	// Note: Body is already read and closed inside http.Get

	// Execute all loaded templates
	for _, tmpl := range s.templates {
		s.executeTemplate(tmpl, target, resp)
	}

	// Fallback basic checks if no templates matched
	if len(s.Results.Items) == 0 {
		s.runBasicChecks(target, resp.Response)
	}
}

func (s *Scanner) runBasicChecks(target string, resp *http.Response) {
	if server := resp.Header.Get("Server"); server != "" {
		s.Results.Add(ScanResult{
			Target:      target,
			Name:        "Server Header Exposure",
			Severity:    "low",
			Matched:     true,
			Description: fmt.Sprintf("Server header exposed: %s", server),
		})
		fmt.Printf("✅ [+] Server Header Exposure [low] → %s\n", server)
	}
}

func (s *Scanner) executeTemplate(tmpl *templates.Template, target string, resp *utils.Response) {
	for _, req := range tmpl.Requests {
		for _, matcher := range req.Matchers {
			matched := false

			if matcher.Type == "word" {
				if matcher.Part == "header" {
					for _, word := range matcher.Words {
						if server := resp.Header.Get("Server"); server != "" && strings.Contains(server, word) {
							matched = true
							break
						}
					}
				} else if matcher.Part == "body" {
					for _, word := range matcher.Words {
						if strings.Contains(resp.BodyContent, word) {
							matched = true
							break
						}
					}
				}
			}

			if matched {
				s.Results.Add(ScanResult{
					Target:      target,
					Name:        tmpl.Info.Name,
					Severity:    tmpl.Info.Severity,
					Matched:     true,
					Description: tmpl.Info.Description,
				})
				fmt.Printf("✅ [+] %s [%s]\n", tmpl.Info.Name, tmpl.Info.Severity)
			}
		}
	}
}

// SaveResults saves results to JSON
func (s *Scanner) SaveResults(outputFile string) {
	fmt.Printf("📊 Saving results to %s\n", outputFile)
	s.Results.Print()

	// Save JSON
	if err := s.Results.SaveJSON(outputFile); err != nil {
		fmt.Printf("❌ Failed to save JSON: %v\n", err)
	} else {
		fmt.Printf("✅ JSON report saved!\n")
	}

	// Save HTML Report
	htmlFile := "reports/results.html"
	if err := s.Results.SaveHTML(htmlFile); err != nil {
		fmt.Printf("❌ Failed to save HTML: %v\n", err)
	} else {
		fmt.Printf("✅ HTML report saved: %s\n", htmlFile)
	}
}

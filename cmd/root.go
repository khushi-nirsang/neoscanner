package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/khushi-nirsang/neoscanner/internal/config"
	"github.com/khushi-nirsang/neoscanner/internal/engine"
	"github.com/spf13/cobra"
)

var (
	target       string
	threads      int
	templateDir  string
)

var rootCmd = &cobra.Command{
	Use:   "neoscanner",
	Short: "NeoScanner - Next Generation Vulnerability Scanner",
	Long:  `A fast, extensible, template-based vulnerability scanner.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("NeoScanner starting...")

		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("Config error: %v", err)
			os.Exit(1)
		}

		if threads > 0 {
			cfg.Threads = threads
		}

		color.Cyan("[*] Scanning %s with %d threads", target, cfg.Threads)

		// Create scanner
		scanner := engine.NewScanner(cfg.Threads, cfg.Timeout)

		// Load templates
		if templateDir == "" {
			templateDir = "templates"
		}
		scanner.LoadTemplates(templateDir)

		// Start scan
		scanner.StartScan(target)

		// Save results
		outputFile := "reports/results.json"
		scanner.SaveResults(outputFile)

		color.Green("\n✅ Done. Findings: %d", len(scanner.Results.Items))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&target, "target", "u", "", "Target URL or IP (required)")
	rootCmd.Flags().IntVarP(&threads, "threads", "c", 50, "Number of concurrent threads")
	rootCmd.Flags().StringVarP(&templateDir, "templates", "t", "templates", "Templates directory")

	rootCmd.MarkFlagRequired("target")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

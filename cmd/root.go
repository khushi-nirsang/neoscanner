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
	target  string
	threads int
	output  string
)

var rootCmd = &cobra.Command{
	Use:   "neoscanner",
	Short: "NeoScanner - Next Generation Vulnerability Scanner",
	Long:  `A fast, extensible, template-based vulnerability scanner.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("NeoScanner starting...")

		// Load configuration
		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("Config error: %v", err)
			os.Exit(1)
		}

		// Override threads if user provided
		if threads > 0 {
			cfg.Threads = threads
		}

		color.Cyan("[*] Scanning %s with %d threads", target, cfg.Threads)

		// Start scanning
		scanner := engine.NewScanner(cfg.Threads)
		scanner.StartScan(target)

		// Save results
		outputFile := "reports/results.json"
		if output != "" {
			outputFile = output
		}
		scanner.SaveResults(outputFile)

		color.Green("Done. Findings: %d", len(scanner.Results.Items))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&target, "target", "u", "", "Target URL or IP (required)")
	rootCmd.Flags().IntVarP(&threads, "threads", "c", 50, "Number of concurrent threads")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Output JSON file path")

	rootCmd.MarkFlagRequired("target")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

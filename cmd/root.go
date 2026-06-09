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
)

var rootCmd = &cobra.Command{
	Use:   "neoscanner",
	Short: "NeoScanner - Next Generation Vulnerability Scanner",
	Long:  `A fast and extensible vulnerability scanner.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("NeoScanner starting...")

		cfg, err := config.LoadConfig()
		if err != nil {
			color.Red("Config error: %v", err)
			os.Exit(1)
		}

		// Override config with flags
		if target != "" {
			cfg.Target = target // This may need adjustment based on your config
		}
		if threads > 0 {
			cfg.Threads = threads
		}

		scanner := engine.NewScanner(cfg.Threads) // Adjust based on your engine
		results := scanner.StartScan(target)      // Passing target directly for now

		color.Green("Done. Findings: %d", len(results))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&target, "target", "u", "", "Target URL (required)")
	rootCmd.Flags().IntVarP(&threads, "threads", "c", 50, "Number of threads")

	rootCmd.MarkFlagRequired("target")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

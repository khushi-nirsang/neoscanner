package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	// Import the scanner engine
	"github.com/khushi-nirsang/neoscanner/internal/scanner"
)

var (
	cfgFile    string
	targetURL  string
	targetFile string
	templates  string
	threads    int
	outputFile string
	severity   string
)

var rootCmd = &cobra.Command{
	Use:   "neoscanner",
	Short: "NeoScanner - Next-Gen Vulnerability Scanner",
	Long:  `NeoScanner - Better than Nessus & Nuclei`,
	Run: func(cmd *cobra.Command, args []string) {
		if targetURL == "" && targetFile == "" {
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println("🚀 NeoScanner Starting...")
		startScan()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	
	rootCmd.Flags().StringVarP(&targetURL, "target", "u", "", "Target URL or IP (e.g. https://example.com)")
	rootCmd.Flags().StringVarP(&targetFile, "list", "l", "", "File containing list of targets")
	rootCmd.Flags().StringVarP(&templates, "templates", "t", "templates", "Templates directory")
	rootCmd.Flags().IntVarP(&threads, "threads", "c", 50, "Number of concurrent threads")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "results.json", "Output file")
	rootCmd.Flags().StringVarP(&severity, "severity", "s", "medium", "Severity filter (info,low,medium,high,critical)")

	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NeoScanner v0.1.0")
	},
}

// Updated startScan function - Connected with Engine
func startScan() {
	// Create scanner engine
	engine := scanner.NewEngine(threads, templates)

	if targetURL != "" {
		engine.StartScan(targetURL)
		engine.SaveResults(outputFile)
	} else if targetFile != "" {
		fmt.Println("Multiple targets support coming soon...")
	}
}

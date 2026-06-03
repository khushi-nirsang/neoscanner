package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	targetURL   string
	targetFile  string
	templates   string
	threads     int
	outputFile  string
	severity    string
)

var rootCmd = &cobra.Command{
	Use:   "neoscanner",
	Short: "NeoScanner - Next-Gen Vulnerability Scanner (Better than Nessus + Nuclei)",
	Long: `NeoScanner is a fast, intelligent, and extensible vulnerability scanner.
It combines the speed of Nuclei with the depth of Nessus using modern techniques.`,
	Run: func(cmd *cobra.Command, args []string) {
		if targetURL == "" && targetFile == "" {
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println("🚀 Starting NeoScanner...")
		// Yahan scanning engine call hoga
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
	rootCmd.Flags().StringVarP(&templates, "templates", "t", "templates/", "Directory containing templates")
	rootCmd.Flags().IntVarP(&threads, "threads", "c", 50, "Number of concurrent threads")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "results.json", "Output file (json/html)")
	rootCmd.Flags().StringVarP(&severity, "severity", "s", "medium", "Filter by severity (info, low, medium, high, critical)")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(templateCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err == nil {
			viper.AddConfigPath(home)
		}
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of NeoScanner",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NeoScanner v0.1.0 - Built for Cryptus Internship")
	},
}

var templateCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📂 Available Templates:")
		files, _ := filepath.Glob("templates/**/*.yaml")
		for _, f := range files {
			fmt.Println("   ✓", f)
		}
	},
}

func startScan() {
	// Yeh function internal/scanner/engine.go mein baad mein implement karenge
	fmt.Printf("Scanning target: %s with %d threads\n", targetURL, threads)
	fmt.Println("Results will be saved to:", outputFile)
	// TODO: Call actual scanning engine here
}

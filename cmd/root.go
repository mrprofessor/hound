package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hound",
	Short: "Hound is a web information gathering tool",
	Long:  `Hound aspires to be the only tool you need to get all the information for a website.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(WhoCmd)
	rootCmd.AddCommand(SslCmd)
	rootCmd.AddCommand(DnsCmd)
}

package cmd

import (
	"github.com/mrprofessor/hound/internal/pkg/ssl"
	"github.com/spf13/cobra"
)

var SslCmd = &cobra.Command{
	Use:     "ssl",
	Aliases: []string{"ssl"},
	Short:   "SSL data for a site",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ssl.RunSSL(args[0])
	},
}

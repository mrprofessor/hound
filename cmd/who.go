package cmd

import (
	"github.com/mrprofessor/hound/internal/pkg/whois"
	"github.com/spf13/cobra"
)

var WhoCmd = &cobra.Command{
	Use:     "who",
	Aliases: []string{"whois"},
	Short:   "Whois data about a domain",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		whois.RunWhoIs(args[0])
	},
}

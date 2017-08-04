package cmd

import "github.com/spf13/cobra"

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install packages",
	Run:   install,
}

var isDryRun bool

func init() {
	RootCmd.AddCommand(InstallCmd)
	InstallCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false,"See the commands as they would be executed")
}

func install(cmd *cobra.Command, args []string) {
}

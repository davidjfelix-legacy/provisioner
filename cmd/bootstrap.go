package cmd

import "github.com/spf13/cobra"

var bootstrapCmd = &cobra.Command{
	Use: "bootstrap",
	Short: "bootstrap the system for provisioning",
	Run: bootstrap,
}

func init() {
	RootCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false, "See the commands as they would be executed")
}

func bootstrap(cmd *cobra.Command, args []string) {
}

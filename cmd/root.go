package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "provisioner is an install script launcher/manager",
}

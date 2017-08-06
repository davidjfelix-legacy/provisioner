package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
	"os/exec"
	"github.com/labstack/gommon/log"
	"strings"
	"fmt"
)


var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install packages",
	Run:   install,
}

var isDryRun bool

func init() {
	RootCmd.AddCommand(InstallCmd)
	InstallCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false, "See the commands as they would be executed")
}

func install(cmd *cobra.Command, args []string) {
}

func getInstallerNames() []string {
	var names []string

	names = append(names, runtime.GOOS)
	names = append(names, runtime.GOOS+"-"+runtime.GOARCH)

	switch runtime.GOOS {
	case "darwin":
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err != nil {
			log.Fatal(err)
		}
		version, err := mapOsxVersions(string(out))
		if err != nil {
			log.Fatal(err)
		}
		names = append(
			names,
			runtime.GOOS+"-any-"+version,
			runtime.GOOS+"-"+runtime.GOARCH+"-"+version,
		)
	}
	return names
}

func mapOsxVersions(versionNum string) (string, error) {
	switch {
	case strings.HasPrefix(versionNum, "17.13"):
		return "high_sierra", nil
	case strings.HasPrefix(versionNum, "17.12"):
		return "sierra", nil
	case strings.HasPrefix(versionNum, "17.11"):
		return "el_capitan", nil
	case strings.HasPrefix(versionNum, "17.10"):
		return "mavericks", nil
	default:
		return "", fmt.Errorf("OSX version number not supported or not recognized: %s", versionNum)
	}
}

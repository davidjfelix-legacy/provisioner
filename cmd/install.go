package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
	"os/exec"
	"github.com/labstack/gommon/log"
	"strings"
	"fmt"
	"bufio"
	"os"
)


var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install packages",
	Run:   install,
}

var isDryRun bool

func init() {
	RootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&isDryRun, "dry", "d", false, "See the commands as they would be executed")
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
			// FIXME: better error message.
			log.Fatal(err)
		}
		version, err := mapOsxVersions(string(out))
		if err != nil {
			// FIXME: better error message.
			log.Fatal(err)
		}
		names = append(
			names,
			runtime.GOOS+"-any-"+version,
			runtime.GOOS+"-"+runtime.GOARCH+"-"+version,
		)
		break
	case "ubuntu":
		file, err := os.Open("/etc/os-release")
		if err != nil {
			// FIXME: better error message.
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "VERSION_ID=") {
				versionName := strings.TrimSuffix(strings.TrimPrefix(scanner.Text(),"VERSION_ID=\""), "\"")
				version, err := mapUbuntuVersions(versionName)
				if err != nil {
					// FIXME: better error mesage.
					log.Fatal(err)
				}
				names = append(
					names,
					runtime.GOOS+"-any-"+version,
					runtime.GOOS+"-"+runtime.GOARCH+"-"+version,
				)
			}
			break
		}
		file.Close()
		break
	case "debian":
		file, err := os.Open("/etc/os-release")
		if err != nil {
			// FIXME: better error message.
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "VERSION_ID=") {
				versionName := strings.TrimSuffix(strings.TrimPrefix(scanner.Text(),"VERSION_ID=\""), "\"")
				version, err := mapDebianVersions(versionName)
				if err != nil {
					// FIXME: better error mesage.
					log.Fatal(err)
				}
				names = append(
					names,
					runtime.GOOS+"-any-"+version,
					runtime.GOOS+"-"+runtime.GOARCH+"-"+version,
				)
			}
			break
		}
		file.Close()
		break
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

func mapUbuntuVersions(versionNum string) (string, error) {
	switch {
	case strings.HasPrefix(versionNum, "17.10"):
		return "artful", nil
	case strings.HasPrefix(versionNum, "17.04"):
		return "zenial", nil
	case strings.HasPrefix(versionNum, "16.10"):
		return "xenial", nil
	case strings.HasPrefix(versionNum, "16.04"):
		return "wily", nil
	case strings.HasPrefix(versionNum, "14.04"):
		return "trusty", nil
	default:
		return "", fmt.Errorf("Ubuntu version number not supported or not recognized: %s", versionNum)
	}
}

func mapDebianVersions(versionNum string) (string, error) {
	switch {
	case strings.HasPrefix(versionNum, "9"):
		return "stretch", nil
	case strings.HasPrefix(versionNum, "8"):
		return "jessie", nil
	case strings.HasPrefix(versionNum, "7"):
		return "wheezy", nil
	default:
		return "", fmt.Errorf("Debian version number not supported or not recognized: %s", versionNum)
	}
}
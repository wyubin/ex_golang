package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	tryCmd = &cobra.Command{
		Use:   "try",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		Run:   try,
	}
)

func try(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		repo.GetGitRepository(args[0], false)
		fmt.Printf("Repo added.\n")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.WriteConfigAs(home + "/jamctl.yaml")
	} else {
		fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		return
	}
}

func init() {
	tryCmd.PersistentFlags().StringVar(&relativePath, "relativePath", "/content", "The directory within the git repository which contains the markdown files.")
	tryCmd.PersistentFlags().StringVar(&targetPath, "targetPath", "/tmp/jamctl", "The absolute path where the git repository is going to cloned to.")
	viper.BindPFlag("relativePath", tryCmd.PersistentFlags().Lookup("relativePath"))
	viper.BindPFlag("targetPath", tryCmd.PersistentFlags().Lookup("targetPath"))
}

func includeAddFlags(cmd *cobra.Command) {

}

package cmd

import (
	"os"
	"os/exec"

	"example.com/command/utils/log"
	"github.com/spf13/cobra"
)

var (
	keyword string
	lsCmd   = &cobra.Command{
		Use:   "ls",
		Short: "ls files with grep",
		Long:  ``,
		Run:   lsFunc,
	}
)

func lsFunc(ccmd *cobra.Command, args []string) {
	if keyword != "" && len(args) > 0 {
		cmdLS := exec.Command("ls", args[0])
		cmdGrep := exec.Command("grep", keyword)
		cmdGrep.Stdin, _ = cmdLS.StdoutPipe()
		cmdGrep.Stdout = os.Stdout
		cmdGrep.Start()
		cmdLS.Run()
		cmdGrep.Wait()
		log.Logger.Info("ls finished")
	} else {
		log.Logger.Warn("no input keyword")
	}
	return
}

func init() {
	lsCmd.PersistentFlags().StringVarP(&keyword, "keyword", "k", "NA", "keyword for grep")
}

package cmd

import (
	"fmt"

	"example.com/command/pkg/add"
	"example.com/command/utils/log"
	"github.com/spf13/cobra"
)

var (
	title    string
	sliceInt []int
	addCmd   = &cobra.Command{
		Use:   "add",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		Run:   addFunc,
	}
)

func addFunc(ccmd *cobra.Command, args []string) {
	if len(sliceInt) > 0 {
		res := add.Sum(sliceInt)
		log.Logger.Info(fmt.Sprintf("Get [%s]: %d\n", title, res))
	} else {
		log.Logger.Warn("no input numbers")
	}
	return
}

func init() {
	addCmd.PersistentFlags().StringVar(&title, "title", "NA", "add title for integer sum")
	addCmd.PersistentFlags().IntSliceVar(&sliceInt, "numbers", []int{}, "input integers to get sum")
}

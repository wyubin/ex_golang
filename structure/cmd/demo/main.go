package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "demo",
	Short: "test tools",
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

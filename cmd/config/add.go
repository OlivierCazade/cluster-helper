package config

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func newAddConfigCmd(home string) *cobra.Command {
	var addConfigCmd = &cobra.Command{
		Use:   "add <name> <path>",
		Short: "add the config files to the list of availables config files",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source, err := os.Open(args[1])
			if err != nil {
				log.Fatal(err)
			}
			defer source.Close()

			dest, err := os.Create(home + "/" + args[0] + ".yaml")
			if err != nil {
				log.Fatal(err)
			}
			defer dest.Close()

			_, err = io.Copy(dest, source)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return addConfigCmd
}

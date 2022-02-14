package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func newListConfigCmd(home string) *cobra.Command {
	var addConfigCmd = &cobra.Command{
		Use:   "list",
		Short: "list the config files available",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(home)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				fmt.Println(strings.Split(file.Name(), ".")[0])
			}
		},
	}
	return addConfigCmd
}

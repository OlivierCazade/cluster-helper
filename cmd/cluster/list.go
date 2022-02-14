package cluster

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

func newListClusterCmd(home string) *cobra.Command {
	var addConfigCmd = &cobra.Command{
		Use:   "list",
		Short: "list the existing clusters",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(home + "/cluster")
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				fmt.Println(file.Name())
			}
		},
	}
	return addConfigCmd
}

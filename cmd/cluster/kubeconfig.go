package cluster

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func newKubeconfigClusterCmd(home string) *cobra.Command {
	var addConfigCmd = &cobra.Command{
		Use:   "kubeconfig",
		Short: "colon delimited list of available kubeconfig files",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(home + "/cluster")
			if err != nil {
				log.Fatal(err)
			}
			var res strings.Builder
			for _, file := range files {
				if res.Len() > 0 {
					res.WriteString(":")
				}
				res.WriteString(home)
				res.WriteString("/cluster/")
				res.WriteString(file.Name())
				res.WriteString("/auth/kubeconfig")
			}
			fmt.Println(res.String())
		},
	}
	return addConfigCmd
}

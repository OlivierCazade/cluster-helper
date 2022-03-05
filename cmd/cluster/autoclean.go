package cluster

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func confirm() (bool, error) {
	var answer string

	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false, err
	}

	answer = strings.ToLower(answer)
	if answer == "y" || answer == "yes" {
		return true, nil
	} else {
		return false, nil
	}
}

func newAutocleanCmd(home string) *cobra.Command {
	var addConfigCmd = &cobra.Command{
		Use:   "autoclean",
		Short: "automatically clean cluster configurations of not reachable clusters",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(home + "/cluster")
			if err != nil {
				log.Fatal(err)
			}
			res := []string{}
			var kubeconfig strings.Builder
			for _, file := range files {
				if kubeconfig.Len() > 0 {
					kubeconfig.WriteString(":")
				}
				kubeconfig.WriteString(home)
				kubeconfig.WriteString("/cluster/")
				kubeconfig.WriteString(file.Name())
				kubeconfig.WriteString("/auth/kubeconfig")
				config, err := clientcmd.BuildConfigFromFlags("", kubeconfig.String())
				if err != nil {
					res = append(res, file.Name())
				} else {
					clientset, err := kubernetes.NewForConfig(config)
					if err != nil {
						res = append(res, file.Name())
					} else {
						_, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
						if err != nil {
							res = append(res, file.Name())
						}
					}
				}
				kubeconfig.Reset()

			}
			fmt.Println("WARNING! This will remove:")
			for _, cluster := range res {
				fmt.Println("  - " + cluster)
			}
			fmt.Print("\nAre you sure you want to continue? [y/N]")
			proceed, err := confirm()
			if proceed {
				for _, cluster := range res {
					if kubeconfig.Len() > 0 {
						kubeconfig.WriteString(":")
					}
					kubeconfig.WriteString(home)
					kubeconfig.WriteString("/cluster/")
					kubeconfig.WriteString(cluster)
					if err = os.RemoveAll(kubeconfig.String()); err != nil {
						log.Fatal(err)
					}
					kubeconfig.Reset()
				}
			}
		},
	}
	return addConfigCmd
}

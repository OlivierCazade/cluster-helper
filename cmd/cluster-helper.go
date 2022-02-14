package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OlivierCazade/cluster-helper/cmd/cluster"
	"github.com/OlivierCazade/cluster-helper/cmd/config"
	"github.com/spf13/cobra"
)

var home string

func initHome() (string, error) {
	home = os.Getenv("CHHOME")
	if len(home) == 0 {
		home = os.Getenv("HOME") + "/.cluster-helper"
	}
	configpath := filepath.Join(home, "config")
	err := os.MkdirAll(configpath, os.ModePerm)
	if err != nil {
		return "", err
	}
	clusterpath := filepath.Join(home, "cluster")
	err = os.MkdirAll(clusterpath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return home, nil
}

func initRootCmd() (*cobra.Command, error) {
	home, err := initHome()
	if err != nil {
		return nil, err
	}
	var rootCmd = &cobra.Command{
		Use:   "cluster-helper",
		Short: "A wrapper to the openshift-install command",
		Long: `A wrapper to the openshift-install command.
                Manage cluster install file and cluster files lifecycle.`,
	}
	rootCmd.AddCommand(config.NewConfigCmd(home + "/config"))
	rootCmd.AddCommand(cluster.NewClusterCmd(home))
	return rootCmd, nil
}

func main() {
	rootCmd, err := initRootCmd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

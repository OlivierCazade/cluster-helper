package cluster

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func createClusterDir(home string, clusterName string) error {
	configpath := filepath.Join(home+"/cluster", clusterName)
	err := os.MkdirAll(configpath, os.ModePerm)
	return err
}

func initInstallConfig(home string, clusterName string, configName string) error {

	configPath := home + "/cluster/" + clusterName + "/install-config.yaml"

	data, err := ioutil.ReadFile(home + "/config/" + configName + ".yaml")
	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	metadata, ok := config["metadata"].(map[interface{}]interface{})
	if ok {
		metadata["name"] = clusterName
	}

	data, err = yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, data, 0640)
	if err != nil {
		return err
	}
	return nil
}

func startCluster(home string, clusterName string) error {
	clusterDir := home + "/cluster/" + clusterName
	cmd := exec.Command("openshift-install", "--dir="+clusterDir, "create", "cluster")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func newCreateClusterCmd(home string) *cobra.Command {
	var configName string
	var addConfigCmd = &cobra.Command{
		Use:   "create [<options] <name>",
		Short: "create a new cluster",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := createClusterDir(home, args[0])
			if err != nil {
				log.Fatal(err)
			}
			err = initInstallConfig(home, args[0], configName)
			if err != nil {
				log.Fatal(err)
			}
			err = startCluster(home, args[0])
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	addConfigCmd.Flags().StringVarP(&configName, "config", "c", "default", "Config name to use to create the cluster")
	return addConfigCmd
}

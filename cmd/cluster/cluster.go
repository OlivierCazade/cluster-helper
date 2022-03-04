package cluster

import "github.com/spf13/cobra"

func NewClusterCmd(home string) *cobra.Command {
	var clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "Manage openshift cluster using openshift-installer",
	}
	// clusterCmd.AddCommand(newAddClusterCmd(home))
	clusterCmd.AddCommand(newListClusterCmd(home))
	clusterCmd.AddCommand(newCreateClusterCmd(home))
	clusterCmd.AddCommand(newKubeconfigClusterCmd(home))
	return clusterCmd
}

package config

import "github.com/spf13/cobra"

func NewConfigCmd(home string) *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage install config files",
		Long:  `Manage install config files`,
	}
	configCmd.AddCommand(newAddConfigCmd(home))
	configCmd.AddCommand(newListConfigCmd(home))
	return configCmd
}

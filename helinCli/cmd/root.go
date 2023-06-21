package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string
var rootCmd = &cobra.Command{
	Use:     "helincli",
	Short:   "hlcli",
	Long:    "何林测试cli工具",
	Version: "betav1版本，不稳定版，在学习K8s",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("name", name)
		fmt.Println("name:", name)
	},
}

func Exec() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "the name to be greeted")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

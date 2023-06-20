package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	fun "hlcli.com/m/v2/internal"
)

var name string
var rootCmd = &cobra.Command{
	Use:     "helincli",
	Short:   "hlcli",
	Long:    "何林测试cli工具",
	Version: "betav1版本，不稳定版，在学习K8s",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("进入这里了")
	},
}

func Exec() {
	rootCmd.AddCommand(fun.FuncCmd)
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "the name to be greeted")
	fmt.Println("name:", name)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

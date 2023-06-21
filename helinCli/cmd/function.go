package cmd

import (
	"github.com/spf13/cobra"
	"hlcli.com/m/v2/internal/function"
)

var FuncCmd = &cobra.Command{
	Use:     "func",
	Aliases: []string{"fn", "hlfn"},
	Short:   "何林的测试",
	Long:    "这是何林来试验cli工具怎么用，只是为了测试，为了测试，测试",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	FuncCmd.AddCommand(function.SayCmd)
	FuncCmd.AddCommand(function.InitCmd)
	rootCmd.AddCommand(FuncCmd)
}

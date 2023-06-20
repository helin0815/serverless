package internal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "说话",
	Long:  "用来说话的工具",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("name", "jjjjj")
		name := viper.GetString("name")
		fmt.Println("hello", name)
	},
}

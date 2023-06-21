package function

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SayCmd = &cobra.Command{
	Use:   "say",
	Short: "说话",
	Long:  "用来说话的工具",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		fmt.Println("hello", name)
	},
}

package function

import "github.com/spf13/cobra"

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化函数配置信息",
	Long:  `初始化函数配置信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

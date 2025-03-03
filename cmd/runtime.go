package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runtimeCmd = &cobra.Command{
	Use:   "runtime [容器ID]",
	Short: "检查容器运行时安全状态",
	Long: `对正在运行的容器进行安全检查，包括：
  - 特权模式检测
  - 敏感挂载点检查
  - 系统调用限制检查
  - 资源限制检查
  - 网络配置检查
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始检查容器：", args[0])
		// TODO: 实现运行时检查逻辑
	},
}

func init() {
	rootCmd.AddCommand(runtimeCmd)

	runtimeCmd.Flags().BoolP("live", "l", false, "实时监控模式")
	runtimeCmd.Flags().BoolP("all", "a", false, "显示所有检查项结果")
	runtimeCmd.Flags().StringP("format", "f", "text", "输出格式 (text/json)")
	runtimeCmd.Flags().StringP("output", "o", "", "输出结果到指定文件")
}
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "securedock",
	Short: "SecureDock - 容器安全检测工具",
	Long: `SecureDock 是一个全面的容器安全检测工具，提供以下功能：
  - 容器镜像扫描
  - 容器运行时安全检查
  更多功能正在开发中...
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "配置文件路径")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "启用详细输出模式")
}
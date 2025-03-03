package cmd

import (
	"fmt"
	"os"

	"github.com/securedock/pkg/server"
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

	// 检查是否以服务模式运行
	serverMode, _ := rootCmd.PersistentFlags().GetBool("server")
	if serverMode {
		host, _ := rootCmd.PersistentFlags().GetString("host")
		port, _ := rootCmd.PersistentFlags().GetInt("port")

		// 创建并启动HTTP服务
		server := server.NewServer(host, port)
		server.SetupRoutes()
		if err := server.Start(); err != nil {
			fmt.Println("启动HTTP服务失败:", err)
			os.Exit(1)
		}
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "配置文件路径")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "启用详细输出模式")
	rootCmd.PersistentFlags().BoolP("server", "s", false, "以HTTP服务模式运行")
	rootCmd.PersistentFlags().StringP("host", "", "localhost", "HTTP服务监听地址")
	rootCmd.PersistentFlags().IntP("port", "p", 8080, "HTTP服务监听端口")
}

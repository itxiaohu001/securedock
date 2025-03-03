package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [镜像名称]",
	Short: "扫描容器镜像的安全漏洞",
	Long: `对指定的容器镜像进行安全扫描，包括：
  - 已知漏洞检测
  - 敏感信息检查
  - 配置风险分析
  - 基础镜像检查
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始扫描镜像：", args[0])
		// TODO: 实现镜像扫描逻辑
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolP("severity", "s", false, "显示漏洞严重程度")
	scanCmd.Flags().BoolP("details", "d", false, "显示详细扫描结果")
	scanCmd.Flags().StringP("output", "o", "", "输出结果到指定文件")
}
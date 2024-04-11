package util

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var desc = strings.Join([]string{
	"该命令行用于在程序启动时读取参数",
}, "\n")

var CC string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动参数",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("配置文件使用：%s \n", CC)
	},
}

func cmdExecute() error {
	startCmd.Flags().StringVarP(&CC, "config", "c", "/config/conf.toml", "配置文件路径")
	err := startCmd.Execute()
	return err
}

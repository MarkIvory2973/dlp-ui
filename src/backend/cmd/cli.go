package cmd

import (
	"dlp-ui/pkg/utils"
	"fmt"
)

var (
	tag    = "dev"
	commit = "none"
)

func printVersion() {
	fmt.Println("  " +
		utils.Style("DLP UI", utils.Strong, utils.Green) +
		" " +
		utils.Style(tag, utils.Green) +
		utils.Style("+git.", utils.Green) +
		utils.Style(commit[:min(len(commit), 7)], utils.Green) +
		"  " +
		utils.Style("by Mark Ivory", utils.Fade),
	)
}

func printURL() {
	fmt.Println("  " +
		utils.Style("➜", utils.Green) +
		"  " +
		utils.Style("本地:", utils.Strong) +
		"   " +
		utils.Style("http://localhost:", utils.Sky) +
		utils.Style("5000", utils.Sky, utils.Strong) +
		utils.Style("/webui", utils.Sky),
	)
}

func printFiles() {
	fmt.Println("  " +
		utils.Style("➜", utils.Green) +
		"  " +
		utils.Style("下载文件:", utils.Strong) +
		" " +
		utils.Style("./down", utils.Sky),
	)
	fmt.Println("  " +
		utils.Style("➜", utils.Green) +
		"  " +
		utils.Style("日志文件:", utils.Strong) +
		" " +
		utils.Style("./dlp-ui.log", utils.Sky),
	)
}

func printShortcuts() {
	fmt.Println("  " +
		utils.Style("➜", utils.Fade, utils.Green) +
		"  " +
		utils.Style("按", utils.Fade) +
		" " +
		utils.Style("CTRL + C", utils.Strong) +
		" " +
		utils.Style("退出程序", utils.Fade),
	)
}

func PrintBanner() {
	fmt.Println()

	printVersion()

	fmt.Println()

	printURL()
	printFiles()
	printShortcuts()

	fmt.Println()
}

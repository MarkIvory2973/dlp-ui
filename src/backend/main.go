package main

import (
	"dlp-ui/log"
	"dlp-ui/utils"
	"dlp-ui/web"
	"dlp-ui/web/view"
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	tag    = "unknown"
	commit = "none"
)

//go:embed ui/*
var ui embed.FS

func printStartupInfo(logger *logrus.Logger) {
	fmt.Println()
	fmt.Println("  " +
		utils.UseStyles("DLP UI", utils.Highlight, utils.Green) +
		" " +
		utils.UseStyles(tag, utils.Green) +
		utils.UseStyles("+git.", utils.Green) +
		utils.UseStyles(commit, utils.Green) +
		"  " +
		utils.UseStyles("by Mark Ivory", utils.Fade),
	)
	fmt.Println()
	fmt.Println("  " +
		utils.UseStyles("➜", utils.Green) +
		"  " +
		utils.UseStyles("本地:", utils.Highlight) +
		"   " +
		utils.UseStyles("http://localhost:", utils.SkyBlue) +
		utils.UseStyles("5000", utils.SkyBlue, utils.Highlight) +
		utils.UseStyles("/ui", utils.SkyBlue),
	)

	ips, err := utils.ListIPs()
	if err != nil {
		logger.Error(err)
	}

	for _, ip := range ips {
		fmt.Println("  " +
			utils.UseStyles("➜", utils.Green) +
			"  " +
			utils.UseStyles("局域网:", utils.Highlight) +
			" " +
			utils.UseStyles("http://", utils.SkyBlue) +
			utils.UseStyles(ip, utils.SkyBlue) +
			":" +
			utils.UseStyles("5000", utils.SkyBlue, utils.Highlight) +
			utils.UseStyles("/ui", utils.SkyBlue),
		)
	}

	fmt.Println("  " +
		utils.UseStyles("➜", utils.Green) +
		"  " +
		utils.UseStyles("下载文件:", utils.Highlight) +
		" " +
		utils.UseStyles("./down", utils.SkyBlue),
	)
	fmt.Println("  " +
		utils.UseStyles("➜", utils.Green) +
		"  " +
		utils.UseStyles("日志文件:", utils.Highlight) +
		" " +
		utils.UseStyles("./dlp-ui.log", utils.SkyBlue),
	)
	fmt.Println("  " +
		utils.UseStyles("➜", utils.Fade, utils.Green) +
		"  " +
		utils.UseStyles("按", utils.Fade) +
		" " +
		utils.UseStyles("CTRL + C", utils.Highlight) +
		" " +
		utils.UseStyles("退出程序", utils.Fade),
	)
	fmt.Println()
}

func main() {
	// create a new logger located at './dlp-ui.log'
	logger, err := log.New("dlp-ui")
	if err != nil {
		panic(err)
	}

	// print startup infomations
	printStartupInfo(logger)

	// load ui from the binary
	ui, err := fs.Sub(ui, "ui")
	if err != nil {
		logger.Fatal(err)
	}

	// create a new router
	router := web.New(logger)

	// route '/api/parse'
	view.Parse(router)
	// route '/api/download'
	view.Download(router)
	// route '/ui'
	router.StaticFS("/ui", http.FS(ui))

	// listening and serving HTTP on :5000
	router.Run(":5000")
}

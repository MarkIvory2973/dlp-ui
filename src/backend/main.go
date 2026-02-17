package main

import (
	"dlp-ui/log"
	"dlp-ui/utils"
	"dlp-ui/web"
	"dlp-ui/web/view"
	"fmt"

	"github.com/sirupsen/logrus"
)

func printInfos(logger *logrus.Logger) {
	fmt.Println()
	fmt.Println("  " + utils.Highlight + utils.Green + "DLP UI" + utils.Reset + " " + utils.Green + "v2.0-alpha" + utils.Reset + "  " + utils.Fade + "by 093" + utils.Reset)
	fmt.Println()
	fmt.Println("  " + utils.Green + "➜" + utils.Reset + "  " + utils.Highlight + "本地:" + utils.Reset + "   " + utils.SkyBlue + "http://localhost:" + utils.Highlight + "5000" + utils.Reset + utils.SkyBlue + "/ui" + utils.Reset)

	ips, err := utils.ListIPs()
	if err != nil {
		logger.Error(err)
	}

	for _, ip := range ips {
		fmt.Println("  " + utils.Green + "➜" + utils.Reset + "  " + utils.Highlight + "局域网:" + utils.Reset + " " + utils.SkyBlue + "http://" + ip + ":" + utils.Highlight + "5000" + utils.Reset + utils.SkyBlue + "/ui" + utils.Reset)
	}

	fmt.Println("  " + utils.Green + "➜" + utils.Reset + "  " + utils.Highlight + "下载文件:" + utils.Reset + " " + utils.SkyBlue + "./down" + utils.Reset)
	fmt.Println("  " + utils.Green + "➜" + utils.Reset + "  " + utils.Highlight + "日志文件:" + utils.Reset + " " + utils.SkyBlue + "./dlp-ui.log" + utils.Reset)
	fmt.Println("  " + utils.Fade + utils.Green + "➜" + utils.Reset + "  " + utils.Fade + "按" + utils.Reset + " " + utils.Highlight + "CTRL + C" + utils.Reset + " " + utils.Fade + "退出程序" + utils.Reset)
	fmt.Println()
}

func main() {
	// create a new logger located at './dlp-ui.log'
	logger, err := log.New("dlp-ui")
	if err != nil {
		panic(err)
	}

	// create a new router
	router := web.New(logger)

	// route '/api/parse'
	view.Parse(router)
	// route '/api/download'
	view.Download(router)
	// route '/'
	view.UI(router)

	// print infomations
	printInfos(logger)

	// listening and serving HTTP on :5000
	router.Run(":5000")
}

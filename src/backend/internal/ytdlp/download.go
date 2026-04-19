package ytdlp

import (
	"dlp-ui/pkg/utils"
	"slices"

	"github.com/sirupsen/logrus"
)

type Download struct {
	URL  string `json:"url"`
	Task struct {
		Title   string `json:"title"`
		Current int    `json:"current"`
		Total   int    `json:"total"`
		Speed   int    `json:"speed"`
		Done    bool   `json:"done"`
	} `json:"task"`
	Errors []string `json:"errors"`
}

func NewDownloader(browser string, url string, format string, downloads []Download) (func(*logrus.Entry), error) {
	var args []string
	if browser != "" {
		args = []string{
			"--cookies-from-browser", browser,
			"-f", format,
			"--downloader-args", "aria2c:--summary-interval=1 --human-readable=false",
			"-O", "before_dl:TITLE: %(title)s",
			"-o", "down/%(title)s.%(ext)s",
			url,
		}
	} else {
		args = []string{
			"-f", format,
			"--downloader-args", "aria2c:--summary-interval=1 --human-readable=false",
			"-O", "before_dl:TITLE: %(title)s",
			"-o", "down/%(title)s.%(ext)s",
			url,
		}
	}

	command, stdout, stderr, err := new(args...)
	if err != nil {
		return nil, err
	}

	err = command.Start()
	if err != nil {
		return nil, err
	}

	index := slices.IndexFunc(downloads, func(download Download) bool {
		return url == download.URL
	})
	download := downloads[index]

	return func(logger *logrus.Entry) {
		go utils.ScanLineFunc(stdout, func(content string) {
			title := utils.ReadTitle(content)
			if title != "" {
				download.Task.Title = title
			}

			current, total, speed := utils.ReadAria2(content)
			if current != -1 && total != -1 && speed != -1 {
				download.Task.Current = current
				download.Task.Total = total
				download.Task.Speed = speed
			}

			downloads[index] = download
		})

		go utils.ScanLineFunc(stderr, func(error string) {
			download.Errors = append(download.Errors, error)
			downloads[index] = download
		})

		err = command.Wait()
		if err != nil {
			download.Errors = append(download.Errors, err.Error())
		}

		download.Task.Done = true
		downloads[index] = download

		for _, error := range download.Errors {
			logger.Errorf("error while downloading: %s", error)
		}
	}, nil
}

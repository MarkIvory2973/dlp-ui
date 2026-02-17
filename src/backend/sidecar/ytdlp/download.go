package ytdlp

import (
	"dlp-ui/utils"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type Progress struct {
	Current int  `json:"current"`
	Total   int  `json:"total"`
	Speed   int  `json:"speed"`
	Done    bool `json:"done"`
}

type Download struct {
	Title    string   `json:"title"`
	Progress Progress `json:"progress"`
	Errors   []string `json:"errors"`
}

func Downloader(url string, format string, downloads *sync.Map) (func(logger *logrus.Entry), error) {
	// create a new command: ... -f "${format}" --downloader-args "aria2c:--summary-interval=1 --human-readable=false" -O "before_dl:TITLE: %(title)s" -o "down/%(title)s.%(ext)s" "${url}"
	command, stdout, stderr, err := new(
		"-f", format,
		"--downloader-args", "aria2c:--summary-interval=1 --human-readable=false",
		"-O", "before_dl:TITLE: %(title)s",
		"-o", "down/%(title)s.%(ext)s",
		url,
	)
	if err != nil {
		return nil, err
	}

	// start to execute the command
	err = command.Start()
	if err != nil {
		return nil, err
	}

	return func(logger *logrus.Entry) {
		var download Download

		// read title and progress from stdout
		for {
			line, err := stdout.ReadString('\n')
			line = strings.Trim(line, "\n")
			line = strings.TrimSpace(line)

			// read title
			title := utils.ReadTitle(line)
			if title != "" {
				download.Title = title
			}

			// read progress
			current, total, speed := utils.ReadProgress(line)
			if current != -1 && total != -1 && speed != -1 {
				download.Progress.Current = current
				download.Progress.Total = total
				download.Progress.Speed = speed
			}

			// update downloads in real time
			downloads.Store(url, download)

			if err != nil {
				if err != io.EOF {
					err := fmt.Sprintf("%v", err)
					download.Errors = append(download.Errors, err)
				}

				break
			}
		}

		// read errors from stderr
		errors, err := utils.ReadLines(stderr)
		if err != nil {
			err := fmt.Sprintf("%v", err)
			download.Errors = append(download.Errors, err)
		}

		download.Errors = append(download.Errors, errors...)

		// wait for the command exits
		err = command.Wait()
		if err != nil {
			err := fmt.Sprintf("%v", err)
			download.Errors = append(download.Errors, err)
		}

		// mark the download as done
		download.Progress.Done = true

		// update downloads
		downloads.Store(url, download)

		for _, error := range download.Errors {
			logger.Errorf("error while downloading files: %s", error)
		}
	}, nil
}

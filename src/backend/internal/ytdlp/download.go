package ytdlp

import (
	"dlp-ui/pkg/outputs"
	"slices"
)

type Download struct {
	URL string `json:"url"`
	Job struct {
		Title   string `json:"title"`
		Current int    `json:"current"`
		Total   int    `json:"total"`
		Speed   int    `json:"speed"`
		Done    bool   `json:"done"`
	} `json:"job"`
	Errors []string `json:"errors"`
}

type Downloads []Download

func (downloads Downloads) Contains(url string) bool {
	return slices.ContainsFunc(downloads, func(download Download) bool {
		return url == download.URL
	})
}

func (downloads *Downloads) Append(url string) {
	download := Download{
		URL: url,
	}
	*downloads = append(*downloads, download)
}

func (downloads Downloads) Index(url string) int {
	return slices.IndexFunc(downloads, func(download Download) bool {
		return url == download.URL
	})
}

func (downloads *Downloads) Delete(url string) {
	*downloads = slices.DeleteFunc(*downloads, func(download Download) bool {
		return url == download.URL
	})
}

func NewDownloader(url string, format string) (func(Downloads), error) {
	extraArgs := []string{
		"-f", format,
		"--downloader-args", "aria2c:--summary-interval=1 --human-readable=false",
		"-O", "before_dl:TITLE: %(title)s",
		"-o", "down/%(title)s.%(ext)s",
		url,
	}
	process, err := new(extraArgs...)
	if err != nil {
		return nil, err
	}

	err = process.Start()
	if err != nil {
		return nil, err
	}

	return func(downloads Downloads) {
		index := downloads.Index(url)

		go outputs.ScanAria2Func(process.Stdout, func(title string, current int, total int, speed int) {
			if title != "" {
				downloads[index].Job.Title = title
			}

			if current != -1 {
				downloads[index].Job.Current = current
			}

			if total != -1 {
				downloads[index].Job.Total = total
			}

			if speed != -1 {
				downloads[index].Job.Speed = speed
			}
		})

		go outputs.ScanTextFunc(process.Stderr, func(error string) {
			downloads[index].Errors = append(downloads[index].Errors, error)
		})

		err := process.Wait()
		if err != nil {
			downloads[index].Errors = append(downloads[index].Errors, err.Error())
		}

		downloads[index].Job.Done = true
	}, nil
}

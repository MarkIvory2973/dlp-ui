package ytdlp

import (
	"dlp-ui/pkg/utils"
	"slices"

	"github.com/sirupsen/logrus"
)

type Parsed struct {
	URL  string `json:"url"`
	Task struct {
		Entries []map[string]any `json:"entries"`
		Done    bool             `json:"done"`
	} `json:"task"`
	Errors []string `json:"errors"`
}

func NewParser(browser string, url string, parseds []Parsed) (func(*logrus.Entry), error) {
	var args []string
	if browser != "" {
		args = []string{
			"--cookies-from-browser", browser,
			"-j", url,
		}
	} else {
		args = []string{
			"-j", url,
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

	index := slices.IndexFunc(parseds, func(parsed Parsed) bool {
		return url == parsed.URL
	})
	parsed := parseds[index]

	return func(logger *logrus.Entry) {
		go utils.ScanJsonFunc(stdout, func(entry map[string]any, err error) {
			if err != nil {
				parsed.Errors = append(parsed.Errors, err.Error())
				parseds[index] = parsed
			}

			parsed.Task.Entries = append(parsed.Task.Entries, entry)
			parseds[index] = parsed
		})

		go utils.ScanLineFunc(stderr, func(error string) {
			parsed.Errors = append(parsed.Errors, error)
			parseds[index] = parsed
		})

		err := command.Wait()
		if err != nil {
			parsed.Errors = append(parsed.Errors, err.Error())
		}

		parsed.Task.Done = true
		parseds[index] = parsed

		for _, entry := range parsed.Task.Entries {
			logger.Debugf("parsed entry '%s'", entry["title"])
		}

		for _, error := range parsed.Errors {
			logger.Errorf("error while parsing: %s", error)
		}
	}, nil
}

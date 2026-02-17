package ytdlp

import (
	"dlp-ui/utils"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Parsed struct {
	Entries []map[string]any `json:"entries"`
	Errors  []string         `json:"errors"`
}

func Parser(url string, parseds *sync.Map) (func(logger *logrus.Entry), error) {
	// create a new command: ... -j "${url}"
	command, stdout, stderr, err := new(
		"-j", url,
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
		var parsed Parsed

		// read entries from stdout
		entries, err := utils.ReadJsons(stdout)
		if err != nil {
			err := fmt.Sprintf("%v", err)
			parsed.Errors = append(parsed.Errors, err)
		}

		// if there are no entries
		if len(entries) == 0 {
			entries = append(entries, map[string]any{
				"title":   "",
				"formats": []string{},
			})
		}

		parsed.Entries = entries

		// read errors from stderr
		errors, err := utils.ReadLines(stderr)
		if err != nil {
			err := fmt.Sprintf("%v", err)
			parsed.Errors = append(parsed.Errors, err)
		}

		parsed.Errors = append(parsed.Errors, errors...)

		// wait for the command exits
		err = command.Wait()
		if err != nil {
			err := fmt.Sprintf("%v", err)
			parsed.Errors = append(parsed.Errors, err)
		}

		// update parseds
		parseds.Store(url, parsed)

		for _, entry := range parsed.Entries {
			logger.Debugf("parsed entry '%s'", entry["title"])
		}

		for _, error := range parsed.Errors {
			logger.Errorf("error while parsing: %s", error)
		}
	}, nil
}

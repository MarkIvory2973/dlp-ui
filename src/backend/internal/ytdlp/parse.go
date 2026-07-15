package ytdlp

import (
	"dlp-ui/pkg/outputs"
	"slices"
)

type Parsed struct {
	URL string `json:"url"`
	Job struct {
		Entries []map[string]any `json:"entries"`
		Done    bool             `json:"done"`
	} `json:"job"`
	Errors []string `json:"errors"`
}

type Parseds []Parsed

func (parseds Parseds) Contains(url string) bool {
	return slices.ContainsFunc(parseds, func(parsed Parsed) bool {
		return url == parsed.URL
	})
}

func (parseds *Parseds) Append(url string) {
	parsed := Parsed{
		URL: url,
	}
	*parseds = append(*parseds, parsed)
}

func (parseds Parseds) Index(url string) int {
	return slices.IndexFunc(parseds, func(parsed Parsed) bool {
		return url == parsed.URL
	})
}

func (parseds *Parseds) Delete(url string) {
	*parseds = slices.DeleteFunc(*parseds, func(parsed Parsed) bool {
		return url == parsed.URL
	})
}

func NewParser(url string) (func(Parseds), error) {
	extraArgs := []string{
		"-j", url,
	}
	process, err := new(extraArgs...)
	if err != nil {
		return nil, err
	}

	err = process.Start()
	if err != nil {
		return nil, err
	}

	return func(parseds Parseds) {
		index := parseds.Index(url)

		go outputs.ScanObjectFunc(process.Stdout, func(entry map[string]any) {
			parseds[index].Job.Entries = append(parseds[index].Job.Entries, entry)
		})

		go outputs.ScanTextFunc(process.Stderr, func(error string) {
			parseds[index].Errors = append(parseds[index].Errors, error)
		})

		err := process.Wait()
		if err != nil {
			parseds[index].Errors = append(parseds[index].Errors, err.Error())
		}

		parseds[index].Job.Done = true
	}, nil
}

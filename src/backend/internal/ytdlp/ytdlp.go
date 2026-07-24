package ytdlp

import (
	"dlp-ui/cmd"
	"io"
	"os/exec"
)

var baseArgs = []string{
	"-q", "--no-warnings",
	"--downloader", "aria2c",
}

type Process struct {
	Command *exec.Cmd
	Stdout  io.ReadCloser
	Stderr  io.ReadCloser
}

func (process Process) Start() error {
	return process.Command.Start()
}

func (process Process) Wait() error {
	return process.Command.Wait()
}

func init() {
	browser := cmd.GetBrowser()
	if browser != "" {
		baseArgs = append(baseArgs, "--cookies-from-browser", browser)
	}
}

func new(extraArgs ...string) (Process, error) {
	args := append(baseArgs, extraArgs...)
	command := exec.Command("yt-dlp", args...)

	stdout, err := command.StdoutPipe()
	if err != nil {
		return Process{}, err
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return Process{}, err
	}

	process := Process{
		Command: command,
		Stdout:  stdout,
		Stderr:  stderr,
	}

	return process, nil
}

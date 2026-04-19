package ytdlp

import (
	"io"
	"os/exec"
)

var baseArgs = []string{
	"-q", "--no-warnings",
	"--downloader", "bin/aria2c",
	"--ffmpeg-location", "bin",
}

func new(extraArgs ...string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	args := append(baseArgs, extraArgs...)
	command := exec.Command("bin/yt-dlp", args...)

	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	return command, stdout, stderr, nil
}

package ytdlp

import (
	"bufio"
	"os"
	"os/exec"
)

var commonArgs = []string{
	"-q", "--no-warnings",
	"--downloader", "bin/aria2c",
	"--ffmpeg-location", "bin",
}

/* TODO: 2.0-alpha test */

func init() {
	browser := os.Getenv("ALPHA_BROWSER")
	if browser != "" {
		commonArgs = append(commonArgs, "--cookies-from-browser", browser)
	}
}

/* TODO: 2.0-alpha test */

func new(extraArgs ...string) (*exec.Cmd, *bufio.Reader, *bufio.Reader, error) {
	// create a new command: bin/yt-dlp -q --no-warnings --downloader "bin/aria2c" --ffmpeg-location "bin" ...
	args := append(commonArgs, extraArgs...)
	command := exec.Command("bin/yt-dlp", args...)

	// get stdout pipe
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	// get stderr pipe
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	// stdout reader
	stdoutReader := bufio.NewReader(stdoutPipe)
	// stderr reader
	stderrReader := bufio.NewReader(stderrPipe)

	return command, stdoutReader, stderrReader, nil
}

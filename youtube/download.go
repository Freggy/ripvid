package youtube

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

const DownloadLocation = "/tmp/"

// DownloadVideo downloads a youtube video with the given id and format using youtube-dl.
// Supported formats are documented in the youtube-dl documentation. The video will be saved at /tmp.
// The path to the file will be returned.
func DownloadVideo(filename, format, id string) (string, error) {
	args := []string{"--format", format}
	if err := runCommand(filename, format, id, args...); err != nil {
		return "", err
	}
	return DownloadLocation + filename, nil
}

// DownloadAudio downloads a youtube video with the given id and extracts the audio using youtube-dl.
// Supported audio formats are documented in the youtube-dl documentation. The audio will be saved at /tmp.
// The path to the file will be returned.
func DownloadAudio(filename, format, id string) (string, error) {
	args := []string{"-x", "--audio-format", format}
	if err := runCommand(filename, format, id, args...); err != nil {
		return "", err
	}
	return DownloadLocation + filename, nil
}

func runCommand(filename, format, id string, extraArgs ...string) error {
	args := buildArgs(filename, id, extraArgs...)
	cmd := exec.Command("youtube-dl", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	log.Println(cmd.String()) // TODO: only log in DEBUG mode
	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

func buildArgs(filename, id string, extraArgs ...string) []string {
	args := make([]string, 0)
	args = append(args, "-o", DownloadLocation+filename)
	if extraArgs != nil {
		args = append(args, extraArgs...)
	}
	args = append(args, fmt.Sprintf("https://youtube.com/watch?v=%s", id))
	return args
}

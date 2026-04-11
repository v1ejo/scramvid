package ffmpeg

import (
	"os"
	"os/exec"
)

func ExtractFrames(path string) error {
	if err := os.RemoveAll("video/frames"); err != nil {
		return err
	}

	if err := os.MkdirAll("video/frames", 0755); err != nil {
		return err
	}

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", path,
		"-vf", "fps=24",
		"video/frames/frame_%04d.png",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ExtractAudio(path string) error {
	if err := os.MkdirAll("video", 0755); err != nil {
		return err
	}
	cmd := exec.Command("ffmpeg", "-y", "-i", path, "-vn", "-acodec", "copy", "video/audio.m4a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func JoinFramesAndAudio(path string) error {
	audioPath := "video/audio.m4a"
	cmd := exec.Command("ffmpeg", "-y", "-framerate", "24", "-i", "video/scrambled/frame_%04d.png", "-i", audioPath,
		"-c:v", "libx264", "-c:a", "copy", "-shortest", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

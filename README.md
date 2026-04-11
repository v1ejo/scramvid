# scramvid

A CLI tool to scramble, unscramble video files using frame-level transformations and a secret key.

## Requirements
- Linux
- Go (1.18+ recommended)
- ffmpeg installed and available in your PATH

## Build

From the project root, build each command:

```bash
go build -o bin/scramble   ./cmd/scramvid/scramble
go build -o bin/unscramble ./cmd/scramvid/unscramble
go build -o bin/clean      ./cmd/scramvid/clean
```

## Usage

### Scramble
Scrambles a video file using a secret key.

```bash
./scramble -in=video.mp4 -key=SECRET -out=scrambled.mp4
```
- `-in`   : Path to the input video file
- `-key`  : Secret key for scrambling
- `-out`  : Path for the scrambled output video

**What it does:**
1. Extracts frames and audio from the input video.
2. Scrambles each frame with the provided key.
3. Saves scrambled frames to `video/scrambled/`.
4. Reassembles the video from scrambled frames and original audio.
5. Cleans up the `video/` temp directory.

### Unscramble
Restores a scrambled video file using the same secret key.

```bash
./unscramble -in=scrambled.mp4 -key=SECRET -out=restored.mp4
```
- `-in`   : Path to the scrambled video file
- `-key`  : Secret key for unscrambling (must match the scramble key)
- `-out`  : Path for the restored output video

**What it does:**
1. Extracts frames and audio from the scrambled video.
2. Unscrambles each frame with the provided key.
3. Saves restored frames to `video/scrambled/`.
4. Reassembles the video from restored frames and original audio.
5. Cleans up the `video/` temp directory.

### Clean
Removes all temporary files and folders used during processing.

```bash
./clean
```

**What it does:**
- Deletes the `video/` temp directory and all its contents.

## Notes
- All commands must be run from the project root or a directory where relative paths resolve correctly.
- On error, temporary files may remain; use `./clean` to remove them.
- For help on any command, use the `-h` flag (e.g., `./scramble -h`).

# DLP UI

DLP UI is a frontend of yt-dlp.

## Installation

Clone this repository:

```bash
git clone https://github.com/MarkIvory2973/dlp-ui.git
cd dlp-ui
```

Install dependencies:

```bash
yarn --cwd src/frontend install
```

⚠ The following binaries are required and must be placed in the `./bin` directory 
(the same directory as the `dlp-ui` executable):

- [x] `yt-dlp`
- [x] `aria2c`
- [x] `ffmpeg`
- [x] `ffprobe`

Build frontend:

```bash
yarn --cwd src/frontend build
mv src/frontend/dist src/backend/ui
```

Build backend:

```bash
cd src/backend
go build .
```

## Usage

```bash
./dlp-ui
```

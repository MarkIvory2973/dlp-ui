# DLP UI

DLP UI is a frontend of yt-dlp.

## Installation

Clone this repository:

```bash
git clone https://github.com/MarkIvory2973/dlp-ui.git
```

Install dependencies:

```bash
cd src/frontend
yarn install
cd ../..
```

⚠ The following binaries are required and must be placed in the `./bin` directory 
(the same directory as the `dlp-ui` executable):

- [x] `yt-dlp`
- [x] `aria2c`
- [x] `ffmpeg`
- [x] `ffprobe`

Build frontend:

```bash
cd src/frontend
yarn build
cd ../..
```

Build backend:

```bash
cd src/backend
go build .
cd ../..
```

## Usage

```bash
./dlp-ui
```
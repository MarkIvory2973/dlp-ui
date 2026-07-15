# DLP UI

DLP UI is a frontend for yt-dlp.

## Installation

> **Note:**
>
> The following binaries are required and must be available in your `$PATH`:
>
> - [x] `yt-dlp`
> - [x] `aria2c`
> - [x] `ffmpeg`
> - [x] `ffprobe`

### GitHub Releases

Download latest release from [GitHub Releases](https://github.com/MarkIvory2973/dlp-ui/releases/latest).

### Build from source

#### Requirements

- Go 1.26+
- UPX
- GNU Make
- Git

Clone the repository:

```bash
git clone https://github.com/MarkIvory2973/dlp-ui.git
cd dlp-ui
```

Install dependencies:

```bash
make install
```

Build binaries:

```bash
make build
```

Clean files:

```bash
make clean
```

## Usage

Run the following command:

```bash
./dlp-ui
```

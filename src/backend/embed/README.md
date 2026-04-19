# DLP UI

DLP UI is a frontend for yt-dlp.

## About this folder

This folder contains embeded files.

The following steps should be done:

1. Build frontend in the `src/frontend` directory:

    ```bash
    npm run build
    ```

2. Move the output to this directory:

    ```bash
    mv src/frontend/dist src/backend/embed/webui
    ```

## Structure

```
embed
 ├─ webui
 │   └─ ...
 └─ README.md
```

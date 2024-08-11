# Overview
- The `crawl/deepl` project is a Go-based application(`go 1.22.0`) designed to fetch text from the clipboard, translate it using Deepl's web interface, and optionally send the translations via a Telegram bot.
- The application leverages Docker to ensure a consistent and isolated environment for development and execution.
- It serves as a project for me to lean the Golang programming lanague. Feel free to give feedback or open pullreuests

## Current Setup

### How It Works
- Initialization: The application first checks for an active internet connection.
- Chrome Context: A Chrome context is created to allow headless browsing.
- Clipboard Access: Text is fetched from the clipboard.
- Translation: The text is translated by navigating to the Deepl website and extracting the translated text using Chromedp.
- Output: The translated text is output to the console, and if a Telegram bot is set up, it sends the translation via Telegram

### Docker Setup
- In this branch with this setup, there is only a docker-compose.dev.yml which starts a container named `goCrawlDevtainer`. The implementation of this programm in the future will also have a production-ready docker-compose.yml

#### Running the Container
##### Start the Container:

```bash
./run.sh
```
##### Connect to the Container:
```bash
docker exec -it goCrawlDevtainer bash
```
##### Start the Translation:
Inside the container, run the following command to start the translation:
```bash
/usr/local/bin/app/deepl
```

## Security Concerns
- The current setup poses potential security risks due to the following reasons:

### X11 Server Exposure: 
- By mounting the host's X11 server socket into the container, the browser inside the container can access the host's display server. This could allow malicious software to interact with the host's display, leading to potential security breaches.
### Elevated Privileges:
- The container is granted the SYS_ADMIN capability, which provides elevated privileges that could be exploited if the container is compromised.
### Network Exposure:
- Running a web browser within the container that has network access increases the risk of exposing sensitive data or becoming a target for network-based attacks.

### Future Plans
- This branch (`feature/mounted_X11_server_solution`) will not be continued due to the aforementioned security concerns. Future development will focus on implementing a more secure solution.

### Requirements
- Docker: Ensure Docker is installed on your system. You can find installation instructions at [Docker Documentation.](https://docs.docker.com/get-docker/)
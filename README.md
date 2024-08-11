# Overview
- The `crawl/deepl` project is a Go-based application (`go 1.22.0`) designed to translate text using Deepl's web interface and optionally send the translations via a Telegram bot.
- The application leverages Docker to ensure a consistent and isolated environment for development and execution.
- It serves as a project for me to learn the Golang programming language. Feel free to give feedback or open pull requests.

## Current Setup

### How It Works
- **Initialization**: The application first checks for an active internet connection.
- **Chrome Context**: A Chrome context is created to allow headless browsing.
- **Translation**: The text is translated by navigating to the Deepl website and extracting the translated text using Chromedp.
- **Output**: The translated text is output to the console, and if a Telegram bot is set up, it sends the translation via Telegram.

### Setup Telegram bot
1. Get your Bot-Token and Chat-ID: https://core.telegram.org/bots
2. Create an .env file

3. Set the following variables in the .env file:
- `BOT_TOKEN`: Your Telegram bot token.
- `CHAT_ID`: Your Telegram chat ID.

### Docker Setup

#### Development Environment
- The development environment is defined in `docker-compose.dev.yml` and starts a container named `goCrawlDevtainer`.

#### Production Environment
- The production environment is defined in `docker-compose.yml` and starts a container named `goCrawlApp`.

### Development
#### Build the Development Container:
```bash
docker compose -f docker-compose.dev.yml build
```

##### Start the Development Container:
```bash
docker compose -f docker-compose.dev.yml up
```
##### Connect to the Development Container:
```bash
docker exec -it goCrawlDevtainer bash
```

##### Run the Application:
Inside the container, run one of the following commands to start the translation:
```bash
go run main.go "word to translate"
```
or
```bash
/usr/local/bin/run-app "word to translate"
```
### Production
#### Build the Production Container:
```bash
./build.sh
```
#### Start the Production Container:
```bash
./run.sh
```
The run script will fetch the content of your clip board and run its translation.


### Requirements
- Docker: Ensure Docker is installed on your system. [Docker Documentation.](https://docs.docker.com/get-docker/)

- Xclip: Ensure the Xclip utility is installed on your system. [xclip](https://wiki.ubuntuusers.de/xclip/)
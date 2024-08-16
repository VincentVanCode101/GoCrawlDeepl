# Overview
- The `crawl/deepl` project is a Go-based application (`go 1.22.0`) designed to translate the contents of your current clipboard-storage using Deepl's web interface from [your desired language to another desired language](#setup-desired-languages-in-an-env-file) and optionally send the translations via a [Telegram bot](#setup-telegram-bot).
- The application leverages Docker to ensure a consistent and isolated environment for development and execution.
- It serves as a project for me to learn the Golang programming language. Feel free to give feedback or open pull requests.
- This app only works as intended if you have xclip installed [(look here)](#requirements), which mostly only will be linux
    - To run it on MacOS or Windows [have a look at the stuff here](#macos-and-windows)

## Current Setup
### How It Works Inside
- **Starting the application**: The application is run in a [docker container](#docker-setup), when starting it the current clipboard-contents is retrieved and passed as an argument to the go-app which is than to be translated.
- **Initialization**: The application first checks for an active internet connection.
- **Chrome Context**: A Chrome context is created to allow headless browsing.
- **Translation**: The text is translated by navigating to the Deepl website and extracting the translated text using Chromedp.
- **Output**: The translated text is output to the console, and if a Telegram bot is set up, it sends the translation via Telegram.

### How it looks
##### 1. Case
- Lets say I copy this into my clipboard and start the app:
```bash
aforementioned
```

*The The Standart output*:

![stdout.jpg](docs/singlephrase_stdout.jpg)

*The Telegram bot*:

![telegrambot.jpg](docs/singlephrase_telegrambot.jpg)

##### 2. Case
- Lets say I copy this into my clipboard and start the app:
```bash
aforementioned
impeccable
```
... phrases separated by a linebreak

*The The Standart output*:

![alt text](docs/twophrases_stdout.jpg)

*The Telegram bot*:

![alt text](docs/twophrases_telegrambot.jpg)

The phrases are split at the linebreaks and are translated separately

### Setup desired Languages in an `.env` file
- `FROM_LANGUAGE` is the language in which the input is expected
- `TO_LANGUAGE` is the language to which the application input is translated

```bash
FROM_LANGUAGE=
TO_LANGUAGE=
```

### Requirements
- Docker: Ensure Docker is installed on your system. [Docker Documentation.](https://docs.docker.com/get-docker/)
- Xclip: Ensure the Xclip utility is installed on your system. [xclip](https://wiki.ubuntuusers.de/xclip/) [(or see this)](#macos-and-windows)

### Setup Telegram bot
1. Get your Bot-Token and Chat-ID: https://core.telegram.org/bots

2. Create an `.env` file

3. Set the following variables in the `.env` file:
- `BOT_TOKEN`: Your Telegram bot token.
- `CHAT_ID`: Your Telegram chat ID.

```bash
BOT_TOKEN=
CHAT_ID=
```

### Docker Setup

#### Development Environment
- The development environment is defined in `docker-compose.dev.yml` and starts a container named `goCrawlDevtainer` and a container named `webServerDevtainer`.

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
- In case you encounter a
```bash
Error response from daemon: Pool overlaps with other one on this address space
```

- this means the ip-subnet defined in the `docker-compose.dev.yml` overlaps with the address pool of an existing docker network on your machine
- since networks need no time to build again, I just execute

```bash
docker network prune
```
-  and run `docker compose -f docker-compose.dev.yml` again, which solves the problem (when i now I do not have any other docker-network in another project configured with this ip-address pool)

##### Connect to the Development Container:
```bash
docker exec -it goCrawlDevtainer bash
```

### Development - Testing
- To run the unit-tests manualy:
```bash
docker exec -it goCrawlDevtainer bash
make test
```
aftert that? Go visit 
```bash
localhost:8080
```
 where the `webServerDevtainer` has spun up a simple webserver, serving the current test coverages

### Development - Run the application:
Inside the container, run one of the following commands to start the translation:
```bash
make run ARGS="word to translate"
```
#### Keeping the browser open during development

- If you need to keep the browser open for debugging or manual inspection, you can set the `KEEP_BROWSER_OPEN` environment variable to true. This will start the Chrome browser in non-headless mode and keep it open until you press 'Enter' in the terminal.

1. Add the following line to your .env file:
```bash
KEEP_BROWSER_OPEN=true
```
2. When the container is started with `KEEP_BROWSER_OPEN=true`, you can:

##### Connect to the Container via VNC
- You can connect to the running container via VNC to interact with the browser visually. The VNC server is exposed on port 5920. Use a VNC viewer to connect to the following address:

```bash
172.28.0.4:5920
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

# Notes
I developed this with:

`Docker:`

Docker version 26.0.1, build d260a54

`Xclip:`

xclip version 0.13

Copyright (C) 2001-2008 Kim Saunders et al.

Distributed under the terms of the GNU GPL

`OS:`

Ubuntu 22.04.4 LTS x86_64

# MacOS and Windows
- When not running this on linux but MacOS or Windows the functionality differs slightly
- You will not be able to just translate your clipboard contents
- On MacOS you will have to call
```bash
./runWithoutXclip.sh "word to translate"
```
(this assumes you have no xclip installed, otherwise you can just call `./run.sh` and your clipboard contents will be translated)
- On Windows you will have to call
```bash
./run.bat "word to translate"
```
this bat-script is written by Chat-GPT since I have to clue about `.bat` scripts and no way of testing them either (just get yourself a unix-based OS and be happy)


# Note for future Chris that still uses i3:

```
bindsym $mod+Ctrl+P exec --no-startup-id /home/christoph/projects/GoCrawlDeepl/run.sh
```

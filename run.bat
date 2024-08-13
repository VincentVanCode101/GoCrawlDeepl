@echo off

:: Define the path to the .env file
set ENV_FILE=%~dp0.env

:: Collect all command-line arguments into a single string
set args_string=%*

echo All arguments: %args_string%

:: Load the environment variables from the .env file
for /F "tokens=*" %%i in (%ENV_FILE%) do (
    set %%i
)

:: Run the Docker container with the environment variables and arguments
docker run -e TO_LANGUAGE=%TO_LANGUAGE% -e FROM_LANGUAGE=%FROM_LANGUAGE% -e BOT_TOKEN=%BOT_TOKEN% -e CHAT_ID=%CHAT_ID% crawl-deepl:app "%args_string%"

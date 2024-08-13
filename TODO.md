### FUNCTIONALITY:
- store already translated words and look them up before scraping deepl for a word that  has been translated before
- extract XPath into seperate file becaus its likly to change from time to time
- extract ```less common``` words and more translation details from deepl
- add a queue for words so when in the middle of the translation an error occurs, the not translated words are processed the next time the program is ran

### CODE:
- make retrieving words and sending them concurrent:
- now:
    - fetch word 1
    - fetch word 2
    - send word 1
    - send word 2
- wanted:
    - fetch word 1
    - send word 1
    - fetch word 2
    - send word 2



### TEST:
- write unit-tests

### BUILD:
- add running of tests into Dockerfile
- add ci pipeline for automated tests?

- get Makefile out of /app directory:
    - Problem: copying the ```./Makefile``` into the container works just fine,
but the ```- ./app:/usr/src/app``` volume mount from the docker-compose.dev.yml overrides
the ```/usr/src/app``` directory in the container, deleting the Makefile
Volume-mounting: ```- Makefile:/usr/src/app/Makefile``` works to get it into the container
but then an empty Makefile is created on my machine (*/app/Makefile*)... which is a problem

    - possible fix? Create a symlink from ./Makefile to ./app/Makefile ? ... does not feel correct -> ask Dockerexpert for help

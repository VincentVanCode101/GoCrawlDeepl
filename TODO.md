### FUNCTIONALITY:
- ❌ store already translated words and look them up before scraping deepl for a word that  has been translated before
- ❌ extract XPath into seperate file becaus its likly to change from time to time
- 🔨 extract ```less common``` words and more translation details from deepl
- ❌ add a queue for words so when in the middle of the translation an error occurs, the not translated words are processed the next time the program is ran
- ✅ add a timeout to the chromebrowser (e.g. after 5s of no connections, exit the programm)
- add posibility to connect to container via vnc to see open chrome browser (add env-var which controlls if browser is kept open)

### CODE:
- 🔨 make scraping words and sending them concurrent:
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

- ❌ get Jaco to rewrite the program in rust to make it faster
- ⛔ get Esa to rewrite the program in c to make it unsafe
    - but he said maybe he will do it in: https://www.emojicode.org/
- 🔨 main.go should only have the main(){}, remove all other functions
- 🔨 restructure project setup? -> (https://www.youtube.com/watch?v=dxPakeBsgl4) (https://www.youtube.com/watch?v=1ZbQS6pOlSQ)

### TEST:
- 🔨 write unit-tests
- ❌ test against live-deepl system? or just create mock-server
- ❌ save xyz.go file -> run main_test.go and xyz_test.go files -> just like quarkus-auto tests 

### BUILD:
- ✅ add running of tests into Dockerfile
- ❌ add ci pipeline for automated tests?
- ❌ add env var to controlle the size of the browser when viewed with VNC (currently: 3840x2160x16 in the dev-entrypoint.sh)

- 🔨 get Makefile out of /app directory:
    - Problem: copying the ```./Makefile``` into the container works just fine,
but the ```- ./app:/usr/src/app``` volume mount from the docker-compose.dev.yml overrides
the ```/usr/src/app``` directory in the container, deleting the Makefile
Volume-mounting: ```- Makefile:/usr/src/app/Makefile``` works to get it into the container
but then an empty Makefile is created on my machine (*/app/Makefile*)... which is a problem

    - possible fix? Create a symlink from ./Makefile to ./app/Makefile ? ... does not feel correct -> ask Dockerexpert for help

- ⛔ stop commiting to main

#### Legend
✅ = done

🔨 = working on it

❌ = haven't been touchd as off yet

⛔ = Started, but now no intentions of completing (won't do anymore)
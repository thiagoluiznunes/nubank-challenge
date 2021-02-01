# Nubank Challenge

---
Description

**Goal**: You are tasked with implementing an application that authorizes a transaction for a specific account following a set of predefined rules.


Tools: Golang | Docker

### Requerements ###

It's necessary to install Docker previously. 

* **[Docker 20.10.x](https://docs.docker.com)** :white_check_mark:

### Project installation steps ###

**Obs.: The follow instructions were tested on macOS Mojave version 10.14.6**

1 - Unzip file nubank on your directory path:
  - user@user:~/home/$ **unzip nubank-challenge.zip -d /path/to/directory**

2 - Run docker command to create nubank image:
  - use@user:~/home/path/to/directory/$ **docker build -t [YOUR-IMAGE-NAME] -f Dockerfile .**

3 - Running docker image passing OPERATIONS as param:
  - use@user:~/home/path/to/directory/$ **cat [OPERATIONS-PATH] | docker run -i --rm [YOUR-IMAGE-NAME]**
**Example:**
```
$  cat operations | docker run -i --rm nubank
```
  - use@user:~/home/path/to/directory/$ **docker run -i --rm [YOUR-IMAGE-NAME] go run main.go < [[OPERATIONS-PATH]**
**Example:**
```
$  docker run -i --rm nubank go run main.go < operations
```

4 - Running tests
  - use@user:~/home/path/to/directory/$ **docker run -i --rm [YOUR-IMAGE-NAME] go test ./... -count=1 -v**
**Example:**
```
$  docker run -i --rm nubank go test ./... -count=1 -v
```

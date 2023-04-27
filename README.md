# Newsfeed

News app based on RSS feeds

## Installation

Prerequisite

- make [link](https://github.com/wkusnierczyk/make)
- podman [link](https://github.com/containers/podman)
- nodejs [link](https://github.com/nodejs/node)
- go [link](https://github.com/golang/go)

Environment

Setup a file called `.env` in the root directory with the following content:

```
POSTGRES_PASSWORD=postgres
POSTGRES_PORT=5432
POSTGRES_DB_NAME=postgres
BASIC_USER=user
BASIC_PASSWORD=password
```

Start the dev server

- Clone this repo
- Run `make fe_install`
- Run `make fe_dev`
- Run `make postgres_run`
- Run `make be_dev` in a second terminal

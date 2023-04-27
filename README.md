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
POSTGRES_PASSWORD=<value>
POSTGRES_PORT=<value>
POSTGRES_DB_NAME=<value>
BASIC_USER=<value>
BASIC_PASSWORD=<value>
```

Start the dev server

- Clone this repo
- Run `make fe_install`
- Run `make fe_dev`
- Run `make be_dev` in a second terminal

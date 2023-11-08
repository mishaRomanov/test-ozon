[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/mishaRomanov/test-ozon)
# This a small link shortener Gin based-app for ozon tech team 
# Prerequisites
Create a Postgres database with 
``CREATE DATABASE linksdb;``
Or choose your own custom database name with editing configure parameters in ``config/conf.env:``

``USER = "misha"
ADDRESS = "localhost:5432"
DATABASE_NAME = "linksdb"
PASSWORD = ""``

Build an app with `go build cmd/main.go` in `cmd` directory and setup the config as shown earlier 

`./main` to run 

## Usage and how to run
To choose the storage type you have to set it in config file at 
``config/conf.env:`` 
### It's either `cache` or `postgres`
 To short a link use POST method on /link/add endpoint:
``curl --location 'localhost:80/link/add'\
--header 'Content-Type: application/json' \
--data '{"url":"google.com"}'``
Which returns a new link 
``New link generated: localhost:80/link/Nm_5MWFiYz``
To get your old link back, you make a GET request to /link/*your_link* endpoint 
``curl --location 'localhost:80/link/NjA2NjkyZT'``
And it returns you full link in JSON
``"google.com"``
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/mishaRomanov/test-ozon)
# THIS CURRENTLY DOESN'T WORK PROPERLY!!! 
# ---------
# This a small link shortener Gin based-app for ozon tech team 
# Prerequisites 
Have Postgres installed and launched
Have Docker installed and launched 
# How to get app working 
Clone my repository to your local machine `git clone github.com/mishaRomanov/test-ozon`

Run with docker compose 
`docker compose up -d`

# How to use
To make link shorter use POST method on /link/add endpoint:
``curl --location 'localhost:8080/link/add'\
--header 'Content-Type: application/json' \
--data '{"url":"google.com"}'``

Which returns a new link 
``New link generated: localhost:8080/link/Nm_5MWFiYz``


To get your old link back, you make a GET request to /link/*your_link* endpoint
``curl --location 'localhost:8080/link/NjA2NjkyZT'``

And it returns you full link in JSON
``"google.com"``

Also feel free to use `/about` endpoint to get all the information you need to test the app.

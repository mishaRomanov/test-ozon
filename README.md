# This a small link shortener Gin based-app for ozon tech team 
## Usage and how to run
To run app use ``go run main.go --storage-type='x'`` *where x is either* ``postgres`` or ``cache`` 
### To short a link use POST method on /link/add endpoint:
``curl --location 'localhost:80/link/add'\
--header 'Content-Type: application/json' \
--data '{"url":"google.com"}'``
### Which should return you something like this
``New link generated: localhost:80/link/Nm_5MWFiYz``
### To get your old link back, you make a GET request to /link/*your_link* endpoint 
``curl --location 'localhost:80/link/NjA2NjkyZT'``
### And it returns you full link in JSON
``"google.com"``
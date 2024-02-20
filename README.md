# Invit.io

## A project to make it easier for you to invite people to your events

use `docker-compose up --build -d`
to build and run the project locally


## Getting started

1) `docker-compose up --build -d`
2) use postman to send a POST request to `http://localhost:3202/invite` with body
    ```
        {
        "organiser":"123", 
        "location":"Manchester", 
        "date":"tomorrow",
        "passphrase": "welcome"
        }
    ```

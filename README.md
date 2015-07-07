Go-starter
=========

A starter structure for web development with Golang. Has some pieces in place for ReactJS development.

##Setup
First make sure you have the following installed:

    https://github.com/codegangsta/gin
    https://github.com/chiedojohn/shell-dotenv

Then to install all pacakages the project depends on:

    go get ./...

Create an `.env` file with the following contents for holding your local environment variables.

    STATIC_DIR=/static
    GO_ENV=development
    # Update the following with accurate data
    MYSQL_DATABASE=database
    MYSQL_USER=username
    MYSQL_PASS=password
    MYSQL_HOSTNAME=host
    MYSQL_PORT=port

Create your local database in accordance with the name specified in your .env file

Run the migrations. We are using Goose for migrations. The documentation is sparse but refer to it. Run all goose commands in the following format in order to ensure that goose takes note of your env variables `dotenv goose args`. Run the below to update your database:
    
    dotenv goose up

##Development
To run the web server and have it automatically update upon making changes run the following (appears convoluted due to our use of gom):

    gin run

Then you can visit the app in your browser at 127.0.0.1:3000 (assumes that port 3001 is serving the app to the gin reverse proxy at port 3000).

###ReactJS
It will be up to you to get your ReactJS code in the static/js directory as react-bundle.js

##Deployment
To build your application, run:

    gom build main.go


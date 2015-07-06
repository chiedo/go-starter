Go-starter
=========

A starter structure for web development with Golang. Has some pieces in place for ReactJS development.

Actively being developed... Currently not ready for anything.

##Setup
First make sure you have gin installed:

    go get github.com/codegangsta/gin

Then to install your pacakages:

    go get ./...

Lastly, create a `.env` file with the following contents for holding your local environment variables.

    STATIC_DIR="/static"

##Development
To run the web server and have it automatically update upon making changes run the following (appears convoluted due to our use of gom):

    gin run

Then you can visit the app in your browser at 127.0.0.1:3000 (assumes that port 3001 is serving the app to the gin reverse proxy at port 3000).

###ReactJS
It will be up to you to get your ReactJS code in the static/js directory as react-bundle.js

##Deployment
To build your application, run:

    gom build main.go


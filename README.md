Go-starter
=========

A starter structure for web development with Golang. Has some pieces in place for ReactJS development.

Actively being developed... Currently not ready for anything.

##Setup
First make sure you have the gom Go manager and gin installed:

    go get github.com/mattn/gom
    go get github.com/codegangsta/gin

Then to install your pacakages (will be stored in \_vendor):

    gom install

Lastly, create a `.env` file with the following contents for holding your local environment variables.

    MESSAGE="hello world"
    STATIC_DIR="/static"

##Development
To run the web server and have it automatically update upon making changes run the following (appears convoluted due to our use of gom):

    env GOPATH='_vendor' gin -a 8080 -p 4000 run main.go

Then you can visit the app in your browser at 127.0.0.1:4000. Note that you using this approach creates a proxy that you must use,
so although your server is set to port 8080 in the main.go file, you must use port 4000.

###ReactJS
It will be up to you to get your ReactJS code in the static/js directory as react-bundle.js

##Deployment
To build your application, run:

    gom build main.go


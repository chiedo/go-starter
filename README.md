Go-starter
=========

A starter structure for web development with Golang.
Actively being developed... Currently not ready for anything.

##Setup
First make sure you have the gom Go manager installed:

    go get github.com/mattn/gom

Then to install your pacakages (will be stored in \_vendor):

    gom install

Lastly, create a `.env` file with the following contents for holding your local environment variables.

    MESSAGE="hello world"

##Development
To run the web application, run the following:

    gom run main.go

Then you can visit the app in your browser at 127.0.0.1:8080 

##Deployment
To build your application, run:

    gom build main.go


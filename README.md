# Simple analysis tool in golang
A simple analysis tool made in golang to find out the spread of questions across related tags

This application entirely depends upon Api's published by StackOverflow.
This app is an example of using Nvd3 charts, gorilla websocket, Stack api with golang's 
templates. 

On executing the application a http server will be launched at port :4748

## Environment set up
Download the latest stable Go distribution(go1.5) from the golang site and follow
the instruction mentioned.
  
	https://golang.org/doc/install
	https://golang.org/dl/

## Building Source Code

Open terminal at application directory, execute "go build" command. This will build 
entire application. More instruction can be found at

	https://golang.org/cmd/go/

## How to Use
 	
Create a configuration file by name app.properties file in application 
directory with the following details.
	
	APP_ID= 
	API_KEY_ID=
	API_KEY_SECRET=
	REDIRECT_URI=(#Replace this with your app location)/authenticated

###### These are the details fetched by registring app to [Stack App](http://stackapps.com/)
			
	
## Starting webserver

On successfull build, an executable will be created at the build location. 
Execute that from terminal. It will launch a webserver at port 4748
Pull the login page.

######  http://localhost:4748	

## Code Structure

	Main.go    :Http Server's entry point 
	hub.go		:Websockets logic
	search.go	:Stackoverlfows corresponding logic 
	session.go	:Stackoverlfows corresponding logic
	Structs.go	:Data structure to hold Json result.
	Utils.go	:Few routines.. 
	

## Visit application deployed to heroku.
######[Application on heroku ](https://soanalysis.herokuapp.com/)

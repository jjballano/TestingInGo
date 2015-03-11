How to deploy in heroku
========

In the directory where you main program is, run:

	$ go get

This should create a executable file in $GOPATH/bin. In this example, it will create a executable called herokudeploy

In the root dir, you have to create the Procfile that Heroku needs:

	$ echo 'web: herokudeploy' > Procfile

Change herokudeploy for the proper one.

Install Godep

	$ go get github.com/kr/godep

Save dependencies

	$ godep save	

Create a .godir	in the root directory, which must contain the full package path to the application. In this example:

	$ echo 'github.com/jjballano/learningGo' > .godir

	
Commit all changes.

	$ git add .
	$ git commit -m "any comment"

Create a new Heroku app, telling it to use Go Heroku Buildpack

	$ heroku create -b https://github.com/kr/heroku-buildpack-go.git

Push changes to heroku:

	$ git push heroku master

References:
	http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku 
	https://flynn.io/docs/how-to-deploy-go

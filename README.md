## Hello World Bday Application

### This is a simple bday application in which you will get a response from app as N days remaining to your birthday or if it's your birthday today, Happy Birthady greeting message. 

## Overview

In this project I am trying to setup Jenkins pipeline for web application on top of Kubernetes cluster.

![Architecture diagram](images/architecture.png)

This is a simple 2-tier application where web application is running using Go and users are stored in backend database MySQL, respectively.

To locally run and test MySQL running into Docker container using the following command(using an official MySQL docker image), 

    $ docker container run --name mysqldb -v /opt/datadir:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

where, -v /opt/datadir:/var/lib/mysql telling that /opt/datadir which is on my host should be mapped to var/lib/mysql which is inside container. We do this so that we would have all MySQL created database files on our computer not just in docker container.

Grab the mysql IP and and use it in your go mysql driver while connecting to this database.

    $ docker container inspect mysql | grep IPAddr

And, for running go lang application I have use multi-stage docker image where we'll build go binary and copy it into tiny alpine docker image to run an application. 

    $ docker build -t SimpleBdayApp -f Dockerfile .

    $ docker container run -t simplebdayapp -p 8080:8080 SimpleBdayApp

Once docker images are ready we can push a docker image into [Docker Hub](https://hub.docker.com/) centralized repository for future reference and using this docker image.

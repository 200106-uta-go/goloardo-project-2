#!/bin/bash
#become root user
sudo su
#upgrade basic packages for a new start of a instance
apt update
apt upgrade

#install go and docker
apt install docker.io

#docker swarm join-token worker

#pull each docker image for webApp, database, and controller
docker pull felicianoej/arkcontroller
docker pull codezipline/dbserver
docker pull danish287/horsocoped

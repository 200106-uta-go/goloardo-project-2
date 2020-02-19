#!/bin/bash
#become root user
sudo su
#upgrade basic packages for a new start of a instance
apt update
apt upgrade
#download java development kit
apt install openjdk-8-jdk
#download jenkins with a key
wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | sudo apt-key add -
#append this to sources list
echo -e "deb https://pkg.jenkins.io/debian binary/\n" >> /etc/apt/sources.list
#UPDATE again due to newly downloaded jenkins
apt update
#install jenkins
apt install jenkins
#install go and docker
apt install golang
apt install docker.io


#export MYIP=`curl ifconfig.me`
#docker swarm join-token worker

#pull each docker image for webApp, database, and controller
docker pull felicianoej/arkcontroller
docker pull codezipline/dbserver
docker pull danish287/horsocoped

#create the overlap network for communication between networks
#docker network create -d overlay --attachable onet

#run arkcontrol
docker service create --name horsocoped --expose 8080 --mode replicated --replicas=1 --network onet danish287/horsocoped
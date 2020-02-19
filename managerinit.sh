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


export MYIP=`curl ifconfig.me`
#init a docker swarm 
docker swarm init --advertise-addr $MYIP
#write to w_token.txt the worker token to accept worker nodes
sudo docker swarm join-token worker -q > w_token.txt  
# docker swarm join --token SWMTKN-1-0kzbg8ex75h1nufvng9589spp0hrnr3i5h5bz5iqizwf62oirl-dpe91dl6yzz1z0wufg9udd27s 54.200.205.192:2377      

docker pull felicianoej/arkcontroller
docker pull codezipline/dbserver
docker pull danish287/horsocoped
docker network create -d overlay --attachable onet

#run arkcontrol
docker service create --name arkcontroller --publish published=7777,target=7777 --mode replicated --replicas=1 --network onet felicianoej/arkcontroller







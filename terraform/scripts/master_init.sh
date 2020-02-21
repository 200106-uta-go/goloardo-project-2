#!/bin/bash

# export my machine IP
# export MYIP=`curl ifconfig.me`
# echo $MYIP

#initialize a docker swarm 
#docker swarm join-token worker
echo Initializing Docker swarm...
sudo docker swarm init

echo Retrieving docker continers...
sudo docker pull felicianoej/arkcontroller
sudo docker pull codezipline/dbserver
sudo docker pull danish287/horsocoped
sudo docker network create -d overlay --attachable onet


#run arkcontrol
echo Initializing ark-controller...
sudo docker service create --name arkcontroller --publish published=7777,target=7777 --mode replicated --replicas=1 --network onet felicianoej/arkcontroller

echo Getting github scripts... 
wget https://raw.githubusercontent.com/200106-uta-go/goloardo-project-2/master/scripts/serviceinit.sh
wget https://raw.githubusercontent.com/200106-uta-go/goloardo-project-2/master/scripts/pop.sh

echo Changing file permissions...
sudo chmod u+s /usr/bin/docker
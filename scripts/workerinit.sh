#!/bin/bash
#upgrade basic packages for a new start of a instance
sudo apt update
sudo apt upgrade

#install go and docker
sudo apt install docker.io

#joining the swarm 
docker swarm join --token SWMTKN-1-38g0e9bdl1ylcxbjbm1miwd3pjaziqgw5u2qduvy0du9iwgmaa-5nxiemkmhsq3goucvmhupw4wq 172.31.43.70:2377

#pull each docker image for webApp, database, and controller
sudo docker pull felicianoej/arkcontroller
sudo docker pull codezipline/dbserver
sudo docker pull danish287/horsocoped

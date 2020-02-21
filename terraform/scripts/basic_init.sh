#!/bin/bash
echo Beginning set-up...
sudo apt-get update 
#upgrade basic packages for a new start of a instance
# sudo apt update -y
# sudo apt-get upgrade


echo Installing Docker...
# install docker 
sudo apt-get -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get -y install docker-ce docker-ce-cli containerd.io


echo Installing Java development kit...
#download java development kit
sudo apt -y install openjdk-8-jdk 
#download jenkins with a key
wget http://mirrors.jenkins.io/war-stable/latest/jenkins.war
# sudo wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | sudo apt-key add -
#append this to sources list
# sudo echo -e "deb https://pkg.jenkins.io/debian binary/\n" >> /etc/apt/sources.list


echo Updating again...
#UPDATE again due to newly downloaded jenkins
sudo apt-get update 

#installing jenkins
echo Installing Jenkins...
sudo apt -y install jenkins

#installing go 
echo Installing Go...
sudo apt -y install golang


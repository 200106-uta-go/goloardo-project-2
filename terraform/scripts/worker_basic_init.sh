echo Updating and upgrading machine...
sudo apt-get update

echo Installing Docker...
# install docker 
sudo apt-get -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get -y install docker-ce docker-ce-cli containerd.io

echo Joining swarm...
#joining the swarm 
docker swarm join --token SWMTKN-1-02hztjrn88qjif9avnzjmnsj6molkyo3l5slhpxww3tj971kt6-9austc43da7gqv78mvteh5sqx 3.20.120.110

echo Downloading Images...
#pull each docker image for webApp, database, and controller
sudo docker pull felicianoej/arkcontroller
sudo docker pull codezipline/dbserver
sudo docker pull danish287/horsocoped
#Makefile called in project-2 directory to start server within a container

#Defualt goal with basic call to make.
all: dockDbServer

#Start the docker container(dbserverContainer) with myDB volume attached. 
dockDbServer:
	docker volume create --name myDB
	docker build -t codezipline/dbserver .
	docker run -it --expose 8081 --name dbserverContainer -v myDB:/app/badger codezipline/dbserver

#Inspects IP of the container
dockContainerIP:
	sudo docker inspect -f "{{ .NetworkSettings.IPAddress }}" dbserverContainer

#Remove all container volumes and images from dockDbServer command from docker
clean:
	docker system prune
	docker volume rm myDB
	docker image rm codezipline/dbserver

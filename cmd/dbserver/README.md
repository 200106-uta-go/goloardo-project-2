## Getting Started
***Note:*** cmd/config.json file stores the data directory configuration in string format defining where the data will be stored
   
## Docker

### Commands for Execution

#### Creating a Docker Volume 
    
Volume will be be shared between host machine and container

Command  ``` sudo docker volume create --name <volumeName> ```

To name volume myDB run the following command
    
``` sudo docker volume create --name myDB ```
    
#### Create a docker image with a tag for dockerhub, with the format of ``` <userName> / <repositoryName>. ```
    
***Note:*** Dockerfile contains preconfigured settings (refer to Makefile section below) 

Command ``` docker build -t codezipline/dbserver ./cmd ```
    
#### Create an apline container with the volume, myDB, attach, to /app/badger in the container, on port 8081

 Command ``` sudo docker run -it expose <portNumber> --name <containerName> -v <volumeName>:<containerDirectoryPath> <imageName> ```

In our case, we run  ``` sudo docker run -it --expose 8081 --name dbserverContainer -v myDB:/app/badger codezipline/dbserver ```
    
    
#### To obtain ip address of the container with Container_Name we run the following command  

``` sudo docker inspect -f "{{ .NetworkSettings.IPAddress }}" dbserverContainer ``` 

giving us the container name ``` dbserverContainer ```

## Makefile

``` Makefile ``` is a convenient configuration script of the docker container with an attached volume. The container hosts a badger database and stores to volume located in /app/badger within the container. 

 ***Note:*** by convention a project should always run the make clean method before executing other commands

1. ```make```

    The default ``` make ``` command starts the database and hosts it on port 8081. 

2. ``` make <target name> ```

    Specificing a method can be done with the command ``` make <target name> ```

3.  ``` make clean ```

    ``` make clean ``` will prune the system and remove the recently made image and container
    
   



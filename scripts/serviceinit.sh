#!/bin/bash

docker service create --name horsocoped --expose 8080 --mode replicated --replicas=1 --network onet danish287/horsocoped
docker service create --name dbserver --expose 8081 --mode replicated --replicas=1 --network onet codezipline/dbserver
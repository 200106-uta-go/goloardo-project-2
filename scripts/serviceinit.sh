#!/bin/bash

# Must replace the ip with private ip of the manager 
docker service create --name dbserver --publish published=8081,target=8081 --mode replicated --replicas=1 --network onet codezipline/dbserver -ark=10.0.1.2
docker service create --name horoscoped --publish published=80,target=8080 --mode replicated --replicas=5 --network onet danish287/horsocoped -ark=10.0.1.2

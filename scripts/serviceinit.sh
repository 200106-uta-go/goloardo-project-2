#!/bin/bash

# Must replace the ip with private ip of the manager 
docker service create --name dbserver --publish published=8081,target=8081 --mode replicated --replicas=1 --network onet codezipline/dbserver -ark=10.0.0.2
docker service create --name horoscoped --publish published=80,target=8080 --mode replicated --replicas=1 --network onet danish287/horsocoped -ark=10.0.0.2
docker service create --name pop --network onet --mode replicated --replicas=1 --restart-condition none felicianoej/p2dbpop -ip=10.0.0.2
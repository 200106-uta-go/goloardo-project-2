#!/bin/bash 

# Needs ip for database(private)
docker service create --name pop --network onet --mode replicated --replicas=1 --restart-condition none felicianoej/p2dbpop -ip=
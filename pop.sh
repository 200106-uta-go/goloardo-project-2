#!/bin/bash 

# Needs ip for database(private)
export DBIP=`sudo docker service inspect dbserver | grep Addr | tail -1 | sed 's/"Addr": "//' | sed 's/"//' | cut -f1 -d"/" | sed -e 's/^[ \t]*//'`
docker service create --name pop --network onet --mode replicated --replicas=1 --restart-condition none felicianoej/p2dbpop -ip=$DBIP
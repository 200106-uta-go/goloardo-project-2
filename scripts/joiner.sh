#!/bin/bash

#Read from Present working directory and run this command to join swarm for worker
#docker swarm join --token ${w_token.txt} EC2PUBLIC_IP:2377
#docker swarm join --token SWMTKN-1-0kzbg8ex75h1nufvng9589spp0hrnr3i5h5bz5iqizwf62oirl-dpe91dl6yzz1z0wufg9udd27s 54.200.205.192:2377


docker swarm join --token "$(<w_token.txt)" "$(<manager_ip.txt)":2377
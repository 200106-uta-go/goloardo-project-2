#!/bin/bash

docker swarm join --token "$(<w_token.txt)" "$(<manager_ip.txt)":2377
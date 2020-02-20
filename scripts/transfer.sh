#!/bin/bash

#Pull from manager to .(pwd)
#scp -i ec2key.pem username@ec2ip:/path/to/file . 
#scp -i YOURKEY.pem ubuntu@EC2PUBLIC_IP:/home/ubuntu .

scp -i YOURKEY.pem ubuntu@EC2PUBLIC_IP:/home/ubuntu .


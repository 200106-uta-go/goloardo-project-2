Dania, Josh, Emilio

# Project 2 - Overlay Networking
Working in groups of 3 or 4, deploy a multi-instance server and VNFs to a production environment on a cloud services platform of your choice. There must be communication across separate virtual machines managed by an SDN controller and service mesh, separate applications controlling and routing traffic between them. Infrastructure and application changes should be managed by an automated CI/CD pipeline.

## Recommended Tools
- DockerHub 
- AWS/GCE/Azure 
- Jenkins/CircleCI/TravisCI
- Terraform
- Kubernetes

## Requirements
- [ ] Documentation
- [ ] Agile/Scrum Project Management
- [ ] Git Branching & Semantic Versioning
- [ ] (Cloud) Production Environment
- [ ] CI/CD Pipeline
- [ ] Infrastructure as Code
- [ ] Orchestration

## Presentation
- [ ] 15-minute Demonstration
- [ ] Presentation Slides


#Pull from manager to .(pwd)
scp -i ec2key.pem username@ec2ip:/path/to/file . 
scp -i YOURKEY.pem ubuntu@EC2PUBLIC_IP:/home/ubuntu .
#Read from Present working directory and run this command 
docker swarm join --token ${w_token.txt} EC2PUBLIC_IP:2377
docker swarm join --token SWMTKN-1-0kzbg8ex75h1nufvng9589spp0hrnr3i5h5bz5iqizwf62oirl-dpe91dl6yzz1z0wufg9udd27s 54.200.205.192:2377
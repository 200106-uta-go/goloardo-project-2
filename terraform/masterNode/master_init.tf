# pull access and ecret keys from local .txt files && adjust the region
provider "aws" {
  access_key = "${file("../keys/accesskey.txt")}"
  secret_key = "${file("../keys/privatekey.txt")}"
  region     = "us-east-2"
}

# create security group so we can ssh from any machine at initialization
resource "aws_security_group" "master_SSH" {
  description = "Allow ALL SSH traffic through"
  name = "ssh-master"

  ingress {
    from_port   = 0 
    to_port     = 0
    protocol =   "-1"

    cidr_blocks =  ["0.0.0.0/0"]
  }

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}

# assign an static ip to the master ec2 instance
resource "aws_eip" "ip" {
    vpc = true
    instance = aws_instance.master.id
}

#initialize master node
resource "aws_instance" "master" {
  key_name      = "myMac"
  ami           = "ami-0fc20dd1da406780b"
  instance_type = "t2.medium"

# create security group to be able to ssh
  security_groups = [aws_security_group.master_SSH.name]

# establish ssh connection to machine 
  connection {
    type        = "ssh"
    user        = "ubuntu"
    private_key = "${file("../keys/myMac.pem")}" 
    host        = self.public_ip
  }

# script file to run inside our machine and set up master
  provisioner "file" {
    source      = "../scripts/master_init.sh"
    destination = "/tmp/master_init.sh"
  }

# script file to run inside our machine for basic init
  provisioner "file" {
    source      = "../scripts/basic_init.sh"
    destination = "/tmp/basic_init.sh"
  }

# terraform master file 
  provisioner "file" {
    source      = "master_init.tf"
    destination = "/home/ubuntu/terraform/master_init.tf"
  }

# save ip address of instance in a txt file to save to our machine
  provisioner "local-exec" {
    command = "echo ${aws_instance.master.public_ip} > ip_address.txt"
  }

# run these commands inside our instance
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /tmp/basic_init.sh",      
      "sudo /bin/bash /tmp/master_init.sh",
    ]
  }
}





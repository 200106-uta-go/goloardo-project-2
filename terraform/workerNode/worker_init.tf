# pull access and ecret keys from local .txt files && adjust the region
provider "aws" {
  access_key = "${file("../keys/accesskey.txt")}"
  secret_key = "${file("../keys/privatekey.txt")}"
  region     = "us-east-2"
}

# create security group so we can ssh from any machine to worker worker 
resource "aws_security_group" "worker_SSH" {
  description = "Allow ALL SSH traffic through"
  name = "ssh-workers"

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

variable "instance_count" {
  default = "2"
}

#initialize db worker node
resource "aws_instance" "worker" {
  count         = "${var.instance_count}"
  key_name      = "myMac"
  ami           = "ami-0fc20dd1da406780b"
  instance_type = "t2.micro"

# create security group 
  security_groups = [aws_security_group.worker_SSH.name]

# establish ssh connection to machine 
  connection {
    type        = "ssh"
    user        = "ubuntu"
    private_key = "${file("../keys/myMac.pem")}" 
    host        = self.public_ip
  }


# script file to run inside our machine for basic setup
  provisioner "file" {
    source      = "../scripts/worker_basic_init.sh"
    destination = "/tmp/worker_basic_init.sh"
  }



# run these commands inside our instance
  provisioner "remote-exec" {
    inline = [
      "sudo /bin/bash /tmp/worker_basic_init.sh",      
    ]
  }
}





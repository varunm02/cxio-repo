provider "aws" {
	region = "us-east-1"
  }

 
  resource "aws_security_group" "sg3" {
	name        = "allow_tls"
	
	dynamic "ingress"{
		for_each = [22,80,443]
		itertator = port
		content {
		description = "Allow TLS inbound traffic"
		from_port   = port.value
        to_port     = port.value
		protocol    = "tcp"
		cidr_blocks = ["0.0.0.0/0"]
			
		}
	}
	tags = {
	  Name = "mysg3"
	}
  }

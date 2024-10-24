provider "aws" {
	region = "us-east-1"
  }

  terraform {
	backend "s3" {
	  bucket = "varunbuck1"
	  key    = "path/to/terraformtf.state"
	  region = "us-east-1"
	}
  }
  

  resource "aws_instance" "ec2_instance" {
    ami = "ami-03a6eaae9938c858c"
    count = 1
    instance_type = "t2.micro"
    key_name = "link2"
	tags = {
    Name = "ec21"
  }
	
} 
provider "aws" {
	region = "us-east-1"
  }
  resource "aws_instance" "ec2_example" {

    ami = "ami-03a6eaae9938c858c"  
    instance_type = "t2.micro" 
    key_name= "key2"
    vpc_security_group_ids = [aws_security_group.main.id]
	provisioner "remote-exec" {
		inline = [
		  "touch hello.txt",
		  "echo helloworld remote provisioner >> hello.txt",
		]
	  }
	  connection {
		type        = "ssh"
		host        = self.public_ip
		user        = "ec2-user"
		private_key = file("/home/ec2-user/varun/key2")   
		timeout     = "4m"
	 }
  }
  resource "aws_security_group" "main" {
	egress = [
	  {
		cidr_blocks      = [ "0.0.0.0/0", ]
		description      = ""
		from_port        = 0
		ipv6_cidr_blocks = []
		prefix_list_ids  = []
		protocol         = "-1"
		security_groups  = []
		self             = false
		to_port          = 0
	  }
	]
   ingress                = [
	 {
	   cidr_blocks      = [ "0.0.0.0/0", ]
	   description      = ""
	   from_port        = 22
	   ipv6_cidr_blocks = []
	   prefix_list_ids  = []
	   protocol         = "tcp"
	   security_groups  = []
	   self             = false
	   to_port          = 22
	}
	]
  }
  resource "aws_key_pair" "deployer" {
	key_name   = "key2"
	public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCdrDptQ88VZz08jw73h740aRgHtfl1Ph9i+9+WPDHMjsYDTD1zRpvGrgU9KRxktEGTPnmAvvbL+6x+5u7D/oKXv5HcAfFpOEEGUNJODr8J0If7kFOdPzWdyaAAL/jaqiIw+cZshUOQF8ESOnJfy+NRyQ4yCfvewSnbTIS9bP6tmClzN0pUY69fOxJGQ8eFe7qUOpKwS1sC4jwkWrbpvVTouFLaSvC8iRhVBMWbkvJw4oz5M6uUQu3sTGYTBMnojIpFKJxGUtUgfjFdHP4OwpzLYIKGZBUnVN7sEa6GrqZX9ekJNnrYxMSr++gsvsIGCV7kho90sxj0Ae4c8eAr7/YOYN/RUqx9dnRlAIBGW4X8v1lgZY+QlQ9acMMfkMa7E0uUwVYMB45a9UHnBsPN8kVvby9UgPN56UpTFnkIFhFD8N9O7/7J5C86r1Ob/xd7b3Pr0sa34R2enr6YsegEE9fEN7YOFItuiO+Jl9G3a6W1zkY+k853WOgS6WAEv4avokc= root@ip-172-31-94-174.ec2.internal
	"
  }


####    INSTRUCTIONS FOR PROVISION #######
  #####create ssh-keygen -t rsa in working directory of terraform only   
  ##that is if your terraform file is in varun then your keys should also be generated in varun only see the path of code and act accordingly
  #### aslo you need to create vpc security groups instead of Normal security groups
  ##in public key make sure to copy paste contents of public which was ccreated in ssshkeygen
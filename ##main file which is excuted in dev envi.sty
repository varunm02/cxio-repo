##main file which is excuted in dev environment main.tf file##

provider "aws" {
	region = "us-east-1"
}

module "my_vpc" {
	source = "../modules/vpc"
  vpc_cidr = "10.0.0.0/16"
subnet_cidr = "10.0.1.0/24"
	vpc_id = "${module.my_vpc.vpc_id}"
}
module "myec2"{
	source = "../modules/ec2"
 subnet_id = "${module.my_vpc.subnet_id}"


## ec2directory file as ec2.tf##

 resource "aws_instance" "ec2_instance" {
    ami = "ami-03a6eaae9938c858c"
    count = 1
    subnet_id = "${var.subnet_id}"
    instance_type = "${var.instance_type}"
    key_name = "jkl"
                tags = {
    Name = "ec21"
  }
##ec2 var.tf file ##
}
variable "instance_type" {
	default = "t2.micro"
	}
	
	variable "subnet_id" {}

	

## vpc directory file as vpcnetwork.tf.tf##



	resource "aws_vpc" "this" {
        cidr_block       = "${var.vpc_cidr}"
        instance_tenancy = "default"

        tags = {
          Name = "varun"
        }
  }
resource "aws_subnet" "sub1" {
        vpc_id     = "${var.vpc_id}"
        cidr_block = "${var.subnet_cidr}"

        tags = {
          Name = "subnet_pub"
        }
  }

output "vpc_id"{
value = "${aws_vpc.this.id}"
}

output "subnet_id"{
value = "${aws_subnet.sub1.id}"
}

## vpc var.tf file##

variable "vpc_cidr" {
	default = "10.0.0.0/16"
	}
variable "vpc_id" {}

variable "subnet_cidr"{
default = "10.0.1.0/24"
}

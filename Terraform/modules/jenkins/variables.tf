variable "vpc_id" {
  description = "VPC ID where Jenkins will be deployed"
  type        = string
}

variable "subnet_id" {
  description = "Subnet ID where Jenkins will be deployed"
  type        = string
}

variable "ami_id" {
  description = "AMI ID for Jenkins server (Amazon Linux 2 recommended)"
  type        = string
  default     = "ami-0735c191cf914754d"  # Amazon Linux 2 in us-west-2
}

variable "instance_type" {
  description = "Instance type for Jenkins server"
  type        = string
  default     = "t3.medium"
}

variable "key_name" {
  description = "Name of the SSH key pair"
  type        = string
}

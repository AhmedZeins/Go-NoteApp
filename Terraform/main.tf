terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.23"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.11"
    }
  }
  required_version = ">= 1.0"
}

provider "aws" {
  region = var.aws_region
}

provider "kubernetes" {
  host                   = module.eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    args = [
      "eks",
      "get-token",
      "--cluster-name",
      module.eks.cluster_name
    ]
  }
}

provider "helm" {
  kubernetes {
    host                   = module.eks.cluster_endpoint
    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
    exec {
      api_version = "client.authentication.k8s.io/v1beta1"
      command     = "aws"
      args = [
        "eks",
        "get-token",
        "--cluster-name",
        module.eks.cluster_name
      ]
    }
  }
}

# VPC Module
module "vpc" {
  source = "./modules/vpc"
  
  vpc_cidr             = var.vpc_cidr
  availability_zones   = var.availability_zones
  private_subnet_cidrs = var.private_subnet_cidrs
  public_subnet_cidrs  = var.public_subnet_cidrs
  cluster_name         = var.cluster_name
}

# Security Module
module "security" {
  source = "./modules/security"
  
  vpc_id = module.vpc.vpc_id
}

# IAM Module
module "iam" {
  source = "./modules/iam"
  
  cluster_name = var.cluster_name
}

# Bastion Host Module
module "bastion" {
  source = "./modules/bastion"
  
  cluster_name          = var.cluster_name
  vpc_id               = module.vpc.vpc_id
  public_subnet_id     = module.vpc.public_subnet_ids[0]
  allowed_ssh_cidr_blocks = var.allowed_ssh_cidr_blocks
  key_name             = var.ssh_key_name
}

# EKS Module
module "eks" {
  source = "./modules/eks"
  
  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version
  
  vpc_id          = module.vpc.vpc_id
  private_subnets = module.vpc.private_subnet_ids
  
  cluster_role_arn     = module.iam.cluster_role_arn
  node_group_role_arn  = module.iam.node_group_role_arn
  
  security_group_ids        = [module.security.cluster_security_group_id]
  bastion_security_group_id = module.bastion.bastion_security_group_id
}

# Jenkins Module
module "jenkins" {
  source = "./modules/jenkins"

  vpc_id    = module.vpc.vpc_id
  subnet_id = module.vpc.public_subnets[0]
  key_name  = "your-key-pair-name"  # Replace with your SSH key pair name
}

# EC2 Module
module "ec2_instance" {
  source = "./modules/ec2"

  instance_name    = "jenkins-server"
  instance_type    = "t2.medium"  # Recommended for Jenkins
  aws_region      = "us-east-1"
  environment     = "dev"
}

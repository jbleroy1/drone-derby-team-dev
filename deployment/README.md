# Deployment

This folder contains Terraform configuration for deploying the base infrastructure for the Drone Derby.

This Terraform will deploy all of the resources required to enable your autonomous drone.

# Instructions
1. Git clone the public repo
2. Terraform init
3. Terraform apply
4. gcloud trigger deploy

5. Copy `terraform.tfvars.example` to `terraform.tfvars`, updating variables as appropriate 
6. Ensure you have gcloud configured and authenticated locally
7. Run `terraform apply`


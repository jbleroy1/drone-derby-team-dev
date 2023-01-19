# Infrastructure

This folder contains Terraform configuration for deploying the base infrastructure for the Team Drone Derby repository. 

State is stored in the drone-derby-dev [drone-derby-dev project](https://console.cloud.google.com/compute/instances?project=drone-derby-dev) in a GCS Bucket named `gs://drone-derby-terraform-state`

## Initial Setup

1. Setup Environment
```
# Set below vars
PROJECT_ID=drone-derby-dev
STATE_BUCKET_NAME=drone-derby-terraform-state

gcloud config set project $PROJECT_ID
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format 'value(projectNumber)')
```

2. Create GCS Bucket
```
gcloud config set project $PROJECT_ID
gsutil mb -l europe-west1 gs://$STATE_BUCKET_NAME
gsutil versioning set on gs://$STATE_BUCKET_NAME
```

3. Grant GCB in `drone-derby-dev` (or alternate project) access to GCS Bucket
```
gsutil iam ch serviceAccount:$PROJECT_NUMBER@cloudbuild.gserviceaccount.com:objectAdmin gs://$STATE_BUCKET_NAME
```

4. Grant GCP Permission on Projects (Repeat for additional environments)
```
gcloud projects add-iam-policy-binding $PROJECT_ID --member=serviceAccount:$PROJECT_NUMBER@cloudbuild.gserviceaccount.com --role=roles/editor
```

5. Apply Build Triggers
```
gcloud beta builds triggers import --source=build/triggers/infrastructure/base/dev.yaml --project=$PROJECT_ID
```

## Running Locally

1. Set Environment
```
_ENVIRONMENT_=dev
_DEPLOY_PROJECT_=drone-derby-dev
_DEPLOY_REGION_=europe-west1
```


2. Initialise Terraform
```
terraform init -backend-config=prefix=$_ENVIRONMENT_
```

2. Apply Config
```
terraform apply -var project=$_DEPLOY_PROJECT_ -var region=$_DEPLOY_REGION_
```

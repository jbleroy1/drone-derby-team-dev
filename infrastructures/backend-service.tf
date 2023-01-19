/*
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
 */

resource "google_service_account" "cloud-run-backend-service" {
  account_id   = "bootstrap-service"
  display_name = "Bootstrap Service"
  description  = "[Terraform Managed] Service Account for the Bootstrap Service, responsible for initialising new projects"
}


resource "google_cloud_run_service" "cr-backend-service" {
  name     = "cr-backend-service"
  location = var.region

  template {
    spec {
      containers {
        image = var.image
      }
      service_account_name = google_service_account.cloud-run-backend-service.email
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  // We need this to be able to reference it in other resources
  // but it is deployed via gcloud and GCB, so we will ignore all
  // changes
  lifecycle {
    ignore_changes = all
  }
}

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

// Configure Google Provider
provider "google" {
  project = var.project
}

// Configure Google Beta Provider
provider "google-beta" {
  project = var.project
}

// Specify Required Providers
terraform {
  required_providers {
    google = {
      version = "4.42.1"
    }
    google-beta = {
      version = "4.42.1"
    }
  }
}

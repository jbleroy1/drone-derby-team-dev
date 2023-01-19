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

variable "dcu_mac_address" {

  type        = string
  description = "The Mac Address of the DCU (Drone Control Unit)."

  validation {
    condition     = can(regex("^([0-9a-f]{2}[-]){5}([0-9a-f]{2})$", var.dcu_mac_address))
    error_message = "The Mac Address must be in lower case, and hyphen seperated."
  }

}

variable "deploy_project" {

  type        = string
  description = "The Google Cloud Project to deploy the resources to."

}


variable "deploy_region" {

  type        = string
  description = "The Google Cloud Region to deploy the resources to."
  default     = "europe-west1"

}

variable "control_project" {

  type        = string
  description = "The name of the Drone Derby Control Project. You should never need to edit this. Speak with a Googler if confused."
  default     = "drone-derby-prod"

}

# Create  Cloud Build Trigger
variable "github-repo" {
  type        = string
  description = "The url of the Github repository of the backend service"
  default = ""
}

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

locals {
  cmd_channels   = ["APP-CMD", "HOST-CMD"]     // Topics that the DCU can subscribe to
  telem_channels = ["APP-TELEM", "HOST-TELEM"] // Topics that the DCU can publish to
}

// Create the PubSub topics the DCU will subscribe to
module "dcu_pubsub_telem_channels" {

  for_each = toset(local.cmd_channels)

  source = "./modules/dcu_pubsub"

  dcu_association = "subscribe"
  dcu_mac_address = var.dcu_mac_address
  control_project = var.control_project
  topic_name      = each.key

}

// Create the PubSub topics the DCU will publish to
module "dcu_comms_telem_channels" {

  for_each = toset(local.telem_channels)

  source = "./modules/dcu_pubsub"

  dcu_association = "publish"
  dcu_mac_address = var.dcu_mac_address
  control_project = var.control_project
  topic_name      = each.key

}

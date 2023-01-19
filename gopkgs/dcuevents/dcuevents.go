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

package dcuevents

// The events are expected in the PubSub Message Attributes under the
// key "type". They inform the DCU Agent which event to process
const (
	// HostControlTopic is the PubSub topic that is used for sending
	// commands to the DCU
	c = "HOST-CMD"
	// EventAttributeName is the PubSub Attribute key that is used
	// to identify the event type for a given PubSub message
	EventAttributeName = "type"
	// EventShutdownHost is the PubSub type value for a message that
	// shuts down the DCU
	EventShutdownHost = "dcu.shutdown_host"
	// EventRunContainer is the PubSub type value for a message to run
	// a Docker Container on the Coral host
	EventRunContainer = "dcu.run_container"
	// EventJoinNetwork is the PubSub type value for a message to
	// connect to a given wireless network
	EventJoinNetwork = "dcu.join_network"

	EventSendCommand="dcu.send_event"
)

// EventRunContainerPayload is the expected payload for the EventRunContainer
// event, specifying information needed to run the container on the Coral
type EventRunContainerPayload struct {
	// Image is the fully qualified path from where the container image can
	// be pulled. For example europe-west1-docker.pkg.dev/drone-derby-dev/docker-dcu/testimg:latest
	Image string `json:"image"`
	// Environment represents key value pairs to pass into the container as
	// environment variables (format "KEY=VALUE")
	Environment []string `json:"environment"`
	// RequiresTPU specifies whether or not the TPU needs to be passed to the
	// container as a hardware device
	RequireTPU bool `json:"requireTPU"`
	// RequiresServiceAccountKey specifies whether or not a Service Account Key
	// needs to be passed into the container for authentication to Google Cloud
	// Services
	RequireServiceAccountKey bool `json:"requireServiceAccountKey"`
	// ExposePorts lists is a slice of map values. Key should be the port number
	// and value should be the protocol
	ExposePorts map[string]string `json:"exposePorts"`
}

// EventJoinNetworkPayload is the expected payload for the EventJoinNetwork event,
// specifying information needed to join a specific network on the Coral
type EventJoinNetworkPayload struct {
	// SSID is the base station network name to join
	SSID string `json:"ssid"`
	// PreSharedKey is the pre-shared key for WPA(2/3) networks that use a supplied credential
	PreSharedKey string `json:"presharedkey"`
	// DHCP is a boolean flag to tell the unit to grab a IP address from the connected network
	DHCP bool `json:"dhcp"`
	// InterfaceNumber is the interface on which to initiate the connection 0 is the built in
	// Coral boards wifi which should join the network with internet access and n is the interface
	// Which should connect to the drone, if omitted defaults to the first interface that isn't the built in
	InterfaceNumber int `json:"interfacenumber"`
	// IPv4StaticAddress is ignored if DHCP is set to true, will set a static ipv4 address for the interface
	IPv4StaticAddress string `json:"ipv4staticaddress"`
	// IPv4StaticSubnet is ignored if DHCP is set to true, will set a static ipv4 subnetwork for the interface
	IPv4StaticSubnet string `json:"ipv4staticsubnet"`
}

type EventSendCommandPayload struct {
	// Action sent to the DCU. Could be takeOff, land, stop, turnLeft, turnRight
	Action string `json:"action"`

	// Angle is ignored if Action is takeOff, land or stop. Then angle represent the rotation angle of the drone if action is set to turnLeft, or TurnRight
	Angle int `json:"angle"`
}

// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"net/url"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/vic/lib/config"
	"github.com/vmware/vic/lib/config/executor"
	"github.com/vmware/vic/lib/portlayer/event"
)

var Config Configuration

// Configuration is a slice of the VCH config that is relevant to the exec part of the port layer
type Configuration struct {
	// Turn on debug logging
	DebugLevel int `vic:"0.1" scope:"read-only" key:"init/common/debug"`

	// Port Layer - exec
	// Default containerVM capacity
	ContainerVMSize config.Resources `vic:"0.1" scope:"read-only" recurse:"depth=0"`

	// Permitted datastore URLs for container storage for this virtual container host
	ContainerStores []url.URL `vic:"0.1" scope:"read-only" recurse:"depth=0"`

	// Resource pools under which all containers will be created
	ComputeResources []types.ManagedObjectReference `vic:"0.1" scope:"read-only"`
	// Resource pool is the working version of the compute resource config
	ResourcePool *object.ResourcePool
	// Parent resource will be a VirtualApp on VC
	VirtualApp *object.VirtualApp

	// Path of the ISO to use for bootstrapping containers
	BootstrapImagePath string `vic:"0.1" scope:"read-only" key:"bootstrap_image_path"`

	// Allow custom naming convention for containerVMs
	ContainerNameConvention string

	// FIXME: temporary work around for injecting network path of debug nic
	Networks     map[string]*executor.NetworkEndpoint `vic:"0.1" scope:"read-only" key:"init/networks"`
	DebugNetwork object.NetworkReference

	// For now throw the Event Manager here
	EventManager event.EventManager

	// Information about the VCH resource pool and about the real host that we want
	// tol retrieve just once.
	VCHMhz          int64
	VCHMemoryLimit  int64
	HostOS          string
	HostOSVersion   string
	HostProductName string //'VMware vCenter Server' or 'VMare ESXi'

	// Datastore URLs for image stores - the top layer is [0], the bottom layer is [len-1]
	ImageStores []url.URL `vic:"0.1" scope:"read-only" key:"image_stores"`

	// Size of scratch layer in KB
	ScratchSize int64 `vic:"0.1" scope:"read-only" key:"scratch_size"`
}

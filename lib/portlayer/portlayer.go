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

package portlayer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/vmware/vic/lib/portlayer/exec"
	"github.com/vmware/vic/lib/portlayer/network"
	"github.com/vmware/vic/lib/portlayer/storage"
	"github.com/vmware/vic/pkg/vsphere/extraconfig"
	"github.com/vmware/vic/pkg/vsphere/session"
	"golang.org/x/net/context"
)

// XXX TODO(FA) use this in the _handlers the swagger server includes.
//
// API defines the interface the REST server used by the portlayer expects the
// implementation side to export
type API interface {
	storage.ImageStorer
	storage.VolumeStorer
}

func Init(ctx context.Context, sess *session.Session) error {
	source, err := extraconfig.GuestInfoSource()
	if err != nil {
		return err
	}

	sink, err := extraconfig.GuestInfoSink()
	if err != nil {
		return err
	}

	if err := exec.Init(ctx, sess, source, sink); err != nil {
		return err
	}

	if err = network.Init(ctx, sess, source, sink); err != nil {
		return err
	}

	// Grab the storage layer config blobs from extra config
	extraconfig.Decode(source, &storage.Config)
	log.Debugf("Decoded VCH config for storage: %#v", storage.Config)

	return nil
}

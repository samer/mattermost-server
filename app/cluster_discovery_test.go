// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"testing"

	"time"

	"github.com/mattermost/mattermost-server/model"
)

func TestClusterDiscoveryService(t *testing.T) {
	a := Global()
	a.Setup()

	ds := NewClusterDiscoveryService()
	ds.Type = model.CDS_TYPE_APP
	ds.ClusterName = "ClusterA"
	ds.AutoFillHostname()

	ds.Start()
	time.Sleep(2 * time.Second)

	ds.Stop()
	time.Sleep(2 * time.Second)
}

// Copyright 2020 The Lokomotive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build packet_fluo
// +build e2e

package rook

import (
	"fmt"
	"testing"

	testutil "github.com/kinvolk/lokomotive/test/components/util"
)

func TestRookCephDeployment(t *testing.T) {
	namespace := "rook"
	client := testutil.CreateKubeClient(t)
	testCases := []struct {
		deployment string
	}{
		{"csi-cephfsplugin-provisioner"},
		{"csi-rbdplugin-provisioner"},
		{"rook-ceph-crashcollector-suraj-lk-cluster-storage-worker-0"},
		{"rook-ceph-crashcollector-suraj-lk-cluster-storage-worker-1"},
		{"rook-ceph-crashcollector-suraj-lk-cluster-storage-worker-2"},
		{"rook-ceph-mgr-a"},
		{"rook-ceph-mon-a"},
		{"rook-ceph-mon-b"},
		{"rook-ceph-mon-c"},
		{"rook-ceph-osd-0"},
		{"rook-ceph-osd-1"},
		{"rook-ceph-osd-2"},
		{"rook-ceph-osd-3"},
		{"rook-ceph-osd-4"},
		{"rook-ceph-osd-5"},
		{"rook-ceph-osd-6"},
		{"rook-ceph-osd-7"},
		{"rook-ceph-osd-8"},
		{"rook-ceph-tools"},
	}

	for _, test := range testCases {
		test := test
		t.Run(fmt.Sprintf("rook-ceph deployment:%s", test.deployment), func(t *testing.T) {
			t.Parallel()
			testutil.WaitForDeployment(t, client, namespace, test.deployment, testutil.RetryInterval, testutil.Timeout)
		})
	}
}

func TestRookCephDaemonset(t *testing.T) {
	namespace := "rook"
	client := testutil.CreateKubeClient(t)
	testCases := []struct {
		daemonset string
	}{
		{"csi-cephfsplugin"},
		{"csi-rbdplugin"},
		{"rook-discover"},
	}

	for _, test := range testCases {
		test := test
		t.Run(fmt.Sprintf("rook-ceph daemonset:%s", test.daemonset), func(t *testing.T) {
			t.Parallel()
			testutil.WaitForDaemonSet(t, client, namespace, test.daemonset, testutil.RetryInterval, testutil.Timeout)
		})
	}
}

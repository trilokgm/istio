// Copyright 2019 Istio Authors
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

package bookinfo

import (
	"testing"

	"istio.io/istio/pkg/test/framework2/core"
)

// Instance represents a deployed Bookinfo app instance in a Kubernetes cluster.
type Instance interface {
	core.Resource

	Namespace() core.Namespace

	DeployRatingsV2(ctx core.Context) error
	DeployMongoDb(ctx core.Context) error
}


// New returns a new instance of Apps
func New(ctx core.Context) (i Instance, err error) {
	err = core.UnsupportedEnvironment(ctx.Environment())

	ctx.Environment().Case(core.Kube, func() {
		i, err = newKube(ctx)
	})

	return
}

// NewOrFail returns a new instance of BookInfo or fails test
func NewOrFail(t *testing.T, ctx core.Context) Instance {
	t.Helper()

	i, err := New(ctx)
	if err != nil {
		t.Fatalf("bookinfo.NewOrFail: %v", err)
	}

	return i
}
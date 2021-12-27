/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package server

import (
	"golang.org/x/net/context"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/containerd/containerd/errdefs"
	sandboxstore "github.com/containerd/containerd/pkg/cri/store/sandbox"
	"github.com/pkg/errors"
)

func (c *criService) podSandboxStats(ctx context.Context, sandbox sandboxstore.Sandbox, stats interface{}) (*runtime.PodSandboxStats, error) {
	return nil, errors.Wrap(errdefs.ErrNotImplemented, "pod sandbox stats not implemented on windows")
}

func metricsForSandbox(sandbox sandboxstore.Sandbox) (interface{}, error) {
	return nil, errors.Wrap(errdefs.ErrNotImplemented, "metrics for sandbox not implemented on windows")
}

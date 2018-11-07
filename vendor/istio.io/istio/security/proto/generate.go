// Copyright 2018 Istio Authors.
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

//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f security/proto/ca_service.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f security/proto/workload_service.proto
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f security/proto/istioca.proto
// nolint
package istio_v1_auth

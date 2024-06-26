// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	bpt := tft.NewTFBlueprintTest(t)

	bpt.DefineVerify(func(assert *assert.Assertions) {
		bpt.DefaultVerify(assert)

		projectId := bpt.GetStringOutput("project_id")
		envVars := bpt.GetStringOutput("output_env_vars")

		op := gcloud.Runf(t, "redis instances describe test-redis --project=%s --region=us-east1", projectId)
		assert.True(op.Get("authEnabled").Bool())
		assert.Equal(op.Get("memorySizeGb").String(), "1")
		assert.Equal(op.Get("transitEncryptionMode").String(), "SERVER_AUTHENTICATION")
		assert.Contains(envVars, "REDIS_HOST")
		assert.Contains(envVars, "REDIS_PORT")
	})

	bpt.Test()
}

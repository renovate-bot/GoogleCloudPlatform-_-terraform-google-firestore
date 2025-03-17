// Copyright 2025 Google LLC
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

package firestore_resource

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

type IndexField struct {
	FieldPath string `json:"fieldPath"`
	Order     string `json:"order"`
}

func TestExampleWithCompositeIndexes(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)

	example.DefineVerify(func(assert *assert.Assertions) {
		example.DefaultVerify(assert)

		projectId := example.GetTFSetupStringOutput("project_id")
		databaseId := example.GetStringOutput("database_id")

		databaseIdPrefix := fmt.Sprintf("projects/%s/databases/", projectId)
		databaseName := strings.TrimPrefix(databaseId, databaseIdPrefix)

		databaseInfo := gcloud.Run(
			t,
			"firestore databases describe",
			gcloud.WithCommonArgs([]string{"--project", projectId, "--database", databaseName, "--format", "json"}),
		)

		indexInfo := gcloud.Run(
			t,
			"firestore indexes composite list",
			gcloud.WithCommonArgs([]string{"--project", projectId, "--database", databaseName, "--format", "json"}),
		).Array()

		assert.Equal(
			databaseId,
			databaseInfo.Get("name").String(),
			"Database ID does not match.",
		)

		assert.Equal(
			1,
			len(indexInfo),
		)

		assert.Equal(
			"COLLECTION",
			indexInfo[0].Get("queryScope").String(),
		)

		var indexFieldList []IndexField
		json.Unmarshal([]byte(indexInfo[0].Get("fields").String()), &indexFieldList)

		//field1 - ASC, field2 - DESC, __name__ - DESC (auto-added)
		assert.Equal(
			3,
			len(indexFieldList),
		)

		assert.Equal(
			"field1",
			indexFieldList[0].FieldPath,
		)

		assert.Equal(
			"ASCENDING",
			indexFieldList[0].Order,
		)

		assert.Equal(
			"field2",
			indexFieldList[1].FieldPath,
		)

		assert.Equal(
			"DESCENDING",
			indexFieldList[1].Order,
		)

		assert.Equal(
			"__name__",
			indexFieldList[2].FieldPath,
		)

		assert.Equal(
			"DESCENDING",
			indexFieldList[2].Order,
		)
	})
	example.Test()
}

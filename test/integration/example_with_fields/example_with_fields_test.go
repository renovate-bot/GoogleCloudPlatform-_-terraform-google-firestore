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
	"regexp"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

type IndexConfig struct {
	AncestorField string `json:"ancestorField"`
	Indexes       []IndexInfo
}

type IndexInfo struct {
	Fields     []IndexField `json:"fields"`
	QueryScope string       `json:"queryScope"`
	State      string       `json:"state"`
}

type IndexField struct {
	FieldPath string `json:"fieldPath"`
	Order     string `json:"order"`
}

func extractCollectionGroupAndFieldName(fieldId string) (string, string) {
	re := regexp.MustCompile(`collectionGroups/([^/]+)/fields/([^/]+)`)
	match := re.FindStringSubmatch(fieldId)

	if len(match) == 3 {
		return match[1], match[2]
	}

	return "", ""
}

func TestExampleWithFields(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)

	example.DefineVerify(func(assert *assert.Assertions) {
		example.DefaultVerify(assert)

		projectId := example.GetTFSetupStringOutput("project_id")
		databaseId := example.GetStringOutput("database_id")
		fieldIds := example.GetJsonOutput("field_ids").Array()

		databaseIdPrefix := fmt.Sprintf("projects/%s/databases/", projectId)
		databaseName := strings.TrimPrefix(databaseId, databaseIdPrefix)
		collectionGroupName, fieldName := extractCollectionGroupAndFieldName(fieldIds[0].String())

		assert.Equal("collection-1", collectionGroupName, "Collection group name does not match in TF output.")
		assert.Equal("field1", fieldName, "Field name does not match in TF output.")

		databaseInfo := gcloud.Run(
			t,
			"firestore databases describe",
			gcloud.WithCommonArgs([]string{"--project", projectId, "--database", databaseName, "--format", "json"}),
		)

		fieldInfo := gcloud.Run(
			t,
			fmt.Sprintf("firestore indexes fields describe %s", fieldName),
			gcloud.WithCommonArgs([]string{"--project", projectId, "--database", databaseName, "--collection-group", collectionGroupName, "--format", "json"}),
		)

		var indexConfig IndexConfig
		json.Unmarshal([]byte(fieldInfo.Get("indexConfig").String()), &indexConfig)

		assert.Equal(
			databaseId,
			databaseInfo.Get("name").String(),
			"Database ID does not match.",
		)

		assert.Equal(
			4,
			len(indexConfig.Indexes),
			"Fields were not created properly - missing indexes for field.",
		)

		var collectionScopedIndexes []IndexInfo
		var collectionGroupScopedIndexes []IndexInfo

		for _, index := range indexConfig.Indexes {
			if index.QueryScope == "COLLECTION" {
				collectionScopedIndexes = append(collectionScopedIndexes, index)
			} else if index.QueryScope == "COLLECTION_GROUP" {
				collectionGroupScopedIndexes = append(collectionGroupScopedIndexes, index)
			}
		}

		assert.Equal(
			2,
			len(collectionScopedIndexes),
			"Fields were not created properly - collection scoped indexes missing.",
		)

		assert.Equal(
			2,
			len(collectionGroupScopedIndexes),
			"Fields were not created properly - collection group scoped indexes missing.",
		)

		assert.True(
			collectionGroupScopedIndexes[0].Fields[0].Order == "ASCENDING" || collectionGroupScopedIndexes[1].Fields[0].Order == "ASCENDING",
			"Fields were not created properly - Ascending index missing in collection group scope.",
		)

		assert.True(
			collectionGroupScopedIndexes[0].Fields[0].Order == "DESCENDING" || collectionGroupScopedIndexes[1].Fields[0].Order == "DESCENDING",
			"Fields were not created properly - Descending index missing in collection group scope.",
		)

		assert.True(
			collectionScopedIndexes[0].Fields[0].Order == "ASCENDING" || collectionScopedIndexes[1].Fields[0].Order == "ASCENDING",
			"Fields were not created properly - Ascending index missing in collection scope.",
		)

		assert.True(
			collectionScopedIndexes[0].Fields[0].Order == "DESCENDING" || collectionScopedIndexes[1].Fields[0].Order == "DESCENDING",
			"Fields were not created properly - Descending index missing in collection scope.",
		)
	})
	example.Test()
}

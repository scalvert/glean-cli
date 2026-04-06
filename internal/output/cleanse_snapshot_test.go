package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanseSearchResponse_Snapshots(t *testing.T) {
	fixtures := []struct {
		name     string
		rawFile  string
		wantFile string
	}{
		{"people query", "raw_people.json", "cleansed_people.json"},
		{"jira query", "raw_jira.json", "cleansed_jira.json"},
		{"github query", "raw_github.json", "cleansed_github.json"},
		{"mixed query", "raw_mixed.json", "cleansed_mixed.json"},
	}

	for _, tt := range fixtures {
		t.Run(tt.name, func(t *testing.T) {
			rawBytes, err := os.ReadFile(filepath.Join("testdata", tt.rawFile))
			require.NoError(t, err)

			wantBytes, err := os.ReadFile(filepath.Join("testdata", tt.wantFile))
			require.NoError(t, err)

			var rawInput map[string]any
			require.NoError(t, json.Unmarshal(rawBytes, &rawInput))

			got, err := CleanseSearchResponse(rawInput)
			require.NoError(t, err)

			gotBytes, err := json.MarshalIndent(got, "", "  ")
			require.NoError(t, err)

			var want any
			require.NoError(t, json.Unmarshal(wantBytes, &want))
			wantNorm, err := json.MarshalIndent(want, "", "  ")
			require.NoError(t, err)

			if !assert.JSONEq(t, string(wantNorm), string(gotBytes)) {
				// On failure, write the actual output for easy diffing
				actualPath := filepath.Join("testdata", "actual_"+tt.wantFile)
				_ = os.WriteFile(actualPath, append(gotBytes, '\n'), 0600)
				t.Logf("Actual output written to %s for diffing", actualPath)
			}
		})
	}
}

func TestCleanseSearchResponse_SnapshotStructure(t *testing.T) {
	allowedResponse := map[string]bool{"results": true, "cursor": true, "hasMoreResults": true, "requestID": true}
	allowedResult := map[string]bool{"title": true, "url": true, "snippets": true, "document": true}
	allowedDocument := map[string]bool{"title": true, "url": true, "datasource": true, "docType": true, "metadata": true}
	allowedMetadata := map[string]bool{
		"datasource": true, "objectType": true, "author": true, "owner": true,
		"assignedTo": true, "updatedBy": true, "updateTime": true, "createTime": true,
		"status": true, "priority": true, "container": true, "datasourceId": true,
	}
	allowedPerson := map[string]bool{"name": true}
	allowedSnippet := map[string]bool{"text": true, "mimeType": true}

	fixtures := []string{"raw_people.json", "raw_jira.json", "raw_github.json", "raw_mixed.json"}

	for _, fixture := range fixtures {
		t.Run(fixture, func(t *testing.T) {
			rawBytes, err := os.ReadFile(filepath.Join("testdata", fixture))
			require.NoError(t, err)

			var rawInput map[string]any
			require.NoError(t, json.Unmarshal(rawBytes, &rawInput))

			got, err := CleanseSearchResponse(rawInput)
			require.NoError(t, err)

			m := got.(map[string]any)

			// Response level
			for k := range m {
				assert.True(t, allowedResponse[k], "disallowed response key: %s", k)
			}

			results, _ := m["results"].([]any)
			for i, r := range results {
				rm := r.(map[string]any)

				// Result level
				for k := range rm {
					assert.True(t, allowedResult[k], "result[%d] disallowed key: %s", i, k)
				}

				// Document level
				if doc, ok := rm["document"].(map[string]any); ok {
					for k := range doc {
						assert.True(t, allowedDocument[k], "result[%d].document disallowed key: %s", i, k)
					}

					// Metadata level
					if meta, ok := doc["metadata"].(map[string]any); ok {
						for k := range meta {
							assert.True(t, allowedMetadata[k], "result[%d].metadata disallowed key: %s", i, k)
						}

						// Person fields
						for _, pf := range []string{"author", "owner", "assignedTo", "updatedBy"} {
							if p, ok := meta[pf].(map[string]any); ok {
								for k := range p {
									assert.True(t, allowedPerson[k], "result[%d].metadata.%s disallowed key: %s", i, pf, k)
								}
							}
						}
					}
				}

				// Snippet level
				if snippets, ok := rm["snippets"].([]any); ok {
					for j, s := range snippets {
						sm := s.(map[string]any)
						for k := range sm {
							assert.True(t, allowedSnippet[k], "result[%d].snippets[%d] disallowed key: %s", i, j, k)
						}
						text, _ := sm["text"].(string)
						assert.NotEmpty(t, text, "result[%d].snippets[%d] has empty text", i, j)
					}
				}
			}
		})
	}
}

package output

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanseSearchResponse_StripsUIFields(t *testing.T) {
	input := map[string]any{
		"trackingToken":     "tok_abc",
		"sessionInfo":       map[string]any{"sessionId": "sess123"},
		"experimentIds":     []any{float64(1), float64(2)},
		"backendTimeMillis": float64(42),
		"structuredResults": []any{map[string]any{"foo": "bar"}},
		"generatedQnaResult": map[string]any{
			"question": "what?",
		},
		"metadata":              map[string]any{"source": "internal"},
		"facetResults":          []any{},
		"resultTabs":            []any{},
		"resultTabIds":          []any{"tab1"},
		"resultsDescription":    map[string]any{"text": "showing results"},
		"rewrittenFacetFilters": []any{},
		"requestID":             "req_123",
		"cursor":                "eyJwIjo0",
		"hasMoreResults":        true,
		"results": []any{
			map[string]any{
				"trackingToken":       "result_tok",
				"structuredResults":   []any{},
				"clusteredResults":    []any{},
				"backlinkResults":     []any{},
				"prominence":          "HIGH",
				"querySuggestion":     map[string]any{},
				"clusterType":         "THREAD",
				"attachmentCount":     float64(3),
				"attachments":         []any{},
				"pins":                []any{},
				"nativeAppUrl":        "slack://foo",
				"fullText":            "some full text",
				"fullTextList":        []any{"line1"},
				"relatedResults":      []any{},
				"allClusteredResults": []any{},
				"title":               "Q2 Planning",
				"url":                 "https://example.com/doc",
				"snippets": []any{
					map[string]any{
						"text":                "The platform team will focus on...",
						"snippet":             "deprecated snippet value",
						"mimeType":            "text/plain",
						"snippetTextOrdering": float64(1),
						"ranges":              []any{map[string]any{"start": float64(0), "end": float64(5)}},
					},
					map[string]any{
						"text":     "",
						"mimeType": "text/plain",
					},
				},
				"document": map[string]any{
					"id":                "doc_123",
					"title":             "Q2 Planning Doc",
					"url":               "https://example.com/doc",
					"datasource":        "confluence",
					"docType":           "page",
					"connectorType":     "API_CRAWL",
					"content":           map[string]any{"fullTextList": []any{"text"}},
					"containerDocument": map[string]any{"id": "parent"},
					"parentDocument":    map[string]any{"id": "parent2"},
					"sections":          []any{map[string]any{"title": "Intro"}},
					"metadata": map[string]any{
						"datasource": "confluence",
						"objectType": "page",
						"author": map[string]any{
							"name":         "Steve",
							"obfuscatedId": "ABC123",
							"metadata":     map[string]any{"loggingId": "XYZ"},
						},
						"owner": map[string]any{
							"name":         "Jane",
							"obfuscatedId": "DEF456",
							"metadata":     map[string]any{"loggingId": "UVW"},
						},
						"assignedTo": map[string]any{
							"name":         "Bob",
							"obfuscatedId": "GHI789",
						},
						"updatedBy": map[string]any{
							"name":         "Alice",
							"obfuscatedId": "JKL012",
						},
						"updateTime":       "2026-03-28T14:30:00Z",
						"createTime":       "2026-01-15T09:00:00Z",
						"status":           "In Progress",
						"priority":         "P1",
						"container":        "Engineering Space",
						"datasourceId":     "JIRA-123",
						"interactions":     map[string]any{"numViews": float64(42)},
						"documentCategory": "PUBLISHED_CONTENT",
						"loggingId":        "log_abc",
						"documentId":       "did_123",
						"visibility":       map[string]any{"level": "PUBLIC"},
						"pins":             []any{},
						"collections":      []any{},
					},
				},
			},
		},
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m, ok := result.(map[string]any)
	require.True(t, ok)

	// Top-level: only allowed keys
	assert.Contains(t, m, "results")
	assert.Contains(t, m, "cursor")
	assert.Contains(t, m, "hasMoreResults")
	assert.Contains(t, m, "requestID")
	assert.NotContains(t, m, "trackingToken")
	assert.NotContains(t, m, "sessionInfo")
	assert.NotContains(t, m, "experimentIds")
	assert.NotContains(t, m, "backendTimeMillis")
	assert.NotContains(t, m, "structuredResults")
	assert.NotContains(t, m, "generatedQnaResult")
	assert.NotContains(t, m, "metadata")
	assert.NotContains(t, m, "facetResults")
	assert.NotContains(t, m, "resultTabs")
	assert.NotContains(t, m, "resultTabIds")
	assert.NotContains(t, m, "resultsDescription")
	assert.NotContains(t, m, "rewrittenFacetFilters")

	// Results array
	results, ok := m["results"].([]any)
	require.True(t, ok)
	require.Len(t, results, 1)

	r := results[0].(map[string]any)
	assert.Contains(t, r, "title")
	assert.Contains(t, r, "url")
	assert.Contains(t, r, "snippets")
	assert.Contains(t, r, "document")
	assert.NotContains(t, r, "trackingToken")
	assert.NotContains(t, r, "structuredResults")
	assert.NotContains(t, r, "clusteredResults")
	assert.NotContains(t, r, "backlinkResults")
	assert.NotContains(t, r, "prominence")
	assert.NotContains(t, r, "pins")
	assert.NotContains(t, r, "nativeAppUrl")
	assert.NotContains(t, r, "fullText")
	assert.NotContains(t, r, "attachments")
	assert.NotContains(t, r, "clusterType")

	// Snippets: only text and mimeType, empty snippets filtered out
	snippets := r["snippets"].([]any)
	require.Len(t, snippets, 1, "empty snippet should be filtered out")
	snip := snippets[0].(map[string]any)
	assert.Contains(t, snip, "text")
	assert.Contains(t, snip, "mimeType")
	assert.NotContains(t, snip, "snippet", "deprecated snippet field should be stripped")
	assert.NotContains(t, snip, "ranges")
	assert.NotContains(t, snip, "snippetTextOrdering")

	// Document: only allowed keys
	doc := r["document"].(map[string]any)
	assert.Contains(t, doc, "title")
	assert.Contains(t, doc, "url")
	assert.Contains(t, doc, "datasource")
	assert.Contains(t, doc, "docType")
	assert.Contains(t, doc, "metadata")
	assert.NotContains(t, doc, "id")
	assert.NotContains(t, doc, "connectorType")
	assert.NotContains(t, doc, "content")
	assert.NotContains(t, doc, "containerDocument")
	assert.NotContains(t, doc, "parentDocument")
	assert.NotContains(t, doc, "sections")

	// Document metadata: allowed keys kept, noise stripped
	meta := doc["metadata"].(map[string]any)
	assert.Contains(t, meta, "datasource")
	assert.Contains(t, meta, "objectType")
	assert.Contains(t, meta, "author")
	assert.Contains(t, meta, "owner")
	assert.Contains(t, meta, "assignedTo")
	assert.Contains(t, meta, "updatedBy")
	assert.Contains(t, meta, "updateTime")
	assert.Contains(t, meta, "createTime")
	assert.Contains(t, meta, "status")
	assert.Contains(t, meta, "priority")
	assert.Contains(t, meta, "container")
	assert.Contains(t, meta, "datasourceId")
	assert.NotContains(t, meta, "interactions")
	assert.NotContains(t, meta, "documentCategory")
	assert.NotContains(t, meta, "loggingId")
	assert.NotContains(t, meta, "documentId")
	assert.NotContains(t, meta, "visibility")
	assert.NotContains(t, meta, "pins")
	assert.NotContains(t, meta, "collections")

	// Author filtered to name only (email doesn't exist in API responses)
	author := meta["author"].(map[string]any)
	assert.Equal(t, "Steve", author["name"])
	assert.NotContains(t, author, "obfuscatedId")
	assert.NotContains(t, author, "metadata")

	// Other person fields also filtered to name only
	owner := meta["owner"].(map[string]any)
	assert.Equal(t, "Jane", owner["name"])
	assert.NotContains(t, owner, "obfuscatedId")
	assert.NotContains(t, owner, "metadata")

	assignedTo := meta["assignedTo"].(map[string]any)
	assert.Equal(t, "Bob", assignedTo["name"])
	assert.NotContains(t, assignedTo, "obfuscatedId")

	updatedBy := meta["updatedBy"].(map[string]any)
	assert.Equal(t, "Alice", updatedBy["name"])
	assert.NotContains(t, updatedBy, "obfuscatedId")
}

func TestCleanseSearchResponse_EmptyResults(t *testing.T) {
	input := map[string]any{
		"results":        []any{},
		"hasMoreResults": false,
		"requestID":      "req_456",
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m := result.(map[string]any)
	assert.Equal(t, []any{}, m["results"])
	assert.Equal(t, false, m["hasMoreResults"])
	assert.Equal(t, "req_456", m["requestID"])
}

func TestCleanseSearchResponse_FiltersEmptyResults(t *testing.T) {
	input := map[string]any{
		"results": []any{
			// Structured result with no document or title — should be removed
			map[string]any{
				"url":               "",
				"structuredResults": []any{map[string]any{"foo": "bar"}},
				"trackingToken":     "tok",
			},
			// Real result with document — should be kept
			map[string]any{
				"url":   "https://example.com/doc",
				"title": "Real Doc",
				"document": map[string]any{
					"title":      "Real Doc",
					"datasource": "confluence",
				},
			},
			// Result with title but no document — should be kept
			map[string]any{
				"url":   "https://example.com/other",
				"title": "Has Title",
			},
			// Empty result, url present but no title/document — should be removed
			map[string]any{
				"url": "https://example.com/empty",
			},
		},
		"requestID": "req_filter",
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m := result.(map[string]any)
	results := m["results"].([]any)
	require.Len(t, results, 2, "should keep only results with document or non-empty title")

	r0 := results[0].(map[string]any)
	assert.Equal(t, "Real Doc", r0["title"])

	r1 := results[1].(map[string]any)
	assert.Equal(t, "Has Title", r1["title"])
}

func TestCleanseSearchResponse_MissingOptionalFields(t *testing.T) {
	input := map[string]any{
		"results": []any{
			map[string]any{
				"url":   "https://example.com",
				"title": "Test",
			},
		},
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m := result.(map[string]any)
	results := m["results"].([]any)
	r := results[0].(map[string]any)
	assert.Equal(t, "Test", r["title"])
	assert.Equal(t, "https://example.com", r["url"])
	assert.NotContains(t, r, "document")
	assert.NotContains(t, r, "snippets")
}

func TestCleanseSearchResponse_SDKStruct(t *testing.T) {
	// Simulate passing an SDK struct (not a map) — verifies the marshal round-trip works.
	type fakeResult struct {
		Title         string `json:"title"`
		URL           string `json:"url"`
		TrackingToken string `json:"trackingToken"`
	}
	type fakeResponse struct {
		Results       []fakeResult `json:"results"`
		TrackingToken string       `json:"trackingToken"`
		RequestID     string       `json:"requestID"`
	}

	resp := fakeResponse{
		Results: []fakeResult{
			{Title: "Doc", URL: "https://x.com", TrackingToken: "tok"},
		},
		TrackingToken: "resp_tok",
		RequestID:     "req_789",
	}

	result, err := CleanseSearchResponse(resp)
	require.NoError(t, err)

	// Should be a map now
	m, ok := result.(map[string]any)
	require.True(t, ok)
	assert.Contains(t, m, "requestID")
	assert.NotContains(t, m, "trackingToken")

	results := m["results"].([]any)
	r := results[0].(map[string]any)
	assert.Equal(t, "Doc", r["title"])
	assert.NotContains(t, r, "trackingToken")
}

func TestWarnStrippedFields_AllowedFields(t *testing.T) {
	stripped := WarnStrippedFields("results.title,results.url,results.document.datasource,results.document.metadata.status,results.document.metadata.owner.name")
	assert.Empty(t, stripped)
}

func TestWarnStrippedFields_StrippedFields(t *testing.T) {
	stripped := WarnStrippedFields("results.trackingToken,results.title,results.structuredResults")
	assert.Equal(t, []string{"results.trackingToken", "results.structuredResults"}, stripped)
}

func TestWarnStrippedFields_DeepAllowedPath(t *testing.T) {
	stripped := WarnStrippedFields("results.document.metadata.author.name")
	assert.Empty(t, stripped)
}

func TestWarnStrippedFields_DeepStrippedPath(t *testing.T) {
	stripped := WarnStrippedFields("results.document.content")
	assert.Equal(t, []string{"results.document.content"}, stripped)
}

func TestWarnStrippedFields_TopLevelStripped(t *testing.T) {
	stripped := WarnStrippedFields("trackingToken,sessionInfo,requestID")
	assert.Equal(t, []string{"trackingToken", "sessionInfo"}, stripped)
}

func TestWarnStrippedFields_Empty(t *testing.T) {
	stripped := WarnStrippedFields("")
	assert.Nil(t, stripped)
}

func TestCleanseSearchResponse_FiltersEmptySnippets(t *testing.T) {
	input := map[string]any{
		"results": []any{
			map[string]any{
				"title": "Doc With Snippets",
				"url":   "https://example.com",
				"document": map[string]any{
					"title":      "Doc",
					"datasource": "gdrive",
				},
				"snippets": []any{
					map[string]any{"text": "real content", "mimeType": "text/plain"},
					map[string]any{"text": "", "mimeType": "text/plain"},
					map[string]any{"mimeType": "text/plain"},
					map[string]any{"text": "another real one", "mimeType": "text/html"},
				},
			},
		},
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m := result.(map[string]any)
	results := m["results"].([]any)
	r := results[0].(map[string]any)
	snippets := r["snippets"].([]any)
	require.Len(t, snippets, 2, "should keep only snippets with non-empty text")
	assert.Equal(t, "real content", snippets[0].(map[string]any)["text"])
	assert.Equal(t, "another real one", snippets[1].(map[string]any)["text"])
}

func TestCleanseSearchResponse_AllSnippetsEmptyRemovesKey(t *testing.T) {
	input := map[string]any{
		"results": []any{
			map[string]any{
				"title": "Doc",
				"url":   "https://example.com",
				"document": map[string]any{
					"title":      "Doc",
					"datasource": "gdrive",
				},
				"snippets": []any{
					map[string]any{"text": "", "mimeType": "text/plain"},
					map[string]any{"text": "", "mimeType": "text/plain"},
				},
			},
		},
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	m := result.(map[string]any)
	r := m["results"].([]any)[0].(map[string]any)
	assert.NotContains(t, r, "snippets", "snippets key should be removed when all snippets are empty")
}

func TestCleanseSearchResponse_RoundTripsToJSON(t *testing.T) {
	input := map[string]any{
		"results": []any{
			map[string]any{
				"title": "Hello",
				"url":   "https://example.com",
			},
		},
		"requestID": "req_1",
	}

	result, err := CleanseSearchResponse(input)
	require.NoError(t, err)

	// Must be serializable
	data, err := json.Marshal(result)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"title":"Hello"`)
	assert.Contains(t, string(data), `"requestID":"req_1"`)
}

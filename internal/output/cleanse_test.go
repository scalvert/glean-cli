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
						"snippet":  "The platform team will focus on...",
						"mimeType": "text/plain",
						"ranges":   []any{map[string]any{"start": float64(0), "end": float64(5)}},
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
							"email":        "steve@co.com",
							"obfuscatedId": "ABC123",
							"metadata":     map[string]any{"loggingId": "XYZ"},
						},
						"updateTime":       "2026-03-28T14:30:00Z",
						"createTime":       "2026-01-15T09:00:00Z",
						"container":        "Engineering Space",
						"interactions":     map[string]any{"numViews": float64(42)},
						"documentCategory": "PUBLISHED_CONTENT",
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

	// Snippets: only snippet and mimeType
	snippets := r["snippets"].([]any)
	require.Len(t, snippets, 1)
	snip := snippets[0].(map[string]any)
	assert.Contains(t, snip, "snippet")
	assert.Contains(t, snip, "mimeType")
	assert.NotContains(t, snip, "ranges")

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

	// Document metadata: only allowed keys
	meta := doc["metadata"].(map[string]any)
	assert.Contains(t, meta, "datasource")
	assert.Contains(t, meta, "objectType")
	assert.Contains(t, meta, "author")
	assert.Contains(t, meta, "updateTime")
	assert.Contains(t, meta, "createTime")
	assert.NotContains(t, meta, "container")
	assert.NotContains(t, meta, "interactions")
	assert.NotContains(t, meta, "documentCategory")
	assert.NotContains(t, meta, "pins")
	assert.NotContains(t, meta, "collections")

	// Author filtered to name and email only
	author := meta["author"].(map[string]any)
	assert.Equal(t, "Steve", author["name"])
	assert.Equal(t, "steve@co.com", author["email"])
	assert.NotContains(t, author, "metadata")
	assert.NotContains(t, author, "obfuscatedId")
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
	stripped := WarnStrippedFields("results.title,results.url,results.document.datasource")
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

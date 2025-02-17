// Package api provides types and utilities for interacting with the Glean API.
package api

// Document represents a document in the search or chat results
type Document struct {
	ParentDocument *Document         `json:"parentDocument,omitempty"`
	Metadata       *DocumentMetadata `json:"metadata,omitempty"`
	ID             string            `json:"id"`
	Datasource     string            `json:"datasource"`
	DocType        string            `json:"docType"`
	Title          string            `json:"title"`
	URL            string            `json:"url"`
}

// DocumentMetadata contains additional information about a document
type DocumentMetadata struct {
	Datasource         string                 `json:"datasource"`
	DatasourceInstance string                 `json:"datasourceInstance"`
	ObjectType         string                 `json:"objectType"`
	Container          string                 `json:"container,omitempty"`
	ContainerId        string                 `json:"containerId,omitempty"`
	MimeType           string                 `json:"mimeType"`
	DocumentId         string                 `json:"documentId"`
	LoggingId          string                 `json:"loggingId"`
	CreateTime         string                 `json:"createTime"`
	UpdateTime         string                 `json:"updateTime"`
	Author             *Person                `json:"author,omitempty"`
	Owner              *Person                `json:"owner,omitempty"`
	Visibility         string                 `json:"visibility"`
	Status             string                 `json:"status,omitempty"`
	AssignedTo         *Person                `json:"assignedTo,omitempty"`
	DatasourceId       string                 `json:"datasourceId"`
	Interactions       map[string]interface{} `json:"interactions"`
	DocumentCategory   string                 `json:"documentCategory"`
	CustomData         map[string]interface{} `json:"customData,omitempty"`
	Shortcuts          []Shortcut             `json:"shortcuts,omitempty"`
}

// Person represents a user in the Glean system
type Person struct {
	Metadata     *PersonMetadata `json:"metadata,omitempty"`
	Name         string          `json:"name"`
	ObfuscatedId string          `json:"obfuscatedId"`
}

// PersonMetadata contains additional information about a person
type PersonMetadata struct {
	RelatedDocuments []RelatedDocument `json:"relatedDocuments,omitempty"`
}

// RelatedDocument represents a document related to a person
type RelatedDocument struct {
	// Add fields as needed
}

// Shortcut represents a go link or other shortcut
type Shortcut struct {
	InputAlias     string `json:"inputAlias"`
	DestinationUrl string `json:"destinationUrl"`
	Description    string `json:"description"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
	ViewPrefix     string `json:"viewPrefix"`
	Alias          string `json:"alias"`
	Title          string `json:"title"`
}

// StructuredResult represents a document result in search or chat responses
type StructuredResult struct {
	Document *Document      `json:"document,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Type     string         `json:"type"`
}

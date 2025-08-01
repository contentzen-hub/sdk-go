package contentzen

// Document represents a ContentZen document.
type Document struct {
	UUID    string                 `json:"uuid"`
	Payload map[string]interface{} `json:"payload"`
	Lang    string                 `json:"lang"`
	State   string                 `json:"state"`
}

// Collection represents a ContentZen collection.
type Collection struct {
	UUID        string                 `json:"uuid"`
	Name        string                 `json:"name"`
	DisplayName string                 `json:"display_name"`
	Description string                 `json:"description"`
	IsPublic    bool                   `json:"is_public"`
	Fields      []CollectionField      `json:"fields"`
}

// CollectionField represents a field in a collection schema.
type CollectionField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
	Required    bool   `json:"required"`
	Unique      bool   `json:"unique,omitempty"`
}

// Media represents a media file in ContentZen.
type Media struct {
	UUID    string `json:"uuid"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

// Webhook represents a webhook in ContentZen.
type Webhook struct {
	UUID   string   `json:"uuid"`
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Events []string `json:"events"`
	Method string   `json:"method"`
}

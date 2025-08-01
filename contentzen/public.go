package contentzen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPublicDocuments fetches all published documents from a public collection.
func (c *Client) GetPublicDocuments(collectionUUID string) ([]Document, error) {
	url := fmt.Sprintf("%s/api/v1/documents/collection/%s?state=published", c.BaseURL, collectionUUID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	// The API returns {"data": [...], ...}
	var result struct {
		Data []Document `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetPublicDocument fetches a specific published document from a public collection.
func (c *Client) GetPublicDocument(collectionUUID, documentUUID string) (*Document, error) {
	url := fmt.Sprintf("%s/api/v1/documents/collection/%s/%s", c.BaseURL, collectionUUID, documentUUID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var doc Document
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

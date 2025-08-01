package contentzen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// GetDocuments fetches documents from a collection (requires API token).
func (c *Client) GetDocuments(collectionUUID string) ([]Document, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/documents/%s", c.BaseURL, collectionUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var docs []Document
	if err := json.NewDecoder(resp.Body).Decode(&docs); err != nil {
		return nil, err
	}
	return docs, nil
}

// CreateDocument creates a new document in a collection (requires API token).
func (c *Client) CreateDocument(collectionUUID string, doc *Document) (*Document, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/documents/%s", c.BaseURL, collectionUUID)
	body, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var created Document
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

// UpdateDocument updates an existing document (requires API token).
func (c *Client) UpdateDocument(collectionUUID, documentUUID string, doc *Document) (*Document, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/documents/%s/%s", c.BaseURL, collectionUUID, documentUUID)
	body, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var updated Document
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteDocument deletes a document (requires API token).
func (c *Client) DeleteDocument(collectionUUID, documentUUID string) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/documents/%s/%s", c.BaseURL, collectionUUID, documentUUID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// GetCollections fetches all collections for the authenticated project.
func (c *Client) GetCollections() ([]Collection, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var collections []Collection
	if err := json.NewDecoder(resp.Body).Decode(&collections); err != nil {
		return nil, err
	}
	return collections, nil
}

// GetCollection fetches a specific collection by UUID.
func (c *Client) GetCollection(collectionUUID string) (*Collection, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/%s", c.BaseURL, collectionUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var collection Collection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// CreateCollection creates a new collection.
func (c *Client) CreateCollection(col *Collection) (*Collection, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections", c.BaseURL)
	body, err := json.Marshal(col)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var created Collection
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

// UpdateCollection updates an existing collection.
func (c *Client) UpdateCollection(collectionUUID string, col *Collection) (*Collection, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/%s", c.BaseURL, collectionUUID)
	body, err := json.Marshal(col)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var updated Collection
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteCollection deletes a collection by UUID.
func (c *Client) DeleteCollection(collectionUUID string) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/%s", c.BaseURL, collectionUUID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// ListMedia fetches all media files for the authenticated project.
func (c *Client) ListMedia() ([]Media, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/media/ls", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var media []Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, err
	}
	return media, nil
}

// UploadMedia uploads a media file. Accepts a file path.
func (c *Client) UploadMedia(filePath string) (*Media, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return nil, err
	}
	w.Close()

	url := fmt.Sprintf("%s/api/v1/media/upload", c.BaseURL)
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var media Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, err
	}
	return &media, nil
}

// GetMedia fetches a specific media file by UUID.
func (c *Client) GetMedia(mediaUUID string) (*Media, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/media/%s", c.BaseURL, mediaUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var media Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, err
	}
	return &media, nil
}

// UpdateMedia updates metadata for a media file.
func (c *Client) UpdateMedia(mediaUUID string, media *Media) (*Media, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/media/%s", c.BaseURL, mediaUUID)
	body, err := json.Marshal(media)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var updated Media
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteMedia deletes a media file by UUID.
func (c *Client) DeleteMedia(mediaUUID string) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/media/%s", c.BaseURL, mediaUUID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// DownloadMedia downloads a media file by UUID to the specified destination path.
func (c *Client) DownloadMedia(mediaUUID, destPath string) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/media/%s/download", c.BaseURL, mediaUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// ListWebhooks fetches all webhooks for the authenticated project.
func (c *Client) ListWebhooks() ([]Webhook, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/webhooks", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var webhooks []Webhook
	if err := json.NewDecoder(resp.Body).Decode(&webhooks); err != nil {
		return nil, err
	}
	return webhooks, nil
}

// CreateWebhook creates a new webhook.
func (c *Client) CreateWebhook(wh *Webhook) (*Webhook, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/webhooks", c.BaseURL)
	body, err := json.Marshal(wh)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var created Webhook
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

// UpdateWebhook updates an existing webhook by UUID.
func (c *Client) UpdateWebhook(webhookUUID string, wh *Webhook) (*Webhook, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/webhooks/%s", c.BaseURL, webhookUUID)
	body, err := json.Marshal(wh)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var updated Webhook
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteWebhook deletes a webhook by UUID.
func (c *Client) DeleteWebhook(webhookUUID string) error {
	if c.APIToken == "" {
		return fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/webhooks/%s", c.BaseURL, webhookUUID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	return nil
}

// GetCollectionSchema fetches the schema for a collection.
func (c *Client) GetCollectionSchema(collectionUUID string) (map[string]interface{}, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/%s/schema", c.BaseURL, collectionUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var schema map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&schema); err != nil {
		return nil, err
	}
	return schema, nil
}

// GetCollectionFields fetches the fields for a collection.
func (c *Client) GetCollectionFields(collectionUUID string) ([]CollectionField, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/%s/fields", c.BaseURL, collectionUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var fields []CollectionField
	if err := json.NewDecoder(resp.Body).Decode(&fields); err != nil {
		return nil, err
	}
	return fields, nil
}

// GetFieldTypes fetches available field types for collections.
func (c *Client) GetFieldTypes() ([]string, error) {
	if c.APIToken == "" {
		return nil, fmt.Errorf("API token required for this endpoint")
	}
	url := fmt.Sprintf("%s/api/v1/collections/field-types", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var types []string
	if err := json.NewDecoder(resp.Body).Decode(&types); err != nil {
		return nil, err
	}
	return types, nil
}

# ContentZen Go SDK

Go SDK for ContentZen CMS. Built for speed, concurrency, and seamless integration into microservices and backend systems.

## Features
- Access public and private ContentZen API endpoints
- API token authentication for private endpoints
- Manage documents, collections, media, and webhooks
- Designed for high performance and concurrency

## Installation

```
go get github.com/yourusername/sdk-go
```

## Authentication
Some endpoints require an API token. You can obtain your API token after registering at [contentzen.io](https://contentzen.io).

## Usage Examples

### Client Initialization

```go
import "github.com/yourusername/sdk-go/contentzen"

// For public endpoints only:
publicClient := contentzen.NewClient("")

// For authenticated endpoints:
authClient := contentzen.NewClient("<your-api-token>")
```

### Documents

```go
// Get public documents
publicDocs, err := publicClient.GetPublicDocuments("collection-uuid")

// Get a public document
publicDoc, err := publicClient.GetPublicDocument("collection-uuid", "document-uuid")

// Get private documents (requires API token)
docs, err := authClient.GetDocuments("collection-uuid")

// Create a document
newDoc := &contentzen.Document{Payload: map[string]interface{}{"title": "Test"}, Lang: "en", State: "draft"}
createdDoc, err := authClient.CreateDocument("collection-uuid", newDoc)

// Update a document
updatedDoc, err := authClient.UpdateDocument("collection-uuid", "document-uuid", createdDoc)

// Delete a document
err = authClient.DeleteDocument("collection-uuid", "document-uuid")
```

### Collections

```go
// List collections
collections, err := authClient.GetCollections()

// Get a collection
col, err := authClient.GetCollection("collection-uuid")

// Create a collection
newCol := &contentzen.Collection{Name: "products", DisplayName: "Products", Description: "Product catalog", IsPublic: false}
createdCol, err := authClient.CreateCollection(newCol)

// Update a collection
updatedCol, err := authClient.UpdateCollection("collection-uuid", createdCol)

// Delete a collection
err = authClient.DeleteCollection("collection-uuid")
```

### Media

```go
// List media
mediaList, err := authClient.ListMedia()

// Upload media
uploaded, err := authClient.UploadMedia("/path/to/file.jpg")

// Get media
media, err := authClient.GetMedia("media-uuid")

// Update media
media.AltText = "New alt text"
updatedMedia, err := authClient.UpdateMedia("media-uuid", media)

// Delete media
err = authClient.DeleteMedia("media-uuid")

// Download media
err = authClient.DownloadMedia("media-uuid", "/path/to/save.jpg")
```

### Webhooks

```go
// List webhooks
webhooks, err := authClient.ListWebhooks()

// Create webhook
newWebhook := &contentzen.Webhook{Name: "My Webhook", URL: "https://example.com/webhook", Events: []string{"document.created"}, Method: "POST"}
createdWebhook, err := authClient.CreateWebhook(newWebhook)

// Update webhook
createdWebhook.Name = "Updated Webhook"
updatedWebhook, err := authClient.UpdateWebhook("webhook-uuid", createdWebhook)

// Delete webhook
err = authClient.DeleteWebhook("webhook-uuid")
```

## Error Handling
All methods return Go errors. Always check the error value before using the result.

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## Documentation
See [ContentZen API Docs](https://www.postman.com/winter-meteor-7066631/contentzen/collection/9m7enab/contentzen-api) for full API reference.

## License
MIT
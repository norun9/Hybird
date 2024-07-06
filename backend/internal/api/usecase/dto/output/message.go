package output

// MessageOutput call from Controller layer
type MessageOutput struct {
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

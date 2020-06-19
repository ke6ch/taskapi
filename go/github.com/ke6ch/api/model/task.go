package model

// Task task type
type (
	Task struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Status    bool   `json:"status"`
		Order     int    `json:"order"`
		Timestamp string `json:"timestamp"`
	}
)

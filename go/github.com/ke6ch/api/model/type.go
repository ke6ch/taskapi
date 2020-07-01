package model

type (
	// Message message
	Message struct {
		Message string "json:message"
	}

	// Greeting Greeting
	Greeting struct {
		Message string "json:message"
	}

	// Payload ユーザ情報
	Payload struct {
		Message string `json:"message"`
	}

	// Task task type
	Task struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Status    bool   `json:"status"`
		Order     int    `json:"order"`
		Timestamp string `json:"timestamp"`
	}

	// User ユーザ情報
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

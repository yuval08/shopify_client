package models

type UserError struct {
	Field   []string `json:"field"`
	Message string   `json:"message"`
}

func (e UserError) DisplayError() string {
	message := e.Message
	if len(e.Field) > 0 {
		message += "\nFields: "
		for i, field := range e.Field {
			if i > 0 {
				message += ", "
			}
			message += field
		}
	}
	return message
}

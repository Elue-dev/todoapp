package helpers

func ValidateRequestBody(title, description string) bool {
	if title == "" || description == "" {
		return false
	} else {
		return true
	}
}

func ValidateReuestBodyForUpdate(title, description string, isCompleted interface{}) bool {
	if title == "" && description == "" && isCompleted == nil {
		return false
	} else {
		return true
	}
}
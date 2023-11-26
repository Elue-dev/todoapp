package helpers

// TODO: create helper functions for json responses
// TODO: send appropriate status codes for json responses

func ValidateRequestBody(title, description string) bool {
	if title == "" || description == "" {
		return false
	} else {
		return true
	}
}

func ValidateReuestBodyForUpdate(title, description string) bool {
	if title == "" && description == "" {
		return false
	} else {
		return true
	}
}

func UpdateFieldBasedOfValuePresence(newVal, oldVal string) string {
	if newVal != "" {
		return newVal
	}
	return oldVal
}

func ValidateSignUpFields(username, email, password string) bool {
	if username == "" || email == "" || password == "" {
		return false
	} else {
		return true
	}
}
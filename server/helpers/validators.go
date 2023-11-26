package helpers

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
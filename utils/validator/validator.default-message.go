package utilsValidator

func GetValidatorErrorMessage(errTag string) string {
	switch errTag {
	case "required":
			return "This field is required"
	case "email":
			return "Invalid email"
	}
	return "Invalid Format"
}
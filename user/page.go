package user

func Handle(inputs map[string]string) (filename string, placeholders map[string]string) {
	userProvider := Provider{}
	
	userId, ok := inputs["id"]
	if !ok {
		variables := map[string]string{
			"%errorMessage%": "User not found",
		}
		return "html/error.html", variables
	}
	
	usr, err := userProvider.FindUser(userId)
	if err != nil {
		variables := map[string]string{
			"%errorMessage%": err.Error(),
		}
		return "html/error.html", variables
	} else {
		variables := map[string]string{
			"%name%":            usr.GetFullName(),
			"%profile-picture%": usr.GetProfilePic(),
			"%profileColor%":    usr.GetProfileColor(),
			"%experience%":      usr.GetExperience(),
			"%summary%":         usr.GetSummary(),
			"%education%":       usr.GetEducation(),
		}

		// we output the contents of index.html
		return "html/user.html", variables
	}
}
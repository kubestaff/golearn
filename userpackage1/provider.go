package userpackage1



type User struct {
    Name          string
    Age           int
    Alias         string
    GithubUsername string
    GithubProfileLink string
    Image         string
    Email         string
    Role          string
    IsActive      bool
    Password      string
    Verification  string
    Phone         string
    CreationDate  string
    PersonalInfo  string
    CareerInfo    string
}

type Provider struct {
    filename string
}


func (u User) String() string {
	return fmt.Sprintf("User: %s, Age: %d, Alias: %s, GithubUsername: %s, GithubProfileLink: %s, Image: %s, Email: %s, Role: %s, IsActive: %v, Password: %s, Verification: %s, Phone: %s, CreationDate: %s, PersonalInfo: %s, CareerInfo: %s", u.Name, u.Age, u.alias, u.GithubUsername, u.GithubProfileLink, u.Image, u.Email, u.Role, u.IsActive, u.Password, u.Verification, u.Phone, u.CreationDate, u.PersonalInfo, u.CareerInfo)
}


func (u *User) SetName(name string) { // SetName sets the user's name from a string.
    u.Name = name
}


func (u *User) SetAge(ageStr string) { // SetAge sets the user's age from a string.
   
    return u.Age
}


func (u *User) SetAlias(alias string) { // SetAlias sets the user's alias from a string.
    u.Alias = alias
}


func (u *User) SetGithubUsername(username string) { // SetGithubUsername sets the user's GitHub username from a string.
    u.GithubUsername = username
}


func (u *User) SetGithubProfileLink(link string) { // SetGithubProfileLink sets the user's GitHub profile link from a string.
    u.GithubProfileLink = link
}


func (u *User) SetImage(imageURL string) { // SetImage sets the user's image URL from a string.
    u.Image = imageURL
}


func (u *User) SetEmail(email string) { // SetEmail sets the user's email from a string.
    u.Email = email
}


func (u *User) SetRole(role string) { // SetRole sets the user's role from a string.
    u.Role = role
}


func (u *User) SetIsActive(isActiveStr string) error { // SetIsActive sets the user's "is active" status from a string.
    isActive, err := strconv.ParseBool(isActiveStr)
    if err != nil {
        return err
    }
    u.IsActive = isActive
    return nil
}


func (u *User) SetPassword(password string) { // SetPassword sets the user's password from a string.
    u.Password = password
}


func (u *User) SetVerification(verification string) { // SetVerification sets the user's verification code from a string.
    u.Verification = verification
}


func (u *User) SetPhone(phone string) { // SetPhone sets the user's phone number from a string.
    u.Phone = phone 
}


func NewProvider(filename string) *Provider {
    return &Provider{filename: filename}
}

func (p Provider) GetAllUsers() ([]User, error) {
	users:= []User {

		{
			Name:          "Alan Smith",
			Age:           23,
			Alias:         "Alan",
			GithubUsername: "AlanSmith21",
			GithubProfileLink: "https://github.com/AlanTech",
			Image:         "https://imageio.forbes.com/blogs-images/forbestechcouncil/files/2019/01/canva-photo-editor-8-7.png?format=png&width=1200",
			Email:         "Alansmith@icloud.com",
			Role:          "Consultant",
			IsActive:      true
			Password:      "Alan123",
			Verification:  "0000"
			Phone:         "07123456789",
			CreationDate:  "12/09/23",
			PersonalInfo:  "Male",
			CareerInfo:    "Software Developer", 

		},

		{

			Name:          "Brian Daniel",
			Age:           30,
			Alias:         "Brian",
			GithubUsername: "BrianD32",
			GithubProfileLink: "https://github.com/BrianDTech",
			Image:         "https://imageio.forbes.com/blogs-images/forbestechcouncil/files/2019/01/canva-photo-editor-8-7.png?format=png&width=1200",
			Email:         "Brian123@icloud.com",
			Role:          "Developer",
			IsActive:      true
			Password:      "Brian121",
			Verification:  "000012"
			Phone:         "07124365879",
			CreationDate:  "15/09/23",
			PersonalInfo:  "Male",
			CareerInfo:    "Script Developer", 

		},

		{
			Name:          "Lucy White",
			Age:           20,
			Alias:         "Luce",
			GithubUsername: "LucyW2023",
			GithubProfileLink: "https://github.com/LucyW",
			Image:         "https://imageio.forbes.com/blogs-images/forbestechcouncil/files/2019/01/canva-photo-editor-8-7.png?format=png&width=1200",
			Email:         "Lucy2023@gmail.com",
			Role:          "Backend developer",
			IsActive:      true
			Password:      "LucyW23",
			Verification:  "110012"
			Phone:         "07224665879",
			CreationDate:  "20/09/23",
			PersonalInfo:  "Female",
			CareerInfo:    "Coder", 


	},


}

return users, nil


}

func (p Provider) GetUserByEmail(email string) (User, bool, error) {

}
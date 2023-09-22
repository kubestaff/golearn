package userinterface

import (

)

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


func (u *User) SetCreationDate(creationDateStr string) error { // SetCreationDate sets the user's creation date from a string in a specific format.
    creationDate, err := 
    if err != nil {
        return err
    }
    u.CreationDate = creationDate
    return nil
}

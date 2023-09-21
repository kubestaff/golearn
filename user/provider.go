package user

type Provider struct{}

func (p Provider) GetCurrentUser() (User, error) {
	users := Users{
		{
			Name:       "Adam",
			Surname:    "Smith",
			ProfilePic: "user1.png",
			ProfileColor: "#004daa",
			Experience:   []string{"Office 365, Coding, Piloting"},
		},
		{
			Name:       "Kariana",
			Surname:    "Corky",
			ProfilePic: "user2.jpg",
			ProfileColor: "#ffffff",
			Experience:   []string{"Pilot"},
		},
		{
			Name:       "Mariana",
			Surname:    "Orkey",
			ProfilePic: "user3.jpeg",
			ProfileColor: "#e8e6e3",
			Experience:   []string{"Interior Designer"},
		},
		{
			Name:       "Oman",
			Surname:    "Osas",
			ProfilePic: "user4.jpeg",
			ProfileColor: "#454a4d",
			Experience:   []string{"Developed and maintained company websites"},
		},
		{
			Name:       "Adam",
			Surname:    "Loqwick",
			ProfilePic: "user5.jpeg",
			ProfileColor: "#0f0f0f",
			Experience:   []string{"Web Developer, XYZ Company, 2016 - Present"},
		},
		{
			Name:       "Ariana",
			Surname:    "Grande",
			ProfilePic: "user6.jpeg",
			ProfileColor: "#f1f1f1",
			Experience:   []string{"Developer, Instructor"},
		},
		{
			Name:    "Josh",
			Surname: "Smith",
			Experience:   []string{"GunMan, Driving"},
		},
		{
			Name:    "John",
			Surname: "Doe",
			Experience:   []string{"Managment, Teaching, Coding"},
		},
		{
			Name:    "Bart",
			Surname: "Simpson",
			Experience:   []string{"Office 365, Managment"},
		},
		{
			Name:    "Ojan",
			Surname: "Olypmic",
			Experience:   []string{"Office 365, Web Developer"},
		},
	}
	user := users.GetRandomUser()

	return user, nil
}

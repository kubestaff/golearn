package user

type Provider struct{}

func (p Provider) GetCurrentUser() (User, error) {
	users := Users{
		{
			Name:         "Adam",
			Surname:      "Smith",
			ProfilePic:   "user1.png",
			ProfileColor: "#337df5",
			Experience:   []string{"Office 365, Coding, Piloting"},
		},
		{
			Name:         "Fernanda",
			Surname:      "Cross",
			ProfilePic:   "user2.jpeg",
			ProfileColor: "#f533e8",
			Experience:   []string{"Pilot"},
		},
		{
			Name:         "Andrew",
			Surname:      "Potter",
			ProfilePic:   "user3.jpg",
			ProfileColor: "#33f543",
			Experience:   []string{"Interior Designer"},
		},
		{
			Name:         "Caylee",
			Surname:      "Arias",
			ProfilePic:   "user4.png",
			ProfileColor: "#33f543",
			Experience:   []string{"Developed and maintained company websites"},
		},
		{
			Name:         "Zaiden",
			Surname:      "Winters",
			ProfilePic:   "user5.jpg",
			ProfileColor: "#507553",
			Experience:   []string{"Web Developer, XYZ Company, 2016 - Present"},
		},
		{
			Name:         "Ibrahim",
			Surname:      "House",
			ProfilePic:   "user6.jpeg",
			ProfileColor: "#506c75",
			Experience:   []string{"Developer, Instructor"},
		},
		{
			Name:         "Jalen",
			Surname:      "Hanna",
			ProfilePic:   "user7.jpg",
			ProfileColor: "#bb8215",
			Experience:   []string{"GunMan, Driving"},
		},
		{
			Name:    "German",
			Surname: "Ali",
			ProfileColor: "#383734",
			Experience:   []string{"Managment, Teaching, Coding"},
		},
		{
			Name:    "Shane",
			Surname: "Snyder",
			Experience:   []string{"Office 365, Managment"},
		},
		{
			Name:    "Evie",
			Surname: "Riggs",
			Experience:   []string{"Office 365, Web Developer"},
		},
		{
			Name:    "Lyric",
			Surname: "Rowe",
		},
		{
			Name:    "Roger",
			Surname: "Morse",
		},
		{
			Name:    "Jameson",
			Surname: "Pruitt",
		},
	}

	user := users.GetRandomUser()

	return user, nil
}

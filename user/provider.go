package user

type Provider struct{}

func (p Provider) GetCurrentUser() (User, error) {
	users := Users{
		{
			Name:         "Adam",
			Surname:      "Smith",
			ProfilePic:   "user1.png",
			ProfileColor: "#337df5",
		},
		{
			Name:         "Fernanda",
			Surname:      "Cross",
			ProfilePic:   "user2.jpeg",
			ProfileColor: "#f533e8",
		},
		{
			Name:         "Andrew",
			Surname:      "Potter",
			ProfilePic:   "user3.jpg",
			ProfileColor: "#33f543",
		},
		{
			Name:         "Caylee",
			Surname:      "Arias",
			ProfilePic:   "user4.png",
			ProfileColor: "#33f543",
		},
		{
			Name:         "Zaiden",
			Surname:      "Winters",
			ProfilePic:   "user5.jpg",
			ProfileColor: "#507553",
		},
		{
			Name:         "Ibrahim",
			Surname:      "House",
			ProfilePic:   "user6.jpeg",
			ProfileColor: "#506c75",
		},
		{
			Name:         "Jalen",
			Surname:      "Hanna",
			ProfilePic:   "user7.jpg",
			ProfileColor: "#bb8215",
		},
		{
			Name:    "German",
			Surname: "Ali",
			ProfileColor: "#383734",
		},
		{
			Name:    "Shane",
			Surname: "Snyder",
		},
		{
			Name:    "Evie",
			Surname: "Riggs",
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

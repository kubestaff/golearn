package user

type Provider struct{}

func (p Provider) GetCurrentUser() (User, error) {
	users := Users{
		{
			Name:         "Adam",
			Surname:      "Smith",
			ProfilePic:   "user1.png",
			ProfileColor: "#004daa",
			Experience:   []string{"Office 365, Coding, Piloting"},
			Summary:      []string{"Been coding for 2 years and just started my pilot licience"},
			Education:    []string{"University"},
		},
		{
			Name:         "Kariana",
			Surname:      "Corky",
			ProfilePic:   "user2.jpg",
			ProfileColor: "#ffffff",
			Experience:   []string{"Pilot"},
			Summary:      []string{"Been flying for 12 years. Aprox. 4500h flown with the A320."},
			Education:    []string{"University"},
		},
		{
			Name:         "Mariana",
			Surname:      "Orkey",
			ProfilePic:   "user3.jpeg",
			ProfileColor: "#e8e6e3",
			Experience:   []string{"Interior Designer"},
			Summary:      []string{"Designed 100+ houses across the US"},
			Education:    []string{"High School"},
		},
		{
			Name:         "Oman",
			Surname:      "Osas",
			ProfilePic:   "user4.jpeg",
			ProfileColor: "#454a4d",
			Experience:   []string{"Developed and maintained company websites"},
			Summary:      []string{"Junior Role in Amazon. Have been coding for 2 year"},
			Education:    []string{"Collage"},
		},
		{
			Name:         "Adam",
			Surname:      "Loqwick",
			ProfilePic:   "user5.jpeg",
			ProfileColor: "#0f0f0f",
			Experience:   []string{"Web Developer, XYZ Company, 2016 - Present"},
			Summary:      []string{"Junior Role in DHL. Have been coding for 1 year"},
			Education:    []string{"Collage"},
		},
		{
			Name:         "Ariana",
			Surname:      "Grande",
			ProfilePic:   "user6.jpeg",
			ProfileColor: "#f1f1f1",
			Experience:   []string{"Developer, Instructor"},
			Summary:      []string{"Been developing for 8 years, main language is Golang."},
			Education:    []string{"Secondary School"},
		},
		{
			Name:       "Josh",
			Surname:    "Smith",
			Experience: []string{"GunMan, Driving"},
			Summary:    []string{"Person who is very talented driver. Drove in F2"},
			Education:  []string{"University"},
		},
		{
			Name:       "John",
			Surname:    "Doe",
			Experience: []string{"Managment, Teaching, Coding"},
			Summary:    []string{"This person is a dedicated environmentalist who spends their free time volunteering for conservation organizations, enjoys hiking in national parks, and practices sustainable living in their everyday life."},
			Education:  []string{"Collage"},
		},
		{
			Name:       "Bart",
			Surname:    "Simpson",
			Experience: []string{"Office 365, Managment"},
			Summary:    []string{"Managed the biggest parties in Germany. Been managing people sice I was a small boy."},
			Education:  []string{"High School"},
		},
		{
			Name:       "Ojan",
			Surname:    "Olypmic",
			Experience: []string{"Office 365, Web Developer"},
			Summary:    []string{"Web Developer for Amazon. Senior Role been at Amazon for 5 years."},
			Education:  []string{"University"},
		},
	}
	user := users.GetRandomUser()

	return user, nil
}

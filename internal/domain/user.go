package domain

type Role struct {
	ID       int `json:"id"`
	RoleID   int `json:"roleid"`
	PersonID int `json:"personid"`
}

type Person struct {
	PersonID int    `json:"personid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	// education??
}

//?????????????????
type Admin struct {
}

type TeamMembers struct {
	TeamID   int `json:"teamid"`
	PersonID int `json:"personid"`
}

type Team struct {
	TeamID   int    `json:"teamid"`
	Name     string `json:"name"`
	LeaderID int    `json:"leaderid"`
	Token    int    `json:"token"`
}

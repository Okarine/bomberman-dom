package model

type Request struct {
	Type         string `json:"type"`
	Username     string `json:"username"`
	Message      string `json:"message"`
	SessionToken string `json:"sessionToken"`
	Coordinates  `json:"coordinates"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

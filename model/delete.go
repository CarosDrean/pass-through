package model

type RequestDeleteInput struct {
	Action          string `json:"action"`
	Commentary      string `json:"commentary"`
	CurrentState    string `json:"currentState"`
	FollowingNumber string `json:"followingNumber"`
	IdDoc           string `json:"idDoc"`
	Profile         string `json:"profile"`
}

type RequestDeleteOutput struct {
	Action          string `json:"accion"`
	Commentary      string `json:"comentario"`
	CurrentState    string `json:"estadoActual"`
	FollowingNumber string `json:"numeroSeguimiento"`
	IdDoc           string `json:"idDoc"`
	Profile         string `json:"perfil"`
}

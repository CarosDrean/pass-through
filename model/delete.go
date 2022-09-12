package model

type RequestDelete struct {
	Action          string `json:"action" json-out:"accion"`
	Commentary      string `json:"commentary" json-out:"comentario"`
	CurrentState    string `json:"currentState" json-out:"estadoActual"`
	FollowingNumber string `json:"followingNumber" json-out:"numeroSeguimiento"`
	IdDoc           string `json:"idDoc" json-out:"idDoc"`
	Profile         string `json:"profile" json-out:"perfil"`
}

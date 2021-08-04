package models

/* RespuestaLogin has the responseBody returned at login */
type RespuetaLogin struct {
	Token string `json:"token,omitempty"`
}

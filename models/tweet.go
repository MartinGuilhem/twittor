package models

/* Tweet captures the message coming from body */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}

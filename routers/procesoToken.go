package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/martinguilhem/twittor/bd"
	"github.com/martinguilhem/twittor/models"
)

var (
	Email     string
	IDUsuario string
)

/* ProcesoToken process the token to extact its values */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid format token")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}
	return claims, false, "", err
}

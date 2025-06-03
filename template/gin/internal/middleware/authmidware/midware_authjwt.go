package authmidware

import (
	"boilerplate/config"
	errorUc "boilerplate/internal/error"
	"boilerplate/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/saucon/sauron/v2/pkg/log"
	"strings"
)

type AuthHandler struct {
	cfg config.Config
	log *log.LogCustom
}

func New(r *gin.RouterGroup, c config.Config, l *log.LogCustom) {
	handler := AuthHandler{
		cfg: c,
		log: l,
	}

	r.Use(handler.AuthJwt)
}

func (au *AuthHandler) AuthJwt(c *gin.Context) {
	v := c.Request.Header["Authorization"]
	if v == nil {
		au.log.Error(log.LogData{
			Err: utils.MakeError(errorUc.ErrUnauthorized),
		})
		utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
		c.Abort()
		return
	} else {
		authVal := strings.Split(v[0], " ")
		if authVal[0] != "Bearer" || len(authVal) == 1 {
			utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
			c.Abort()
			return
		}
		tokenHeader := authVal[1]
		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				au.log.Error(log.LogData{
					Err: utils.MakeError(errorUc.ErrUnauthorized),
				})
				utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
				c.Abort()
				return nil, utils.MakeError(errorUc.ErrUnauthorized)
			}

			return []byte("secret"), nil
		})
		// parsing errors result
		if err != nil {
			au.log.Error(log.LogData{
				Err: err,
			})
			utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name := claims["name"].(string)
			userId := claims["userid"].(string)
			username := claims["username"].(string)

			c.Set("name", name)
			c.Set("userid", userId)
			c.Set("username", username)
		} else {
			au.log.Error(log.LogData{
				Err: utils.MakeError(errorUc.ErrUnauthorized),
			})
			utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
			c.Abort()
			return
		}

		// if there's a token
		if !token.Valid {
			au.log.Error(log.LogData{
				Err: utils.MakeError(errorUc.ErrUnauthorized),
			})
			utils.Failed(c, utils.CustomError(errorUc.ErrorCustom(utils.MakeError(errorUc.ErrUnauthorized))))
			c.Abort()
			return
		}
	}
	c.Next()
}

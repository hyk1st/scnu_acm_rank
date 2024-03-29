package middle

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"scnu_acm_rank/biz/model"
	"scnu_acm_rank/biz/reqModel"
	"time"
)

//type userStatus struct {
//	name   string
//	vjName string
//	stuID  string
//	level  int
//}

var identityKey = "token"

func GetJWT() (*jwt.HertzJWTMiddleware, error) {
	temp, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("hyk1st"),
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[identityKey]
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals reqModel.LoginReq
			if err := c.BindForm(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.StuId
			password := loginVals.Password
			user := model.User{}
			model.DB.Model(&model.User{}).Where("stu_id = ?", userID).Find(&user)

			if password == user.Password {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			mp := data.(map[string]interface{})
			c.Set("user", &model.User{
				//Email:  mp["email"].(string),
				//VjName: mp["vj_name"].(string),
				StuId: int64(mp["stu_id"].(float64)),
				//Name:   mp["name"].(string),
				Level: int(mp["level"].(float64)),
			})
			return true
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, FailResp(errors.New("未登录")))
		},
	})

	temp.LoginResponse = func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
		tk, err := temp.ParseTokenString(token)
		if err != nil {
			c.JSON(http.StatusOK, utils.H{
				"status": 0,
				"data": map[string]interface{}{
					"token":  "Bearer " + token,
					"expire": expire.Format(time.RFC3339),
				},
				"msg": "登录成功",
			})
			return
		}
		mp := jwt.ExtractClaimsFromToken(tk)
		mp = mp["token"].(map[string]interface{})
		delete(mp, "password")
		c.JSON(http.StatusOK, utils.H{
			"status": 0,
			"data": map[string]interface{}{
				"token":  "Bearer " + token,
				"expire": expire.Format(time.RFC3339),
				"user":   mp,
			},
			"msg": "登录成功",
		})
	}

	return temp, err
}

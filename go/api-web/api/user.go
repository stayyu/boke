package api

import (
	"boke/go/api-web/froms"
	"boke/go/api-web/global"
	"boke/go/api-web/middlewares"
	"boke/go/api-web/models"
	"boke/go/api-web/proto"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func Register(c *gin.Context) {

	registerForm := froms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
	}

	////验证码
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:"127.0.0.1:8081",
	//})
	//value, err := rdb.Get(context.Background(),registerForm.Code).Result()
	//if err == redis.Nil{
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":"验证码错误",
	//	})
	//	return
	//}else{
	//	if value != registerForm.Code {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"code":"验证码错误",
	//		})
	//		return
	//	}
	//}

	user, err := global.UsersrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NiceName: registerForm.NickName,
		PassWord: registerForm.PassWord,
	})

	if err != nil {
		fmt.Println("创建用户失败")
		fmt.Println(err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func PassWordLogin(c *gin.Context) {
	passwordLoginForm := froms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	//登录的逻辑
	rsp, err := global.UsersrvClient.GetUsername(context.Background(), &proto.NameRequest{
		NickName: passwordLoginForm.NickName,
	})

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"name": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"name": "登录失败",
				})
			}
			return
		}
	} else {
		//只是查询到用户了而已，并没有检查密码
		if passRsp, pasErr := global.UsersrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Suess {
				//生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "imooc",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}

}

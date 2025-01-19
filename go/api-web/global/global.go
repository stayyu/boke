package global

import (
	"boke/go/api-web/proto"
	ut "github.com/go-playground/universal-translator"
)

var (
	BlogSrvClient proto.BlogClient
	UsersrvClient proto.UserClient

	Trans ut.Translator
)

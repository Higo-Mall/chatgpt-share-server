package admin

import (
	"backend/modules/chatgpt/service"

	"github.com/cool-team-official/cool-admin-go/cool"
)

type ChatgptSessionController struct {
	*cool.Controller
}

func init() {
	var chatgpt_session_controller = &ChatgptSessionController{
		&cool.Controller{
			Prefix:  "/admin/chatgpt/session",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewChatgptSessionService(),
		},
	}
	// 注册路由
	cool.RegisterController(chatgpt_session_controller)
}

// // 增加 Welcome 演示 方法
// type ChatgptSessionWelcomeReq struct {
// 	g.Meta `path:"/welcome" method:"GET"`
// }
// type ChatgptSessionWelcomeRes struct {
// 	*cool.BaseRes
// 	Data interface{} `json:"data"`
// }

// func (c *ChatgptSessionController) Welcome(ctx context.Context, req *ChatgptSessionWelcomeReq) (res *ChatgptSessionWelcomeRes, err error) {
// 	res = &ChatgptSessionWelcomeRes{
// 		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
// 		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
// 	}
// 	return
// }

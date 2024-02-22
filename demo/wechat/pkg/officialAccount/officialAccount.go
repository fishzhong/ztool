package officialAccount

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gowechat/example/config"
	"github.com/gowechat/example/pkg/util"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
)

//ExampleOfficialAccount 公众号操作样例
type ExampleOfficialAccount struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

//ExampleOfficialAccount new
func NewExampleOfficialAccount(wc *wechat.Wechat) *ExampleOfficialAccount {
	//init config
	globalCfg := config.GetConfig()
	offCfg := &offConfig.Config{
		AppID:          globalCfg.AppID,
		AppSecret:      globalCfg.AppSecret,
		Token:          globalCfg.Token,
		EncodingAESKey: globalCfg.EncodingAESKey,
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &ExampleOfficialAccount{
		wc:              wc,
		officialAccount: officialAccount,
	}
}

//Serve 处理消息
func (ex *ExampleOfficialAccount) Serve(c *gin.Context) {
	// 传入request和responseWriter
	server := ex.officialAccount.GetServer(c.Request, c.Writer)
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO
		fmt.Println("消息雷系", msg.MsgType)
		fmt.Println(msg.Event)
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
		//articles := []*message.Article{article1}
		//news := message.NewNews(articles)
		//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

		//voice := message.NewVoice(mediaID)
		//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

		//
		//video := message.NewVideo(mediaID, "标题", "描述")
		//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

		//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
		//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

		//多客服消息转发
		//transferCustomer := message.NewTransferCustomer("")
		//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Error("Serve Error, err=%+v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Error("Send Error, err=%+v", err)
		return
	}
}

func (ex *ExampleOfficialAccount) Redirect(c *gin.Context) {
	err := ex.officialAccount.GetOauth().Redirect(c.Writer, c.Request, "http://yhf.nat300.top/api/v1/token", "snsapi_base", "123")
	if err != nil {
		fmt.Println("dayincuowu", err.Error())
		return
	}
}

func (ex *ExampleOfficialAccount) AccToken(c *gin.Context) {
	code := c.Query("code")
	token, err := ex.officialAccount.GetOauth().GetUserAccessToken(code)
	if err != nil {
		return
	}
	fmt.Println(token.OpenID, token.UnionID)
	//根据openId查询用户信息

	c.JSON(200, "openId:"+token.OpenID+" uniId:"+token.UnionID)
}

func (ex *ExampleOfficialAccount) Template(c *gin.Context) {
	send, err := ex.officialAccount.GetTemplate().Send(&message.TemplateMessage{
		ToUser:     "o2SLY6OnPjrCemOyGZ2xsW3zEGf4",
		TemplateID: "ZJEfenbk5hn2FVzR2GQplCHHFeeHdITeFWgxHaDj1g4",
		URL:        "https://www.baidu.com",
		Data: map[string]*message.TemplateDataItem{
			"thing":        {"孙悟空", "#173177"},
			"phone_number": {"13991398989", "#173177"},
		},
	})
	if err != nil {
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, send)
}

func (ex *ExampleOfficialAccount) GetQrTicket(c *gin.Context) {
	send, err := ex.officialAccount.GetBasic().GetQRTicket(&basic.Request{
		ExpireSeconds: 300,
		ActionName:    "QR_STR_SCENE",
		ActionInfo: struct {
			Scene struct {
				SceneStr string `json:"scene_str,omitempty"`
				SceneID  int    `json:"scene_id,omitempty"`
			} `json:"scene"`
		}{
			//{"login_" + fmt.Sprintf(`{"channel_number": "%s", "ct": "%s"}`, "PV00001", 1)},
		},
	})

	//	data := make(map[string]interface{})
	//	data["expire_seconds"] = 300
	//	data["action_name"] = "QR_STR_SCENE"
	//	action := make(map[string]interface{})
	//	action["scene"] = map[string]string{"scene_str": "login_" + fmt.Sprintf(`{"channel_number": "%s", "ct": "%s"}`, params.ChannelNumber, params.Ct)}
	//	data["action_info"] = action
	//	TicketDataBate, _ := json.Marshal(data)
	//	return TicketDataBate

	if err != nil {
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, send)
}

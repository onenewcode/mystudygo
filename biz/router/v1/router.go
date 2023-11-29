package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	//"github.com/hertz-contrib/swagger"
	//swaggerFiles "github.com/swaggo/files"
)

//context.Context 与 RequestContext 都有存储值的能力，具体选择使用哪一个上下文有个简单依据：所储存值的生命周期和所选择的上下文要匹配。
//
//ctx 主要用来存储请求级别的变量，请求结束就回收了，特点是查询效率高（底层是 map），协程不安全，且未实现 context.Context 接口。
//
//c 作为上下文在中间件 /handler 之间传递，协程安全。所有需要 context.Context 接口作为入参的地方，直接传递 c 即可。

func InitRouter(r *server.Hertz) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	//
	//apiRouter := r.Group("/douyin")
	//
	//// basic apis
	//apiRouter.GET("/feed/", dto.Feed)
	//apiRouter.GET("/user/", handler.UserInfo)
	//apiRouter.POST("/user/register/", handler.Register)
	//apiRouter.POST("/user/login/", handler.Login)
	//apiRouter.POST("/publish/action/", handler.Publish)
	//apiRouter.GET("/publish/list/", handler.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", handler.FavoriteAction)
	//apiRouter.GET("/favorite/list/", handler.FavoriteList)
	//apiRouter.POST("/comment/action/", handler.CommentAction)
	//apiRouter.GET("/comment/list/", handler.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", handler.RelationAction)
	//apiRouter.GET("/relation/follow/list/", handler.FollowList)
	//apiRouter.GET("/relation/follower/list/", handler.FollowerList)
	//apiRouter.GET("/relation/friend/list/", handler.FriendList)
	//apiRouter.GET("/message/chat/", handler.MessageChat)
	//apiRouter.POST("/message/action/", handler.MessageAction)
	//
	//apiRouter.GET("/test/", func(c context.Context, ctx *app.RequestContext) {
	//	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	//})
	////swagger
	//url := swagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	//apiRouter.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}

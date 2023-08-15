// Code generated by hertz generator.

package api

import (
	"context"

	api "GoYin/server/service/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishVideo .
// @router /douyin/publish/action/ [POST]
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinPublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoList .
// @router /douyin/publish/list/ [GET]
func VideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinPublishListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Favorite .
// @router /douyin/favorite/action/ [POST]
func Favorite(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinFavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Comment .
// @router /douyin/comment/action/ [POST]
func Comment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinCommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// Action .
// @router /douyin/relation/action/ [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowingList .
// @router /douyin/relation/follow/list/ [GET]
func FollowingList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FriendList .
// @router /douyin/relation/friend/list/ [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}

// ChatHistory .
// @router /douyin/message/chat/ [GET]
func ChatHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinMessageChatResponse)

	c.JSON(consts.StatusOK, resp)
}

// SentMessage .
// @router /douyin/message/action/ [POST]
func SentMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinMessageActionResponse)

	c.JSON(consts.StatusOK, resp)
}
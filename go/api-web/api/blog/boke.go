package blog

import (
	"boke/go/api-web/api"
	"boke/go/api-web/froms"
	"boke/go/api-web/global"
	"boke/go/api-web/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Detail(ctx *gin.Context) {
	pages := ctx.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(pages)

	rsp, err := global.BlogSrvClient.GetBlogDetail(context.Background(), &proto.BlogInfoRequest{Id: int32(id)})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, rsp)

}
func New(ctx *gin.Context) {
	bokeForm := froms.BokeForm{}
	if err := ctx.ShouldBindJSON(&bokeForm); err != nil {
		api.HandleValidatorError(ctx, err)

		fmt.Println(err)
		return
	}

	rsp, err := global.BlogSrvClient.CreateBlog(context.Background(), &proto.CreateBlogInfo{
		Title:      bokeForm.Title,
		Context:    bokeForm.Content,
		CategoryId: bokeForm.CategoryId,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
	}
	ctx.JSON(http.StatusOK, rsp)
}
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
	}
	_, err = global.BlogSrvClient.DeleteBlog(context.Background(), &proto.DeleteBlogInfo{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
	}
	ctx.Status(http.StatusOK)
	return
}
func Update(ctx *gin.Context) {
	blokform := froms.BokeForm1{}
	if err := ctx.ShouldBindJSON(&blokform); err != nil {
		api.HandleValidatorError(ctx, err)
		fmt.Println(err)
		return
	}

	fmt.Println("a2")
	if _, err := global.BlogSrvClient.UpdateBlog(context.Background(), &proto.UpdateBlogsInfo{
		Id:         blokform.Id,
		CategoryId: blokform.CategoryId,
		Context:    blokform.Content,
		Title:      blokform.Title,
	}); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func List(ctx *gin.Context) {
	request := &proto.BlogFilterRequest{}
	//页数
	pages := ctx.DefaultQuery("p", "0")
	pageInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pageInt)
	//页面数量
	panums := ctx.DefaultQuery("pnum", "10")
	panumsInt, _ := strconv.Atoi(panums)
	request.PagePerNums = int32(panumsInt)

	categoryId := ctx.DefaultQuery("cid", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.Categoryid = int32(categoryIdInt)

	contents := ctx.DefaultQuery("content", "0")
	request.Content = contents

	KeyWords := ctx.DefaultQuery("keyword", "0")
	request.Content = KeyWords

	r, err := global.BlogSrvClient.GetBlogList(context.WithValue(context.Background(), "ginCoontext", ctx), request)
	if err != nil {
		fmt.Println("查询失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := map[string]interface{}{
		"total": r.Total,
	}
	bokeList := make([]interface{}, 0)
	for _, value := range r.Data {
		if len(value.Content) > 100 {
			bokeList = append(bokeList, map[string]interface{}{
				"pages":       request.Pages,
				"content":     value.Content[:100],
				"pagePerNums": request.PagePerNums,
				"categoryid":  value.Categoryid,
				"title":       value.Title,
				"createtime":  value.CreateTime,
				"id":          value.Id,
			})
		} else {
			bokeList = append(bokeList, map[string]interface{}{
				"pages":       request.Pages,
				"content":     value.Content,
				"pagePerNums": request.PagePerNums,
				"categoryid":  value.Categoryid,
				"title":       value.Title,
				"createtime":  value.CreateTime,
				"id":          value.Id,
			})
		}

	}
	reMap["date"] = bokeList
	ctx.JSON(http.StatusOK, reMap)
}

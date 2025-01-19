package cate

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

func DetailList(ctx *gin.Context) {
	request := &proto.CategoryInfoRequest{}

	Id := ctx.DefaultQuery("id", "0")
	categoryIdInt, _ := strconv.Atoi(Id)
	request.Id = int32(categoryIdInt)

	name := ctx.DefaultQuery("name", "0")
	request.Name = name
	r, err := global.BlogSrvClient.GetCategoryDetail(context.WithValue(context.Background(), "ginCoontext", ctx), request)
	if err != nil {
		fmt.Println("a1")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": r.Total,
	}
	cateList := make([]interface{}, 0)
	for _, value := range r.Data {
		cateList = append(cateList, map[string]interface{}{
			"id":   value.Id,
			"name": value.Name,
		})
	}
	reMap["date"] = cateList
	ctx.JSON(http.StatusOK, reMap)
}
func New(ctx *gin.Context) {

	categoryFrom := froms.CategoryFrom{}
	if err := ctx.ShouldBindJSON(&categoryFrom); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	blogClient := global.BlogSrvClient
	i, err := strconv.ParseInt(categoryFrom.Id, 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	rsp, err := blogClient.CreateCategory(context.Background(), &proto.CreateCategoryInfo{
		Name: categoryFrom.Name,
		Id:   int32(i),
	})
	if err != nil {
		fmt.Println("创建类别失败")
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
	_, err = global.BlogSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryInfo{Id: int32(i)})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
	}
	ctx.Status(http.StatusOK)
	return
}
func Update(ctx *gin.Context) {
	categoryFrom := froms.CategoryFrom1{}
	if err := ctx.ShouldBindJSON(&categoryFrom); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	if _, err := global.BlogSrvClient.UpdateCategory(context.Background(), &proto.CreateCategoryInfo{Name: categoryFrom.Name, Id: categoryFrom.Id}); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

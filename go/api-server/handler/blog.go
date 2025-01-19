package handler

import (
	"boke/go/api-server/global"
	"boke/go/api-server/model"
	"boke/go/api-server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type BlogServer struct {
	proto.UnimplementedBlogServer
}

func (b *BlogServer) CreateBlog(ctx context.Context, req *proto.CreateBlogInfo) (*proto.BlogInfoResponse, error) {
	var bl model.Blog
	bl.Title = req.Title
	bl.Content = req.Context
	bl.Category.ID = req.CategoryId
	su := global.DB.Create(&bl)

	if su.Error != nil {
		return nil, status.Errorf(codes.Internal, su.Error.Error())
	}
	return &proto.BlogInfoResponse{Susee: true}, nil
}

func (b *BlogServer) DeleteBlog(ctx context.Context, req *proto.DeleteBlogInfo) (*proto.BlogInfoResponse, error) {
	result := global.DB.Delete(&model.Blog{BaseModel: model.BaseModel{ID: req.Id}}, req.Id)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.BlogInfoResponse{Susee: true}, nil
}

func (b *BlogServer) UpdateBlog(ctx context.Context, req *proto.UpdateBlogsInfo) (*proto.BlogInfoResponse, error) {
	var bl model.Blog
	//result := global.DB.Where(&model.Blog{BaseModel: model.BaseModel{ID: req.Id}}).First(&bl)
	result := global.DB.First(&bl, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	bl.Title = req.Title
	bl.Content = req.Context
	bl.CategoryID = int32(req.CategoryId)
	result = global.DB.Save(&bl)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.BlogInfoResponse{Susee: true}, nil
}

func (b *BlogServer) GetBlogDetail(ctx context.Context, req *proto.BlogInfoRequest) (*proto.BlogInfoResponses, error) {
	var bl model.Blog
	fmt.Print(req.Id)
	result := global.DB.Where(&model.BaseModel{ID: req.Id}).First(&bl)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "类别不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.BlogInfoResponses{
		Title:      bl.Title,
		CreateTime: bl.CreatedAt.String(),
		Content:    bl.Content,
		Id:         bl.ID,
		Categoryid: bl.CategoryID,
	}, nil
}

func (b *BlogServer) GetBlogList(ctx context.Context, req *proto.BlogFilterRequest) (*proto.BlogListResponse, error) {
	var blogs []model.Blog
	result := global.DB.Find(&blogs)
	rsp := &proto.BlogListResponse{}
	rsp.Total = int32(result.RowsAffected)
	if req.KeyWords == "0" {
		result = global.DB.Find(&blogs)
	} else {
		result = global.DB.Where("title LIKE ?", req.KeyWords).Or("content LIKE ?", req.KeyWords).Find(&blogs)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&blogs)
	for _, blo := range blogs {
		bloginfo := proto.BlogFilterRequest{
			Categoryid: blo.CategoryID,
			Content:    blo.Content,
			Title:      blo.Title,
			CreateTime: blo.CreatedAt.String(),
			Id:         blo.ID,
		}
		rsp.Data = append(rsp.Data, &bloginfo)
	}
	return rsp, nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (c *BlogServer) CreateCategory(ctx context.Context, req *proto.CreateCategoryInfo) (*proto.CategoryInfoResponse, error) {
	var cate model.Category
	result := global.DB.Where(&model.Category{Name: req.Name}).First(&cate)
	if result.RowsAffected == 1 {
		return nil, status.Error(codes.AlreadyExists, "类别已经存在")
	}
	cate.Name = req.Name
	result = global.DB.Create(&cate)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.CategoryInfoResponse{
		Susee: true,
	}, nil
}

func (c *BlogServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryInfo) (*proto.CategoryInfoResponse, error) {
	result := global.DB.Delete(&model.Category{BaseModel: model.BaseModel{ID: req.Id}}, req.Id)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.CategoryInfoResponse{
		Susee: true,
	}, nil
}

func (c *BlogServer) UpdateCategory(ctx context.Context, req *proto.CreateCategoryInfo) (*proto.CategoryInfoResponse, error) {
	var cate model.Category
	result := global.DB.Where(&model.Category{BaseModel: model.BaseModel{ID: req.Id}}).First(&cate)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.AlreadyExists, "类别不存在")
	}
	cate.Name = req.Name
	result.Save(&cate)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.CategoryInfoResponse{
		Susee: true,
	}, nil
}

func (c *BlogServer) GetCategoryDetail(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponses, error) {

	var cate []model.Category
	result := global.DB.Find(&cate)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.CategoryInfoResponses{}
	rsp.Total = int32(result.RowsAffected)
	for _, blo := range cate {
		cateinfo := proto.CateGoryFilterRequest{
			Id:   blo.ID,
			Name: blo.Name,
		}

		rsp.Data = append(rsp.Data, &cateinfo)
	}

	return rsp, nil
}

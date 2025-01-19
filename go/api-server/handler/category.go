package handler

//
//// // 博客分类
//type CategoryServer struct {
//	proto.UnimplementedBlogServer
//}
//
//func (c *CategoryServer) CreateCategory(ctx context.Context, req *proto.CreateCategoryInfo) (*proto.CategoryInfoResponse, error) {
//	var cate model.Category
//	result := global.DB.Where(&model.Category{Name: req.Name}).First(&cate)
//	if result.RowsAffected == 1 {
//		return nil, status.Error(codes.AlreadyExists, "类别已经存在")
//	}
//	cate.Name = req.Name
//	cate.ID = req.Id
//	result = global.DB.Create(&cate)
//	if result.Error != nil {
//		return nil, status.Errorf(codes.Internal, result.Error.Error())
//	}
//	return &proto.CategoryInfoResponse{
//		Susee: true,
//	}, nil
//}
//
//func (c *CategoryServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryInfo) (*proto.CategoryInfoResponse, error) {
//	result := global.DB.Delete(&req.Id)
//	if result.Error != nil {
//		return nil, status.Errorf(codes.Internal, result.Error.Error())
//	}
//	return nil, nil
//}
//
//func (c *CategoryServer) UpdateCategory(ctx context.Context, req *proto.CreateCategoryInfo) (*proto.CategoryInfoResponse, error) {
//	//var cate model.Category
//	//result := global.DB.Where(&model.Category{BaseModel: model.BaseModel{ID: req.Id}}).First(&cate)
//	//if result.RowsAffected == 0 {
//	//	return nil, status.Error(codes.AlreadyExists, "类别不存在")
//	//}
//	//cate.Name = req.Name
//	//result.Save(&cate)
//	//if result != nil {
//	//	return nil, status.Errorf(codes.Internal, result.Error.Error())
//	//}
//	//return nil, nil
//
//}
//func (c *CategoryServer) GetCategoryDetail(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponses, error) {
//
//	var cate []model.Category
//	fmt.Println("aa1")
//	result := global.DB.Find(&cate)
//	fmt.Println("aa2")
//	if result != nil {
//		fmt.Println("aa3")
//		return nil, result.Error
//	}
//	rsp := &proto.CategoryInfoResponses{}
//	rsp.Total = int32(result.RowsAffected)
//	for _, blo := range cate {
//		cateinfo := proto.CateGoryFilterRequest{
//			Id:   blo.ID,
//			Name: blo.Name,
//		}
//
//		rsp.Data = append(rsp.Data, &cateinfo)
//	}
//	return rsp, nil
//}

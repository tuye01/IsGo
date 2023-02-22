package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/repository"
	"ginEssential/response"
	"ginEssential/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	//var requestCategory model.Category
	//ctx.BindJSON(&requestCategory)
	//if requestCategory.Name == "" {
	//	response.Fail(ctx, "数据验证错误，分类必填", nil)
	//	return
	//}
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类必填", nil)
		return
	}

	if _, err := c.Repository.Create(requestCategory.Name); err != nil {
		response.Fail(ctx, "创建失败", nil)
		return
	}

	response.Success(ctx, nil, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//parm 穿id body穿对象
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类必填", nil)
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}
	//更新分类
	//map
	//struct
	//name value
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		//response.Fail(ctx, "更新失败", nil)
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "id不存在", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, "删除失败", nil)
		return
	}

	response.Success(ctx, nil, "删除成功")
}

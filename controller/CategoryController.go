package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
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

	category := model.Category{Name: requestCategory.Name}
	c.DB.Create(&category)

	response.Success(ctx, nil, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//parm 穿id body穿对象
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类必填", nil)
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category

	if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}
	//更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category

	if err := c.DB.First(&category, categoryId).Error; err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "删除失败", nil)
		return
	}

	response.Success(ctx, nil, "删除成功")
}

package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
	"sync"
	"time"
)

type CategoryServiceImpl struct {
	wg sync.WaitGroup
}

func (receiver *CategoryServiceImpl) DeleteCategory(ids *pb.CsDeleteCategory) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		_, err := model.DeleteCategory(tx, "id in (?)", ids.CategoryIdList)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (receiver *CategoryServiceImpl) AddOrUpdateCategory(category *pb.CsCategory) error {
	if category.Id == 0 {
		//新增
		err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
			err := model.AddCategory(tx, &model.Category{
				CategoryName: category.GetCategoryName(),
				CreateTime:   time.Now().Unix(),
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		err := model.UpdateCategory(tx, &model.Category{
			ID:           int(category.Id),
			CategoryName: category.CategoryName,
			UpdateTime:   time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (receiver *CategoryServiceImpl) GetAdminCategory(csCondition *pb.CsCondition) (*pb.ScAdminCategories, error) {
	categories, err := model.GetCategoriesByConditionWithPage(csCondition.GetKeywords(), &page.IPage{Current: int(csCondition.GetCurrent()), Size: int(csCondition.GetSize())}, "%"+csCondition.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var categorySlice []*pb.Category
	for _, category := range categories {
		categorySlice = append(categorySlice, &pb.Category{
			Id:           int32(category.ID),
			CategoryName: category.CategoryName,
			CreateTIme:   category.CreateTime,
		})
	}
	categoryCount, err := model.GetCategoriesCountByCondition(csCondition.GetKeywords(), "%"+csCondition.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	return &pb.ScAdminCategories{
		CategoryList: categorySlice,
		Count:        categoryCount,
	}, nil
}

func NewCategoryServiceImpl() *CategoryServiceImpl {
	return &CategoryServiceImpl{}
}

func (receiver *CategoryServiceImpl) GetCategories() ([]*pb.Category, error) {
	categories, err := model.GetCategories("1 = 1")
	if err != nil {
		return nil, err
	}
	var categorySlice []*pb.Category
	for _, category := range categories {
		count, err := model.GetArticlesCountByCondition("category_id = ? AND is_delete = 0 AND is_publish = 1", category.ID)
		logging.Logger.Debug(count)
		if err != nil {
			return nil, err
		}
		c := &pb.Category{
			Id:           int32(category.ID),
			CategoryName: category.CategoryName,
			ArticleCount: int32(count),
		}
		categorySlice = append(categorySlice, c)
	}
	return categorySlice, err
}

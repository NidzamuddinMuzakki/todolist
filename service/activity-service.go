package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/repository"
)

type ActivityService interface {
	Insert(ctx context.Context, activity entity.ActivityEntity) interface{}
	Update(ctx context.Context, activity entity.ActivityEntity) interface{}

	FindById(ctx context.Context, byId int) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string) interface{}
	Delete(ctx context.Context, id int) interface{}
}

type ActivityServiceImpl struct {
	ActivityRepository repository.ActivityRepository
	DB                 *sql.DB
}

func NewActivityService(activityRepo repository.ActivityRepository, DB *sql.DB) ActivityService {
	return &ActivityServiceImpl{
		ActivityRepository: activityRepo,
		DB:                 DB,
	}
}

func (service *ActivityServiceImpl) Delete(ctx context.Context, id int) interface{} {
	ss := service.ActivityRepository.Delete(ctx, service.DB, id)
	if ss == "gagal" {
		return "not found"
	}
	datas := make(map[string]interface{})
	return datas
}
func (service *ActivityServiceImpl) Update(ctx context.Context, activity entity.ActivityEntity) interface{} {

	data := service.ActivityRepository.FindById(ctx, service.DB, activity.RowId)
	// fmt.Println(data)
	if len(data) == 0 {
		return "not found"
	}
	data[0].Title = activity.Title
	update := service.ActivityRepository.Update(ctx, service.DB, data[0])
	return update
}
func (service *ActivityServiceImpl) Insert(ctx context.Context, activity entity.ActivityEntity) interface{} {

	insertData := service.ActivityRepository.Insert(ctx, service.DB, activity)
	return insertData
}
func (service *ActivityServiceImpl) FindById(ctx context.Context, byId int) interface{} {

	getData := service.ActivityRepository.FindById(ctx, service.DB, byId)
	if len(getData) == 0 {
		return nil
	}
	return getData[0]
}

func (service *ActivityServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string) interface{} {

	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "id asc"
	}

	where := fmt.Sprintf("%s ", filter)
	getData := service.ActivityRepository.FindAll(ctx, service.DB, where)
	if len(getData) == 0 {
		return make([]string, 0)
	}
	return getData
}

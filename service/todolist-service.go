package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/repository"
)

type TodolistService interface {
	Insert(ctx context.Context, activity entity.TodolistEntity) interface{}
	Update(ctx context.Context, activity entity.TodolistEntity, isactive interface{}) interface{}

	FindById(ctx context.Context, byId int) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string, group int) interface{}
	Delete(ctx context.Context, id int) (interface{}, int)
}

type TodolistServiceImpl struct {
	TodolistRepository repository.TodolistRepository
	DB                 *sql.DB
}

func NewTodolistService(activityRepo repository.TodolistRepository, DB *sql.DB) TodolistService {
	return &TodolistServiceImpl{
		TodolistRepository: activityRepo,
		DB:                 DB,
	}
}

func (service *TodolistServiceImpl) Delete(ctx context.Context, id int) (interface{}, int) {

	data := service.TodolistRepository.FindById(ctx, service.DB, id)
	ss := service.TodolistRepository.Delete(ctx, service.DB, id)
	if ss == "gagal" || len(data) == 0 {
		return "not found", 0
	}
	datas := make(map[string]interface{})
	return datas, data[0].ActivityGroupId
}
func (service *TodolistServiceImpl) Update(ctx context.Context, activity entity.TodolistEntity, isactive interface{}) interface{} {

	data := service.TodolistRepository.FindById(ctx, service.DB, activity.RowId)
	// fmt.Println(data)
	if len(data) == 0 {
		return "not found"
	}
	if activity.Title != "" {
		data[0].Title = activity.Title
	}
	// fmt.Println(activity.Prority, "n")
	if activity.Prority != "" {
		data[0].Prority = activity.Prority
	}
	if isactive != nil {
		data[0].IsActive = activity.IsActive

	}

	update := service.TodolistRepository.Update(ctx, service.DB, data[0])
	return update
}
func (service *TodolistServiceImpl) Insert(ctx context.Context, activity entity.TodolistEntity) interface{} {

	if activity.Prority == "" {
		activity.Prority = "very-high"
	}
	activity.Status = "ok"
	activity.IsActive = true
	insertData := service.TodolistRepository.Insert(ctx, service.DB, activity)
	return insertData
}
func (service *TodolistServiceImpl) FindById(ctx context.Context, byId int) interface{} {

	getData := service.TodolistRepository.FindById(ctx, service.DB, byId)
	if len(getData) == 0 {
		return nil
	}
	return getData[0]
}

func (service *TodolistServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string, group int) interface{} {

	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "id asc"
	}
	if group != 0 {
		filter += fmt.Sprintf(" and activity_group_id=%d", group)
	}

	where := fmt.Sprintf("%s ", filter)
	getData := service.TodolistRepository.FindAll(ctx, service.DB, where)
	if len(getData) == 0 {
		return make([]string, 0)
	}
	return getData
}

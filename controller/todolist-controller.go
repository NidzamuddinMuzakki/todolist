package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/helper"
	"github.com/ReneKroon/ttlcache"

	"github.com/NidzamuddinMuzakki/todolist/service"

	"github.com/labstack/echo/v4"
)

type TodolistController interface {
	Insert(ctx echo.Context, cache *ttlcache.Cache)
	Delete(ctx echo.Context, cache *ttlcache.Cache)
	Update(ctx echo.Context, cache *ttlcache.Cache)
	FindAll(ctx echo.Context, cache *ttlcache.Cache)
	FindById(ctx echo.Context, cache *ttlcache.Cache)
}

type TodolistControllerImpl struct {
	TodolistService service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) TodolistController {
	return &TodolistControllerImpl{
		TodolistService: todolistService,
	}
}

func (controller *TodolistControllerImpl) FindAll(ctx echo.Context, cache *ttlcache.Cache) {
	getall := entity.ReqList{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)
	key := fmt.Sprintf("todos-%d", getall.ActivityGroupId)
	activities, errNot := cache.Get(key)
	if !errNot {
		activities = controller.TodolistService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order, getall.ActivityGroupId)
		go cache.SetWithTTL(key, activities, time.Hour)
	}

	// fmt.Println(resultData)
	status := "Success"
	message := "Success"

	webResponse := entity.WebResponseListAndDetail{
		Message: message,
		Status:  status,
		Data:    activities,
	}

	helper.WriteToResponseBody(ctx, webResponse, 200)
}

func (controller *TodolistControllerImpl) FindById(ctx echo.Context, cache *ttlcache.Cache) {

	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	key := fmt.Sprintf("todo-id-%d", id)
	activities, errNot := cache.Get(key)
	if !errNot {
		activities = controller.TodolistService.FindById(ctx.Request().Context(), id)
		go cache.SetWithTTL(key, activities, time.Hour)
	}

	// fmt.Println(resultData)
	status := "Success"
	message := "Success"
	code := 200
	var webResponse interface{}
	webResponse = entity.WebResponseListAndDetail{
		Message: message,
		Status:  status,
		Data:    activities,
	}
	if activities == nil {
		status = "Not Found"
		code = 404
		message = fmt.Sprintf("Todo with ID %d Not Found", id)
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *TodolistControllerImpl) Insert(ctx echo.Context, cache *ttlcache.Cache) {
	create := entity.TodolistEntity{}
	err := ctx.Bind(&create)
	helper.PanicIfError(err)

	var resultData interface{}
	status := "Success"
	message := "Success"
	code := 201
	var webResponse interface{}
	if create.Title != "" && create.ActivityGroupId != 0 {
		resultData = controller.TodolistService.Insert(ctx.Request().Context(), create)
		key := fmt.Sprintf("todo-id-%d", resultData.(entity.TodolistEntity).RowId)
		go cache.SetWithTTL(key, resultData, time.Hour)
		go cache.Remove(fmt.Sprintf("todos-%d", create.ActivityGroupId))
		webResponse = entity.WebResponseListAndDetail{
			Message: message,
			Status:  status,
			Data:    resultData,
		}
	} else {
		code = 400
		status = "Bad Request"
		message = "title cannot be null"
		if create.ActivityGroupId == 0 {
			message = "activity_group_id cannot be null"
		}
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *TodolistControllerImpl) Update(ctx echo.Context, cache *ttlcache.Cache) {
	json_map := make(map[string]interface{})
	errs := json.NewDecoder(ctx.Request().Body).Decode(&json_map)
	helper.PanicIfError(errs)

	create := entity.TodolistEntity{}
	if json_map["activity_group_id"] != nil {
		create.ActivityGroupId = json_map["activity_group_id"].(int)

	}
	if json_map["is_active"] != nil {
		create.IsActive = json_map["is_active"].(bool)

	}
	if json_map["priority"] != nil {

		create.Prority = json_map["priority"].(string)
	}
	if json_map["status"] != nil {
		create.Status = json_map["status"].(string)

	}
	if json_map["title"] != nil {
		create.Title = json_map["title"].(string)

	}
	// fmt.Println(json_map["is_active"], "cek")

	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	var resultData interface{}
	status := "Success"
	message := "Success"
	var webResponse interface{}
	code := 200
	// fmt.Println(create.Status)
	if create.Title != "" || json_map["is_active"] != nil {

		create.RowId = id

		resultData = controller.TodolistService.Update(ctx.Request().Context(), create, json_map["is_active"])

		if resultData == "not found" {
			status = "Not Found"
			code = 404
			message = fmt.Sprintf("Todo with ID %d Not Found", id)
			webResponse = entity.WebResponseListAndDetailNotFound{
				Message: message,
				Status:  status,
			}
		} else {
			key := fmt.Sprintf("todo-id-%d", resultData.(entity.TodolistEntity).RowId)
			go cache.SetWithTTL(key, resultData, time.Hour)
			go cache.Remove(fmt.Sprintf("todos-%d", resultData.(entity.TodolistEntity).ActivityGroupId))
			webResponse = entity.WebResponseListAndDetail{
				Message: message,
				Status:  status,
				Data:    resultData,
			}

		}

	} else {
		code = 400
		status = "Bad Request"
		message = "title cannot be null"
		// fmt.Println(400, create, "nidzam")
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *TodolistControllerImpl) Delete(ctx echo.Context, cache *ttlcache.Cache) {
	code := 200
	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	var resultData interface{}
	status := "Success"
	message := "Success"
	var webResponse interface{}

	resultData, grup := controller.TodolistService.Delete(ctx.Request().Context(), id)
	if resultData == "not found" {
		status = "Not Found"
		code = 404
		message = fmt.Sprintf("Todo with ID %d Not Found", id)
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	} else {
		key := fmt.Sprintf("todo-id-%d", id)
		go cache.Remove(key)
		go cache.Remove(fmt.Sprintf("todos-%d", grup))
		webResponse = entity.WebResponseListAndDetail{
			Message: message,
			Status:  status,
			Data:    resultData,
		}

	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

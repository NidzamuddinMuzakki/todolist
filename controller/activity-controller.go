package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/helper"
	"github.com/NidzamuddinMuzakki/todolist/service"
	"github.com/ReneKroon/ttlcache"
	"github.com/labstack/echo/v4"
)

type ActivityController interface {
	Insert(ctx echo.Context, cache *ttlcache.Cache)
	Delete(ctx echo.Context, cache *ttlcache.Cache)
	Update(ctx echo.Context, cache *ttlcache.Cache)
	FindAll(ctx echo.Context, cache *ttlcache.Cache)
	FindById(ctx echo.Context, cache *ttlcache.Cache)
}

type ActivityControllerImpl struct {
	ActivityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		ActivityService: activityService,
	}
}

func (controller *ActivityControllerImpl) FindAll(ctx echo.Context, cache *ttlcache.Cache) {
	key := "activities"

	getall := entity.ReqList{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)
	activities, errNot := cache.Get(key)
	if !errNot {
		activities = controller.ActivityService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order)
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

func (controller *ActivityControllerImpl) FindById(ctx echo.Context, cache *ttlcache.Cache) {

	// getall := entity.ReqList{}
	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	key := fmt.Sprintf("activity-id-%d", id)
	activities, errNot := cache.Get(key)
	if !errNot {
		activities = controller.ActivityService.FindById(ctx.Request().Context(), id)
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
		code = 404
		status = "Not Found"
		message = fmt.Sprintf("Activity with ID %d Not Found", id)
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *ActivityControllerImpl) Insert(ctx echo.Context, cache *ttlcache.Cache) {
	create := entity.ActivityEntity{}
	err := ctx.Bind(&create)
	helper.PanicIfError(err)

	var resultData interface{}
	status := "Success"
	message := "Success"
	code := 201
	var webResponse interface{}
	if create.Title != "" {

		resultData = controller.ActivityService.Insert(ctx.Request().Context(), create)
		key := fmt.Sprintf("activity-id-%d", resultData.(entity.ActivityEntity).RowId)
		go cache.SetWithTTL(key, resultData, time.Hour)
		go cache.Remove("activities")
		webResponse = entity.WebResponseListAndDetail{
			Message: message,
			Status:  status,
			Data:    resultData,
		}
	} else {
		code = 400
		status = "Bad Request"
		message = "title cannot be null"
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *ActivityControllerImpl) Update(ctx echo.Context, cache *ttlcache.Cache) {
	create := entity.ActivityEntity{}
	err := ctx.Bind(&create)
	helper.PanicIfError(err)
	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	var resultData interface{}
	status := "Success"
	message := "Success"
	code := 200
	var webResponse interface{}
	if create.Title != "" {
		create.RowId = id
		resultData = controller.ActivityService.Update(ctx.Request().Context(), create)
		if resultData == "not found" {
			code = 404
			status = "Not Found"
			message = fmt.Sprintf("Activity with ID %d Not Found", id)
			webResponse = entity.WebResponseListAndDetailNotFound{
				Message: message,
				Status:  status,
			}
		} else {
			key := fmt.Sprintf("activity-id-%d", resultData.(entity.ActivityEntity).RowId)
			go cache.SetWithTTL(key, resultData, time.Hour)
			go cache.Remove("activities")
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
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

func (controller *ActivityControllerImpl) Delete(ctx echo.Context, cache *ttlcache.Cache) {

	getId := ctx.ParamValues()[0]
	id, err := strconv.Atoi(getId)
	helper.PanicIfError(err)
	var resultData interface{}
	status := "Success"
	message := "Success"
	var webResponse interface{}
	code := 200
	resultData = controller.ActivityService.Delete(ctx.Request().Context(), id)
	if resultData == "not found" {
		status = "Not Found"
		code = 404
		message = fmt.Sprintf("Activity with ID %d Not Found", id)
		webResponse = entity.WebResponseListAndDetailNotFound{
			Message: message,
			Status:  status,
		}
	} else {
		key := fmt.Sprintf("activity-id-%d", id)
		go cache.Remove(key)
		go cache.Remove("activities")
		webResponse = entity.WebResponseListAndDetail{
			Message: message,
			Status:  status,
			Data:    resultData,
		}

	}

	helper.WriteToResponseBody(ctx, webResponse, code)
}

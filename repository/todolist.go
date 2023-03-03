package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/helper"
)

type TodolistRepository interface {
	Insert(ctx context.Context, tx *sql.DB, user entity.TodolistEntity) interface{}
	Update(ctx context.Context, tx *sql.DB, user entity.TodolistEntity) interface{}
	FindById(ctx context.Context, tx *sql.DB, byID int) []entity.TodolistEntity
	FindAll(ctx context.Context, tx *sql.DB, where string) []entity.TodolistEntity
	Delete(ctx context.Context, tx *sql.DB, byID int) string
}

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() TodolistRepository {
	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.DB, byId int) string {

	SQL := fmt.Sprintf("delete from todos  where id='%d'", byId)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {

		return "berhasil"
	} else {
		return "gagal"
	}

}
func (repository *TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.DB, activity entity.TodolistEntity) interface{} {
	waktuNow := helper.TimePlus7(time.Now())
	SQL := fmt.Sprintf("update todos set title='%s',priority='%s',is_active=%t,updatedAt='%s' where id=%d", activity.Title, activity.Prority, activity.IsActive, waktuNow, activity.RowId)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {
		activity.UpdatedAt = waktuNow
		activity.Status = "ok"
		return activity
	} else {
		return nil
	}

}
func (repository *TodolistRepositoryImpl) Insert(ctx context.Context, tx *sql.DB, activity entity.TodolistEntity) interface{} {
	waktuNow := helper.TimePlus7(time.Now())
	SQL := fmt.Sprintf("insert into todos (title,activity_group_id,is_active,priority,createdAt,updatedAt) values ('%s',%d,%t,'%s','%s','%s') ", activity.Title, activity.ActivityGroupId, activity.IsActive, activity.Prority, waktuNow, waktuNow)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {
		id, err := row.LastInsertId()
		helper.PanicIfError(err)
		activity.RowId = int(id)

		activity.CreatedAt = waktuNow
		activity.UpdatedAt = waktuNow
		return activity
	} else {
		return nil
	}

}

func (repository *TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.DB, byID int) []entity.TodolistEntity {
	SQL := fmt.Sprintf("select id,activity_group_id,title,is_active,priority, createdAt, updatedAt from todos where id=%d ", byID)
	var datas []entity.TodolistEntity
	var data entity.TodolistEntity
	row, err := tx.QueryContext(ctx, SQL)
	// fmt.Println(err, byID)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.ActivityGroupId, &data.Title, &data.IsActive, &data.Prority, &data.CreatedAt, &data.UpdatedAt)
		helper.PanicIfError(err)
		datas = append(datas, data)
	}
	return datas
}

func (repository *TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.DB, where string) []entity.TodolistEntity {
	SQL := fmt.Sprintf("select id,activity_group_id,title,is_active,priority, createdAt, updatedAt from todos where %s", where)
	// fmt.Println(SQL)
	var datas []entity.TodolistEntity
	var data entity.TodolistEntity
	row, err := tx.QueryContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.ActivityGroupId, &data.Title, &data.IsActive, &data.Prority, &data.CreatedAt, &data.UpdatedAt)
		// fmt.Println(err)

		helper.PanicIfError(err)
		// data.CreatedTime = helper.ConvertDateTime(data.CreatedTime)
		// data.UpdatedTime = helper.ConvertDateTime(data.UpdatedTime)

		datas = append(datas, data)
	}
	return datas
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/helper"
)

type ActivityRepository interface {
	Insert(ctx context.Context, tx *sql.DB, user entity.ActivityEntity) interface{}
	Update(ctx context.Context, tx *sql.DB, user entity.ActivityEntity) interface{}
	FindById(ctx context.Context, tx *sql.DB, byID int) []entity.ActivityEntity

	FindAll(ctx context.Context, tx *sql.DB, where string) []entity.ActivityEntity
	Delete(ctx context.Context, tx *sql.DB, byID int) string
}

type ActivityRepositoryImpl struct {
}

func NewActivityRepository() ActivityRepository {
	return &ActivityRepositoryImpl{}
}

func (repository *ActivityRepositoryImpl) Delete(ctx context.Context, tx *sql.DB, byId int) string {

	SQL := fmt.Sprintf("delete from activities  where id='%d'", byId)
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
func (repository *ActivityRepositoryImpl) Update(ctx context.Context, tx *sql.DB, activity entity.ActivityEntity) interface{} {
	waktuNow := helper.TimePlus7(time.Now())
	SQL := fmt.Sprintf("update activities set title='%s',email='%s',updatedAt='%s' where id='%d'", activity.Title, activity.Email, waktuNow, activity.RowId)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {
		activity.UpdatedAt = waktuNow
		return activity
	} else {
		return nil
	}

}
func (repository *ActivityRepositoryImpl) Insert(ctx context.Context, tx *sql.DB, activity entity.ActivityEntity) interface{} {
	waktuNow := helper.TimePlus7(time.Now())
	SQL := fmt.Sprintf("insert into activities (title,email,createdAt,updatedAt) values ('%s','%s','%s','%s') ", activity.Title, activity.Email, waktuNow, waktuNow)
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

func (repository *ActivityRepositoryImpl) FindById(ctx context.Context, tx *sql.DB, byID int) []entity.ActivityEntity {
	SQL := fmt.Sprintf("select id,title,email, createdAt, updatedAt from activities where id=%d", byID)
	var datas []entity.ActivityEntity
	var data entity.ActivityEntity
	row, err := tx.QueryContext(ctx, SQL)
	// fmt.Println(err, byID)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Title, &data.Email, &data.CreatedAt, &data.UpdatedAt)
		helper.PanicIfError(err)
		datas = append(datas, data)
	}
	return datas
}

func (repository *ActivityRepositoryImpl) FindAll(ctx context.Context, tx *sql.DB, where string) []entity.ActivityEntity {
	SQL := fmt.Sprintf("select id,title,email, createdAt, updatedAt from activities where %s", where)
	// fmt.Println(SQL)
	var datas []entity.ActivityEntity
	var data entity.ActivityEntity
	row, err := tx.QueryContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Title, &data.Email, &data.CreatedAt, &data.UpdatedAt)
		// fmt.Println(err)

		helper.PanicIfError(err)
		// data.CreatedTime = helper.ConvertDateTime(data.CreatedTime)
		// data.UpdatedTime = helper.ConvertDateTime(data.UpdatedTime)

		datas = append(datas, data)
	}
	return datas
}

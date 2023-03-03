package entity

type TodolistEntity struct {
	RowId           int    `json:"id"`
	ActivityGroupId int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	Prority         string `json:"priority"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type TodolistEntityss struct {
	RowId           int    `json:"id"`
	ActivityGroupId int    `json:"activity_group_id"`
	IsActive        string `json:"is_active"`
	Title           string `json:"title"`
	Status          string `json:"status"`
	Prority         string `json:"priority"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}
type TodolistEntityS struct {
	RowId           int    `json:"id"`
	ActivityGroupId int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Title           string `json:"title"`
	Prority         string `json:"priority"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type ActivityEntity struct {
	RowId     int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

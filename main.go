package main

import (
	"database/sql"

	"github.com/NidzamuddinMuzakki/todolist/app"
	"github.com/NidzamuddinMuzakki/todolist/controller"
	"github.com/NidzamuddinMuzakki/todolist/repository"
	"github.com/NidzamuddinMuzakki/todolist/service"
	"github.com/ReneKroon/ttlcache"
)

var (
	db                 *sql.DB                       = app.Init()
	ActivityRepository repository.ActivityRepository = repository.NewActivityRepository()
	ActivityService    service.ActivityService       = service.NewActivityService(ActivityRepository, db)
	ActivityController controller.ActivityController = controller.NewActivityController(ActivityService)

	TodolistRepository repository.TodolistRepository = repository.NewTodolistRepository()
	TodolistService    service.TodolistService       = service.NewTodolistService(TodolistRepository, db)
	TodolistController controller.TodolistController = controller.NewTodolistController(TodolistService)
)

func main() {
	defer db.Close()
	db.Query("CREATE TABLE `todos` (`id` int(11) NOT NULL,`activity_group_id` int(11) NOT NULL,`title` varchar(200) NOT NULL,`is_active` tinyint(1) NOT NULL,`priority` varchar(100) NOT NULL,`createdAt` timestamp NULL DEFAULT NULL,`updatedAt` timestamp NULL DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")
	db.Query("CREATE TABLE `activities` (`id` int(11) NOT NULL,`title` varchar(200) NOT NULL,`email` varchar(200) NOT NULL,`createdAt` timestamp NULL DEFAULT NULL,`updatedAt` timestamp NULL DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")
	db.Query("ALTER TABLE `activities` ADD PRIMARY KEY (`id`);")
	db.Query("ALTER TABLE `todos` ADD PRIMARY KEY (`id`);")
	db.Query("ALTER TABLE `activities` MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;")
	db.Query("ALTER TABLE `todos` MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;")
	cache := ttlcache.NewCache()
	r := app.InitRouter(ActivityController, TodolistController, cache)
	r.Start(":3030")

}

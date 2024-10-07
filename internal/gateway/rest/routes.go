package rest

import (
	"codechallenge/config"
	"codechallenge/internal/repository/database"
	"codechallenge/internal/repository/queue"
	"codechallenge/internal/repository/storage"
	"codechallenge/internal/service"
	"codechallenge/utils"
	"github.com/labstack/echo/v4"
)

func registerRoutes(app *echo.Echo) {
	// ---------------- Dependency   ----------------
	db, err := utils.PostgresConnection(
		config.AppConfig.Databases.Postgres.Host,
		config.AppConfig.Databases.Postgres.Port,
		config.AppConfig.Databases.Postgres.User,
		config.AppConfig.Databases.Postgres.Pass,
		config.AppConfig.Databases.Postgres.DatabaseName,
		config.AppConfig.Databases.Postgres.SslMode,
		config.AppConfig.Databases.Postgres.MaxOpenConns,
		config.AppConfig.Databases.Postgres.MaxIdleConns,
		config.AppConfig.Databases.Postgres.Timeout)
	if err != nil {
		panic(err)
	}
	queueClient, err := utils.LoadSQS(config.AppConfig.Queue.Address)
	if err != nil {
		panic(err)
	}
	storageClient, err := utils.LoadS3(config.AppConfig.Storage.Endpoint)
	if err != nil {
		panic(err)
	}

	// ---------------- Repositories ----------------
	todoRepository := database.NewTodoRepository(db, db)
	queueRepository := queue.NewQueue(config.AppConfig.Queue.QueueAddress, queueClient)
	storageRepository := storage.NewStorageRepository(storageClient)

	// ---------------- Services     ----------------
	queueService := service.NewQueueService(queueRepository)
	storageService := service.NewStorageService(storageRepository, config.AppConfig.Opt.MaxUploadSize, config.AppConfig.Opt.ValidFileExtensions)
	todoService := service.NewTodoService(queueService, todoRepository)

	// ---------------- Controllers  ----------------
	todoController := NewTodoController(todoService)
	storageController := NewStorageController(storageService)

	app.GET("/", func(c echo.Context) error {
		return c.JSON(200, utils.StandardHttpResponse{
			Message: "",
			Status:  200,
			Data:    nil,
		})
	})

	v1 := app.Group("/v1")

	// TODO: We need authentication over here, but not mentioned in task description
	v1.POST("/todo", todoController.AddItem)
	v1.POST("/upload", storageController.UploadHandler)
}

package main

import (
	"context"
	"log"
	"os"

	"task_itv/config"
	"task_itv/database"
	_ "task_itv/docs" // Импорт Swagger документации
	"task_itv/handlers"
	"task_itv/middleware"
	"task_itv/models"
	"task_itv/routes"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	// Создаем логгер
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}

	// Инициализируем приложение с UberFx
	app := fx.New(
		// Предоставляем зависимости
		fx.Provide(
			// Конфигурация
			config.LoadConfig,
			// База данных
			database.NewDatabase,
			// Модели и репозитории
			models.NewMovieRepository,
			// Middleware
			middleware.NewAuthMiddleware,
			// Обработчики
			handlers.NewMovieHandler,
			// Роутер
			routes.NewRouter,
			// Логгер
			func() *zap.Logger { return logger },
		),
		// Запускаем приложение
		fx.Invoke(
			// Миграция базы данных
			func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Movie{})
			},
			// Настройка маршрутов
			func(router *routes.Router) {
				router.SetupRoutes()
			},
		),
		// Настраиваем логирование
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	// Запускаем приложение
	if err := app.Start(context.Background()); err != nil {
		logger.Fatal("Failed to start application", zap.Error(err))
		os.Exit(1)
	}

	// Ждем сигнала завершения
	<-app.Done()
}

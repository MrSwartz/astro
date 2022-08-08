package main

import (
	"astro/pkg/handler"
	repo "astro/pkg/repository"
	srvc "astro/pkg/service"
	"astro/pkg/tools"
	"context"
	"os"
	"os/signal"
	"syscall"

	"astro"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cnf := repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := repo.NewPostgresDB(cnf)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	handlers := handler.NewHandler(srvc.NewService(repo.NewRepository(db)))

	srv := new(astro.Server)
	go func() {
		if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	uuid := tools.NewUUID()
	logrus.Printf("service astro Started, UUID: %s", uuid)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("astro Shutting Down, UUID: %s", uuid)

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s, UUID: %s", err.Error(), uuid)
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s, UUID: %s", err.Error(), uuid)
	}
}

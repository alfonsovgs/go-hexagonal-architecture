package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/alfonsovgs/go-hexagonal-architecture/internal"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/creating"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/increasing"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/bus/inmemory"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/server"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

const (
	host = "localhost"
	port = 8080

	dbUser          = "codely"
	dbPass          = "codely"
	dbHost          = "localhost"
	dbPort          = "3306"
	dbName          = "codely"
	shutdownTimeout = time.Second * 10
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	fmt.Println("Data:", cfg)

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, shutdownTimeout)
	courseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(courseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`

	// Database configuration
	DbUser    string        `default:"codely"`
	DbPass    string        `default:"codely"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"codely"`
	DbTimeout time.Duration `default:"5s"`
}

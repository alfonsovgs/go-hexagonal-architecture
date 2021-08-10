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
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
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

package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/alfonsovgs/go-hexagonal-architecture/internal/creating"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/bus/inmemory"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/server"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)
	courseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(courseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}

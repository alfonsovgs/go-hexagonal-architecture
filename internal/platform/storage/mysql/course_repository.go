package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	mooc "github.com/alfonsovgs/go-hexagonal-architecture/internal"
	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID(),
		Name:     course.Name(),
		Duration: course.Duration(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return errors.New("")
}

func (r *CourseRepository) GetAll(ctx context.Context) []mooc.Course {
	return make([]mooc.Course, 0)
}

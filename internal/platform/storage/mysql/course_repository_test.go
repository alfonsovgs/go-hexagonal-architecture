package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/alfonsovgs/go-hexagonal-architecture/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const Duration = time.Second * 10

func Test_Course_Repository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "0866e3be-340b-4abc-97ca-326c79ae4716", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db, Duration)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Succeed(t *testing.T) {
	courseID, courseName, courseDuration := "0866e3be-340b-4abc-97ca-326c79ae4716", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db, Duration)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)

}

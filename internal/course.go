package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// CourseRepository defines the expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	GetAll(ctx context.Context) []Course
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)

	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id.String()
}

// Name returns the coruse name,
func (c Course) Name() string {
	return c.name.String()
}

// Duration returns the course duration.
func (c Course) Duration() string {
	return c.duration.String()
}

var ErrInvalidCourseId = errors.New("invalid Course ID")
var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")
var ErrEmptyDuration = errors.New("the field Duration can not be empty")

// CourseID represernts the course unique identifier
type CourseID struct {
	value string
}

//NewCourseId instantiate the VO for the CourseId
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseId, value)

	}

	return CourseID{
		value: v.String(),
	}, nil
}

func (id CourseID) String() string {
	return id.value
}

type CourseName struct {
	value string
}

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

func (c CourseName) String() string {
	return c.value
}

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	return CourseDuration{
		value: value,
	}, nil
}

func (c CourseDuration) String() string {
	return c.value
}

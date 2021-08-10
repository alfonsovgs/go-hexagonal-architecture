package creating

import (
	"context"
	"errors"

	mooc "github.com/alfonsovgs/go-hexagonal-architecture/internal"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/increasing"
	"github.com/alfonsovgs/go-hexagonal-architecture/kit/event"
)

type IncreaseCoursesCounternCourseCreated struct {
	increaserService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterService) IncreaseCoursesCounternCourseCreated {
	return IncreaseCoursesCounternCourseCreated{
		increaserService: increaserService,
	}
}

func (e IncreaseCoursesCounternCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)

	if !ok {
		return errors.New("unexpected event")
	}

	return e.increaserService.Increase(courseCreatedEvt.ID())
}

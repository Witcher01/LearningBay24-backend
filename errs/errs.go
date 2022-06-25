package errs

import (
	"errors"
)

var (
	ErrFileExtensionNotAllowed error = errors.New("File extension is not allowed")
	ErrNoFileExtension         error = errors.New("File has no extension")

	ErrNotAdmin           error = errors.New("Admin permission required")
	ErrNotModerator       error = errors.New("Moderator permissions required")
	ErrNotUser            error = errors.New("User permissions required")
	ErrNotCourseAdmin     error = errors.New("Course admin permissions required")
	ErrNotCourseModerator error = errors.New("Course moderator permissions required")
	ErrNotCourseUser      error = errors.New("Course user permissions required")

	ErrParameterConversion error = errors.New("Unable to convert parameter item")
	ErrNoQuery             error = errors.New("Unable to find query parameter")
	ErrRawData             error = errors.New("Unable to get raw data from request")
	ErrNoFileInRequest     error = errors.New("Unable to find file in request")
	ErrBodyConversion      error = errors.New("Unable to convert body")

	// TODO: add to `handleApiError`
	ErrSelfRegisterExam error = errors.New("Cannot register for own exam")
	// TODO: add to `handleApiError`
	ErrRegisterDeadlinePassed error = errors.New("Cannot register to exam past deadline")
)
package failure

const (
	FailureTypeUnknown              FailureCode = "unknown_error"
	FailureTypeInvalidAnswerType    FailureCode = "invalid_answer_type"
	FailureTypeUnauthorized         FailureCode = "unauthorized"
	FailureTypeNotFound             FailureCode = "not_found"
	FailureTypeUnsupportedMediaType FailureCode = "unsupported_media_type"
	FailureTypeValidation           FailureCode = "validation"
	FailureTypeInvalidFormat        FailureCode = "invalid_format"
	FailureTypeValueNotUnique       FailureCode = "value_not_unique"
	FailureTypeFileNotFound         FailureCode = "file_not_found"
	FailureTypeProfileNotFound      FailureCode = "profile_not_found"
	FailureTypeTaskNotFound         FailureCode = "task_not_found"
	FailureTypeUsernameNotUnique    FailureCode = "username_not_unique"
)

package failure

const (
	FailureTypeUnknown             FailureType = "unknown_error"
	FailureTypeNotFound            FailureType = "not_found"
	FailureTypeUnauthorized        FailureType = "unauthorized"
	FailureTypeInvalidInput        FailureType = "invalid_input"
	FailureTypeValueNotUnique      FailureType = "value_not_unique"
	FailureTypeUnsupportedFileType FailureType = "unsupported_file_type"
)

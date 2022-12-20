package errors

//ErrLoadingConfiguringForAwsSecretsManager
type ErrLoadingConfiguringForAwsSecretsManager struct {
	Code    int
	Message string
}

func NewErrLoadingConfiguringForAwsSecretsManager() ErrLoadingConfiguringForAwsSecretsManager {
	return ErrLoadingConfiguringForAwsSecretsManager{
		Code:    500,
		Message: "Error loading configuration for Secrets Manager",
	}
}

func (e ErrLoadingConfiguringForAwsSecretsManager) Error() string {
	return e.Message
}

//
type ErrRegionNotSpecifiedForSecretsManager struct {
	Code    int
	Message string
}

func NewErrRegionNotSpecifiedForSecretsManager() ErrRegionNotSpecifiedForSecretsManager {
	return ErrRegionNotSpecifiedForSecretsManager{
		Code:    500,
		Message: "Region not specified for Secrets Manager",
	}
}

func (e ErrRegionNotSpecifiedForSecretsManager) Error() string {
	return e.Message
}

//
type ErrRetrievingAwsSecretsManagerClient struct {
	Code    int
	Message string
}

func NewErrRetrievingAwsSecretsManagerClient() ErrRetrievingAwsSecretsManagerClient {
	return ErrRetrievingAwsSecretsManagerClient{
		Code:    500,
		Message: "Error getting client for Secrets Manager",
	}
}

func (e ErrRetrievingAwsSecretsManagerClient) Error() string {
	return e.Message
}

type ErrRetrievingSecretFromAwsSecretsManager struct {
	Code    int
	Message string
}

func NewErrRetrievingSecretFromAwsSecretsManager() ErrRetrievingSecretFromAwsSecretsManager {
	return ErrRetrievingSecretFromAwsSecretsManager{
		Code:    500,
		Message: "Error getting secret value from secrets manager",
	}
}

func (e ErrRetrievingSecretFromAwsSecretsManager) Error() string {
	return e.Message
}

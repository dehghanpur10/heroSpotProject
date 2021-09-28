package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetDynamoDBFailOnMissingInfo(t *testing.T) {
	// Arrange
	defer func() {
		AWS_REGION = os.Getenv("AWS_REGION")
		ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
		SECRET_KEY = os.Getenv("SECRET_KEY")

	}()
	AWS_REGION = ""
	ACCESS_TOKEN = ""
	SECRET_KEY = ""
	// Act
	_, err := GetDynamoDB()
	// Assert
	assert.Contains(t, err.Error(), "env")
}
func TestGetDynamoDBSuccess(t *testing.T) {
	// Arrange
	defer func() {
		AWS_REGION = os.Getenv("AWS_REGION")
		ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
		SECRET_KEY = os.Getenv("SECRET_KEY")

	}()
	AWS_REGION = "test"
	ACCESS_TOKEN = "test"
	SECRET_KEY = "test"
	// Act
	_, err := GetDynamoDB()
	fmt.Println(err)
	// Assert
	assert.NoError(t, err)
}

package searchService

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type SearchService struct {
	db dynamodbiface.DynamoDBAPI
}

func New(db dynamodbiface.DynamoDBAPI) *SearchService {
	return &SearchService{
		db: db,
	}
}

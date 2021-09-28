package searchService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *SearchService) GetFacilityWithLatAndLon(lat float64, lon float64) ([]models.Facility, error) {
	filter := expression.Name("latitude").Equal(expression.Value(lat)).And(expression.Name("longitude").Equal(expression.Value(lon)))
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		fmt.Print("searchService.GetFacilityWithLatAndLon - expressionBuild - ")
		return nil, fmt.Errorf("%v - %v", lib.ErrInternal, err)
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(lib.FACILITY_TABLE_NAME),
		IndexName:                 aws.String("Facility_index"),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		KeyConditionExpression:    expr.Filter(),
	}

	result, err := s.db.Query(queryInput)
	if err != nil {
		fmt.Print("searchService.GetFacilityWithLatAndLon - query - ")
		return nil, fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}

	var facilities []models.Facility
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &facilities)
	if err != nil {
		fmt.Print("searchService.GetFacilityWithLatAndLon  - UnmarshalListOfMaps - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return nil, err
	}

	return facilities, nil

}

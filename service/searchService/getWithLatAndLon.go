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
		return nil, fmt.Errorf("searchService.GetFacilityWithLatAndLon - expressionBuild - %v - %v", lib.ErrInternal, err)
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String("FacilitySpot"),
		IndexName:                 aws.String("Facility_index"),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		KeyConditionExpression:    expr.Filter(),
	}

	result, err := s.db.Query(queryInput)
	if err != nil {
		return nil, fmt.Errorf("searchService.GetFacilityWithLatAndLon - query - %w - %v", lib.ErrInternal, err)
	}

	var facilities []models.Facility
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &facilities)
	if err != nil {
		err = fmt.Errorf("searchService.GetFacilityWithLatAndLon  - UnmarshalListOfMaps - %w - %v", lib.ErrInternal, err)
		return nil, err
	}

	return facilities, nil

}

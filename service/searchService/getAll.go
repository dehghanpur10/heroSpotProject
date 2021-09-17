package searchService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *SearchService) GetAllFacility() ([]models.Facility, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("FacilitySpot"),
	}

	result, err := s.db.Scan(scanInput)

	if err != nil {
		return nil, fmt.Errorf("searchService.GetAllFacility - Scan - %v - %v", lib.ErrInternal, err)
	}

	var facilities []models.Facility
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &facilities)
	if err != nil {
		err = fmt.Errorf("searchService.GetAllFacility - UnmarshalListOfMaps - %v - %v", lib.ErrInternal, err)
		return nil, err
	}

	return facilities, nil
}

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
		TableName: aws.String(lib.FACILITY_TABLE_NAME),
	}

	result, err := s.db.Scan(scanInput)

	if err != nil {
		fmt.Print("searchService.GetAllFacility - Scan - ")
		return nil, fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}

	var facilities []models.Facility
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &facilities)
	if err != nil {
		fmt.Print("searchService.GetAllFacility - UnmarshalListOfMaps - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return nil, err
	}

	return facilities, nil
}

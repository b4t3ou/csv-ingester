package database

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/b4t3ou/csv-ingester/storage"
)

const (
	// DefaultRegion representing the default aws region
	DefaultRegion = "eu-west-1"

	// DefaultReadCapacity representing the default aws dynamo read capacity
	DefaultReadCapacity = 10

	// DefaultWriteCapacity representing the default aws dynamo write capacity
	DefaultWriteCapacity = 10
)

// Dynamo table instance for saving the stream position
type Dynamo struct {
	name          string
	region        string
	endpoint      string
	writeCapacity int64
	readCapacity  int64
	session       *dynamodb.DynamoDB
}

// DynamoOption representing the Dynamo constructor options
type DynamoOption func(*Dynamo)

// NewDynamoDB build a new Dynamo db handler
func NewDynamoDB(name string, options ...DynamoOption) (*Dynamo, error) {
	storage := &Dynamo{
		name:          name,
		region:        DefaultRegion,
		writeCapacity: DefaultReadCapacity,
		readCapacity:  DefaultWriteCapacity,
	}

	for _, option := range options {
		option(storage)
	}

	if storage.session == nil {
		s, err := session.NewSession(&aws.Config{
			Region:   aws.String(storage.region),
			Endpoint: aws.String(storage.endpoint),
		})
		if err != nil {
			return nil, err
		}

		storage.session = dynamodb.New(s)
	}

	return storage, nil
}

// WithDynamoRegion setting the region
func WithDynamoRegion(region string) DynamoOption {
	return func(dynamo *Dynamo) { dynamo.region = region }
}

// WithDynamoEndpoint setting the endpoint
func WithDynamoEndpoint(endpoint string) DynamoOption {
	return func(dynamo *Dynamo) { dynamo.endpoint = endpoint }
}

// WithDynamoWriteCapacity setting the write capacity of the table
func WithDynamoWriteCapacity(capacity int64) DynamoOption {
	return func(dynamo *Dynamo) { dynamo.writeCapacity = capacity }
}

// WithDynamoReadCapacity setting the read capacity of the table
func WithDynamoReadCapacity(capacity int64) DynamoOption {
	return func(dynamo *Dynamo) { dynamo.readCapacity = capacity }
}

// WithDynamoSession set the aws dynamo session
func WithDynamoSession(session *dynamodb.DynamoDB) DynamoOption {
	return func(dynamo *Dynamo) { dynamo.session = session }
}

// DeleteTable deleting a table
// WARNING: only uso for testing
func (d *Dynamo) DeleteTable() {
	d.session.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(d.name),
	})
}

// CreateTable is creating a new table and returning when it is ready top use
func (d *Dynamo) CreateTable() error {
	d.session.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(d.name),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("email"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(d.readCapacity),
			WriteCapacityUnits: aws.Int64(d.writeCapacity),
		},
	})

	for {
		output, err := d.session.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String(d.name)})
		if err != nil {
			return err
		}

		if *output.Table.TableStatus == dynamodb.TableStatusActive {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

// Save is saving a record
func (d *Dynamo) Save(item storage.Record) error {
	_, err := d.session.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.name),
		Item: map[string]*dynamodb.AttributeValue{
			"id":           {S: aws.String(item.ID)},
			"name":         {S: aws.String(item.Name)},
			"email":        {S: aws.String(item.Name)},
			"mobileNumber": {N: aws.String(item.MobileNumber)},
		},
	})

	return err
}

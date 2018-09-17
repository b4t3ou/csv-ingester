package ingester

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/b4t3ou/csv-ingester/config"
	pb "github.com/b4t3ou/csv-ingester/proto"
)

// Client representing the client
type Client struct {
	config        *config.Config
	storageClient pb.StorageClient
	reader        ReaderInterface
	items         []*pb.Item
	dataChannel   chan *pb.Item
	doneChannel   chan struct{}
}

// NewClient returning with a new client
func NewClient(cfg *config.Config, reader ReaderInterface) (*Client, error) {
	c := &Client{
		config:      cfg,
		reader:      reader,
		items:       []*pb.Item{},
		dataChannel: make(chan *pb.Item),
		doneChannel: make(chan struct{}),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c.storageClient = pb.NewStorageClient(conn)

	return c, err
}

// Run running the client and the data reader
func (c *Client) Run() {
	go c.readDataChan()

	c.reader.Receive(c.callback)
	close(c.dataChannel)

	<-c.doneChannel
}

func (c *Client) readDataChan() {
	var data []*pb.Item

	for {
		record, open := <-c.dataChannel
		if !open {
			c.send(data)
			c.doneChannel <- struct{}{}
		}

		data = append(data, record)

		if len(data) == 10 {
			c.send(data)
			data = []*pb.Item{}
		}
	}
}

func (c *Client) callback(record Record) {
	c.dataChannel <- &pb.Item{
		Id:           record.ID,
		Name:         record.Name,
		Email:        record.Email,
		MobileNumber: record.MobileNumber,
	}
}

func (c *Client) send(data []*pb.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, err := c.storageClient.Save(ctx, &pb.StorageRequest{
		Items: data,
	})

	return err
}

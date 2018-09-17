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
}

// NewClient returning with a new client
func NewClient(cfg *config.Config, reader ReaderInterface) (*Client, error) {
	c := &Client{
		config: cfg,
		reader: reader,
		items:  []*pb.Item{},
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
	c.reader.Receive(c.callback)
	c.send()
}

func (c *Client) callback(record Record) {
	c.items = append(c.items, &pb.Item{
		Id:           record.ID,
		Name:         record.Name,
		Email:        record.Email,
		MobileNumber: record.MobileNumber,
	})

	if len(c.items) == 10 {
		c.send()
		c.items = []*pb.Item{}
	}
}

func (c *Client) send() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, err := c.storageClient.Save(ctx, &pb.StorageRequest{
		Items: c.items,
	})

	return err
}

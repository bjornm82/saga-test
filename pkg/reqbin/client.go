package reqbin

import (
	"math/rand"
	"time"
)

func New() *Client {
	return &Client{}
}

type Client struct {
	ReturnError error
	ReturnInt   int
}

func (c *Client) Get() (int, error) {
	if c.ReturnError != nil {
		return -0, c.ReturnError
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100), nil
}

func (c *Client) Post(id int) (int, error) {
	if c.ReturnError != nil {
		return -0, c.ReturnError
	}
	return id, nil
}

package adapter2

import "errors"

type Client struct {
	lst map[string]IClient
}

func (c *Client) Traverse() error {
	for _, client := range c.lst {
		if client == nil {
			return errors.New("client is nil")
		}
	}
	return nil
}

func (c *Client) AddClient(str IClient) {
	if c.lst == nil {
		c.lst = make(map[string]IClient)
	}
	c.lst[str.ConvertToStringInterface()] = str
}

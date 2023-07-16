// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IClientHolder interface {
	GetClient() IClient
}

type ClientHolder struct {
	client IClient
}

func (c *ClientHolder) Init(client IClient) *ClientHolder {
	c.client = client
	return c
}

func (c *ClientHolder) GetClient() IClient {
	return c.client
}

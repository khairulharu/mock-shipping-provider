package webhook

import (
	"context"
	"mock-shipping-provider/repository"
)

type Client struct {
	Url string
}

func NewWebHookClient() repository.WebhookClient {
	return &Client{
		Url: "jsjssjsj",
	}
}

func (c *Client) SendStatusUpdate(context.Context, repository.StatusUpdate) error {
	panic("halo")
}

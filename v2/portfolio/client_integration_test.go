package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseIntegrationTestSuite struct {
	suite.Suite
	client    *Client
}

func SetupTest(t *testing.T) *baseIntegrationTestSuite {
	// TODO: get from env
	apiKey := "z2aSOxQ9qPS0RIxS17ns2zN7UjMVoPF7jvTvpUJmnTlTZoj25GcnHdyPQXeQDJk9"
	secretKey := "X1p9TkTnfbPTmYsbO3eJGI5I338XMmiFCF9AE3U5NxbS5dVZE6qZy8RqWXg1dcGF"

	if apiKey == "" || secretKey == "" {
		t.Skip("API key and secret are required for integration tests")
	}

	client := NewClient(apiKey, secretKey)
	client.Debug = true

	return &baseIntegrationTestSuite{
		client: client,
	}
}

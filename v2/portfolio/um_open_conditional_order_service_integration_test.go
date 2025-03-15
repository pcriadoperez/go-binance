package portfolio

import (
	"context"
	"testing"
)

type umOpenConditionalOrderServiceIntegrationTestSuite struct {
	*baseIntegrationTestSuite
}

func TestUMOpenConditionalOrderServiceIntegration(t *testing.T) {
	base := SetupTest(t)
	suite := &umOpenConditionalOrderServiceIntegrationTestSuite{
		baseIntegrationTestSuite: base,
	}

	t.Run("GetOpenConditionalOrder_ByStrategyID", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrderService()
		res, err := service.Symbol("BTCUSDT").
			StrategyID(123445).
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for canceled/triggered/expired orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get open conditional order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.StrategyID != 123445 {
			t.Errorf("Expected strategy ID 123445, got %d", res.StrategyID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
		if res.StrategyStatus != "NEW" {
			t.Errorf("Expected status NEW, got %s", res.StrategyStatus)
		}
	})

	t.Run("GetOpenConditionalOrder_ByClientStrategyID", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrderService()
		res, err := service.Symbol("BTCUSDT").
			NewClientStrategyID("abc").
			Do(context.Background())
		if err != nil {
			// Check if error is "Order does not exist" which is expected for canceled/triggered/expired orders
			if err.Error() != "Order does not exist" {
				t.Fatalf("Failed to get open conditional order: %v", err)
			}
			return
		}

		// Basic validation of returned data
		if res.NewClientStrategyID != "abc" {
			t.Errorf("Expected client strategy ID abc, got %s", res.NewClientStrategyID)
		}
		if res.Symbol != "BTCUSDT" {
			t.Errorf("Expected symbol BTCUSDT, got %s", res.Symbol)
		}
	})

	t.Run("GetOpenConditionalOrder_Error_NoIDs", func(t *testing.T) {
		service := suite.client.NewUMOpenConditionalOrderService()
		_, err := service.Symbol("BTCUSDT").
			Do(context.Background())
		if err == nil {
			t.Fatal("Expected an error when neither strategyId nor newClientStrategyId is provided")
		}

		// Verify it's a Portfolio error
		portfolioErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("Expected Error, got %T", err)
		}
		if portfolioErr.Code != ErrMandatoryParamEmptyOrMalformed {
			t.Errorf("Expected error code %d, got %d", ErrMandatoryParamEmptyOrMalformed, portfolioErr.Code)
		}
	})
}

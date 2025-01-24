package mongoex

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type health struct {
	client *mongo.Client
}

func (m *health) HealthChecks() (name string, ready, live func(ctx context.Context) error) {
	ready = func(ctx context.Context) error {
		ctxPing, cancelPing := context.WithTimeout(ctx, 5*time.Second)
		defer cancelPing()

		err := m.client.Ping(ctxPing, readpref.Secondary())
		if err != nil {
			return fmt.Errorf("mongoDB health check failed on ping: %w", err)
		}

		return nil
	}
	return "mongo", ready, nil
}

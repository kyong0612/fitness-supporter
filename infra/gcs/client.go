package gcs

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
)

type Client interface {
	Upload(ctx context.Context, bucket, object string, data []byte) error
	GetContentURL(ctx context.Context, bucket, object string) (string, error)
}

type client struct {
	*storage.Client
}

func NewClient(ctx context.Context) (Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gcs client")
	}

	return client{c}, nil
}

func (c client) Upload(ctx context.Context, bucket, object string, data []byte) error {
	if config.IsLocal() {
		bucket = fmt.Sprintf("%s-test", bucket)
	}

	wc := c.Bucket(bucket).Object(object).NewWriter(ctx)

	if _, err := wc.Write(data); err != nil {
		return errors.Wrap(err, "failed to write data")
	}

	if err := wc.Close(); err != nil {
		return errors.Wrap(err, "failed to close writer")
	}

	return nil
}

func (c client) GetContentURL(ctx context.Context, bucket, object string) (string, error) {
	return "https://storage.googleapis.com/" + bucket + "/" + object, nil
}

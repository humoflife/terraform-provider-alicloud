package xpprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetProvider(ctx context.Context) (*schema.Provider, error) {
	return GetSchema(), nil
}

package xpprovider

import (
	"context"

	"github.com/aliyun/terraform-provider-alicloud/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetProvider(ctx context.Context) (*schema.Provider, error) {
	return sdkv2.AlibabaCloudProvider(), nil
}

package xpprovider

import (
	"context"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetProvider(_ context.Context) (*schema.Provider, error) {
	return alicloud.Provider(), nil
}

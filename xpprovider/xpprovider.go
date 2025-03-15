package xpprovider

import (
	"context"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	schemaV2 "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetProviderSchema(_ context.Context) (*schemaV2.Provider, error) {
	var d *schema.ResourceData
	var p *schema.Provider
	var pV2 *schemaV2.Provider

	p = alicloud.Provider().(*schema.Provider), nil
	alicloud.ProviderConfigure(d, p)
	pV2.Schema = p.Schema
	return pV2
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Market product item details of Alibaba Cloud.
//
// > **NOTE:** Available in 1.69.0+
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/market_product.html.markdown.
func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	var rv GetProductResult
	err := ctx.Invoke("alicloud:marketplace/getProduct:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProduct.
type GetProductArgs struct {
	// A available region id used to filter market place Ecs images.
	AvailableRegion *string `pulumi:"availableRegion"`
	// The product code of the market product.
	ProductCode string `pulumi:"productCode"`
}

// A collection of values returned by getProduct.
type GetProductResult struct {
	AvailableRegion *string `pulumi:"availableRegion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ProductCode string `pulumi:"productCode"`
	// A product. It contains the following attributes:
	Products []GetProductProduct `pulumi:"products"`
}

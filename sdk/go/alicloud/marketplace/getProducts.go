// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Market product items of Alibaba Cloud.
//
// > **NOTE:** Available in 1.64.0+
func GetProducts(ctx *pulumi.Context, args *GetProductsArgs, opts ...pulumi.InvokeOption) (*GetProductsResult, error) {
	var rv GetProductsResult
	err := ctx.Invoke("alicloud:marketplace/getProducts:getProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProducts.
type GetProductsArgs struct {
	// The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
	CategoryId *string `pulumi:"categoryId"`
	// A list of product code.
	Ids []string `pulumi:"ids"`
	// A regex string to apply to the product name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The type of products, Valid values: `APP`, `SERVICE`, `MIRROR`, `DOWNLOAD` and `API_SERVICE`.
	ProductType *string `pulumi:"productType"`
	// Search term in this query.
	SearchTerm *string `pulumi:"searchTerm"`
	// This field determines how to sort the filtered results, Valid values: `user_count-desc`, `created_on-desc`, `price-desc` and `score-desc`.
	Sort *string `pulumi:"sort"`
	// The suggested price of the product.
	SuggestedPrice *float64 `pulumi:"suggestedPrice"`
	// The supplier id of the product.
	SupplierId *string `pulumi:"supplierId"`
	// The supplier name keyword of the product.
	SupplierNameKeyword *string `pulumi:"supplierNameKeyword"`
}

// A collection of values returned by getProducts.
type GetProductsResult struct {
	// The category id of the product.
	CategoryId *string `pulumi:"categoryId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of product codes.
	Ids         []string `pulumi:"ids"`
	NameRegex   *string  `pulumi:"nameRegex"`
	OutputFile  *string  `pulumi:"outputFile"`
	ProductType *string  `pulumi:"productType"`
	// A list of products. Each element contains the following attributes:
	Products   []GetProductsProduct `pulumi:"products"`
	SearchTerm *string              `pulumi:"searchTerm"`
	Sort       *string              `pulumi:"sort"`
	// The suggested price of the product.
	SuggestedPrice *float64 `pulumi:"suggestedPrice"`
	// The supplier id of the product.
	SupplierId          *string `pulumi:"supplierId"`
	SupplierNameKeyword *string `pulumi:"supplierNameKeyword"`
}

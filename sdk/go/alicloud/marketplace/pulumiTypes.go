// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package marketplace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GetProductProduct struct {
	// The code of the product.
	Code string `pulumi:"code"`
	// The description of the product.
	Description string `pulumi:"description"`
	// The name of the product.
	Name string `pulumi:"name"`
	// A list of one element containing sku attributes of an object. Each element contains the following attributes:
	Skuses []GetProductProductSkus `pulumi:"skuses"`
}

type GetProductProductInput interface {
	pulumi.Input

	ToGetProductProductOutput() GetProductProductOutput
	ToGetProductProductOutputWithContext(context.Context) GetProductProductOutput
}

type GetProductProductArgs struct {
	// The code of the product.
	Code pulumi.StringInput `pulumi:"code"`
	// The description of the product.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the product.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of one element containing sku attributes of an object. Each element contains the following attributes:
	Skuses GetProductProductSkusArrayInput `pulumi:"skuses"`
}

func (GetProductProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProduct)(nil)).Elem()
}

func (i GetProductProductArgs) ToGetProductProductOutput() GetProductProductOutput {
	return i.ToGetProductProductOutputWithContext(context.Background())
}

func (i GetProductProductArgs) ToGetProductProductOutputWithContext(ctx context.Context) GetProductProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductOutput)
}

type GetProductProductArrayInput interface {
	pulumi.Input

	ToGetProductProductArrayOutput() GetProductProductArrayOutput
	ToGetProductProductArrayOutputWithContext(context.Context) GetProductProductArrayOutput
}

type GetProductProductArray []GetProductProductInput

func (GetProductProductArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProduct)(nil)).Elem()
}

func (i GetProductProductArray) ToGetProductProductArrayOutput() GetProductProductArrayOutput {
	return i.ToGetProductProductArrayOutputWithContext(context.Background())
}

func (i GetProductProductArray) ToGetProductProductArrayOutputWithContext(ctx context.Context) GetProductProductArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductArrayOutput)
}

type GetProductProductOutput struct{ *pulumi.OutputState }

func (GetProductProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProduct)(nil)).Elem()
}

func (o GetProductProductOutput) ToGetProductProductOutput() GetProductProductOutput {
	return o
}

func (o GetProductProductOutput) ToGetProductProductOutputWithContext(ctx context.Context) GetProductProductOutput {
	return o
}

// The code of the product.
func (o GetProductProductOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProduct) string { return v.Code }).(pulumi.StringOutput)
}

// The description of the product.
func (o GetProductProductOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProduct) string { return v.Description }).(pulumi.StringOutput)
}

// The name of the product.
func (o GetProductProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProduct) string { return v.Name }).(pulumi.StringOutput)
}

// A list of one element containing sku attributes of an object. Each element contains the following attributes:
func (o GetProductProductOutput) Skuses() GetProductProductSkusArrayOutput {
	return o.ApplyT(func(v GetProductProduct) []GetProductProductSkus { return v.Skuses }).(GetProductProductSkusArrayOutput)
}

type GetProductProductArrayOutput struct{ *pulumi.OutputState }

func (GetProductProductArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProduct)(nil)).Elem()
}

func (o GetProductProductArrayOutput) ToGetProductProductArrayOutput() GetProductProductArrayOutput {
	return o
}

func (o GetProductProductArrayOutput) ToGetProductProductArrayOutputWithContext(ctx context.Context) GetProductProductArrayOutput {
	return o
}

func (o GetProductProductArrayOutput) Index(i pulumi.IntInput) GetProductProductOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProductProduct {
		return vs[0].([]GetProductProduct)[vs[1].(int)]
	}).(GetProductProductOutput)
}

type GetProductProductSkus struct {
	// The list of custom ECS images, Each element contains the following attributes:
	Images []GetProductProductSkusImage `pulumi:"images"`
	// The list of package version details of this product sku, Each element contains the following attributes:
	PackageVersions []GetProductProductSkusPackageVersion `pulumi:"packageVersions"`
	// The sku code of this product sku.
	SkuCode string `pulumi:"skuCode"`
	// The sku name of this product sku.
	SkuName string `pulumi:"skuName"`
}

type GetProductProductSkusInput interface {
	pulumi.Input

	ToGetProductProductSkusOutput() GetProductProductSkusOutput
	ToGetProductProductSkusOutputWithContext(context.Context) GetProductProductSkusOutput
}

type GetProductProductSkusArgs struct {
	// The list of custom ECS images, Each element contains the following attributes:
	Images GetProductProductSkusImageArrayInput `pulumi:"images"`
	// The list of package version details of this product sku, Each element contains the following attributes:
	PackageVersions GetProductProductSkusPackageVersionArrayInput `pulumi:"packageVersions"`
	// The sku code of this product sku.
	SkuCode pulumi.StringInput `pulumi:"skuCode"`
	// The sku name of this product sku.
	SkuName pulumi.StringInput `pulumi:"skuName"`
}

func (GetProductProductSkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkus)(nil)).Elem()
}

func (i GetProductProductSkusArgs) ToGetProductProductSkusOutput() GetProductProductSkusOutput {
	return i.ToGetProductProductSkusOutputWithContext(context.Background())
}

func (i GetProductProductSkusArgs) ToGetProductProductSkusOutputWithContext(ctx context.Context) GetProductProductSkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusOutput)
}

type GetProductProductSkusArrayInput interface {
	pulumi.Input

	ToGetProductProductSkusArrayOutput() GetProductProductSkusArrayOutput
	ToGetProductProductSkusArrayOutputWithContext(context.Context) GetProductProductSkusArrayOutput
}

type GetProductProductSkusArray []GetProductProductSkusInput

func (GetProductProductSkusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkus)(nil)).Elem()
}

func (i GetProductProductSkusArray) ToGetProductProductSkusArrayOutput() GetProductProductSkusArrayOutput {
	return i.ToGetProductProductSkusArrayOutputWithContext(context.Background())
}

func (i GetProductProductSkusArray) ToGetProductProductSkusArrayOutputWithContext(ctx context.Context) GetProductProductSkusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusArrayOutput)
}

type GetProductProductSkusOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkus)(nil)).Elem()
}

func (o GetProductProductSkusOutput) ToGetProductProductSkusOutput() GetProductProductSkusOutput {
	return o
}

func (o GetProductProductSkusOutput) ToGetProductProductSkusOutputWithContext(ctx context.Context) GetProductProductSkusOutput {
	return o
}

// The list of custom ECS images, Each element contains the following attributes:
func (o GetProductProductSkusOutput) Images() GetProductProductSkusImageArrayOutput {
	return o.ApplyT(func(v GetProductProductSkus) []GetProductProductSkusImage { return v.Images }).(GetProductProductSkusImageArrayOutput)
}

// The list of package version details of this product sku, Each element contains the following attributes:
func (o GetProductProductSkusOutput) PackageVersions() GetProductProductSkusPackageVersionArrayOutput {
	return o.ApplyT(func(v GetProductProductSkus) []GetProductProductSkusPackageVersion { return v.PackageVersions }).(GetProductProductSkusPackageVersionArrayOutput)
}

// The sku code of this product sku.
func (o GetProductProductSkusOutput) SkuCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkus) string { return v.SkuCode }).(pulumi.StringOutput)
}

// The sku name of this product sku.
func (o GetProductProductSkusOutput) SkuName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkus) string { return v.SkuName }).(pulumi.StringOutput)
}

type GetProductProductSkusArrayOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkus)(nil)).Elem()
}

func (o GetProductProductSkusArrayOutput) ToGetProductProductSkusArrayOutput() GetProductProductSkusArrayOutput {
	return o
}

func (o GetProductProductSkusArrayOutput) ToGetProductProductSkusArrayOutputWithContext(ctx context.Context) GetProductProductSkusArrayOutput {
	return o
}

func (o GetProductProductSkusArrayOutput) Index(i pulumi.IntInput) GetProductProductSkusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProductProductSkus {
		return vs[0].([]GetProductProductSkus)[vs[1].(int)]
	}).(GetProductProductSkusOutput)
}

type GetProductProductSkusImage struct {
	// The Ecs image id.
	ImageId string `pulumi:"imageId"`
	// The Ecs image display name.
	ImageName string `pulumi:"imageName"`
	// The Ecs image region.
	RegionId string `pulumi:"regionId"`
}

type GetProductProductSkusImageInput interface {
	pulumi.Input

	ToGetProductProductSkusImageOutput() GetProductProductSkusImageOutput
	ToGetProductProductSkusImageOutputWithContext(context.Context) GetProductProductSkusImageOutput
}

type GetProductProductSkusImageArgs struct {
	// The Ecs image id.
	ImageId pulumi.StringInput `pulumi:"imageId"`
	// The Ecs image display name.
	ImageName pulumi.StringInput `pulumi:"imageName"`
	// The Ecs image region.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (GetProductProductSkusImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkusImage)(nil)).Elem()
}

func (i GetProductProductSkusImageArgs) ToGetProductProductSkusImageOutput() GetProductProductSkusImageOutput {
	return i.ToGetProductProductSkusImageOutputWithContext(context.Background())
}

func (i GetProductProductSkusImageArgs) ToGetProductProductSkusImageOutputWithContext(ctx context.Context) GetProductProductSkusImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusImageOutput)
}

type GetProductProductSkusImageArrayInput interface {
	pulumi.Input

	ToGetProductProductSkusImageArrayOutput() GetProductProductSkusImageArrayOutput
	ToGetProductProductSkusImageArrayOutputWithContext(context.Context) GetProductProductSkusImageArrayOutput
}

type GetProductProductSkusImageArray []GetProductProductSkusImageInput

func (GetProductProductSkusImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkusImage)(nil)).Elem()
}

func (i GetProductProductSkusImageArray) ToGetProductProductSkusImageArrayOutput() GetProductProductSkusImageArrayOutput {
	return i.ToGetProductProductSkusImageArrayOutputWithContext(context.Background())
}

func (i GetProductProductSkusImageArray) ToGetProductProductSkusImageArrayOutputWithContext(ctx context.Context) GetProductProductSkusImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusImageArrayOutput)
}

type GetProductProductSkusImageOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkusImage)(nil)).Elem()
}

func (o GetProductProductSkusImageOutput) ToGetProductProductSkusImageOutput() GetProductProductSkusImageOutput {
	return o
}

func (o GetProductProductSkusImageOutput) ToGetProductProductSkusImageOutputWithContext(ctx context.Context) GetProductProductSkusImageOutput {
	return o
}

// The Ecs image id.
func (o GetProductProductSkusImageOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkusImage) string { return v.ImageId }).(pulumi.StringOutput)
}

// The Ecs image display name.
func (o GetProductProductSkusImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkusImage) string { return v.ImageName }).(pulumi.StringOutput)
}

// The Ecs image region.
func (o GetProductProductSkusImageOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkusImage) string { return v.RegionId }).(pulumi.StringOutput)
}

type GetProductProductSkusImageArrayOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkusImage)(nil)).Elem()
}

func (o GetProductProductSkusImageArrayOutput) ToGetProductProductSkusImageArrayOutput() GetProductProductSkusImageArrayOutput {
	return o
}

func (o GetProductProductSkusImageArrayOutput) ToGetProductProductSkusImageArrayOutputWithContext(ctx context.Context) GetProductProductSkusImageArrayOutput {
	return o
}

func (o GetProductProductSkusImageArrayOutput) Index(i pulumi.IntInput) GetProductProductSkusImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProductProductSkusImage {
		return vs[0].([]GetProductProductSkusImage)[vs[1].(int)]
	}).(GetProductProductSkusImageOutput)
}

type GetProductProductSkusPackageVersion struct {
	// The package name of this product sku package.
	PackageName string `pulumi:"packageName"`
	// The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.
	PackageVersion string `pulumi:"packageVersion"`
}

type GetProductProductSkusPackageVersionInput interface {
	pulumi.Input

	ToGetProductProductSkusPackageVersionOutput() GetProductProductSkusPackageVersionOutput
	ToGetProductProductSkusPackageVersionOutputWithContext(context.Context) GetProductProductSkusPackageVersionOutput
}

type GetProductProductSkusPackageVersionArgs struct {
	// The package name of this product sku package.
	PackageName pulumi.StringInput `pulumi:"packageName"`
	// The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.
	PackageVersion pulumi.StringInput `pulumi:"packageVersion"`
}

func (GetProductProductSkusPackageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkusPackageVersion)(nil)).Elem()
}

func (i GetProductProductSkusPackageVersionArgs) ToGetProductProductSkusPackageVersionOutput() GetProductProductSkusPackageVersionOutput {
	return i.ToGetProductProductSkusPackageVersionOutputWithContext(context.Background())
}

func (i GetProductProductSkusPackageVersionArgs) ToGetProductProductSkusPackageVersionOutputWithContext(ctx context.Context) GetProductProductSkusPackageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusPackageVersionOutput)
}

type GetProductProductSkusPackageVersionArrayInput interface {
	pulumi.Input

	ToGetProductProductSkusPackageVersionArrayOutput() GetProductProductSkusPackageVersionArrayOutput
	ToGetProductProductSkusPackageVersionArrayOutputWithContext(context.Context) GetProductProductSkusPackageVersionArrayOutput
}

type GetProductProductSkusPackageVersionArray []GetProductProductSkusPackageVersionInput

func (GetProductProductSkusPackageVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkusPackageVersion)(nil)).Elem()
}

func (i GetProductProductSkusPackageVersionArray) ToGetProductProductSkusPackageVersionArrayOutput() GetProductProductSkusPackageVersionArrayOutput {
	return i.ToGetProductProductSkusPackageVersionArrayOutputWithContext(context.Background())
}

func (i GetProductProductSkusPackageVersionArray) ToGetProductProductSkusPackageVersionArrayOutputWithContext(ctx context.Context) GetProductProductSkusPackageVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductProductSkusPackageVersionArrayOutput)
}

type GetProductProductSkusPackageVersionOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusPackageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductProductSkusPackageVersion)(nil)).Elem()
}

func (o GetProductProductSkusPackageVersionOutput) ToGetProductProductSkusPackageVersionOutput() GetProductProductSkusPackageVersionOutput {
	return o
}

func (o GetProductProductSkusPackageVersionOutput) ToGetProductProductSkusPackageVersionOutputWithContext(ctx context.Context) GetProductProductSkusPackageVersionOutput {
	return o
}

// The package name of this product sku package.
func (o GetProductProductSkusPackageVersionOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkusPackageVersion) string { return v.PackageName }).(pulumi.StringOutput)
}

// The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.
func (o GetProductProductSkusPackageVersionOutput) PackageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductProductSkusPackageVersion) string { return v.PackageVersion }).(pulumi.StringOutput)
}

type GetProductProductSkusPackageVersionArrayOutput struct{ *pulumi.OutputState }

func (GetProductProductSkusPackageVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductProductSkusPackageVersion)(nil)).Elem()
}

func (o GetProductProductSkusPackageVersionArrayOutput) ToGetProductProductSkusPackageVersionArrayOutput() GetProductProductSkusPackageVersionArrayOutput {
	return o
}

func (o GetProductProductSkusPackageVersionArrayOutput) ToGetProductProductSkusPackageVersionArrayOutputWithContext(ctx context.Context) GetProductProductSkusPackageVersionArrayOutput {
	return o
}

func (o GetProductProductSkusPackageVersionArrayOutput) Index(i pulumi.IntInput) GetProductProductSkusPackageVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProductProductSkusPackageVersion {
		return vs[0].([]GetProductProductSkusPackageVersion)[vs[1].(int)]
	}).(GetProductProductSkusPackageVersionOutput)
}

type GetProductsProduct struct {
	// The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
	CategoryId int `pulumi:"categoryId"`
	// The code of the product.
	Code string `pulumi:"code"`
	// The delivery date of the product.
	DeliveryDate string `pulumi:"deliveryDate"`
	// The delivery way of the product.
	DeliveryWay string `pulumi:"deliveryWay"`
	// The image URL of the product.
	ImageUrl string `pulumi:"imageUrl"`
	// The name of the product.
	Name string `pulumi:"name"`
	// The operation system of the product.
	OperationSystem string `pulumi:"operationSystem"`
	// The rating information of the product.
	Score string `pulumi:"score"`
	// The short description of the product.
	ShortDescription string `pulumi:"shortDescription"`
	// The suggested price of the product.
	SuggestedPrice string `pulumi:"suggestedPrice"`
	// The supplier id of the product.
	SupplierId int `pulumi:"supplierId"`
	// The supplier name of the product.
	SupplierName string `pulumi:"supplierName"`
	// The tags of the product.
	Tags string `pulumi:"tags"`
	// The detail page URL of the product.
	TargetUrl string `pulumi:"targetUrl"`
	// The warranty date of the product.
	WarrantyDate string `pulumi:"warrantyDate"`
}

type GetProductsProductInput interface {
	pulumi.Input

	ToGetProductsProductOutput() GetProductsProductOutput
	ToGetProductsProductOutputWithContext(context.Context) GetProductsProductOutput
}

type GetProductsProductArgs struct {
	// The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
	CategoryId pulumi.IntInput `pulumi:"categoryId"`
	// The code of the product.
	Code pulumi.StringInput `pulumi:"code"`
	// The delivery date of the product.
	DeliveryDate pulumi.StringInput `pulumi:"deliveryDate"`
	// The delivery way of the product.
	DeliveryWay pulumi.StringInput `pulumi:"deliveryWay"`
	// The image URL of the product.
	ImageUrl pulumi.StringInput `pulumi:"imageUrl"`
	// The name of the product.
	Name pulumi.StringInput `pulumi:"name"`
	// The operation system of the product.
	OperationSystem pulumi.StringInput `pulumi:"operationSystem"`
	// The rating information of the product.
	Score pulumi.StringInput `pulumi:"score"`
	// The short description of the product.
	ShortDescription pulumi.StringInput `pulumi:"shortDescription"`
	// The suggested price of the product.
	SuggestedPrice pulumi.StringInput `pulumi:"suggestedPrice"`
	// The supplier id of the product.
	SupplierId pulumi.IntInput `pulumi:"supplierId"`
	// The supplier name of the product.
	SupplierName pulumi.StringInput `pulumi:"supplierName"`
	// The tags of the product.
	Tags pulumi.StringInput `pulumi:"tags"`
	// The detail page URL of the product.
	TargetUrl pulumi.StringInput `pulumi:"targetUrl"`
	// The warranty date of the product.
	WarrantyDate pulumi.StringInput `pulumi:"warrantyDate"`
}

func (GetProductsProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductsProduct)(nil)).Elem()
}

func (i GetProductsProductArgs) ToGetProductsProductOutput() GetProductsProductOutput {
	return i.ToGetProductsProductOutputWithContext(context.Background())
}

func (i GetProductsProductArgs) ToGetProductsProductOutputWithContext(ctx context.Context) GetProductsProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductsProductOutput)
}

type GetProductsProductArrayInput interface {
	pulumi.Input

	ToGetProductsProductArrayOutput() GetProductsProductArrayOutput
	ToGetProductsProductArrayOutputWithContext(context.Context) GetProductsProductArrayOutput
}

type GetProductsProductArray []GetProductsProductInput

func (GetProductsProductArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductsProduct)(nil)).Elem()
}

func (i GetProductsProductArray) ToGetProductsProductArrayOutput() GetProductsProductArrayOutput {
	return i.ToGetProductsProductArrayOutputWithContext(context.Background())
}

func (i GetProductsProductArray) ToGetProductsProductArrayOutputWithContext(ctx context.Context) GetProductsProductArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductsProductArrayOutput)
}

type GetProductsProductOutput struct{ *pulumi.OutputState }

func (GetProductsProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductsProduct)(nil)).Elem()
}

func (o GetProductsProductOutput) ToGetProductsProductOutput() GetProductsProductOutput {
	return o
}

func (o GetProductsProductOutput) ToGetProductsProductOutputWithContext(ctx context.Context) GetProductsProductOutput {
	return o
}

// The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
func (o GetProductsProductOutput) CategoryId() pulumi.IntOutput {
	return o.ApplyT(func(v GetProductsProduct) int { return v.CategoryId }).(pulumi.IntOutput)
}

// The code of the product.
func (o GetProductsProductOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.Code }).(pulumi.StringOutput)
}

// The delivery date of the product.
func (o GetProductsProductOutput) DeliveryDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.DeliveryDate }).(pulumi.StringOutput)
}

// The delivery way of the product.
func (o GetProductsProductOutput) DeliveryWay() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.DeliveryWay }).(pulumi.StringOutput)
}

// The image URL of the product.
func (o GetProductsProductOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.ImageUrl }).(pulumi.StringOutput)
}

// The name of the product.
func (o GetProductsProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.Name }).(pulumi.StringOutput)
}

// The operation system of the product.
func (o GetProductsProductOutput) OperationSystem() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.OperationSystem }).(pulumi.StringOutput)
}

// The rating information of the product.
func (o GetProductsProductOutput) Score() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.Score }).(pulumi.StringOutput)
}

// The short description of the product.
func (o GetProductsProductOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.ShortDescription }).(pulumi.StringOutput)
}

// The suggested price of the product.
func (o GetProductsProductOutput) SuggestedPrice() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.SuggestedPrice }).(pulumi.StringOutput)
}

// The supplier id of the product.
func (o GetProductsProductOutput) SupplierId() pulumi.IntOutput {
	return o.ApplyT(func(v GetProductsProduct) int { return v.SupplierId }).(pulumi.IntOutput)
}

// The supplier name of the product.
func (o GetProductsProductOutput) SupplierName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.SupplierName }).(pulumi.StringOutput)
}

// The tags of the product.
func (o GetProductsProductOutput) Tags() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.Tags }).(pulumi.StringOutput)
}

// The detail page URL of the product.
func (o GetProductsProductOutput) TargetUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.TargetUrl }).(pulumi.StringOutput)
}

// The warranty date of the product.
func (o GetProductsProductOutput) WarrantyDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductsProduct) string { return v.WarrantyDate }).(pulumi.StringOutput)
}

type GetProductsProductArrayOutput struct{ *pulumi.OutputState }

func (GetProductsProductArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductsProduct)(nil)).Elem()
}

func (o GetProductsProductArrayOutput) ToGetProductsProductArrayOutput() GetProductsProductArrayOutput {
	return o
}

func (o GetProductsProductArrayOutput) ToGetProductsProductArrayOutputWithContext(ctx context.Context) GetProductsProductArrayOutput {
	return o
}

func (o GetProductsProductArrayOutput) Index(i pulumi.IntInput) GetProductsProductOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProductsProduct {
		return vs[0].([]GetProductsProduct)[vs[1].(int)]
	}).(GetProductsProductOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductProductOutput{})
	pulumi.RegisterOutputType(GetProductProductArrayOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusArrayOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusImageOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusImageArrayOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusPackageVersionOutput{})
	pulumi.RegisterOutputType(GetProductProductSkusPackageVersionArrayOutput{})
	pulumi.RegisterOutputType(GetProductsProductOutput{})
	pulumi.RegisterOutputType(GetProductsProductArrayOutput{})
}

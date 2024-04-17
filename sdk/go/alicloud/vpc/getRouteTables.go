// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of Route Tables owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.36.0+.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// cfg := config.New(ctx, "")
// name := "route-tables-datasource-example-name";
// if param := cfg.Get("name"); param != ""{
// name = param
// }
// fooNetwork, err := vpc.NewNetwork(ctx, "foo", &vpc.NetworkArgs{
// CidrBlock: pulumi.String("172.16.0.0/12"),
// VpcName: pulumi.String(name),
// })
// if err != nil {
// return err
// }
// fooRouteTable, err := vpc.NewRouteTable(ctx, "foo", &vpc.RouteTableArgs{
// VpcId: fooNetwork.ID(),
// RouteTableName: pulumi.String(name),
// Description: pulumi.String(name),
// })
// if err != nil {
// return err
// }
// foo := vpc.GetRouteTablesOutput(ctx, vpc.GetRouteTablesOutputArgs{
// Ids: pulumi.StringArray{
// fooRouteTable.ID(),
// },
// }, nil);
// ctx.Export("routeTableIds", foo.ApplyT(func(foo vpc.GetRouteTablesResult) (interface{}, error) {
// return foo.Ids, nil
// }).(pulumi.Interface{}Output))
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
func GetRouteTables(ctx *pulumi.Context, args *GetRouteTablesArgs, opts ...pulumi.InvokeOption) (*GetRouteTablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouteTablesResult
	err := ctx.Invoke("alicloud:vpc/getRouteTables:getRouteTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTables.
type GetRouteTablesArgs struct {
	// A list of Route Tables IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter route tables by name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The Id of resource group which route tables belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table name.
	RouteTableName *string `pulumi:"routeTableName"`
	// The router ID.
	RouterId *string `pulumi:"routerId"`
	// The route type of route table. Valid values: `VRouter` and `VBR`.
	RouterType *string `pulumi:"routerType"`
	// The status of resource. Valid values: `Available` and `Pending`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc id of the route table.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getRouteTables.
type GetRouteTablesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of Route Tables IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Route Tables names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	// The Id of resource group which route tables belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table name.
	RouteTableName *string `pulumi:"routeTableName"`
	// Router Id of the route table.
	RouterId *string `pulumi:"routerId"`
	// The route type.
	RouterType *string `pulumi:"routerType"`
	// The status of route table.
	Status *string `pulumi:"status"`
	// A list of Route Tables. Each element contains the following attributes:
	Tables     []GetRouteTablesTable  `pulumi:"tables"`
	Tags       map[string]interface{} `pulumi:"tags"`
	TotalCount int                    `pulumi:"totalCount"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

func GetRouteTablesOutput(ctx *pulumi.Context, args GetRouteTablesOutputArgs, opts ...pulumi.InvokeOption) GetRouteTablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouteTablesResult, error) {
			args := v.(GetRouteTablesArgs)
			r, err := GetRouteTables(ctx, &args, opts...)
			var s GetRouteTablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouteTablesResultOutput)
}

// A collection of arguments for invoking getRouteTables.
type GetRouteTablesOutputArgs struct {
	// A list of Route Tables IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter route tables by name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The Id of resource group which route tables belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The route table name.
	RouteTableName pulumi.StringPtrInput `pulumi:"routeTableName"`
	// The router ID.
	RouterId pulumi.StringPtrInput `pulumi:"routerId"`
	// The route type of route table. Valid values: `VRouter` and `VBR`.
	RouterType pulumi.StringPtrInput `pulumi:"routerType"`
	// The status of resource. Valid values: `Available` and `Pending`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Vpc id of the route table.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetRouteTablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteTablesArgs)(nil)).Elem()
}

// A collection of values returned by getRouteTables.
type GetRouteTablesResultOutput struct{ *pulumi.OutputState }

func (GetRouteTablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteTablesResult)(nil)).Elem()
}

func (o GetRouteTablesResultOutput) ToGetRouteTablesResultOutput() GetRouteTablesResultOutput {
	return o
}

func (o GetRouteTablesResultOutput) ToGetRouteTablesResultOutputWithContext(ctx context.Context) GetRouteTablesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouteTablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouteTablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Optional) A list of Route Tables IDs.
func (o GetRouteTablesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouteTablesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRouteTablesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Route Tables names.
func (o GetRouteTablesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouteTablesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRouteTablesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRouteTablesResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetRouteTablesResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// The Id of resource group which route tables belongs.
func (o GetRouteTablesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The route table name.
func (o GetRouteTablesResultOutput) RouteTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.RouteTableName }).(pulumi.StringPtrOutput)
}

// Router Id of the route table.
func (o GetRouteTablesResultOutput) RouterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.RouterId }).(pulumi.StringPtrOutput)
}

// The route type.
func (o GetRouteTablesResultOutput) RouterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.RouterType }).(pulumi.StringPtrOutput)
}

// The status of route table.
func (o GetRouteTablesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of Route Tables. Each element contains the following attributes:
func (o GetRouteTablesResultOutput) Tables() GetRouteTablesTableArrayOutput {
	return o.ApplyT(func(v GetRouteTablesResult) []GetRouteTablesTable { return v.Tables }).(GetRouteTablesTableArrayOutput)
}

func (o GetRouteTablesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetRouteTablesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetRouteTablesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetRouteTablesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The VPC ID.
func (o GetRouteTablesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteTablesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouteTablesResultOutput{})
}

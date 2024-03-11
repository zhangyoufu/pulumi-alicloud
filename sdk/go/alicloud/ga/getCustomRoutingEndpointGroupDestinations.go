// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Global Accelerator (GA) Custom Routing Endpoint Group Destinations of the current Alibaba Cloud user.
//
// > **NOTE:** Available in 1.197.0+
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ga.GetCustomRoutingEndpointGroupDestinations(ctx, &ga.GetCustomRoutingEndpointGroupDestinationsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				AcceleratorId: "your_accelerator_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("gaCustomRoutingEndpointGroupDestinationsId1", ids.CustomRoutingEndpointGroupDestinations[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetCustomRoutingEndpointGroupDestinations(ctx *pulumi.Context, args *GetCustomRoutingEndpointGroupDestinationsArgs, opts ...pulumi.InvokeOption) (*GetCustomRoutingEndpointGroupDestinationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomRoutingEndpointGroupDestinationsResult
	err := ctx.Invoke("alicloud:ga/getCustomRoutingEndpointGroupDestinations:getCustomRoutingEndpointGroupDestinations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomRoutingEndpointGroupDestinations.
type GetCustomRoutingEndpointGroupDestinationsArgs struct {
	// The ID of the GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The ID of the endpoint group.
	EndpointGroupId *string `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort *int `pulumi:"fromPort"`
	// A list of Custom Routing Endpoint Group Destination IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols []string `pulumi:"protocols"`
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort *int `pulumi:"toPort"`
}

// A collection of values returned by getCustomRoutingEndpointGroupDestinations.
type GetCustomRoutingEndpointGroupDestinationsResult struct {
	// The ID of the GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// A list of Custom Routing Endpoint Group Destinations. Each element contains the following attributes:
	CustomRoutingEndpointGroupDestinations []GetCustomRoutingEndpointGroupDestinationsCustomRoutingEndpointGroupDestination `pulumi:"customRoutingEndpointGroupDestinations"`
	// The ID of the Custom Routing Endpoint Group.
	EndpointGroupId *string `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group.
	FromPort *int `pulumi:"fromPort"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The backend service protocol of the endpoint group.
	Protocols []string `pulumi:"protocols"`
	// The end port of the backend service port range of the endpoint group.
	ToPort *int `pulumi:"toPort"`
}

func GetCustomRoutingEndpointGroupDestinationsOutput(ctx *pulumi.Context, args GetCustomRoutingEndpointGroupDestinationsOutputArgs, opts ...pulumi.InvokeOption) GetCustomRoutingEndpointGroupDestinationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomRoutingEndpointGroupDestinationsResult, error) {
			args := v.(GetCustomRoutingEndpointGroupDestinationsArgs)
			r, err := GetCustomRoutingEndpointGroupDestinations(ctx, &args, opts...)
			var s GetCustomRoutingEndpointGroupDestinationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomRoutingEndpointGroupDestinationsResultOutput)
}

// A collection of arguments for invoking getCustomRoutingEndpointGroupDestinations.
type GetCustomRoutingEndpointGroupDestinationsOutputArgs struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringInput `pulumi:"acceleratorId"`
	// The ID of the endpoint group.
	EndpointGroupId pulumi.StringPtrInput `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort pulumi.IntPtrInput `pulumi:"fromPort"`
	// A list of Custom Routing Endpoint Group Destination IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the listener.
	ListenerId pulumi.StringPtrInput `pulumi:"listenerId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols pulumi.StringArrayInput `pulumi:"protocols"`
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort pulumi.IntPtrInput `pulumi:"toPort"`
}

func (GetCustomRoutingEndpointGroupDestinationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomRoutingEndpointGroupDestinationsArgs)(nil)).Elem()
}

// A collection of values returned by getCustomRoutingEndpointGroupDestinations.
type GetCustomRoutingEndpointGroupDestinationsResultOutput struct{ *pulumi.OutputState }

func (GetCustomRoutingEndpointGroupDestinationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomRoutingEndpointGroupDestinationsResult)(nil)).Elem()
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) ToGetCustomRoutingEndpointGroupDestinationsResultOutput() GetCustomRoutingEndpointGroupDestinationsResultOutput {
	return o
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) ToGetCustomRoutingEndpointGroupDestinationsResultOutputWithContext(ctx context.Context) GetCustomRoutingEndpointGroupDestinationsResultOutput {
	return o
}

// The ID of the GA instance.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) string { return v.AcceleratorId }).(pulumi.StringOutput)
}

// A list of Custom Routing Endpoint Group Destinations. Each element contains the following attributes:
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) CustomRoutingEndpointGroupDestinations() GetCustomRoutingEndpointGroupDestinationsCustomRoutingEndpointGroupDestinationArrayOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) []GetCustomRoutingEndpointGroupDestinationsCustomRoutingEndpointGroupDestination {
		return v.CustomRoutingEndpointGroupDestinations
	}).(GetCustomRoutingEndpointGroupDestinationsCustomRoutingEndpointGroupDestinationArrayOutput)
}

// The ID of the Custom Routing Endpoint Group.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) EndpointGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *string { return v.EndpointGroupId }).(pulumi.StringPtrOutput)
}

// The start port of the backend service port range of the endpoint group.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) FromPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *int { return v.FromPort }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The ID of the listener.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) ListenerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *string { return v.ListenerId }).(pulumi.StringPtrOutput)
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// The backend service protocol of the endpoint group.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

// The end port of the backend service port range of the endpoint group.
func (o GetCustomRoutingEndpointGroupDestinationsResultOutput) ToPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetCustomRoutingEndpointGroupDestinationsResult) *int { return v.ToPort }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomRoutingEndpointGroupDestinationsResultOutput{})
}

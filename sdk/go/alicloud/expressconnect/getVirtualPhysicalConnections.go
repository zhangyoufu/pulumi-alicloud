// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Express Connect Virtual Physical Connection available to the user.
//
// > **NOTE:** Available in 1.196.0+
func GetVirtualPhysicalConnections(ctx *pulumi.Context, args *GetVirtualPhysicalConnectionsArgs, opts ...pulumi.InvokeOption) (*GetVirtualPhysicalConnectionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualPhysicalConnectionsResult
	err := ctx.Invoke("alicloud:expressconnect/getVirtualPhysicalConnections:getVirtualPhysicalConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualPhysicalConnections.
type GetVirtualPhysicalConnectionsArgs struct {
	// The commercial status of the physical line. Value:
	// - **Normal**: activated.
	// - **Financialized**: Arrears locked.
	// - **SecurityLocked**: locked for security reasons.
	BusinessStatus *string `pulumi:"businessStatus"`
	// A list of Virtual Physical Connection IDs.
	Ids         []string `pulumi:"ids"`
	IsConfirmed *bool    `pulumi:"isConfirmed"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the instance of the physical connection.
	ParentPhysicalConnectionId *string `pulumi:"parentPhysicalConnectionId"`
	// The ID of the hosted connection. You can specify multiple hosted connection IDs.
	VirtualPhysicalConnectionIds []string `pulumi:"virtualPhysicalConnectionIds"`
	// The business status of the shared line. Value:
	// - **Confirmed**: The shared line has been Confirmed to receive.
	// - **UnConfirmed**: The shared line has not been confirmed to be received.
	// - **Deleted**: The shared line has been Deleted.
	VirtualPhysicalConnectionStatus *string `pulumi:"virtualPhysicalConnectionStatus"`
	// The VLAN ID of the hosted connection. You can specify multiple VLAN IDs.
	VlanIds []int `pulumi:"vlanIds"`
	// The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
	VpconnAliUid *string `pulumi:"vpconnAliUid"`
}

// A collection of values returned by getVirtualPhysicalConnections.
type GetVirtualPhysicalConnectionsResult struct {
	// The commercial status of the physical line. Value:-**Normal**: activated.-**Financialized**: Arrears locked.-**SecurityLocked**: locked for security reasons.
	BusinessStatus *string `pulumi:"businessStatus"`
	// A list of Virtual Physical Connection Entries. Each element contains the following attributes:
	Connections []GetVirtualPhysicalConnectionsConnection `pulumi:"connections"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Virtual Physical Connection IDs.
	Ids         []string `pulumi:"ids"`
	IsConfirmed *bool    `pulumi:"isConfirmed"`
	NameRegex   *string  `pulumi:"nameRegex"`
	// A list of name of Virtual Physical Connections.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the instance of the physical connection.
	ParentPhysicalConnectionId   *string  `pulumi:"parentPhysicalConnectionId"`
	VirtualPhysicalConnectionIds []string `pulumi:"virtualPhysicalConnectionIds"`
	// The business status of the shared line. Value:-**Confirmed**: The shared line has been Confirmed to receive.-**UnConfirmed**: The shared line has not been confirmed to be received.-**Deleted**: The shared line has been Deleted.
	VirtualPhysicalConnectionStatus *string `pulumi:"virtualPhysicalConnectionStatus"`
	VlanIds                         []int   `pulumi:"vlanIds"`
	// The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
	VpconnAliUid *string `pulumi:"vpconnAliUid"`
}

func GetVirtualPhysicalConnectionsOutput(ctx *pulumi.Context, args GetVirtualPhysicalConnectionsOutputArgs, opts ...pulumi.InvokeOption) GetVirtualPhysicalConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualPhysicalConnectionsResult, error) {
			args := v.(GetVirtualPhysicalConnectionsArgs)
			r, err := GetVirtualPhysicalConnections(ctx, &args, opts...)
			var s GetVirtualPhysicalConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualPhysicalConnectionsResultOutput)
}

// A collection of arguments for invoking getVirtualPhysicalConnections.
type GetVirtualPhysicalConnectionsOutputArgs struct {
	// The commercial status of the physical line. Value:
	// - **Normal**: activated.
	// - **Financialized**: Arrears locked.
	// - **SecurityLocked**: locked for security reasons.
	BusinessStatus pulumi.StringPtrInput `pulumi:"businessStatus"`
	// A list of Virtual Physical Connection IDs.
	Ids         pulumi.StringArrayInput `pulumi:"ids"`
	IsConfirmed pulumi.BoolPtrInput     `pulumi:"isConfirmed"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the instance of the physical connection.
	ParentPhysicalConnectionId pulumi.StringPtrInput `pulumi:"parentPhysicalConnectionId"`
	// The ID of the hosted connection. You can specify multiple hosted connection IDs.
	VirtualPhysicalConnectionIds pulumi.StringArrayInput `pulumi:"virtualPhysicalConnectionIds"`
	// The business status of the shared line. Value:
	// - **Confirmed**: The shared line has been Confirmed to receive.
	// - **UnConfirmed**: The shared line has not been confirmed to be received.
	// - **Deleted**: The shared line has been Deleted.
	VirtualPhysicalConnectionStatus pulumi.StringPtrInput `pulumi:"virtualPhysicalConnectionStatus"`
	// The VLAN ID of the hosted connection. You can specify multiple VLAN IDs.
	VlanIds pulumi.IntArrayInput `pulumi:"vlanIds"`
	// The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
	VpconnAliUid pulumi.StringPtrInput `pulumi:"vpconnAliUid"`
}

func (GetVirtualPhysicalConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualPhysicalConnectionsArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualPhysicalConnections.
type GetVirtualPhysicalConnectionsResultOutput struct{ *pulumi.OutputState }

func (GetVirtualPhysicalConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualPhysicalConnectionsResult)(nil)).Elem()
}

func (o GetVirtualPhysicalConnectionsResultOutput) ToGetVirtualPhysicalConnectionsResultOutput() GetVirtualPhysicalConnectionsResultOutput {
	return o
}

func (o GetVirtualPhysicalConnectionsResultOutput) ToGetVirtualPhysicalConnectionsResultOutputWithContext(ctx context.Context) GetVirtualPhysicalConnectionsResultOutput {
	return o
}

// The commercial status of the physical line. Value:-**Normal**: activated.-**Financialized**: Arrears locked.-**SecurityLocked**: locked for security reasons.
func (o GetVirtualPhysicalConnectionsResultOutput) BusinessStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.BusinessStatus }).(pulumi.StringPtrOutput)
}

// A list of Virtual Physical Connection Entries. Each element contains the following attributes:
func (o GetVirtualPhysicalConnectionsResultOutput) Connections() GetVirtualPhysicalConnectionsConnectionArrayOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) []GetVirtualPhysicalConnectionsConnection {
		return v.Connections
	}).(GetVirtualPhysicalConnectionsConnectionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualPhysicalConnectionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Virtual Physical Connection IDs.
func (o GetVirtualPhysicalConnectionsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVirtualPhysicalConnectionsResultOutput) IsConfirmed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *bool { return v.IsConfirmed }).(pulumi.BoolPtrOutput)
}

func (o GetVirtualPhysicalConnectionsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Virtual Physical Connections.
func (o GetVirtualPhysicalConnectionsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVirtualPhysicalConnectionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the instance of the physical connection.
func (o GetVirtualPhysicalConnectionsResultOutput) ParentPhysicalConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.ParentPhysicalConnectionId }).(pulumi.StringPtrOutput)
}

func (o GetVirtualPhysicalConnectionsResultOutput) VirtualPhysicalConnectionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) []string { return v.VirtualPhysicalConnectionIds }).(pulumi.StringArrayOutput)
}

// The business status of the shared line. Value:-**Confirmed**: The shared line has been Confirmed to receive.-**UnConfirmed**: The shared line has not been confirmed to be received.-**Deleted**: The shared line has been Deleted.
func (o GetVirtualPhysicalConnectionsResultOutput) VirtualPhysicalConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.VirtualPhysicalConnectionStatus }).(pulumi.StringPtrOutput)
}

func (o GetVirtualPhysicalConnectionsResultOutput) VlanIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) []int { return v.VlanIds }).(pulumi.IntArrayOutput)
}

// The ID of the Alibaba Cloud account (primary account) of the owner of the shared line.
func (o GetVirtualPhysicalConnectionsResultOutput) VpconnAliUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualPhysicalConnectionsResult) *string { return v.VpconnAliUid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualPhysicalConnectionsResultOutput{})
}

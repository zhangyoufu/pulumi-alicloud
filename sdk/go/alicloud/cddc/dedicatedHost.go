// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cddc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// ApsaraDB for MyBase Dedicated Host can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cddc/dedicatedHost:DedicatedHost example <dedicated_host_group_id>:<dedicated_host_id>
//
// ```
type DedicatedHost struct {
	pulumi.CustomResourceState

	// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
	AllocationStatus pulumi.StringOutput `pulumi:"allocationStatus"`
	// Specifies whether to enable the auto-renewal feature.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The ID of the dedicated cluster.
	DedicatedHostGroupId pulumi.StringOutput `pulumi:"dedicatedHostGroupId"`
	// The ID of the host.
	DedicatedHostId pulumi.StringOutput `pulumi:"dedicatedHostId"`
	// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
	HostClass pulumi.StringOutput `pulumi:"hostClass"`
	// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
	ImageCategory pulumi.StringPtrOutput `pulumi:"imageCategory"`
	// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
	OsPassword pulumi.StringPtrOutput `pulumi:"osPassword"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The subscription duration of the host. Valid values:
	// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
	// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
	// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
	UsedTime pulumi.IntPtrOutput `pulumi:"usedTime"`
	// The ID of the vSwitch to which the host is connected.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The ID of the zone.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewDedicatedHost registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DedicatedHostGroupId == nil {
		return nil, errors.New("invalid value for required argument 'DedicatedHostGroupId'")
	}
	if args.HostClass == nil {
		return nil, errors.New("invalid value for required argument 'HostClass'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	if args.OsPassword != nil {
		args.OsPassword = pulumi.ToSecret(args.OsPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"osPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedHost
	err := ctx.RegisterResource("alicloud:cddc/dedicatedHost:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHost gets an existing DedicatedHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("alicloud:cddc/dedicatedHost:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHost resources.
type dedicatedHostState struct {
	// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
	AllocationStatus *string `pulumi:"allocationStatus"`
	// Specifies whether to enable the auto-renewal feature.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The ID of the dedicated cluster.
	DedicatedHostGroupId *string `pulumi:"dedicatedHostGroupId"`
	// The ID of the host.
	DedicatedHostId *string `pulumi:"dedicatedHostId"`
	// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
	HostClass *string `pulumi:"hostClass"`
	// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	HostName *string `pulumi:"hostName"`
	// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
	ImageCategory *string `pulumi:"imageCategory"`
	// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
	OsPassword *string `pulumi:"osPassword"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
	Period *string `pulumi:"period"`
	// The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The subscription duration of the host. Valid values:
	// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
	// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
	// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
	UsedTime *int `pulumi:"usedTime"`
	// The ID of the vSwitch to which the host is connected.
	VswitchId *string `pulumi:"vswitchId"`
	// The ID of the zone.
	ZoneId *string `pulumi:"zoneId"`
}

type DedicatedHostState struct {
	// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
	AllocationStatus pulumi.StringPtrInput
	// Specifies whether to enable the auto-renewal feature.
	AutoRenew pulumi.BoolPtrInput
	// The ID of the dedicated cluster.
	DedicatedHostGroupId pulumi.StringPtrInput
	// The ID of the host.
	DedicatedHostId pulumi.StringPtrInput
	// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
	HostClass pulumi.StringPtrInput
	// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	HostName pulumi.StringPtrInput
	// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
	ImageCategory pulumi.StringPtrInput
	// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
	OsPassword pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
	Period pulumi.StringPtrInput
	// The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The subscription duration of the host. Valid values:
	// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
	// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
	// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
	UsedTime pulumi.IntPtrInput
	// The ID of the vSwitch to which the host is connected.
	VswitchId pulumi.StringPtrInput
	// The ID of the zone.
	ZoneId pulumi.StringPtrInput
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
	AllocationStatus *string `pulumi:"allocationStatus"`
	// Specifies whether to enable the auto-renewal feature.
	AutoRenew *bool `pulumi:"autoRenew"`
	// The ID of the dedicated cluster.
	DedicatedHostGroupId string `pulumi:"dedicatedHostGroupId"`
	// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
	HostClass string `pulumi:"hostClass"`
	// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	HostName *string `pulumi:"hostName"`
	// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
	ImageCategory *string `pulumi:"imageCategory"`
	// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
	OsPassword *string `pulumi:"osPassword"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
	Period *string `pulumi:"period"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The subscription duration of the host. Valid values:
	// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
	// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
	// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
	UsedTime *int `pulumi:"usedTime"`
	// The ID of the vSwitch to which the host is connected.
	VswitchId string `pulumi:"vswitchId"`
	// The ID of the zone.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a DedicatedHost resource.
type DedicatedHostArgs struct {
	// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
	AllocationStatus pulumi.StringPtrInput
	// Specifies whether to enable the auto-renewal feature.
	AutoRenew pulumi.BoolPtrInput
	// The ID of the dedicated cluster.
	DedicatedHostGroupId pulumi.StringInput
	// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
	HostClass pulumi.StringInput
	// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	HostName pulumi.StringPtrInput
	// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
	ImageCategory pulumi.StringPtrInput
	// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
	OsPassword pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringInput
	// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
	Period pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The subscription duration of the host. Valid values:
	// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
	// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
	// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
	UsedTime pulumi.IntPtrInput
	// The ID of the vSwitch to which the host is connected.
	VswitchId pulumi.StringInput
	// The ID of the zone.
	ZoneId pulumi.StringInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}

type DedicatedHostInput interface {
	pulumi.Input

	ToDedicatedHostOutput() DedicatedHostOutput
	ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput
}

func (*DedicatedHost) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (i *DedicatedHost) ToDedicatedHostOutput() DedicatedHostOutput {
	return i.ToDedicatedHostOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostOutput)
}

// DedicatedHostArrayInput is an input type that accepts DedicatedHostArray and DedicatedHostArrayOutput values.
// You can construct a concrete instance of `DedicatedHostArrayInput` via:
//
//	DedicatedHostArray{ DedicatedHostArgs{...} }
type DedicatedHostArrayInput interface {
	pulumi.Input

	ToDedicatedHostArrayOutput() DedicatedHostArrayOutput
	ToDedicatedHostArrayOutputWithContext(context.Context) DedicatedHostArrayOutput
}

type DedicatedHostArray []DedicatedHostInput

func (DedicatedHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHost)(nil)).Elem()
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return i.ToDedicatedHostArrayOutputWithContext(context.Background())
}

func (i DedicatedHostArray) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostArrayOutput)
}

// DedicatedHostMapInput is an input type that accepts DedicatedHostMap and DedicatedHostMapOutput values.
// You can construct a concrete instance of `DedicatedHostMapInput` via:
//
//	DedicatedHostMap{ "key": DedicatedHostArgs{...} }
type DedicatedHostMapInput interface {
	pulumi.Input

	ToDedicatedHostMapOutput() DedicatedHostMapOutput
	ToDedicatedHostMapOutputWithContext(context.Context) DedicatedHostMapOutput
}

type DedicatedHostMap map[string]DedicatedHostInput

func (DedicatedHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHost)(nil)).Elem()
}

func (i DedicatedHostMap) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return i.ToDedicatedHostMapOutputWithContext(context.Background())
}

func (i DedicatedHostMap) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostMapOutput)
}

type DedicatedHostOutput struct{ *pulumi.OutputState }

func (DedicatedHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostOutput) ToDedicatedHostOutput() DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return o
}

// Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
func (o DedicatedHostOutput) AllocationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.AllocationStatus }).(pulumi.StringOutput)
}

// Specifies whether to enable the auto-renewal feature.
func (o DedicatedHostOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The ID of the dedicated cluster.
func (o DedicatedHostOutput) DedicatedHostGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.DedicatedHostGroupId }).(pulumi.StringOutput)
}

// The ID of the host.
func (o DedicatedHostOutput) DedicatedHostId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.DedicatedHostId }).(pulumi.StringOutput)
}

// The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
func (o DedicatedHostOutput) HostClass() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.HostClass }).(pulumi.StringOutput)
}

// The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
func (o DedicatedHostOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
func (o DedicatedHostOutput) ImageCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.ImageCategory }).(pulumi.StringPtrOutput)
}

// Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
func (o DedicatedHostOutput) OsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.OsPassword }).(pulumi.StringPtrOutput)
}

// The payment type of the resource. Valid values: `Subscription`.
func (o DedicatedHostOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
func (o DedicatedHostOutput) Period() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.Period }).(pulumi.StringPtrOutput)
}

// The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
func (o DedicatedHostOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o DedicatedHostOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The subscription duration of the host. Valid values:
// * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
// * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
// * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
func (o DedicatedHostOutput) UsedTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.IntPtrOutput { return v.UsedTime }).(pulumi.IntPtrOutput)
}

// The ID of the vSwitch to which the host is connected.
func (o DedicatedHostOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The ID of the zone.
func (o DedicatedHostOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type DedicatedHostArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutput() DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) ToDedicatedHostArrayOutputWithContext(ctx context.Context) DedicatedHostArrayOutput {
	return o
}

func (o DedicatedHostArrayOutput) Index(i pulumi.IntInput) DedicatedHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedHost {
		return vs[0].([]*DedicatedHost)[vs[1].(int)]
	}).(DedicatedHostOutput)
}

type DedicatedHostMapOutput struct{ *pulumi.OutputState }

func (DedicatedHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutput() DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) ToDedicatedHostMapOutputWithContext(ctx context.Context) DedicatedHostMapOutput {
	return o
}

func (o DedicatedHostMapOutput) MapIndex(k pulumi.StringInput) DedicatedHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedHost {
		return vs[0].(map[string]*DedicatedHost)[vs[1].(string)]
	}).(DedicatedHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostInput)(nil)).Elem(), &DedicatedHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostArrayInput)(nil)).Elem(), DedicatedHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedHostMapInput)(nil)).Elem(), DedicatedHostMap{})
	pulumi.RegisterOutputType(DedicatedHostOutput{})
	pulumi.RegisterOutputType(DedicatedHostArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostMapOutput{})
}

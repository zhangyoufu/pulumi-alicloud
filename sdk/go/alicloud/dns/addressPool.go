// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Alidns Address Pool resource.
//
// For information about Alidns Address Pool and how to use it, see [What is Address Pool](https://www.alibabacloud.com/help/doc-detail/189621.html).
//
// > **NOTE:** Available since v1.152.0.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dns"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			domainName := "alicloud-provider.com"
//			if param := cfg.Get("domainName"); param != "" {
//				domainName = param
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultAlarmContactGroup, err := cms.NewAlarmContactGroup(ctx, "defaultAlarmContactGroup", &cms.AlarmContactGroupArgs{
//				AlarmContactGroupName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGtmInstance, err := dns.NewGtmInstance(ctx, "defaultGtmInstance", &dns.GtmInstanceArgs{
//				InstanceName:         pulumi.String(name),
//				PaymentType:          pulumi.String("Subscription"),
//				Period:               pulumi.Int(1),
//				RenewalStatus:        pulumi.String("ManualRenewal"),
//				PackageEdition:       pulumi.String("standard"),
//				HealthCheckTaskCount: pulumi.Int(100),
//				SmsNotificationCount: pulumi.Int(1000),
//				PublicCnameMode:      pulumi.String("SYSTEM_ASSIGN"),
//				Ttl:                  pulumi.Int(60),
//				CnameType:            pulumi.String("PUBLIC"),
//				ResourceGroupId:      *pulumi.String(defaultResourceGroups.Groups[0].Id),
//				AlertGroups: pulumi.StringArray{
//					defaultAlarmContactGroup.AlarmContactGroupName,
//				},
//				PublicUserDomainName: pulumi.String(domainName),
//				AlertConfigs: dns.GtmInstanceAlertConfigArray{
//					&dns.GtmInstanceAlertConfigArgs{
//						SmsNotice:      pulumi.Bool(true),
//						NoticeType:     pulumi.String("ADDR_ALERT"),
//						EmailNotice:    pulumi.Bool(true),
//						DingtalkNotice: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewAddressPool(ctx, "defaultAddressPool", &dns.AddressPoolArgs{
//				AddressPoolName: pulumi.String(name),
//				InstanceId:      defaultGtmInstance.ID(),
//				LbaStrategy:     pulumi.String("RATIO"),
//				Type:            pulumi.String("IPV4"),
//				Addresses: dns.AddressPoolAddressArray{
//					&dns.AddressPoolAddressArgs{
//						AttributeInfo: pulumi.String("{\"lineCodeRectifyType\":\"RECTIFIED\",\"lineCodes\":[\"os_namerica_us\"]}"),
//						Remark:        pulumi.String("address_remark"),
//						Address:       pulumi.String("1.1.1.1"),
//						Mode:          pulumi.String("SMART"),
//						LbaWeight:     pulumi.Int(1),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Alidns Address Pool can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dns/addressPool:AddressPool example <id>
//
// ```
type AddressPool struct {
	pulumi.CustomResourceState

	// The name of the address pool.
	AddressPoolName pulumi.StringOutput `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See `address` below for details.
	Addresses AddressPoolAddressArrayOutput `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringOutput `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddressPool registers a new resource with the given unique name, arguments, and options.
func NewAddressPool(ctx *pulumi.Context,
	name string, args *AddressPoolArgs, opts ...pulumi.ResourceOption) (*AddressPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressPoolName == nil {
		return nil, errors.New("invalid value for required argument 'AddressPoolName'")
	}
	if args.Addresses == nil {
		return nil, errors.New("invalid value for required argument 'Addresses'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.LbaStrategy == nil {
		return nil, errors.New("invalid value for required argument 'LbaStrategy'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AddressPool
	err := ctx.RegisterResource("alicloud:dns/addressPool:AddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressPool gets an existing AddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressPoolState, opts ...pulumi.ResourceOption) (*AddressPool, error) {
	var resource AddressPool
	err := ctx.ReadResource("alicloud:dns/addressPool:AddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressPool resources.
type addressPoolState struct {
	// The name of the address pool.
	AddressPoolName *string `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See `address` below for details.
	Addresses []AddressPoolAddress `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy *string `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type *string `pulumi:"type"`
}

type AddressPoolState struct {
	// The name of the address pool.
	AddressPoolName pulumi.StringPtrInput
	// The address lists of the Address Pool. See `address` below for details.
	Addresses AddressPoolAddressArrayInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringPtrInput
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringPtrInput
}

func (AddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressPoolState)(nil)).Elem()
}

type addressPoolArgs struct {
	// The name of the address pool.
	AddressPoolName string `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See `address` below for details.
	Addresses []AddressPoolAddress `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy string `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AddressPool resource.
type AddressPoolArgs struct {
	// The name of the address pool.
	AddressPoolName pulumi.StringInput
	// The address lists of the Address Pool. See `address` below for details.
	Addresses AddressPoolAddressArrayInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringInput
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringInput
}

func (AddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressPoolArgs)(nil)).Elem()
}

type AddressPoolInput interface {
	pulumi.Input

	ToAddressPoolOutput() AddressPoolOutput
	ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput
}

func (*AddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPool)(nil)).Elem()
}

func (i *AddressPool) ToAddressPoolOutput() AddressPoolOutput {
	return i.ToAddressPoolOutputWithContext(context.Background())
}

func (i *AddressPool) ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolOutput)
}

func (i *AddressPool) ToOutput(ctx context.Context) pulumix.Output[*AddressPool] {
	return pulumix.Output[*AddressPool]{
		OutputState: i.ToAddressPoolOutputWithContext(ctx).OutputState,
	}
}

// AddressPoolArrayInput is an input type that accepts AddressPoolArray and AddressPoolArrayOutput values.
// You can construct a concrete instance of `AddressPoolArrayInput` via:
//
//	AddressPoolArray{ AddressPoolArgs{...} }
type AddressPoolArrayInput interface {
	pulumi.Input

	ToAddressPoolArrayOutput() AddressPoolArrayOutput
	ToAddressPoolArrayOutputWithContext(context.Context) AddressPoolArrayOutput
}

type AddressPoolArray []AddressPoolInput

func (AddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressPool)(nil)).Elem()
}

func (i AddressPoolArray) ToAddressPoolArrayOutput() AddressPoolArrayOutput {
	return i.ToAddressPoolArrayOutputWithContext(context.Background())
}

func (i AddressPoolArray) ToAddressPoolArrayOutputWithContext(ctx context.Context) AddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolArrayOutput)
}

func (i AddressPoolArray) ToOutput(ctx context.Context) pulumix.Output[[]*AddressPool] {
	return pulumix.Output[[]*AddressPool]{
		OutputState: i.ToAddressPoolArrayOutputWithContext(ctx).OutputState,
	}
}

// AddressPoolMapInput is an input type that accepts AddressPoolMap and AddressPoolMapOutput values.
// You can construct a concrete instance of `AddressPoolMapInput` via:
//
//	AddressPoolMap{ "key": AddressPoolArgs{...} }
type AddressPoolMapInput interface {
	pulumi.Input

	ToAddressPoolMapOutput() AddressPoolMapOutput
	ToAddressPoolMapOutputWithContext(context.Context) AddressPoolMapOutput
}

type AddressPoolMap map[string]AddressPoolInput

func (AddressPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressPool)(nil)).Elem()
}

func (i AddressPoolMap) ToAddressPoolMapOutput() AddressPoolMapOutput {
	return i.ToAddressPoolMapOutputWithContext(context.Background())
}

func (i AddressPoolMap) ToAddressPoolMapOutputWithContext(ctx context.Context) AddressPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolMapOutput)
}

func (i AddressPoolMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AddressPool] {
	return pulumix.Output[map[string]*AddressPool]{
		OutputState: i.ToAddressPoolMapOutputWithContext(ctx).OutputState,
	}
}

type AddressPoolOutput struct{ *pulumi.OutputState }

func (AddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPool)(nil)).Elem()
}

func (o AddressPoolOutput) ToAddressPoolOutput() AddressPoolOutput {
	return o
}

func (o AddressPoolOutput) ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput {
	return o
}

func (o AddressPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*AddressPool] {
	return pulumix.Output[*AddressPool]{
		OutputState: o.OutputState,
	}
}

// The name of the address pool.
func (o AddressPoolOutput) AddressPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.AddressPoolName }).(pulumi.StringOutput)
}

// The address lists of the Address Pool. See `address` below for details.
func (o AddressPoolOutput) Addresses() AddressPoolAddressArrayOutput {
	return o.ApplyT(func(v *AddressPool) AddressPoolAddressArrayOutput { return v.Addresses }).(AddressPoolAddressArrayOutput)
}

// The ID of the instance.
func (o AddressPoolOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
func (o AddressPoolOutput) LbaStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.LbaStrategy }).(pulumi.StringOutput)
}

// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
func (o AddressPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AddressPoolArrayOutput struct{ *pulumi.OutputState }

func (AddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressPool)(nil)).Elem()
}

func (o AddressPoolArrayOutput) ToAddressPoolArrayOutput() AddressPoolArrayOutput {
	return o
}

func (o AddressPoolArrayOutput) ToAddressPoolArrayOutputWithContext(ctx context.Context) AddressPoolArrayOutput {
	return o
}

func (o AddressPoolArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AddressPool] {
	return pulumix.Output[[]*AddressPool]{
		OutputState: o.OutputState,
	}
}

func (o AddressPoolArrayOutput) Index(i pulumi.IntInput) AddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AddressPool {
		return vs[0].([]*AddressPool)[vs[1].(int)]
	}).(AddressPoolOutput)
}

type AddressPoolMapOutput struct{ *pulumi.OutputState }

func (AddressPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressPool)(nil)).Elem()
}

func (o AddressPoolMapOutput) ToAddressPoolMapOutput() AddressPoolMapOutput {
	return o
}

func (o AddressPoolMapOutput) ToAddressPoolMapOutputWithContext(ctx context.Context) AddressPoolMapOutput {
	return o
}

func (o AddressPoolMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AddressPool] {
	return pulumix.Output[map[string]*AddressPool]{
		OutputState: o.OutputState,
	}
}

func (o AddressPoolMapOutput) MapIndex(k pulumi.StringInput) AddressPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AddressPool {
		return vs[0].(map[string]*AddressPool)[vs[1].(string)]
	}).(AddressPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolInput)(nil)).Elem(), &AddressPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolArrayInput)(nil)).Elem(), AddressPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolMapInput)(nil)).Elem(), AddressPoolMap{})
	pulumi.RegisterOutputType(AddressPoolOutput{})
	pulumi.RegisterOutputType(AddressPoolArrayOutput{})
	pulumi.RegisterOutputType(AddressPoolMapOutput{})
}

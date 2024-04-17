// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB endpoint address resource to allocate an Internet endpoint address string for PolarDB instance.
//
// > **NOTE:** Available since v1.68.0. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
//
//	To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := polardb.GetNodeClasses(ctx, &polardb.GetNodeClassesArgs{
//				DbType:    pulumi.StringRef("MySQL"),
//				DbVersion: pulumi.StringRef("8.0"),
//				PayType:   "PostPaid",
//				Category:  pulumi.StringRef("Normal"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      pulumi.String(_default.Classes[0].ZoneId),
//				VswitchName: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultCluster, err := polardb.NewCluster(ctx, "default", &polardb.ClusterArgs{
//				DbType:      pulumi.String("MySQL"),
//				DbVersion:   pulumi.String("8.0"),
//				DbNodeClass: pulumi.String(_default.Classes[0].SupportedEngines[0].AvailableResources[0].DbNodeClass),
//				PayType:     pulumi.String("PostPaid"),
//				VswitchId:   defaultSwitch.ID(),
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGetEndpoints := polardb.GetEndpointsOutput(ctx, polardb.GetEndpointsOutputArgs{
//				DbClusterId: defaultCluster.ID(),
//			}, nil)
//			_, err = polardb.NewEndpointAddress(ctx, "default", &polardb.EndpointAddressArgs{
//				DbClusterId: defaultCluster.ID(),
//				DbEndpointId: defaultGetEndpoints.ApplyT(func(defaultGetEndpoints polardb.GetEndpointsResult) (*string, error) {
//					return &defaultGetEndpoints.Endpoints[0].DbEndpointId, nil
//				}).(pulumi.StringPtrOutput),
//				ConnectionPrefix: pulumi.String("polardbexample"),
//				NetType:          pulumi.String("Public"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// PolarDB endpoint address can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:polardb/endpointAddress:EndpointAddress example pc-abc123456:pe-abc123456
// ```
type EndpointAddress struct {
	pulumi.CustomResourceState

	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringOutput `pulumi:"connectionPrefix"`
	// Connection cluster or endpoint string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringOutput `pulumi:"dbEndpointId"`
	// The ip address of connection string.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrOutput `pulumi:"netType"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringOutput `pulumi:"port"`
}

// NewEndpointAddress registers a new resource with the given unique name, arguments, and options.
func NewEndpointAddress(ctx *pulumi.Context,
	name string, args *EndpointAddressArgs, opts ...pulumi.ResourceOption) (*EndpointAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	if args.DbEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'DbEndpointId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointAddress
	err := ctx.RegisterResource("alicloud:polardb/endpointAddress:EndpointAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAddress gets an existing EndpointAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAddressState, opts ...pulumi.ResourceOption) (*EndpointAddress, error) {
	var resource EndpointAddress
	err := ctx.ReadResource("alicloud:polardb/endpointAddress:EndpointAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAddress resources.
type endpointAddressState struct {
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Connection cluster or endpoint string.
	ConnectionString *string `pulumi:"connectionString"`
	// The Id of cluster that can run database.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId *string `pulumi:"dbEndpointId"`
	// The ip address of connection string.
	IpAddress *string `pulumi:"ipAddress"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType *string `pulumi:"netType"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port *string `pulumi:"port"`
}

type EndpointAddressState struct {
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringPtrInput
	// Connection cluster or endpoint string.
	ConnectionString pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringPtrInput
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringPtrInput
	// The ip address of connection string.
	IpAddress pulumi.StringPtrInput
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrInput
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringPtrInput
}

func (EndpointAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAddressState)(nil)).Elem()
}

type endpointAddressArgs struct {
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId string `pulumi:"dbClusterId"`
	// The Id of endpoint that can run database.
	DbEndpointId string `pulumi:"dbEndpointId"`
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType *string `pulumi:"netType"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port *string `pulumi:"port"`
}

// The set of arguments for constructing a EndpointAddress resource.
type EndpointAddressArgs struct {
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringInput
	// The Id of endpoint that can run database.
	DbEndpointId pulumi.StringInput
	// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
	NetType pulumi.StringPtrInput
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringPtrInput
}

func (EndpointAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAddressArgs)(nil)).Elem()
}

type EndpointAddressInput interface {
	pulumi.Input

	ToEndpointAddressOutput() EndpointAddressOutput
	ToEndpointAddressOutputWithContext(ctx context.Context) EndpointAddressOutput
}

func (*EndpointAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAddress)(nil)).Elem()
}

func (i *EndpointAddress) ToEndpointAddressOutput() EndpointAddressOutput {
	return i.ToEndpointAddressOutputWithContext(context.Background())
}

func (i *EndpointAddress) ToEndpointAddressOutputWithContext(ctx context.Context) EndpointAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAddressOutput)
}

// EndpointAddressArrayInput is an input type that accepts EndpointAddressArray and EndpointAddressArrayOutput values.
// You can construct a concrete instance of `EndpointAddressArrayInput` via:
//
//	EndpointAddressArray{ EndpointAddressArgs{...} }
type EndpointAddressArrayInput interface {
	pulumi.Input

	ToEndpointAddressArrayOutput() EndpointAddressArrayOutput
	ToEndpointAddressArrayOutputWithContext(context.Context) EndpointAddressArrayOutput
}

type EndpointAddressArray []EndpointAddressInput

func (EndpointAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAddress)(nil)).Elem()
}

func (i EndpointAddressArray) ToEndpointAddressArrayOutput() EndpointAddressArrayOutput {
	return i.ToEndpointAddressArrayOutputWithContext(context.Background())
}

func (i EndpointAddressArray) ToEndpointAddressArrayOutputWithContext(ctx context.Context) EndpointAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAddressArrayOutput)
}

// EndpointAddressMapInput is an input type that accepts EndpointAddressMap and EndpointAddressMapOutput values.
// You can construct a concrete instance of `EndpointAddressMapInput` via:
//
//	EndpointAddressMap{ "key": EndpointAddressArgs{...} }
type EndpointAddressMapInput interface {
	pulumi.Input

	ToEndpointAddressMapOutput() EndpointAddressMapOutput
	ToEndpointAddressMapOutputWithContext(context.Context) EndpointAddressMapOutput
}

type EndpointAddressMap map[string]EndpointAddressInput

func (EndpointAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAddress)(nil)).Elem()
}

func (i EndpointAddressMap) ToEndpointAddressMapOutput() EndpointAddressMapOutput {
	return i.ToEndpointAddressMapOutputWithContext(context.Background())
}

func (i EndpointAddressMap) ToEndpointAddressMapOutputWithContext(ctx context.Context) EndpointAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAddressMapOutput)
}

type EndpointAddressOutput struct{ *pulumi.OutputState }

func (EndpointAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAddress)(nil)).Elem()
}

func (o EndpointAddressOutput) ToEndpointAddressOutput() EndpointAddressOutput {
	return o
}

func (o EndpointAddressOutput) ToEndpointAddressOutputWithContext(ctx context.Context) EndpointAddressOutput {
	return o
}

// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
func (o EndpointAddressOutput) ConnectionPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.ConnectionPrefix }).(pulumi.StringOutput)
}

// Connection cluster or endpoint string.
func (o EndpointAddressOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// The Id of cluster that can run database.
func (o EndpointAddressOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// The Id of endpoint that can run database.
func (o EndpointAddressOutput) DbEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.DbEndpointId }).(pulumi.StringOutput)
}

// The ip address of connection string.
func (o EndpointAddressOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
func (o EndpointAddressOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringPtrOutput { return v.NetType }).(pulumi.StringPtrOutput)
}

// Port of the specified endpoint. Valid values: 3000 to 5999.
func (o EndpointAddressOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAddress) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

type EndpointAddressArrayOutput struct{ *pulumi.OutputState }

func (EndpointAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAddress)(nil)).Elem()
}

func (o EndpointAddressArrayOutput) ToEndpointAddressArrayOutput() EndpointAddressArrayOutput {
	return o
}

func (o EndpointAddressArrayOutput) ToEndpointAddressArrayOutputWithContext(ctx context.Context) EndpointAddressArrayOutput {
	return o
}

func (o EndpointAddressArrayOutput) Index(i pulumi.IntInput) EndpointAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointAddress {
		return vs[0].([]*EndpointAddress)[vs[1].(int)]
	}).(EndpointAddressOutput)
}

type EndpointAddressMapOutput struct{ *pulumi.OutputState }

func (EndpointAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAddress)(nil)).Elem()
}

func (o EndpointAddressMapOutput) ToEndpointAddressMapOutput() EndpointAddressMapOutput {
	return o
}

func (o EndpointAddressMapOutput) ToEndpointAddressMapOutputWithContext(ctx context.Context) EndpointAddressMapOutput {
	return o
}

func (o EndpointAddressMapOutput) MapIndex(k pulumi.StringInput) EndpointAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointAddress {
		return vs[0].(map[string]*EndpointAddress)[vs[1].(string)]
	}).(EndpointAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAddressInput)(nil)).Elem(), &EndpointAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAddressArrayInput)(nil)).Elem(), EndpointAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAddressMapInput)(nil)).Elem(), EndpointAddressMap{})
	pulumi.RegisterOutputType(EndpointAddressOutput{})
	pulumi.RegisterOutputType(EndpointAddressArrayOutput{})
	pulumi.RegisterOutputType(EndpointAddressMapOutput{})
}

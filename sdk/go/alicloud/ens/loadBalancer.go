// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ens

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ENS Load Balancer resource. Load balancing. When you use it for the first time, please contact the product classmates to add a resource whitelist.
//
// For information about ENS Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createloadbalancer).
//
// > **NOTE:** Available since v1.213.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ens"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			network, err := ens.NewNetwork(ctx, "network", &ens.NetworkArgs{
//				NetworkName: pulumi.String(name),
//				Description: pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.2.0/24"),
//				EnsRegionId: pulumi.String("cn-chenzhou-telecom_unicom_cmcc"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ens.NewVswitch(ctx, "switch", &ens.VswitchArgs{
//				Description: pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.2.0/24"),
//				VswitchName: pulumi.String(name),
//				EnsRegionId: pulumi.String("cn-chenzhou-telecom_unicom_cmcc"),
//				NetworkId:   network.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ens.NewLoadBalancer(ctx, "default", &ens.LoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				PaymentType:      pulumi.String("PayAsYouGo"),
//				EnsRegionId:      pulumi.String("cn-chenzhou-telecom_unicom_cmcc"),
//				LoadBalancerSpec: pulumi.String("elb.s1.small"),
//				VswitchId:        _switch.ID(),
//				NetworkId:        network.ID(),
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
// ENS Load Balancer can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ens/loadBalancer:LoadBalancer example <id>
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// The creation Time (UTC) of the load balancing instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The ID of the ENS node.
	EnsRegionId pulumi.StringOutput `pulumi:"ensRegionId"`
	// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
	LoadBalancerName pulumi.StringPtrOutput `pulumi:"loadBalancerName"`
	// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
	LoadBalancerSpec pulumi.StringOutput `pulumi:"loadBalancerSpec"`
	// The network ID of the created edge load balancing (ELB) instance.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The status of the SLB instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the vSwitch to which the VPC instance belongs.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnsRegionId == nil {
		return nil, errors.New("invalid value for required argument 'EnsRegionId'")
	}
	if args.LoadBalancerSpec == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerSpec'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("alicloud:ens/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("alicloud:ens/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// The creation Time (UTC) of the load balancing instance.
	CreateTime *string `pulumi:"createTime"`
	// The ID of the ENS node.
	EnsRegionId *string `pulumi:"ensRegionId"`
	// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
	LoadBalancerSpec *string `pulumi:"loadBalancerSpec"`
	// The network ID of the created edge load balancing (ELB) instance.
	NetworkId *string `pulumi:"networkId"`
	// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
	PaymentType *string `pulumi:"paymentType"`
	// The status of the SLB instance.
	Status *string `pulumi:"status"`
	// The ID of the vSwitch to which the VPC instance belongs.
	VswitchId *string `pulumi:"vswitchId"`
}

type LoadBalancerState struct {
	// The creation Time (UTC) of the load balancing instance.
	CreateTime pulumi.StringPtrInput
	// The ID of the ENS node.
	EnsRegionId pulumi.StringPtrInput
	// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
	LoadBalancerName pulumi.StringPtrInput
	// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
	LoadBalancerSpec pulumi.StringPtrInput
	// The network ID of the created edge load balancing (ELB) instance.
	NetworkId pulumi.StringPtrInput
	// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
	PaymentType pulumi.StringPtrInput
	// The status of the SLB instance.
	Status pulumi.StringPtrInput
	// The ID of the vSwitch to which the VPC instance belongs.
	VswitchId pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// The ID of the ENS node.
	EnsRegionId string `pulumi:"ensRegionId"`
	// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
	LoadBalancerSpec string `pulumi:"loadBalancerSpec"`
	// The network ID of the created edge load balancing (ELB) instance.
	NetworkId string `pulumi:"networkId"`
	// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
	PaymentType string `pulumi:"paymentType"`
	// The ID of the vSwitch to which the VPC instance belongs.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// The ID of the ENS node.
	EnsRegionId pulumi.StringInput
	// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
	LoadBalancerName pulumi.StringPtrInput
	// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
	LoadBalancerSpec pulumi.StringInput
	// The network ID of the created edge load balancing (ELB) instance.
	NetworkId pulumi.StringInput
	// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
	PaymentType pulumi.StringInput
	// The ID of the vSwitch to which the VPC instance belongs.
	VswitchId pulumi.StringInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

// LoadBalancerArrayInput is an input type that accepts LoadBalancerArray and LoadBalancerArrayOutput values.
// You can construct a concrete instance of `LoadBalancerArrayInput` via:
//
//	LoadBalancerArray{ LoadBalancerArgs{...} }
type LoadBalancerArrayInput interface {
	pulumi.Input

	ToLoadBalancerArrayOutput() LoadBalancerArrayOutput
	ToLoadBalancerArrayOutputWithContext(context.Context) LoadBalancerArrayOutput
}

type LoadBalancerArray []LoadBalancerInput

func (LoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return i.ToLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerArrayOutput)
}

// LoadBalancerMapInput is an input type that accepts LoadBalancerMap and LoadBalancerMapOutput values.
// You can construct a concrete instance of `LoadBalancerMapInput` via:
//
//	LoadBalancerMap{ "key": LoadBalancerArgs{...} }
type LoadBalancerMapInput interface {
	pulumi.Input

	ToLoadBalancerMapOutput() LoadBalancerMapOutput
	ToLoadBalancerMapOutputWithContext(context.Context) LoadBalancerMapOutput
}

type LoadBalancerMap map[string]LoadBalancerInput

func (LoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerMap) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return i.ToLoadBalancerMapOutputWithContext(context.Background())
}

func (i LoadBalancerMap) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerMapOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// The creation Time (UTC) of the load balancing instance.
func (o LoadBalancerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of the ENS node.
func (o LoadBalancerOutput) EnsRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.EnsRegionId }).(pulumi.StringOutput)
}

// Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
func (o LoadBalancerOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

// Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
func (o LoadBalancerOutput) LoadBalancerSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerSpec }).(pulumi.StringOutput)
}

// The network ID of the created edge load balancing (ELB) instance.
func (o LoadBalancerOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
func (o LoadBalancerOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The status of the SLB instance.
func (o LoadBalancerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the vSwitch to which the VPC instance belongs.
func (o LoadBalancerOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type LoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) Index(i pulumi.IntInput) LoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].([]*LoadBalancer)[vs[1].(int)]
	}).(LoadBalancerOutput)
}

type LoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].(map[string]*LoadBalancer)[vs[1].(string)]
	}).(LoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerArrayInput)(nil)).Elem(), LoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerMapInput)(nil)).Elem(), LoadBalancerMap{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
	pulumi.RegisterOutputType(LoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerMapOutput{})
}

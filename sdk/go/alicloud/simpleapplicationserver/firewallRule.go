// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleapplicationserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Simple Application Server Firewall Rule resource.
//
// For information about Simple Application Server Firewall Rule and how to use it, see [What is Firewall Rule](https://www.alibabacloud.com/help/doc-detail/190449.htm).
//
// > **NOTE:** Available since v1.143.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/simpleapplicationserver"
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
//			defaultImages, err := simpleapplicationserver.GetImages(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultServerPlans, err := simpleapplicationserver.GetServerPlans(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := simpleapplicationserver.NewInstance(ctx, "defaultInstance", &simpleapplicationserver.InstanceArgs{
//				PaymentType:  pulumi.String("Subscription"),
//				PlanId:       *pulumi.String(defaultServerPlans.Plans[0].Id),
//				InstanceName: pulumi.String(name),
//				ImageId:      *pulumi.String(defaultImages.Images[0].Id),
//				Period:       pulumi.Int(1),
//				DataDiskSize: pulumi.Int(100),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = simpleapplicationserver.NewFirewallRule(ctx, "defaultFirewallRule", &simpleapplicationserver.FirewallRuleArgs{
//				InstanceId:   defaultInstance.ID(),
//				RuleProtocol: pulumi.String("Tcp"),
//				Port:         pulumi.String("9999"),
//				Remark:       pulumi.String(name),
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
// Simple Application Server Firewall Rule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:simpleapplicationserver/firewallRule:FirewallRule example <instance_id>:<firewall_rule_id>
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// The ID of the firewall rule.
	FirewallRuleId pulumi.StringOutput `pulumi:"firewallRuleId"`
	// Alibaba Cloud simple application server instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
	Port pulumi.StringOutput `pulumi:"port"`
	// The remarks of the firewall rule.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
	RuleProtocol pulumi.StringOutput `pulumi:"ruleProtocol"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.RuleProtocol == nil {
		return nil, errors.New("invalid value for required argument 'RuleProtocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallRule
	err := ctx.RegisterResource("alicloud:simpleapplicationserver/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("alicloud:simpleapplicationserver/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The ID of the firewall rule.
	FirewallRuleId *string `pulumi:"firewallRuleId"`
	// Alibaba Cloud simple application server instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
	Port *string `pulumi:"port"`
	// The remarks of the firewall rule.
	Remark *string `pulumi:"remark"`
	// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
	RuleProtocol *string `pulumi:"ruleProtocol"`
}

type FirewallRuleState struct {
	// The ID of the firewall rule.
	FirewallRuleId pulumi.StringPtrInput
	// Alibaba Cloud simple application server instance ID.
	InstanceId pulumi.StringPtrInput
	// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
	Port pulumi.StringPtrInput
	// The remarks of the firewall rule.
	Remark pulumi.StringPtrInput
	// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
	RuleProtocol pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// Alibaba Cloud simple application server instance ID.
	InstanceId string `pulumi:"instanceId"`
	// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
	Port string `pulumi:"port"`
	// The remarks of the firewall rule.
	Remark *string `pulumi:"remark"`
	// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
	RuleProtocol string `pulumi:"ruleProtocol"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Alibaba Cloud simple application server instance ID.
	InstanceId pulumi.StringInput
	// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
	Port pulumi.StringInput
	// The remarks of the firewall rule.
	Remark pulumi.StringPtrInput
	// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
	RuleProtocol pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//	FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//	FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

// The ID of the firewall rule.
func (o FirewallRuleOutput) FirewallRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.FirewallRuleId }).(pulumi.StringOutput)
}

// Alibaba Cloud simple application server instance ID.
func (o FirewallRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The port range. Valid values of port numbers: `1` to `65535`. Specify a port range in the format of `<start port number>/<end port number>`. Example: `1024/1055`, which indicates the port range of `1024` through `1055`.
func (o FirewallRuleOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The remarks of the firewall rule.
func (o FirewallRuleOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// The transport layer protocol. Valid values: `Tcp`, `Udp`, `TcpAndUdp`.
func (o FirewallRuleOutput) RuleProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.RuleProtocol }).(pulumi.StringOutput)
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}

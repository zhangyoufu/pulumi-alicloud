// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Traffic Mirror Filter Ingress Rule resource. Traffic mirror entry rule.
//
// For information about VPC Traffic Mirror Filter Ingress Rule and how to use it, see [What is Traffic Mirror Filter Ingress Rule](https://www.alibabacloud.com/help/doc-detail/261357.htm).
//
// > **NOTE:** Available since v1.141.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleTrafficMirrorFilter, err := vpc.NewTrafficMirrorFilter(ctx, "exampleTrafficMirrorFilter", &vpc.TrafficMirrorFilterArgs{
//				TrafficMirrorFilterName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewTrafficMirrorFilterIngressRule(ctx, "exampleTrafficMirrorFilterIngressRule", &vpc.TrafficMirrorFilterIngressRuleArgs{
//				TrafficMirrorFilterId: exampleTrafficMirrorFilter.ID(),
//				Priority:              pulumi.Int(1),
//				RuleAction:            pulumi.String("accept"),
//				Protocol:              pulumi.String("UDP"),
//				DestinationCidrBlock:  pulumi.String("10.0.0.0/24"),
//				SourceCidrBlock:       pulumi.String("10.0.0.0/24"),
//				DestinationPortRange:  pulumi.String("1/120"),
//				SourcePortRange:       pulumi.String("1/120"),
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
// VPC Traffic Mirror Filter Ingress Rule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule example <traffic_mirror_filter_id>:<traffic_mirror_filter_ingress_rule_id>
// ```
type TrafficMirrorFilterIngressRule struct {
	pulumi.CustomResourceState

	// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
	Action pulumi.StringOutput `pulumi:"action"`
	// The destination CIDR block of the inbound traffic.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	DestinationPortRange pulumi.StringOutput `pulumi:"destinationPortRange"`
	// Whether to PreCheck this request only. Value:
	// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
	//
	// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// The source CIDR block of the inbound traffic.
	SourceCidrBlock pulumi.StringOutput `pulumi:"sourceCidrBlock"`
	// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	SourcePortRange pulumi.StringOutput `pulumi:"sourcePortRange"`
	// The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the filter.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
	// The ID of the outbound rule.
	TrafficMirrorFilterIngressRuleId pulumi.StringOutput `pulumi:"trafficMirrorFilterIngressRuleId"`
}

// NewTrafficMirrorFilterIngressRule registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilterIngressRule(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterIngressRuleArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterIngressRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.SourceCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidrBlock'")
	}
	if args.TrafficMirrorFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorFilterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficMirrorFilterIngressRule
	err := ctx.RegisterResource("alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilterIngressRule gets an existing TrafficMirrorFilterIngressRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilterIngressRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterIngressRuleState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterIngressRule, error) {
	var resource TrafficMirrorFilterIngressRule
	err := ctx.ReadResource("alicloud:vpc/trafficMirrorFilterIngressRule:TrafficMirrorFilterIngressRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilterIngressRule resources.
type trafficMirrorFilterIngressRuleState struct {
	// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
	Action *string `pulumi:"action"`
	// The destination CIDR block of the inbound traffic.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Whether to PreCheck this request only. Value:
	// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
	DryRun *bool `pulumi:"dryRun"`
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
	Priority *int `pulumi:"priority"`
	// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
	Protocol *string `pulumi:"protocol"`
	// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
	//
	// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
	RuleAction *string `pulumi:"ruleAction"`
	// The source CIDR block of the inbound traffic.
	SourceCidrBlock *string `pulumi:"sourceCidrBlock"`
	// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
	Status *string `pulumi:"status"`
	// The ID of the filter.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
	// The ID of the outbound rule.
	TrafficMirrorFilterIngressRuleId *string `pulumi:"trafficMirrorFilterIngressRuleId"`
}

type TrafficMirrorFilterIngressRuleState struct {
	// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
	Action pulumi.StringPtrInput
	// The destination CIDR block of the inbound traffic.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	DestinationPortRange pulumi.StringPtrInput
	// Whether to PreCheck this request only. Value:
	// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
	DryRun pulumi.BoolPtrInput
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
	Priority pulumi.IntPtrInput
	// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
	Protocol pulumi.StringPtrInput
	// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
	//
	// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
	RuleAction pulumi.StringPtrInput
	// The source CIDR block of the inbound traffic.
	SourceCidrBlock pulumi.StringPtrInput
	// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	SourcePortRange pulumi.StringPtrInput
	// The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
	Status pulumi.StringPtrInput
	// The ID of the filter.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TrafficMirrorFilterId pulumi.StringPtrInput
	// The ID of the outbound rule.
	TrafficMirrorFilterIngressRuleId pulumi.StringPtrInput
}

func (TrafficMirrorFilterIngressRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterIngressRuleState)(nil)).Elem()
}

type trafficMirrorFilterIngressRuleArgs struct {
	// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
	Action *string `pulumi:"action"`
	// The destination CIDR block of the inbound traffic.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Whether to PreCheck this request only. Value:
	// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
	DryRun *bool `pulumi:"dryRun"`
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
	Priority int `pulumi:"priority"`
	// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
	Protocol string `pulumi:"protocol"`
	// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
	//
	// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
	RuleAction *string `pulumi:"ruleAction"`
	// The source CIDR block of the inbound traffic.
	SourceCidrBlock string `pulumi:"sourceCidrBlock"`
	// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The ID of the filter.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
}

// The set of arguments for constructing a TrafficMirrorFilterIngressRule resource.
type TrafficMirrorFilterIngressRuleArgs struct {
	// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
	Action pulumi.StringPtrInput
	// The destination CIDR block of the inbound traffic.
	DestinationCidrBlock pulumi.StringInput
	// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	DestinationPortRange pulumi.StringPtrInput
	// Whether to PreCheck this request only. Value:
	// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
	DryRun pulumi.BoolPtrInput
	// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
	Priority pulumi.IntInput
	// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
	Protocol pulumi.StringInput
	// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
	//
	// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
	RuleAction pulumi.StringPtrInput
	// The source CIDR block of the inbound traffic.
	SourceCidrBlock pulumi.StringInput
	// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
	SourcePortRange pulumi.StringPtrInput
	// The ID of the filter.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TrafficMirrorFilterId pulumi.StringInput
}

func (TrafficMirrorFilterIngressRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterIngressRuleArgs)(nil)).Elem()
}

type TrafficMirrorFilterIngressRuleInput interface {
	pulumi.Input

	ToTrafficMirrorFilterIngressRuleOutput() TrafficMirrorFilterIngressRuleOutput
	ToTrafficMirrorFilterIngressRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleOutput
}

func (*TrafficMirrorFilterIngressRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (i *TrafficMirrorFilterIngressRule) ToTrafficMirrorFilterIngressRuleOutput() TrafficMirrorFilterIngressRuleOutput {
	return i.ToTrafficMirrorFilterIngressRuleOutputWithContext(context.Background())
}

func (i *TrafficMirrorFilterIngressRule) ToTrafficMirrorFilterIngressRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterIngressRuleOutput)
}

// TrafficMirrorFilterIngressRuleArrayInput is an input type that accepts TrafficMirrorFilterIngressRuleArray and TrafficMirrorFilterIngressRuleArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterIngressRuleArrayInput` via:
//
//	TrafficMirrorFilterIngressRuleArray{ TrafficMirrorFilterIngressRuleArgs{...} }
type TrafficMirrorFilterIngressRuleArrayInput interface {
	pulumi.Input

	ToTrafficMirrorFilterIngressRuleArrayOutput() TrafficMirrorFilterIngressRuleArrayOutput
	ToTrafficMirrorFilterIngressRuleArrayOutputWithContext(context.Context) TrafficMirrorFilterIngressRuleArrayOutput
}

type TrafficMirrorFilterIngressRuleArray []TrafficMirrorFilterIngressRuleInput

func (TrafficMirrorFilterIngressRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (i TrafficMirrorFilterIngressRuleArray) ToTrafficMirrorFilterIngressRuleArrayOutput() TrafficMirrorFilterIngressRuleArrayOutput {
	return i.ToTrafficMirrorFilterIngressRuleArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterIngressRuleArray) ToTrafficMirrorFilterIngressRuleArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterIngressRuleArrayOutput)
}

// TrafficMirrorFilterIngressRuleMapInput is an input type that accepts TrafficMirrorFilterIngressRuleMap and TrafficMirrorFilterIngressRuleMapOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterIngressRuleMapInput` via:
//
//	TrafficMirrorFilterIngressRuleMap{ "key": TrafficMirrorFilterIngressRuleArgs{...} }
type TrafficMirrorFilterIngressRuleMapInput interface {
	pulumi.Input

	ToTrafficMirrorFilterIngressRuleMapOutput() TrafficMirrorFilterIngressRuleMapOutput
	ToTrafficMirrorFilterIngressRuleMapOutputWithContext(context.Context) TrafficMirrorFilterIngressRuleMapOutput
}

type TrafficMirrorFilterIngressRuleMap map[string]TrafficMirrorFilterIngressRuleInput

func (TrafficMirrorFilterIngressRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (i TrafficMirrorFilterIngressRuleMap) ToTrafficMirrorFilterIngressRuleMapOutput() TrafficMirrorFilterIngressRuleMapOutput {
	return i.ToTrafficMirrorFilterIngressRuleMapOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterIngressRuleMap) ToTrafficMirrorFilterIngressRuleMapOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterIngressRuleMapOutput)
}

type TrafficMirrorFilterIngressRuleOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterIngressRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (o TrafficMirrorFilterIngressRuleOutput) ToTrafficMirrorFilterIngressRuleOutput() TrafficMirrorFilterIngressRuleOutput {
	return o
}

func (o TrafficMirrorFilterIngressRuleOutput) ToTrafficMirrorFilterIngressRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleOutput {
	return o
}

// The collection policy of the inbound rule. Valid values: `accept` or `drop`. `accept`: collects network traffic. `drop`: does not collect network traffic.
func (o TrafficMirrorFilterIngressRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The destination CIDR block of the inbound traffic.
func (o TrafficMirrorFilterIngressRuleOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// The destination CIDR block of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
func (o TrafficMirrorFilterIngressRuleOutput) DestinationPortRange() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.DestinationPortRange }).(pulumi.StringOutput)
}

// Whether to PreCheck this request only. Value:
// - **true**: sends a check request and does not create inbound or outbound rules. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
// - **false** (default): Sends a normal request and directly creates an inbound or outbound direction rule after checking.
func (o TrafficMirrorFilterIngressRuleOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The priority of the inbound rule. A smaller value indicates a higher priority. The maximum value is `10`, which indicates that you can configure at most 10 inbound rules for a filter.
func (o TrafficMirrorFilterIngressRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The transport protocol used by inbound traffic that needs to be mirrored. Valid values: `ALL`, `ICMP`, `TCP`, `UDP`.
func (o TrafficMirrorFilterIngressRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// . Field 'rule_action' has been deprecated from provider version 1.211.0. New field 'action' instead.
//
// Deprecated: Field 'rule_action' has been deprecated since provider version 1.211.0. New field 'action' instead.
func (o TrafficMirrorFilterIngressRuleOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// The source CIDR block of the inbound traffic.
func (o TrafficMirrorFilterIngressRuleOutput) SourceCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.SourceCidrBlock }).(pulumi.StringOutput)
}

// The source port range of the inbound traffic. Valid values: `1` to `65535`. Separate the first port and last port with a forward slash (/), for example, `1/200` or `80/80`. A value of `-1/-1` indicates that all ports are available. Therefore, do not set the value to `-1/-1`. **NOTE:** When `protocol` is `ICMP`, this parameter is invalid.
func (o TrafficMirrorFilterIngressRuleOutput) SourcePortRange() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.SourcePortRange }).(pulumi.StringOutput)
}

// The state of the inbound rule. `Creating`, `Created`, `Modifying` and `Deleting`.
func (o TrafficMirrorFilterIngressRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the filter.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o TrafficMirrorFilterIngressRuleOutput) TrafficMirrorFilterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.TrafficMirrorFilterId }).(pulumi.StringOutput)
}

// The ID of the outbound rule.
func (o TrafficMirrorFilterIngressRuleOutput) TrafficMirrorFilterIngressRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterIngressRule) pulumi.StringOutput { return v.TrafficMirrorFilterIngressRuleId }).(pulumi.StringOutput)
}

type TrafficMirrorFilterIngressRuleArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterIngressRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (o TrafficMirrorFilterIngressRuleArrayOutput) ToTrafficMirrorFilterIngressRuleArrayOutput() TrafficMirrorFilterIngressRuleArrayOutput {
	return o
}

func (o TrafficMirrorFilterIngressRuleArrayOutput) ToTrafficMirrorFilterIngressRuleArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleArrayOutput {
	return o
}

func (o TrafficMirrorFilterIngressRuleArrayOutput) Index(i pulumi.IntInput) TrafficMirrorFilterIngressRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorFilterIngressRule {
		return vs[0].([]*TrafficMirrorFilterIngressRule)[vs[1].(int)]
	}).(TrafficMirrorFilterIngressRuleOutput)
}

type TrafficMirrorFilterIngressRuleMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterIngressRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilterIngressRule)(nil)).Elem()
}

func (o TrafficMirrorFilterIngressRuleMapOutput) ToTrafficMirrorFilterIngressRuleMapOutput() TrafficMirrorFilterIngressRuleMapOutput {
	return o
}

func (o TrafficMirrorFilterIngressRuleMapOutput) ToTrafficMirrorFilterIngressRuleMapOutputWithContext(ctx context.Context) TrafficMirrorFilterIngressRuleMapOutput {
	return o
}

func (o TrafficMirrorFilterIngressRuleMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorFilterIngressRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorFilterIngressRule {
		return vs[0].(map[string]*TrafficMirrorFilterIngressRule)[vs[1].(string)]
	}).(TrafficMirrorFilterIngressRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterIngressRuleInput)(nil)).Elem(), &TrafficMirrorFilterIngressRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterIngressRuleArrayInput)(nil)).Elem(), TrafficMirrorFilterIngressRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterIngressRuleMapInput)(nil)).Elem(), TrafficMirrorFilterIngressRuleMap{})
	pulumi.RegisterOutputType(TrafficMirrorFilterIngressRuleOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterIngressRuleArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterIngressRuleMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Traffic Mirror Filter resource. Traffic mirror filter criteria.
//
// For information about VPC Traffic Mirror Filter and how to use it, see [What is Traffic Mirror Filter](https://www.alibabacloud.com/help/doc-detail/207513.htm).
//
// > **NOTE:** Available in v1.140.0+.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			default3iXhoa, err := resourcemanager.NewResourceGroup(ctx, "default3iXhoa", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("testname03"),
//				ResourceGroupName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewResourceGroup(ctx, "defaultdNz2qk", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("testname04"),
//				ResourceGroupName: pulumi.Sprintf("%v1", name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewTrafficMirrorFilter(ctx, "default", &vpc.TrafficMirrorFilterArgs{
//				TrafficMirrorFilterDescription: pulumi.String("test"),
//				TrafficMirrorFilterName:        pulumi.String(name),
//				ResourceGroupId:                default3iXhoa.ID(),
//				EgressRules: vpc.TrafficMirrorFilterEgressRuleTypeArray{
//					&vpc.TrafficMirrorFilterEgressRuleTypeArgs{
//						Priority:             pulumi.Int(1),
//						Protocol:             pulumi.String("TCP"),
//						Action:               pulumi.String("accept"),
//						DestinationCidrBlock: pulumi.String("32.0.0.0/4"),
//						DestinationPortRange: pulumi.String("80/80"),
//						SourceCidrBlock:      pulumi.String("16.0.0.0/4"),
//						SourcePortRange:      pulumi.String("80/80"),
//					},
//				},
//				IngressRules: vpc.TrafficMirrorFilterIngressRuleTypeArray{
//					&vpc.TrafficMirrorFilterIngressRuleTypeArgs{
//						Priority:             pulumi.Int(1),
//						Protocol:             pulumi.String("TCP"),
//						Action:               pulumi.String("accept"),
//						DestinationCidrBlock: pulumi.String("10.64.0.0/10"),
//						DestinationPortRange: pulumi.String("80/80"),
//						SourceCidrBlock:      pulumi.String("10.0.0.0/8"),
//						SourcePortRange:      pulumi.String("80/80"),
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
// VPC Traffic Mirror Filter can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter example <id>
// ```
type TrafficMirrorFilter struct {
	pulumi.CustomResourceState

	// Whether to PreCheck only this request. Value:
	// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Information about the outbound rule. See the following `Block EgressRules`.
	EgressRules TrafficMirrorFilterEgressRuleTypeArrayOutput `pulumi:"egressRules"`
	// Inward direction rule information. See the following `Block IngressRules`.
	IngressRules TrafficMirrorFilterIngressRuleTypeArrayOutput `pulumi:"ingressRules"`
	// The ID of the resource group to which the VPC belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of this resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The description of the TrafficMirrorFilter.
	TrafficMirrorFilterDescription pulumi.StringPtrOutput `pulumi:"trafficMirrorFilterDescription"`
	// The name of the TrafficMirrorFilter.
	TrafficMirrorFilterName pulumi.StringPtrOutput `pulumi:"trafficMirrorFilterName"`
}

// NewTrafficMirrorFilter registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilter(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilter, error) {
	if args == nil {
		args = &TrafficMirrorFilterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficMirrorFilter
	err := ctx.RegisterResource("alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilter gets an existing TrafficMirrorFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilter, error) {
	var resource TrafficMirrorFilter
	err := ctx.ReadResource("alicloud:vpc/trafficMirrorFilter:TrafficMirrorFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilter resources.
type trafficMirrorFilterState struct {
	// Whether to PreCheck only this request. Value:
	// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
	DryRun *bool `pulumi:"dryRun"`
	// Information about the outbound rule. See the following `Block EgressRules`.
	EgressRules []TrafficMirrorFilterEgressRuleType `pulumi:"egressRules"`
	// Inward direction rule information. See the following `Block IngressRules`.
	IngressRules []TrafficMirrorFilterIngressRuleType `pulumi:"ingressRules"`
	// The ID of the resource group to which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The tags of this resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The description of the TrafficMirrorFilter.
	TrafficMirrorFilterDescription *string `pulumi:"trafficMirrorFilterDescription"`
	// The name of the TrafficMirrorFilter.
	TrafficMirrorFilterName *string `pulumi:"trafficMirrorFilterName"`
}

type TrafficMirrorFilterState struct {
	// Whether to PreCheck only this request. Value:
	// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
	DryRun pulumi.BoolPtrInput
	// Information about the outbound rule. See the following `Block EgressRules`.
	EgressRules TrafficMirrorFilterEgressRuleTypeArrayInput
	// Inward direction rule information. See the following `Block IngressRules`.
	IngressRules TrafficMirrorFilterIngressRuleTypeArrayInput
	// The ID of the resource group to which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The tags of this resource.
	Tags pulumi.MapInput
	// The description of the TrafficMirrorFilter.
	TrafficMirrorFilterDescription pulumi.StringPtrInput
	// The name of the TrafficMirrorFilter.
	TrafficMirrorFilterName pulumi.StringPtrInput
}

func (TrafficMirrorFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterState)(nil)).Elem()
}

type trafficMirrorFilterArgs struct {
	// Whether to PreCheck only this request. Value:
	// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
	DryRun *bool `pulumi:"dryRun"`
	// Information about the outbound rule. See the following `Block EgressRules`.
	EgressRules []TrafficMirrorFilterEgressRuleType `pulumi:"egressRules"`
	// Inward direction rule information. See the following `Block IngressRules`.
	IngressRules []TrafficMirrorFilterIngressRuleType `pulumi:"ingressRules"`
	// The ID of the resource group to which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The tags of this resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The description of the TrafficMirrorFilter.
	TrafficMirrorFilterDescription *string `pulumi:"trafficMirrorFilterDescription"`
	// The name of the TrafficMirrorFilter.
	TrafficMirrorFilterName *string `pulumi:"trafficMirrorFilterName"`
}

// The set of arguments for constructing a TrafficMirrorFilter resource.
type TrafficMirrorFilterArgs struct {
	// Whether to PreCheck only this request. Value:
	// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
	// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
	DryRun pulumi.BoolPtrInput
	// Information about the outbound rule. See the following `Block EgressRules`.
	EgressRules TrafficMirrorFilterEgressRuleTypeArrayInput
	// Inward direction rule information. See the following `Block IngressRules`.
	IngressRules TrafficMirrorFilterIngressRuleTypeArrayInput
	// The ID of the resource group to which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The tags of this resource.
	Tags pulumi.MapInput
	// The description of the TrafficMirrorFilter.
	TrafficMirrorFilterDescription pulumi.StringPtrInput
	// The name of the TrafficMirrorFilter.
	TrafficMirrorFilterName pulumi.StringPtrInput
}

func (TrafficMirrorFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterArgs)(nil)).Elem()
}

type TrafficMirrorFilterInput interface {
	pulumi.Input

	ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput
	ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput
}

func (*TrafficMirrorFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilter)(nil)).Elem()
}

func (i *TrafficMirrorFilter) ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput {
	return i.ToTrafficMirrorFilterOutputWithContext(context.Background())
}

func (i *TrafficMirrorFilter) ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterOutput)
}

// TrafficMirrorFilterArrayInput is an input type that accepts TrafficMirrorFilterArray and TrafficMirrorFilterArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterArrayInput` via:
//
//	TrafficMirrorFilterArray{ TrafficMirrorFilterArgs{...} }
type TrafficMirrorFilterArrayInput interface {
	pulumi.Input

	ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput
	ToTrafficMirrorFilterArrayOutputWithContext(context.Context) TrafficMirrorFilterArrayOutput
}

type TrafficMirrorFilterArray []TrafficMirrorFilterInput

func (TrafficMirrorFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilter)(nil)).Elem()
}

func (i TrafficMirrorFilterArray) ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput {
	return i.ToTrafficMirrorFilterArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterArray) ToTrafficMirrorFilterArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterArrayOutput)
}

// TrafficMirrorFilterMapInput is an input type that accepts TrafficMirrorFilterMap and TrafficMirrorFilterMapOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterMapInput` via:
//
//	TrafficMirrorFilterMap{ "key": TrafficMirrorFilterArgs{...} }
type TrafficMirrorFilterMapInput interface {
	pulumi.Input

	ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput
	ToTrafficMirrorFilterMapOutputWithContext(context.Context) TrafficMirrorFilterMapOutput
}

type TrafficMirrorFilterMap map[string]TrafficMirrorFilterInput

func (TrafficMirrorFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilter)(nil)).Elem()
}

func (i TrafficMirrorFilterMap) ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput {
	return i.ToTrafficMirrorFilterMapOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterMap) ToTrafficMirrorFilterMapOutputWithContext(ctx context.Context) TrafficMirrorFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterMapOutput)
}

type TrafficMirrorFilterOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterOutput) ToTrafficMirrorFilterOutput() TrafficMirrorFilterOutput {
	return o
}

func (o TrafficMirrorFilterOutput) ToTrafficMirrorFilterOutputWithContext(ctx context.Context) TrafficMirrorFilterOutput {
	return o
}

// Whether to PreCheck only this request. Value:
// - **true**: The check request is sent without creating traffic Image filter conditions. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
// - **false** (default): Sends a normal request, returns a 2xx HTTP status code after passing the check, and directly creates a filter condition.
func (o TrafficMirrorFilterOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Information about the outbound rule. See the following `Block EgressRules`.
func (o TrafficMirrorFilterOutput) EgressRules() TrafficMirrorFilterEgressRuleTypeArrayOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) TrafficMirrorFilterEgressRuleTypeArrayOutput { return v.EgressRules }).(TrafficMirrorFilterEgressRuleTypeArrayOutput)
}

// Inward direction rule information. See the following `Block IngressRules`.
func (o TrafficMirrorFilterOutput) IngressRules() TrafficMirrorFilterIngressRuleTypeArrayOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) TrafficMirrorFilterIngressRuleTypeArrayOutput { return v.IngressRules }).(TrafficMirrorFilterIngressRuleTypeArrayOutput)
}

// The ID of the resource group to which the VPC belongs.
func (o TrafficMirrorFilterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o TrafficMirrorFilterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of this resource.
func (o TrafficMirrorFilterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The description of the TrafficMirrorFilter.
func (o TrafficMirrorFilterOutput) TrafficMirrorFilterDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringPtrOutput { return v.TrafficMirrorFilterDescription }).(pulumi.StringPtrOutput)
}

// The name of the TrafficMirrorFilter.
func (o TrafficMirrorFilterOutput) TrafficMirrorFilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilter) pulumi.StringPtrOutput { return v.TrafficMirrorFilterName }).(pulumi.StringPtrOutput)
}

type TrafficMirrorFilterArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterArrayOutput) ToTrafficMirrorFilterArrayOutput() TrafficMirrorFilterArrayOutput {
	return o
}

func (o TrafficMirrorFilterArrayOutput) ToTrafficMirrorFilterArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterArrayOutput {
	return o
}

func (o TrafficMirrorFilterArrayOutput) Index(i pulumi.IntInput) TrafficMirrorFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorFilter {
		return vs[0].([]*TrafficMirrorFilter)[vs[1].(int)]
	}).(TrafficMirrorFilterOutput)
}

type TrafficMirrorFilterMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilter)(nil)).Elem()
}

func (o TrafficMirrorFilterMapOutput) ToTrafficMirrorFilterMapOutput() TrafficMirrorFilterMapOutput {
	return o
}

func (o TrafficMirrorFilterMapOutput) ToTrafficMirrorFilterMapOutputWithContext(ctx context.Context) TrafficMirrorFilterMapOutput {
	return o
}

func (o TrafficMirrorFilterMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorFilter {
		return vs[0].(map[string]*TrafficMirrorFilter)[vs[1].(string)]
	}).(TrafficMirrorFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterInput)(nil)).Elem(), &TrafficMirrorFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterArrayInput)(nil)).Elem(), TrafficMirrorFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterMapInput)(nil)).Elem(), TrafficMirrorFilterMap{})
	pulumi.RegisterOutputType(TrafficMirrorFilterOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterMapOutput{})
}

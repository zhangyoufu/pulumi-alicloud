// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule resource.
//
// For information about Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule and how to use it, see [What is Prometheus Alert Rule](https://www.alibabacloud.com/help/en/doc-detail/212056.htm).
//
// > **NOTE:** Available since v1.136.0.
//
// ## Import
//
// Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:arms/prometheusAlertRule:PrometheusAlertRule example <cluster_id>:<prometheus_alert_rule_id>
// ```
type PrometheusAlertRule struct {
	pulumi.CustomResourceState

	// The annotations of the alert rule. See `annotations` below.
	Annotations PrometheusAlertRuleAnnotationArrayOutput `pulumi:"annotations"`
	// The ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
	DispatchRuleId pulumi.StringPtrOutput `pulumi:"dispatchRuleId"`
	// The duration of the alert.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// The alert rule expression that follows the PromQL syntax.
	Expression pulumi.StringOutput `pulumi:"expression"`
	// The labels of the resource. See `labels` below.
	Labels PrometheusAlertRuleLabelArrayOutput `pulumi:"labels"`
	// The message of the alert notification.
	Message pulumi.StringOutput `pulumi:"message"`
	// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
	NotifyType pulumi.StringPtrOutput `pulumi:"notifyType"`
	// The first ID of the resource.
	PrometheusAlertRuleId pulumi.IntOutput `pulumi:"prometheusAlertRuleId"`
	// The name of the resource.
	PrometheusAlertRuleName pulumi.StringOutput `pulumi:"prometheusAlertRuleName"`
	// The status of the resource. Valid values: `0`, `1`.
	Status pulumi.IntOutput `pulumi:"status"`
	// The type of the alert rule.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrometheusAlertRule registers a new resource with the given unique name, arguments, and options.
func NewPrometheusAlertRule(ctx *pulumi.Context,
	name string, args *PrometheusAlertRuleArgs, opts ...pulumi.ResourceOption) (*PrometheusAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	if args.Message == nil {
		return nil, errors.New("invalid value for required argument 'Message'")
	}
	if args.PrometheusAlertRuleName == nil {
		return nil, errors.New("invalid value for required argument 'PrometheusAlertRuleName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrometheusAlertRule
	err := ctx.RegisterResource("alicloud:arms/prometheusAlertRule:PrometheusAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusAlertRule gets an existing PrometheusAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusAlertRuleState, opts ...pulumi.ResourceOption) (*PrometheusAlertRule, error) {
	var resource PrometheusAlertRule
	err := ctx.ReadResource("alicloud:arms/prometheusAlertRule:PrometheusAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusAlertRule resources.
type prometheusAlertRuleState struct {
	// The annotations of the alert rule. See `annotations` below.
	Annotations []PrometheusAlertRuleAnnotation `pulumi:"annotations"`
	// The ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
	DispatchRuleId *string `pulumi:"dispatchRuleId"`
	// The duration of the alert.
	Duration *string `pulumi:"duration"`
	// The alert rule expression that follows the PromQL syntax.
	Expression *string `pulumi:"expression"`
	// The labels of the resource. See `labels` below.
	Labels []PrometheusAlertRuleLabel `pulumi:"labels"`
	// The message of the alert notification.
	Message *string `pulumi:"message"`
	// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
	NotifyType *string `pulumi:"notifyType"`
	// The first ID of the resource.
	PrometheusAlertRuleId *int `pulumi:"prometheusAlertRuleId"`
	// The name of the resource.
	PrometheusAlertRuleName *string `pulumi:"prometheusAlertRuleName"`
	// The status of the resource. Valid values: `0`, `1`.
	Status *int `pulumi:"status"`
	// The type of the alert rule.
	Type *string `pulumi:"type"`
}

type PrometheusAlertRuleState struct {
	// The annotations of the alert rule. See `annotations` below.
	Annotations PrometheusAlertRuleAnnotationArrayInput
	// The ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
	DispatchRuleId pulumi.StringPtrInput
	// The duration of the alert.
	Duration pulumi.StringPtrInput
	// The alert rule expression that follows the PromQL syntax.
	Expression pulumi.StringPtrInput
	// The labels of the resource. See `labels` below.
	Labels PrometheusAlertRuleLabelArrayInput
	// The message of the alert notification.
	Message pulumi.StringPtrInput
	// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
	NotifyType pulumi.StringPtrInput
	// The first ID of the resource.
	PrometheusAlertRuleId pulumi.IntPtrInput
	// The name of the resource.
	PrometheusAlertRuleName pulumi.StringPtrInput
	// The status of the resource. Valid values: `0`, `1`.
	Status pulumi.IntPtrInput
	// The type of the alert rule.
	Type pulumi.StringPtrInput
}

func (PrometheusAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAlertRuleState)(nil)).Elem()
}

type prometheusAlertRuleArgs struct {
	// The annotations of the alert rule. See `annotations` below.
	Annotations []PrometheusAlertRuleAnnotation `pulumi:"annotations"`
	// The ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
	DispatchRuleId *string `pulumi:"dispatchRuleId"`
	// The duration of the alert.
	Duration string `pulumi:"duration"`
	// The alert rule expression that follows the PromQL syntax.
	Expression string `pulumi:"expression"`
	// The labels of the resource. See `labels` below.
	Labels []PrometheusAlertRuleLabel `pulumi:"labels"`
	// The message of the alert notification.
	Message string `pulumi:"message"`
	// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
	NotifyType *string `pulumi:"notifyType"`
	// The name of the resource.
	PrometheusAlertRuleName string `pulumi:"prometheusAlertRuleName"`
	// The type of the alert rule.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a PrometheusAlertRule resource.
type PrometheusAlertRuleArgs struct {
	// The annotations of the alert rule. See `annotations` below.
	Annotations PrometheusAlertRuleAnnotationArrayInput
	// The ID of the cluster.
	ClusterId pulumi.StringInput
	// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
	DispatchRuleId pulumi.StringPtrInput
	// The duration of the alert.
	Duration pulumi.StringInput
	// The alert rule expression that follows the PromQL syntax.
	Expression pulumi.StringInput
	// The labels of the resource. See `labels` below.
	Labels PrometheusAlertRuleLabelArrayInput
	// The message of the alert notification.
	Message pulumi.StringInput
	// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
	NotifyType pulumi.StringPtrInput
	// The name of the resource.
	PrometheusAlertRuleName pulumi.StringInput
	// The type of the alert rule.
	Type pulumi.StringPtrInput
}

func (PrometheusAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAlertRuleArgs)(nil)).Elem()
}

type PrometheusAlertRuleInput interface {
	pulumi.Input

	ToPrometheusAlertRuleOutput() PrometheusAlertRuleOutput
	ToPrometheusAlertRuleOutputWithContext(ctx context.Context) PrometheusAlertRuleOutput
}

func (*PrometheusAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAlertRule)(nil)).Elem()
}

func (i *PrometheusAlertRule) ToPrometheusAlertRuleOutput() PrometheusAlertRuleOutput {
	return i.ToPrometheusAlertRuleOutputWithContext(context.Background())
}

func (i *PrometheusAlertRule) ToPrometheusAlertRuleOutputWithContext(ctx context.Context) PrometheusAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAlertRuleOutput)
}

// PrometheusAlertRuleArrayInput is an input type that accepts PrometheusAlertRuleArray and PrometheusAlertRuleArrayOutput values.
// You can construct a concrete instance of `PrometheusAlertRuleArrayInput` via:
//
//	PrometheusAlertRuleArray{ PrometheusAlertRuleArgs{...} }
type PrometheusAlertRuleArrayInput interface {
	pulumi.Input

	ToPrometheusAlertRuleArrayOutput() PrometheusAlertRuleArrayOutput
	ToPrometheusAlertRuleArrayOutputWithContext(context.Context) PrometheusAlertRuleArrayOutput
}

type PrometheusAlertRuleArray []PrometheusAlertRuleInput

func (PrometheusAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAlertRule)(nil)).Elem()
}

func (i PrometheusAlertRuleArray) ToPrometheusAlertRuleArrayOutput() PrometheusAlertRuleArrayOutput {
	return i.ToPrometheusAlertRuleArrayOutputWithContext(context.Background())
}

func (i PrometheusAlertRuleArray) ToPrometheusAlertRuleArrayOutputWithContext(ctx context.Context) PrometheusAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAlertRuleArrayOutput)
}

// PrometheusAlertRuleMapInput is an input type that accepts PrometheusAlertRuleMap and PrometheusAlertRuleMapOutput values.
// You can construct a concrete instance of `PrometheusAlertRuleMapInput` via:
//
//	PrometheusAlertRuleMap{ "key": PrometheusAlertRuleArgs{...} }
type PrometheusAlertRuleMapInput interface {
	pulumi.Input

	ToPrometheusAlertRuleMapOutput() PrometheusAlertRuleMapOutput
	ToPrometheusAlertRuleMapOutputWithContext(context.Context) PrometheusAlertRuleMapOutput
}

type PrometheusAlertRuleMap map[string]PrometheusAlertRuleInput

func (PrometheusAlertRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAlertRule)(nil)).Elem()
}

func (i PrometheusAlertRuleMap) ToPrometheusAlertRuleMapOutput() PrometheusAlertRuleMapOutput {
	return i.ToPrometheusAlertRuleMapOutputWithContext(context.Background())
}

func (i PrometheusAlertRuleMap) ToPrometheusAlertRuleMapOutputWithContext(ctx context.Context) PrometheusAlertRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAlertRuleMapOutput)
}

type PrometheusAlertRuleOutput struct{ *pulumi.OutputState }

func (PrometheusAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAlertRule)(nil)).Elem()
}

func (o PrometheusAlertRuleOutput) ToPrometheusAlertRuleOutput() PrometheusAlertRuleOutput {
	return o
}

func (o PrometheusAlertRuleOutput) ToPrometheusAlertRuleOutputWithContext(ctx context.Context) PrometheusAlertRuleOutput {
	return o
}

// The annotations of the alert rule. See `annotations` below.
func (o PrometheusAlertRuleOutput) Annotations() PrometheusAlertRuleAnnotationArrayOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) PrometheusAlertRuleAnnotationArrayOutput { return v.Annotations }).(PrometheusAlertRuleAnnotationArrayOutput)
}

// The ID of the cluster.
func (o PrometheusAlertRuleOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The ID of the notification policy. This parameter is required when the `notifyType` parameter is set to `DISPATCH_RULE`.
func (o PrometheusAlertRuleOutput) DispatchRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringPtrOutput { return v.DispatchRuleId }).(pulumi.StringPtrOutput)
}

// The duration of the alert.
func (o PrometheusAlertRuleOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

// The alert rule expression that follows the PromQL syntax.
func (o PrometheusAlertRuleOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

// The labels of the resource. See `labels` below.
func (o PrometheusAlertRuleOutput) Labels() PrometheusAlertRuleLabelArrayOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) PrometheusAlertRuleLabelArrayOutput { return v.Labels }).(PrometheusAlertRuleLabelArrayOutput)
}

// The message of the alert notification.
func (o PrometheusAlertRuleOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
func (o PrometheusAlertRuleOutput) NotifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringPtrOutput { return v.NotifyType }).(pulumi.StringPtrOutput)
}

// The first ID of the resource.
func (o PrometheusAlertRuleOutput) PrometheusAlertRuleId() pulumi.IntOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.IntOutput { return v.PrometheusAlertRuleId }).(pulumi.IntOutput)
}

// The name of the resource.
func (o PrometheusAlertRuleOutput) PrometheusAlertRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.PrometheusAlertRuleName }).(pulumi.StringOutput)
}

// The status of the resource. Valid values: `0`, `1`.
func (o PrometheusAlertRuleOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// The type of the alert rule.
func (o PrometheusAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type PrometheusAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (PrometheusAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAlertRule)(nil)).Elem()
}

func (o PrometheusAlertRuleArrayOutput) ToPrometheusAlertRuleArrayOutput() PrometheusAlertRuleArrayOutput {
	return o
}

func (o PrometheusAlertRuleArrayOutput) ToPrometheusAlertRuleArrayOutputWithContext(ctx context.Context) PrometheusAlertRuleArrayOutput {
	return o
}

func (o PrometheusAlertRuleArrayOutput) Index(i pulumi.IntInput) PrometheusAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrometheusAlertRule {
		return vs[0].([]*PrometheusAlertRule)[vs[1].(int)]
	}).(PrometheusAlertRuleOutput)
}

type PrometheusAlertRuleMapOutput struct{ *pulumi.OutputState }

func (PrometheusAlertRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAlertRule)(nil)).Elem()
}

func (o PrometheusAlertRuleMapOutput) ToPrometheusAlertRuleMapOutput() PrometheusAlertRuleMapOutput {
	return o
}

func (o PrometheusAlertRuleMapOutput) ToPrometheusAlertRuleMapOutputWithContext(ctx context.Context) PrometheusAlertRuleMapOutput {
	return o
}

func (o PrometheusAlertRuleMapOutput) MapIndex(k pulumi.StringInput) PrometheusAlertRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrometheusAlertRule {
		return vs[0].(map[string]*PrometheusAlertRule)[vs[1].(string)]
	}).(PrometheusAlertRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAlertRuleInput)(nil)).Elem(), &PrometheusAlertRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAlertRuleArrayInput)(nil)).Elem(), PrometheusAlertRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAlertRuleMapInput)(nil)).Elem(), PrometheusAlertRuleMap{})
	pulumi.RegisterOutputType(PrometheusAlertRuleOutput{})
	pulumi.RegisterOutputType(PrometheusAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(PrometheusAlertRuleMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Quotas Quota Alarm resource.
//
// For information about Quotas Quota Alarm and how to use it, see [What is Quota Alarm](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaalarm).
//
// > **NOTE:** Available since v1.116.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
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
//			_, err := quotas.NewQuotaAlarm(ctx, "default", &quotas.QuotaAlarmArgs{
//				QuotaActionCode: pulumi.String("q_desktop-count"),
//				QuotaDimensions: quotas.QuotaAlarmQuotaDimensionArray{
//					&quotas.QuotaAlarmQuotaDimensionArgs{
//						Key:   pulumi.String("regionId"),
//						Value: pulumi.String("cn-hangzhou"),
//					},
//				},
//				ThresholdPercent: pulumi.Float64(80),
//				ProductCode:      pulumi.String("gws"),
//				QuotaAlarmName:   pulumi.String(name),
//				ThresholdType:    pulumi.String("used"),
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
// Quotas Quota Alarm can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:quotas/quotaAlarm:QuotaAlarm example <id>
// ```
type QuotaAlarm struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The Product Code.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode pulumi.StringOutput `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringOutput `pulumi:"quotaAlarmName"`
	// The Quota Dimensions. See `quotaDimensions` below.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayOutput `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrOutput `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrOutput `pulumi:"thresholdPercent"`
	// Quota alarm type. Value:
	// - used: Quota used alarm.
	// - usable: alarm for the remaining available quota.
	ThresholdType pulumi.StringOutput `pulumi:"thresholdType"`
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrOutput `pulumi:"webHook"`
}

// NewQuotaAlarm registers a new resource with the given unique name, arguments, and options.
func NewQuotaAlarm(ctx *pulumi.Context,
	name string, args *QuotaAlarmArgs, opts ...pulumi.ResourceOption) (*QuotaAlarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	if args.QuotaActionCode == nil {
		return nil, errors.New("invalid value for required argument 'QuotaActionCode'")
	}
	if args.QuotaAlarmName == nil {
		return nil, errors.New("invalid value for required argument 'QuotaAlarmName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QuotaAlarm
	err := ctx.RegisterResource("alicloud:quotas/quotaAlarm:QuotaAlarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaAlarm gets an existing QuotaAlarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaAlarmState, opts ...pulumi.ResourceOption) (*QuotaAlarm, error) {
	var resource QuotaAlarm
	err := ctx.ReadResource("alicloud:quotas/quotaAlarm:QuotaAlarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaAlarm resources.
type quotaAlarmState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The Product Code.
	ProductCode *string `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName *string `pulumi:"quotaAlarmName"`
	// The Quota Dimensions. See `quotaDimensions` below.
	QuotaDimensions []QuotaAlarmQuotaDimension `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold *float64 `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent *float64 `pulumi:"thresholdPercent"`
	// Quota alarm type. Value:
	// - used: Quota used alarm.
	// - usable: alarm for the remaining available quota.
	ThresholdType *string `pulumi:"thresholdType"`
	// The WebHook of Quota Alarm.
	WebHook *string `pulumi:"webHook"`
}

type QuotaAlarmState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The Product Code.
	ProductCode pulumi.StringPtrInput
	// The Quota Action Code.
	QuotaActionCode pulumi.StringPtrInput
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringPtrInput
	// The Quota Dimensions. See `quotaDimensions` below.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayInput
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrInput
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrInput
	// Quota alarm type. Value:
	// - used: Quota used alarm.
	// - usable: alarm for the remaining available quota.
	ThresholdType pulumi.StringPtrInput
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrInput
}

func (QuotaAlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaAlarmState)(nil)).Elem()
}

type quotaAlarmArgs struct {
	// The Product Code.
	ProductCode string `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode string `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName string `pulumi:"quotaAlarmName"`
	// The Quota Dimensions. See `quotaDimensions` below.
	QuotaDimensions []QuotaAlarmQuotaDimension `pulumi:"quotaDimensions"`
	// The threshold of Quota Alarm.
	Threshold *float64 `pulumi:"threshold"`
	// The threshold percent of Quota Alarm.
	ThresholdPercent *float64 `pulumi:"thresholdPercent"`
	// Quota alarm type. Value:
	// - used: Quota used alarm.
	// - usable: alarm for the remaining available quota.
	ThresholdType *string `pulumi:"thresholdType"`
	// The WebHook of Quota Alarm.
	WebHook *string `pulumi:"webHook"`
}

// The set of arguments for constructing a QuotaAlarm resource.
type QuotaAlarmArgs struct {
	// The Product Code.
	ProductCode pulumi.StringInput
	// The Quota Action Code.
	QuotaActionCode pulumi.StringInput
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringInput
	// The Quota Dimensions. See `quotaDimensions` below.
	QuotaDimensions QuotaAlarmQuotaDimensionArrayInput
	// The threshold of Quota Alarm.
	Threshold pulumi.Float64PtrInput
	// The threshold percent of Quota Alarm.
	ThresholdPercent pulumi.Float64PtrInput
	// Quota alarm type. Value:
	// - used: Quota used alarm.
	// - usable: alarm for the remaining available quota.
	ThresholdType pulumi.StringPtrInput
	// The WebHook of Quota Alarm.
	WebHook pulumi.StringPtrInput
}

func (QuotaAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaAlarmArgs)(nil)).Elem()
}

type QuotaAlarmInput interface {
	pulumi.Input

	ToQuotaAlarmOutput() QuotaAlarmOutput
	ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput
}

func (*QuotaAlarm) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaAlarm)(nil)).Elem()
}

func (i *QuotaAlarm) ToQuotaAlarmOutput() QuotaAlarmOutput {
	return i.ToQuotaAlarmOutputWithContext(context.Background())
}

func (i *QuotaAlarm) ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmOutput)
}

// QuotaAlarmArrayInput is an input type that accepts QuotaAlarmArray and QuotaAlarmArrayOutput values.
// You can construct a concrete instance of `QuotaAlarmArrayInput` via:
//
//	QuotaAlarmArray{ QuotaAlarmArgs{...} }
type QuotaAlarmArrayInput interface {
	pulumi.Input

	ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput
	ToQuotaAlarmArrayOutputWithContext(context.Context) QuotaAlarmArrayOutput
}

type QuotaAlarmArray []QuotaAlarmInput

func (QuotaAlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaAlarm)(nil)).Elem()
}

func (i QuotaAlarmArray) ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput {
	return i.ToQuotaAlarmArrayOutputWithContext(context.Background())
}

func (i QuotaAlarmArray) ToQuotaAlarmArrayOutputWithContext(ctx context.Context) QuotaAlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmArrayOutput)
}

// QuotaAlarmMapInput is an input type that accepts QuotaAlarmMap and QuotaAlarmMapOutput values.
// You can construct a concrete instance of `QuotaAlarmMapInput` via:
//
//	QuotaAlarmMap{ "key": QuotaAlarmArgs{...} }
type QuotaAlarmMapInput interface {
	pulumi.Input

	ToQuotaAlarmMapOutput() QuotaAlarmMapOutput
	ToQuotaAlarmMapOutputWithContext(context.Context) QuotaAlarmMapOutput
}

type QuotaAlarmMap map[string]QuotaAlarmInput

func (QuotaAlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaAlarm)(nil)).Elem()
}

func (i QuotaAlarmMap) ToQuotaAlarmMapOutput() QuotaAlarmMapOutput {
	return i.ToQuotaAlarmMapOutputWithContext(context.Background())
}

func (i QuotaAlarmMap) ToQuotaAlarmMapOutputWithContext(ctx context.Context) QuotaAlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaAlarmMapOutput)
}

type QuotaAlarmOutput struct{ *pulumi.OutputState }

func (QuotaAlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaAlarm)(nil)).Elem()
}

func (o QuotaAlarmOutput) ToQuotaAlarmOutput() QuotaAlarmOutput {
	return o
}

func (o QuotaAlarmOutput) ToQuotaAlarmOutputWithContext(ctx context.Context) QuotaAlarmOutput {
	return o
}

// The creation time of the resource.
func (o QuotaAlarmOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The Product Code.
func (o QuotaAlarmOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringOutput { return v.ProductCode }).(pulumi.StringOutput)
}

// The Quota Action Code.
func (o QuotaAlarmOutput) QuotaActionCode() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringOutput { return v.QuotaActionCode }).(pulumi.StringOutput)
}

// The name of Quota Alarm.
func (o QuotaAlarmOutput) QuotaAlarmName() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringOutput { return v.QuotaAlarmName }).(pulumi.StringOutput)
}

// The Quota Dimensions. See `quotaDimensions` below.
func (o QuotaAlarmOutput) QuotaDimensions() QuotaAlarmQuotaDimensionArrayOutput {
	return o.ApplyT(func(v *QuotaAlarm) QuotaAlarmQuotaDimensionArrayOutput { return v.QuotaDimensions }).(QuotaAlarmQuotaDimensionArrayOutput)
}

// The threshold of Quota Alarm.
func (o QuotaAlarmOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.Float64PtrOutput { return v.Threshold }).(pulumi.Float64PtrOutput)
}

// The threshold percent of Quota Alarm.
func (o QuotaAlarmOutput) ThresholdPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.Float64PtrOutput { return v.ThresholdPercent }).(pulumi.Float64PtrOutput)
}

// Quota alarm type. Value:
// - used: Quota used alarm.
// - usable: alarm for the remaining available quota.
func (o QuotaAlarmOutput) ThresholdType() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringOutput { return v.ThresholdType }).(pulumi.StringOutput)
}

// The WebHook of Quota Alarm.
func (o QuotaAlarmOutput) WebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuotaAlarm) pulumi.StringPtrOutput { return v.WebHook }).(pulumi.StringPtrOutput)
}

type QuotaAlarmArrayOutput struct{ *pulumi.OutputState }

func (QuotaAlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaAlarm)(nil)).Elem()
}

func (o QuotaAlarmArrayOutput) ToQuotaAlarmArrayOutput() QuotaAlarmArrayOutput {
	return o
}

func (o QuotaAlarmArrayOutput) ToQuotaAlarmArrayOutputWithContext(ctx context.Context) QuotaAlarmArrayOutput {
	return o
}

func (o QuotaAlarmArrayOutput) Index(i pulumi.IntInput) QuotaAlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuotaAlarm {
		return vs[0].([]*QuotaAlarm)[vs[1].(int)]
	}).(QuotaAlarmOutput)
}

type QuotaAlarmMapOutput struct{ *pulumi.OutputState }

func (QuotaAlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaAlarm)(nil)).Elem()
}

func (o QuotaAlarmMapOutput) ToQuotaAlarmMapOutput() QuotaAlarmMapOutput {
	return o
}

func (o QuotaAlarmMapOutput) ToQuotaAlarmMapOutputWithContext(ctx context.Context) QuotaAlarmMapOutput {
	return o
}

func (o QuotaAlarmMapOutput) MapIndex(k pulumi.StringInput) QuotaAlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuotaAlarm {
		return vs[0].(map[string]*QuotaAlarm)[vs[1].(string)]
	}).(QuotaAlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaAlarmInput)(nil)).Elem(), &QuotaAlarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaAlarmArrayInput)(nil)).Elem(), QuotaAlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaAlarmMapInput)(nil)).Elem(), QuotaAlarmMap{})
	pulumi.RegisterOutputType(QuotaAlarmOutput{})
	pulumi.RegisterOutputType(QuotaAlarmArrayOutput{})
	pulumi.RegisterOutputType(QuotaAlarmMapOutput{})
}

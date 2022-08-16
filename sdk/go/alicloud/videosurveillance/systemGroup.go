// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videosurveillance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Video Surveillance System Group resource.
//
// For information about Video Surveillance System Group and how to use it, see [What is Group](https://help.aliyun.com/product/108765.html).
//
// > **NOTE:** Available in v1.135.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/videosurveillance"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := videosurveillance.NewSystemGroup(ctx, "default", &videosurveillance.SystemGroupArgs{
//				GroupName:   pulumi.String("your_group_name"),
//				InProtocol:  pulumi.String("rtmp"),
//				OutProtocol: pulumi.String("flv"),
//				PlayDomain:  pulumi.String("your_plan_domain"),
//				PushDomain:  pulumi.String("your_push_domain"),
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
// Video Surveillance System Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:videosurveillance/systemGroup:SystemGroup example <id>
//
// ```
type SystemGroup struct {
	pulumi.CustomResourceState

	// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
	Callback pulumi.StringPtrOutput `pulumi:"callback"`
	// The capture image.
	CaptureImage pulumi.IntOutput `pulumi:"captureImage"`
	// The capture interval.
	CaptureInterval pulumi.IntOutput `pulumi:"captureInterval"`
	// The capture oss bucket.
	CaptureOssBucket pulumi.StringOutput `pulumi:"captureOssBucket"`
	// The capture oss path.
	CaptureOssPath pulumi.StringOutput `pulumi:"captureOssPath"`
	// The capture video.
	CaptureVideo pulumi.IntOutput `pulumi:"captureVideo"`
	// The description of Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to open Group.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The Group Name.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
	InProtocol pulumi.StringOutput `pulumi:"inProtocol"`
	// Whether to enable on-demand streaming. Default value:`false`.
	LazyPull pulumi.BoolOutput `pulumi:"lazyPull"`
	// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
	OutProtocol pulumi.StringOutput `pulumi:"outProtocol"`
	// The domain name of plan streaming used by the group.
	PlayDomain pulumi.StringOutput `pulumi:"playDomain"`
	// The domain name of push streaming used by the group.
	PushDomain pulumi.StringOutput `pulumi:"pushDomain"`
	// Whether to open Group. Valid values: `on`,`off`.
	Status pulumi.BoolOutput `pulumi:"status"`
}

// NewSystemGroup registers a new resource with the given unique name, arguments, and options.
func NewSystemGroup(ctx *pulumi.Context,
	name string, args *SystemGroupArgs, opts ...pulumi.ResourceOption) (*SystemGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.InProtocol == nil {
		return nil, errors.New("invalid value for required argument 'InProtocol'")
	}
	if args.OutProtocol == nil {
		return nil, errors.New("invalid value for required argument 'OutProtocol'")
	}
	if args.PlayDomain == nil {
		return nil, errors.New("invalid value for required argument 'PlayDomain'")
	}
	if args.PushDomain == nil {
		return nil, errors.New("invalid value for required argument 'PushDomain'")
	}
	var resource SystemGroup
	err := ctx.RegisterResource("alicloud:videosurveillance/systemGroup:SystemGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemGroup gets an existing SystemGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemGroupState, opts ...pulumi.ResourceOption) (*SystemGroup, error) {
	var resource SystemGroup
	err := ctx.ReadResource("alicloud:videosurveillance/systemGroup:SystemGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemGroup resources.
type systemGroupState struct {
	// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
	Callback *string `pulumi:"callback"`
	// The capture image.
	CaptureImage *int `pulumi:"captureImage"`
	// The capture interval.
	CaptureInterval *int `pulumi:"captureInterval"`
	// The capture oss bucket.
	CaptureOssBucket *string `pulumi:"captureOssBucket"`
	// The capture oss path.
	CaptureOssPath *string `pulumi:"captureOssPath"`
	// The capture video.
	CaptureVideo *int `pulumi:"captureVideo"`
	// The description of Group.
	Description *string `pulumi:"description"`
	// Whether to open Group.
	Enabled *bool `pulumi:"enabled"`
	// The Group Name.
	GroupName *string `pulumi:"groupName"`
	// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
	InProtocol *string `pulumi:"inProtocol"`
	// Whether to enable on-demand streaming. Default value:`false`.
	LazyPull *bool `pulumi:"lazyPull"`
	// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
	OutProtocol *string `pulumi:"outProtocol"`
	// The domain name of plan streaming used by the group.
	PlayDomain *string `pulumi:"playDomain"`
	// The domain name of push streaming used by the group.
	PushDomain *string `pulumi:"pushDomain"`
	// Whether to open Group. Valid values: `on`,`off`.
	Status *bool `pulumi:"status"`
}

type SystemGroupState struct {
	// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
	Callback pulumi.StringPtrInput
	// The capture image.
	CaptureImage pulumi.IntPtrInput
	// The capture interval.
	CaptureInterval pulumi.IntPtrInput
	// The capture oss bucket.
	CaptureOssBucket pulumi.StringPtrInput
	// The capture oss path.
	CaptureOssPath pulumi.StringPtrInput
	// The capture video.
	CaptureVideo pulumi.IntPtrInput
	// The description of Group.
	Description pulumi.StringPtrInput
	// Whether to open Group.
	Enabled pulumi.BoolPtrInput
	// The Group Name.
	GroupName pulumi.StringPtrInput
	// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
	InProtocol pulumi.StringPtrInput
	// Whether to enable on-demand streaming. Default value:`false`.
	LazyPull pulumi.BoolPtrInput
	// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
	OutProtocol pulumi.StringPtrInput
	// The domain name of plan streaming used by the group.
	PlayDomain pulumi.StringPtrInput
	// The domain name of push streaming used by the group.
	PushDomain pulumi.StringPtrInput
	// Whether to open Group. Valid values: `on`,`off`.
	Status pulumi.BoolPtrInput
}

func (SystemGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGroupState)(nil)).Elem()
}

type systemGroupArgs struct {
	// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
	Callback *string `pulumi:"callback"`
	// The description of Group.
	Description *string `pulumi:"description"`
	// Whether to open Group.
	Enabled *bool `pulumi:"enabled"`
	// The Group Name.
	GroupName string `pulumi:"groupName"`
	// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
	InProtocol string `pulumi:"inProtocol"`
	// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
	OutProtocol string `pulumi:"outProtocol"`
	// The domain name of plan streaming used by the group.
	PlayDomain string `pulumi:"playDomain"`
	// The domain name of push streaming used by the group.
	PushDomain string `pulumi:"pushDomain"`
}

// The set of arguments for constructing a SystemGroup resource.
type SystemGroupArgs struct {
	// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
	Callback pulumi.StringPtrInput
	// The description of Group.
	Description pulumi.StringPtrInput
	// Whether to open Group.
	Enabled pulumi.BoolPtrInput
	// The Group Name.
	GroupName pulumi.StringInput
	// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
	InProtocol pulumi.StringInput
	// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
	OutProtocol pulumi.StringInput
	// The domain name of plan streaming used by the group.
	PlayDomain pulumi.StringInput
	// The domain name of push streaming used by the group.
	PushDomain pulumi.StringInput
}

func (SystemGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemGroupArgs)(nil)).Elem()
}

type SystemGroupInput interface {
	pulumi.Input

	ToSystemGroupOutput() SystemGroupOutput
	ToSystemGroupOutputWithContext(ctx context.Context) SystemGroupOutput
}

func (*SystemGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGroup)(nil)).Elem()
}

func (i *SystemGroup) ToSystemGroupOutput() SystemGroupOutput {
	return i.ToSystemGroupOutputWithContext(context.Background())
}

func (i *SystemGroup) ToSystemGroupOutputWithContext(ctx context.Context) SystemGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGroupOutput)
}

// SystemGroupArrayInput is an input type that accepts SystemGroupArray and SystemGroupArrayOutput values.
// You can construct a concrete instance of `SystemGroupArrayInput` via:
//
//	SystemGroupArray{ SystemGroupArgs{...} }
type SystemGroupArrayInput interface {
	pulumi.Input

	ToSystemGroupArrayOutput() SystemGroupArrayOutput
	ToSystemGroupArrayOutputWithContext(context.Context) SystemGroupArrayOutput
}

type SystemGroupArray []SystemGroupInput

func (SystemGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGroup)(nil)).Elem()
}

func (i SystemGroupArray) ToSystemGroupArrayOutput() SystemGroupArrayOutput {
	return i.ToSystemGroupArrayOutputWithContext(context.Background())
}

func (i SystemGroupArray) ToSystemGroupArrayOutputWithContext(ctx context.Context) SystemGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGroupArrayOutput)
}

// SystemGroupMapInput is an input type that accepts SystemGroupMap and SystemGroupMapOutput values.
// You can construct a concrete instance of `SystemGroupMapInput` via:
//
//	SystemGroupMap{ "key": SystemGroupArgs{...} }
type SystemGroupMapInput interface {
	pulumi.Input

	ToSystemGroupMapOutput() SystemGroupMapOutput
	ToSystemGroupMapOutputWithContext(context.Context) SystemGroupMapOutput
}

type SystemGroupMap map[string]SystemGroupInput

func (SystemGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGroup)(nil)).Elem()
}

func (i SystemGroupMap) ToSystemGroupMapOutput() SystemGroupMapOutput {
	return i.ToSystemGroupMapOutputWithContext(context.Background())
}

func (i SystemGroupMap) ToSystemGroupMapOutputWithContext(ctx context.Context) SystemGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemGroupMapOutput)
}

type SystemGroupOutput struct{ *pulumi.OutputState }

func (SystemGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemGroup)(nil)).Elem()
}

func (o SystemGroupOutput) ToSystemGroupOutput() SystemGroupOutput {
	return o
}

func (o SystemGroupOutput) ToSystemGroupOutputWithContext(ctx context.Context) SystemGroupOutput {
	return o
}

// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
func (o SystemGroupOutput) Callback() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringPtrOutput { return v.Callback }).(pulumi.StringPtrOutput)
}

// The capture image.
func (o SystemGroupOutput) CaptureImage() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.IntOutput { return v.CaptureImage }).(pulumi.IntOutput)
}

// The capture interval.
func (o SystemGroupOutput) CaptureInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.IntOutput { return v.CaptureInterval }).(pulumi.IntOutput)
}

// The capture oss bucket.
func (o SystemGroupOutput) CaptureOssBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.CaptureOssBucket }).(pulumi.StringOutput)
}

// The capture oss path.
func (o SystemGroupOutput) CaptureOssPath() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.CaptureOssPath }).(pulumi.StringOutput)
}

// The capture video.
func (o SystemGroupOutput) CaptureVideo() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.IntOutput { return v.CaptureVideo }).(pulumi.IntOutput)
}

// The description of Group.
func (o SystemGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to open Group.
func (o SystemGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The Group Name.
func (o SystemGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
func (o SystemGroupOutput) InProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.InProtocol }).(pulumi.StringOutput)
}

// Whether to enable on-demand streaming. Default value:`false`.
func (o SystemGroupOutput) LazyPull() pulumi.BoolOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.BoolOutput { return v.LazyPull }).(pulumi.BoolOutput)
}

// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
func (o SystemGroupOutput) OutProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.OutProtocol }).(pulumi.StringOutput)
}

// The domain name of plan streaming used by the group.
func (o SystemGroupOutput) PlayDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.PlayDomain }).(pulumi.StringOutput)
}

// The domain name of push streaming used by the group.
func (o SystemGroupOutput) PushDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.StringOutput { return v.PushDomain }).(pulumi.StringOutput)
}

// Whether to open Group. Valid values: `on`,`off`.
func (o SystemGroupOutput) Status() pulumi.BoolOutput {
	return o.ApplyT(func(v *SystemGroup) pulumi.BoolOutput { return v.Status }).(pulumi.BoolOutput)
}

type SystemGroupArrayOutput struct{ *pulumi.OutputState }

func (SystemGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemGroup)(nil)).Elem()
}

func (o SystemGroupArrayOutput) ToSystemGroupArrayOutput() SystemGroupArrayOutput {
	return o
}

func (o SystemGroupArrayOutput) ToSystemGroupArrayOutputWithContext(ctx context.Context) SystemGroupArrayOutput {
	return o
}

func (o SystemGroupArrayOutput) Index(i pulumi.IntInput) SystemGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemGroup {
		return vs[0].([]*SystemGroup)[vs[1].(int)]
	}).(SystemGroupOutput)
}

type SystemGroupMapOutput struct{ *pulumi.OutputState }

func (SystemGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemGroup)(nil)).Elem()
}

func (o SystemGroupMapOutput) ToSystemGroupMapOutput() SystemGroupMapOutput {
	return o
}

func (o SystemGroupMapOutput) ToSystemGroupMapOutputWithContext(ctx context.Context) SystemGroupMapOutput {
	return o
}

func (o SystemGroupMapOutput) MapIndex(k pulumi.StringInput) SystemGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemGroup {
		return vs[0].(map[string]*SystemGroup)[vs[1].(string)]
	}).(SystemGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGroupInput)(nil)).Elem(), &SystemGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGroupArrayInput)(nil)).Elem(), SystemGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemGroupMapInput)(nil)).Elem(), SystemGroupMap{})
	pulumi.RegisterOutputType(SystemGroupOutput{})
	pulumi.RegisterOutputType(SystemGroupArrayOutput{})
	pulumi.RegisterOutputType(SystemGroupMapOutput{})
}

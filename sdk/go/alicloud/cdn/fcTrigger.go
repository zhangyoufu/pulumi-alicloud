// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CDN Fc Trigger resource.
//
// For information about CDN Fc Trigger and how to use it, see [What is Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger).
//
// > **NOTE:** Available in v1.165.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAccount, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultRegions, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cdn.NewFcTrigger(ctx, "example", &cdn.FcTriggerArgs{
//				EventMetaName:    pulumi.String("LogFileCreated"),
//				EventMetaVersion: pulumi.String("1.0.0"),
//				Notes:            pulumi.String("example_value"),
//				RoleArn:          pulumi.String(fmt.Sprintf("acs:ram::%v:role/aliyuncdneventnotificationrole", defaultAccount.Id)),
//				SourceArn:        pulumi.String(fmt.Sprintf("acs:cdn:*:%v:domain/example.com", defaultAccount.Id)),
//				TriggerArn:       pulumi.String(fmt.Sprintf("acs:fc:%v:%v:services/FCTestService/functions/printEvent/triggers/testtrigger", defaultRegions.Regions[0].Id, defaultAccount.Id)),
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
// CDN Fc Trigger can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cdn/fcTrigger:FcTrigger example <trigger_arn>
// ```
type FcTrigger struct {
	pulumi.CustomResourceState

	// The name of the Event.
	EventMetaName pulumi.StringOutput `pulumi:"eventMetaName"`
	// The version of the Event.
	EventMetaVersion pulumi.StringOutput `pulumi:"eventMetaVersion"`
	// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
	FunctionArn pulumi.StringPtrOutput `pulumi:"functionArn"`
	// The Note information.
	Notes pulumi.StringOutput `pulumi:"notes"`
	// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
	SourceArn pulumi.StringOutput `pulumi:"sourceArn"`
	// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
	TriggerArn pulumi.StringOutput `pulumi:"triggerArn"`
}

// NewFcTrigger registers a new resource with the given unique name, arguments, and options.
func NewFcTrigger(ctx *pulumi.Context,
	name string, args *FcTriggerArgs, opts ...pulumi.ResourceOption) (*FcTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventMetaName == nil {
		return nil, errors.New("invalid value for required argument 'EventMetaName'")
	}
	if args.EventMetaVersion == nil {
		return nil, errors.New("invalid value for required argument 'EventMetaVersion'")
	}
	if args.Notes == nil {
		return nil, errors.New("invalid value for required argument 'Notes'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.SourceArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceArn'")
	}
	if args.TriggerArn == nil {
		return nil, errors.New("invalid value for required argument 'TriggerArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FcTrigger
	err := ctx.RegisterResource("alicloud:cdn/fcTrigger:FcTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFcTrigger gets an existing FcTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFcTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FcTriggerState, opts ...pulumi.ResourceOption) (*FcTrigger, error) {
	var resource FcTrigger
	err := ctx.ReadResource("alicloud:cdn/fcTrigger:FcTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FcTrigger resources.
type fcTriggerState struct {
	// The name of the Event.
	EventMetaName *string `pulumi:"eventMetaName"`
	// The version of the Event.
	EventMetaVersion *string `pulumi:"eventMetaVersion"`
	// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
	FunctionArn *string `pulumi:"functionArn"`
	// The Note information.
	Notes *string `pulumi:"notes"`
	// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
	RoleArn *string `pulumi:"roleArn"`
	// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
	SourceArn *string `pulumi:"sourceArn"`
	// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
	TriggerArn *string `pulumi:"triggerArn"`
}

type FcTriggerState struct {
	// The name of the Event.
	EventMetaName pulumi.StringPtrInput
	// The version of the Event.
	EventMetaVersion pulumi.StringPtrInput
	// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
	FunctionArn pulumi.StringPtrInput
	// The Note information.
	Notes pulumi.StringPtrInput
	// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
	RoleArn pulumi.StringPtrInput
	// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
	SourceArn pulumi.StringPtrInput
	// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
	TriggerArn pulumi.StringPtrInput
}

func (FcTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fcTriggerState)(nil)).Elem()
}

type fcTriggerArgs struct {
	// The name of the Event.
	EventMetaName string `pulumi:"eventMetaName"`
	// The version of the Event.
	EventMetaVersion string `pulumi:"eventMetaVersion"`
	// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
	FunctionArn *string `pulumi:"functionArn"`
	// The Note information.
	Notes string `pulumi:"notes"`
	// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
	RoleArn string `pulumi:"roleArn"`
	// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
	SourceArn string `pulumi:"sourceArn"`
	// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
	TriggerArn string `pulumi:"triggerArn"`
}

// The set of arguments for constructing a FcTrigger resource.
type FcTriggerArgs struct {
	// The name of the Event.
	EventMetaName pulumi.StringInput
	// The version of the Event.
	EventMetaVersion pulumi.StringInput
	// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
	FunctionArn pulumi.StringPtrInput
	// The Note information.
	Notes pulumi.StringInput
	// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
	RoleArn pulumi.StringInput
	// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
	SourceArn pulumi.StringInput
	// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
	TriggerArn pulumi.StringInput
}

func (FcTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fcTriggerArgs)(nil)).Elem()
}

type FcTriggerInput interface {
	pulumi.Input

	ToFcTriggerOutput() FcTriggerOutput
	ToFcTriggerOutputWithContext(ctx context.Context) FcTriggerOutput
}

func (*FcTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**FcTrigger)(nil)).Elem()
}

func (i *FcTrigger) ToFcTriggerOutput() FcTriggerOutput {
	return i.ToFcTriggerOutputWithContext(context.Background())
}

func (i *FcTrigger) ToFcTriggerOutputWithContext(ctx context.Context) FcTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FcTriggerOutput)
}

// FcTriggerArrayInput is an input type that accepts FcTriggerArray and FcTriggerArrayOutput values.
// You can construct a concrete instance of `FcTriggerArrayInput` via:
//
//	FcTriggerArray{ FcTriggerArgs{...} }
type FcTriggerArrayInput interface {
	pulumi.Input

	ToFcTriggerArrayOutput() FcTriggerArrayOutput
	ToFcTriggerArrayOutputWithContext(context.Context) FcTriggerArrayOutput
}

type FcTriggerArray []FcTriggerInput

func (FcTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FcTrigger)(nil)).Elem()
}

func (i FcTriggerArray) ToFcTriggerArrayOutput() FcTriggerArrayOutput {
	return i.ToFcTriggerArrayOutputWithContext(context.Background())
}

func (i FcTriggerArray) ToFcTriggerArrayOutputWithContext(ctx context.Context) FcTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FcTriggerArrayOutput)
}

// FcTriggerMapInput is an input type that accepts FcTriggerMap and FcTriggerMapOutput values.
// You can construct a concrete instance of `FcTriggerMapInput` via:
//
//	FcTriggerMap{ "key": FcTriggerArgs{...} }
type FcTriggerMapInput interface {
	pulumi.Input

	ToFcTriggerMapOutput() FcTriggerMapOutput
	ToFcTriggerMapOutputWithContext(context.Context) FcTriggerMapOutput
}

type FcTriggerMap map[string]FcTriggerInput

func (FcTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FcTrigger)(nil)).Elem()
}

func (i FcTriggerMap) ToFcTriggerMapOutput() FcTriggerMapOutput {
	return i.ToFcTriggerMapOutputWithContext(context.Background())
}

func (i FcTriggerMap) ToFcTriggerMapOutputWithContext(ctx context.Context) FcTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FcTriggerMapOutput)
}

type FcTriggerOutput struct{ *pulumi.OutputState }

func (FcTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FcTrigger)(nil)).Elem()
}

func (o FcTriggerOutput) ToFcTriggerOutput() FcTriggerOutput {
	return o
}

func (o FcTriggerOutput) ToFcTriggerOutputWithContext(ctx context.Context) FcTriggerOutput {
	return o
}

// The name of the Event.
func (o FcTriggerOutput) EventMetaName() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.EventMetaName }).(pulumi.StringOutput)
}

// The version of the Event.
func (o FcTriggerOutput) EventMetaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.EventMetaVersion }).(pulumi.StringOutput)
}

// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
func (o FcTriggerOutput) FunctionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringPtrOutput { return v.FunctionArn }).(pulumi.StringPtrOutput)
}

// The Note information.
func (o FcTriggerOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
func (o FcTriggerOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
func (o FcTriggerOutput) SourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.SourceArn }).(pulumi.StringOutput)
}

// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/en/cdn/developer-reference/api-cdn-2018-05-10-addfctrigger) for more details.
func (o FcTriggerOutput) TriggerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FcTrigger) pulumi.StringOutput { return v.TriggerArn }).(pulumi.StringOutput)
}

type FcTriggerArrayOutput struct{ *pulumi.OutputState }

func (FcTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FcTrigger)(nil)).Elem()
}

func (o FcTriggerArrayOutput) ToFcTriggerArrayOutput() FcTriggerArrayOutput {
	return o
}

func (o FcTriggerArrayOutput) ToFcTriggerArrayOutputWithContext(ctx context.Context) FcTriggerArrayOutput {
	return o
}

func (o FcTriggerArrayOutput) Index(i pulumi.IntInput) FcTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FcTrigger {
		return vs[0].([]*FcTrigger)[vs[1].(int)]
	}).(FcTriggerOutput)
}

type FcTriggerMapOutput struct{ *pulumi.OutputState }

func (FcTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FcTrigger)(nil)).Elem()
}

func (o FcTriggerMapOutput) ToFcTriggerMapOutput() FcTriggerMapOutput {
	return o
}

func (o FcTriggerMapOutput) ToFcTriggerMapOutputWithContext(ctx context.Context) FcTriggerMapOutput {
	return o
}

func (o FcTriggerMapOutput) MapIndex(k pulumi.StringInput) FcTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FcTrigger {
		return vs[0].(map[string]*FcTrigger)[vs[1].(string)]
	}).(FcTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FcTriggerInput)(nil)).Elem(), &FcTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*FcTriggerArrayInput)(nil)).Elem(), FcTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FcTriggerMapInput)(nil)).Elem(), FcTriggerMap{})
	pulumi.RegisterOutputType(FcTriggerOutput{})
	pulumi.RegisterOutputType(FcTriggerArrayOutput{})
	pulumi.RegisterOutputType(FcTriggerMapOutput{})
}

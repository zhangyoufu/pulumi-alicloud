// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Alicloud Serverless App Engine (SAE) Application Load Balancer Attachment resource.
//
// For information about Serverless App Engine (SAE) Load Balancer Internet Attachment and how to use it, see [sae.LoadBalancerInternet](https://help.aliyun.com/document_detail/126360.html).
//
// > **NOTE:** Available in v1.164.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sae"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sae.NewLoadBalancerInternet(ctx, "example", &sae.LoadBalancerInternetArgs{
//				AppId: pulumi.String("your_application_id"),
//				Internets: sae.LoadBalancerInternetInternetArray{
//					&sae.LoadBalancerInternetInternetArgs{
//						Port:       pulumi.Int(80),
//						Protocol:   pulumi.String("TCP"),
//						TargetPort: pulumi.Int(8080),
//					},
//				},
//				InternetSlbId: pulumi.String("your_internet_slb_id"),
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
// The resource can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sae/loadBalancerInternet:LoadBalancerInternet example <id>
//
// ```
type LoadBalancerInternet struct {
	pulumi.CustomResourceState

	// The target application ID that needs to be bound to the SLB.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Use designated public network SLBs that have been purchased to support non-shared instances.
	InternetIp pulumi.StringOutput `pulumi:"internetIp"`
	// The internet SLB ID.
	InternetSlbId pulumi.StringPtrOutput `pulumi:"internetSlbId"`
	// The bound private network SLB. See the following `Block internet`.
	Internets LoadBalancerInternetInternetArrayOutput `pulumi:"internets"`
}

// NewLoadBalancerInternet registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerInternet(ctx *pulumi.Context,
	name string, args *LoadBalancerInternetArgs, opts ...pulumi.ResourceOption) (*LoadBalancerInternet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Internets == nil {
		return nil, errors.New("invalid value for required argument 'Internets'")
	}
	var resource LoadBalancerInternet
	err := ctx.RegisterResource("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerInternet gets an existing LoadBalancerInternet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerInternet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerInternetState, opts ...pulumi.ResourceOption) (*LoadBalancerInternet, error) {
	var resource LoadBalancerInternet
	err := ctx.ReadResource("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerInternet resources.
type loadBalancerInternetState struct {
	// The target application ID that needs to be bound to the SLB.
	AppId *string `pulumi:"appId"`
	// Use designated public network SLBs that have been purchased to support non-shared instances.
	InternetIp *string `pulumi:"internetIp"`
	// The internet SLB ID.
	InternetSlbId *string `pulumi:"internetSlbId"`
	// The bound private network SLB. See the following `Block internet`.
	Internets []LoadBalancerInternetInternet `pulumi:"internets"`
}

type LoadBalancerInternetState struct {
	// The target application ID that needs to be bound to the SLB.
	AppId pulumi.StringPtrInput
	// Use designated public network SLBs that have been purchased to support non-shared instances.
	InternetIp pulumi.StringPtrInput
	// The internet SLB ID.
	InternetSlbId pulumi.StringPtrInput
	// The bound private network SLB. See the following `Block internet`.
	Internets LoadBalancerInternetInternetArrayInput
}

func (LoadBalancerInternetState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerInternetState)(nil)).Elem()
}

type loadBalancerInternetArgs struct {
	// The target application ID that needs to be bound to the SLB.
	AppId string `pulumi:"appId"`
	// The internet SLB ID.
	InternetSlbId *string `pulumi:"internetSlbId"`
	// The bound private network SLB. See the following `Block internet`.
	Internets []LoadBalancerInternetInternet `pulumi:"internets"`
}

// The set of arguments for constructing a LoadBalancerInternet resource.
type LoadBalancerInternetArgs struct {
	// The target application ID that needs to be bound to the SLB.
	AppId pulumi.StringInput
	// The internet SLB ID.
	InternetSlbId pulumi.StringPtrInput
	// The bound private network SLB. See the following `Block internet`.
	Internets LoadBalancerInternetInternetArrayInput
}

func (LoadBalancerInternetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerInternetArgs)(nil)).Elem()
}

type LoadBalancerInternetInput interface {
	pulumi.Input

	ToLoadBalancerInternetOutput() LoadBalancerInternetOutput
	ToLoadBalancerInternetOutputWithContext(ctx context.Context) LoadBalancerInternetOutput
}

func (*LoadBalancerInternet) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerInternet)(nil)).Elem()
}

func (i *LoadBalancerInternet) ToLoadBalancerInternetOutput() LoadBalancerInternetOutput {
	return i.ToLoadBalancerInternetOutputWithContext(context.Background())
}

func (i *LoadBalancerInternet) ToLoadBalancerInternetOutputWithContext(ctx context.Context) LoadBalancerInternetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerInternetOutput)
}

// LoadBalancerInternetArrayInput is an input type that accepts LoadBalancerInternetArray and LoadBalancerInternetArrayOutput values.
// You can construct a concrete instance of `LoadBalancerInternetArrayInput` via:
//
//	LoadBalancerInternetArray{ LoadBalancerInternetArgs{...} }
type LoadBalancerInternetArrayInput interface {
	pulumi.Input

	ToLoadBalancerInternetArrayOutput() LoadBalancerInternetArrayOutput
	ToLoadBalancerInternetArrayOutputWithContext(context.Context) LoadBalancerInternetArrayOutput
}

type LoadBalancerInternetArray []LoadBalancerInternetInput

func (LoadBalancerInternetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerInternet)(nil)).Elem()
}

func (i LoadBalancerInternetArray) ToLoadBalancerInternetArrayOutput() LoadBalancerInternetArrayOutput {
	return i.ToLoadBalancerInternetArrayOutputWithContext(context.Background())
}

func (i LoadBalancerInternetArray) ToLoadBalancerInternetArrayOutputWithContext(ctx context.Context) LoadBalancerInternetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerInternetArrayOutput)
}

// LoadBalancerInternetMapInput is an input type that accepts LoadBalancerInternetMap and LoadBalancerInternetMapOutput values.
// You can construct a concrete instance of `LoadBalancerInternetMapInput` via:
//
//	LoadBalancerInternetMap{ "key": LoadBalancerInternetArgs{...} }
type LoadBalancerInternetMapInput interface {
	pulumi.Input

	ToLoadBalancerInternetMapOutput() LoadBalancerInternetMapOutput
	ToLoadBalancerInternetMapOutputWithContext(context.Context) LoadBalancerInternetMapOutput
}

type LoadBalancerInternetMap map[string]LoadBalancerInternetInput

func (LoadBalancerInternetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerInternet)(nil)).Elem()
}

func (i LoadBalancerInternetMap) ToLoadBalancerInternetMapOutput() LoadBalancerInternetMapOutput {
	return i.ToLoadBalancerInternetMapOutputWithContext(context.Background())
}

func (i LoadBalancerInternetMap) ToLoadBalancerInternetMapOutputWithContext(ctx context.Context) LoadBalancerInternetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerInternetMapOutput)
}

type LoadBalancerInternetOutput struct{ *pulumi.OutputState }

func (LoadBalancerInternetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerInternet)(nil)).Elem()
}

func (o LoadBalancerInternetOutput) ToLoadBalancerInternetOutput() LoadBalancerInternetOutput {
	return o
}

func (o LoadBalancerInternetOutput) ToLoadBalancerInternetOutputWithContext(ctx context.Context) LoadBalancerInternetOutput {
	return o
}

// The target application ID that needs to be bound to the SLB.
func (o LoadBalancerInternetOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerInternet) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// Use designated public network SLBs that have been purchased to support non-shared instances.
func (o LoadBalancerInternetOutput) InternetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerInternet) pulumi.StringOutput { return v.InternetIp }).(pulumi.StringOutput)
}

// The internet SLB ID.
func (o LoadBalancerInternetOutput) InternetSlbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerInternet) pulumi.StringPtrOutput { return v.InternetSlbId }).(pulumi.StringPtrOutput)
}

// The bound private network SLB. See the following `Block internet`.
func (o LoadBalancerInternetOutput) Internets() LoadBalancerInternetInternetArrayOutput {
	return o.ApplyT(func(v *LoadBalancerInternet) LoadBalancerInternetInternetArrayOutput { return v.Internets }).(LoadBalancerInternetInternetArrayOutput)
}

type LoadBalancerInternetArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerInternetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerInternet)(nil)).Elem()
}

func (o LoadBalancerInternetArrayOutput) ToLoadBalancerInternetArrayOutput() LoadBalancerInternetArrayOutput {
	return o
}

func (o LoadBalancerInternetArrayOutput) ToLoadBalancerInternetArrayOutputWithContext(ctx context.Context) LoadBalancerInternetArrayOutput {
	return o
}

func (o LoadBalancerInternetArrayOutput) Index(i pulumi.IntInput) LoadBalancerInternetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancerInternet {
		return vs[0].([]*LoadBalancerInternet)[vs[1].(int)]
	}).(LoadBalancerInternetOutput)
}

type LoadBalancerInternetMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerInternetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerInternet)(nil)).Elem()
}

func (o LoadBalancerInternetMapOutput) ToLoadBalancerInternetMapOutput() LoadBalancerInternetMapOutput {
	return o
}

func (o LoadBalancerInternetMapOutput) ToLoadBalancerInternetMapOutputWithContext(ctx context.Context) LoadBalancerInternetMapOutput {
	return o
}

func (o LoadBalancerInternetMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerInternetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancerInternet {
		return vs[0].(map[string]*LoadBalancerInternet)[vs[1].(string)]
	}).(LoadBalancerInternetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInternetInput)(nil)).Elem(), &LoadBalancerInternet{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInternetArrayInput)(nil)).Elem(), LoadBalancerInternetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInternetMapInput)(nil)).Elem(), LoadBalancerInternetMap{})
	pulumi.RegisterOutputType(LoadBalancerInternetOutput{})
	pulumi.RegisterOutputType(LoadBalancerInternetArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerInternetMapOutput{})
}

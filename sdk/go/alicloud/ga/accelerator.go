// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Ga Accelerator resource.
//
// For information about Ga Accelerator and how to use it, see [What is Accelerator](https://help.aliyun.com/document_detail/153235.html).
//
// > **NOTE:** Available in v1.111.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ga"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ga.NewAccelerator(ctx, "example", &ga.AcceleratorArgs{
// 			AutoUseCoupon: pulumi.Bool(true),
// 			Duration:      pulumi.Int(1),
// 			Spec:          pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Ga Accelerator can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ga/accelerator:Accelerator example <accelerator_id>
// ```
type Accelerator struct {
	pulumi.CustomResourceState

	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrOutput `pulumi:"acceleratorName"`
	// The auto use coupon.
	AutoUseCoupon pulumi.BoolPtrOutput `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The duration. The value range is 1-9.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringOutput `pulumi:"spec"`
	// The status of the GA instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAccelerator registers a new resource with the given unique name, arguments, and options.
func NewAccelerator(ctx *pulumi.Context,
	name string, args *AcceleratorArgs, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	if args == nil || args.Duration == nil {
		return nil, errors.New("missing required argument 'Duration'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &AcceleratorArgs{}
	}
	var resource Accelerator
	err := ctx.RegisterResource("alicloud:ga/accelerator:Accelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccelerator gets an existing Accelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcceleratorState, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	var resource Accelerator
	err := ctx.ReadResource("alicloud:ga/accelerator:Accelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accelerator resources.
type acceleratorState struct {
	// The Name of the GA instance.
	AcceleratorName *string `pulumi:"acceleratorName"`
	// The auto use coupon.
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description *string `pulumi:"description"`
	// The duration. The value range is 1-9.
	Duration *int `pulumi:"duration"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec *string `pulumi:"spec"`
	// The status of the GA instance.
	Status *string `pulumi:"status"`
}

type AcceleratorState struct {
	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrInput
	// The auto use coupon.
	AutoUseCoupon pulumi.BoolPtrInput
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrInput
	// The duration. The value range is 1-9.
	Duration pulumi.IntPtrInput
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringPtrInput
	// The status of the GA instance.
	Status pulumi.StringPtrInput
}

func (AcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorState)(nil)).Elem()
}

type acceleratorArgs struct {
	// The Name of the GA instance.
	AcceleratorName *string `pulumi:"acceleratorName"`
	// The auto use coupon.
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description *string `pulumi:"description"`
	// The duration. The value range is 1-9.
	Duration int `pulumi:"duration"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec string `pulumi:"spec"`
}

// The set of arguments for constructing a Accelerator resource.
type AcceleratorArgs struct {
	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrInput
	// The auto use coupon.
	AutoUseCoupon pulumi.BoolPtrInput
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrInput
	// The duration. The value range is 1-9.
	Duration pulumi.IntInput
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringInput
}

func (AcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorArgs)(nil)).Elem()
}

type AcceleratorInput interface {
	pulumi.Input

	ToAcceleratorOutput() AcceleratorOutput
	ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput
}

func (Accelerator) ElementType() reflect.Type {
	return reflect.TypeOf((*Accelerator)(nil)).Elem()
}

func (i Accelerator) ToAcceleratorOutput() AcceleratorOutput {
	return i.ToAcceleratorOutputWithContext(context.Background())
}

func (i Accelerator) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorOutput)
}

type AcceleratorOutput struct {
	*pulumi.OutputState
}

func (AcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorOutput)(nil)).Elem()
}

func (o AcceleratorOutput) ToAcceleratorOutput() AcceleratorOutput {
	return o
}

func (o AcceleratorOutput) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AcceleratorOutput{})
}

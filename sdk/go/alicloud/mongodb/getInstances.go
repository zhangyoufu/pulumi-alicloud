// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `mongodb.getInstances` data source provides a collection of MongoDB instances available in Alicloud account.
// Filters support regular expression for the instance name, engine or instance type.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.GetInstances(ctx, &mongodb.GetInstancesArgs{
//				NameRegex:        pulumi.StringRef("dds-.+\\d+"),
//				InstanceType:     pulumi.StringRef("replicate"),
//				InstanceClass:    pulumi.StringRef("dds.mongo.mid"),
//				AvailabilityZone: pulumi.StringRef("eu-central-1a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:mongodb/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Instance availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ids list of MongoDB instances
	Ids []string `pulumi:"ids"`
	// Sizing of the instance to be queried.
	InstanceClass *string `pulumi:"instanceClass"`
	// Type of the instance to be queried. If it is set to `sharding`, the sharded cluster instances are listed. If it is set to `replicate`, replica set instances are listed. Default value `replicate`.
	InstanceType *string `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex *string `pulumi:"nameRegex"`
	// The name of file that can save the collection of instances after running `pulumi preview`.
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// Instance availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ids list of MongoDB instances
	Ids []string `pulumi:"ids"`
	// Sizing of the MongoDB instance.
	InstanceClass *string `pulumi:"instanceClass"`
	// Instance type. Optional values `sharding` or `replicate`.
	InstanceType *string `pulumi:"instanceType"`
	// A list of MongoDB instances. Its every element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// The names list of MongoDB instances
	Names      []string          `pulumi:"names"`
	OutputFile *string           `pulumi:"outputFile"`
	Tags       map[string]string `pulumi:"tags"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// Instance availability zone.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// The ids list of MongoDB instances
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Sizing of the instance to be queried.
	InstanceClass pulumi.StringPtrInput `pulumi:"instanceClass"`
	// Type of the instance to be queried. If it is set to `sharding`, the sharded cluster instances are listed. If it is set to `replicate`, replica set instances are listed. Default value `replicate`.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The name of file that can save the collection of instances after running `pulumi preview`.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// Instance availability zone.
func (o GetInstancesResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ids list of MongoDB instances
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Sizing of the MongoDB instance.
func (o GetInstancesResultOutput) InstanceClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceClass }).(pulumi.StringPtrOutput)
}

// Instance type. Optional values `sharding` or `replicate`.
func (o GetInstancesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// A list of MongoDB instances. Its every element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The names list of MongoDB instances
func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}

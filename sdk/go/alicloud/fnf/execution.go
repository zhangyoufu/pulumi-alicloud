// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fnf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Serverless Workflow Execution resource.
//
// For information about Serverless Workflow Execution and how to use it, see [What is Execution](https://www.alibabacloud.com/help/en/doc-detail/122628.html).
//
// > **NOTE:** Available since v1.149.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fnf"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example-fnfflow"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := ram.NewRole(ctx, "default", &ram.RoleArgs{
//				Name: pulumi.String(name),
//				Document: pulumi.String(`  {
//	    "Statement": [
//	      {
//	        "Action": "sts:AssumeRole",
//	        "Effect": "Allow",
//	        "Principal": {
//	          "Service": [
//	            "fnf.aliyuncs.com"
//	          ]
//	        }
//	      }
//	    ],
//	    "Version": "1"
//	  }
//
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			defaultFlow, err := fnf.NewFlow(ctx, "default", &fnf.FlowArgs{
//				Definition: pulumi.String(`  version: v1beta1
//	  type: flow
//	  steps:
//	    - type: wait
//	      name: custom_wait
//	      duration: $.wait
//
// `),
//
//				RoleArn:     _default.Arn,
//				Description: pulumi.String("Test for terraform fnf_flow."),
//				Name:        pulumi.String(name),
//				Type:        pulumi.String("FDL"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fnf.NewExecution(ctx, "default", &fnf.ExecutionArgs{
//				ExecutionName: pulumi.String(name),
//				FlowName:      defaultFlow.Name,
//				Input:         pulumi.String("{\"wait\": 600}"),
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
// Serverless Workflow Execution can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:fnf/execution:Execution example <flow_name>:<execution_name>
// ```
type Execution struct {
	pulumi.CustomResourceState

	// The name of the execution.
	ExecutionName pulumi.StringOutput `pulumi:"executionName"`
	// The name of the flow.
	FlowName pulumi.StringOutput `pulumi:"flowName"`
	// The Input information for this execution.
	Input pulumi.StringPtrOutput `pulumi:"input"`
	// The status of the resource. Valid values: `Stopped`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionName == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionName'")
	}
	if args.FlowName == nil {
		return nil, errors.New("invalid value for required argument 'FlowName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Execution
	err := ctx.RegisterResource("alicloud:fnf/execution:Execution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecution gets an existing Execution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionState, opts ...pulumi.ResourceOption) (*Execution, error) {
	var resource Execution
	err := ctx.ReadResource("alicloud:fnf/execution:Execution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Execution resources.
type executionState struct {
	// The name of the execution.
	ExecutionName *string `pulumi:"executionName"`
	// The name of the flow.
	FlowName *string `pulumi:"flowName"`
	// The Input information for this execution.
	Input *string `pulumi:"input"`
	// The status of the resource. Valid values: `Stopped`.
	Status *string `pulumi:"status"`
}

type ExecutionState struct {
	// The name of the execution.
	ExecutionName pulumi.StringPtrInput
	// The name of the flow.
	FlowName pulumi.StringPtrInput
	// The Input information for this execution.
	Input pulumi.StringPtrInput
	// The status of the resource. Valid values: `Stopped`.
	Status pulumi.StringPtrInput
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	// The name of the execution.
	ExecutionName string `pulumi:"executionName"`
	// The name of the flow.
	FlowName string `pulumi:"flowName"`
	// The Input information for this execution.
	Input *string `pulumi:"input"`
	// The status of the resource. Valid values: `Stopped`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	// The name of the execution.
	ExecutionName pulumi.StringInput
	// The name of the flow.
	FlowName pulumi.StringInput
	// The Input information for this execution.
	Input pulumi.StringPtrInput
	// The status of the resource. Valid values: `Stopped`.
	Status pulumi.StringPtrInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}

type ExecutionInput interface {
	pulumi.Input

	ToExecutionOutput() ExecutionOutput
	ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput
}

func (*Execution) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (i *Execution) ToExecutionOutput() ExecutionOutput {
	return i.ToExecutionOutputWithContext(context.Background())
}

func (i *Execution) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionOutput)
}

// ExecutionArrayInput is an input type that accepts ExecutionArray and ExecutionArrayOutput values.
// You can construct a concrete instance of `ExecutionArrayInput` via:
//
//	ExecutionArray{ ExecutionArgs{...} }
type ExecutionArrayInput interface {
	pulumi.Input

	ToExecutionArrayOutput() ExecutionArrayOutput
	ToExecutionArrayOutputWithContext(context.Context) ExecutionArrayOutput
}

type ExecutionArray []ExecutionInput

func (ExecutionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Execution)(nil)).Elem()
}

func (i ExecutionArray) ToExecutionArrayOutput() ExecutionArrayOutput {
	return i.ToExecutionArrayOutputWithContext(context.Background())
}

func (i ExecutionArray) ToExecutionArrayOutputWithContext(ctx context.Context) ExecutionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionArrayOutput)
}

// ExecutionMapInput is an input type that accepts ExecutionMap and ExecutionMapOutput values.
// You can construct a concrete instance of `ExecutionMapInput` via:
//
//	ExecutionMap{ "key": ExecutionArgs{...} }
type ExecutionMapInput interface {
	pulumi.Input

	ToExecutionMapOutput() ExecutionMapOutput
	ToExecutionMapOutputWithContext(context.Context) ExecutionMapOutput
}

type ExecutionMap map[string]ExecutionInput

func (ExecutionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Execution)(nil)).Elem()
}

func (i ExecutionMap) ToExecutionMapOutput() ExecutionMapOutput {
	return i.ToExecutionMapOutputWithContext(context.Background())
}

func (i ExecutionMap) ToExecutionMapOutputWithContext(ctx context.Context) ExecutionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionMapOutput)
}

type ExecutionOutput struct{ *pulumi.OutputState }

func (ExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (o ExecutionOutput) ToExecutionOutput() ExecutionOutput {
	return o
}

func (o ExecutionOutput) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return o
}

// The name of the execution.
func (o ExecutionOutput) ExecutionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.ExecutionName }).(pulumi.StringOutput)
}

// The name of the flow.
func (o ExecutionOutput) FlowName() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.FlowName }).(pulumi.StringOutput)
}

// The Input information for this execution.
func (o ExecutionOutput) Input() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringPtrOutput { return v.Input }).(pulumi.StringPtrOutput)
}

// The status of the resource. Valid values: `Stopped`.
func (o ExecutionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ExecutionArrayOutput struct{ *pulumi.OutputState }

func (ExecutionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Execution)(nil)).Elem()
}

func (o ExecutionArrayOutput) ToExecutionArrayOutput() ExecutionArrayOutput {
	return o
}

func (o ExecutionArrayOutput) ToExecutionArrayOutputWithContext(ctx context.Context) ExecutionArrayOutput {
	return o
}

func (o ExecutionArrayOutput) Index(i pulumi.IntInput) ExecutionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Execution {
		return vs[0].([]*Execution)[vs[1].(int)]
	}).(ExecutionOutput)
}

type ExecutionMapOutput struct{ *pulumi.OutputState }

func (ExecutionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Execution)(nil)).Elem()
}

func (o ExecutionMapOutput) ToExecutionMapOutput() ExecutionMapOutput {
	return o
}

func (o ExecutionMapOutput) ToExecutionMapOutputWithContext(ctx context.Context) ExecutionMapOutput {
	return o
}

func (o ExecutionMapOutput) MapIndex(k pulumi.StringInput) ExecutionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Execution {
		return vs[0].(map[string]*Execution)[vs[1].(string)]
	}).(ExecutionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionInput)(nil)).Elem(), &Execution{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionArrayInput)(nil)).Elem(), ExecutionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionMapInput)(nil)).Elem(), ExecutionMap{})
	pulumi.RegisterOutputType(ExecutionOutput{})
	pulumi.RegisterOutputType(ExecutionArrayOutput{})
	pulumi.RegisterOutputType(ExecutionMapOutput{})
}

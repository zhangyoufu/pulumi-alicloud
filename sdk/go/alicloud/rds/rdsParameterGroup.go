// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RDS Parameter Group resource.
//
// For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).
//
// > **NOTE:** Available since v1.119.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := rds.NewRdsParameterGroup(ctx, "default", &rds.RdsParameterGroupArgs{
//				Engine:        pulumi.String("mysql"),
//				EngineVersion: pulumi.String("5.7"),
//				ParamDetails: rds.RdsParameterGroupParamDetailArray{
//					&rds.RdsParameterGroupParamDetailArgs{
//						ParamName:  pulumi.String("back_log"),
//						ParamValue: pulumi.String("4000"),
//					},
//					&rds.RdsParameterGroupParamDetailArgs{
//						ParamName:  pulumi.String("wait_timeout"),
//						ParamValue: pulumi.String("86460"),
//					},
//				},
//				ParameterGroupDesc: pulumi.String(name),
//				ParameterGroupName: pulumi.String(name),
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
// RDS Parameter Group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:rds/rdsParameterGroup:RdsParameterGroup example <id>
// ```
type RdsParameterGroup struct {
	pulumi.CustomResourceState

	// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Parameter list. See `paramDetail` below.
	ParamDetails RdsParameterGroupParamDetailArrayOutput `pulumi:"paramDetails"`
	// The description of the parameter template.
	ParameterGroupDesc pulumi.StringPtrOutput `pulumi:"parameterGroupDesc"`
	// The name of the parameter template.
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
}

// NewRdsParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewRdsParameterGroup(ctx *pulumi.Context,
	name string, args *RdsParameterGroupArgs, opts ...pulumi.ResourceOption) (*RdsParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.ParamDetails == nil {
		return nil, errors.New("invalid value for required argument 'ParamDetails'")
	}
	if args.ParameterGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ParameterGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsParameterGroup
	err := ctx.RegisterResource("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsParameterGroup gets an existing RdsParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsParameterGroupState, opts ...pulumi.ResourceOption) (*RdsParameterGroup, error) {
	var resource RdsParameterGroup
	err := ctx.ReadResource("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsParameterGroup resources.
type rdsParameterGroupState struct {
	// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
	Engine *string `pulumi:"engine"`
	// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Parameter list. See `paramDetail` below.
	ParamDetails []RdsParameterGroupParamDetail `pulumi:"paramDetails"`
	// The description of the parameter template.
	ParameterGroupDesc *string `pulumi:"parameterGroupDesc"`
	// The name of the parameter template.
	ParameterGroupName *string `pulumi:"parameterGroupName"`
}

type RdsParameterGroupState struct {
	// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
	Engine pulumi.StringPtrInput
	// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
	EngineVersion pulumi.StringPtrInput
	// Parameter list. See `paramDetail` below.
	ParamDetails RdsParameterGroupParamDetailArrayInput
	// The description of the parameter template.
	ParameterGroupDesc pulumi.StringPtrInput
	// The name of the parameter template.
	ParameterGroupName pulumi.StringPtrInput
}

func (RdsParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsParameterGroupState)(nil)).Elem()
}

type rdsParameterGroupArgs struct {
	// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
	Engine string `pulumi:"engine"`
	// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
	EngineVersion string `pulumi:"engineVersion"`
	// Parameter list. See `paramDetail` below.
	ParamDetails []RdsParameterGroupParamDetail `pulumi:"paramDetails"`
	// The description of the parameter template.
	ParameterGroupDesc *string `pulumi:"parameterGroupDesc"`
	// The name of the parameter template.
	ParameterGroupName string `pulumi:"parameterGroupName"`
}

// The set of arguments for constructing a RdsParameterGroup resource.
type RdsParameterGroupArgs struct {
	// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
	Engine pulumi.StringInput
	// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
	EngineVersion pulumi.StringInput
	// Parameter list. See `paramDetail` below.
	ParamDetails RdsParameterGroupParamDetailArrayInput
	// The description of the parameter template.
	ParameterGroupDesc pulumi.StringPtrInput
	// The name of the parameter template.
	ParameterGroupName pulumi.StringInput
}

func (RdsParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsParameterGroupArgs)(nil)).Elem()
}

type RdsParameterGroupInput interface {
	pulumi.Input

	ToRdsParameterGroupOutput() RdsParameterGroupOutput
	ToRdsParameterGroupOutputWithContext(ctx context.Context) RdsParameterGroupOutput
}

func (*RdsParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsParameterGroup)(nil)).Elem()
}

func (i *RdsParameterGroup) ToRdsParameterGroupOutput() RdsParameterGroupOutput {
	return i.ToRdsParameterGroupOutputWithContext(context.Background())
}

func (i *RdsParameterGroup) ToRdsParameterGroupOutputWithContext(ctx context.Context) RdsParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsParameterGroupOutput)
}

// RdsParameterGroupArrayInput is an input type that accepts RdsParameterGroupArray and RdsParameterGroupArrayOutput values.
// You can construct a concrete instance of `RdsParameterGroupArrayInput` via:
//
//	RdsParameterGroupArray{ RdsParameterGroupArgs{...} }
type RdsParameterGroupArrayInput interface {
	pulumi.Input

	ToRdsParameterGroupArrayOutput() RdsParameterGroupArrayOutput
	ToRdsParameterGroupArrayOutputWithContext(context.Context) RdsParameterGroupArrayOutput
}

type RdsParameterGroupArray []RdsParameterGroupInput

func (RdsParameterGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsParameterGroup)(nil)).Elem()
}

func (i RdsParameterGroupArray) ToRdsParameterGroupArrayOutput() RdsParameterGroupArrayOutput {
	return i.ToRdsParameterGroupArrayOutputWithContext(context.Background())
}

func (i RdsParameterGroupArray) ToRdsParameterGroupArrayOutputWithContext(ctx context.Context) RdsParameterGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsParameterGroupArrayOutput)
}

// RdsParameterGroupMapInput is an input type that accepts RdsParameterGroupMap and RdsParameterGroupMapOutput values.
// You can construct a concrete instance of `RdsParameterGroupMapInput` via:
//
//	RdsParameterGroupMap{ "key": RdsParameterGroupArgs{...} }
type RdsParameterGroupMapInput interface {
	pulumi.Input

	ToRdsParameterGroupMapOutput() RdsParameterGroupMapOutput
	ToRdsParameterGroupMapOutputWithContext(context.Context) RdsParameterGroupMapOutput
}

type RdsParameterGroupMap map[string]RdsParameterGroupInput

func (RdsParameterGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsParameterGroup)(nil)).Elem()
}

func (i RdsParameterGroupMap) ToRdsParameterGroupMapOutput() RdsParameterGroupMapOutput {
	return i.ToRdsParameterGroupMapOutputWithContext(context.Background())
}

func (i RdsParameterGroupMap) ToRdsParameterGroupMapOutputWithContext(ctx context.Context) RdsParameterGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsParameterGroupMapOutput)
}

type RdsParameterGroupOutput struct{ *pulumi.OutputState }

func (RdsParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsParameterGroup)(nil)).Elem()
}

func (o RdsParameterGroupOutput) ToRdsParameterGroupOutput() RdsParameterGroupOutput {
	return o
}

func (o RdsParameterGroupOutput) ToRdsParameterGroupOutputWithContext(ctx context.Context) RdsParameterGroupOutput {
	return o
}

// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
func (o RdsParameterGroupOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsParameterGroup) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
func (o RdsParameterGroupOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsParameterGroup) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// Parameter list. See `paramDetail` below.
func (o RdsParameterGroupOutput) ParamDetails() RdsParameterGroupParamDetailArrayOutput {
	return o.ApplyT(func(v *RdsParameterGroup) RdsParameterGroupParamDetailArrayOutput { return v.ParamDetails }).(RdsParameterGroupParamDetailArrayOutput)
}

// The description of the parameter template.
func (o RdsParameterGroupOutput) ParameterGroupDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsParameterGroup) pulumi.StringPtrOutput { return v.ParameterGroupDesc }).(pulumi.StringPtrOutput)
}

// The name of the parameter template.
func (o RdsParameterGroupOutput) ParameterGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsParameterGroup) pulumi.StringOutput { return v.ParameterGroupName }).(pulumi.StringOutput)
}

type RdsParameterGroupArrayOutput struct{ *pulumi.OutputState }

func (RdsParameterGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsParameterGroup)(nil)).Elem()
}

func (o RdsParameterGroupArrayOutput) ToRdsParameterGroupArrayOutput() RdsParameterGroupArrayOutput {
	return o
}

func (o RdsParameterGroupArrayOutput) ToRdsParameterGroupArrayOutputWithContext(ctx context.Context) RdsParameterGroupArrayOutput {
	return o
}

func (o RdsParameterGroupArrayOutput) Index(i pulumi.IntInput) RdsParameterGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsParameterGroup {
		return vs[0].([]*RdsParameterGroup)[vs[1].(int)]
	}).(RdsParameterGroupOutput)
}

type RdsParameterGroupMapOutput struct{ *pulumi.OutputState }

func (RdsParameterGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsParameterGroup)(nil)).Elem()
}

func (o RdsParameterGroupMapOutput) ToRdsParameterGroupMapOutput() RdsParameterGroupMapOutput {
	return o
}

func (o RdsParameterGroupMapOutput) ToRdsParameterGroupMapOutputWithContext(ctx context.Context) RdsParameterGroupMapOutput {
	return o
}

func (o RdsParameterGroupMapOutput) MapIndex(k pulumi.StringInput) RdsParameterGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsParameterGroup {
		return vs[0].(map[string]*RdsParameterGroup)[vs[1].(string)]
	}).(RdsParameterGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsParameterGroupInput)(nil)).Elem(), &RdsParameterGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsParameterGroupArrayInput)(nil)).Elem(), RdsParameterGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsParameterGroupMapInput)(nil)).Elem(), RdsParameterGroupMap{})
	pulumi.RegisterOutputType(RdsParameterGroupOutput{})
	pulumi.RegisterOutputType(RdsParameterGroupArrayOutput{})
	pulumi.RegisterOutputType(RdsParameterGroupMapOutput{})
}

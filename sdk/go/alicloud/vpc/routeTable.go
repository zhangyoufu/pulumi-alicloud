// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Route Table resource. Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.
//
// For information about VPC Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).
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
//			defaultVpc, err := vpc.NewNetwork(ctx, "defaultVpc", &vpc.NetworkArgs{
//				VpcName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewRouteTable(ctx, "default", &vpc.RouteTableArgs{
//				Description:    pulumi.String("test-description"),
//				VpcId:          defaultVpc.ID(),
//				RouteTableName: pulumi.String(name),
//				AssociateType:  pulumi.String("VSwitch"),
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
// VPC Route Table can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/routeTable:RouteTable example <id>
// ```
type RouteTable struct {
	pulumi.CustomResourceState

	// The type of cloud resource that is bound to the routing table. Value:
	// - **VSwitch**: switch.
	// - **Gateway**:IPv4 Gateway.
	AssociateType pulumi.StringOutput `pulumi:"associateType"`
	// The creation time of the routing table.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the routing table.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource group ID.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The name of the routing table.
	RouteTableName pulumi.StringOutput `pulumi:"routeTableName"`
	// Routing table state.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tag.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteTable
	err := ctx.RegisterResource("alicloud:vpc/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("alicloud:vpc/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// The type of cloud resource that is bound to the routing table. Value:
	// - **VSwitch**: switch.
	// - **Gateway**:IPv4 Gateway.
	AssociateType *string `pulumi:"associateType"`
	// The creation time of the routing table.
	CreateTime *string `pulumi:"createTime"`
	// Description of the routing table.
	Description *string `pulumi:"description"`
	// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	Name *string `pulumi:"name"`
	// Resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The name of the routing table.
	RouteTableName *string `pulumi:"routeTableName"`
	// Routing table state.
	Status *string `pulumi:"status"`
	// The tag.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId *string `pulumi:"vpcId"`
}

type RouteTableState struct {
	// The type of cloud resource that is bound to the routing table. Value:
	// - **VSwitch**: switch.
	// - **Gateway**:IPv4 Gateway.
	AssociateType pulumi.StringPtrInput
	// The creation time of the routing table.
	CreateTime pulumi.StringPtrInput
	// Description of the routing table.
	Description pulumi.StringPtrInput
	// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	Name pulumi.StringPtrInput
	// Resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// The name of the routing table.
	RouteTableName pulumi.StringPtrInput
	// Routing table state.
	Status pulumi.StringPtrInput
	// The tag.
	Tags pulumi.MapInput
	// The ID of VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringPtrInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// The type of cloud resource that is bound to the routing table. Value:
	// - **VSwitch**: switch.
	// - **Gateway**:IPv4 Gateway.
	AssociateType *string `pulumi:"associateType"`
	// Description of the routing table.
	Description *string `pulumi:"description"`
	// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	Name *string `pulumi:"name"`
	// The name of the routing table.
	RouteTableName *string `pulumi:"routeTableName"`
	// The tag.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// The type of cloud resource that is bound to the routing table. Value:
	// - **VSwitch**: switch.
	// - **Gateway**:IPv4 Gateway.
	AssociateType pulumi.StringPtrInput
	// Description of the routing table.
	Description pulumi.StringPtrInput
	// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
	Name pulumi.StringPtrInput
	// The name of the routing table.
	RouteTableName pulumi.StringPtrInput
	// The tag.
	Tags pulumi.MapInput
	// The ID of VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

// RouteTableArrayInput is an input type that accepts RouteTableArray and RouteTableArrayOutput values.
// You can construct a concrete instance of `RouteTableArrayInput` via:
//
//	RouteTableArray{ RouteTableArgs{...} }
type RouteTableArrayInput interface {
	pulumi.Input

	ToRouteTableArrayOutput() RouteTableArrayOutput
	ToRouteTableArrayOutputWithContext(context.Context) RouteTableArrayOutput
}

type RouteTableArray []RouteTableInput

func (RouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (i RouteTableArray) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return i.ToRouteTableArrayOutputWithContext(context.Background())
}

func (i RouteTableArray) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableArrayOutput)
}

// RouteTableMapInput is an input type that accepts RouteTableMap and RouteTableMapOutput values.
// You can construct a concrete instance of `RouteTableMapInput` via:
//
//	RouteTableMap{ "key": RouteTableArgs{...} }
type RouteTableMapInput interface {
	pulumi.Input

	ToRouteTableMapOutput() RouteTableMapOutput
	ToRouteTableMapOutputWithContext(context.Context) RouteTableMapOutput
}

type RouteTableMap map[string]RouteTableInput

func (RouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (i RouteTableMap) ToRouteTableMapOutput() RouteTableMapOutput {
	return i.ToRouteTableMapOutputWithContext(context.Background())
}

func (i RouteTableMap) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableMapOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

// The type of cloud resource that is bound to the routing table. Value:
// - **VSwitch**: switch.
// - **Gateway**:IPv4 Gateway.
func (o RouteTableOutput) AssociateType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.AssociateType }).(pulumi.StringOutput)
}

// The creation time of the routing table.
func (o RouteTableOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the routing table.
func (o RouteTableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
func (o RouteTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource group ID.
func (o RouteTableOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The name of the routing table.
func (o RouteTableOutput) RouteTableName() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.RouteTableName }).(pulumi.StringOutput)
}

// Routing table state.
func (o RouteTableOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tag.
func (o RouteTableOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of VPC.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o RouteTableOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RouteTableArrayOutput struct{ *pulumi.OutputState }

func (RouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) Index(i pulumi.IntInput) RouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].([]*RouteTable)[vs[1].(int)]
	}).(RouteTableOutput)
}

type RouteTableMapOutput struct{ *pulumi.OutputState }

func (RouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (o RouteTableMapOutput) ToRouteTableMapOutput() RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) MapIndex(k pulumi.StringInput) RouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].(map[string]*RouteTable)[vs[1].(string)]
	}).(RouteTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableInput)(nil)).Elem(), &RouteTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableArrayInput)(nil)).Elem(), RouteTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableMapInput)(nil)).Elem(), RouteTableMap{})
	pulumi.RegisterOutputType(RouteTableOutput{})
	pulumi.RegisterOutputType(RouteTableArrayOutput{})
	pulumi.RegisterOutputType(RouteTableMapOutput{})
}

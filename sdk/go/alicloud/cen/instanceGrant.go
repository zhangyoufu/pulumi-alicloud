// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CEN child instance grant resource, which allow you to authorize a VPC or VBR to a CEN of a different account.
//
// For more information about how to use it, see [Attach a network in a different account](https://www.alibabacloud.com/help/doc-detail/73645.htm).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/providers"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := providers.Newalicloud(ctx, "account1", &providers.alicloudArgs{
//				AccessKey: "access123",
//				SecretKey: "secret123",
//			})
//			if err != nil {
//				return err
//			}
//			_, err = providers.Newalicloud(ctx, "account2", &providers.alicloudArgs{
//				AccessKey: "access456",
//				SecretKey: "secret456",
//			})
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			name := "tf-testAccCenInstanceGrantBasic"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			cen, err := cen.NewInstance(ctx, "cen", nil, pulumi.Provider(alicloud.Account2))
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("192.168.0.0/16"),
//			}, pulumi.Provider(alicloud.Account1))
//			if err != nil {
//				return err
//			}
//			fooInstanceGrant, err := cen.NewInstanceGrant(ctx, "fooInstanceGrant", &cen.InstanceGrantArgs{
//				CenId:           cen.ID(),
//				ChildInstanceId: vpc.ID(),
//				CenOwnerId:      pulumi.String("uid2"),
//			}, pulumi.Provider(alicloud.Account1))
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewInstanceAttachment(ctx, "fooInstanceAttachment", &cen.InstanceAttachmentArgs{
//				InstanceId:            cen.ID(),
//				ChildInstanceId:       vpc.ID(),
//				ChildInstanceType:     pulumi.String("VPC"),
//				ChildInstanceRegionId: pulumi.String("cn-qingdao"),
//				ChildInstanceOwnerId:  pulumi.Int("uid1"),
//			}, pulumi.Provider(alicloud.Account2), pulumi.DependsOn([]pulumi.Resource{
//				fooInstanceGrant,
//			}))
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
// CEN instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/instanceGrant:InstanceGrant example cen-abc123456:vpc-abc123456:uid123456
//
// ```
type InstanceGrant struct {
	pulumi.CustomResourceState

	// The ID of the CEN.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The owner UID of the  CEN which the child instance granted to.
	CenOwnerId pulumi.StringOutput `pulumi:"cenOwnerId"`
	// The ID of the child instance to grant.
	ChildInstanceId pulumi.StringOutput `pulumi:"childInstanceId"`
}

// NewInstanceGrant registers a new resource with the given unique name, arguments, and options.
func NewInstanceGrant(ctx *pulumi.Context,
	name string, args *InstanceGrantArgs, opts ...pulumi.ResourceOption) (*InstanceGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.CenOwnerId == nil {
		return nil, errors.New("invalid value for required argument 'CenOwnerId'")
	}
	if args.ChildInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'ChildInstanceId'")
	}
	var resource InstanceGrant
	err := ctx.RegisterResource("alicloud:cen/instanceGrant:InstanceGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGrant gets an existing InstanceGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGrantState, opts ...pulumi.ResourceOption) (*InstanceGrant, error) {
	var resource InstanceGrant
	err := ctx.ReadResource("alicloud:cen/instanceGrant:InstanceGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGrant resources.
type instanceGrantState struct {
	// The ID of the CEN.
	CenId *string `pulumi:"cenId"`
	// The owner UID of the  CEN which the child instance granted to.
	CenOwnerId *string `pulumi:"cenOwnerId"`
	// The ID of the child instance to grant.
	ChildInstanceId *string `pulumi:"childInstanceId"`
}

type InstanceGrantState struct {
	// The ID of the CEN.
	CenId pulumi.StringPtrInput
	// The owner UID of the  CEN which the child instance granted to.
	CenOwnerId pulumi.StringPtrInput
	// The ID of the child instance to grant.
	ChildInstanceId pulumi.StringPtrInput
}

func (InstanceGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGrantState)(nil)).Elem()
}

type instanceGrantArgs struct {
	// The ID of the CEN.
	CenId string `pulumi:"cenId"`
	// The owner UID of the  CEN which the child instance granted to.
	CenOwnerId string `pulumi:"cenOwnerId"`
	// The ID of the child instance to grant.
	ChildInstanceId string `pulumi:"childInstanceId"`
}

// The set of arguments for constructing a InstanceGrant resource.
type InstanceGrantArgs struct {
	// The ID of the CEN.
	CenId pulumi.StringInput
	// The owner UID of the  CEN which the child instance granted to.
	CenOwnerId pulumi.StringInput
	// The ID of the child instance to grant.
	ChildInstanceId pulumi.StringInput
}

func (InstanceGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGrantArgs)(nil)).Elem()
}

type InstanceGrantInput interface {
	pulumi.Input

	ToInstanceGrantOutput() InstanceGrantOutput
	ToInstanceGrantOutputWithContext(ctx context.Context) InstanceGrantOutput
}

func (*InstanceGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGrant)(nil)).Elem()
}

func (i *InstanceGrant) ToInstanceGrantOutput() InstanceGrantOutput {
	return i.ToInstanceGrantOutputWithContext(context.Background())
}

func (i *InstanceGrant) ToInstanceGrantOutputWithContext(ctx context.Context) InstanceGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGrantOutput)
}

// InstanceGrantArrayInput is an input type that accepts InstanceGrantArray and InstanceGrantArrayOutput values.
// You can construct a concrete instance of `InstanceGrantArrayInput` via:
//
//	InstanceGrantArray{ InstanceGrantArgs{...} }
type InstanceGrantArrayInput interface {
	pulumi.Input

	ToInstanceGrantArrayOutput() InstanceGrantArrayOutput
	ToInstanceGrantArrayOutputWithContext(context.Context) InstanceGrantArrayOutput
}

type InstanceGrantArray []InstanceGrantInput

func (InstanceGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGrant)(nil)).Elem()
}

func (i InstanceGrantArray) ToInstanceGrantArrayOutput() InstanceGrantArrayOutput {
	return i.ToInstanceGrantArrayOutputWithContext(context.Background())
}

func (i InstanceGrantArray) ToInstanceGrantArrayOutputWithContext(ctx context.Context) InstanceGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGrantArrayOutput)
}

// InstanceGrantMapInput is an input type that accepts InstanceGrantMap and InstanceGrantMapOutput values.
// You can construct a concrete instance of `InstanceGrantMapInput` via:
//
//	InstanceGrantMap{ "key": InstanceGrantArgs{...} }
type InstanceGrantMapInput interface {
	pulumi.Input

	ToInstanceGrantMapOutput() InstanceGrantMapOutput
	ToInstanceGrantMapOutputWithContext(context.Context) InstanceGrantMapOutput
}

type InstanceGrantMap map[string]InstanceGrantInput

func (InstanceGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGrant)(nil)).Elem()
}

func (i InstanceGrantMap) ToInstanceGrantMapOutput() InstanceGrantMapOutput {
	return i.ToInstanceGrantMapOutputWithContext(context.Background())
}

func (i InstanceGrantMap) ToInstanceGrantMapOutputWithContext(ctx context.Context) InstanceGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGrantMapOutput)
}

type InstanceGrantOutput struct{ *pulumi.OutputState }

func (InstanceGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGrant)(nil)).Elem()
}

func (o InstanceGrantOutput) ToInstanceGrantOutput() InstanceGrantOutput {
	return o
}

func (o InstanceGrantOutput) ToInstanceGrantOutputWithContext(ctx context.Context) InstanceGrantOutput {
	return o
}

// The ID of the CEN.
func (o InstanceGrantOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGrant) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The owner UID of the  CEN which the child instance granted to.
func (o InstanceGrantOutput) CenOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGrant) pulumi.StringOutput { return v.CenOwnerId }).(pulumi.StringOutput)
}

// The ID of the child instance to grant.
func (o InstanceGrantOutput) ChildInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGrant) pulumi.StringOutput { return v.ChildInstanceId }).(pulumi.StringOutput)
}

type InstanceGrantArrayOutput struct{ *pulumi.OutputState }

func (InstanceGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGrant)(nil)).Elem()
}

func (o InstanceGrantArrayOutput) ToInstanceGrantArrayOutput() InstanceGrantArrayOutput {
	return o
}

func (o InstanceGrantArrayOutput) ToInstanceGrantArrayOutputWithContext(ctx context.Context) InstanceGrantArrayOutput {
	return o
}

func (o InstanceGrantArrayOutput) Index(i pulumi.IntInput) InstanceGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceGrant {
		return vs[0].([]*InstanceGrant)[vs[1].(int)]
	}).(InstanceGrantOutput)
}

type InstanceGrantMapOutput struct{ *pulumi.OutputState }

func (InstanceGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGrant)(nil)).Elem()
}

func (o InstanceGrantMapOutput) ToInstanceGrantMapOutput() InstanceGrantMapOutput {
	return o
}

func (o InstanceGrantMapOutput) ToInstanceGrantMapOutputWithContext(ctx context.Context) InstanceGrantMapOutput {
	return o
}

func (o InstanceGrantMapOutput) MapIndex(k pulumi.StringInput) InstanceGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceGrant {
		return vs[0].(map[string]*InstanceGrant)[vs[1].(string)]
	}).(InstanceGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGrantInput)(nil)).Elem(), &InstanceGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGrantArrayInput)(nil)).Elem(), InstanceGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGrantMapInput)(nil)).Elem(), InstanceGrantMap{})
	pulumi.RegisterOutputType(InstanceGrantOutput{})
	pulumi.RegisterOutputType(InstanceGrantArrayOutput{})
	pulumi.RegisterOutputType(InstanceGrantMapOutput{})
}

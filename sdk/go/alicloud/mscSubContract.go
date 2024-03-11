// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Msc Sub Contact resource.
//
// > **NOTE:** Available since v1.132.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tfexample"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := alicloud.NewMscSubContract(ctx, "default", &alicloud.MscSubContractArgs{
//				ContactName: pulumi.String(name),
//				Position:    pulumi.String("CEO"),
//				Email:       pulumi.String("123@163.com"),
//				Mobile:      pulumi.String("15388888888"),
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
// Msc Sub Contact can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:index/mscSubContract:MscSubContract example <id>
// ```
type MscSubContract struct {
	pulumi.CustomResourceState

	// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
	ContactName pulumi.StringOutput `pulumi:"contactName"`
	// The User's Contact Email Address.
	Email pulumi.StringOutput `pulumi:"email"`
	// The User's Telephone.
	Mobile pulumi.StringOutput `pulumi:"mobile"`
	// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
	//
	// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
	Position pulumi.StringOutput `pulumi:"position"`
}

// NewMscSubContract registers a new resource with the given unique name, arguments, and options.
func NewMscSubContract(ctx *pulumi.Context,
	name string, args *MscSubContractArgs, opts ...pulumi.ResourceOption) (*MscSubContract, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactName == nil {
		return nil, errors.New("invalid value for required argument 'ContactName'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Mobile == nil {
		return nil, errors.New("invalid value for required argument 'Mobile'")
	}
	if args.Position == nil {
		return nil, errors.New("invalid value for required argument 'Position'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MscSubContract
	err := ctx.RegisterResource("alicloud:index/mscSubContract:MscSubContract", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMscSubContract gets an existing MscSubContract resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMscSubContract(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MscSubContractState, opts ...pulumi.ResourceOption) (*MscSubContract, error) {
	var resource MscSubContract
	err := ctx.ReadResource("alicloud:index/mscSubContract:MscSubContract", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MscSubContract resources.
type mscSubContractState struct {
	// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
	ContactName *string `pulumi:"contactName"`
	// The User's Contact Email Address.
	Email *string `pulumi:"email"`
	// The User's Telephone.
	Mobile *string `pulumi:"mobile"`
	// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
	//
	// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
	Position *string `pulumi:"position"`
}

type MscSubContractState struct {
	// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
	ContactName pulumi.StringPtrInput
	// The User's Contact Email Address.
	Email pulumi.StringPtrInput
	// The User's Telephone.
	Mobile pulumi.StringPtrInput
	// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
	//
	// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
	Position pulumi.StringPtrInput
}

func (MscSubContractState) ElementType() reflect.Type {
	return reflect.TypeOf((*mscSubContractState)(nil)).Elem()
}

type mscSubContractArgs struct {
	// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
	ContactName string `pulumi:"contactName"`
	// The User's Contact Email Address.
	Email string `pulumi:"email"`
	// The User's Telephone.
	Mobile string `pulumi:"mobile"`
	// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
	//
	// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
	Position string `pulumi:"position"`
}

// The set of arguments for constructing a MscSubContract resource.
type MscSubContractArgs struct {
	// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
	ContactName pulumi.StringInput
	// The User's Contact Email Address.
	Email pulumi.StringInput
	// The User's Telephone.
	Mobile pulumi.StringInput
	// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
	//
	// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
	Position pulumi.StringInput
}

func (MscSubContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mscSubContractArgs)(nil)).Elem()
}

type MscSubContractInput interface {
	pulumi.Input

	ToMscSubContractOutput() MscSubContractOutput
	ToMscSubContractOutputWithContext(ctx context.Context) MscSubContractOutput
}

func (*MscSubContract) ElementType() reflect.Type {
	return reflect.TypeOf((**MscSubContract)(nil)).Elem()
}

func (i *MscSubContract) ToMscSubContractOutput() MscSubContractOutput {
	return i.ToMscSubContractOutputWithContext(context.Background())
}

func (i *MscSubContract) ToMscSubContractOutputWithContext(ctx context.Context) MscSubContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MscSubContractOutput)
}

// MscSubContractArrayInput is an input type that accepts MscSubContractArray and MscSubContractArrayOutput values.
// You can construct a concrete instance of `MscSubContractArrayInput` via:
//
//	MscSubContractArray{ MscSubContractArgs{...} }
type MscSubContractArrayInput interface {
	pulumi.Input

	ToMscSubContractArrayOutput() MscSubContractArrayOutput
	ToMscSubContractArrayOutputWithContext(context.Context) MscSubContractArrayOutput
}

type MscSubContractArray []MscSubContractInput

func (MscSubContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MscSubContract)(nil)).Elem()
}

func (i MscSubContractArray) ToMscSubContractArrayOutput() MscSubContractArrayOutput {
	return i.ToMscSubContractArrayOutputWithContext(context.Background())
}

func (i MscSubContractArray) ToMscSubContractArrayOutputWithContext(ctx context.Context) MscSubContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MscSubContractArrayOutput)
}

// MscSubContractMapInput is an input type that accepts MscSubContractMap and MscSubContractMapOutput values.
// You can construct a concrete instance of `MscSubContractMapInput` via:
//
//	MscSubContractMap{ "key": MscSubContractArgs{...} }
type MscSubContractMapInput interface {
	pulumi.Input

	ToMscSubContractMapOutput() MscSubContractMapOutput
	ToMscSubContractMapOutputWithContext(context.Context) MscSubContractMapOutput
}

type MscSubContractMap map[string]MscSubContractInput

func (MscSubContractMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MscSubContract)(nil)).Elem()
}

func (i MscSubContractMap) ToMscSubContractMapOutput() MscSubContractMapOutput {
	return i.ToMscSubContractMapOutputWithContext(context.Background())
}

func (i MscSubContractMap) ToMscSubContractMapOutputWithContext(ctx context.Context) MscSubContractMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MscSubContractMapOutput)
}

type MscSubContractOutput struct{ *pulumi.OutputState }

func (MscSubContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MscSubContract)(nil)).Elem()
}

func (o MscSubContractOutput) ToMscSubContractOutput() MscSubContractOutput {
	return o
}

func (o MscSubContractOutput) ToMscSubContractOutputWithContext(ctx context.Context) MscSubContractOutput {
	return o
}

// The User's Contact Name. **Note:** The name must be 2 to 12 characters in length.
func (o MscSubContractOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v *MscSubContract) pulumi.StringOutput { return v.ContactName }).(pulumi.StringOutput)
}

// The User's Contact Email Address.
func (o MscSubContractOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *MscSubContract) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The User's Telephone.
func (o MscSubContractOutput) Mobile() pulumi.StringOutput {
	return o.ApplyT(func(v *MscSubContract) pulumi.StringOutput { return v.Mobile }).(pulumi.StringOutput)
}

// The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Other`.
//
// > **NOTE:** When the user creates a contact, the user should use `getMscSubContactVerificationMessage` to receive the verification message and confirm it.
func (o MscSubContractOutput) Position() pulumi.StringOutput {
	return o.ApplyT(func(v *MscSubContract) pulumi.StringOutput { return v.Position }).(pulumi.StringOutput)
}

type MscSubContractArrayOutput struct{ *pulumi.OutputState }

func (MscSubContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MscSubContract)(nil)).Elem()
}

func (o MscSubContractArrayOutput) ToMscSubContractArrayOutput() MscSubContractArrayOutput {
	return o
}

func (o MscSubContractArrayOutput) ToMscSubContractArrayOutputWithContext(ctx context.Context) MscSubContractArrayOutput {
	return o
}

func (o MscSubContractArrayOutput) Index(i pulumi.IntInput) MscSubContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MscSubContract {
		return vs[0].([]*MscSubContract)[vs[1].(int)]
	}).(MscSubContractOutput)
}

type MscSubContractMapOutput struct{ *pulumi.OutputState }

func (MscSubContractMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MscSubContract)(nil)).Elem()
}

func (o MscSubContractMapOutput) ToMscSubContractMapOutput() MscSubContractMapOutput {
	return o
}

func (o MscSubContractMapOutput) ToMscSubContractMapOutputWithContext(ctx context.Context) MscSubContractMapOutput {
	return o
}

func (o MscSubContractMapOutput) MapIndex(k pulumi.StringInput) MscSubContractOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MscSubContract {
		return vs[0].(map[string]*MscSubContract)[vs[1].(string)]
	}).(MscSubContractOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MscSubContractInput)(nil)).Elem(), &MscSubContract{})
	pulumi.RegisterInputType(reflect.TypeOf((*MscSubContractArrayInput)(nil)).Elem(), MscSubContractArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MscSubContractMapInput)(nil)).Elem(), MscSubContractMap{})
	pulumi.RegisterOutputType(MscSubContractOutput{})
	pulumi.RegisterOutputType(MscSubContractArrayOutput{})
	pulumi.RegisterOutputType(MscSubContractMapOutput{})
}

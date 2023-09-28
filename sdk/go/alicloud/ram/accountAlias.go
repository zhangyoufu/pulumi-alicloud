// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a RAM cloud account alias.
//
// > **NOTE:** Available since v1.0.0.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
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
//			_, err := ram.NewAccountAlias(ctx, "alias", &ram.AccountAliasArgs{
//				AccountAlias: pulumi.String(name),
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
// RAM account alias can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/accountAlias:AccountAlias example my-alias
//
// ```
type AccountAlias struct {
	pulumi.CustomResourceState

	// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
	AccountAlias pulumi.StringOutput `pulumi:"accountAlias"`
}

// NewAccountAlias registers a new resource with the given unique name, arguments, and options.
func NewAccountAlias(ctx *pulumi.Context,
	name string, args *AccountAliasArgs, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountAlias == nil {
		return nil, errors.New("invalid value for required argument 'AccountAlias'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountAlias
	err := ctx.RegisterResource("alicloud:ram/accountAlias:AccountAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountAlias gets an existing AccountAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountAliasState, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	var resource AccountAlias
	err := ctx.ReadResource("alicloud:ram/accountAlias:AccountAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountAlias resources.
type accountAliasState struct {
	// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
	AccountAlias *string `pulumi:"accountAlias"`
}

type AccountAliasState struct {
	// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
	AccountAlias pulumi.StringPtrInput
}

func (AccountAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasState)(nil)).Elem()
}

type accountAliasArgs struct {
	// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
	AccountAlias string `pulumi:"accountAlias"`
}

// The set of arguments for constructing a AccountAlias resource.
type AccountAliasArgs struct {
	// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
	AccountAlias pulumi.StringInput
}

func (AccountAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasArgs)(nil)).Elem()
}

type AccountAliasInput interface {
	pulumi.Input

	ToAccountAliasOutput() AccountAliasOutput
	ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput
}

func (*AccountAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountAlias)(nil)).Elem()
}

func (i *AccountAlias) ToAccountAliasOutput() AccountAliasOutput {
	return i.ToAccountAliasOutputWithContext(context.Background())
}

func (i *AccountAlias) ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasOutput)
}

func (i *AccountAlias) ToOutput(ctx context.Context) pulumix.Output[*AccountAlias] {
	return pulumix.Output[*AccountAlias]{
		OutputState: i.ToAccountAliasOutputWithContext(ctx).OutputState,
	}
}

// AccountAliasArrayInput is an input type that accepts AccountAliasArray and AccountAliasArrayOutput values.
// You can construct a concrete instance of `AccountAliasArrayInput` via:
//
//	AccountAliasArray{ AccountAliasArgs{...} }
type AccountAliasArrayInput interface {
	pulumi.Input

	ToAccountAliasArrayOutput() AccountAliasArrayOutput
	ToAccountAliasArrayOutputWithContext(context.Context) AccountAliasArrayOutput
}

type AccountAliasArray []AccountAliasInput

func (AccountAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountAlias)(nil)).Elem()
}

func (i AccountAliasArray) ToAccountAliasArrayOutput() AccountAliasArrayOutput {
	return i.ToAccountAliasArrayOutputWithContext(context.Background())
}

func (i AccountAliasArray) ToAccountAliasArrayOutputWithContext(ctx context.Context) AccountAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasArrayOutput)
}

func (i AccountAliasArray) ToOutput(ctx context.Context) pulumix.Output[[]*AccountAlias] {
	return pulumix.Output[[]*AccountAlias]{
		OutputState: i.ToAccountAliasArrayOutputWithContext(ctx).OutputState,
	}
}

// AccountAliasMapInput is an input type that accepts AccountAliasMap and AccountAliasMapOutput values.
// You can construct a concrete instance of `AccountAliasMapInput` via:
//
//	AccountAliasMap{ "key": AccountAliasArgs{...} }
type AccountAliasMapInput interface {
	pulumi.Input

	ToAccountAliasMapOutput() AccountAliasMapOutput
	ToAccountAliasMapOutputWithContext(context.Context) AccountAliasMapOutput
}

type AccountAliasMap map[string]AccountAliasInput

func (AccountAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountAlias)(nil)).Elem()
}

func (i AccountAliasMap) ToAccountAliasMapOutput() AccountAliasMapOutput {
	return i.ToAccountAliasMapOutputWithContext(context.Background())
}

func (i AccountAliasMap) ToAccountAliasMapOutputWithContext(ctx context.Context) AccountAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountAliasMapOutput)
}

func (i AccountAliasMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountAlias] {
	return pulumix.Output[map[string]*AccountAlias]{
		OutputState: i.ToAccountAliasMapOutputWithContext(ctx).OutputState,
	}
}

type AccountAliasOutput struct{ *pulumi.OutputState }

func (AccountAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountAlias)(nil)).Elem()
}

func (o AccountAliasOutput) ToAccountAliasOutput() AccountAliasOutput {
	return o
}

func (o AccountAliasOutput) ToAccountAliasOutputWithContext(ctx context.Context) AccountAliasOutput {
	return o
}

func (o AccountAliasOutput) ToOutput(ctx context.Context) pulumix.Output[*AccountAlias] {
	return pulumix.Output[*AccountAlias]{
		OutputState: o.OutputState,
	}
}

// Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
func (o AccountAliasOutput) AccountAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountAlias) pulumi.StringOutput { return v.AccountAlias }).(pulumi.StringOutput)
}

type AccountAliasArrayOutput struct{ *pulumi.OutputState }

func (AccountAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountAlias)(nil)).Elem()
}

func (o AccountAliasArrayOutput) ToAccountAliasArrayOutput() AccountAliasArrayOutput {
	return o
}

func (o AccountAliasArrayOutput) ToAccountAliasArrayOutputWithContext(ctx context.Context) AccountAliasArrayOutput {
	return o
}

func (o AccountAliasArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AccountAlias] {
	return pulumix.Output[[]*AccountAlias]{
		OutputState: o.OutputState,
	}
}

func (o AccountAliasArrayOutput) Index(i pulumi.IntInput) AccountAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountAlias {
		return vs[0].([]*AccountAlias)[vs[1].(int)]
	}).(AccountAliasOutput)
}

type AccountAliasMapOutput struct{ *pulumi.OutputState }

func (AccountAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountAlias)(nil)).Elem()
}

func (o AccountAliasMapOutput) ToAccountAliasMapOutput() AccountAliasMapOutput {
	return o
}

func (o AccountAliasMapOutput) ToAccountAliasMapOutputWithContext(ctx context.Context) AccountAliasMapOutput {
	return o
}

func (o AccountAliasMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountAlias] {
	return pulumix.Output[map[string]*AccountAlias]{
		OutputState: o.OutputState,
	}
}

func (o AccountAliasMapOutput) MapIndex(k pulumi.StringInput) AccountAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountAlias {
		return vs[0].(map[string]*AccountAlias)[vs[1].(string)]
	}).(AccountAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasInput)(nil)).Elem(), &AccountAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasArrayInput)(nil)).Elem(), AccountAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountAliasMapInput)(nil)).Elem(), AccountAliasMap{})
	pulumi.RegisterOutputType(AccountAliasOutput{})
	pulumi.RegisterOutputType(AccountAliasArrayOutput{})
	pulumi.RegisterOutputType(AccountAliasMapOutput{})
}

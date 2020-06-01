// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GetAliasesAlias struct {
	// The unique identifier of the alias.
	AliasName string `pulumi:"aliasName"`
	// ID of the alias. The value is same as KMS alias_name.
	Id string `pulumi:"id"`
	// ID of the key.
	KeyId string `pulumi:"keyId"`
}

// GetAliasesAliasInput is an input type that accepts GetAliasesAliasArgs and GetAliasesAliasOutput values.
// You can construct a concrete instance of `GetAliasesAliasInput` via:
//
// 		 GetAliasesAliasArgs{...}
//
type GetAliasesAliasInput interface {
	pulumi.Input

	ToGetAliasesAliasOutput() GetAliasesAliasOutput
	ToGetAliasesAliasOutputWithContext(context.Context) GetAliasesAliasOutput
}

type GetAliasesAliasArgs struct {
	// The unique identifier of the alias.
	AliasName pulumi.StringInput `pulumi:"aliasName"`
	// ID of the alias. The value is same as KMS alias_name.
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the key.
	KeyId pulumi.StringInput `pulumi:"keyId"`
}

func (GetAliasesAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAliasesAlias)(nil)).Elem()
}

func (i GetAliasesAliasArgs) ToGetAliasesAliasOutput() GetAliasesAliasOutput {
	return i.ToGetAliasesAliasOutputWithContext(context.Background())
}

func (i GetAliasesAliasArgs) ToGetAliasesAliasOutputWithContext(ctx context.Context) GetAliasesAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAliasesAliasOutput)
}

// GetAliasesAliasArrayInput is an input type that accepts GetAliasesAliasArray and GetAliasesAliasArrayOutput values.
// You can construct a concrete instance of `GetAliasesAliasArrayInput` via:
//
// 		 GetAliasesAliasArray{ GetAliasesAliasArgs{...} }
//
type GetAliasesAliasArrayInput interface {
	pulumi.Input

	ToGetAliasesAliasArrayOutput() GetAliasesAliasArrayOutput
	ToGetAliasesAliasArrayOutputWithContext(context.Context) GetAliasesAliasArrayOutput
}

type GetAliasesAliasArray []GetAliasesAliasInput

func (GetAliasesAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAliasesAlias)(nil)).Elem()
}

func (i GetAliasesAliasArray) ToGetAliasesAliasArrayOutput() GetAliasesAliasArrayOutput {
	return i.ToGetAliasesAliasArrayOutputWithContext(context.Background())
}

func (i GetAliasesAliasArray) ToGetAliasesAliasArrayOutputWithContext(ctx context.Context) GetAliasesAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAliasesAliasArrayOutput)
}

type GetAliasesAliasOutput struct{ *pulumi.OutputState }

func (GetAliasesAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAliasesAlias)(nil)).Elem()
}

func (o GetAliasesAliasOutput) ToGetAliasesAliasOutput() GetAliasesAliasOutput {
	return o
}

func (o GetAliasesAliasOutput) ToGetAliasesAliasOutputWithContext(ctx context.Context) GetAliasesAliasOutput {
	return o
}

// The unique identifier of the alias.
func (o GetAliasesAliasOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAliasesAlias) string { return v.AliasName }).(pulumi.StringOutput)
}

// ID of the alias. The value is same as KMS alias_name.
func (o GetAliasesAliasOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAliasesAlias) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the key.
func (o GetAliasesAliasOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAliasesAlias) string { return v.KeyId }).(pulumi.StringOutput)
}

type GetAliasesAliasArrayOutput struct{ *pulumi.OutputState }

func (GetAliasesAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAliasesAlias)(nil)).Elem()
}

func (o GetAliasesAliasArrayOutput) ToGetAliasesAliasArrayOutput() GetAliasesAliasArrayOutput {
	return o
}

func (o GetAliasesAliasArrayOutput) ToGetAliasesAliasArrayOutputWithContext(ctx context.Context) GetAliasesAliasArrayOutput {
	return o
}

func (o GetAliasesAliasArrayOutput) Index(i pulumi.IntInput) GetAliasesAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAliasesAlias {
		return vs[0].([]GetAliasesAlias)[vs[1].(int)]
	}).(GetAliasesAliasOutput)
}

type GetKeyVersionsVersion struct {
	// Date and time when the key version was created (UTC time).
	CreationDate string `pulumi:"creationDate"`
	// ID of the KMS KeyVersion resource.
	Id string `pulumi:"id"`
	// The id of kms key.
	KeyId string `pulumi:"keyId"`
	// ID of the key version.
	KeyVersionId string `pulumi:"keyVersionId"`
}

// GetKeyVersionsVersionInput is an input type that accepts GetKeyVersionsVersionArgs and GetKeyVersionsVersionOutput values.
// You can construct a concrete instance of `GetKeyVersionsVersionInput` via:
//
// 		 GetKeyVersionsVersionArgs{...}
//
type GetKeyVersionsVersionInput interface {
	pulumi.Input

	ToGetKeyVersionsVersionOutput() GetKeyVersionsVersionOutput
	ToGetKeyVersionsVersionOutputWithContext(context.Context) GetKeyVersionsVersionOutput
}

type GetKeyVersionsVersionArgs struct {
	// Date and time when the key version was created (UTC time).
	CreationDate pulumi.StringInput `pulumi:"creationDate"`
	// ID of the KMS KeyVersion resource.
	Id pulumi.StringInput `pulumi:"id"`
	// The id of kms key.
	KeyId pulumi.StringInput `pulumi:"keyId"`
	// ID of the key version.
	KeyVersionId pulumi.StringInput `pulumi:"keyVersionId"`
}

func (GetKeyVersionsVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyVersionsVersion)(nil)).Elem()
}

func (i GetKeyVersionsVersionArgs) ToGetKeyVersionsVersionOutput() GetKeyVersionsVersionOutput {
	return i.ToGetKeyVersionsVersionOutputWithContext(context.Background())
}

func (i GetKeyVersionsVersionArgs) ToGetKeyVersionsVersionOutputWithContext(ctx context.Context) GetKeyVersionsVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyVersionsVersionOutput)
}

// GetKeyVersionsVersionArrayInput is an input type that accepts GetKeyVersionsVersionArray and GetKeyVersionsVersionArrayOutput values.
// You can construct a concrete instance of `GetKeyVersionsVersionArrayInput` via:
//
// 		 GetKeyVersionsVersionArray{ GetKeyVersionsVersionArgs{...} }
//
type GetKeyVersionsVersionArrayInput interface {
	pulumi.Input

	ToGetKeyVersionsVersionArrayOutput() GetKeyVersionsVersionArrayOutput
	ToGetKeyVersionsVersionArrayOutputWithContext(context.Context) GetKeyVersionsVersionArrayOutput
}

type GetKeyVersionsVersionArray []GetKeyVersionsVersionInput

func (GetKeyVersionsVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyVersionsVersion)(nil)).Elem()
}

func (i GetKeyVersionsVersionArray) ToGetKeyVersionsVersionArrayOutput() GetKeyVersionsVersionArrayOutput {
	return i.ToGetKeyVersionsVersionArrayOutputWithContext(context.Background())
}

func (i GetKeyVersionsVersionArray) ToGetKeyVersionsVersionArrayOutputWithContext(ctx context.Context) GetKeyVersionsVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyVersionsVersionArrayOutput)
}

type GetKeyVersionsVersionOutput struct{ *pulumi.OutputState }

func (GetKeyVersionsVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyVersionsVersion)(nil)).Elem()
}

func (o GetKeyVersionsVersionOutput) ToGetKeyVersionsVersionOutput() GetKeyVersionsVersionOutput {
	return o
}

func (o GetKeyVersionsVersionOutput) ToGetKeyVersionsVersionOutputWithContext(ctx context.Context) GetKeyVersionsVersionOutput {
	return o
}

// Date and time when the key version was created (UTC time).
func (o GetKeyVersionsVersionOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyVersionsVersion) string { return v.CreationDate }).(pulumi.StringOutput)
}

// ID of the KMS KeyVersion resource.
func (o GetKeyVersionsVersionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyVersionsVersion) string { return v.Id }).(pulumi.StringOutput)
}

// The id of kms key.
func (o GetKeyVersionsVersionOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyVersionsVersion) string { return v.KeyId }).(pulumi.StringOutput)
}

// ID of the key version.
func (o GetKeyVersionsVersionOutput) KeyVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyVersionsVersion) string { return v.KeyVersionId }).(pulumi.StringOutput)
}

type GetKeyVersionsVersionArrayOutput struct{ *pulumi.OutputState }

func (GetKeyVersionsVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyVersionsVersion)(nil)).Elem()
}

func (o GetKeyVersionsVersionArrayOutput) ToGetKeyVersionsVersionArrayOutput() GetKeyVersionsVersionArrayOutput {
	return o
}

func (o GetKeyVersionsVersionArrayOutput) ToGetKeyVersionsVersionArrayOutputWithContext(ctx context.Context) GetKeyVersionsVersionArrayOutput {
	return o
}

func (o GetKeyVersionsVersionArrayOutput) Index(i pulumi.IntInput) GetKeyVersionsVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetKeyVersionsVersion {
		return vs[0].([]GetKeyVersionsVersion)[vs[1].(int)]
	}).(GetKeyVersionsVersionOutput)
}

type GetKeysKey struct {
	// The Alibaba Cloud Resource Name (ARN) of the key.
	Arn string `pulumi:"arn"`
	// Creation date of key.
	CreationDate string `pulumi:"creationDate"`
	// The owner of the key.
	Creator string `pulumi:"creator"`
	// Deletion date of key.
	DeleteDate string `pulumi:"deleteDate"`
	// Description of the key.
	Description string `pulumi:"description"`
	// ID of the key.
	Id string `pulumi:"id"`
	// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
	Status string `pulumi:"status"`
}

// GetKeysKeyInput is an input type that accepts GetKeysKeyArgs and GetKeysKeyOutput values.
// You can construct a concrete instance of `GetKeysKeyInput` via:
//
// 		 GetKeysKeyArgs{...}
//
type GetKeysKeyInput interface {
	pulumi.Input

	ToGetKeysKeyOutput() GetKeysKeyOutput
	ToGetKeysKeyOutputWithContext(context.Context) GetKeysKeyOutput
}

type GetKeysKeyArgs struct {
	// The Alibaba Cloud Resource Name (ARN) of the key.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Creation date of key.
	CreationDate pulumi.StringInput `pulumi:"creationDate"`
	// The owner of the key.
	Creator pulumi.StringInput `pulumi:"creator"`
	// Deletion date of key.
	DeleteDate pulumi.StringInput `pulumi:"deleteDate"`
	// Description of the key.
	Description pulumi.StringInput `pulumi:"description"`
	// ID of the key.
	Id pulumi.StringInput `pulumi:"id"`
	// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetKeysKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeysKey)(nil)).Elem()
}

func (i GetKeysKeyArgs) ToGetKeysKeyOutput() GetKeysKeyOutput {
	return i.ToGetKeysKeyOutputWithContext(context.Background())
}

func (i GetKeysKeyArgs) ToGetKeysKeyOutputWithContext(ctx context.Context) GetKeysKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeysKeyOutput)
}

// GetKeysKeyArrayInput is an input type that accepts GetKeysKeyArray and GetKeysKeyArrayOutput values.
// You can construct a concrete instance of `GetKeysKeyArrayInput` via:
//
// 		 GetKeysKeyArray{ GetKeysKeyArgs{...} }
//
type GetKeysKeyArrayInput interface {
	pulumi.Input

	ToGetKeysKeyArrayOutput() GetKeysKeyArrayOutput
	ToGetKeysKeyArrayOutputWithContext(context.Context) GetKeysKeyArrayOutput
}

type GetKeysKeyArray []GetKeysKeyInput

func (GetKeysKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeysKey)(nil)).Elem()
}

func (i GetKeysKeyArray) ToGetKeysKeyArrayOutput() GetKeysKeyArrayOutput {
	return i.ToGetKeysKeyArrayOutputWithContext(context.Background())
}

func (i GetKeysKeyArray) ToGetKeysKeyArrayOutputWithContext(ctx context.Context) GetKeysKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeysKeyArrayOutput)
}

type GetKeysKeyOutput struct{ *pulumi.OutputState }

func (GetKeysKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeysKey)(nil)).Elem()
}

func (o GetKeysKeyOutput) ToGetKeysKeyOutput() GetKeysKeyOutput {
	return o
}

func (o GetKeysKeyOutput) ToGetKeysKeyOutputWithContext(ctx context.Context) GetKeysKeyOutput {
	return o
}

// The Alibaba Cloud Resource Name (ARN) of the key.
func (o GetKeysKeyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.Arn }).(pulumi.StringOutput)
}

// Creation date of key.
func (o GetKeysKeyOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The owner of the key.
func (o GetKeysKeyOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.Creator }).(pulumi.StringOutput)
}

// Deletion date of key.
func (o GetKeysKeyOutput) DeleteDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.DeleteDate }).(pulumi.StringOutput)
}

// Description of the key.
func (o GetKeysKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.Description }).(pulumi.StringOutput)
}

// ID of the key.
func (o GetKeysKeyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.Id }).(pulumi.StringOutput)
}

// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
func (o GetKeysKeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysKey) string { return v.Status }).(pulumi.StringOutput)
}

type GetKeysKeyArrayOutput struct{ *pulumi.OutputState }

func (GetKeysKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeysKey)(nil)).Elem()
}

func (o GetKeysKeyArrayOutput) ToGetKeysKeyArrayOutput() GetKeysKeyArrayOutput {
	return o
}

func (o GetKeysKeyArrayOutput) ToGetKeysKeyArrayOutputWithContext(ctx context.Context) GetKeysKeyArrayOutput {
	return o
}

func (o GetKeysKeyArrayOutput) Index(i pulumi.IntInput) GetKeysKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetKeysKey {
		return vs[0].([]GetKeysKey)[vs[1].(int)]
	}).(GetKeysKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAliasesAliasOutput{})
	pulumi.RegisterOutputType(GetAliasesAliasArrayOutput{})
	pulumi.RegisterOutputType(GetKeyVersionsVersionOutput{})
	pulumi.RegisterOutputType(GetKeyVersionsVersionArrayOutput{})
	pulumi.RegisterOutputType(GetKeysKeyOutput{})
	pulumi.RegisterOutputType(GetKeysKeyArrayOutput{})
}

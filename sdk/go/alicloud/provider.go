// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the alicloud package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.EcsRoleName == nil {
		args.EcsRoleName = pulumi.StringPtr(getEnvOrDefault("", nil, "ALICLOUD_ECS_ROLE_NAME").(string))
	}
	if args.Profile == nil {
		args.Profile = pulumi.StringPtr(getEnvOrDefault("", nil, "ALICLOUD_PROFILE").(string))
	}
	if args.Region == nil {
		args.Region = pulumi.StringPtr(getEnvOrDefault("", nil, "ALICLOUD_REGION").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:alicloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	AccessKey *string `pulumi:"accessKey"`
	// The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
	// Alibaba Cloud console.
	AccountId  *string             `pulumi:"accountId"`
	AssumeRole *ProviderAssumeRole `pulumi:"assumeRole"`
	// Use this to mark a terraform configuration file source.
	ConfigurationSource *string `pulumi:"configurationSource"`
	// The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
	// of the Alibaba Cloud console.
	EcsRoleName *string            `pulumi:"ecsRoleName"`
	Endpoints   []ProviderEndpoint `pulumi:"endpoints"`
	// Deprecated: Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.
	Fc *string `pulumi:"fc"`
	// Deprecated: Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.
	LogEndpoint *string `pulumi:"logEndpoint"`
	// Deprecated: Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.
	MnsEndpoint *string `pulumi:"mnsEndpoint"`
	// Deprecated: Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.
	OtsInstanceName *string `pulumi:"otsInstanceName"`
	// The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
	Profile  *string `pulumi:"profile"`
	Protocol *string `pulumi:"protocol"`
	// The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
	Region *string `pulumi:"region"`
	// The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	SecretKey *string `pulumi:"secretKey"`
	// security token. A security token is only required if you are using Security Token Service.
	SecurityToken *string `pulumi:"securityToken"`
	// The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
	SharedCredentialsFile *string `pulumi:"sharedCredentialsFile"`
	// Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
	// that are not public (yet).
	SkipRegionValidation *bool `pulumi:"skipRegionValidation"`
	// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	SourceIp *string `pulumi:"sourceIp"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	AccessKey pulumi.StringPtrInput
	// The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
	// Alibaba Cloud console.
	AccountId  pulumi.StringPtrInput
	AssumeRole ProviderAssumeRolePtrInput
	// Use this to mark a terraform configuration file source.
	ConfigurationSource pulumi.StringPtrInput
	// The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
	// of the Alibaba Cloud console.
	EcsRoleName pulumi.StringPtrInput
	Endpoints   ProviderEndpointArrayInput
	// Deprecated: Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.
	Fc pulumi.StringPtrInput
	// Deprecated: Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.
	LogEndpoint pulumi.StringPtrInput
	// Deprecated: Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.
	MnsEndpoint pulumi.StringPtrInput
	// Deprecated: Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.
	OtsInstanceName pulumi.StringPtrInput
	// The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
	Profile  pulumi.StringPtrInput
	Protocol pulumi.StringPtrInput
	// The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
	Region pulumi.StringPtrInput
	// The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	SecretKey pulumi.StringPtrInput
	// security token. A security token is only required if you are using Security Token Service.
	SecurityToken pulumi.StringPtrInput
	// The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
	SharedCredentialsFile pulumi.StringPtrInput
	// Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
	// that are not public (yet).
	SkipRegionValidation pulumi.BoolPtrInput
	// The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
	// console.
	SourceIp pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyT(func(v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct {
	*pulumi.OutputState
}

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}

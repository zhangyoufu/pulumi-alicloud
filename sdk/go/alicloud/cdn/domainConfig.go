// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cdn

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CDN Accelerated Domain resource.
//
// For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/zh/doc-detail/90915.htm)
//
// > **NOTE:** Available in v1.34.0+.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cdn_domain_config.html.markdown.
type DomainConfig struct {
	pulumi.CustomResourceState

	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The args of the domain config.
	FunctionArgs DomainConfigFunctionArgArrayOutput `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
}

// NewDomainConfig registers a new resource with the given unique name, arguments, and options.
func NewDomainConfig(ctx *pulumi.Context,
	name string, args *DomainConfigArgs, opts ...pulumi.ResourceOption) (*DomainConfig, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil || args.FunctionArgs == nil {
		return nil, errors.New("missing required argument 'FunctionArgs'")
	}
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil {
		args = &DomainConfigArgs{}
	}
	var resource DomainConfig
	err := ctx.RegisterResource("alicloud:cdn/domainConfig:DomainConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainConfig gets an existing DomainConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainConfigState, opts ...pulumi.ResourceOption) (*DomainConfig, error) {
	var resource DomainConfig
	err := ctx.ReadResource("alicloud:cdn/domainConfig:DomainConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainConfig resources.
type domainConfigState struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName *string `pulumi:"domainName"`
	// The args of the domain config.
	FunctionArgs []DomainConfigFunctionArg `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName *string `pulumi:"functionName"`
}

type DomainConfigState struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringPtrInput
	// The args of the domain config.
	FunctionArgs DomainConfigFunctionArgArrayInput
	// The name of the domain config.
	FunctionName pulumi.StringPtrInput
}

func (DomainConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigState)(nil)).Elem()
}

type domainConfigArgs struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName string `pulumi:"domainName"`
	// The args of the domain config.
	FunctionArgs []DomainConfigFunctionArg `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName string `pulumi:"functionName"`
}

// The set of arguments for constructing a DomainConfig resource.
type DomainConfigArgs struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringInput
	// The args of the domain config.
	FunctionArgs DomainConfigFunctionArgArrayInput
	// The name of the domain config.
	FunctionName pulumi.StringInput
}

func (DomainConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigArgs)(nil)).Elem()
}

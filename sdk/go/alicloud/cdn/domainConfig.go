// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cdn Domain Config resource.
//
// For information about Cdn Domain Config and how to use it, see [What is Domain Config](https://www.alibabacloud.com/help/en/doc-detail/90915.htm)
//
// > **NOTE:** Available since v1.34.0.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cdn"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewRandomInteger(ctx, "default", &random.RandomIntegerArgs{
//				Min: pulumi.Int(10000),
//				Max: pulumi.Int(99999),
//			})
//			if err != nil {
//				return err
//			}
//			// Create a new Domain config.
//			domain, err := cdn.NewDomainNew(ctx, "domain", &cdn.DomainNewArgs{
//				DomainName: _default.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("mycdndomain-%v.alicloud-provider.cn", result), nil
//				}).(pulumi.StringOutput),
//				CdnType: pulumi.String("web"),
//				Scope:   pulumi.String("overseas"),
//				Sources: cdn.DomainNewSourceArray{
//					&cdn.DomainNewSourceArgs{
//						Content:  pulumi.String("1.1.1.1"),
//						Type:     pulumi.String("ipaddr"),
//						Priority: pulumi.Int(20),
//						Port:     pulumi.Int(80),
//						Weight:   pulumi.Int(15),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cdn.NewDomainConfig(ctx, "config", &cdn.DomainConfigArgs{
//				DomainName:   domain.DomainName,
//				FunctionName: pulumi.String("ip_allow_list_set"),
//				FunctionArgs: cdn.DomainConfigFunctionArgArray{
//					&cdn.DomainConfigFunctionArgArgs{
//						ArgName:  pulumi.String("ip_list"),
//						ArgValue: pulumi.String("110.110.110.110"),
//					},
//				},
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
// CDN domain config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cdn/domainConfig:DomainConfig example <domain_name>:<function_name>:<config_id>
// ```
//
// ```sh
// $ pulumi import alicloud:cdn/domainConfig:DomainConfig example <domain_name>:<function_name>
// ```
type DomainConfig struct {
	pulumi.CustomResourceState

	// (Available since v1.132.0) The ID of the domain config function.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The args of the domain config. See `functionArgs` below.
	FunctionArgs DomainConfigFunctionArgArrayOutput `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// (Available since v1.132.0) The Status of the function.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewDomainConfig registers a new resource with the given unique name, arguments, and options.
func NewDomainConfig(ctx *pulumi.Context,
	name string, args *DomainConfigArgs, opts ...pulumi.ResourceOption) (*DomainConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.FunctionArgs == nil {
		return nil, errors.New("invalid value for required argument 'FunctionArgs'")
	}
	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// (Available since v1.132.0) The ID of the domain config function.
	ConfigId *string `pulumi:"configId"`
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName *string `pulumi:"domainName"`
	// The args of the domain config. See `functionArgs` below.
	FunctionArgs []DomainConfigFunctionArg `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName *string `pulumi:"functionName"`
	// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
	ParentId *string `pulumi:"parentId"`
	// (Available since v1.132.0) The Status of the function.
	Status *string `pulumi:"status"`
}

type DomainConfigState struct {
	// (Available since v1.132.0) The ID of the domain config function.
	ConfigId pulumi.StringPtrInput
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringPtrInput
	// The args of the domain config. See `functionArgs` below.
	FunctionArgs DomainConfigFunctionArgArrayInput
	// The name of the domain config.
	FunctionName pulumi.StringPtrInput
	// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
	ParentId pulumi.StringPtrInput
	// (Available since v1.132.0) The Status of the function.
	Status pulumi.StringPtrInput
}

func (DomainConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigState)(nil)).Elem()
}

type domainConfigArgs struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName string `pulumi:"domainName"`
	// The args of the domain config. See `functionArgs` below.
	FunctionArgs []DomainConfigFunctionArg `pulumi:"functionArgs"`
	// The name of the domain config.
	FunctionName string `pulumi:"functionName"`
	// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
	ParentId *string `pulumi:"parentId"`
}

// The set of arguments for constructing a DomainConfig resource.
type DomainConfigArgs struct {
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringInput
	// The args of the domain config. See `functionArgs` below.
	FunctionArgs DomainConfigFunctionArgArrayInput
	// The name of the domain config.
	FunctionName pulumi.StringInput
	// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
	ParentId pulumi.StringPtrInput
}

func (DomainConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainConfigArgs)(nil)).Elem()
}

type DomainConfigInput interface {
	pulumi.Input

	ToDomainConfigOutput() DomainConfigOutput
	ToDomainConfigOutputWithContext(ctx context.Context) DomainConfigOutput
}

func (*DomainConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainConfig)(nil)).Elem()
}

func (i *DomainConfig) ToDomainConfigOutput() DomainConfigOutput {
	return i.ToDomainConfigOutputWithContext(context.Background())
}

func (i *DomainConfig) ToDomainConfigOutputWithContext(ctx context.Context) DomainConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigOutput)
}

// DomainConfigArrayInput is an input type that accepts DomainConfigArray and DomainConfigArrayOutput values.
// You can construct a concrete instance of `DomainConfigArrayInput` via:
//
//	DomainConfigArray{ DomainConfigArgs{...} }
type DomainConfigArrayInput interface {
	pulumi.Input

	ToDomainConfigArrayOutput() DomainConfigArrayOutput
	ToDomainConfigArrayOutputWithContext(context.Context) DomainConfigArrayOutput
}

type DomainConfigArray []DomainConfigInput

func (DomainConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainConfig)(nil)).Elem()
}

func (i DomainConfigArray) ToDomainConfigArrayOutput() DomainConfigArrayOutput {
	return i.ToDomainConfigArrayOutputWithContext(context.Background())
}

func (i DomainConfigArray) ToDomainConfigArrayOutputWithContext(ctx context.Context) DomainConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigArrayOutput)
}

// DomainConfigMapInput is an input type that accepts DomainConfigMap and DomainConfigMapOutput values.
// You can construct a concrete instance of `DomainConfigMapInput` via:
//
//	DomainConfigMap{ "key": DomainConfigArgs{...} }
type DomainConfigMapInput interface {
	pulumi.Input

	ToDomainConfigMapOutput() DomainConfigMapOutput
	ToDomainConfigMapOutputWithContext(context.Context) DomainConfigMapOutput
}

type DomainConfigMap map[string]DomainConfigInput

func (DomainConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainConfig)(nil)).Elem()
}

func (i DomainConfigMap) ToDomainConfigMapOutput() DomainConfigMapOutput {
	return i.ToDomainConfigMapOutputWithContext(context.Background())
}

func (i DomainConfigMap) ToDomainConfigMapOutputWithContext(ctx context.Context) DomainConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigMapOutput)
}

type DomainConfigOutput struct{ *pulumi.OutputState }

func (DomainConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainConfig)(nil)).Elem()
}

func (o DomainConfigOutput) ToDomainConfigOutput() DomainConfigOutput {
	return o
}

func (o DomainConfigOutput) ToDomainConfigOutputWithContext(ctx context.Context) DomainConfigOutput {
	return o
}

// (Available since v1.132.0) The ID of the domain config function.
func (o DomainConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfig) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
func (o DomainConfigOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfig) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The args of the domain config. See `functionArgs` below.
func (o DomainConfigOutput) FunctionArgs() DomainConfigFunctionArgArrayOutput {
	return o.ApplyT(func(v *DomainConfig) DomainConfigFunctionArgArrayOutput { return v.FunctionArgs }).(DomainConfigFunctionArgArrayOutput)
}

// The name of the domain config.
func (o DomainConfigOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfig) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

// By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
func (o DomainConfigOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfig) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// (Available since v1.132.0) The Status of the function.
func (o DomainConfigOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainConfig) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type DomainConfigArrayOutput struct{ *pulumi.OutputState }

func (DomainConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainConfig)(nil)).Elem()
}

func (o DomainConfigArrayOutput) ToDomainConfigArrayOutput() DomainConfigArrayOutput {
	return o
}

func (o DomainConfigArrayOutput) ToDomainConfigArrayOutputWithContext(ctx context.Context) DomainConfigArrayOutput {
	return o
}

func (o DomainConfigArrayOutput) Index(i pulumi.IntInput) DomainConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainConfig {
		return vs[0].([]*DomainConfig)[vs[1].(int)]
	}).(DomainConfigOutput)
}

type DomainConfigMapOutput struct{ *pulumi.OutputState }

func (DomainConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainConfig)(nil)).Elem()
}

func (o DomainConfigMapOutput) ToDomainConfigMapOutput() DomainConfigMapOutput {
	return o
}

func (o DomainConfigMapOutput) ToDomainConfigMapOutputWithContext(ctx context.Context) DomainConfigMapOutput {
	return o
}

func (o DomainConfigMapOutput) MapIndex(k pulumi.StringInput) DomainConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainConfig {
		return vs[0].(map[string]*DomainConfig)[vs[1].(string)]
	}).(DomainConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigInput)(nil)).Elem(), &DomainConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigArrayInput)(nil)).Elem(), DomainConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigMapInput)(nil)).Elem(), DomainConfigMap{})
	pulumi.RegisterOutputType(DomainConfigOutput{})
	pulumi.RegisterOutputType(DomainConfigArrayOutput{})
	pulumi.RegisterOutputType(DomainConfigMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CR Chain resource.
//
// For information about CR Chain and how to use it, see [What is Chain](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchain).
//
// > **NOTE:** Available since v1.161.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := cr.NewRegistryEnterpriseInstance(ctx, "default", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewPeriod:   pulumi.Int(0),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRegistryEnterpriseNamespace, err := cs.NewRegistryEnterpriseNamespace(ctx, "default", &cs.RegistryEnterpriseNamespaceArgs{
//				InstanceId:        _default.ID(),
//				Name:              pulumi.String(name),
//				AutoCreate:        pulumi.Bool(false),
//				DefaultVisibility: pulumi.String("PUBLIC"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRegistryEnterpriseRepo, err := cs.NewRegistryEnterpriseRepo(ctx, "default", &cs.RegistryEnterpriseRepoArgs{
//				InstanceId: _default.ID(),
//				Namespace:  defaultRegistryEnterpriseNamespace.Name,
//				Name:       pulumi.String(name),
//				Summary:    pulumi.String("this is summary of my new repo"),
//				RepoType:   pulumi.String("PUBLIC"),
//				Detail:     pulumi.String("this is a public repo"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cr.NewChain(ctx, "default", &cr.ChainArgs{
//				ChainConfigs: cr.ChainChainConfigArray{
//					&cr.ChainChainConfigArgs{
//						Nodes: cr.ChainChainConfigNodeArray{
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(true),
//								NodeName: pulumi.String("DOCKER_IMAGE_BUILD"),
//							},
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(true),
//								NodeName: pulumi.String("DOCKER_IMAGE_PUSH"),
//							},
//							&cr.ChainChainConfigNodeArgs{
//								Enable:   pulumi.Bool(true),
//								NodeName: pulumi.String("VULNERABILITY_SCANNING"),
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											&cr.ChainChainConfigNodeNodeConfigDenyPolicyArgs{
//												IssueLevel: pulumi.String("MEDIUM"),
//												IssueCount: pulumi.String("1"),
//												Action:     pulumi.String("BLOCK_DELETE_TAG"),
//												Logic:      pulumi.String("AND"),
//											},
//										},
//									},
//								},
//							},
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(true),
//								NodeName: pulumi.String("ACTIVATE_REPLICATION"),
//							},
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(true),
//								NodeName: pulumi.String("TRIGGER"),
//							},
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(false),
//								NodeName: pulumi.String("SNAPSHOT"),
//							},
//							&cr.ChainChainConfigNodeArgs{
//								NodeConfigs: cr.ChainChainConfigNodeNodeConfigArray{
//									&cr.ChainChainConfigNodeNodeConfigArgs{
//										DenyPolicies: cr.ChainChainConfigNodeNodeConfigDenyPolicyArray{
//											nil,
//										},
//									},
//								},
//								Enable:   pulumi.Bool(false),
//								NodeName: pulumi.String("TRIGGER_SNAPSHOT"),
//							},
//						},
//						Routers: cr.ChainChainConfigRouterArray{
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("DOCKER_IMAGE_BUILD"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("DOCKER_IMAGE_PUSH"),
//									},
//								},
//							},
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("DOCKER_IMAGE_PUSH"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("VULNERABILITY_SCANNING"),
//									},
//								},
//							},
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("VULNERABILITY_SCANNING"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("ACTIVATE_REPLICATION"),
//									},
//								},
//							},
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("ACTIVATE_REPLICATION"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("TRIGGER"),
//									},
//								},
//							},
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("VULNERABILITY_SCANNING"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("SNAPSHOT"),
//									},
//								},
//							},
//							&cr.ChainChainConfigRouterArgs{
//								Froms: cr.ChainChainConfigRouterFromArray{
//									&cr.ChainChainConfigRouterFromArgs{
//										NodeName: pulumi.String("SNAPSHOT"),
//									},
//								},
//								Tos: cr.ChainChainConfigRouterToArray{
//									&cr.ChainChainConfigRouterToArgs{
//										NodeName: pulumi.String("TRIGGER_SNAPSHOT"),
//									},
//								},
//							},
//						},
//					},
//				},
//				ChainName:         pulumi.String(name),
//				Description:       pulumi.String(name),
//				InstanceId:        defaultRegistryEnterpriseNamespace.InstanceId,
//				RepoName:          defaultRegistryEnterpriseRepo.Name,
//				RepoNamespaceName: defaultRegistryEnterpriseNamespace.Name,
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
// CR Chain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cr/chain:Chain example <instance_id>:<chain_id>
// ```
type Chain struct {
	pulumi.CustomResourceState

	// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	ChainConfigs ChainChainConfigArrayOutput `pulumi:"chainConfigs"`
	// Delivery chain ID.
	ChainId pulumi.StringOutput `pulumi:"chainId"`
	// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
	ChainName pulumi.StringOutput `pulumi:"chainName"`
	// The description delivery chain.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of CR Enterprise Edition instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
	RepoName pulumi.StringPtrOutput `pulumi:"repoName"`
	// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	RepoNamespaceName pulumi.StringPtrOutput `pulumi:"repoNamespaceName"`
}

// NewChain registers a new resource with the given unique name, arguments, and options.
func NewChain(ctx *pulumi.Context,
	name string, args *ChainArgs, opts ...pulumi.ResourceOption) (*Chain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChainName == nil {
		return nil, errors.New("invalid value for required argument 'ChainName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Chain
	err := ctx.RegisterResource("alicloud:cr/chain:Chain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChain gets an existing Chain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChainState, opts ...pulumi.ResourceOption) (*Chain, error) {
	var resource Chain
	err := ctx.ReadResource("alicloud:cr/chain:Chain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Chain resources.
type chainState struct {
	// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	ChainConfigs []ChainChainConfig `pulumi:"chainConfigs"`
	// Delivery chain ID.
	ChainId *string `pulumi:"chainId"`
	// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
	ChainName *string `pulumi:"chainName"`
	// The description delivery chain.
	Description *string `pulumi:"description"`
	// The ID of CR Enterprise Edition instance.
	InstanceId *string `pulumi:"instanceId"`
	// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
	RepoName *string `pulumi:"repoName"`
	// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	RepoNamespaceName *string `pulumi:"repoNamespaceName"`
}

type ChainState struct {
	// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	ChainConfigs ChainChainConfigArrayInput
	// Delivery chain ID.
	ChainId pulumi.StringPtrInput
	// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
	ChainName pulumi.StringPtrInput
	// The description delivery chain.
	Description pulumi.StringPtrInput
	// The ID of CR Enterprise Edition instance.
	InstanceId pulumi.StringPtrInput
	// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
	RepoName pulumi.StringPtrInput
	// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	RepoNamespaceName pulumi.StringPtrInput
}

func (ChainState) ElementType() reflect.Type {
	return reflect.TypeOf((*chainState)(nil)).Elem()
}

type chainArgs struct {
	// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	ChainConfigs []ChainChainConfig `pulumi:"chainConfigs"`
	// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
	ChainName string `pulumi:"chainName"`
	// The description delivery chain.
	Description *string `pulumi:"description"`
	// The ID of CR Enterprise Edition instance.
	InstanceId string `pulumi:"instanceId"`
	// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
	RepoName *string `pulumi:"repoName"`
	// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	RepoNamespaceName *string `pulumi:"repoNamespaceName"`
}

// The set of arguments for constructing a Chain resource.
type ChainArgs struct {
	// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	ChainConfigs ChainChainConfigArrayInput
	// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
	ChainName pulumi.StringInput
	// The description delivery chain.
	Description pulumi.StringPtrInput
	// The ID of CR Enterprise Edition instance.
	InstanceId pulumi.StringInput
	// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
	RepoName pulumi.StringPtrInput
	// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
	RepoNamespaceName pulumi.StringPtrInput
}

func (ChainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chainArgs)(nil)).Elem()
}

type ChainInput interface {
	pulumi.Input

	ToChainOutput() ChainOutput
	ToChainOutputWithContext(ctx context.Context) ChainOutput
}

func (*Chain) ElementType() reflect.Type {
	return reflect.TypeOf((**Chain)(nil)).Elem()
}

func (i *Chain) ToChainOutput() ChainOutput {
	return i.ToChainOutputWithContext(context.Background())
}

func (i *Chain) ToChainOutputWithContext(ctx context.Context) ChainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChainOutput)
}

// ChainArrayInput is an input type that accepts ChainArray and ChainArrayOutput values.
// You can construct a concrete instance of `ChainArrayInput` via:
//
//	ChainArray{ ChainArgs{...} }
type ChainArrayInput interface {
	pulumi.Input

	ToChainArrayOutput() ChainArrayOutput
	ToChainArrayOutputWithContext(context.Context) ChainArrayOutput
}

type ChainArray []ChainInput

func (ChainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Chain)(nil)).Elem()
}

func (i ChainArray) ToChainArrayOutput() ChainArrayOutput {
	return i.ToChainArrayOutputWithContext(context.Background())
}

func (i ChainArray) ToChainArrayOutputWithContext(ctx context.Context) ChainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChainArrayOutput)
}

// ChainMapInput is an input type that accepts ChainMap and ChainMapOutput values.
// You can construct a concrete instance of `ChainMapInput` via:
//
//	ChainMap{ "key": ChainArgs{...} }
type ChainMapInput interface {
	pulumi.Input

	ToChainMapOutput() ChainMapOutput
	ToChainMapOutputWithContext(context.Context) ChainMapOutput
}

type ChainMap map[string]ChainInput

func (ChainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Chain)(nil)).Elem()
}

func (i ChainMap) ToChainMapOutput() ChainMapOutput {
	return i.ToChainMapOutputWithContext(context.Background())
}

func (i ChainMap) ToChainMapOutputWithContext(ctx context.Context) ChainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChainMapOutput)
}

type ChainOutput struct{ *pulumi.OutputState }

func (ChainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Chain)(nil)).Elem()
}

func (o ChainOutput) ToChainOutput() ChainOutput {
	return o
}

func (o ChainOutput) ToChainOutputWithContext(ctx context.Context) ChainOutput {
	return o
}

// The configuration of delivery chain. See `chainConfig` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
func (o ChainOutput) ChainConfigs() ChainChainConfigArrayOutput {
	return o.ApplyT(func(v *Chain) ChainChainConfigArrayOutput { return v.ChainConfigs }).(ChainChainConfigArrayOutput)
}

// Delivery chain ID.
func (o ChainOutput) ChainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringOutput { return v.ChainId }).(pulumi.StringOutput)
}

// The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators "_", "-", "." can be used, noted that the separator cannot be at the first or last position.
func (o ChainOutput) ChainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringOutput { return v.ChainName }).(pulumi.StringOutput)
}

// The description delivery chain.
func (o ChainOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of CR Enterprise Edition instance.
func (o ChainOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
func (o ChainOutput) RepoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringPtrOutput { return v.RepoName }).(pulumi.StringPtrOutput)
}

// The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
func (o ChainOutput) RepoNamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Chain) pulumi.StringPtrOutput { return v.RepoNamespaceName }).(pulumi.StringPtrOutput)
}

type ChainArrayOutput struct{ *pulumi.OutputState }

func (ChainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Chain)(nil)).Elem()
}

func (o ChainArrayOutput) ToChainArrayOutput() ChainArrayOutput {
	return o
}

func (o ChainArrayOutput) ToChainArrayOutputWithContext(ctx context.Context) ChainArrayOutput {
	return o
}

func (o ChainArrayOutput) Index(i pulumi.IntInput) ChainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Chain {
		return vs[0].([]*Chain)[vs[1].(int)]
	}).(ChainOutput)
}

type ChainMapOutput struct{ *pulumi.OutputState }

func (ChainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Chain)(nil)).Elem()
}

func (o ChainMapOutput) ToChainMapOutput() ChainMapOutput {
	return o
}

func (o ChainMapOutput) ToChainMapOutputWithContext(ctx context.Context) ChainMapOutput {
	return o
}

func (o ChainMapOutput) MapIndex(k pulumi.StringInput) ChainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Chain {
		return vs[0].(map[string]*Chain)[vs[1].(string)]
	}).(ChainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChainInput)(nil)).Elem(), &Chain{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChainArrayInput)(nil)).Elem(), ChainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChainMapInput)(nil)).Elem(), ChainMap{})
	pulumi.RegisterOutputType(ChainOutput{})
	pulumi.RegisterOutputType(ChainArrayOutput{})
	pulumi.RegisterOutputType(ChainMapOutput{})
}

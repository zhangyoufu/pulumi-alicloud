// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manager Container Registry Enterprise Edition sync rules.
//
// For information about Container Registry Enterprise Edition sync rules and how to use it, see [Create a Sync Rule](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createreposynctaskbyrule)
//
// > **NOTE:** Available since v1.90.0.
//
// > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
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
//			source, err := cr.NewRegistryEnterpriseInstance(ctx, "source", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewPeriod:   pulumi.Int(0),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(fmt.Sprintf("%v-source", name)),
//			})
//			if err != nil {
//				return err
//			}
//			target, err := cr.NewRegistryEnterpriseInstance(ctx, "target", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewPeriod:   pulumi.Int(0),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(fmt.Sprintf("%v-target", name)),
//			})
//			if err != nil {
//				return err
//			}
//			sourceRegistryEnterpriseNamespace, err := cs.NewRegistryEnterpriseNamespace(ctx, "source", &cs.RegistryEnterpriseNamespaceArgs{
//				InstanceId:        source.ID(),
//				Name:              pulumi.String(name),
//				AutoCreate:        pulumi.Bool(false),
//				DefaultVisibility: pulumi.String("PUBLIC"),
//			})
//			if err != nil {
//				return err
//			}
//			targetRegistryEnterpriseNamespace, err := cs.NewRegistryEnterpriseNamespace(ctx, "target", &cs.RegistryEnterpriseNamespaceArgs{
//				InstanceId:        target.ID(),
//				Name:              pulumi.String(name),
//				AutoCreate:        pulumi.Bool(false),
//				DefaultVisibility: pulumi.String("PUBLIC"),
//			})
//			if err != nil {
//				return err
//			}
//			sourceRegistryEnterpriseRepo, err := cs.NewRegistryEnterpriseRepo(ctx, "source", &cs.RegistryEnterpriseRepoArgs{
//				InstanceId: source.ID(),
//				Namespace:  sourceRegistryEnterpriseNamespace.Name,
//				Name:       pulumi.String(name),
//				Summary:    pulumi.String("this is summary of my new repo"),
//				RepoType:   pulumi.String("PUBLIC"),
//				Detail:     pulumi.String("this is a public repo"),
//			})
//			if err != nil {
//				return err
//			}
//			targetRegistryEnterpriseRepo, err := cs.NewRegistryEnterpriseRepo(ctx, "target", &cs.RegistryEnterpriseRepoArgs{
//				InstanceId: target.ID(),
//				Namespace:  targetRegistryEnterpriseNamespace.Name,
//				Name:       pulumi.String(name),
//				Summary:    pulumi.String("this is summary of my new repo"),
//				RepoType:   pulumi.String("PUBLIC"),
//				Detail:     pulumi.String("this is a public repo"),
//			})
//			if err != nil {
//				return err
//			}
//			_default, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cs.NewRegistryEnterpriseSyncRule(ctx, "default", &cs.RegistryEnterpriseSyncRuleArgs{
//				InstanceId:          source.ID(),
//				NamespaceName:       sourceRegistryEnterpriseNamespace.Name,
//				Name:                pulumi.String(name),
//				TargetRegionId:      pulumi.String(_default.Regions[0].Id),
//				TargetInstanceId:    target.ID(),
//				TargetNamespaceName: targetRegistryEnterpriseNamespace.Name,
//				TagFilter:           pulumi.String(".*"),
//				RepoName:            sourceRegistryEnterpriseRepo.Name,
//				TargetRepoName:      targetRegistryEnterpriseRepo.Name,
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
// Container Registry Enterprise Edition sync rule can be imported using the id. Format to `{instance_id}:{namespace_name}:{rule_id}`, e.g.
//
// ```sh
// $ pulumi import alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule default `cri-xxx:my-namespace:crsr-yyy`
// ```
type RegistryEnterpriseSyncRule struct {
	pulumi.CustomResourceState

	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrOutput `pulumi:"repoName"`
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection pulumi.StringOutput `pulumi:"syncDirection"`
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope pulumi.StringOutput `pulumi:"syncScope"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringOutput `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringOutput `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringOutput `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId pulumi.StringOutput `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrOutput `pulumi:"targetRepoName"`
}

// NewRegistryEnterpriseSyncRule registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnterpriseSyncRule(ctx *pulumi.Context,
	name string, args *RegistryEnterpriseSyncRuleArgs, opts ...pulumi.ResourceOption) (*RegistryEnterpriseSyncRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.TagFilter == nil {
		return nil, errors.New("invalid value for required argument 'TagFilter'")
	}
	if args.TargetInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetInstanceId'")
	}
	if args.TargetNamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'TargetNamespaceName'")
	}
	if args.TargetRegionId == nil {
		return nil, errors.New("invalid value for required argument 'TargetRegionId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegistryEnterpriseSyncRule
	err := ctx.RegisterResource("alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnterpriseSyncRule gets an existing RegistryEnterpriseSyncRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnterpriseSyncRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnterpriseSyncRuleState, opts ...pulumi.ResourceOption) (*RegistryEnterpriseSyncRule, error) {
	var resource RegistryEnterpriseSyncRule
	err := ctx.ReadResource("alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnterpriseSyncRule resources.
type registryEnterpriseSyncRuleState struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId *string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName *string `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName *string `pulumi:"repoName"`
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId *string `pulumi:"ruleId"`
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection *string `pulumi:"syncDirection"`
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope *string `pulumi:"syncScope"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter *string `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId *string `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName *string `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId *string `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName *string `pulumi:"targetRepoName"`
}

type RegistryEnterpriseSyncRuleState struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringPtrInput
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrInput
	// The uuid of Container Registry Enterprise Edition sync rule.
	RuleId pulumi.StringPtrInput
	// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
	SyncDirection pulumi.StringPtrInput
	// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
	SyncScope pulumi.StringPtrInput
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringPtrInput
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringPtrInput
	// The target region to be synchronized.
	TargetRegionId pulumi.StringPtrInput
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrInput
}

func (RegistryEnterpriseSyncRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseSyncRuleState)(nil)).Elem()
}

type registryEnterpriseSyncRuleArgs struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition sync rule.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName *string `pulumi:"repoName"`
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter string `pulumi:"tagFilter"`
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId string `pulumi:"targetInstanceId"`
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName string `pulumi:"targetNamespaceName"`
	// The target region to be synchronized.
	TargetRegionId string `pulumi:"targetRegionId"`
	// Name of the target repository.
	TargetRepoName *string `pulumi:"targetRepoName"`
}

// The set of arguments for constructing a RegistryEnterpriseSyncRule resource.
type RegistryEnterpriseSyncRuleArgs struct {
	// ID of Container Registry Enterprise Edition source instance.
	InstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition sync rule.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
	NamespaceName pulumi.StringInput
	// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
	RepoName pulumi.StringPtrInput
	// The regular expression used to filter image tags for synchronization in the source repository.
	TagFilter pulumi.StringInput
	// ID of Container Registry Enterprise Edition target instance to be synchronized.
	TargetInstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
	TargetNamespaceName pulumi.StringInput
	// The target region to be synchronized.
	TargetRegionId pulumi.StringInput
	// Name of the target repository.
	TargetRepoName pulumi.StringPtrInput
}

func (RegistryEnterpriseSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseSyncRuleArgs)(nil)).Elem()
}

type RegistryEnterpriseSyncRuleInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput
	ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput
}

func (*RegistryEnterpriseSyncRule) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput {
	return i.ToRegistryEnterpriseSyncRuleOutputWithContext(context.Background())
}

func (i *RegistryEnterpriseSyncRule) ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleOutput)
}

// RegistryEnterpriseSyncRuleArrayInput is an input type that accepts RegistryEnterpriseSyncRuleArray and RegistryEnterpriseSyncRuleArrayOutput values.
// You can construct a concrete instance of `RegistryEnterpriseSyncRuleArrayInput` via:
//
//	RegistryEnterpriseSyncRuleArray{ RegistryEnterpriseSyncRuleArgs{...} }
type RegistryEnterpriseSyncRuleArrayInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput
	ToRegistryEnterpriseSyncRuleArrayOutputWithContext(context.Context) RegistryEnterpriseSyncRuleArrayOutput
}

type RegistryEnterpriseSyncRuleArray []RegistryEnterpriseSyncRuleInput

func (RegistryEnterpriseSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (i RegistryEnterpriseSyncRuleArray) ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput {
	return i.ToRegistryEnterpriseSyncRuleArrayOutputWithContext(context.Background())
}

func (i RegistryEnterpriseSyncRuleArray) ToRegistryEnterpriseSyncRuleArrayOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleArrayOutput)
}

// RegistryEnterpriseSyncRuleMapInput is an input type that accepts RegistryEnterpriseSyncRuleMap and RegistryEnterpriseSyncRuleMapOutput values.
// You can construct a concrete instance of `RegistryEnterpriseSyncRuleMapInput` via:
//
//	RegistryEnterpriseSyncRuleMap{ "key": RegistryEnterpriseSyncRuleArgs{...} }
type RegistryEnterpriseSyncRuleMapInput interface {
	pulumi.Input

	ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput
	ToRegistryEnterpriseSyncRuleMapOutputWithContext(context.Context) RegistryEnterpriseSyncRuleMapOutput
}

type RegistryEnterpriseSyncRuleMap map[string]RegistryEnterpriseSyncRuleInput

func (RegistryEnterpriseSyncRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (i RegistryEnterpriseSyncRuleMap) ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput {
	return i.ToRegistryEnterpriseSyncRuleMapOutputWithContext(context.Background())
}

func (i RegistryEnterpriseSyncRuleMap) ToRegistryEnterpriseSyncRuleMapOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseSyncRuleMapOutput)
}

type RegistryEnterpriseSyncRuleOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRuleOutput() RegistryEnterpriseSyncRuleOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleOutput) ToRegistryEnterpriseSyncRuleOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleOutput {
	return o
}

// ID of Container Registry Enterprise Edition source instance.
func (o RegistryEnterpriseSyncRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition sync rule.
func (o RegistryEnterpriseSyncRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.
func (o RegistryEnterpriseSyncRuleOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// Name of the source repository which should be set together with `targetRepoName`, if empty means that the synchronization scope is the entire namespace level.
func (o RegistryEnterpriseSyncRuleOutput) RepoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringPtrOutput { return v.RepoName }).(pulumi.StringPtrOutput)
}

// The uuid of Container Registry Enterprise Edition sync rule.
func (o RegistryEnterpriseSyncRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// `FROM` or `TO`, the direction of synchronization. `FROM` means source instance, `TO` means target instance.
func (o RegistryEnterpriseSyncRuleOutput) SyncDirection() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.SyncDirection }).(pulumi.StringOutput)
}

// `REPO` or `NAMESPACE`,the scope that the synchronization rule applies.
func (o RegistryEnterpriseSyncRuleOutput) SyncScope() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.SyncScope }).(pulumi.StringOutput)
}

// The regular expression used to filter image tags for synchronization in the source repository.
func (o RegistryEnterpriseSyncRuleOutput) TagFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.TagFilter }).(pulumi.StringOutput)
}

// ID of Container Registry Enterprise Edition target instance to be synchronized.
func (o RegistryEnterpriseSyncRuleOutput) TargetInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.TargetInstanceId }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.
func (o RegistryEnterpriseSyncRuleOutput) TargetNamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.TargetNamespaceName }).(pulumi.StringOutput)
}

// The target region to be synchronized.
func (o RegistryEnterpriseSyncRuleOutput) TargetRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringOutput { return v.TargetRegionId }).(pulumi.StringOutput)
}

// Name of the target repository.
func (o RegistryEnterpriseSyncRuleOutput) TargetRepoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseSyncRule) pulumi.StringPtrOutput { return v.TargetRepoName }).(pulumi.StringPtrOutput)
}

type RegistryEnterpriseSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (o RegistryEnterpriseSyncRuleArrayOutput) ToRegistryEnterpriseSyncRuleArrayOutput() RegistryEnterpriseSyncRuleArrayOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleArrayOutput) ToRegistryEnterpriseSyncRuleArrayOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleArrayOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleArrayOutput) Index(i pulumi.IntInput) RegistryEnterpriseSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegistryEnterpriseSyncRule {
		return vs[0].([]*RegistryEnterpriseSyncRule)[vs[1].(int)]
	}).(RegistryEnterpriseSyncRuleOutput)
}

type RegistryEnterpriseSyncRuleMapOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseSyncRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseSyncRule)(nil)).Elem()
}

func (o RegistryEnterpriseSyncRuleMapOutput) ToRegistryEnterpriseSyncRuleMapOutput() RegistryEnterpriseSyncRuleMapOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleMapOutput) ToRegistryEnterpriseSyncRuleMapOutputWithContext(ctx context.Context) RegistryEnterpriseSyncRuleMapOutput {
	return o
}

func (o RegistryEnterpriseSyncRuleMapOutput) MapIndex(k pulumi.StringInput) RegistryEnterpriseSyncRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegistryEnterpriseSyncRule {
		return vs[0].(map[string]*RegistryEnterpriseSyncRule)[vs[1].(string)]
	}).(RegistryEnterpriseSyncRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleInput)(nil)).Elem(), &RegistryEnterpriseSyncRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleArrayInput)(nil)).Elem(), RegistryEnterpriseSyncRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseSyncRuleMapInput)(nil)).Elem(), RegistryEnterpriseSyncRuleMap{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleArrayOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseSyncRuleMapOutput{})
}

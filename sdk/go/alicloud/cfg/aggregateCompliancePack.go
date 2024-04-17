// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Config Aggregate Compliance Pack resource.
//
// For information about Cloud Config Aggregate Compliance Pack and how to use it, see [What is Aggregate Compliance Pack](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createaggregatecompliancepack).
//
// > **NOTE:** Available since v1.124.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := resourcemanager.GetAccounts(ctx, &resourcemanager.GetAccountsArgs{
//				Status: pulumi.StringRef("CreateSuccess"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultAggregator, err := cfg.NewAggregator(ctx, "default", &cfg.AggregatorArgs{
//				AggregatorAccounts: cfg.AggregatorAggregatorAccountArray{
//					&cfg.AggregatorAggregatorAccountArgs{
//						AccountId:   pulumi.String(_default.Accounts[0].AccountId),
//						AccountName: pulumi.String(_default.Accounts[0].DisplayName),
//						AccountType: pulumi.String("ResourceDirectory"),
//					},
//				},
//				AggregatorName: pulumi.String(name),
//				Description:    pulumi.String(name),
//				AggregatorType: pulumi.String("CUSTOM"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultAggregateConfigRule, err := cfg.NewAggregateConfigRule(ctx, "default", &cfg.AggregateConfigRuleArgs{
//				AggregateConfigRuleName: pulumi.String("contains-tag"),
//				AggregatorId:            defaultAggregator.ID(),
//				ConfigRuleTriggerTypes:  pulumi.String("ConfigurationItemChangeNotification"),
//				SourceOwner:             pulumi.String("ALIYUN"),
//				SourceIdentifier:        pulumi.String("contains-tag"),
//				Description:             pulumi.String(name),
//				RiskLevel:               pulumi.Int(1),
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::ECS::Instance"),
//				},
//				InputParameters: pulumi.Map{
//					"key":   pulumi.Any("example"),
//					"value": pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewAggregateCompliancePack(ctx, "default", &cfg.AggregateCompliancePackArgs{
//				AggregateCompliancePackName: pulumi.String(name),
//				AggregatorId:                defaultAggregator.ID(),
//				Description:                 pulumi.String(name),
//				RiskLevel:                   pulumi.Int(1),
//				ConfigRuleIds: cfg.AggregateCompliancePackConfigRuleIdArray{
//					&cfg.AggregateCompliancePackConfigRuleIdArgs{
//						ConfigRuleId: defaultAggregateConfigRule.ConfigRuleId,
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
// Cloud Config Aggregate Compliance Pack can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack example <aggregator_id>:<aggregator_compliance_pack_id>
// ```
type AggregateCompliancePack struct {
	pulumi.CustomResourceState

	// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
	AggregateCompliancePackName pulumi.StringOutput `pulumi:"aggregateCompliancePackName"`
	// The ID of the compliance package.
	AggregatorCompliancePackId pulumi.StringOutput `pulumi:"aggregatorCompliancePackId"`
	// The ID of aggregator.
	AggregatorId pulumi.StringOutput `pulumi:"aggregatorId"`
	// The Template ID of compliance package.
	CompliancePackTemplateId pulumi.StringPtrOutput `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds AggregateCompliancePackConfigRuleIdArrayOutput `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	ConfigRules AggregateCompliancePackConfigRuleArrayOutput `pulumi:"configRules"`
	// The description of compliance package.
	Description pulumi.StringOutput `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntOutput `pulumi:"riskLevel"`
	// The status of the Aggregate Compliance Pack.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAggregateCompliancePack registers a new resource with the given unique name, arguments, and options.
func NewAggregateCompliancePack(ctx *pulumi.Context,
	name string, args *AggregateCompliancePackArgs, opts ...pulumi.ResourceOption) (*AggregateCompliancePack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AggregateCompliancePackName == nil {
		return nil, errors.New("invalid value for required argument 'AggregateCompliancePackName'")
	}
	if args.AggregatorId == nil {
		return nil, errors.New("invalid value for required argument 'AggregatorId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RiskLevel == nil {
		return nil, errors.New("invalid value for required argument 'RiskLevel'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AggregateCompliancePack
	err := ctx.RegisterResource("alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregateCompliancePack gets an existing AggregateCompliancePack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregateCompliancePack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregateCompliancePackState, opts ...pulumi.ResourceOption) (*AggregateCompliancePack, error) {
	var resource AggregateCompliancePack
	err := ctx.ReadResource("alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregateCompliancePack resources.
type aggregateCompliancePackState struct {
	// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
	AggregateCompliancePackName *string `pulumi:"aggregateCompliancePackName"`
	// The ID of the compliance package.
	AggregatorCompliancePackId *string `pulumi:"aggregatorCompliancePackId"`
	// The ID of aggregator.
	AggregatorId *string `pulumi:"aggregatorId"`
	// The Template ID of compliance package.
	CompliancePackTemplateId *string `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds []AggregateCompliancePackConfigRuleId `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	ConfigRules []AggregateCompliancePackConfigRule `pulumi:"configRules"`
	// The description of compliance package.
	Description *string `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel *int `pulumi:"riskLevel"`
	// The status of the Aggregate Compliance Pack.
	Status *string `pulumi:"status"`
}

type AggregateCompliancePackState struct {
	// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
	AggregateCompliancePackName pulumi.StringPtrInput
	// The ID of the compliance package.
	AggregatorCompliancePackId pulumi.StringPtrInput
	// The ID of aggregator.
	AggregatorId pulumi.StringPtrInput
	// The Template ID of compliance package.
	CompliancePackTemplateId pulumi.StringPtrInput
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds AggregateCompliancePackConfigRuleIdArrayInput
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	ConfigRules AggregateCompliancePackConfigRuleArrayInput
	// The description of compliance package.
	Description pulumi.StringPtrInput
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntPtrInput
	// The status of the Aggregate Compliance Pack.
	Status pulumi.StringPtrInput
}

func (AggregateCompliancePackState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateCompliancePackState)(nil)).Elem()
}

type aggregateCompliancePackArgs struct {
	// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
	AggregateCompliancePackName string `pulumi:"aggregateCompliancePackName"`
	// The ID of aggregator.
	AggregatorId string `pulumi:"aggregatorId"`
	// The Template ID of compliance package.
	CompliancePackTemplateId *string `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds []AggregateCompliancePackConfigRuleId `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	ConfigRules []AggregateCompliancePackConfigRule `pulumi:"configRules"`
	// The description of compliance package.
	Description string `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel int `pulumi:"riskLevel"`
}

// The set of arguments for constructing a AggregateCompliancePack resource.
type AggregateCompliancePackArgs struct {
	// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
	AggregateCompliancePackName pulumi.StringInput
	// The ID of aggregator.
	AggregatorId pulumi.StringInput
	// The Template ID of compliance package.
	CompliancePackTemplateId pulumi.StringPtrInput
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds AggregateCompliancePackConfigRuleIdArrayInput
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	ConfigRules AggregateCompliancePackConfigRuleArrayInput
	// The description of compliance package.
	Description pulumi.StringInput
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntInput
}

func (AggregateCompliancePackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateCompliancePackArgs)(nil)).Elem()
}

type AggregateCompliancePackInput interface {
	pulumi.Input

	ToAggregateCompliancePackOutput() AggregateCompliancePackOutput
	ToAggregateCompliancePackOutputWithContext(ctx context.Context) AggregateCompliancePackOutput
}

func (*AggregateCompliancePack) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateCompliancePack)(nil)).Elem()
}

func (i *AggregateCompliancePack) ToAggregateCompliancePackOutput() AggregateCompliancePackOutput {
	return i.ToAggregateCompliancePackOutputWithContext(context.Background())
}

func (i *AggregateCompliancePack) ToAggregateCompliancePackOutputWithContext(ctx context.Context) AggregateCompliancePackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateCompliancePackOutput)
}

// AggregateCompliancePackArrayInput is an input type that accepts AggregateCompliancePackArray and AggregateCompliancePackArrayOutput values.
// You can construct a concrete instance of `AggregateCompliancePackArrayInput` via:
//
//	AggregateCompliancePackArray{ AggregateCompliancePackArgs{...} }
type AggregateCompliancePackArrayInput interface {
	pulumi.Input

	ToAggregateCompliancePackArrayOutput() AggregateCompliancePackArrayOutput
	ToAggregateCompliancePackArrayOutputWithContext(context.Context) AggregateCompliancePackArrayOutput
}

type AggregateCompliancePackArray []AggregateCompliancePackInput

func (AggregateCompliancePackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateCompliancePack)(nil)).Elem()
}

func (i AggregateCompliancePackArray) ToAggregateCompliancePackArrayOutput() AggregateCompliancePackArrayOutput {
	return i.ToAggregateCompliancePackArrayOutputWithContext(context.Background())
}

func (i AggregateCompliancePackArray) ToAggregateCompliancePackArrayOutputWithContext(ctx context.Context) AggregateCompliancePackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateCompliancePackArrayOutput)
}

// AggregateCompliancePackMapInput is an input type that accepts AggregateCompliancePackMap and AggregateCompliancePackMapOutput values.
// You can construct a concrete instance of `AggregateCompliancePackMapInput` via:
//
//	AggregateCompliancePackMap{ "key": AggregateCompliancePackArgs{...} }
type AggregateCompliancePackMapInput interface {
	pulumi.Input

	ToAggregateCompliancePackMapOutput() AggregateCompliancePackMapOutput
	ToAggregateCompliancePackMapOutputWithContext(context.Context) AggregateCompliancePackMapOutput
}

type AggregateCompliancePackMap map[string]AggregateCompliancePackInput

func (AggregateCompliancePackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateCompliancePack)(nil)).Elem()
}

func (i AggregateCompliancePackMap) ToAggregateCompliancePackMapOutput() AggregateCompliancePackMapOutput {
	return i.ToAggregateCompliancePackMapOutputWithContext(context.Background())
}

func (i AggregateCompliancePackMap) ToAggregateCompliancePackMapOutputWithContext(ctx context.Context) AggregateCompliancePackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateCompliancePackMapOutput)
}

type AggregateCompliancePackOutput struct{ *pulumi.OutputState }

func (AggregateCompliancePackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateCompliancePack)(nil)).Elem()
}

func (o AggregateCompliancePackOutput) ToAggregateCompliancePackOutput() AggregateCompliancePackOutput {
	return o
}

func (o AggregateCompliancePackOutput) ToAggregateCompliancePackOutputWithContext(ctx context.Context) AggregateCompliancePackOutput {
	return o
}

// The name of compliance package name. **NOTE:** From version 1.145.0, `aggregateCompliancePackName` can be modified.
func (o AggregateCompliancePackOutput) AggregateCompliancePackName() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringOutput { return v.AggregateCompliancePackName }).(pulumi.StringOutput)
}

// The ID of the compliance package.
func (o AggregateCompliancePackOutput) AggregatorCompliancePackId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringOutput { return v.AggregatorCompliancePackId }).(pulumi.StringOutput)
}

// The ID of aggregator.
func (o AggregateCompliancePackOutput) AggregatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringOutput { return v.AggregatorId }).(pulumi.StringOutput)
}

// The Template ID of compliance package.
func (o AggregateCompliancePackOutput) CompliancePackTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringPtrOutput { return v.CompliancePackTemplateId }).(pulumi.StringPtrOutput)
}

// A list of Config Rule IDs. See `configRuleIds` below.
func (o AggregateCompliancePackOutput) ConfigRuleIds() AggregateCompliancePackConfigRuleIdArrayOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) AggregateCompliancePackConfigRuleIdArrayOutput {
		return v.ConfigRuleIds
	}).(AggregateCompliancePackConfigRuleIdArrayOutput)
}

// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
//
// Deprecated: Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
func (o AggregateCompliancePackOutput) ConfigRules() AggregateCompliancePackConfigRuleArrayOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) AggregateCompliancePackConfigRuleArrayOutput { return v.ConfigRules }).(AggregateCompliancePackConfigRuleArrayOutput)
}

// The description of compliance package.
func (o AggregateCompliancePackOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The Risk Level. Valid values:
func (o AggregateCompliancePackOutput) RiskLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.IntOutput { return v.RiskLevel }).(pulumi.IntOutput)
}

// The status of the Aggregate Compliance Pack.
func (o AggregateCompliancePackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateCompliancePack) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AggregateCompliancePackArrayOutput struct{ *pulumi.OutputState }

func (AggregateCompliancePackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateCompliancePack)(nil)).Elem()
}

func (o AggregateCompliancePackArrayOutput) ToAggregateCompliancePackArrayOutput() AggregateCompliancePackArrayOutput {
	return o
}

func (o AggregateCompliancePackArrayOutput) ToAggregateCompliancePackArrayOutputWithContext(ctx context.Context) AggregateCompliancePackArrayOutput {
	return o
}

func (o AggregateCompliancePackArrayOutput) Index(i pulumi.IntInput) AggregateCompliancePackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AggregateCompliancePack {
		return vs[0].([]*AggregateCompliancePack)[vs[1].(int)]
	}).(AggregateCompliancePackOutput)
}

type AggregateCompliancePackMapOutput struct{ *pulumi.OutputState }

func (AggregateCompliancePackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateCompliancePack)(nil)).Elem()
}

func (o AggregateCompliancePackMapOutput) ToAggregateCompliancePackMapOutput() AggregateCompliancePackMapOutput {
	return o
}

func (o AggregateCompliancePackMapOutput) ToAggregateCompliancePackMapOutputWithContext(ctx context.Context) AggregateCompliancePackMapOutput {
	return o
}

func (o AggregateCompliancePackMapOutput) MapIndex(k pulumi.StringInput) AggregateCompliancePackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AggregateCompliancePack {
		return vs[0].(map[string]*AggregateCompliancePack)[vs[1].(string)]
	}).(AggregateCompliancePackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateCompliancePackInput)(nil)).Elem(), &AggregateCompliancePack{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateCompliancePackArrayInput)(nil)).Elem(), AggregateCompliancePackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateCompliancePackMapInput)(nil)).Elem(), AggregateCompliancePackMap{})
	pulumi.RegisterOutputType(AggregateCompliancePackOutput{})
	pulumi.RegisterOutputType(AggregateCompliancePackArrayOutput{})
	pulumi.RegisterOutputType(AggregateCompliancePackMapOutput{})
}

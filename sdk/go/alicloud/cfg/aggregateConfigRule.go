// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Config Aggregate Config Rule resource.
//
// For information about Cloud Config Aggregate Config Rule and how to use it, see [What is Aggregate Config Rule](https://www.alibabacloud.com/help/doc-detail/154216.html).
//
// > **NOTE:** Available in v1.124.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleAggregator, err := cfg.NewAggregator(ctx, "exampleAggregator", &cfg.AggregatorArgs{
//				AggregatorAccounts: cfg.AggregatorAggregatorAccountArray{
//					&cfg.AggregatorAggregatorAccountArgs{
//						AccountId:   pulumi.String("140278452670****"),
//						AccountName: pulumi.String("test-2"),
//						AccountType: pulumi.String("ResourceDirectory"),
//					},
//				},
//				AggregatorName: pulumi.String("tf-testaccaggregator"),
//				Description:    pulumi.String("tf-testaccaggregator"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewAggregateConfigRule(ctx, "exampleAggregateConfigRule", &cfg.AggregateConfigRuleArgs{
//				AggregateConfigRuleName: pulumi.String("tf-testaccconfig1234"),
//				AggregatorId:            exampleAggregator.ID(),
//				ConfigRuleTriggerTypes:  pulumi.String("ConfigurationItemChangeNotification"),
//				SourceOwner:             pulumi.String("ALIYUN"),
//				SourceIdentifier:        pulumi.String("ecs-cpu-min-count-limit"),
//				RiskLevel:               pulumi.Int(1),
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::ECS::Instance"),
//				},
//				InputParameters: pulumi.AnyMap{
//					"cpuCount": pulumi.Any("4"),
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
//
// ## Import
//
// Cloud Config Aggregate Config Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cfg/aggregateConfigRule:AggregateConfigRule example <aggregator_id>:<config_rule_id>
//
// ```
type AggregateConfigRule struct {
	pulumi.CustomResourceState

	// The name of the rule.
	AggregateConfigRuleName pulumi.StringOutput `pulumi:"aggregateConfigRuleName"`
	// The Aggregator Id.
	AggregatorId pulumi.StringOutput `pulumi:"aggregatorId"`
	// (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
	ConfigRuleId pulumi.StringOutput `pulumi:"configRuleId"`
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringOutput `pulumi:"configRuleTriggerTypes"`
	// The description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrOutput `pulumi:"excludeResourceIdsScope"`
	// The settings map of the input parameters for the rule.
	InputParameters pulumi.MapOutput `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringOutput `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrOutput `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrOutput `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayOutput `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
	RiskLevel pulumi.IntOutput `pulumi:"riskLevel"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringOutput `pulumi:"sourceIdentifier"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringOutput `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrOutput `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrOutput `pulumi:"tagValueScope"`
}

// NewAggregateConfigRule registers a new resource with the given unique name, arguments, and options.
func NewAggregateConfigRule(ctx *pulumi.Context,
	name string, args *AggregateConfigRuleArgs, opts ...pulumi.ResourceOption) (*AggregateConfigRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AggregateConfigRuleName == nil {
		return nil, errors.New("invalid value for required argument 'AggregateConfigRuleName'")
	}
	if args.AggregatorId == nil {
		return nil, errors.New("invalid value for required argument 'AggregatorId'")
	}
	if args.ConfigRuleTriggerTypes == nil {
		return nil, errors.New("invalid value for required argument 'ConfigRuleTriggerTypes'")
	}
	if args.ResourceTypesScopes == nil {
		return nil, errors.New("invalid value for required argument 'ResourceTypesScopes'")
	}
	if args.RiskLevel == nil {
		return nil, errors.New("invalid value for required argument 'RiskLevel'")
	}
	if args.SourceIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'SourceIdentifier'")
	}
	if args.SourceOwner == nil {
		return nil, errors.New("invalid value for required argument 'SourceOwner'")
	}
	var resource AggregateConfigRule
	err := ctx.RegisterResource("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregateConfigRule gets an existing AggregateConfigRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregateConfigRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregateConfigRuleState, opts ...pulumi.ResourceOption) (*AggregateConfigRule, error) {
	var resource AggregateConfigRule
	err := ctx.ReadResource("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregateConfigRule resources.
type aggregateConfigRuleState struct {
	// The name of the rule.
	AggregateConfigRuleName *string `pulumi:"aggregateConfigRuleName"`
	// The Aggregator Id.
	AggregatorId *string `pulumi:"aggregatorId"`
	// (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
	ConfigRuleId *string `pulumi:"configRuleId"`
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes *string `pulumi:"configRuleTriggerTypes"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// The settings map of the input parameters for the rule.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
	RiskLevel *int `pulumi:"riskLevel"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier *string `pulumi:"sourceIdentifier"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner *string `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope *string `pulumi:"tagValueScope"`
}

type AggregateConfigRuleState struct {
	// The name of the rule.
	AggregateConfigRuleName pulumi.StringPtrInput
	// The Aggregator Id.
	AggregatorId pulumi.StringPtrInput
	// (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
	ConfigRuleId pulumi.StringPtrInput
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// The settings map of the input parameters for the rule.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
	RiskLevel pulumi.IntPtrInput
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringPtrInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringPtrInput
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrInput
}

func (AggregateConfigRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateConfigRuleState)(nil)).Elem()
}

type aggregateConfigRuleArgs struct {
	// The name of the rule.
	AggregateConfigRuleName string `pulumi:"aggregateConfigRuleName"`
	// The Aggregator Id.
	AggregatorId string `pulumi:"aggregatorId"`
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes string `pulumi:"configRuleTriggerTypes"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// The settings map of the input parameters for the rule.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
	RiskLevel int `pulumi:"riskLevel"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier string `pulumi:"sourceIdentifier"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner string `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope *string `pulumi:"tagValueScope"`
}

// The set of arguments for constructing a AggregateConfigRule resource.
type AggregateConfigRuleArgs struct {
	// The name of the rule.
	AggregateConfigRuleName pulumi.StringInput
	// The Aggregator Id.
	AggregatorId pulumi.StringInput
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// The settings map of the input parameters for the rule.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
	RiskLevel pulumi.IntInput
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringInput
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrInput
}

func (AggregateConfigRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateConfigRuleArgs)(nil)).Elem()
}

type AggregateConfigRuleInput interface {
	pulumi.Input

	ToAggregateConfigRuleOutput() AggregateConfigRuleOutput
	ToAggregateConfigRuleOutputWithContext(ctx context.Context) AggregateConfigRuleOutput
}

func (*AggregateConfigRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateConfigRule)(nil)).Elem()
}

func (i *AggregateConfigRule) ToAggregateConfigRuleOutput() AggregateConfigRuleOutput {
	return i.ToAggregateConfigRuleOutputWithContext(context.Background())
}

func (i *AggregateConfigRule) ToAggregateConfigRuleOutputWithContext(ctx context.Context) AggregateConfigRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateConfigRuleOutput)
}

// AggregateConfigRuleArrayInput is an input type that accepts AggregateConfigRuleArray and AggregateConfigRuleArrayOutput values.
// You can construct a concrete instance of `AggregateConfigRuleArrayInput` via:
//
//	AggregateConfigRuleArray{ AggregateConfigRuleArgs{...} }
type AggregateConfigRuleArrayInput interface {
	pulumi.Input

	ToAggregateConfigRuleArrayOutput() AggregateConfigRuleArrayOutput
	ToAggregateConfigRuleArrayOutputWithContext(context.Context) AggregateConfigRuleArrayOutput
}

type AggregateConfigRuleArray []AggregateConfigRuleInput

func (AggregateConfigRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateConfigRule)(nil)).Elem()
}

func (i AggregateConfigRuleArray) ToAggregateConfigRuleArrayOutput() AggregateConfigRuleArrayOutput {
	return i.ToAggregateConfigRuleArrayOutputWithContext(context.Background())
}

func (i AggregateConfigRuleArray) ToAggregateConfigRuleArrayOutputWithContext(ctx context.Context) AggregateConfigRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateConfigRuleArrayOutput)
}

// AggregateConfigRuleMapInput is an input type that accepts AggregateConfigRuleMap and AggregateConfigRuleMapOutput values.
// You can construct a concrete instance of `AggregateConfigRuleMapInput` via:
//
//	AggregateConfigRuleMap{ "key": AggregateConfigRuleArgs{...} }
type AggregateConfigRuleMapInput interface {
	pulumi.Input

	ToAggregateConfigRuleMapOutput() AggregateConfigRuleMapOutput
	ToAggregateConfigRuleMapOutputWithContext(context.Context) AggregateConfigRuleMapOutput
}

type AggregateConfigRuleMap map[string]AggregateConfigRuleInput

func (AggregateConfigRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateConfigRule)(nil)).Elem()
}

func (i AggregateConfigRuleMap) ToAggregateConfigRuleMapOutput() AggregateConfigRuleMapOutput {
	return i.ToAggregateConfigRuleMapOutputWithContext(context.Background())
}

func (i AggregateConfigRuleMap) ToAggregateConfigRuleMapOutputWithContext(ctx context.Context) AggregateConfigRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateConfigRuleMapOutput)
}

type AggregateConfigRuleOutput struct{ *pulumi.OutputState }

func (AggregateConfigRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateConfigRule)(nil)).Elem()
}

func (o AggregateConfigRuleOutput) ToAggregateConfigRuleOutput() AggregateConfigRuleOutput {
	return o
}

func (o AggregateConfigRuleOutput) ToAggregateConfigRuleOutputWithContext(ctx context.Context) AggregateConfigRuleOutput {
	return o
}

// The name of the rule.
func (o AggregateConfigRuleOutput) AggregateConfigRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.AggregateConfigRuleName }).(pulumi.StringOutput)
}

// The Aggregator Id.
func (o AggregateConfigRuleOutput) AggregatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.AggregatorId }).(pulumi.StringOutput)
}

// (Available in 1.141.0+) The rule ID of Aggregate Config Rule.
func (o AggregateConfigRuleOutput) ConfigRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.ConfigRuleId }).(pulumi.StringOutput)
}

// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
func (o AggregateConfigRuleOutput) ConfigRuleTriggerTypes() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.ConfigRuleTriggerTypes }).(pulumi.StringOutput)
}

// The description of the rule.
func (o AggregateConfigRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
func (o AggregateConfigRuleOutput) ExcludeResourceIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.ExcludeResourceIdsScope }).(pulumi.StringPtrOutput)
}

// The settings map of the input parameters for the rule.
func (o AggregateConfigRuleOutput) InputParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.MapOutput { return v.InputParameters }).(pulumi.MapOutput)
}

// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `configRuleTriggerTypes` is `ScheduledNotification`.
func (o AggregateConfigRuleOutput) MaximumExecutionFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.MaximumExecutionFrequency }).(pulumi.StringOutput)
}

// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
func (o AggregateConfigRuleOutput) RegionIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.RegionIdsScope }).(pulumi.StringPtrOutput)
}

// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
func (o AggregateConfigRuleOutput) ResourceGroupIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.ResourceGroupIdsScope }).(pulumi.StringPtrOutput)
}

// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
func (o AggregateConfigRuleOutput) ResourceTypesScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringArrayOutput { return v.ResourceTypesScopes }).(pulumi.StringArrayOutput)
}

// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
func (o AggregateConfigRuleOutput) RiskLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.IntOutput { return v.RiskLevel }).(pulumi.IntOutput)
}

// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
func (o AggregateConfigRuleOutput) SourceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.SourceIdentifier }).(pulumi.StringOutput)
}

// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
func (o AggregateConfigRuleOutput) SourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.SourceOwner }).(pulumi.StringOutput)
}

// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
func (o AggregateConfigRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The rule monitors the tag key, only applies to rules created based on managed rules.
func (o AggregateConfigRuleOutput) TagKeyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.TagKeyScope }).(pulumi.StringPtrOutput)
}

// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
func (o AggregateConfigRuleOutput) TagValueScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AggregateConfigRule) pulumi.StringPtrOutput { return v.TagValueScope }).(pulumi.StringPtrOutput)
}

type AggregateConfigRuleArrayOutput struct{ *pulumi.OutputState }

func (AggregateConfigRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateConfigRule)(nil)).Elem()
}

func (o AggregateConfigRuleArrayOutput) ToAggregateConfigRuleArrayOutput() AggregateConfigRuleArrayOutput {
	return o
}

func (o AggregateConfigRuleArrayOutput) ToAggregateConfigRuleArrayOutputWithContext(ctx context.Context) AggregateConfigRuleArrayOutput {
	return o
}

func (o AggregateConfigRuleArrayOutput) Index(i pulumi.IntInput) AggregateConfigRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AggregateConfigRule {
		return vs[0].([]*AggregateConfigRule)[vs[1].(int)]
	}).(AggregateConfigRuleOutput)
}

type AggregateConfigRuleMapOutput struct{ *pulumi.OutputState }

func (AggregateConfigRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateConfigRule)(nil)).Elem()
}

func (o AggregateConfigRuleMapOutput) ToAggregateConfigRuleMapOutput() AggregateConfigRuleMapOutput {
	return o
}

func (o AggregateConfigRuleMapOutput) ToAggregateConfigRuleMapOutputWithContext(ctx context.Context) AggregateConfigRuleMapOutput {
	return o
}

func (o AggregateConfigRuleMapOutput) MapIndex(k pulumi.StringInput) AggregateConfigRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AggregateConfigRule {
		return vs[0].(map[string]*AggregateConfigRule)[vs[1].(string)]
	}).(AggregateConfigRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateConfigRuleInput)(nil)).Elem(), &AggregateConfigRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateConfigRuleArrayInput)(nil)).Elem(), AggregateConfigRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateConfigRuleMapInput)(nil)).Elem(), AggregateConfigRuleMap{})
	pulumi.RegisterOutputType(AggregateConfigRuleOutput{})
	pulumi.RegisterOutputType(AggregateConfigRuleArrayOutput{})
	pulumi.RegisterOutputType(AggregateConfigRuleMapOutput{})
}

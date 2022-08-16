// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a a Alicloud Config Rule resource. Cloud Config checks the validity of resources based on rules. You can create rules to evaluate resources as needed.
// For information about Alicloud Config Rule and how to use it, see [What is Alicloud Config Rule](https://www.alibabacloud.com/help/doc-detail/154216.html).
//
// > **NOTE:** Available in v1.99.0+.
//
// > **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-southeast-1`.
//
// > **NOTE:** If you use custom rules, you need to create your own rule functions in advance. Please refer to the link for [Create a custom rule.](https://www.alibabacloud.com/help/en/doc-detail/127405.htm)
//
// ## Example Usage
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
//			_, err := cfg.NewRule(ctx, "example", &cfg.RuleArgs{
//				ConfigRuleTriggerTypes: pulumi.String("ConfigurationItemChangeNotification"),
//				Description:            pulumi.String("ecs instances in vpc"),
//				InputParameters: pulumi.AnyMap{
//					"vpcIds": pulumi.Any("vpc-uf6gksw4ctjd******"),
//				},
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::ECS::Instance"),
//				},
//				RiskLevel:        pulumi.Int(1),
//				RuleName:         pulumi.String("instances-in-vpc"),
//				SourceIdentifier: pulumi.String("ecs-instances-in-vpc"),
//				SourceOwner:      pulumi.String("ALIYUN"),
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
// Alicloud Config Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cfg/rule:Rule this cr-ed4bad756057********
//
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringOutput `pulumi:"configRuleTriggerTypes"`
	// The description of the Config Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrOutput `pulumi:"excludeResourceIdsScope"`
	// Threshold value for managed rule triggering.
	InputParameters pulumi.MapOutput `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringOutput `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrOutput `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrOutput `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayOutput `pulumi:"resourceTypesScopes"`
	// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
	RiskLevel pulumi.IntOutput `pulumi:"riskLevel"`
	// The name of the Config Rule.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringArrayOutput `pulumi:"scopeComplianceResourceTypes"`
	// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringOutput `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringOutput `pulumi:"sourceIdentifier"`
	// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringOutput `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringOutput `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrOutput `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrOutput `pulumi:"tagValueScope"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RiskLevel == nil {
		return nil, errors.New("invalid value for required argument 'RiskLevel'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.SourceIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'SourceIdentifier'")
	}
	if args.SourceOwner == nil {
		return nil, errors.New("invalid value for required argument 'SourceOwner'")
	}
	var resource Rule
	err := ctx.RegisterResource("alicloud:cfg/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("alicloud:cfg/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes *string `pulumi:"configRuleTriggerTypes"`
	// The description of the Config Rule.
	Description *string `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// Threshold value for managed rule triggering.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
	RiskLevel *int `pulumi:"riskLevel"`
	// The name of the Config Rule.
	RuleName *string `pulumi:"ruleName"`
	// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes []string `pulumi:"scopeComplianceResourceTypes"`
	// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType *string `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier *string `pulumi:"sourceIdentifier"`
	// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency *string `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner *string `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope *string `pulumi:"tagValueScope"`
}

type RuleState struct {
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringPtrInput
	// The description of the Config Rule.
	Description pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// Threshold value for managed rule triggering.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
	RiskLevel pulumi.IntPtrInput
	// The name of the Config Rule.
	RuleName pulumi.StringPtrInput
	// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringArrayInput
	// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringPtrInput
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringPtrInput
	// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringPtrInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringPtrInput
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes *string `pulumi:"configRuleTriggerTypes"`
	// The description of the Config Rule.
	Description *string `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// Threshold value for managed rule triggering.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
	RiskLevel int `pulumi:"riskLevel"`
	// The name of the Config Rule.
	RuleName string `pulumi:"ruleName"`
	// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes []string `pulumi:"scopeComplianceResourceTypes"`
	// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType *string `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier string `pulumi:"sourceIdentifier"`
	// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency *string `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner string `pulumi:"sourceOwner"`
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope *string `pulumi:"tagValueScope"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringPtrInput
	// The description of the Config Rule.
	Description pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// Threshold value for managed rule triggering.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
	RiskLevel pulumi.IntInput
	// The name of the Config Rule.
	RuleName pulumi.StringInput
	// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringArrayInput
	// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringPtrInput
	// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
	SourceIdentifier pulumi.StringInput
	// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringPtrInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
	SourceOwner pulumi.StringInput
	// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
	TagValueScope pulumi.StringPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
func (o RuleOutput) ConfigRuleTriggerTypes() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ConfigRuleTriggerTypes }).(pulumi.StringOutput)
}

// The description of the Config Rule.
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, custom rule this field is empty.
func (o RuleOutput) ExcludeResourceIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ExcludeResourceIdsScope }).(pulumi.StringPtrOutput)
}

// Threshold value for managed rule triggering.
func (o RuleOutput) InputParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *Rule) pulumi.MapOutput { return v.InputParameters }).(pulumi.MapOutput)
}

// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
func (o RuleOutput) MaximumExecutionFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.MaximumExecutionFrequency }).(pulumi.StringOutput)
}

// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
func (o RuleOutput) RegionIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.RegionIdsScope }).(pulumi.StringPtrOutput)
}

// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
func (o RuleOutput) ResourceGroupIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ResourceGroupIdsScope }).(pulumi.StringPtrOutput)
}

// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
func (o RuleOutput) ResourceTypesScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringArrayOutput { return v.ResourceTypesScopes }).(pulumi.StringArrayOutput)
}

// The risk level of the Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
func (o RuleOutput) RiskLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.RiskLevel }).(pulumi.IntOutput)
}

// The name of the Config Rule.
func (o RuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Field `scopeComplianceResourceTypes` has been deprecated from provider version 1.124.1. New field `resourceTypesScope` instead.
//
// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
func (o RuleOutput) ScopeComplianceResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringArrayOutput { return v.ScopeComplianceResourceTypes }).(pulumi.StringArrayOutput)
}

// Field `sourceDetailMessageType` has been deprecated from provider version 1.124.1. New field `configRuleTriggerTypes` instead.
//
// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
func (o RuleOutput) SourceDetailMessageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceDetailMessageType }).(pulumi.StringOutput)
}

// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
func (o RuleOutput) SourceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceIdentifier }).(pulumi.StringOutput)
}

// Field `sourceMaximumExecutionFrequency` has been deprecated from provider version 1.124.1. New field `maximumExecutionFrequency` instead.
//
// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
func (o RuleOutput) SourceMaximumExecutionFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceMaximumExecutionFrequency }).(pulumi.StringOutput)
}

// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
func (o RuleOutput) SourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceOwner }).(pulumi.StringOutput)
}

// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
func (o RuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The rule monitors the tag key, only applies to rules created based on managed rules.
func (o RuleOutput) TagKeyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.TagKeyScope }).(pulumi.StringPtrOutput)
}

// The rule monitors the tag value, use with the `tagKeyScope` options. only applies to rules created based on managed rules.
func (o RuleOutput) TagValueScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.TagValueScope }).(pulumi.StringPtrOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}

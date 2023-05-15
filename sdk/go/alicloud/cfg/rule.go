// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Config Rule resource.
//
// For information about Config Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/).
//
// > **NOTE:** Available in v1.204.0+.
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
//			_, err := cfg.NewRule(ctx, "default", &cfg.RuleArgs{
//				ConfigRuleTriggerTypes:  pulumi.String("ConfigurationItemChangeNotification"),
//				Description:             pulumi.String("关联的资源类型下实体资源均已有指定标签，存在没有指定标签的资源则视为“不合规”。"),
//				ExcludeResourceIdsScope: pulumi.String("test"),
//				InputParameters: pulumi.AnyMap{
//					"foo": pulumi.Any("terraform"),
//					"var": pulumi.Any("terraform"),
//				},
//				RegionIdsScope:        pulumi.String("cn-hangzhou"),
//				ResourceGroupIdsScope: pulumi.String("rg-acfmvoh45rhcfly"),
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::RDS::DBInstance"),
//				},
//				RiskLevel:        pulumi.Int(1),
//				RuleName:         pulumi.String("tf-cicd-rule-by-required-tags"),
//				SourceIdentifier: pulumi.String("required-tags"),
//				SourceOwner:      pulumi.String("ALIYUN"),
//				TagKeyScope:      pulumi.String("test"),
//				TagValueScope:    pulumi.String("test"),
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
// Config Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cfg/rule:Rule example <id>
//
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The ID of Alicloud account.
	AccountId pulumi.IntOutput `pulumi:"accountId"`
	// compliance information.
	Compliance RuleComplianceOutput `pulumi:"compliance"`
	// Compliance Package ID.
	CompliancePackId pulumi.StringOutput `pulumi:"compliancePackId"`
	// config rule arn.
	ConfigRuleArn pulumi.StringOutput `pulumi:"configRuleArn"`
	// The ID of the rule.
	ConfigRuleId pulumi.StringOutput `pulumi:"configRuleId"`
	// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringOutput `pulumi:"configRuleTriggerTypes"`
	// The timestamp when the rule was created.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// The description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The event source of the rule.
	EventSource pulumi.StringOutput `pulumi:"eventSource"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrOutput `pulumi:"excludeResourceIdsScope"`
	// The settings of the input parameters for the rule.
	InputParameters pulumi.MapOutput `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringOutput `pulumi:"maximumExecutionFrequency"`
	// The timestamp when the rule was last modified.
	ModifiedTimestamp pulumi.IntOutput `pulumi:"modifiedTimestamp"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrOutput `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrOutput `pulumi:"resourceGroupIdsScope"`
	// The types of the resources to be evaluated against the rule.
	ResourceTypesScopes pulumi.StringArrayOutput `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
	RiskLevel pulumi.IntOutput `pulumi:"riskLevel"`
	// The name of the rule.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringOutput `pulumi:"scopeComplianceResourceTypes"`
	// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringOutput `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
	SourceIdentifier pulumi.StringOutput `pulumi:"sourceIdentifier"`
	// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringOutput `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
	SourceOwner pulumi.StringOutput `pulumi:"sourceOwner"`
	// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
	Status pulumi.StringOutput `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrOutput `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, only applies to rules created based on managed rules.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
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
	// The ID of Alicloud account.
	AccountId *int `pulumi:"accountId"`
	// compliance information.
	Compliance *RuleCompliance `pulumi:"compliance"`
	// Compliance Package ID.
	CompliancePackId *string `pulumi:"compliancePackId"`
	// config rule arn.
	ConfigRuleArn *string `pulumi:"configRuleArn"`
	// The ID of the rule.
	ConfigRuleId *string `pulumi:"configRuleId"`
	// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes *string `pulumi:"configRuleTriggerTypes"`
	// The timestamp when the rule was created.
	CreateTime *int `pulumi:"createTime"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The event source of the rule.
	EventSource *string `pulumi:"eventSource"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// The settings of the input parameters for the rule.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The timestamp when the rule was last modified.
	ModifiedTimestamp *int `pulumi:"modifiedTimestamp"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// The types of the resources to be evaluated against the rule.
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
	RiskLevel *int `pulumi:"riskLevel"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
	// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes *string `pulumi:"scopeComplianceResourceTypes"`
	// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType *string `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
	SourceIdentifier *string `pulumi:"sourceIdentifier"`
	// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency *string `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
	SourceOwner *string `pulumi:"sourceOwner"`
	// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, only applies to rules created based on managed rules.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TagValueScope *string `pulumi:"tagValueScope"`
}

type RuleState struct {
	// The ID of Alicloud account.
	AccountId pulumi.IntPtrInput
	// compliance information.
	Compliance RuleCompliancePtrInput
	// Compliance Package ID.
	CompliancePackId pulumi.StringPtrInput
	// config rule arn.
	ConfigRuleArn pulumi.StringPtrInput
	// The ID of the rule.
	ConfigRuleId pulumi.StringPtrInput
	// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringPtrInput
	// The timestamp when the rule was created.
	CreateTime pulumi.IntPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The event source of the rule.
	EventSource pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// The settings of the input parameters for the rule.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The timestamp when the rule was last modified.
	ModifiedTimestamp pulumi.IntPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// The types of the resources to be evaluated against the rule.
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
	RiskLevel pulumi.IntPtrInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
	// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringPtrInput
	// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringPtrInput
	// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
	SourceIdentifier pulumi.StringPtrInput
	// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringPtrInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
	SourceOwner pulumi.StringPtrInput
	// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, only applies to rules created based on managed rules.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TagValueScope pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes *string `pulumi:"configRuleTriggerTypes"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope *string `pulumi:"excludeResourceIdsScope"`
	// The settings of the input parameters for the rule.
	InputParameters map[string]interface{} `pulumi:"inputParameters"`
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope *string `pulumi:"regionIdsScope"`
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope *string `pulumi:"resourceGroupIdsScope"`
	// The types of the resources to be evaluated against the rule.
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
	RiskLevel int `pulumi:"riskLevel"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
	// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes *string `pulumi:"scopeComplianceResourceTypes"`
	// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType *string `pulumi:"sourceDetailMessageType"`
	// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
	SourceIdentifier string `pulumi:"sourceIdentifier"`
	// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency *string `pulumi:"sourceMaximumExecutionFrequency"`
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
	SourceOwner string `pulumi:"sourceOwner"`
	// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
	Status *string `pulumi:"status"`
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// The rule monitors the tag value, only applies to rules created based on managed rules.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	TagValueScope *string `pulumi:"tagValueScope"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
	ConfigRuleTriggerTypes pulumi.StringPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
	ExcludeResourceIdsScope pulumi.StringPtrInput
	// The settings of the input parameters for the rule.
	InputParameters pulumi.MapInput
	// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
	RegionIdsScope pulumi.StringPtrInput
	// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
	ResourceGroupIdsScope pulumi.StringPtrInput
	// The types of the resources to be evaluated against the rule.
	ResourceTypesScopes pulumi.StringArrayInput
	// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
	RiskLevel pulumi.IntInput
	// The name of the rule.
	RuleName pulumi.StringInput
	// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	//
	// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
	ScopeComplianceResourceTypes pulumi.StringPtrInput
	// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	//
	// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
	SourceDetailMessageType pulumi.StringPtrInput
	// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
	SourceIdentifier pulumi.StringInput
	// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	//
	// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
	SourceMaximumExecutionFrequency pulumi.StringPtrInput
	// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
	SourceOwner pulumi.StringInput
	// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
	Status pulumi.StringPtrInput
	// The rule monitors the tag key, only applies to rules created based on managed rules.
	TagKeyScope pulumi.StringPtrInput
	// The rule monitors the tag value, only applies to rules created based on managed rules.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
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

// The ID of Alicloud account.
func (o RuleOutput) AccountId() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.AccountId }).(pulumi.IntOutput)
}

// compliance information.
func (o RuleOutput) Compliance() RuleComplianceOutput {
	return o.ApplyT(func(v *Rule) RuleComplianceOutput { return v.Compliance }).(RuleComplianceOutput)
}

// Compliance Package ID.
func (o RuleOutput) CompliancePackId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.CompliancePackId }).(pulumi.StringOutput)
}

// config rule arn.
func (o RuleOutput) ConfigRuleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ConfigRuleArn }).(pulumi.StringOutput)
}

// The ID of the rule.
func (o RuleOutput) ConfigRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ConfigRuleId }).(pulumi.StringOutput)
}

// The trigger type of the rule. Valid values:  `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
func (o RuleOutput) ConfigRuleTriggerTypes() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ConfigRuleTriggerTypes }).(pulumi.StringOutput)
}

// The timestamp when the rule was created.
func (o RuleOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// The description of the rule.
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The event source of the rule.
func (o RuleOutput) EventSource() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.EventSource }).(pulumi.StringOutput)
}

// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
func (o RuleOutput) ExcludeResourceIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ExcludeResourceIdsScope }).(pulumi.StringPtrOutput)
}

// The settings of the input parameters for the rule.
func (o RuleOutput) InputParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *Rule) pulumi.MapOutput { return v.InputParameters }).(pulumi.MapOutput)
}

// The frequency of the compliance evaluations, it is required if the ConfigRuleTriggerTypes value is ScheduledNotification. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`.
func (o RuleOutput) MaximumExecutionFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.MaximumExecutionFrequency }).(pulumi.StringOutput)
}

// The timestamp when the rule was last modified.
func (o RuleOutput) ModifiedTimestamp() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.ModifiedTimestamp }).(pulumi.IntOutput)
}

// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
func (o RuleOutput) RegionIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.RegionIdsScope }).(pulumi.StringPtrOutput)
}

// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
func (o RuleOutput) ResourceGroupIdsScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ResourceGroupIdsScope }).(pulumi.StringPtrOutput)
}

// The types of the resources to be evaluated against the rule.
func (o RuleOutput) ResourceTypesScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringArrayOutput { return v.ResourceTypesScopes }).(pulumi.StringArrayOutput)
}

// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info
func (o RuleOutput) RiskLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.RiskLevel }).(pulumi.IntOutput)
}

// The name of the rule.
func (o RuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
//
// Deprecated: Field 'scope_compliance_resource_types' has been deprecated from provider version 1.124.1. New field 'resource_types_scope' instead.
func (o RuleOutput) ScopeComplianceResourceTypes() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ScopeComplianceResourceTypes }).(pulumi.StringOutput)
}

// Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
//
// Deprecated: Field 'source_detail_message_type' has been deprecated from provider version 1.124.1. New field 'config_rule_trigger_types' instead.
func (o RuleOutput) SourceDetailMessageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceDetailMessageType }).(pulumi.StringOutput)
}

// The identifier of the rule.  For a managed rule, the value is the name of the managed rule. For a custom rule, the value is the ARN of the custom rule.
func (o RuleOutput) SourceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceIdentifier }).(pulumi.StringOutput)
}

// Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
//
// Deprecated: Field 'source_maximum_execution_frequency' has been deprecated from provider version 1.124.1. New field 'maximum_execution_frequency' instead.
func (o RuleOutput) SourceMaximumExecutionFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceMaximumExecutionFrequency }).(pulumi.StringOutput)
}

// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values:  `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule
func (o RuleOutput) SourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.SourceOwner }).(pulumi.StringOutput)
}

// The status of the rule. Valid values: ACTIVE: The rule is monitoring the configurations of target resources. DELETING_RESULTS: The compliance evaluation result returned by the rule is being deleted. EVALUATING: The rule is triggered and is evaluating whether the configurations of target resources are compliant. INACTIVE: The rule is disabled from monitoring the configurations of target resources.
func (o RuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The rule monitors the tag key, only applies to rules created based on managed rules.
func (o RuleOutput) TagKeyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.TagKeyScope }).(pulumi.StringPtrOutput)
}

// The rule monitors the tag value, only applies to rules created based on managed rules.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
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

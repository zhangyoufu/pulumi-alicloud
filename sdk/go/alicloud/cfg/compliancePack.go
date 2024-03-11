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

// Provides a Cloud Config Compliance Pack resource.
//
// For information about Cloud Config Compliance Pack and how to use it, see [What is Compliance Pack](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createcompliancepack).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example-config-name"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultRegions, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			rule1, err := cfg.NewRule(ctx, "rule1", &cfg.RuleArgs{
//				Description:               pulumi.String(name),
//				SourceOwner:               pulumi.String("ALIYUN"),
//				SourceIdentifier:          pulumi.String("ram-user-ak-create-date-expired-check"),
//				RiskLevel:                 pulumi.Int(1),
//				MaximumExecutionFrequency: pulumi.String("TwentyFour_Hours"),
//				RegionIdsScope:            *pulumi.String(defaultRegions.Regions[0].Id),
//				ConfigRuleTriggerTypes:    pulumi.String("ScheduledNotification"),
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::RAM::User"),
//				},
//				RuleName: pulumi.String("ciscompliancecheck_ram-user-ak-create-date-expired-check"),
//				InputParameters: pulumi.Map{
//					"days": pulumi.Any("90"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			rule2, err := cfg.NewRule(ctx, "rule2", &cfg.RuleArgs{
//				Description:            pulumi.String(name),
//				SourceOwner:            pulumi.String("ALIYUN"),
//				SourceIdentifier:       pulumi.String("adb-cluster-maintain-time-check"),
//				RiskLevel:              pulumi.Int(2),
//				RegionIdsScope:         *pulumi.String(defaultRegions.Regions[0].Id),
//				ConfigRuleTriggerTypes: pulumi.String("ScheduledNotification"),
//				ResourceTypesScopes: pulumi.StringArray{
//					pulumi.String("ACS::ADB::DBCluster"),
//				},
//				RuleName: pulumi.String("governance-evaluation-adb-cluster-maintain-time-check"),
//				InputParameters: pulumi.Map{
//					"maintainTimes": pulumi.Any("02:00-04:00,06:00-08:00,12:00-13:00"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewCompliancePack(ctx, "defaultCompliancePack", &cfg.CompliancePackArgs{
//				CompliancePackName: pulumi.String(name),
//				Description:        pulumi.String("CloudGovernanceCenter evaluation"),
//				RiskLevel:          pulumi.Int(2),
//				ConfigRuleIds: cfg.CompliancePackConfigRuleIdArray{
//					&cfg.CompliancePackConfigRuleIdArgs{
//						ConfigRuleId: rule1.ID(),
//					},
//					&cfg.CompliancePackConfigRuleIdArgs{
//						ConfigRuleId: rule2.ID(),
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
// Cloud Config Compliance Pack can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cfg/compliancePack:CompliancePack example <id>
// ```
type CompliancePack struct {
	pulumi.CustomResourceState

	// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
	CompliancePackName pulumi.StringOutput `pulumi:"compliancePackName"`
	// Compliance Package Template Id.
	CompliancePackTemplateId pulumi.StringPtrOutput `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds CompliancePackConfigRuleIdArrayOutput `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
	ConfigRules CompliancePackConfigRuleArrayOutput `pulumi:"configRules"`
	// The Description of compliance pack.
	Description pulumi.StringOutput `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntOutput `pulumi:"riskLevel"`
	// The status of the Compliance Pack.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCompliancePack registers a new resource with the given unique name, arguments, and options.
func NewCompliancePack(ctx *pulumi.Context,
	name string, args *CompliancePackArgs, opts ...pulumi.ResourceOption) (*CompliancePack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompliancePackName == nil {
		return nil, errors.New("invalid value for required argument 'CompliancePackName'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RiskLevel == nil {
		return nil, errors.New("invalid value for required argument 'RiskLevel'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CompliancePack
	err := ctx.RegisterResource("alicloud:cfg/compliancePack:CompliancePack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompliancePack gets an existing CompliancePack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompliancePack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompliancePackState, opts ...pulumi.ResourceOption) (*CompliancePack, error) {
	var resource CompliancePack
	err := ctx.ReadResource("alicloud:cfg/compliancePack:CompliancePack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CompliancePack resources.
type compliancePackState struct {
	// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
	CompliancePackName *string `pulumi:"compliancePackName"`
	// Compliance Package Template Id.
	CompliancePackTemplateId *string `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds []CompliancePackConfigRuleId `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
	ConfigRules []CompliancePackConfigRule `pulumi:"configRules"`
	// The Description of compliance pack.
	Description *string `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel *int `pulumi:"riskLevel"`
	// The status of the Compliance Pack.
	Status *string `pulumi:"status"`
}

type CompliancePackState struct {
	// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
	CompliancePackName pulumi.StringPtrInput
	// Compliance Package Template Id.
	CompliancePackTemplateId pulumi.StringPtrInput
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds CompliancePackConfigRuleIdArrayInput
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
	ConfigRules CompliancePackConfigRuleArrayInput
	// The Description of compliance pack.
	Description pulumi.StringPtrInput
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntPtrInput
	// The status of the Compliance Pack.
	Status pulumi.StringPtrInput
}

func (CompliancePackState) ElementType() reflect.Type {
	return reflect.TypeOf((*compliancePackState)(nil)).Elem()
}

type compliancePackArgs struct {
	// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
	CompliancePackName string `pulumi:"compliancePackName"`
	// Compliance Package Template Id.
	CompliancePackTemplateId *string `pulumi:"compliancePackTemplateId"`
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds []CompliancePackConfigRuleId `pulumi:"configRuleIds"`
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
	ConfigRules []CompliancePackConfigRule `pulumi:"configRules"`
	// The Description of compliance pack.
	Description string `pulumi:"description"`
	// The Risk Level. Valid values:
	RiskLevel int `pulumi:"riskLevel"`
}

// The set of arguments for constructing a CompliancePack resource.
type CompliancePackArgs struct {
	// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
	CompliancePackName pulumi.StringInput
	// Compliance Package Template Id.
	CompliancePackTemplateId pulumi.StringPtrInput
	// A list of Config Rule IDs. See `configRuleIds` below.
	ConfigRuleIds CompliancePackConfigRuleIdArrayInput
	// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
	//
	// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
	ConfigRules CompliancePackConfigRuleArrayInput
	// The Description of compliance pack.
	Description pulumi.StringInput
	// The Risk Level. Valid values:
	RiskLevel pulumi.IntInput
}

func (CompliancePackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*compliancePackArgs)(nil)).Elem()
}

type CompliancePackInput interface {
	pulumi.Input

	ToCompliancePackOutput() CompliancePackOutput
	ToCompliancePackOutputWithContext(ctx context.Context) CompliancePackOutput
}

func (*CompliancePack) ElementType() reflect.Type {
	return reflect.TypeOf((**CompliancePack)(nil)).Elem()
}

func (i *CompliancePack) ToCompliancePackOutput() CompliancePackOutput {
	return i.ToCompliancePackOutputWithContext(context.Background())
}

func (i *CompliancePack) ToCompliancePackOutputWithContext(ctx context.Context) CompliancePackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompliancePackOutput)
}

// CompliancePackArrayInput is an input type that accepts CompliancePackArray and CompliancePackArrayOutput values.
// You can construct a concrete instance of `CompliancePackArrayInput` via:
//
//	CompliancePackArray{ CompliancePackArgs{...} }
type CompliancePackArrayInput interface {
	pulumi.Input

	ToCompliancePackArrayOutput() CompliancePackArrayOutput
	ToCompliancePackArrayOutputWithContext(context.Context) CompliancePackArrayOutput
}

type CompliancePackArray []CompliancePackInput

func (CompliancePackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompliancePack)(nil)).Elem()
}

func (i CompliancePackArray) ToCompliancePackArrayOutput() CompliancePackArrayOutput {
	return i.ToCompliancePackArrayOutputWithContext(context.Background())
}

func (i CompliancePackArray) ToCompliancePackArrayOutputWithContext(ctx context.Context) CompliancePackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompliancePackArrayOutput)
}

// CompliancePackMapInput is an input type that accepts CompliancePackMap and CompliancePackMapOutput values.
// You can construct a concrete instance of `CompliancePackMapInput` via:
//
//	CompliancePackMap{ "key": CompliancePackArgs{...} }
type CompliancePackMapInput interface {
	pulumi.Input

	ToCompliancePackMapOutput() CompliancePackMapOutput
	ToCompliancePackMapOutputWithContext(context.Context) CompliancePackMapOutput
}

type CompliancePackMap map[string]CompliancePackInput

func (CompliancePackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompliancePack)(nil)).Elem()
}

func (i CompliancePackMap) ToCompliancePackMapOutput() CompliancePackMapOutput {
	return i.ToCompliancePackMapOutputWithContext(context.Background())
}

func (i CompliancePackMap) ToCompliancePackMapOutputWithContext(ctx context.Context) CompliancePackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompliancePackMapOutput)
}

type CompliancePackOutput struct{ *pulumi.OutputState }

func (CompliancePackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompliancePack)(nil)).Elem()
}

func (o CompliancePackOutput) ToCompliancePackOutput() CompliancePackOutput {
	return o
}

func (o CompliancePackOutput) ToCompliancePackOutputWithContext(ctx context.Context) CompliancePackOutput {
	return o
}

// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliancePackName` can be modified.
func (o CompliancePackOutput) CompliancePackName() pulumi.StringOutput {
	return o.ApplyT(func(v *CompliancePack) pulumi.StringOutput { return v.CompliancePackName }).(pulumi.StringOutput)
}

// Compliance Package Template Id.
func (o CompliancePackOutput) CompliancePackTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompliancePack) pulumi.StringPtrOutput { return v.CompliancePackTemplateId }).(pulumi.StringPtrOutput)
}

// A list of Config Rule IDs. See `configRuleIds` below.
func (o CompliancePackOutput) ConfigRuleIds() CompliancePackConfigRuleIdArrayOutput {
	return o.ApplyT(func(v *CompliancePack) CompliancePackConfigRuleIdArrayOutput { return v.ConfigRuleIds }).(CompliancePackConfigRuleIdArrayOutput)
}

// A list of Config Rules. See `configRules` below. **NOTE:** Field `configRules` has been deprecated from provider version 1.141.0. New field `configRuleIds` instead.
//
// Deprecated: Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
func (o CompliancePackOutput) ConfigRules() CompliancePackConfigRuleArrayOutput {
	return o.ApplyT(func(v *CompliancePack) CompliancePackConfigRuleArrayOutput { return v.ConfigRules }).(CompliancePackConfigRuleArrayOutput)
}

// The Description of compliance pack.
func (o CompliancePackOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CompliancePack) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The Risk Level. Valid values:
func (o CompliancePackOutput) RiskLevel() pulumi.IntOutput {
	return o.ApplyT(func(v *CompliancePack) pulumi.IntOutput { return v.RiskLevel }).(pulumi.IntOutput)
}

// The status of the Compliance Pack.
func (o CompliancePackOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CompliancePack) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CompliancePackArrayOutput struct{ *pulumi.OutputState }

func (CompliancePackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompliancePack)(nil)).Elem()
}

func (o CompliancePackArrayOutput) ToCompliancePackArrayOutput() CompliancePackArrayOutput {
	return o
}

func (o CompliancePackArrayOutput) ToCompliancePackArrayOutputWithContext(ctx context.Context) CompliancePackArrayOutput {
	return o
}

func (o CompliancePackArrayOutput) Index(i pulumi.IntInput) CompliancePackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CompliancePack {
		return vs[0].([]*CompliancePack)[vs[1].(int)]
	}).(CompliancePackOutput)
}

type CompliancePackMapOutput struct{ *pulumi.OutputState }

func (CompliancePackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompliancePack)(nil)).Elem()
}

func (o CompliancePackMapOutput) ToCompliancePackMapOutput() CompliancePackMapOutput {
	return o
}

func (o CompliancePackMapOutput) ToCompliancePackMapOutputWithContext(ctx context.Context) CompliancePackMapOutput {
	return o
}

func (o CompliancePackMapOutput) MapIndex(k pulumi.StringInput) CompliancePackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CompliancePack {
		return vs[0].(map[string]*CompliancePack)[vs[1].(string)]
	}).(CompliancePackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CompliancePackInput)(nil)).Elem(), &CompliancePack{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompliancePackArrayInput)(nil)).Elem(), CompliancePackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompliancePackMapInput)(nil)).Elem(), CompliancePackMap{})
	pulumi.RegisterOutputType(CompliancePackOutput{})
	pulumi.RegisterOutputType(CompliancePackArrayOutput{})
	pulumi.RegisterOutputType(CompliancePackMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg
{
    /// <summary>
    /// Provides a Cloud Config Aggregate Config Rule resource.
    /// 
    /// For information about Cloud Config Aggregate Config Rule and how to use it, see [What is Aggregate Config Rule](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createaggregateconfigrule).
    /// 
    /// &gt; **NOTE:** Available since v1.124.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultAccounts = AliCloud.ResourceManager.GetAccounts.Invoke(new()
    ///     {
    ///         Status = "CreateSuccess",
    ///     });
    /// 
    ///     var defaultAggregator = new AliCloud.Cfg.Aggregator("defaultAggregator", new()
    ///     {
    ///         AggregatorAccounts = new[]
    ///         {
    ///             new AliCloud.Cfg.Inputs.AggregatorAggregatorAccountArgs
    ///             {
    ///                 AccountId = defaultAccounts.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.AccountId),
    ///                 AccountName = defaultAccounts.Apply(getAccountsResult =&gt; getAccountsResult.Accounts[0]?.DisplayName),
    ///                 AccountType = "ResourceDirectory",
    ///             },
    ///         },
    ///         AggregatorName = name,
    ///         Description = name,
    ///         AggregatorType = "CUSTOM",
    ///     });
    /// 
    ///     var defaultAggregateConfigRule = new AliCloud.Cfg.AggregateConfigRule("defaultAggregateConfigRule", new()
    ///     {
    ///         AggregateConfigRuleName = "contains-tag",
    ///         AggregatorId = defaultAggregator.Id,
    ///         ConfigRuleTriggerTypes = "ConfigurationItemChangeNotification",
    ///         SourceOwner = "ALIYUN",
    ///         SourceIdentifier = "contains-tag",
    ///         RiskLevel = 1,
    ///         ResourceTypesScopes = new[]
    ///         {
    ///             "ACS::ECS::Instance",
    ///         },
    ///         InputParameters = 
    ///         {
    ///             { "key", "example" },
    ///             { "value", "example" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Config Aggregate Config Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cfg/aggregateConfigRule:AggregateConfigRule example "&lt;aggregator_id&gt;:&lt;config_rule_id&gt;"
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cfg/aggregateConfigRule:AggregateConfigRule")]
    public partial class AggregateConfigRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Output("aggregateConfigRuleName")]
        public Output<string> AggregateConfigRuleName { get; private set; } = null!;

        /// <summary>
        /// The Aggregator Id.
        /// </summary>
        [Output("aggregatorId")]
        public Output<string> AggregatorId { get; private set; } = null!;

        /// <summary>
        /// (Available since v1.141.0) The rule ID of Aggregate Config Rule.
        /// </summary>
        [Output("configRuleId")]
        public Output<string> ConfigRuleId { get; private set; } = null!;

        /// <summary>
        /// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
        /// </summary>
        [Output("configRuleTriggerTypes")]
        public Output<string> ConfigRuleTriggerTypes { get; private set; } = null!;

        /// <summary>
        /// The description of the rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
        /// </summary>
        [Output("excludeResourceIdsScope")]
        public Output<string?> ExcludeResourceIdsScope { get; private set; } = null!;

        /// <summary>
        /// The settings map of the input parameters for the rule.
        /// </summary>
        [Output("inputParameters")]
        public Output<ImmutableDictionary<string, object>?> InputParameters { get; private set; } = null!;

        /// <summary>
        /// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `config_rule_trigger_types` is `ScheduledNotification`.
        /// </summary>
        [Output("maximumExecutionFrequency")]
        public Output<string> MaximumExecutionFrequency { get; private set; } = null!;

        /// <summary>
        /// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Output("regionIdsScope")]
        public Output<string?> RegionIdsScope { get; private set; } = null!;

        /// <summary>
        /// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Output("resourceGroupIdsScope")]
        public Output<string?> ResourceGroupIdsScope { get; private set; } = null!;

        /// <summary>
        /// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        /// </summary>
        [Output("resourceTypesScopes")]
        public Output<ImmutableArray<string>> ResourceTypesScopes { get; private set; } = null!;

        /// <summary>
        /// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
        /// </summary>
        [Output("riskLevel")]
        public Output<int> RiskLevel { get; private set; } = null!;

        /// <summary>
        /// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
        /// </summary>
        [Output("sourceIdentifier")]
        public Output<string> SourceIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
        /// </summary>
        [Output("sourceOwner")]
        public Output<string> SourceOwner { get; private set; } = null!;

        /// <summary>
        /// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The rule monitors the tag key, only applies to rules created based on managed rules.
        /// </summary>
        [Output("tagKeyScope")]
        public Output<string?> TagKeyScope { get; private set; } = null!;

        /// <summary>
        /// The rule monitors the tag value, use with the `tag_key_scope` options. only applies to rules created based on managed rules.
        /// </summary>
        [Output("tagValueScope")]
        public Output<string?> TagValueScope { get; private set; } = null!;


        /// <summary>
        /// Create a AggregateConfigRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AggregateConfigRule(string name, AggregateConfigRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, args ?? new AggregateConfigRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AggregateConfigRule(string name, Input<string> id, AggregateConfigRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cfg/aggregateConfigRule:AggregateConfigRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AggregateConfigRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AggregateConfigRule Get(string name, Input<string> id, AggregateConfigRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AggregateConfigRule(name, id, state, options);
        }
    }

    public sealed class AggregateConfigRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("aggregateConfigRuleName", required: true)]
        public Input<string> AggregateConfigRuleName { get; set; } = null!;

        /// <summary>
        /// The Aggregator Id.
        /// </summary>
        [Input("aggregatorId", required: true)]
        public Input<string> AggregatorId { get; set; } = null!;

        /// <summary>
        /// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
        /// </summary>
        [Input("configRuleTriggerTypes", required: true)]
        public Input<string> ConfigRuleTriggerTypes { get; set; } = null!;

        /// <summary>
        /// The description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
        /// </summary>
        [Input("excludeResourceIdsScope")]
        public Input<string>? ExcludeResourceIdsScope { get; set; }

        [Input("inputParameters")]
        private InputMap<object>? _inputParameters;

        /// <summary>
        /// The settings map of the input parameters for the rule.
        /// </summary>
        public InputMap<object> InputParameters
        {
            get => _inputParameters ?? (_inputParameters = new InputMap<object>());
            set => _inputParameters = value;
        }

        /// <summary>
        /// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `config_rule_trigger_types` is `ScheduledNotification`.
        /// </summary>
        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        /// <summary>
        /// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Input("regionIdsScope")]
        public Input<string>? RegionIdsScope { get; set; }

        /// <summary>
        /// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Input("resourceGroupIdsScope")]
        public Input<string>? ResourceGroupIdsScope { get; set; }

        [Input("resourceTypesScopes", required: true)]
        private InputList<string>? _resourceTypesScopes;

        /// <summary>
        /// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        /// </summary>
        public InputList<string> ResourceTypesScopes
        {
            get => _resourceTypesScopes ?? (_resourceTypesScopes = new InputList<string>());
            set => _resourceTypesScopes = value;
        }

        /// <summary>
        /// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
        /// </summary>
        [Input("riskLevel", required: true)]
        public Input<int> RiskLevel { get; set; } = null!;

        /// <summary>
        /// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
        /// </summary>
        [Input("sourceIdentifier", required: true)]
        public Input<string> SourceIdentifier { get; set; } = null!;

        /// <summary>
        /// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
        /// </summary>
        [Input("sourceOwner", required: true)]
        public Input<string> SourceOwner { get; set; } = null!;

        /// <summary>
        /// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The rule monitors the tag key, only applies to rules created based on managed rules.
        /// </summary>
        [Input("tagKeyScope")]
        public Input<string>? TagKeyScope { get; set; }

        /// <summary>
        /// The rule monitors the tag value, use with the `tag_key_scope` options. only applies to rules created based on managed rules.
        /// </summary>
        [Input("tagValueScope")]
        public Input<string>? TagValueScope { get; set; }

        public AggregateConfigRuleArgs()
        {
        }
        public static new AggregateConfigRuleArgs Empty => new AggregateConfigRuleArgs();
    }

    public sealed class AggregateConfigRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("aggregateConfigRuleName")]
        public Input<string>? AggregateConfigRuleName { get; set; }

        /// <summary>
        /// The Aggregator Id.
        /// </summary>
        [Input("aggregatorId")]
        public Input<string>? AggregatorId { get; set; }

        /// <summary>
        /// (Available since v1.141.0) The rule ID of Aggregate Config Rule.
        /// </summary>
        [Input("configRuleId")]
        public Input<string>? ConfigRuleId { get; set; }

        /// <summary>
        /// The trigger type of the rule. Valid values: `ConfigurationItemChangeNotification`: The rule is triggered upon configuration changes. `ScheduledNotification`: The rule is triggered as scheduled.
        /// </summary>
        [Input("configRuleTriggerTypes")]
        public Input<string>? ConfigRuleTriggerTypes { get; set; }

        /// <summary>
        /// The description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The rule monitors excluded resource IDs, multiple of which are separated by commas, only applies to rules created based on managed rules, , custom rule this field is empty.
        /// </summary>
        [Input("excludeResourceIdsScope")]
        public Input<string>? ExcludeResourceIdsScope { get; set; }

        [Input("inputParameters")]
        private InputMap<object>? _inputParameters;

        /// <summary>
        /// The settings map of the input parameters for the rule.
        /// </summary>
        public InputMap<object> InputParameters
        {
            get => _inputParameters ?? (_inputParameters = new InputMap<object>());
            set => _inputParameters = value;
        }

        /// <summary>
        /// The frequency of the compliance evaluations. Valid values:  `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, `TwentyFour_Hours`. System default value is `TwentyFour_Hours` and valid when the `config_rule_trigger_types` is `ScheduledNotification`.
        /// </summary>
        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        /// <summary>
        /// The rule monitors region IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Input("regionIdsScope")]
        public Input<string>? RegionIdsScope { get; set; }

        /// <summary>
        /// The rule monitors resource group IDs, separated by commas, only applies to rules created based on managed rules.
        /// </summary>
        [Input("resourceGroupIdsScope")]
        public Input<string>? ResourceGroupIdsScope { get; set; }

        [Input("resourceTypesScopes")]
        private InputList<string>? _resourceTypesScopes;

        /// <summary>
        /// Resource types to be evaluated. [Alibaba Cloud services that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        /// </summary>
        public InputList<string> ResourceTypesScopes
        {
            get => _resourceTypesScopes ?? (_resourceTypesScopes = new InputList<string>());
            set => _resourceTypesScopes = value;
        }

        /// <summary>
        /// The risk level of the resources that are not compliant with the rule. Valid values:  `1`: critical `2`: warning `3`: info.
        /// </summary>
        [Input("riskLevel")]
        public Input<int>? RiskLevel { get; set; }

        /// <summary>
        /// The identifier of the rule. For a managed rule, the value is the identifier of the managed rule. For a custom rule, the value is the ARN of the custom rule. Using managed rules, refer to [List of Managed rules.](https://www.alibabacloud.com/help/en/doc-detail/127404.htm)
        /// </summary>
        [Input("sourceIdentifier")]
        public Input<string>? SourceIdentifier { get; set; }

        /// <summary>
        /// Specifies whether you or Alibaba Cloud owns and manages the rule. Valid values: `CUSTOM_FC`: The rule is a custom rule and you own the rule. `ALIYUN`: The rule is a managed rule and Alibaba Cloud owns the rule.
        /// </summary>
        [Input("sourceOwner")]
        public Input<string>? SourceOwner { get; set; }

        /// <summary>
        /// The rule status. The valid values: `ACTIVE`, `INACTIVE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The rule monitors the tag key, only applies to rules created based on managed rules.
        /// </summary>
        [Input("tagKeyScope")]
        public Input<string>? TagKeyScope { get; set; }

        /// <summary>
        /// The rule monitors the tag value, use with the `tag_key_scope` options. only applies to rules created based on managed rules.
        /// </summary>
        [Input("tagValueScope")]
        public Input<string>? TagValueScope { get; set; }

        public AggregateConfigRuleState()
        {
        }
        public static new AggregateConfigRuleState Empty => new AggregateConfigRuleState();
    }
}

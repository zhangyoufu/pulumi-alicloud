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
    /// Provides a Cloud Config Compliance Pack resource.
    /// 
    /// For information about Cloud Config Compliance Pack and how to use it, see [What is Compliance Pack](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createcompliancepack).
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
    ///     var name = config.Get("name") ?? "tf-example-config-name";
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var rule1 = new AliCloud.Cfg.Rule("rule1", new()
    ///     {
    ///         Description = name,
    ///         SourceOwner = "ALIYUN",
    ///         SourceIdentifier = "ram-user-ak-create-date-expired-check",
    ///         RiskLevel = 1,
    ///         MaximumExecutionFrequency = "TwentyFour_Hours",
    ///         RegionIdsScope = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///         ConfigRuleTriggerTypes = "ScheduledNotification",
    ///         ResourceTypesScopes = new[]
    ///         {
    ///             "ACS::RAM::User",
    ///         },
    ///         RuleName = "ciscompliancecheck_ram-user-ak-create-date-expired-check",
    ///         InputParameters = 
    ///         {
    ///             { "days", "90" },
    ///         },
    ///     });
    /// 
    ///     var rule2 = new AliCloud.Cfg.Rule("rule2", new()
    ///     {
    ///         Description = name,
    ///         SourceOwner = "ALIYUN",
    ///         SourceIdentifier = "adb-cluster-maintain-time-check",
    ///         RiskLevel = 2,
    ///         RegionIdsScope = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///         ConfigRuleTriggerTypes = "ScheduledNotification",
    ///         ResourceTypesScopes = new[]
    ///         {
    ///             "ACS::ADB::DBCluster",
    ///         },
    ///         RuleName = "governance-evaluation-adb-cluster-maintain-time-check",
    ///         InputParameters = 
    ///         {
    ///             { "maintainTimes", "02:00-04:00,06:00-08:00,12:00-13:00" },
    ///         },
    ///     });
    /// 
    ///     var defaultCompliancePack = new AliCloud.Cfg.CompliancePack("defaultCompliancePack", new()
    ///     {
    ///         CompliancePackName = name,
    ///         Description = "CloudGovernanceCenter evaluation",
    ///         RiskLevel = 2,
    ///         ConfigRuleIds = new[]
    ///         {
    ///             new AliCloud.Cfg.Inputs.CompliancePackConfigRuleIdArgs
    ///             {
    ///                 ConfigRuleId = rule1.Id,
    ///             },
    ///             new AliCloud.Cfg.Inputs.CompliancePackConfigRuleIdArgs
    ///             {
    ///                 ConfigRuleId = rule2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Config Compliance Pack can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cfg/compliancePack:CompliancePack example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cfg/compliancePack:CompliancePack")]
    public partial class CompliancePack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliance_pack_name` can be modified.
        /// </summary>
        [Output("compliancePackName")]
        public Output<string> CompliancePackName { get; private set; } = null!;

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Output("compliancePackTemplateId")]
        public Output<string?> CompliancePackTemplateId { get; private set; } = null!;

        /// <summary>
        /// A list of Config Rule IDs. See `config_rule_ids` below.
        /// </summary>
        [Output("configRuleIds")]
        public Output<ImmutableArray<Outputs.CompliancePackConfigRuleId>> ConfigRuleIds { get; private set; } = null!;

        /// <summary>
        /// A list of Config Rules. See `config_rules` below. **NOTE:** Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
        /// </summary>
        [Output("configRules")]
        public Output<ImmutableArray<Outputs.CompliancePackConfigRule>> ConfigRules { get; private set; } = null!;

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The Risk Level. Valid values:
        /// </summary>
        [Output("riskLevel")]
        public Output<int> RiskLevel { get; private set; } = null!;

        /// <summary>
        /// The status of the Compliance Pack.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CompliancePack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CompliancePack(string name, CompliancePackArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cfg/compliancePack:CompliancePack", name, args ?? new CompliancePackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CompliancePack(string name, Input<string> id, CompliancePackState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cfg/compliancePack:CompliancePack", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CompliancePack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CompliancePack Get(string name, Input<string> id, CompliancePackState? state = null, CustomResourceOptions? options = null)
        {
            return new CompliancePack(name, id, state, options);
        }
    }

    public sealed class CompliancePackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliance_pack_name` can be modified.
        /// </summary>
        [Input("compliancePackName", required: true)]
        public Input<string> CompliancePackName { get; set; } = null!;

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Input("compliancePackTemplateId")]
        public Input<string>? CompliancePackTemplateId { get; set; }

        [Input("configRuleIds")]
        private InputList<Inputs.CompliancePackConfigRuleIdArgs>? _configRuleIds;

        /// <summary>
        /// A list of Config Rule IDs. See `config_rule_ids` below.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleIdArgs> ConfigRuleIds
        {
            get => _configRuleIds ?? (_configRuleIds = new InputList<Inputs.CompliancePackConfigRuleIdArgs>());
            set => _configRuleIds = value;
        }

        [Input("configRules")]
        private InputList<Inputs.CompliancePackConfigRuleArgs>? _configRules;

        /// <summary>
        /// A list of Config Rules. See `config_rules` below. **NOTE:** Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
        /// </summary>
        [Obsolete(@"Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.")]
        public InputList<Inputs.CompliancePackConfigRuleArgs> ConfigRules
        {
            get => _configRules ?? (_configRules = new InputList<Inputs.CompliancePackConfigRuleArgs>());
            set => _configRules = value;
        }

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The Risk Level. Valid values:
        /// </summary>
        [Input("riskLevel", required: true)]
        public Input<int> RiskLevel { get; set; } = null!;

        public CompliancePackArgs()
        {
        }
        public static new CompliancePackArgs Empty => new CompliancePackArgs();
    }

    public sealed class CompliancePackState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Compliance Package Name. **NOTE:** From version 1.146.0, `compliance_pack_name` can be modified.
        /// </summary>
        [Input("compliancePackName")]
        public Input<string>? CompliancePackName { get; set; }

        /// <summary>
        /// Compliance Package Template Id.
        /// </summary>
        [Input("compliancePackTemplateId")]
        public Input<string>? CompliancePackTemplateId { get; set; }

        [Input("configRuleIds")]
        private InputList<Inputs.CompliancePackConfigRuleIdGetArgs>? _configRuleIds;

        /// <summary>
        /// A list of Config Rule IDs. See `config_rule_ids` below.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleIdGetArgs> ConfigRuleIds
        {
            get => _configRuleIds ?? (_configRuleIds = new InputList<Inputs.CompliancePackConfigRuleIdGetArgs>());
            set => _configRuleIds = value;
        }

        [Input("configRules")]
        private InputList<Inputs.CompliancePackConfigRuleGetArgs>? _configRules;

        /// <summary>
        /// A list of Config Rules. See `config_rules` below. **NOTE:** Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.
        /// </summary>
        [Obsolete(@"Field `config_rules` has been deprecated from provider version 1.141.0. New field `config_rule_ids` instead.")]
        public InputList<Inputs.CompliancePackConfigRuleGetArgs> ConfigRules
        {
            get => _configRules ?? (_configRules = new InputList<Inputs.CompliancePackConfigRuleGetArgs>());
            set => _configRules = value;
        }

        /// <summary>
        /// The Description of compliance pack.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Risk Level. Valid values:
        /// </summary>
        [Input("riskLevel")]
        public Input<int>? RiskLevel { get; set; }

        /// <summary>
        /// The status of the Compliance Pack.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CompliancePackState()
        {
        }
        public static new CompliancePackState Empty => new CompliancePackState();
    }
}

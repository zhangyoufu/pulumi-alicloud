// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    /// <summary>
    /// Provides a Dcdn Waf Rule resource.
    /// 
    /// For information about Dcdn Waf Rule and how to use it, see [What is Waf Rule](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchcreatedcdnwafrules).
    /// 
    /// &gt; **NOTE:** Available since v1.201.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var example = new AliCloud.Dcdn.WafPolicy("example", new()
    ///     {
    ///         DefenseScene = "waf_group",
    ///         PolicyName = $"{name}_{@default.Result}",
    ///         PolicyType = "custom",
    ///         Status = "on",
    ///     });
    /// 
    ///     var exampleWafRule = new AliCloud.Dcdn.WafRule("example", new()
    ///     {
    ///         PolicyId = example.Id,
    ///         RuleName = name,
    ///         Conditions = new[]
    ///         {
    ///             new AliCloud.Dcdn.Inputs.WafRuleConditionArgs
    ///             {
    ///                 Key = "URI",
    ///                 OpValue = "ne",
    ///                 Values = "/login.php",
    ///             },
    ///             new AliCloud.Dcdn.Inputs.WafRuleConditionArgs
    ///             {
    ///                 Key = "Header",
    ///                 SubKey = "a",
    ///                 OpValue = "eq",
    ///                 Values = "b",
    ///             },
    ///         },
    ///         Status = "on",
    ///         Action = "monitor",
    ///         RateLimit = new AliCloud.Dcdn.Inputs.WafRuleRateLimitArgs
    ///         {
    ///             Target = "IP",
    ///             Interval = 5,
    ///             Threshold = 5,
    ///             Ttl = 1800,
    ///             Status = new AliCloud.Dcdn.Inputs.WafRuleRateLimitStatusArgs
    ///             {
    ///                 Code = "200",
    ///                 Ratio = 60,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Dcdn Waf Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dcdn/wafRule:WafRule example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/wafRule:WafRule")]
    public partial class WafRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `custom_acl`.
        /// </summary>
        [Output("ccStatus")]
        public Output<string> CcStatus { get; private set; } = null!;

        /// <summary>
        /// The blocked regions in the Chinese mainland, separated by commas (,).
        /// </summary>
        [Output("cnRegionList")]
        public Output<string?> CnRegionList { get; private set; } = null!;

        /// <summary>
        /// Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `custom_acl` or `whitelist`.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.WafRuleCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
        /// </summary>
        [Output("defenseScene")]
        public Output<string> DefenseScene { get; private set; } = null!;

        /// <summary>
        /// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
        /// </summary>
        [Output("effect")]
        public Output<string?> Effect { get; private set; } = null!;

        /// <summary>
        /// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
        /// </summary>
        [Output("gmtModified")]
        public Output<string> GmtModified { get; private set; } = null!;

        /// <summary>
        /// Blocked regions outside the Chinese mainland, separated by commas (,).
        /// </summary>
        [Output("otherRegionList")]
        public Output<string?> OtherRegionList { get; private set; } = null!;

        /// <summary>
        /// The protection policy ID.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The rules of rate limiting. If you set `cc_status` to on, you must configure this parameter. See `rate_limit` below.
        /// </summary>
        [Output("rateLimit")]
        public Output<Outputs.WafRuleRateLimit?> RateLimit { get; private set; } = null!;

        /// <summary>
        /// The regular expression.e, when waf_group appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
        /// </summary>
        [Output("regularRules")]
        public Output<ImmutableArray<string>> RegularRules { get; private set; } = null!;

        /// <summary>
        /// Regular rule type, when waf_group appears in tags, this value can be filled in, optional values:["sqli", "xss", "code_exec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
        /// </summary>
        [Output("regularTypes")]
        public Output<ImmutableArray<string>> RegularTypes { get; private set; } = null!;

        /// <summary>
        /// Filter by IP address.
        /// </summary>
        [Output("remoteAddrs")]
        public Output<ImmutableArray<string>> RemoteAddrs { get; private set; } = null!;

        /// <summary>
        /// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `region_block`.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// The types of the protection policies.
        /// </summary>
        [Output("scenes")]
        public Output<ImmutableArray<string>> Scenes { get; private set; } = null!;

        /// <summary>
        /// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
        /// </summary>
        [Output("wafGroupIds")]
        public Output<string?> WafGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a WafRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WafRule(string name, WafRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafRule:WafRule", name, args ?? new WafRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WafRule(string name, Input<string> id, WafRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/wafRule:WafRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WafRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WafRule Get(string name, Input<string> id, WafRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new WafRule(name, id, state, options);
        }
    }

    public sealed class WafRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `custom_acl`.
        /// </summary>
        [Input("ccStatus")]
        public Input<string>? CcStatus { get; set; }

        /// <summary>
        /// The blocked regions in the Chinese mainland, separated by commas (,).
        /// </summary>
        [Input("cnRegionList")]
        public Input<string>? CnRegionList { get; set; }

        [Input("conditions")]
        private InputList<Inputs.WafRuleConditionArgs>? _conditions;

        /// <summary>
        /// Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `custom_acl` or `whitelist`.
        /// </summary>
        public InputList<Inputs.WafRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.WafRuleConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        /// <summary>
        /// Blocked regions outside the Chinese mainland, separated by commas (,).
        /// </summary>
        [Input("otherRegionList")]
        public Input<string>? OtherRegionList { get; set; }

        /// <summary>
        /// The protection policy ID.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// The rules of rate limiting. If you set `cc_status` to on, you must configure this parameter. See `rate_limit` below.
        /// </summary>
        [Input("rateLimit")]
        public Input<Inputs.WafRuleRateLimitArgs>? RateLimit { get; set; }

        [Input("regularRules")]
        private InputList<string>? _regularRules;

        /// <summary>
        /// The regular expression.e, when waf_group appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
        /// </summary>
        public InputList<string> RegularRules
        {
            get => _regularRules ?? (_regularRules = new InputList<string>());
            set => _regularRules = value;
        }

        [Input("regularTypes")]
        private InputList<string>? _regularTypes;

        /// <summary>
        /// Regular rule type, when waf_group appears in tags, this value can be filled in, optional values:["sqli", "xss", "code_exec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
        /// </summary>
        public InputList<string> RegularTypes
        {
            get => _regularTypes ?? (_regularTypes = new InputList<string>());
            set => _regularTypes = value;
        }

        [Input("remoteAddrs")]
        private InputList<string>? _remoteAddrs;

        /// <summary>
        /// Filter by IP address.
        /// </summary>
        public InputList<string> RemoteAddrs
        {
            get => _remoteAddrs ?? (_remoteAddrs = new InputList<string>());
            set => _remoteAddrs = value;
        }

        /// <summary>
        /// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `region_block`.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        [Input("scenes")]
        private InputList<string>? _scenes;

        /// <summary>
        /// The types of the protection policies.
        /// </summary>
        public InputList<string> Scenes
        {
            get => _scenes ?? (_scenes = new InputList<string>());
            set => _scenes = value;
        }

        /// <summary>
        /// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
        /// </summary>
        [Input("wafGroupIds")]
        public Input<string>? WafGroupIds { get; set; }

        public WafRuleArgs()
        {
        }
        public static new WafRuleArgs Empty => new WafRuleArgs();
    }

    public sealed class WafRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `custom_acl`.
        /// </summary>
        [Input("ccStatus")]
        public Input<string>? CcStatus { get; set; }

        /// <summary>
        /// The blocked regions in the Chinese mainland, separated by commas (,).
        /// </summary>
        [Input("cnRegionList")]
        public Input<string>? CnRegionList { get; set; }

        [Input("conditions")]
        private InputList<Inputs.WafRuleConditionGetArgs>? _conditions;

        /// <summary>
        /// Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `custom_acl` or `whitelist`.
        /// </summary>
        public InputList<Inputs.WafRuleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.WafRuleConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
        /// </summary>
        [Input("defenseScene")]
        public Input<string>? DefenseScene { get; set; }

        /// <summary>
        /// The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        /// <summary>
        /// Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
        /// </summary>
        [Input("gmtModified")]
        public Input<string>? GmtModified { get; set; }

        /// <summary>
        /// Blocked regions outside the Chinese mainland, separated by commas (,).
        /// </summary>
        [Input("otherRegionList")]
        public Input<string>? OtherRegionList { get; set; }

        /// <summary>
        /// The protection policy ID.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The rules of rate limiting. If you set `cc_status` to on, you must configure this parameter. See `rate_limit` below.
        /// </summary>
        [Input("rateLimit")]
        public Input<Inputs.WafRuleRateLimitGetArgs>? RateLimit { get; set; }

        [Input("regularRules")]
        private InputList<string>? _regularRules;

        /// <summary>
        /// The regular expression.e, when waf_group appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
        /// </summary>
        public InputList<string> RegularRules
        {
            get => _regularRules ?? (_regularRules = new InputList<string>());
            set => _regularRules = value;
        }

        [Input("regularTypes")]
        private InputList<string>? _regularTypes;

        /// <summary>
        /// Regular rule type, when waf_group appears in tags, this value can be filled in, optional values:["sqli", "xss", "code_exec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
        /// </summary>
        public InputList<string> RegularTypes
        {
            get => _regularTypes ?? (_regularTypes = new InputList<string>());
            set => _regularTypes = value;
        }

        [Input("remoteAddrs")]
        private InputList<string>? _remoteAddrs;

        /// <summary>
        /// Filter by IP address.
        /// </summary>
        public InputList<string> RemoteAddrs
        {
            get => _remoteAddrs ?? (_remoteAddrs = new InputList<string>());
            set => _remoteAddrs = value;
        }

        /// <summary>
        /// The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `region_block`.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        [Input("scenes")]
        private InputList<string>? _scenes;

        /// <summary>
        /// The types of the protection policies.
        /// </summary>
        public InputList<string> Scenes
        {
            get => _scenes ?? (_scenes = new InputList<string>());
            set => _scenes = value;
        }

        /// <summary>
        /// The status of the waf rule. Valid values: `on` and `off`. Default value: on.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
        /// </summary>
        [Input("wafGroupIds")]
        public Input<string>? WafGroupIds { get; set; }

        public WafRuleState()
        {
        }
        public static new WafRuleState Empty => new WafRuleState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sddp
{
    /// <summary>
    /// Provides a Data Security Center Rule resource.
    /// 
    /// For information about Data Security Center Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/data-security-center/latest/api-sddp-2019-01-03-createrule).
    /// 
    /// &gt; **NOTE:** Available since v1.132.0.
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
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf_example_name";
    ///     var @default = new AliCloud.Sddp.Rule("default", new()
    ///     {
    ///         Category = 0,
    ///         Content = "content",
    ///         RuleName = name,
    ///         RiskLevelId = "4",
    ///         ProductCode = "OSS",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Data Security Center Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:sddp/rule:Rule example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:sddp/rule:Rule")]
    public partial class Rule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sensitive Data Identification Rules for the Type of. Valid values:
        /// </summary>
        [Output("category")]
        public Output<int> Category { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Rules the Content.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The Content Classification.
        /// </summary>
        [Output("contentCategory")]
        public Output<string> ContentCategory { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Rules of Type. Valid values:
        /// </summary>
        [Output("customType")]
        public Output<int> CustomType { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification a Description of the Rule Information.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Request and Receive the Language of the Message Type. Valid values:
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;

        /// <summary>
        /// Product Code. Valid values: `OSS`,`RDS`,`ODPS`(MaxCompute).
        /// </summary>
        [Output("productCode")]
        public Output<string?> ProductCode { get; private set; } = null!;

        /// <summary>
        /// Product ID. Valid values:
        /// </summary>
        [Output("productId")]
        public Output<string?> ProductId { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Rules of Risk Level ID. Valid values:
        /// </summary>
        [Output("riskLevelId")]
        public Output<string?> RiskLevelId { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Name of the Rule.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Output("ruleType")]
        public Output<int?> RuleType { get; private set; } = null!;

        /// <summary>
        /// Triggered the Alarm Conditions.
        /// </summary>
        [Output("statExpress")]
        public Output<string?> StatExpress { get; private set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Rules Detection State of.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// The Target of rule.
        /// </summary>
        [Output("target")]
        public Output<string?> Target { get; private set; } = null!;

        /// <summary>
        /// The Level of Risk. Valid values:
        /// </summary>
        [Output("warnLevel")]
        public Output<int?> WarnLevel { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:sddp/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:sddp/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sensitive Data Identification Rules for the Type of. Valid values:
        /// </summary>
        [Input("category", required: true)]
        public Input<int> Category { get; set; } = null!;

        /// <summary>
        /// Sensitive Data Identification Rules the Content.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The Content Classification.
        /// </summary>
        [Input("contentCategory")]
        public Input<string>? ContentCategory { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Type. Valid values:
        /// </summary>
        [Input("customType")]
        public Input<int>? CustomType { get; set; }

        /// <summary>
        /// Sensitive Data Identification a Description of the Rule Information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Request and Receive the Language of the Message Type. Valid values:
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Product Code. Valid values: `OSS`,`RDS`,`ODPS`(MaxCompute).
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// Product ID. Valid values:
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Risk Level ID. Valid values:
        /// </summary>
        [Input("riskLevelId")]
        public Input<string>? RiskLevelId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Name of the Rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Input("ruleType")]
        public Input<int>? RuleType { get; set; }

        /// <summary>
        /// Triggered the Alarm Conditions.
        /// </summary>
        [Input("statExpress")]
        public Input<string>? StatExpress { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules Detection State of.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// The Target of rule.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The Level of Risk. Valid values:
        /// </summary>
        [Input("warnLevel")]
        public Input<int>? WarnLevel { get; set; }

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }

    public sealed class RuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sensitive Data Identification Rules for the Type of. Valid values:
        /// </summary>
        [Input("category")]
        public Input<int>? Category { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules the Content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The Content Classification.
        /// </summary>
        [Input("contentCategory")]
        public Input<string>? ContentCategory { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Type. Valid values:
        /// </summary>
        [Input("customType")]
        public Input<int>? CustomType { get; set; }

        /// <summary>
        /// Sensitive Data Identification a Description of the Rule Information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Request and Receive the Language of the Message Type. Valid values:
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Product Code. Valid values: `OSS`,`RDS`,`ODPS`(MaxCompute).
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// Product ID. Valid values:
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules of Risk Level ID. Valid values:
        /// </summary>
        [Input("riskLevelId")]
        public Input<string>? RiskLevelId { get; set; }

        /// <summary>
        /// Sensitive Data Identification Name of the Rule.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Input("ruleType")]
        public Input<int>? RuleType { get; set; }

        /// <summary>
        /// Triggered the Alarm Conditions.
        /// </summary>
        [Input("statExpress")]
        public Input<string>? StatExpress { get; set; }

        /// <summary>
        /// Sensitive Data Identification Rules Detection State of.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// The Target of rule.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The Level of Risk. Valid values:
        /// </summary>
        [Input("warnLevel")]
        public Input<int>? WarnLevel { get; set; }

        public RuleState()
        {
        }
        public static new RuleState Empty => new RuleState();
    }
}

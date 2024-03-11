// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    /// <summary>
    /// Provides a Threat Detection Baseline Strategy resource.
    /// 
    /// For information about Threat Detection Baseline Strategy and how to use it, see [What is Baseline Strategy](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-modifystrategy).
    /// 
    /// &gt; **NOTE:** Available since v1.195.0.
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
    ///     var @default = new AliCloud.ThreatDetection.BaselineStrategy("default", new()
    ///     {
    ///         BaselineStrategyName = "apispec",
    ///         CustomType = "custom",
    ///         CycleDays = 3,
    ///         EndTime = "08:00:00",
    ///         RiskSubTypeName = "hc_exploit_redis",
    ///         StartTime = "05:00:00",
    ///         TargetType = "groupId",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Threat Detection Baseline Strategy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:threatdetection/baselineStrategy:BaselineStrategy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:threatdetection/baselineStrategy:BaselineStrategy")]
    public partial class BaselineStrategy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the baseline check policy.
        /// </summary>
        [Output("baselineStrategyId")]
        public Output<string> BaselineStrategyId { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("baselineStrategyName")]
        public Output<string> BaselineStrategyName { get; private set; } = null!;

        /// <summary>
        /// The type of policy. Value:
        /// * **common**: standard policy
        /// * **custom**: custom policy
        /// </summary>
        [Output("customType")]
        public Output<string> CustomType { get; private set; } = null!;

        /// <summary>
        /// The detection period of the policy.
        /// </summary>
        [Output("cycleDays")]
        public Output<int> CycleDays { get; private set; } = null!;

        /// <summary>
        /// The detection period of the policy. Value:
        /// * **0**: 0:00~06:00
        /// * **6**: 6:00~12:00
        /// * **12**: 12:00~18:00
        /// * **18**: 18:00~24:00
        /// </summary>
        [Output("cycleStartTime")]
        public Output<int> CycleStartTime { get; private set; } = null!;

        /// <summary>
        /// The baseline check policy execution end time.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Detection item subtype.
        /// </summary>
        [Output("riskSubTypeName")]
        public Output<string> RiskSubTypeName { get; private set; } = null!;

        /// <summary>
        /// The baseline check policy start time.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The method of adding assets that take effect from the policy. Value:
        /// * **groupId**: Added by asset group.
        /// * **uuid**: Add by single asset.
        /// </summary>
        [Output("targetType")]
        public Output<string> TargetType { get; private set; } = null!;


        /// <summary>
        /// Create a BaselineStrategy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BaselineStrategy(string name, BaselineStrategyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/baselineStrategy:BaselineStrategy", name, args ?? new BaselineStrategyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BaselineStrategy(string name, Input<string> id, BaselineStrategyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/baselineStrategy:BaselineStrategy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BaselineStrategy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BaselineStrategy Get(string name, Input<string> id, BaselineStrategyState? state = null, CustomResourceOptions? options = null)
        {
            return new BaselineStrategy(name, id, state, options);
        }
    }

    public sealed class BaselineStrategyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("baselineStrategyName", required: true)]
        public Input<string> BaselineStrategyName { get; set; } = null!;

        /// <summary>
        /// The type of policy. Value:
        /// * **common**: standard policy
        /// * **custom**: custom policy
        /// </summary>
        [Input("customType", required: true)]
        public Input<string> CustomType { get; set; } = null!;

        /// <summary>
        /// The detection period of the policy.
        /// </summary>
        [Input("cycleDays", required: true)]
        public Input<int> CycleDays { get; set; } = null!;

        /// <summary>
        /// The detection period of the policy. Value:
        /// * **0**: 0:00~06:00
        /// * **6**: 6:00~12:00
        /// * **12**: 12:00~18:00
        /// * **18**: 18:00~24:00
        /// </summary>
        [Input("cycleStartTime")]
        public Input<int>? CycleStartTime { get; set; }

        /// <summary>
        /// The baseline check policy execution end time.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Detection item subtype.
        /// </summary>
        [Input("riskSubTypeName", required: true)]
        public Input<string> RiskSubTypeName { get; set; } = null!;

        /// <summary>
        /// The baseline check policy start time.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        /// <summary>
        /// The method of adding assets that take effect from the policy. Value:
        /// * **groupId**: Added by asset group.
        /// * **uuid**: Add by single asset.
        /// </summary>
        [Input("targetType", required: true)]
        public Input<string> TargetType { get; set; } = null!;

        public BaselineStrategyArgs()
        {
        }
        public static new BaselineStrategyArgs Empty => new BaselineStrategyArgs();
    }

    public sealed class BaselineStrategyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the baseline check policy.
        /// </summary>
        [Input("baselineStrategyId")]
        public Input<string>? BaselineStrategyId { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("baselineStrategyName")]
        public Input<string>? BaselineStrategyName { get; set; }

        /// <summary>
        /// The type of policy. Value:
        /// * **common**: standard policy
        /// * **custom**: custom policy
        /// </summary>
        [Input("customType")]
        public Input<string>? CustomType { get; set; }

        /// <summary>
        /// The detection period of the policy.
        /// </summary>
        [Input("cycleDays")]
        public Input<int>? CycleDays { get; set; }

        /// <summary>
        /// The detection period of the policy. Value:
        /// * **0**: 0:00~06:00
        /// * **6**: 6:00~12:00
        /// * **12**: 12:00~18:00
        /// * **18**: 18:00~24:00
        /// </summary>
        [Input("cycleStartTime")]
        public Input<int>? CycleStartTime { get; set; }

        /// <summary>
        /// The baseline check policy execution end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Detection item subtype.
        /// </summary>
        [Input("riskSubTypeName")]
        public Input<string>? RiskSubTypeName { get; set; }

        /// <summary>
        /// The baseline check policy start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The method of adding assets that take effect from the policy. Value:
        /// * **groupId**: Added by asset group.
        /// * **uuid**: Add by single asset.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        public BaselineStrategyState()
        {
        }
        public static new BaselineStrategyState Empty => new BaselineStrategyState();
    }
}

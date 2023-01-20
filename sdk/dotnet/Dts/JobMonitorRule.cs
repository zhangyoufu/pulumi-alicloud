// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dts
{
    /// <summary>
    /// Provides a DTS Job Monitor Rule resource.
    /// 
    /// For information about DTS Job Monitor Rule and how to use it, see [What is Job Monitor Rule](https://www.aliyun.com/product/dts).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.Dts.JobMonitorRule("example", new()
    ///     {
    ///         DtsJobId = "example_value",
    ///         Type = "delay",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DTS Job Monitor Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dts/jobMonitorRule:JobMonitorRule example &lt;dts_job_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dts/jobMonitorRule:JobMonitorRule")]
    public partial class JobMonitorRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Trigger delay alarm threshold, which is measured in seconds.
        /// </summary>
        [Output("delayRuleTime")]
        public Output<string> DelayRuleTime { get; private set; } = null!;

        /// <summary>
        /// Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        /// </summary>
        [Output("dtsJobId")]
        public Output<string> DtsJobId { get; private set; } = null!;

        /// <summary>
        /// The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        /// <summary>
        /// Whether to enable monitoring rules, valid values: `Y`, `N`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a JobMonitorRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobMonitorRule(string name, JobMonitorRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dts/jobMonitorRule:JobMonitorRule", name, args ?? new JobMonitorRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobMonitorRule(string name, Input<string> id, JobMonitorRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dts/jobMonitorRule:JobMonitorRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing JobMonitorRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobMonitorRule Get(string name, Input<string> id, JobMonitorRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new JobMonitorRule(name, id, state, options);
        }
    }

    public sealed class JobMonitorRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger delay alarm threshold, which is measured in seconds.
        /// </summary>
        [Input("delayRuleTime")]
        public Input<string>? DelayRuleTime { get; set; }

        /// <summary>
        /// Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        /// </summary>
        [Input("dtsJobId", required: true)]
        public Input<string> DtsJobId { get; set; } = null!;

        /// <summary>
        /// The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Whether to enable monitoring rules, valid values: `Y`, `N`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public JobMonitorRuleArgs()
        {
        }
        public static new JobMonitorRuleArgs Empty => new JobMonitorRuleArgs();
    }

    public sealed class JobMonitorRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger delay alarm threshold, which is measured in seconds.
        /// </summary>
        [Input("delayRuleTime")]
        public Input<string>? DelayRuleTime { get; set; }

        /// <summary>
        /// Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        /// </summary>
        [Input("dtsJobId")]
        public Input<string>? DtsJobId { get; set; }

        /// <summary>
        /// The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Whether to enable monitoring rules, valid values: `Y`, `N`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public JobMonitorRuleState()
        {
        }
        public static new JobMonitorRuleState Empty => new JobMonitorRuleState();
    }
}

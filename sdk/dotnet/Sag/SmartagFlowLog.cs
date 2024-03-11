// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sag
{
    /// <summary>
    /// Provides a Smartag Flow Log resource.
    /// 
    /// For information about Smartag Flow Log and how to use it, see [What is Flow Log](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createflowlog).
    /// 
    /// &gt; **NOTE:** Available since v1.168.0.
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
    ///     var example = new AliCloud.Sag.SmartagFlowLog("example", new()
    ///     {
    ///         NetflowServerIp = "192.168.0.2",
    ///         NetflowServerPort = 9995,
    ///         NetflowVersion = "V9",
    ///         OutputType = "netflow",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Smartag Flow Log can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:sag/smartagFlowLog:SmartagFlowLog example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:sag/smartagFlowLog:SmartagFlowLog")]
    public partial class SmartagFlowLog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time interval at which log data of active connections is collected. Valid values: `60` to `6000`. Default value: `300`. Unit: second.
        /// </summary>
        [Output("activeAging")]
        public Output<int> ActiveAging { get; private set; } = null!;

        /// <summary>
        /// The description of the flow log.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the flow log.
        /// </summary>
        [Output("flowLogName")]
        public Output<string?> FlowLogName { get; private set; } = null!;

        /// <summary>
        /// The time interval at which log data of inactive connections is connected. Valid values: `10` to `600`. Default value: `15`. Unit: second.
        /// </summary>
        [Output("inactiveAging")]
        public Output<int> InactiveAging { get; private set; } = null!;

        /// <summary>
        /// The Logstore in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Output("logstoreName")]
        public Output<string?> LogstoreName { get; private set; } = null!;

        /// <summary>
        /// The IP address of the NetFlow collector where the flow log is stored. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Output("netflowServerIp")]
        public Output<string?> NetflowServerIp { get; private set; } = null!;

        /// <summary>
        /// The port of the NetFlow collector. Default value: `9995`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Output("netflowServerPort")]
        public Output<int> NetflowServerPort { get; private set; } = null!;

        /// <summary>
        /// The NetFlow version. Default value: `V9`. Valid values: `V10`, `V5`, `V9`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Output("netflowVersion")]
        public Output<string> NetflowVersion { get; private set; } = null!;

        /// <summary>
        /// The location where the flow log is stored. Valid values:
        /// </summary>
        [Output("outputType")]
        public Output<string> OutputType { get; private set; } = null!;

        /// <summary>
        /// The project in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The ID of the region where Log Service is deployed. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Output("slsRegionId")]
        public Output<string?> SlsRegionId { get; private set; } = null!;

        /// <summary>
        /// The status of the flow log. Valid values:  `Active`: The flow log is enabled. `Inactive`: The flow log is disabled.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SmartagFlowLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmartagFlowLog(string name, SmartagFlowLogArgs args, CustomResourceOptions? options = null)
            : base("alicloud:sag/smartagFlowLog:SmartagFlowLog", name, args ?? new SmartagFlowLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmartagFlowLog(string name, Input<string> id, SmartagFlowLogState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:sag/smartagFlowLog:SmartagFlowLog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmartagFlowLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmartagFlowLog Get(string name, Input<string> id, SmartagFlowLogState? state = null, CustomResourceOptions? options = null)
        {
            return new SmartagFlowLog(name, id, state, options);
        }
    }

    public sealed class SmartagFlowLogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time interval at which log data of active connections is collected. Valid values: `60` to `6000`. Default value: `300`. Unit: second.
        /// </summary>
        [Input("activeAging")]
        public Input<int>? ActiveAging { get; set; }

        /// <summary>
        /// The description of the flow log.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the flow log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// The time interval at which log data of inactive connections is connected. Valid values: `10` to `600`. Default value: `15`. Unit: second.
        /// </summary>
        [Input("inactiveAging")]
        public Input<int>? InactiveAging { get; set; }

        /// <summary>
        /// The Logstore in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("logstoreName")]
        public Input<string>? LogstoreName { get; set; }

        /// <summary>
        /// The IP address of the NetFlow collector where the flow log is stored. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowServerIp")]
        public Input<string>? NetflowServerIp { get; set; }

        /// <summary>
        /// The port of the NetFlow collector. Default value: `9995`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowServerPort")]
        public Input<int>? NetflowServerPort { get; set; }

        /// <summary>
        /// The NetFlow version. Default value: `V9`. Valid values: `V10`, `V5`, `V9`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowVersion")]
        public Input<string>? NetflowVersion { get; set; }

        /// <summary>
        /// The location where the flow log is stored. Valid values:
        /// </summary>
        [Input("outputType", required: true)]
        public Input<string> OutputType { get; set; } = null!;

        /// <summary>
        /// The project in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The ID of the region where Log Service is deployed. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("slsRegionId")]
        public Input<string>? SlsRegionId { get; set; }

        /// <summary>
        /// The status of the flow log. Valid values:  `Active`: The flow log is enabled. `Inactive`: The flow log is disabled.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SmartagFlowLogArgs()
        {
        }
        public static new SmartagFlowLogArgs Empty => new SmartagFlowLogArgs();
    }

    public sealed class SmartagFlowLogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time interval at which log data of active connections is collected. Valid values: `60` to `6000`. Default value: `300`. Unit: second.
        /// </summary>
        [Input("activeAging")]
        public Input<int>? ActiveAging { get; set; }

        /// <summary>
        /// The description of the flow log.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the flow log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// The time interval at which log data of inactive connections is connected. Valid values: `10` to `600`. Default value: `15`. Unit: second.
        /// </summary>
        [Input("inactiveAging")]
        public Input<int>? InactiveAging { get; set; }

        /// <summary>
        /// The Logstore in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("logstoreName")]
        public Input<string>? LogstoreName { get; set; }

        /// <summary>
        /// The IP address of the NetFlow collector where the flow log is stored. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowServerIp")]
        public Input<string>? NetflowServerIp { get; set; }

        /// <summary>
        /// The port of the NetFlow collector. Default value: `9995`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowServerPort")]
        public Input<int>? NetflowServerPort { get; set; }

        /// <summary>
        /// The NetFlow version. Default value: `V9`. Valid values: `V10`, `V5`, `V9`. If `output_type` is set to `netflow` or `all`, this parameter is required.
        /// </summary>
        [Input("netflowVersion")]
        public Input<string>? NetflowVersion { get; set; }

        /// <summary>
        /// The location where the flow log is stored. Valid values:
        /// </summary>
        [Input("outputType")]
        public Input<string>? OutputType { get; set; }

        /// <summary>
        /// The project in Log Service. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The ID of the region where Log Service is deployed. If `output_type` is set to `sls` or `all`, this parameter is required.
        /// </summary>
        [Input("slsRegionId")]
        public Input<string>? SlsRegionId { get; set; }

        /// <summary>
        /// The status of the flow log. Valid values:  `Active`: The flow log is enabled. `Inactive`: The flow log is disabled.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SmartagFlowLogState()
        {
        }
        public static new SmartagFlowLogState Empty => new SmartagFlowLogState();
    }
}

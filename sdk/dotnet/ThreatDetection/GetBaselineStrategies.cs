// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    public static class GetBaselineStrategies
    {
        /// <summary>
        /// This data source provides Threat Detection Baseline Strategy available to the user.[What is Baseline Strategy](https://www.alibabacloud.com/help/zh/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifystrategy)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// </summary>
        public static Task<GetBaselineStrategiesResult> InvokeAsync(GetBaselineStrategiesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBaselineStrategiesResult>("alicloud:threatdetection/getBaselineStrategies:getBaselineStrategies", args ?? new GetBaselineStrategiesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Threat Detection Baseline Strategy available to the user.[What is Baseline Strategy](https://www.alibabacloud.com/help/zh/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifystrategy)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// </summary>
        public static Output<GetBaselineStrategiesResult> Invoke(GetBaselineStrategiesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBaselineStrategiesResult>("alicloud:threatdetection/getBaselineStrategies:getBaselineStrategies", args ?? new GetBaselineStrategiesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBaselineStrategiesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of policy. Value:-**common**: standard policy-**custom**: custom policy
        /// </summary>
        [Input("customType")]
        public string? CustomType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Baseline Strategy IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("strategyIds")]
        public string? StrategyIds { get; set; }

        public GetBaselineStrategiesArgs()
        {
        }
        public static new GetBaselineStrategiesArgs Empty => new GetBaselineStrategiesArgs();
    }

    public sealed class GetBaselineStrategiesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of policy. Value:-**common**: standard policy-**custom**: custom policy
        /// </summary>
        [Input("customType")]
        public Input<string>? CustomType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Baseline Strategy IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("strategyIds")]
        public Input<string>? StrategyIds { get; set; }

        public GetBaselineStrategiesInvokeArgs()
        {
        }
        public static new GetBaselineStrategiesInvokeArgs Empty => new GetBaselineStrategiesInvokeArgs();
    }


    [OutputType]
    public sealed class GetBaselineStrategiesResult
    {
        /// <summary>
        /// The type of policy. Value:
        /// * **common**: standard policy
        /// * **custom**: custom policy
        /// </summary>
        public readonly string? CustomType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Baseline Strategy IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Baseline Strategys.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetBaselineStrategiesStrategyResult> Strategies;
        public readonly string? StrategyIds;

        [OutputConstructor]
        private GetBaselineStrategiesResult(
            string? customType,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetBaselineStrategiesStrategyResult> strategies,

            string? strategyIds)
        {
            CustomType = customType;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Strategies = strategies;
            StrategyIds = strategyIds;
        }
    }
}

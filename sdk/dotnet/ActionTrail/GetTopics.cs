// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    public static class GetTopics
    {
        /// <summary>
        /// This data source provides a list of ALIKAFKA Topics in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.56.0+
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicsResult> InvokeAsync(GetTopicsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicsResult>("alicloud:actiontrail/getTopics:getTopics", args ?? new GetTopicsArgs(), options.WithVersion());
    }


    public sealed class GetTopicsArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the topic name. 
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetTopicsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of topic names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of topics. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicsTopicResult> Topics;

        [OutputConstructor]
        private GetTopicsResult(
            string id,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetTopicsTopicResult> topics)
        {
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Topics = topics;
        }
    }
}

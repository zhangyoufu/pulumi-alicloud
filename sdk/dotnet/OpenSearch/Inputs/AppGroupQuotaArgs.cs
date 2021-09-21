// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.OpenSearch.Inputs
{

    public sealed class AppGroupQuotaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Computing resources. Unit: LCU.
        /// </summary>
        [Input("computeResource", required: true)]
        public Input<int> ComputeResource { get; set; } = null!;

        /// <summary>
        /// Storage Size. Unit: GB.
        /// </summary>
        [Input("docSize", required: true)]
        public Input<int> DocSize { get; set; } = null!;

        /// <summary>
        /// Search request. Unit: times/second.
        /// </summary>
        [Input("qps")]
        public Input<int>? Qps { get; set; }

        /// <summary>
        /// Specification. Valid values: 
        /// * `opensearch.share.junior`: Entry-level.
        /// * `opensearch.share.common`: Shared universal.
        /// * `opensearch.share.compute`: Shared computing.
        /// * `opensearch.share.storage`: Shared storage type.
        /// * `opensearch.private.common`: Exclusive universal type.
        /// * `opensearch.private.compute`: Exclusive computing type.
        /// * `opensearch.private.storage`: Exclusive storage type
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        public AppGroupQuotaArgs()
        {
        }
    }
}

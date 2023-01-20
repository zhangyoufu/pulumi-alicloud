// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log.Inputs
{

    public sealed class StoreShardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The begin value of the shard range(MD5), included in the shard range.
        /// </summary>
        [Input("beginKey")]
        public Input<string>? BeginKey { get; set; }

        /// <summary>
        /// The end value of the shard range(MD5), not included in shard range.
        /// </summary>
        [Input("endKey")]
        public Input<string>? EndKey { get; set; }

        /// <summary>
        /// The ID of the shard.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Shard status, only two status of `readwrite` and `readonly`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public StoreShardArgs()
        {
        }
        public static new StoreShardArgs Empty => new StoreShardArgs();
    }
}

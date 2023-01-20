// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB.Inputs
{

    public sealed class ShardingInstanceConfigServerListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection address of the Config Server node.
        /// </summary>
        [Input("connectString")]
        public Input<string>? ConnectString { get; set; }

        /// <summary>
        /// The max connections of the Config Server node.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// The maximum IOPS of the Config Server node.
        /// </summary>
        [Input("maxIops")]
        public Input<int>? MaxIops { get; set; }

        /// <summary>
        /// Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        /// </summary>
        [Input("nodeClass")]
        public Input<string>? NodeClass { get; set; }

        /// <summary>
        /// The description of the Config Server node.
        /// </summary>
        [Input("nodeDescription")]
        public Input<string>? NodeDescription { get; set; }

        /// <summary>
        /// The ID of the Config Server node.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// - Custom storage space; value range: [10, 1,000]
        /// - 10-GB increments. Unit: GB.
        /// </summary>
        [Input("nodeStorage")]
        public Input<int>? NodeStorage { get; set; }

        /// <summary>
        /// The connection port of the Config Server node.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public ShardingInstanceConfigServerListGetArgs()
        {
        }
        public static new ShardingInstanceConfigServerListGetArgs Empty => new ShardingInstanceConfigServerListGetArgs();
    }
}

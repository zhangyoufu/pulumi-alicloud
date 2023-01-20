// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB.Inputs
{

    public sealed class InstanceReplicaSetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection address of the node.
        /// </summary>
        [Input("connectionDomain")]
        public Input<string>? ConnectionDomain { get; set; }

        /// <summary>
        /// The connection port of the node.
        /// </summary>
        [Input("connectionPort")]
        public Input<string>? ConnectionPort { get; set; }

        /// <summary>
        /// The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The role of the node. Valid values: `Primary`,`Secondary`.
        /// </summary>
        [Input("replicaSetRole")]
        public Input<string>? ReplicaSetRole { get; set; }

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Input("vpcCloudInstanceId")]
        public Input<string>? VpcCloudInstanceId { get; set; }

        /// <summary>
        /// The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public InstanceReplicaSetGetArgs()
        {
        }
        public static new InstanceReplicaSetGetArgs Empty => new InstanceReplicaSetGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Inputs
{

    public sealed class ScalingGroupAlbServerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of ALB server group.
        /// </summary>
        [Input("albServerGroupId")]
        public Input<string>? AlbServerGroupId { get; set; }

        /// <summary>
        /// The port number used by an ECS instance after Auto Scaling adds the ECS instance to ALB server group.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to ALB server group.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ScalingGroupAlbServerGroupArgs()
        {
        }
        public static new ScalingGroupAlbServerGroupArgs Empty => new ScalingGroupAlbServerGroupArgs();
    }
}

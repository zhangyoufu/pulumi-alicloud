// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Compute.Inputs
{

    public sealed class NestServiceInstanceOperationMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the imported service instance.
        /// </summary>
        [Input("operatedServiceInstanceId")]
        public Input<string>? OperatedServiceInstanceId { get; set; }

        /// <summary>
        /// The end time of O&amp;M.
        /// </summary>
        [Input("operationEndTime")]
        public Input<string>? OperationEndTime { get; set; }

        /// <summary>
        /// The start time of O&amp;M.
        /// </summary>
        [Input("operationStartTime")]
        public Input<string>? OperationStartTime { get; set; }

        /// <summary>
        /// The list of imported resources.
        /// </summary>
        [Input("resources")]
        public Input<string>? Resources { get; set; }

        public NestServiceInstanceOperationMetadataArgs()
        {
        }
        public static new NestServiceInstanceOperationMetadataArgs Empty => new NestServiceInstanceOperationMetadataArgs();
    }
}

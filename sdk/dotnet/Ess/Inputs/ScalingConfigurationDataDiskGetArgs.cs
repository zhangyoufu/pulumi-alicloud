// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Inputs
{

    public sealed class ScalingConfigurationDataDiskGetArgs : Pulumi.ResourceArgs
    {
        [Input("autoSnapshotPolicyId")]
        public Input<string>? AutoSnapshotPolicyId { get; set; }

        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("performanceLevel")]
        public Input<string>? PerformanceLevel { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ScalingConfigurationDataDiskGetArgs()
        {
        }
    }
}

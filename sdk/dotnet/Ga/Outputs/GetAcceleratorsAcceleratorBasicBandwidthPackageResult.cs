// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Outputs
{

    [OutputType]
    public sealed class GetAcceleratorsAcceleratorBasicBandwidthPackageResult
    {
        /// <summary>
        /// Bandwidth value of cross-domain acceleration package.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The bandwidth type of the basic bandwidth package.
        /// </summary>
        public readonly string BandwidthType;
        /// <summary>
        /// Instance ID of the cross-domain acceleration package.
        /// </summary>
        public readonly string InstanceId;

        [OutputConstructor]
        private GetAcceleratorsAcceleratorBasicBandwidthPackageResult(
            int bandwidth,

            string bandwidthType,

            string instanceId)
        {
            Bandwidth = bandwidth;
            BandwidthType = bandwidthType;
            InstanceId = instanceId;
        }
    }
}

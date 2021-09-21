// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class ListenerAccessLogTracingConfig
    {
        /// <summary>
        /// Xtrace Function. Value: `True` Or `False` . Default Value: `False`.
        /// </summary>
        public readonly bool? TracingEnabled;
        /// <summary>
        /// Xtrace Sampling Rate. Value: `1` to `10000`.
        /// </summary>
        public readonly int? TracingSample;
        /// <summary>
        /// Xtrace Type Value Is `Zipkin`.
        /// </summary>
        public readonly string? TracingType;

        [OutputConstructor]
        private ListenerAccessLogTracingConfig(
            bool? tracingEnabled,

            int? tracingSample,

            string? tracingType)
        {
            TracingEnabled = tracingEnabled;
            TracingSample = tracingSample;
            TracingType = tracingType;
        }
    }
}

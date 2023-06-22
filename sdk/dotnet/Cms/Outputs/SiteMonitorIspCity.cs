// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Outputs
{

    [OutputType]
    public sealed class SiteMonitorIspCity
    {
        /// <summary>
        /// The ID of the city.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// The ID of the carrier.
        /// </summary>
        public readonly string Isp;

        [OutputConstructor]
        private SiteMonitorIspCity(
            string city,

            string isp)
        {
            City = city;
            Isp = isp;
        }
    }
}

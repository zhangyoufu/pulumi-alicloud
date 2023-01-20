// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MaxCompute.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectIpWhiteListResult
    {
        /// <summary>
        /// Classic network IP white list.
        /// </summary>
        public readonly string IpList;
        /// <summary>
        /// VPC network whitelist.
        /// </summary>
        public readonly string VpcIpList;

        [OutputConstructor]
        private GetProjectsProjectIpWhiteListResult(
            string ipList,

            string vpcIpList)
        {
            IpList = ipList;
            VpcIpList = vpcIpList;
        }
    }
}

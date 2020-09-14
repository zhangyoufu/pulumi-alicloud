// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC.Outputs
{

    [OutputType]
    public sealed class ServiceNasConfigMountPoint
    {
        /// <summary>
        /// The local address where to mount your remote NAS directory.
        /// </summary>
        public readonly string MountDir;
        /// <summary>
        /// The address of the remote NAS directory.
        /// </summary>
        public readonly string ServerAddr;

        [OutputConstructor]
        private ServiceNasConfigMountPoint(
            string mountDir,

            string serverAddr)
        {
            MountDir = mountDir;
            ServerAddr = serverAddr;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh.Outputs
{

    [OutputType]
    public sealed class ServiceMeshMeshConfigControlPlaneLog
    {
        /// <summary>
        /// Whether to enable of the access logging. Valid values: `true` and `false`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The SLS Project of the access logging.
        /// </summary>
        public readonly string? Project;

        [OutputConstructor]
        private ServiceMeshMeshConfigControlPlaneLog(
            bool? enabled,

            string? project)
        {
            Enabled = enabled;
            Project = project;
        }
    }
}

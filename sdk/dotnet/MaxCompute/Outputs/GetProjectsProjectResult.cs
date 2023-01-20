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
    public sealed class GetProjectsProjectResult
    {
        public readonly string Comment;
        /// <summary>
        /// Default Computing Resource Group
        /// </summary>
        public readonly string DefaultQuota;
        /// <summary>
        /// Project ID. The value is the same as `project_name`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP whitelist
        /// </summary>
        public readonly Outputs.GetProjectsProjectIpWhiteListResult IpWhiteList;
        /// <summary>
        /// Project owner
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// Project base attributes
        /// </summary>
        public readonly Outputs.GetProjectsProjectPropertiesResult Properties;
        /// <summary>
        /// Security-related attributes
        /// </summary>
        public readonly Outputs.GetProjectsProjectSecurityPropertiesResult SecurityProperties;
        /// <summary>
        /// The status of the resource
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Project type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProjectsProjectResult(
            string comment,

            string defaultQuota,

            string id,

            Outputs.GetProjectsProjectIpWhiteListResult ipWhiteList,

            string owner,

            string projectName,

            Outputs.GetProjectsProjectPropertiesResult properties,

            Outputs.GetProjectsProjectSecurityPropertiesResult securityProperties,

            string status,

            string type)
        {
            Comment = comment;
            DefaultQuota = defaultQuota;
            Id = id;
            IpWhiteList = ipWhiteList;
            Owner = owner;
            ProjectName = projectName;
            Properties = properties;
            SecurityProperties = securityProperties;
            Status = status;
            Type = type;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Imm.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectResult
    {
        /// <summary>
        /// The billing type. **Note:** This parameter is deprecated from 2021-04-01.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// The maximum number of requests that can be processed per second. **Note:** This parameter is deprecated from 2021-04-01.
        /// </summary>
        public readonly int ComputeUnit;
        /// <summary>
        /// The creation time of project.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The service address of project.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The ID of project.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The modification time of project.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// -The name of project.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The service role authorized to the Intelligent Media Management service to access other cloud resources.
        /// </summary>
        public readonly string ServiceRole;
        /// <summary>
        /// The type of project.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProjectsProjectResult(
            string billingType,

            int computeUnit,

            string createTime,

            string endpoint,

            string id,

            string modifyTime,

            string project,

            string serviceRole,

            string type)
        {
            BillingType = billingType;
            ComputeUnit = computeUnit;
            CreateTime = createTime;
            Endpoint = endpoint;
            Id = id;
            ModifyTime = modifyTime;
            Project = project;
            ServiceRole = serviceRole;
            Type = type;
        }
    }
}

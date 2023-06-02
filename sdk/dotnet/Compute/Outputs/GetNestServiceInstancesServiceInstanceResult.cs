// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Compute.Outputs
{

    [OutputType]
    public sealed class GetNestServiceInstancesServiceInstanceResult
    {
        /// <summary>
        /// Whether the service instance has the O&amp;M function.
        /// </summary>
        public readonly bool EnableInstanceOps;
        /// <summary>
        /// The ID of the Service Instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the imported service instance.
        /// </summary>
        public readonly string OperatedServiceInstanceId;
        /// <summary>
        /// The end time of O&amp;M.
        /// </summary>
        public readonly string OperationEndTime;
        /// <summary>
        /// The start time of O&amp;M.
        /// </summary>
        public readonly string OperationStartTime;
        /// <summary>
        /// The parameters entered by the deployment service instance.
        /// </summary>
        public readonly string Parameters;
        /// <summary>
        /// The list of imported resources.
        /// </summary>
        public readonly string Resources;
        /// <summary>
        /// The ID of the Service Instance.
        /// </summary>
        public readonly string ServiceInstanceId;
        /// <summary>
        /// The name of the Service Instance.
        /// </summary>
        public readonly string ServiceInstanceName;
        /// <summary>
        /// Service details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNestServiceInstancesServiceInstanceServiceResult> Services;
        /// <summary>
        /// The source of the Service Instance.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The status of the Service Instance. Valid Values: `Created`, `Deploying`, `DeployedFailed`, `Deployed`, `Upgrading`, `Deleting`, `Deleted`, `DeletedFailed`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The name of the template.
        /// </summary>
        public readonly string TemplateName;

        [OutputConstructor]
        private GetNestServiceInstancesServiceInstanceResult(
            bool enableInstanceOps,

            string id,

            string operatedServiceInstanceId,

            string operationEndTime,

            string operationStartTime,

            string parameters,

            string resources,

            string serviceInstanceId,

            string serviceInstanceName,

            ImmutableArray<Outputs.GetNestServiceInstancesServiceInstanceServiceResult> services,

            string source,

            string status,

            ImmutableDictionary<string, object> tags,

            string templateName)
        {
            EnableInstanceOps = enableInstanceOps;
            Id = id;
            OperatedServiceInstanceId = operatedServiceInstanceId;
            OperationEndTime = operationEndTime;
            OperationStartTime = operationStartTime;
            Parameters = parameters;
            Resources = resources;
            ServiceInstanceId = serviceInstanceId;
            ServiceInstanceName = serviceInstanceName;
            Services = services;
            Source = source;
            Status = status;
            Tags = tags;
            TemplateName = templateName;
        }
    }
}

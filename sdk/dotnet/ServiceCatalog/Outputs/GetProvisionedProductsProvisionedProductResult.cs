// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class GetProvisionedProductsProvisionedProductResult
    {
        /// <summary>
        /// The creation time of the product instance
        /// </summary>
        public readonly string CreateTime;
        public readonly string Id;
        /// <summary>
        /// The ID of the last instance operation task
        /// </summary>
        public readonly string LastProvisioningTaskId;
        /// <summary>
        /// The ID of the last successful instance operation task
        /// </summary>
        public readonly string LastSuccessfulProvisioningTaskId;
        /// <summary>
        /// The ID of the last task
        /// </summary>
        public readonly string LastTaskId;
        public readonly ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductOutputResult> Outputs;
        /// <summary>
        /// The RAM entity ID of the owner
        /// </summary>
        public readonly string OwnerPrincipalId;
        /// <summary>
        /// The RAM entity type of the owner
        /// </summary>
        public readonly string OwnerPrincipalType;
        public readonly ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductParameterResult> Parameters;
        /// <summary>
        /// Product mix ID.&gt; When there is a default Startup option, there is no need to fill in the portfolio. When there is no default Startup option, you must fill in the portfolio.
        /// </summary>
        public readonly string PortfolioId;
        /// <summary>
        /// Product ID.
        /// </summary>
        public readonly string ProductId;
        /// <summary>
        /// The name of the product
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Product version ID.
        /// </summary>
        public readonly string ProductVersionId;
        /// <summary>
        /// The name of the product version
        /// </summary>
        public readonly string ProductVersionName;
        /// <summary>
        /// The ARN of the product instance
        /// </summary>
        public readonly string ProvisionedProductArn;
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        public readonly string ProvisionedProductId;
        /// <summary>
        /// The name of the instance.The length is 1~128 characters.
        /// </summary>
        public readonly string ProvisionedProductName;
        /// <summary>
        /// Instance type.The value is RosStack, which indicates the stack of Alibaba Cloud resource orchestration service (ROS).
        /// </summary>
        public readonly string ProvisionedProductType;
        /// <summary>
        /// The ID of the ROS stack
        /// </summary>
        public readonly string StackId;
        /// <summary>
        /// The ID of the region to which the resource stack of the Alibaba Cloud resource orchestration service (ROS) belongs.
        /// </summary>
        public readonly string StackRegionId;
        /// <summary>
        /// Instance status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The status message of the product instance
        /// </summary>
        public readonly string StatusMessage;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetProvisionedProductsProvisionedProductResult(
            string createTime,

            string id,

            string lastProvisioningTaskId,

            string lastSuccessfulProvisioningTaskId,

            string lastTaskId,

            ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductOutputResult> outputs,

            string ownerPrincipalId,

            string ownerPrincipalType,

            ImmutableArray<Outputs.GetProvisionedProductsProvisionedProductParameterResult> parameters,

            string portfolioId,

            string productId,

            string productName,

            string productVersionId,

            string productVersionName,

            string provisionedProductArn,

            string provisionedProductId,

            string provisionedProductName,

            string provisionedProductType,

            string stackId,

            string stackRegionId,

            string status,

            string statusMessage,

            ImmutableDictionary<string, string>? tags)
        {
            CreateTime = createTime;
            Id = id;
            LastProvisioningTaskId = lastProvisioningTaskId;
            LastSuccessfulProvisioningTaskId = lastSuccessfulProvisioningTaskId;
            LastTaskId = lastTaskId;
            Outputs = outputs;
            OwnerPrincipalId = ownerPrincipalId;
            OwnerPrincipalType = ownerPrincipalType;
            Parameters = parameters;
            PortfolioId = portfolioId;
            ProductId = productId;
            ProductName = productName;
            ProductVersionId = productVersionId;
            ProductVersionName = productVersionName;
            ProvisionedProductArn = provisionedProductArn;
            ProvisionedProductId = provisionedProductId;
            ProvisionedProductName = provisionedProductName;
            ProvisionedProductType = provisionedProductType;
            StackId = stackId;
            StackRegionId = stackRegionId;
            Status = status;
            StatusMessage = statusMessage;
            Tags = tags;
        }
    }
}

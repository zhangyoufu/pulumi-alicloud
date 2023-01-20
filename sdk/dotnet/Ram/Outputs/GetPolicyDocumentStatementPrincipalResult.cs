// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram.Outputs
{

    [OutputType]
    public sealed class GetPolicyDocumentStatementPrincipalResult
    {
        /// <summary>
        /// The trusted entity. Valid values: `RAM`, `Service` and `Federated`.
        /// </summary>
        public readonly string Entity;
        /// <summary>
        /// The identifiers of the principal.
        /// </summary>
        public readonly ImmutableArray<string> Identifiers;

        [OutputConstructor]
        private GetPolicyDocumentStatementPrincipalResult(
            string entity,

            ImmutableArray<string> identifiers)
        {
            Entity = entity;
            Identifiers = identifiers;
        }
    }
}

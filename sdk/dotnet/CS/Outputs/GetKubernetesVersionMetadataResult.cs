// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Outputs
{

    [OutputType]
    public sealed class GetKubernetesVersionMetadataResult
    {
        /// <summary>
        /// The list of supported runtime.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubernetesVersionMetadataRuntimeResult> Runtimes;
        /// <summary>
        /// The runtime version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetKubernetesVersionMetadataResult(
            ImmutableArray<Outputs.GetKubernetesVersionMetadataRuntimeResult> runtimes,

            string version)
        {
            Runtimes = runtimes;
            Version = version;
        }
    }
}

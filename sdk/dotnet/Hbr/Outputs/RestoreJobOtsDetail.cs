// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Outputs
{

    [OutputType]
    public sealed class RestoreJobOtsDetail
    {
        /// <summary>
        /// Whether to overwrite the existing table storage recovery task. Valid values: `true`, `false`.
        /// </summary>
        public readonly bool? OverwriteExisting;

        [OutputConstructor]
        private RestoreJobOtsDetail(bool? overwriteExisting)
        {
            OverwriteExisting = overwriteExisting;
        }
    }
}

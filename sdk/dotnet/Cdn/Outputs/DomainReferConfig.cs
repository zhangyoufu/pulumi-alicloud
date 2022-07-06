// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainReferConfig
    {
        public readonly string? AllowEmpty;
        public readonly ImmutableArray<string> ReferLists;
        public readonly string? ReferType;

        [OutputConstructor]
        private DomainReferConfig(
            string? allowEmpty,

            ImmutableArray<string> referLists,

            string? referType)
        {
            AllowEmpty = allowEmpty;
            ReferLists = referLists;
            ReferType = referType;
        }
    }
}

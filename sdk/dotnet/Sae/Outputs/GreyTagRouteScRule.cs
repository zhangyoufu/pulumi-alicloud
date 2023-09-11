// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Outputs
{

    [OutputType]
    public sealed class GreyTagRouteScRule
    {
        /// <summary>
        /// The conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// A list of conditions items.See `items` below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GreyTagRouteScRuleItem> Items;
        /// <summary>
        /// The path corresponding to the grayscale rule.
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private GreyTagRouteScRule(
            string? condition,

            ImmutableArray<Outputs.GreyTagRouteScRuleItem> items,

            string? path)
        {
            Condition = condition;
            Items = items;
            Path = path;
        }
    }
}

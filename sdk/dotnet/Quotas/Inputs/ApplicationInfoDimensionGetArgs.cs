// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Quotas.Inputs
{

    public sealed class ApplicationInfoDimensionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationInfoDimensionGetArgs()
        {
        }
        public static new ApplicationInfoDimensionGetArgs Empty => new ApplicationInfoDimensionGetArgs();
    }
}

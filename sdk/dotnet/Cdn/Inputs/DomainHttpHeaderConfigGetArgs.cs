// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn.Inputs
{

    public sealed class DomainHttpHeaderConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerId")]
        public Input<string>? HeaderId { get; set; }

        [Input("headerKey", required: true)]
        public Input<string> HeaderKey { get; set; } = null!;

        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        public DomainHttpHeaderConfigGetArgs()
        {
        }
        public static new DomainHttpHeaderConfigGetArgs Empty => new DomainHttpHeaderConfigGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class SlsGroupSlsGroupConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Log Store.
        /// </summary>
        [Input("slsLogstore", required: true)]
        public Input<string> SlsLogstore { get; set; } = null!;

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Input("slsProject", required: true)]
        public Input<string> SlsProject { get; set; } = null!;

        /// <summary>
        /// The Sls Region.
        /// </summary>
        [Input("slsRegion", required: true)]
        public Input<string> SlsRegion { get; set; } = null!;

        /// <summary>
        /// The ID of the Sls User.
        /// </summary>
        [Input("slsUserId")]
        public Input<string>? SlsUserId { get; set; }

        public SlsGroupSlsGroupConfigGetArgs()
        {
        }
        public static new SlsGroupSlsGroupConfigGetArgs Empty => new SlsGroupSlsGroupConfigGetArgs();
    }
}

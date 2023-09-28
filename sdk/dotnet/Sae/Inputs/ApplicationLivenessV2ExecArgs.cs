// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationLivenessV2ExecArgs : global::Pulumi.ResourceArgs
    {
        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        public ApplicationLivenessV2ExecArgs()
        {
        }
        public static new ApplicationLivenessV2ExecArgs Empty => new ApplicationLivenessV2ExecArgs();
    }
}

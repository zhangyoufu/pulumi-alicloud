// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sls.Inputs
{

    public sealed class AlertConfigurationTemplateConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Template Annotations.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// Template Language.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Template ID.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        [Input("tokens")]
        private InputMap<string>? _tokens;

        /// <summary>
        /// Template Variables.
        /// </summary>
        public InputMap<string> Tokens
        {
            get => _tokens ?? (_tokens = new InputMap<string>());
            set => _tokens = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Template Version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AlertConfigurationTemplateConfigurationGetArgs()
        {
        }
        public static new AlertConfigurationTemplateConfigurationGetArgs Empty => new AlertConfigurationTemplateConfigurationGetArgs();
    }
}

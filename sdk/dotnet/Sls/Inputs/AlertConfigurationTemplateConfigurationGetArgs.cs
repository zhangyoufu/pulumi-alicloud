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
        private InputMap<object>? _annotations;
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        [Input("lang")]
        public Input<string>? Lang { get; set; }

        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        [Input("tokens")]
        private InputMap<object>? _tokens;
        public InputMap<object> Tokens
        {
            get => _tokens ?? (_tokens = new InputMap<object>());
            set => _tokens = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public AlertConfigurationTemplateConfigurationGetArgs()
        {
        }
        public static new AlertConfigurationTemplateConfigurationGetArgs Empty => new AlertConfigurationTemplateConfigurationGetArgs();
    }
}

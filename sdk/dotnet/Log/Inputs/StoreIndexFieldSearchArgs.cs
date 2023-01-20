// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log.Inputs
{

    public sealed class StoreIndexFieldSearchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// Whether to enable field analytics. Default to true.
        /// </summary>
        [Input("enableAnalytics")]
        public Input<bool>? EnableAnalytics { get; set; }

        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        [Input("jsonKeys")]
        private InputList<Inputs.StoreIndexFieldSearchJsonKeyArgs>? _jsonKeys;

        /// <summary>
        /// Use nested index when type is json
        /// </summary>
        public InputList<Inputs.StoreIndexFieldSearchJsonKeyArgs> JsonKeys
        {
            get => _jsonKeys ?? (_jsonKeys = new InputList<Inputs.StoreIndexFieldSearchJsonKeyArgs>());
            set => _jsonKeys = value;
        }

        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StoreIndexFieldSearchArgs()
        {
        }
        public static new StoreIndexFieldSearchArgs Empty => new StoreIndexFieldSearchArgs();
    }
}

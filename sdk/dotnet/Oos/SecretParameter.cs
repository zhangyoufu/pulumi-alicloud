// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    /// <summary>
    /// Provides a OOS Secret Parameter resource.
    /// 
    /// For information about OOS Secret Parameter and how to use it, see [What is Secret Parameter](https://www.alibabacloud.com/help/en/doc-detail/183418.html).
    /// 
    /// &gt; **NOTE:** Available in v1.147.0+.
    /// 
    /// ## Import
    /// 
    /// OOS Secret Parameter can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:oos/secretParameter:SecretParameter example &lt;secret_parameter_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/secretParameter:SecretParameter")]
    public partial class SecretParameter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
        /// </summary>
        [Output("constraints")]
        public Output<string?> Constraints { get; private set; } = null!;

        /// <summary>
        /// The description of the encryption parameter. The description must be `1` to `200` characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
        /// </summary>
        [Output("keyId")]
        public Output<string?> KeyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Resource Group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
        /// </summary>
        [Output("secretParameterName")]
        public Output<string> SecretParameterName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The data type of the encryption parameter. Valid values: `Secret`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a SecretParameter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretParameter(string name, SecretParameterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oos/secretParameter:SecretParameter", name, args ?? new SecretParameterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretParameter(string name, Input<string> id, SecretParameterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/secretParameter:SecretParameter", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "value",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretParameter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretParameter Get(string name, Input<string> id, SecretParameterState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretParameter(name, id, state, options);
        }
    }

    public sealed class SecretParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
        /// </summary>
        [Input("constraints")]
        public Input<string>? Constraints { get; set; }

        /// <summary>
        /// The description of the encryption parameter. The description must be `1` to `200` characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The ID of the Resource Group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
        /// </summary>
        [Input("secretParameterName", required: true)]
        public Input<string> SecretParameterName { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The data type of the encryption parameter. Valid values: `Secret`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value", required: true)]
        private Input<string>? _value;

        /// <summary>
        /// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
        /// </summary>
        public Input<string>? Value
        {
            get => _value;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _value = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretParameterArgs()
        {
        }
        public static new SecretParameterArgs Empty => new SecretParameterArgs();
    }

    public sealed class SecretParameterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
        /// </summary>
        [Input("constraints")]
        public Input<string>? Constraints { get; set; }

        /// <summary>
        /// The description of the encryption parameter. The description must be `1` to `200` characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Customer Master Key (CMK) of Key Management Service (KMS) that is used to encrypt the parameter.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The ID of the Resource Group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The name of the encryption parameter.  The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
        /// </summary>
        [Input("secretParameterName")]
        public Input<string>? SecretParameterName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The data type of the encryption parameter. Valid values: `Secret`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value")]
        private Input<string>? _value;

        /// <summary>
        /// The value of the encryption parameter. The value must be `1` to `4096` characters in length.
        /// </summary>
        public Input<string>? Value
        {
            get => _value;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _value = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretParameterState()
        {
        }
        public static new SecretParameterState Empty => new SecretParameterState();
    }
}

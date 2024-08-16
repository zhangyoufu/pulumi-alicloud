// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// &gt; **DEPRECATED:** This resource has been renamed to alicloud.ecs.EcsKeyPair from version 1.121.0.
    /// 
    /// Provides a key pair resource.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basic = new AliCloud.Ecs.KeyPair("basic", new()
    ///     {
    ///         KeyName = "terraform-test-key-pair",
    ///     });
    /// 
    ///     // Using name prefix to build key pair
    ///     var prefix = new AliCloud.Ecs.KeyPair("prefix", new()
    ///     {
    ///         KeyNamePrefix = "terraform-test-key-pair-prefix",
    ///     });
    /// 
    ///     // Import an existing public key to build a alicloud key pair
    ///     var publickey = new AliCloud.Ecs.KeyPair("publickey", new()
    ///     {
    ///         KeyName = "my_public_key",
    ///         PublicKey = "ssh-rsa AAAAB3Nza12345678qwertyuudsfsg",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Key pair can be imported using the name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/keyPair:KeyPair example my_public_key
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/keyPair:KeyPair")]
    public partial class KeyPair : global::Pulumi.CustomResource
    {
        [Output("fingerPrint")]
        public Output<string> FingerPrint { get; private set; } = null!;

        /// <summary>
        /// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        /// </summary>
        [Output("keyFile")]
        public Output<string?> KeyFile { get; private set; } = null!;

        /// <summary>
        /// The key pair's name. It is the only in one Alicloud account.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;

        [Output("keyNamePrefix")]
        public Output<string?> KeyNamePrefix { get; private set; } = null!;

        [Output("keyPairName")]
        public Output<string> KeyPairName { get; private set; } = null!;

        /// <summary>
        /// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        /// </summary>
        [Output("publicKey")]
        public Output<string?> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the key pair belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a KeyPair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyPair(string name, KeyPairArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/keyPair:KeyPair", name, args ?? new KeyPairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyPair(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/keyPair:KeyPair", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KeyPair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyPair Get(string name, Input<string> id, KeyPairState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyPair(name, id, state, options);
        }
    }

    public sealed class KeyPairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        /// <summary>
        /// The key pair's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("keyNamePrefix")]
        public Input<string>? KeyNamePrefix { get; set; }

        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The Id of resource group which the key pair belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KeyPairArgs()
        {
        }
        public static new KeyPairArgs Empty => new KeyPairArgs();
    }

    public sealed class KeyPairState : global::Pulumi.ResourceArgs
    {
        [Input("fingerPrint")]
        public Input<string>? FingerPrint { get; set; }

        /// <summary>
        /// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        /// <summary>
        /// The key pair's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("keyNamePrefix")]
        public Input<string>? KeyNamePrefix { get; set; }

        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The Id of resource group which the key pair belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KeyPairState()
        {
        }
        public static new KeyPairState Empty => new KeyPairState();
    }
}

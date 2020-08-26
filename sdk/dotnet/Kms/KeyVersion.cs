// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms
{
    /// <summary>
    /// Provides a Alikms Key Version resource. For information about Alikms Key Version and how to use it, see [What is Resource Alikms Key Version](https://www.alibabacloud.com/help/doc-detail/133838.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.85.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @this = new AliCloud.Kms.Key("this", new AliCloud.Kms.KeyArgs
    ///         {
    ///         });
    ///         var keyversion = new AliCloud.Kms.KeyVersion("keyversion", new AliCloud.Kms.KeyVersionArgs
    ///         {
    ///             KeyId = @this.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class KeyVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time (UTC time) when the Alikms key version was created.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The id of the master key (CMK).
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The id of the Alikms key version.
        /// </summary>
        [Output("keyVersionId")]
        public Output<string> KeyVersionId { get; private set; } = null!;


        /// <summary>
        /// Create a KeyVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyVersion(string name, KeyVersionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:kms/keyVersion:KeyVersion", name, args ?? new KeyVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyVersion(string name, Input<string> id, KeyVersionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:kms/keyVersion:KeyVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyVersion Get(string name, Input<string> id, KeyVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyVersion(name, id, state, options);
        }
    }

    public sealed class KeyVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the master key (CMK).
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public KeyVersionArgs()
        {
        }
    }

    public sealed class KeyVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time (UTC time) when the Alikms key version was created.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The id of the master key (CMK).
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The id of the Alikms key version.
        /// </summary>
        [Input("keyVersionId")]
        public Input<string>? KeyVersionId { get; set; }

        public KeyVersionState()
        {
        }
    }
}

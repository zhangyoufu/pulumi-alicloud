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
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new AliCloud.Kms.Key("this");
    /// 
    ///     var keyversion = new AliCloud.Kms.KeyVersion("keyversion", new()
    ///     {
    ///         KeyId = @this.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Alikms key version can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:kms/keyVersion:KeyVersion example 72da539a-2fa8-4f2d-b854-*****	
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:kms/keyVersion:KeyVersion")]
    public partial class KeyVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the master key (CMK).
        /// 
        /// &gt; **NOTE:** The minimum interval for creating a Alikms key version is 7 days.
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

    public sealed class KeyVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the master key (CMK).
        /// 
        /// &gt; **NOTE:** The minimum interval for creating a Alikms key version is 7 days.
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public KeyVersionArgs()
        {
        }
        public static new KeyVersionArgs Empty => new KeyVersionArgs();
    }

    public sealed class KeyVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the master key (CMK).
        /// 
        /// &gt; **NOTE:** The minimum interval for creating a Alikms key version is 7 days.
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
        public static new KeyVersionState Empty => new KeyVersionState();
    }
}

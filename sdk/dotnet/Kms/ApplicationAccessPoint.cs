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
    /// Provides a KMS Application Access Point resource. An application access point (AAP) is used to implement fine-grained access control for Key Management Service (KMS) resources. An application can access a KMS instance only after an AAP is created for the application. .
    /// 
    /// For information about KMS Application Access Point and how to use it, see [What is Application Access Point](https://www.alibabacloud.com/help/zh/key-management-service/latest/api-createapplicationaccesspoint).
    /// 
    /// &gt; **NOTE:** Available since v1.210.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = new AliCloud.Kms.ApplicationAccessPoint("default", new()
    ///     {
    ///         Description = "example aap",
    ///         ApplicationAccessPointName = name,
    ///         Policies = new[]
    ///         {
    ///             "abc",
    ///             "efg",
    ///             "hfc",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// KMS Application Access Point can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:kms/applicationAccessPoint:ApplicationAccessPoint example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:kms/applicationAccessPoint:ApplicationAccessPoint")]
    public partial class ApplicationAccessPoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application Access Point Name.
        /// </summary>
        [Output("applicationAccessPointName")]
        public Output<string> ApplicationAccessPointName { get; private set; } = null!;

        /// <summary>
        /// Description .
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The policies that have bound to the Application Access Point (AAP).
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationAccessPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationAccessPoint(string name, ApplicationAccessPointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:kms/applicationAccessPoint:ApplicationAccessPoint", name, args ?? new ApplicationAccessPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationAccessPoint(string name, Input<string> id, ApplicationAccessPointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:kms/applicationAccessPoint:ApplicationAccessPoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationAccessPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationAccessPoint Get(string name, Input<string> id, ApplicationAccessPointState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationAccessPoint(name, id, state, options);
        }
    }

    public sealed class ApplicationAccessPointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Access Point Name.
        /// </summary>
        [Input("applicationAccessPointName", required: true)]
        public Input<string> ApplicationAccessPointName { get; set; } = null!;

        /// <summary>
        /// Description .
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("policies", required: true)]
        private InputList<string>? _policies;

        /// <summary>
        /// The policies that have bound to the Application Access Point (AAP).
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public ApplicationAccessPointArgs()
        {
        }
        public static new ApplicationAccessPointArgs Empty => new ApplicationAccessPointArgs();
    }

    public sealed class ApplicationAccessPointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Access Point Name.
        /// </summary>
        [Input("applicationAccessPointName")]
        public Input<string>? ApplicationAccessPointName { get; set; }

        /// <summary>
        /// Description .
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The policies that have bound to the Application Access Point (AAP).
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public ApplicationAccessPointState()
        {
        }
        public static new ApplicationAccessPointState Empty => new ApplicationAccessPointState();
    }
}

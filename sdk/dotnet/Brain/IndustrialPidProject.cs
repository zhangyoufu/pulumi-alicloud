// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Brain
{
    /// <summary>
    /// Provides a Brain Industrial Pid Project resource.
    /// 
    /// &gt; **NOTE:** Available since v1.113.0.
    /// 
    /// &gt; **DEPRECATED:**  This resource has been deprecated from version `1.222.0`.
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
    ///     var example = new AliCloud.Brain.IndustrialPidProject("example", new()
    ///     {
    ///         PidOrganizationId = "3e74e684-cbb5-xxxx",
    ///         PidProjectName = "tf-testAcc",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Brain Industrial Pid Project can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:brain/industrialPidProject:IndustrialPidProject example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:brain/industrialPidProject:IndustrialPidProject")]
    public partial class IndustrialPidProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of Pid Organization.
        /// </summary>
        [Output("pidOrganizationId")]
        public Output<string> PidOrganizationId { get; private set; } = null!;

        /// <summary>
        /// The description of Pid Project.
        /// </summary>
        [Output("pidProjectDesc")]
        public Output<string?> PidProjectDesc { get; private set; } = null!;

        /// <summary>
        /// The name of Pid Project.
        /// </summary>
        [Output("pidProjectName")]
        public Output<string> PidProjectName { get; private set; } = null!;


        /// <summary>
        /// Create a IndustrialPidProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IndustrialPidProject(string name, IndustrialPidProjectArgs args, CustomResourceOptions? options = null)
            : base("alicloud:brain/industrialPidProject:IndustrialPidProject", name, args ?? new IndustrialPidProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IndustrialPidProject(string name, Input<string> id, IndustrialPidProjectState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:brain/industrialPidProject:IndustrialPidProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IndustrialPidProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IndustrialPidProject Get(string name, Input<string> id, IndustrialPidProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new IndustrialPidProject(name, id, state, options);
        }
    }

    public sealed class IndustrialPidProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Pid Organization.
        /// </summary>
        [Input("pidOrganizationId", required: true)]
        public Input<string> PidOrganizationId { get; set; } = null!;

        /// <summary>
        /// The description of Pid Project.
        /// </summary>
        [Input("pidProjectDesc")]
        public Input<string>? PidProjectDesc { get; set; }

        /// <summary>
        /// The name of Pid Project.
        /// </summary>
        [Input("pidProjectName", required: true)]
        public Input<string> PidProjectName { get; set; } = null!;

        public IndustrialPidProjectArgs()
        {
        }
        public static new IndustrialPidProjectArgs Empty => new IndustrialPidProjectArgs();
    }

    public sealed class IndustrialPidProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Pid Organization.
        /// </summary>
        [Input("pidOrganizationId")]
        public Input<string>? PidOrganizationId { get; set; }

        /// <summary>
        /// The description of Pid Project.
        /// </summary>
        [Input("pidProjectDesc")]
        public Input<string>? PidProjectDesc { get; set; }

        /// <summary>
        /// The name of Pid Project.
        /// </summary>
        [Input("pidProjectName")]
        public Input<string>? PidProjectName { get; set; }

        public IndustrialPidProjectState()
        {
        }
        public static new IndustrialPidProjectState Empty => new IndustrialPidProjectState();
    }
}

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
    /// Provides a OOS Patch Baseline resource.
    /// 
    /// For information about OOS Patch Baseline and how to use it, see [What is Patch Baseline](https://www.alibabacloud.com/help/en/doc-detail/268700.html).
    /// 
    /// &gt; **NOTE:** Available in v1.146.0+.
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
    ///     var example = new AliCloud.Oos.PatchBaseline("example", new()
    ///     {
    ///         ApprovalRules = "{\"PatchRules\":[{\"PatchFilterGroup\":[{\"Key\":\"PatchSet\",\"Values\":[\"OS\"]},{\"Key\":\"ProductFamily\",\"Values\":[\"Windows\"]},{\"Key\":\"Product\",\"Values\":[\"Windows 10\",\"Windows 7\"]},{\"Key\":\"Classification\",\"Values\":[\"Security Updates\",\"Updates\",\"Update Rollups\",\"Critical Updates\"]},{\"Key\":\"Severity\",\"Values\":[\"Critical\",\"Important\",\"Moderate\"]}],\"ApproveAfterDays\":7,\"EnableNonSecurity\":true,\"ComplianceLevel\":\"Medium\"}]}",
    ///         OperationSystem = "Windows",
    ///         PatchBaselineName = "terraform-example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OOS Patch Baseline can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:oos/patchBaseline:PatchBaseline example &lt;patch_baseline_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/patchBaseline:PatchBaseline")]
    public partial class PatchBaseline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
        /// </summary>
        [Output("approvalRules")]
        public Output<string> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// Patches baseline description information.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
        /// </summary>
        [Output("operationSystem")]
        public Output<string> OperationSystem { get; private set; } = null!;

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Output("patchBaselineName")]
        public Output<string> PatchBaselineName { get; private set; } = null!;


        /// <summary>
        /// Create a PatchBaseline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PatchBaseline(string name, PatchBaselineArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oos/patchBaseline:PatchBaseline", name, args ?? new PatchBaselineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PatchBaseline(string name, Input<string> id, PatchBaselineState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/patchBaseline:PatchBaseline", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PatchBaseline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PatchBaseline Get(string name, Input<string> id, PatchBaselineState? state = null, CustomResourceOptions? options = null)
        {
            return new PatchBaseline(name, id, state, options);
        }
    }

    public sealed class PatchBaselineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
        /// </summary>
        [Input("approvalRules", required: true)]
        public Input<string> ApprovalRules { get; set; } = null!;

        /// <summary>
        /// Patches baseline description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
        /// </summary>
        [Input("operationSystem", required: true)]
        public Input<string> OperationSystem { get; set; } = null!;

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Input("patchBaselineName", required: true)]
        public Input<string> PatchBaselineName { get; set; } = null!;

        public PatchBaselineArgs()
        {
        }
        public static new PatchBaselineArgs Empty => new PatchBaselineArgs();
    }

    public sealed class PatchBaselineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept the rules. This value follows the json format. For more details, see the [description of ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/doc-detail/311002.html).
        /// </summary>
        [Input("approvalRules")]
        public Input<string>? ApprovalRules { get; set; }

        /// <summary>
        /// Patches baseline description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
        /// </summary>
        [Input("operationSystem")]
        public Input<string>? OperationSystem { get; set; }

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Input("patchBaselineName")]
        public Input<string>? PatchBaselineName { get; set; }

        public PatchBaselineState()
        {
        }
        public static new PatchBaselineState Empty => new PatchBaselineState();
    }
}

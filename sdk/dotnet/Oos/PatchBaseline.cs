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
    /// For information about OOS Patch Baseline and how to use it, see [What is Patch Baseline](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/patch-manager-overview).
    /// 
    /// &gt; **NOTE:** Available since v1.146.0.
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
    ///     var @default = new AliCloud.Oos.PatchBaseline("default", new()
    ///     {
    ///         PatchBaselineName = name,
    ///         OperationSystem = "Windows",
    ///         ApprovalRules = "{\"PatchRules\":[{\"EnableNonSecurity\":true,\"PatchFilterGroup\":[{\"Values\":[\"*\"],\"Key\":\"Product\"},{\"Values\":[\"Security\",\"Bugfix\"],\"Key\":\"Classification\"},{\"Values\":[\"Critical\",\"Important\"],\"Key\":\"Severity\"}],\"ApproveAfterDays\":7,\"ComplianceLevel\":\"Unspecified\"}]}",
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
    /// $ pulumi import alicloud:oos/patchBaseline:PatchBaseline example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/patchBaseline:PatchBaseline")]
    public partial class PatchBaseline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
        /// </summary>
        [Output("approvalRules")]
        public Output<string> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// Approved Patch.
        /// </summary>
        [Output("approvedPatches")]
        public Output<ImmutableArray<string>> ApprovedPatches { get; private set; } = null!;

        /// <summary>
        /// ApprovedPatchesEnableNonSecurity.
        /// </summary>
        [Output("approvedPatchesEnableNonSecurity")]
        public Output<bool?> ApprovedPatchesEnableNonSecurity { get; private set; } = null!;

        /// <summary>
        /// Creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

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
        /// Reject patches.
        /// </summary>
        [Output("rejectedPatches")]
        public Output<ImmutableArray<string>> RejectedPatches { get; private set; } = null!;

        /// <summary>
        /// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
        /// </summary>
        [Output("rejectedPatchesAction")]
        public Output<string> RejectedPatchesAction { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Source.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<string>> Sources { get; private set; } = null!;

        /// <summary>
        /// Label.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


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
        /// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
        /// </summary>
        [Input("approvalRules", required: true)]
        public Input<string> ApprovalRules { get; set; } = null!;

        [Input("approvedPatches")]
        private InputList<string>? _approvedPatches;

        /// <summary>
        /// Approved Patch.
        /// </summary>
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        /// <summary>
        /// ApprovedPatchesEnableNonSecurity.
        /// </summary>
        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

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

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;

        /// <summary>
        /// Reject patches.
        /// </summary>
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        /// <summary>
        /// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
        /// </summary>
        [Input("rejectedPatchesAction")]
        public Input<string>? RejectedPatchesAction { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("sources")]
        private InputList<string>? _sources;

        /// <summary>
        /// Source.
        /// </summary>
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PatchBaselineArgs()
        {
        }
        public static new PatchBaselineArgs Empty => new PatchBaselineArgs();
    }

    public sealed class PatchBaselineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
        /// </summary>
        [Input("approvalRules")]
        public Input<string>? ApprovalRules { get; set; }

        [Input("approvedPatches")]
        private InputList<string>? _approvedPatches;

        /// <summary>
        /// Approved Patch.
        /// </summary>
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        /// <summary>
        /// ApprovedPatchesEnableNonSecurity.
        /// </summary>
        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

        /// <summary>
        /// Creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

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

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;

        /// <summary>
        /// Reject patches.
        /// </summary>
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        /// <summary>
        /// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
        /// </summary>
        [Input("rejectedPatchesAction")]
        public Input<string>? RejectedPatchesAction { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("sources")]
        private InputList<string>? _sources;

        /// <summary>
        /// Source.
        /// </summary>
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PatchBaselineState()
        {
        }
        public static new PatchBaselineState Empty => new PatchBaselineState();
    }
}

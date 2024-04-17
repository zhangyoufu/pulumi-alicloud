// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    /// <summary>
    /// Provides a Threat Detection Backup Policy resource.
    /// 
    /// For information about Threat Detection Backup Policy and how to use it, see [What is Backup Policy](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = AliCloud.ThreatDetection.GetAssets.Invoke(new()
    ///     {
    ///         MachineTypes = "ecs",
    ///     });
    /// 
    ///     var defaultBackupPolicy = new AliCloud.ThreatDetection.BackupPolicy("default", new()
    ///     {
    ///         BackupPolicyName = "tf-example-name",
    ///         Policy = "{\"Exclude\":[\"/bin/\",\"/usr/bin/\",\"/sbin/\",\"/boot/\",\"/proc/\",\"/sys/\",\"/srv/\",\"/lib/\",\"/selinux/\",\"/usr/sbin/\",\"/run/\",\"/lib32/\",\"/lib64/\",\"/lost+found/\",\"/var/lib/kubelet/\",\"/var/lib/ntp/proc\",\"/var/lib/container\"],\"ExcludeSystemPath\":true,\"Include\":[],\"IsDefault\":1,\"Retention\":7,\"Schedule\":\"I|1668703620|PT24H\",\"Source\":[],\"SpeedLimiter\":\"\",\"UseVss\":true}",
    ///         PolicyVersion = "2.0.0",
    ///         UuidLists = new[]
    ///         {
    ///             @default.Apply(@default =&gt; @default.Apply(getAssetsResult =&gt; getAssetsResult.Ids[0])),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Threat Detection Backup Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:threatdetection/backupPolicy:BackupPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:threatdetection/backupPolicy:BackupPolicy")]
    public partial class BackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Protection of the Name of the Policy.
        /// </summary>
        [Output("backupPolicyName")]
        public Output<string> BackupPolicyName { get; private set; } = null!;

        /// <summary>
        /// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
        /// </summary>
        [Output("policyRegionId")]
        public Output<string?> PolicyRegionId { get; private set; } = null!;

        /// <summary>
        /// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
        /// </summary>
        [Output("policyVersion")]
        public Output<string> PolicyVersion { get; private set; } = null!;

        /// <summary>
        /// The status of the Backup Policy instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specify the Protection of Server UUID List.
        /// </summary>
        [Output("uuidLists")]
        public Output<ImmutableArray<string>> UuidLists { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPolicy(string name, BackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/backupPolicy:BackupPolicy", name, args ?? new BackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPolicy(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/backupPolicy:BackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPolicy Get(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupPolicy(name, id, state, options);
        }
    }

    public sealed class BackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Protection of the Name of the Policy.
        /// </summary>
        [Input("backupPolicyName", required: true)]
        public Input<string> BackupPolicyName { get; set; } = null!;

        /// <summary>
        /// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
        /// </summary>
        [Input("policyRegionId")]
        public Input<string>? PolicyRegionId { get; set; }

        /// <summary>
        /// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
        /// </summary>
        [Input("policyVersion", required: true)]
        public Input<string> PolicyVersion { get; set; } = null!;

        [Input("uuidLists", required: true)]
        private InputList<string>? _uuidLists;

        /// <summary>
        /// Specify the Protection of Server UUID List.
        /// </summary>
        public InputList<string> UuidLists
        {
            get => _uuidLists ?? (_uuidLists = new InputList<string>());
            set => _uuidLists = value;
        }

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }

    public sealed class BackupPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Protection of the Name of the Policy.
        /// </summary>
        [Input("backupPolicyName")]
        public Input<string>? BackupPolicyName { get; set; }

        /// <summary>
        /// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
        /// </summary>
        [Input("policyRegionId")]
        public Input<string>? PolicyRegionId { get; set; }

        /// <summary>
        /// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
        /// </summary>
        [Input("policyVersion")]
        public Input<string>? PolicyVersion { get; set; }

        /// <summary>
        /// The status of the Backup Policy instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("uuidLists")]
        private InputList<string>? _uuidLists;

        /// <summary>
        /// Specify the Protection of Server UUID List.
        /// </summary>
        public InputList<string> UuidLists
        {
            get => _uuidLists ?? (_uuidLists = new InputList<string>());
            set => _uuidLists = value;
        }

        public BackupPolicyState()
        {
        }
        public static new BackupPolicyState Empty => new BackupPolicyState();
    }
}

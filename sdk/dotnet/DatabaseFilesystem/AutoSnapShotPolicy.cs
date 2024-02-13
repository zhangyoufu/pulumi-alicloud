// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DatabaseFilesystem
{
    /// <summary>
    /// Provides a Dbfs Auto Snap Shot Policy resource.
    /// 
    /// For information about Dbfs Auto Snap Shot Policy and how to use it.
    /// 
    /// &gt; **NOTE:** Available since v1.202.0.
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
    ///     var @default = new AliCloud.DatabaseFilesystem.AutoSnapShotPolicy("default", new()
    ///     {
    ///         PolicyName = "tf-example",
    ///         RepeatWeekdays = new[]
    ///         {
    ///             "2",
    ///         },
    ///         RetentionDays = 1,
    ///         TimePoints = new[]
    ///         {
    ///             "01",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dbfs Auto Snap Shot Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy")]
    public partial class AutoSnapShotPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of database file systems set by the automatic snapshot policy.
        /// </summary>
        [Output("appliedDbfsNumber")]
        public Output<int> AppliedDbfsNumber { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Last modification time of automatic snapshot policy
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// Automatic snapshot policy ID
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Automatic snapshot policy name
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        /// </summary>
        [Output("repeatWeekdays")]
        public Output<ImmutableArray<string>> RepeatWeekdays { get; private set; } = null!;

        /// <summary>
        /// Automatic snapshot retention days.
        /// </summary>
        [Output("retentionDays")]
        public Output<int> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// Automatic snapshot policy status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Automatic snapshot policy status details
        /// </summary>
        [Output("statusDetail")]
        public Output<string> StatusDetail { get; private set; } = null!;

        /// <summary>
        /// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        /// </summary>
        [Output("timePoints")]
        public Output<ImmutableArray<string>> TimePoints { get; private set; } = null!;


        /// <summary>
        /// Create a AutoSnapShotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoSnapShotPolicy(string name, AutoSnapShotPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy", name, args ?? new AutoSnapShotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoSnapShotPolicy(string name, Input<string> id, AutoSnapShotPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoSnapShotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoSnapShotPolicy Get(string name, Input<string> id, AutoSnapShotPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoSnapShotPolicy(name, id, state, options);
        }
    }

    public sealed class AutoSnapShotPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatic snapshot policy name
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        [Input("repeatWeekdays", required: true)]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// Automatic snapshot retention days.
        /// </summary>
        [Input("retentionDays", required: true)]
        public Input<int> RetentionDays { get; set; } = null!;

        [Input("timePoints", required: true)]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        public AutoSnapShotPolicyArgs()
        {
        }
        public static new AutoSnapShotPolicyArgs Empty => new AutoSnapShotPolicyArgs();
    }

    public sealed class AutoSnapShotPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of database file systems set by the automatic snapshot policy.
        /// </summary>
        [Input("appliedDbfsNumber")]
        public Input<int>? AppliedDbfsNumber { get; set; }

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Last modification time of automatic snapshot policy
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        /// <summary>
        /// Automatic snapshot policy ID
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Automatic snapshot policy name
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        [Input("repeatWeekdays")]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// Automatic snapshot retention days.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// Automatic snapshot policy status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Automatic snapshot policy status details
        /// </summary>
        [Input("statusDetail")]
        public Input<string>? StatusDetail { get; set; }

        [Input("timePoints")]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        public AutoSnapShotPolicyState()
        {
        }
        public static new AutoSnapShotPolicyState Empty => new AutoSnapShotPolicyState();
    }
}

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
    /// Provides an ECS snapshot policy resource.
    /// 
    /// For information about snapshot policy and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
    /// 
    /// &gt; **NOTE:** Available in 1.42.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sp = new AliCloud.Ecs.SnapshotPolicy("sp", new AliCloud.Ecs.SnapshotPolicyArgs
    ///         {
    ///             RepeatWeekdays = 
    ///             {
    ///                 "1",
    ///                 "2",
    ///                 "3",
    ///             },
    ///             RetentionDays = -1,
    ///             TimePoints = 
    ///             {
    ///                 "1",
    ///                 "22",
    ///                 "23",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Snapshot can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/snapshotPolicy:SnapshotPolicy snapshot sp-abc1234567890000
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/snapshotPolicy:SnapshotPolicy")]
    public partial class SnapshotPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The snapshot policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
        /// - A maximum of seven time points can be selected.
        /// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        /// </summary>
        [Output("repeatWeekdays")]
        public Output<ImmutableArray<string>> RepeatWeekdays { get; private set; } = null!;

        /// <summary>
        /// The snapshot retention time, and the unit of measurement is day. Optional values:
        /// - -1: The automatic snapshots are retained permanently.
        /// - [1, 65536]: The number of days retained.
        /// </summary>
        [Output("retentionDays")]
        public Output<int> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
        /// - A maximum of 24 time points can be selected.
        /// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        /// </summary>
        [Output("timePoints")]
        public Output<ImmutableArray<string>> TimePoints { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotPolicy(string name, SnapshotPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, args ?? new SnapshotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotPolicy(string name, Input<string> id, SnapshotPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotPolicy Get(string name, Input<string> id, SnapshotPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotPolicy(name, id, state, options);
        }
    }

    public sealed class SnapshotPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The snapshot policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("repeatWeekdays", required: true)]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
        /// - A maximum of seven time points can be selected.
        /// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// The snapshot retention time, and the unit of measurement is day. Optional values:
        /// - -1: The automatic snapshots are retained permanently.
        /// - [1, 65536]: The number of days retained.
        /// </summary>
        [Input("retentionDays", required: true)]
        public Input<int> RetentionDays { get; set; } = null!;

        [Input("timePoints", required: true)]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
        /// - A maximum of 24 time points can be selected.
        /// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        public SnapshotPolicyArgs()
        {
        }
    }

    public sealed class SnapshotPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The snapshot policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("repeatWeekdays")]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
        /// - A maximum of seven time points can be selected.
        /// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// The snapshot retention time, and the unit of measurement is day. Optional values:
        /// - -1: The automatic snapshots are retained permanently.
        /// - [1, 65536]: The number of days retained.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        [Input("timePoints")]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
        /// - A maximum of 24 time points can be selected.
        /// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        public SnapshotPolicyState()
        {
        }
    }
}

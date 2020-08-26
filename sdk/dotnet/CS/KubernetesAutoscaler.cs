// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// This resource will help you to manager cluster-autoscaler in Kubernetes Cluster.
    /// 
    /// &gt; **NOTE:** The scaling group must use CentOS7 or AliyunLinux2 as base image.
    /// 
    /// &gt; **NOTE:** The cluster-autoscaler can only use the same size of instanceTypes in one scaling group.
    /// 
    /// &gt; **NOTE:** Add Policy to RAM role of the node to deploy cluster-autoscaler if you need.
    /// 
    /// &gt; **NOTE:** Available in 1.65.0+.
    /// 
    /// ## Example Usage
    /// 
    /// cluster-autoscaler in Kubernetes Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new AliCloud.CS.KubernetesAutoscaler("default", new AliCloud.CS.KubernetesAutoscalerArgs
    ///         {
    ///             ClusterId = @var.Cluster_id,
    ///             CoolDownDuration = @var.Cool_down_duration,
    ///             DeferScaleInDuration = @var.Defer_scale_in_duration,
    ///             Nodepools = 
    ///             {
    ///                 new AliCloud.CS.Inputs.KubernetesAutoscalerNodepoolArgs
    ///                 {
    ///                     Id = "scaling_group_id",
    ///                     Labels = "a=b",
    ///                     Taints = "c=d:NoSchedule",
    ///                 },
    ///             },
    ///             Utilization = @var.Utilization,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class KubernetesAutoscaler : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The cool_down_duration option of cluster-autoscaler.
        /// </summary>
        [Output("coolDownDuration")]
        public Output<string> CoolDownDuration { get; private set; } = null!;

        /// <summary>
        /// The defer_scale_in_duration option of cluster-autoscaler.
        /// </summary>
        [Output("deferScaleInDuration")]
        public Output<string> DeferScaleInDuration { get; private set; } = null!;

        /// <summary>
        /// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
        /// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
        /// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
        /// </summary>
        [Output("nodepools")]
        public Output<ImmutableArray<Outputs.KubernetesAutoscalerNodepool>> Nodepools { get; private set; } = null!;

        /// <summary>
        /// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
        /// </summary>
        [Output("useEcsRamRoleToken")]
        public Output<bool?> UseEcsRamRoleToken { get; private set; } = null!;

        /// <summary>
        /// The utilization option of cluster-autoscaler.
        /// </summary>
        [Output("utilization")]
        public Output<string> Utilization { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesAutoscaler resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesAutoscaler(string name, KubernetesAutoscalerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, args ?? new KubernetesAutoscalerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesAutoscaler(string name, Input<string> id, KubernetesAutoscalerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KubernetesAutoscaler resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesAutoscaler Get(string name, Input<string> id, KubernetesAutoscalerState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesAutoscaler(name, id, state, options);
        }
    }

    public sealed class KubernetesAutoscalerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The cool_down_duration option of cluster-autoscaler.
        /// </summary>
        [Input("coolDownDuration", required: true)]
        public Input<string> CoolDownDuration { get; set; } = null!;

        /// <summary>
        /// The defer_scale_in_duration option of cluster-autoscaler.
        /// </summary>
        [Input("deferScaleInDuration", required: true)]
        public Input<string> DeferScaleInDuration { get; set; } = null!;

        [Input("nodepools")]
        private InputList<Inputs.KubernetesAutoscalerNodepoolArgs>? _nodepools;

        /// <summary>
        /// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
        /// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
        /// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
        /// </summary>
        public InputList<Inputs.KubernetesAutoscalerNodepoolArgs> Nodepools
        {
            get => _nodepools ?? (_nodepools = new InputList<Inputs.KubernetesAutoscalerNodepoolArgs>());
            set => _nodepools = value;
        }

        /// <summary>
        /// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
        /// </summary>
        [Input("useEcsRamRoleToken")]
        public Input<bool>? UseEcsRamRoleToken { get; set; }

        /// <summary>
        /// The utilization option of cluster-autoscaler.
        /// </summary>
        [Input("utilization", required: true)]
        public Input<string> Utilization { get; set; } = null!;

        public KubernetesAutoscalerArgs()
        {
        }
    }

    public sealed class KubernetesAutoscalerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The cool_down_duration option of cluster-autoscaler.
        /// </summary>
        [Input("coolDownDuration")]
        public Input<string>? CoolDownDuration { get; set; }

        /// <summary>
        /// The defer_scale_in_duration option of cluster-autoscaler.
        /// </summary>
        [Input("deferScaleInDuration")]
        public Input<string>? DeferScaleInDuration { get; set; }

        [Input("nodepools")]
        private InputList<Inputs.KubernetesAutoscalerNodepoolGetArgs>? _nodepools;

        /// <summary>
        /// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
        /// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
        /// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
        /// </summary>
        public InputList<Inputs.KubernetesAutoscalerNodepoolGetArgs> Nodepools
        {
            get => _nodepools ?? (_nodepools = new InputList<Inputs.KubernetesAutoscalerNodepoolGetArgs>());
            set => _nodepools = value;
        }

        /// <summary>
        /// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
        /// </summary>
        [Input("useEcsRamRoleToken")]
        public Input<bool>? UseEcsRamRoleToken { get; set; }

        /// <summary>
        /// The utilization option of cluster-autoscaler.
        /// </summary>
        [Input("utilization")]
        public Input<string>? Utilization { get; set; }

        public KubernetesAutoscalerState()
        {
        }
    }
}

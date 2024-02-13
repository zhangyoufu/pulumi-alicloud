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
    /// This resource will help you to manage addon in Kubernetes Cluster, see [What is kubernetes addon](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/api-install-a-component-in-an-ack-cluster).
    /// 
    /// &gt; **NOTE:** Available since v1.150.0.
    /// 
    /// &gt; **NOTE:** From version 1.166.0, support specifying addon customizable configuration.
    /// 
    /// ## Import
    /// 
    /// Cluster addon can be imported by cluster id and addon name. Then write the addon.tf file according to the result of `pulumi preview`.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cs/kubernetesAddon:KubernetesAddon my_addon &lt;cluster_id&gt;:&lt;addon_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cs/kubernetesAddon:KubernetesAddon")]
    public partial class KubernetesAddon : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Is the addon ready for upgrade.
        /// </summary>
        [Output("canUpgrade")]
        public Output<bool> CanUpgrade { get; private set; } = null!;

        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The version which addon can be upgraded to.
        /// </summary>
        [Output("nextVersion")]
        public Output<string> NextVersion { get; private set; } = null!;

        /// <summary>
        /// Is it a mandatory addon to be installed.
        /// </summary>
        [Output("required")]
        public Output<bool> Required { get; private set; } = null!;

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesAddon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesAddon(string name, KubernetesAddonArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAddon:KubernetesAddon", name, args ?? new KubernetesAddonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesAddon(string name, Input<string> id, KubernetesAddonState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAddon:KubernetesAddon", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KubernetesAddon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesAddon Get(string name, Input<string> id, KubernetesAddonState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesAddon(name, id, state, options);
        }
    }

    public sealed class KubernetesAddonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public KubernetesAddonArgs()
        {
        }
        public static new KubernetesAddonArgs Empty => new KubernetesAddonArgs();
    }

    public sealed class KubernetesAddonState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the addon ready for upgrade.
        /// </summary>
        [Input("canUpgrade")]
        public Input<bool>? CanUpgrade { get; set; }

        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The version which addon can be upgraded to.
        /// </summary>
        [Input("nextVersion")]
        public Input<string>? NextVersion { get; set; }

        /// <summary>
        /// Is it a mandatory addon to be installed.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubernetesAddonState()
        {
        }
        public static new KubernetesAddonState Empty => new KubernetesAddonState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a network acl resource to add network acls.
    /// 
    /// &gt; **NOTE:** Available in 1.43.0+. Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
    /// 
    /// ## Import
    /// 
    /// The network acl can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/networkAcl:NetworkAcl default nacl-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/networkAcl:NetworkAcl")]
    public partial class NetworkAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the network acl instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block `egress_acl_entries`.
        /// </summary>
        [Output("egressAclEntries")]
        public Output<ImmutableArray<Outputs.NetworkAclEgressAclEntry>> EgressAclEntries { get; private set; } = null!;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block `ingress_acl_entries`.
        /// </summary>
        [Output("ingressAclEntries")]
        public Output<ImmutableArray<Outputs.NetworkAclIngressAclEntry>> IngressAclEntries { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.122.0. New field `network_acl_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the network acl.
        /// </summary>
        [Output("networkAclName")]
        public Output<string> NetworkAclName { get; private set; } = null!;

        /// <summary>
        /// The associated resources. See the following `Block resources`. **NOTE:** "Field `resources` has been deprecated from provider version 1.193.0 and it will be removed in the future version. Please use the new resource `alicloud.vpc.VpcNetworkAclAttachment`."
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.NetworkAclResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.122.0+) The status of the network acl.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The vpc_id of the network acl, the field can't be changed.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAcl(string name, NetworkAclArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAcl:NetworkAcl", name, args ?? new NetworkAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkAcl(string name, Input<string> id, NetworkAclState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAcl:NetworkAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAcl Get(string name, Input<string> id, NetworkAclState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAcl(name, id, state, options);
        }
    }

    public sealed class NetworkAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the network acl instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("egressAclEntries")]
        private InputList<Inputs.NetworkAclEgressAclEntryArgs>? _egressAclEntries;

        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block `egress_acl_entries`.
        /// </summary>
        public InputList<Inputs.NetworkAclEgressAclEntryArgs> EgressAclEntries
        {
            get => _egressAclEntries ?? (_egressAclEntries = new InputList<Inputs.NetworkAclEgressAclEntryArgs>());
            set => _egressAclEntries = value;
        }

        [Input("ingressAclEntries")]
        private InputList<Inputs.NetworkAclIngressAclEntryArgs>? _ingressAclEntries;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block `ingress_acl_entries`.
        /// </summary>
        public InputList<Inputs.NetworkAclIngressAclEntryArgs> IngressAclEntries
        {
            get => _ingressAclEntries ?? (_ingressAclEntries = new InputList<Inputs.NetworkAclIngressAclEntryArgs>());
            set => _ingressAclEntries = value;
        }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.122.0. New field `network_acl_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the network acl.
        /// </summary>
        [Input("networkAclName")]
        public Input<string>? NetworkAclName { get; set; }

        [Input("resources")]
        private InputList<Inputs.NetworkAclResourceArgs>? _resources;

        /// <summary>
        /// The associated resources. See the following `Block resources`. **NOTE:** "Field `resources` has been deprecated from provider version 1.193.0 and it will be removed in the future version. Please use the new resource `alicloud.vpc.VpcNetworkAclAttachment`."
        /// </summary>
        [Obsolete(@"Field 'resources' has been deprecated from provider version 1.193.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_network_acl_attachment'.")]
        public InputList<Inputs.NetworkAclResourceArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NetworkAclResourceArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// The vpc_id of the network acl, the field can't be changed.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public NetworkAclArgs()
        {
        }
        public static new NetworkAclArgs Empty => new NetworkAclArgs();
    }

    public sealed class NetworkAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the network acl instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("egressAclEntries")]
        private InputList<Inputs.NetworkAclEgressAclEntryGetArgs>? _egressAclEntries;

        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block `egress_acl_entries`.
        /// </summary>
        public InputList<Inputs.NetworkAclEgressAclEntryGetArgs> EgressAclEntries
        {
            get => _egressAclEntries ?? (_egressAclEntries = new InputList<Inputs.NetworkAclEgressAclEntryGetArgs>());
            set => _egressAclEntries = value;
        }

        [Input("ingressAclEntries")]
        private InputList<Inputs.NetworkAclIngressAclEntryGetArgs>? _ingressAclEntries;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block `ingress_acl_entries`.
        /// </summary>
        public InputList<Inputs.NetworkAclIngressAclEntryGetArgs> IngressAclEntries
        {
            get => _ingressAclEntries ?? (_ingressAclEntries = new InputList<Inputs.NetworkAclIngressAclEntryGetArgs>());
            set => _ingressAclEntries = value;
        }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.122.0. New field `network_acl_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the network acl.
        /// </summary>
        [Input("networkAclName")]
        public Input<string>? NetworkAclName { get; set; }

        [Input("resources")]
        private InputList<Inputs.NetworkAclResourceGetArgs>? _resources;

        /// <summary>
        /// The associated resources. See the following `Block resources`. **NOTE:** "Field `resources` has been deprecated from provider version 1.193.0 and it will be removed in the future version. Please use the new resource `alicloud.vpc.VpcNetworkAclAttachment`."
        /// </summary>
        [Obsolete(@"Field 'resources' has been deprecated from provider version 1.193.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_network_acl_attachment'.")]
        public InputList<Inputs.NetworkAclResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NetworkAclResourceGetArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// (Available in 1.122.0+) The status of the network acl.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vpc_id of the network acl, the field can't be changed.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public NetworkAclState()
        {
        }
        public static new NetworkAclState Empty => new NetworkAclState();
    }
}

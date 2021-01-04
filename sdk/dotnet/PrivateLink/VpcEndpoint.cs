// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    /// <summary>
    /// Provides a Private Link Vpc Endpoint resource.
    /// 
    /// For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://help.aliyun.com/document_detail/120479.html).
    /// 
    /// &gt; **NOTE:** Available in v1.109.0+.
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:privatelink/vpcEndpoint:VpcEndpoint example &lt;endpoint_id&gt;
    /// ```
    /// </summary>
    public partial class VpcEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The Bandwidth.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The status of Connection.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// The dry run. Default to: `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The status of Endpoint Business.
        /// </summary>
        [Output("endpointBusinessStatus")]
        public Output<string> EndpointBusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
        /// </summary>
        [Output("endpointDescription")]
        public Output<string?> EndpointDescription { get; private set; } = null!;

        /// <summary>
        /// The Endpoint Domain.
        /// </summary>
        [Output("endpointDomain")]
        public Output<string> EndpointDomain { get; private set; } = null!;

        /// <summary>
        /// The security group associated with the terminal node network card.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The terminal node service associated with the terminal node.
        /// </summary>
        [Output("serviceId")]
        public Output<string?> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The name of the terminal node service associated with the terminal node.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The status of Vpc Endpoint.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
        /// </summary>
        [Output("vpcEndpointName")]
        public Output<string?> VpcEndpointName { get; private set; } = null!;

        /// <summary>
        /// The private network to which the terminal node belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpoint(string name, VpcEndpointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, args ?? new VpcEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpoint(string name, Input<string> id, VpcEndpointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpoint Get(string name, Input<string> id, VpcEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpoint(name, id, state, options);
        }
    }

    public sealed class VpcEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dry run. Default to: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
        /// </summary>
        [Input("endpointDescription")]
        public Input<string>? EndpointDescription { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group associated with the terminal node network card.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The terminal node service associated with the terminal node.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The name of the terminal node service associated with the terminal node.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
        /// </summary>
        [Input("vpcEndpointName")]
        public Input<string>? VpcEndpointName { get; set; }

        /// <summary>
        /// The private network to which the terminal node belongs.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcEndpointArgs()
        {
        }
    }

    public sealed class VpcEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Bandwidth.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The status of Connection.
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// The dry run. Default to: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The status of Endpoint Business.
        /// </summary>
        [Input("endpointBusinessStatus")]
        public Input<string>? EndpointBusinessStatus { get; set; }

        /// <summary>
        /// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
        /// </summary>
        [Input("endpointDescription")]
        public Input<string>? EndpointDescription { get; set; }

        /// <summary>
        /// The Endpoint Domain.
        /// </summary>
        [Input("endpointDomain")]
        public Input<string>? EndpointDomain { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group associated with the terminal node network card.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The terminal node service associated with the terminal node.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The name of the terminal node service associated with the terminal node.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The status of Vpc Endpoint.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
        /// </summary>
        [Input("vpcEndpointName")]
        public Input<string>? VpcEndpointName { get; set; }

        /// <summary>
        /// The private network to which the terminal node belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcEndpointState()
        {
        }
    }
}

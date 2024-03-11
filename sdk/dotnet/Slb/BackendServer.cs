// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    /// <summary>
    /// Add a group of backend servers (ECS or ENI instance) to the Server Load Balancer or remove them from it.
    /// 
    /// &gt; **NOTE:** Available in 1.53.0+
    /// 
    /// ## Import
    /// 
    /// Load balancer backend server can be imported using the load balancer id.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:slb/backendServer:BackendServer example &lt;load_balancer_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/backendServer:BackendServer")]
    public partial class BackendServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        [Output("backendServers")]
        public Output<ImmutableArray<Outputs.BackendServerBackendServer>> BackendServers { get; private set; } = null!;

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;


        /// <summary>
        /// Create a BackendServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendServer(string name, BackendServerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/backendServer:BackendServer", name, args ?? new BackendServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendServer(string name, Input<string> id, BackendServerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/backendServer:BackendServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendServer Get(string name, Input<string> id, BackendServerState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendServer(name, id, state, options);
        }
    }

    public sealed class BackendServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("backendServers")]
        private InputList<Inputs.BackendServerBackendServerArgs>? _backendServers;

        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.BackendServerBackendServerArgs> BackendServers
        {
            get => _backendServers ?? (_backendServers = new InputList<Inputs.BackendServerBackendServerArgs>());
            set => _backendServers = value;
        }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        public BackendServerArgs()
        {
        }
        public static new BackendServerArgs Empty => new BackendServerArgs();
    }

    public sealed class BackendServerState : global::Pulumi.ResourceArgs
    {
        [Input("backendServers")]
        private InputList<Inputs.BackendServerBackendServerGetArgs>? _backendServers;

        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.BackendServerBackendServerGetArgs> BackendServers
        {
            get => _backendServers ?? (_backendServers = new InputList<Inputs.BackendServerBackendServerGetArgs>());
            set => _backendServers = value;
        }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        public BackendServerState()
        {
        }
        public static new BackendServerState Empty => new BackendServerState();
    }
}

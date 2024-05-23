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
    /// For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-createvpcendpoint).
    /// 
    /// &gt; **NOTE:** Available since v1.109.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultbFzA4a = new AliCloud.Vpc.Network("defaultbFzA4a", new()
    ///     {
    ///         Description = "example-terraform",
    ///         CidrBlock = "172.16.0.0/12",
    ///         VpcName = name,
    ///     });
    /// 
    ///     var default1FTFrP = new AliCloud.Ecs.SecurityGroup("default1FTFrP", new()
    ///     {
    ///         Name = name,
    ///         VpcId = defaultbFzA4a.Id,
    ///     });
    /// 
    ///     var defaultjljY5S = new AliCloud.Ecs.SecurityGroup("defaultjljY5S", new()
    ///     {
    ///         Name = name,
    ///         VpcId = defaultbFzA4a.Id,
    ///     });
    /// 
    ///     var defaultVpcEndpoint = new AliCloud.PrivateLink.VpcEndpoint("default", new()
    ///     {
    ///         EndpointDescription = name,
    ///         VpcEndpointName = name,
    ///         ResourceGroupId = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0])),
    ///         EndpointType = "Interface",
    ///         VpcId = defaultbFzA4a.Id,
    ///         ServiceName = "com.aliyuncs.privatelink.ap-southeast-5.oss",
    ///         DryRun = false,
    ///         ZonePrivateIpAddressCount = 1,
    ///         PolicyDocument = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "1",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Effect"] = "Allow",
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "*",
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         "*",
    ///                     },
    ///                     ["Principal"] = "*",
    ///                 },
    ///             },
    ///         }),
    ///         SecurityGroupIds = new[]
    ///         {
    ///             default1FTFrP.Id,
    ///         },
    ///         ServiceId = "epsrv-k1apjysze8u1l9t6uyg9",
    ///         ProtectedEnabled = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:privatelink/vpcEndpoint:VpcEndpoint example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:privatelink/vpcEndpoint:VpcEndpoint")]
    public partial class VpcEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s. Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The state of the endpoint connection.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// The time when the endpoint was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The service state of the endpoint.
        /// </summary>
        [Output("endpointBusinessStatus")]
        public Output<string> EndpointBusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The description of the endpoint.
        /// </summary>
        [Output("endpointDescription")]
        public Output<string?> EndpointDescription { get; private set; } = null!;

        /// <summary>
        /// The domain name of the endpoint.
        /// </summary>
        [Output("endpointDomain")]
        public Output<string> EndpointDomain { get; private set; } = null!;

        /// <summary>
        /// The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// RAM access policies.
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
        /// - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
        /// - **false (default)**: disables user authentication.
        /// </summary>
        [Output("protectedEnabled")]
        public Output<bool?> ProtectedEnabled { get; private set; } = null!;

        /// <summary>
        /// The resource group ID.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Output("serviceId")]
        public Output<string?> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The state of the endpoint.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The list of tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Output("vpcEndpointName")]
        public Output<string?> VpcEndpointName { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC to which the endpoint belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
        /// </summary>
        [Output("zonePrivateIpAddressCount")]
        public Output<int> ZonePrivateIpAddressCount { get; private set; } = null!;


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

    public sealed class VpcEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The description of the endpoint.
        /// </summary>
        [Input("endpointDescription")]
        public Input<string>? EndpointDescription { get; set; }

        /// <summary>
        /// The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// RAM access policies.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
        /// - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
        /// - **false (default)**: disables user authentication.
        /// </summary>
        [Input("protectedEnabled")]
        public Input<bool>? ProtectedEnabled { get; set; }

        /// <summary>
        /// The resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The ID of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The name of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("vpcEndpointName")]
        public Input<string>? VpcEndpointName { get; set; }

        /// <summary>
        /// The ID of the VPC to which the endpoint belongs.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
        /// </summary>
        [Input("zonePrivateIpAddressCount")]
        public Input<int>? ZonePrivateIpAddressCount { get; set; }

        public VpcEndpointArgs()
        {
        }
        public static new VpcEndpointArgs Empty => new VpcEndpointArgs();
    }

    public sealed class VpcEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s. Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The state of the endpoint connection.
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// The time when the endpoint was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The service state of the endpoint.
        /// </summary>
        [Input("endpointBusinessStatus")]
        public Input<string>? EndpointBusinessStatus { get; set; }

        /// <summary>
        /// The description of the endpoint.
        /// </summary>
        [Input("endpointDescription")]
        public Input<string>? EndpointDescription { get; set; }

        /// <summary>
        /// The domain name of the endpoint.
        /// </summary>
        [Input("endpointDomain")]
        public Input<string>? EndpointDomain { get; set; }

        /// <summary>
        /// The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// RAM access policies.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
        /// - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
        /// - **false (default)**: disables user authentication.
        /// </summary>
        [Input("protectedEnabled")]
        public Input<bool>? ProtectedEnabled { get; set; }

        /// <summary>
        /// The resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The ID of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The name of the endpoint service with which the endpoint is associated.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The state of the endpoint.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("vpcEndpointName")]
        public Input<string>? VpcEndpointName { get; set; }

        /// <summary>
        /// The ID of the VPC to which the endpoint belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
        /// </summary>
        [Input("zonePrivateIpAddressCount")]
        public Input<int>? ZonePrivateIpAddressCount { get; set; }

        public VpcEndpointState()
        {
        }
        public static new VpcEndpointState Empty => new VpcEndpointState();
    }
}

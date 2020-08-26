// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// This topic provides an overview of the route map function of Cloud Enterprise Networks (CENs).
    /// You can use the route map function to filter routes and modify route attributes.
    /// By doing so, you can manage the communication between networks attached to a CEN.
    /// 
    /// For information about CEN Route Map and how to use it, see [Manage CEN Route Map](https://www.alibabacloud.com/help/doc-detail/124157.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.82.0+
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new AliCloud.Cen.InstanceArgs
    ///         {
    ///         });
    ///         var vpc00Region = new AliCloud.Provider("vpc00Region", new AliCloud.ProviderArgs
    ///         {
    ///             Region = "cn-hangzhou",
    ///         });
    ///         var vpc01Region = new AliCloud.Provider("vpc01Region", new AliCloud.ProviderArgs
    ///         {
    ///             Region = "cn-shanghai",
    ///         });
    ///         var vpc00 = new AliCloud.Vpc.Network("vpc00", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = "alicloud.vpc00_region",
    ///         });
    ///         var vpc01 = new AliCloud.Vpc.Network("vpc01", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = "alicloud.vpc01_region",
    ///         });
    ///         var default00 = new AliCloud.Cen.InstanceAttachment("default00", new AliCloud.Cen.InstanceAttachmentArgs
    ///         {
    ///             ChildInstanceId = vpc00.Id,
    ///             ChildInstanceRegionId = "cn-hangzhou",
    ///             InstanceId = defaultInstance.Id,
    ///         });
    ///         var default01 = new AliCloud.Cen.InstanceAttachment("default01", new AliCloud.Cen.InstanceAttachmentArgs
    ///         {
    ///             ChildInstanceId = vpc01.Id,
    ///             ChildInstanceRegionId = "cn-shanghai",
    ///             InstanceId = defaultInstance.Id,
    ///         });
    ///         var defaultRouteMap = new AliCloud.Cen.RouteMap("defaultRouteMap", new AliCloud.Cen.RouteMapArgs
    ///         {
    ///             AsPathMatchMode = "Include",
    ///             CenId = alicloud_cen_instance.Cen.Id,
    ///             CenRegionId = "cn-hangzhou",
    ///             CidrMatchMode = "Include",
    ///             CommunityMatchMode = "Include",
    ///             CommunityOperateMode = "Additive",
    ///             Description = "test-desc",
    ///             DestinationChildInstanceTypes = 
    ///             {
    ///                 "VPC",
    ///             },
    ///             DestinationCidrBlocks = 
    ///             {
    ///                 vpc01.CidrBlock,
    ///             },
    ///             DestinationInstanceIds = 
    ///             {
    ///                 vpc01.Id,
    ///             },
    ///             DestinationInstanceIdsReverseMatch = false,
    ///             DestinationRouteTableIds = 
    ///             {
    ///                 vpc01.RouteTableId,
    ///             },
    ///             MapResult = "Permit",
    ///             MatchAsns = 
    ///             {
    ///                 "65501",
    ///             },
    ///             MatchCommunitySets = 
    ///             {
    ///                 "65501:1",
    ///             },
    ///             NextPriority = 1,
    ///             OperateCommunitySets = 
    ///             {
    ///                 "65501:1",
    ///             },
    ///             Preference = 20,
    ///             PrependAsPaths = 
    ///             {
    ///                 "65501",
    ///             },
    ///             Priority = 1,
    ///             RouteTypes = 
    ///             {
    ///                 "System",
    ///             },
    ///             SourceChildInstanceTypes = 
    ///             {
    ///                 "VPC",
    ///             },
    ///             SourceInstanceIds = 
    ///             {
    ///                 vpc00.Id,
    ///             },
    ///             SourceInstanceIdsReverseMatch = false,
    ///             SourceRegionIds = 
    ///             {
    ///                 "cn-hangzhou",
    ///             },
    ///             SourceRouteTableIds = 
    ///             {
    ///                 vpc00.RouteTableId,
    ///             },
    ///             TransmitDirection = "RegionIn",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 "alicloud_cen_instance_attachment.default00",
    ///                 "alicloud_cen_instance_attachment.default01",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RouteMap : Pulumi.CustomResource
    {
        /// <summary>
        /// A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Output("asPathMatchMode")]
        public Output<string?> AsPathMatchMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The ID of the region to which the CEN instance belongs.
        /// </summary>
        [Output("cenRegionId")]
        public Output<string> CenRegionId { get; private set; } = null!;

        /// <summary>
        /// A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Output("cidrMatchMode")]
        public Output<string?> CidrMatchMode { get; private set; } = null!;

        /// <summary>
        /// A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Output("communityMatchMode")]
        public Output<string?> CommunityMatchMode { get; private set; } = null!;

        /// <summary>
        /// An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        /// </summary>
        [Output("communityOperateMode")]
        public Output<string?> CommunityOperateMode { get; private set; } = null!;

        /// <summary>
        /// The description of the route map.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        [Output("destinationChildInstanceTypes")]
        public Output<ImmutableArray<string>> DestinationChildInstanceTypes { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        /// </summary>
        [Output("destinationCidrBlocks")]
        public Output<ImmutableArray<string>> DestinationCidrBlocks { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination instances.
        /// </summary>
        [Output("destinationInstanceIds")]
        public Output<ImmutableArray<string>> DestinationInstanceIds { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Output("destinationInstanceIdsReverseMatch")]
        public Output<bool?> DestinationInstanceIdsReverseMatch { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        [Output("destinationRouteTableIds")]
        public Output<ImmutableArray<string>> DestinationRouteTableIds { get; private set; } = null!;

        /// <summary>
        /// The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        /// </summary>
        [Output("mapResult")]
        public Output<string> MapResult { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        /// </summary>
        [Output("matchAsns")]
        public Output<ImmutableArray<string>> MatchAsns { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        [Output("matchCommunitySets")]
        public Output<ImmutableArray<string>> MatchCommunitySets { get; private set; } = null!;

        /// <summary>
        /// The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        /// </summary>
        [Output("nextPriority")]
        public Output<int?> NextPriority { get; private set; } = null!;

        /// <summary>
        /// An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        [Output("operateCommunitySets")]
        public Output<ImmutableArray<string>> OperateCommunitySets { get; private set; } = null!;

        /// <summary>
        /// An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        /// </summary>
        [Output("preference")]
        public Output<int?> Preference { get; private set; } = null!;

        /// <summary>
        /// An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        /// </summary>
        [Output("prependAsPaths")]
        public Output<ImmutableArray<string>> PrependAsPaths { get; private set; } = null!;

        /// <summary>
        /// The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("routeMapId")]
        public Output<string> RouteMapId { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        /// </summary>
        [Output("routeTypes")]
        public Output<ImmutableArray<string>> RouteTypes { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        [Output("sourceChildInstanceTypes")]
        public Output<ImmutableArray<string>> SourceChildInstanceTypes { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source instances.
        /// </summary>
        [Output("sourceInstanceIds")]
        public Output<ImmutableArray<string>> SourceInstanceIds { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Output("sourceInstanceIdsReverseMatch")]
        public Output<bool?> SourceInstanceIdsReverseMatch { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        /// </summary>
        [Output("sourceRegionIds")]
        public Output<ImmutableArray<string>> SourceRegionIds { get; private set; } = null!;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        [Output("sourceRouteTableIds")]
        public Output<ImmutableArray<string>> SourceRouteTableIds { get; private set; } = null!;

        /// <summary>
        /// (Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"].
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        /// </summary>
        [Output("transmitDirection")]
        public Output<string> TransmitDirection { get; private set; } = null!;


        /// <summary>
        /// Create a RouteMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteMap(string name, RouteMapArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeMap:RouteMap", name, args ?? new RouteMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteMap(string name, Input<string> id, RouteMapState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeMap:RouteMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteMap Get(string name, Input<string> id, RouteMapState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteMap(name, id, state, options);
        }
    }

    public sealed class RouteMapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("asPathMatchMode")]
        public Input<string>? AsPathMatchMode { get; set; }

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The ID of the region to which the CEN instance belongs.
        /// </summary>
        [Input("cenRegionId", required: true)]
        public Input<string> CenRegionId { get; set; } = null!;

        /// <summary>
        /// A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("cidrMatchMode")]
        public Input<string>? CidrMatchMode { get; set; }

        /// <summary>
        /// A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("communityMatchMode")]
        public Input<string>? CommunityMatchMode { get; set; }

        /// <summary>
        /// An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        /// </summary>
        [Input("communityOperateMode")]
        public Input<string>? CommunityOperateMode { get; set; }

        /// <summary>
        /// The description of the route map.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationChildInstanceTypes")]
        private InputList<string>? _destinationChildInstanceTypes;

        /// <summary>
        /// A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        public InputList<string> DestinationChildInstanceTypes
        {
            get => _destinationChildInstanceTypes ?? (_destinationChildInstanceTypes = new InputList<string>());
            set => _destinationChildInstanceTypes = value;
        }

        [Input("destinationCidrBlocks")]
        private InputList<string>? _destinationCidrBlocks;

        /// <summary>
        /// A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        /// </summary>
        public InputList<string> DestinationCidrBlocks
        {
            get => _destinationCidrBlocks ?? (_destinationCidrBlocks = new InputList<string>());
            set => _destinationCidrBlocks = value;
        }

        [Input("destinationInstanceIds")]
        private InputList<string>? _destinationInstanceIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination instances.
        /// </summary>
        public InputList<string> DestinationInstanceIds
        {
            get => _destinationInstanceIds ?? (_destinationInstanceIds = new InputList<string>());
            set => _destinationInstanceIds = value;
        }

        /// <summary>
        /// Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Input("destinationInstanceIdsReverseMatch")]
        public Input<bool>? DestinationInstanceIdsReverseMatch { get; set; }

        [Input("destinationRouteTableIds")]
        private InputList<string>? _destinationRouteTableIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        public InputList<string> DestinationRouteTableIds
        {
            get => _destinationRouteTableIds ?? (_destinationRouteTableIds = new InputList<string>());
            set => _destinationRouteTableIds = value;
        }

        /// <summary>
        /// The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        /// </summary>
        [Input("mapResult", required: true)]
        public Input<string> MapResult { get; set; } = null!;

        [Input("matchAsns")]
        private InputList<string>? _matchAsns;

        /// <summary>
        /// A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        /// </summary>
        public InputList<string> MatchAsns
        {
            get => _matchAsns ?? (_matchAsns = new InputList<string>());
            set => _matchAsns = value;
        }

        [Input("matchCommunitySets")]
        private InputList<string>? _matchCommunitySets;

        /// <summary>
        /// A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        public InputList<string> MatchCommunitySets
        {
            get => _matchCommunitySets ?? (_matchCommunitySets = new InputList<string>());
            set => _matchCommunitySets = value;
        }

        /// <summary>
        /// The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        /// </summary>
        [Input("nextPriority")]
        public Input<int>? NextPriority { get; set; }

        [Input("operateCommunitySets")]
        private InputList<string>? _operateCommunitySets;

        /// <summary>
        /// An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        public InputList<string> OperateCommunitySets
        {
            get => _operateCommunitySets ?? (_operateCommunitySets = new InputList<string>());
            set => _operateCommunitySets = value;
        }

        /// <summary>
        /// An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        /// </summary>
        [Input("preference")]
        public Input<int>? Preference { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("routeTypes")]
        private InputList<string>? _routeTypes;

        /// <summary>
        /// A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        /// </summary>
        public InputList<string> RouteTypes
        {
            get => _routeTypes ?? (_routeTypes = new InputList<string>());
            set => _routeTypes = value;
        }

        [Input("sourceChildInstanceTypes")]
        private InputList<string>? _sourceChildInstanceTypes;

        /// <summary>
        /// A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        public InputList<string> SourceChildInstanceTypes
        {
            get => _sourceChildInstanceTypes ?? (_sourceChildInstanceTypes = new InputList<string>());
            set => _sourceChildInstanceTypes = value;
        }

        [Input("sourceInstanceIds")]
        private InputList<string>? _sourceInstanceIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source instances.
        /// </summary>
        public InputList<string> SourceInstanceIds
        {
            get => _sourceInstanceIds ?? (_sourceInstanceIds = new InputList<string>());
            set => _sourceInstanceIds = value;
        }

        /// <summary>
        /// Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Input("sourceInstanceIdsReverseMatch")]
        public Input<bool>? SourceInstanceIdsReverseMatch { get; set; }

        [Input("sourceRegionIds")]
        private InputList<string>? _sourceRegionIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        /// </summary>
        public InputList<string> SourceRegionIds
        {
            get => _sourceRegionIds ?? (_sourceRegionIds = new InputList<string>());
            set => _sourceRegionIds = value;
        }

        [Input("sourceRouteTableIds")]
        private InputList<string>? _sourceRouteTableIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        public InputList<string> SourceRouteTableIds
        {
            get => _sourceRouteTableIds ?? (_sourceRouteTableIds = new InputList<string>());
            set => _sourceRouteTableIds = value;
        }

        /// <summary>
        /// The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        /// </summary>
        [Input("transmitDirection", required: true)]
        public Input<string> TransmitDirection { get; set; } = null!;

        public RouteMapArgs()
        {
        }
    }

    public sealed class RouteMapState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("asPathMatchMode")]
        public Input<string>? AsPathMatchMode { get; set; }

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The ID of the region to which the CEN instance belongs.
        /// </summary>
        [Input("cenRegionId")]
        public Input<string>? CenRegionId { get; set; }

        /// <summary>
        /// A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("cidrMatchMode")]
        public Input<string>? CidrMatchMode { get; set; }

        /// <summary>
        /// A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        /// </summary>
        [Input("communityMatchMode")]
        public Input<string>? CommunityMatchMode { get; set; }

        /// <summary>
        /// An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        /// </summary>
        [Input("communityOperateMode")]
        public Input<string>? CommunityOperateMode { get; set; }

        /// <summary>
        /// The description of the route map.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationChildInstanceTypes")]
        private InputList<string>? _destinationChildInstanceTypes;

        /// <summary>
        /// A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        public InputList<string> DestinationChildInstanceTypes
        {
            get => _destinationChildInstanceTypes ?? (_destinationChildInstanceTypes = new InputList<string>());
            set => _destinationChildInstanceTypes = value;
        }

        [Input("destinationCidrBlocks")]
        private InputList<string>? _destinationCidrBlocks;

        /// <summary>
        /// A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        /// </summary>
        public InputList<string> DestinationCidrBlocks
        {
            get => _destinationCidrBlocks ?? (_destinationCidrBlocks = new InputList<string>());
            set => _destinationCidrBlocks = value;
        }

        [Input("destinationInstanceIds")]
        private InputList<string>? _destinationInstanceIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination instances.
        /// </summary>
        public InputList<string> DestinationInstanceIds
        {
            get => _destinationInstanceIds ?? (_destinationInstanceIds = new InputList<string>());
            set => _destinationInstanceIds = value;
        }

        /// <summary>
        /// Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Input("destinationInstanceIdsReverseMatch")]
        public Input<bool>? DestinationInstanceIdsReverseMatch { get; set; }

        [Input("destinationRouteTableIds")]
        private InputList<string>? _destinationRouteTableIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        public InputList<string> DestinationRouteTableIds
        {
            get => _destinationRouteTableIds ?? (_destinationRouteTableIds = new InputList<string>());
            set => _destinationRouteTableIds = value;
        }

        /// <summary>
        /// The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        /// </summary>
        [Input("mapResult")]
        public Input<string>? MapResult { get; set; }

        [Input("matchAsns")]
        private InputList<string>? _matchAsns;

        /// <summary>
        /// A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        /// </summary>
        public InputList<string> MatchAsns
        {
            get => _matchAsns ?? (_matchAsns = new InputList<string>());
            set => _matchAsns = value;
        }

        [Input("matchCommunitySets")]
        private InputList<string>? _matchCommunitySets;

        /// <summary>
        /// A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        public InputList<string> MatchCommunitySets
        {
            get => _matchCommunitySets ?? (_matchCommunitySets = new InputList<string>());
            set => _matchCommunitySets = value;
        }

        /// <summary>
        /// The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        /// </summary>
        [Input("nextPriority")]
        public Input<int>? NextPriority { get; set; }

        [Input("operateCommunitySets")]
        private InputList<string>? _operateCommunitySets;

        /// <summary>
        /// An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        /// </summary>
        public InputList<string> OperateCommunitySets
        {
            get => _operateCommunitySets ?? (_operateCommunitySets = new InputList<string>());
            set => _operateCommunitySets = value;
        }

        /// <summary>
        /// An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        /// </summary>
        [Input("preference")]
        public Input<int>? Preference { get; set; }

        [Input("prependAsPaths")]
        private InputList<string>? _prependAsPaths;

        /// <summary>
        /// An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        /// </summary>
        public InputList<string> PrependAsPaths
        {
            get => _prependAsPaths ?? (_prependAsPaths = new InputList<string>());
            set => _prependAsPaths = value;
        }

        /// <summary>
        /// The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("routeMapId")]
        public Input<string>? RouteMapId { get; set; }

        [Input("routeTypes")]
        private InputList<string>? _routeTypes;

        /// <summary>
        /// A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        /// </summary>
        public InputList<string> RouteTypes
        {
            get => _routeTypes ?? (_routeTypes = new InputList<string>());
            set => _routeTypes = value;
        }

        [Input("sourceChildInstanceTypes")]
        private InputList<string>? _sourceChildInstanceTypes;

        /// <summary>
        /// A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        /// </summary>
        public InputList<string> SourceChildInstanceTypes
        {
            get => _sourceChildInstanceTypes ?? (_sourceChildInstanceTypes = new InputList<string>());
            set => _sourceChildInstanceTypes = value;
        }

        [Input("sourceInstanceIds")]
        private InputList<string>? _sourceInstanceIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source instances.
        /// </summary>
        public InputList<string> SourceInstanceIds
        {
            get => _sourceInstanceIds ?? (_sourceInstanceIds = new InputList<string>());
            set => _sourceInstanceIds = value;
        }

        /// <summary>
        /// Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        /// </summary>
        [Input("sourceInstanceIdsReverseMatch")]
        public Input<bool>? SourceInstanceIdsReverseMatch { get; set; }

        [Input("sourceRegionIds")]
        private InputList<string>? _sourceRegionIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        /// </summary>
        public InputList<string> SourceRegionIds
        {
            get => _sourceRegionIds ?? (_sourceRegionIds = new InputList<string>());
            set => _sourceRegionIds = value;
        }

        [Input("sourceRouteTableIds")]
        private InputList<string>? _sourceRouteTableIds;

        /// <summary>
        /// A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        /// </summary>
        public InputList<string> SourceRouteTableIds
        {
            get => _sourceRouteTableIds ?? (_sourceRouteTableIds = new InputList<string>());
            set => _sourceRouteTableIds = value;
        }

        /// <summary>
        /// (Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"].
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        /// </summary>
        [Input("transmitDirection")]
        public Input<string>? TransmitDirection { get; set; }

        public RouteMapState()
        {
        }
    }
}

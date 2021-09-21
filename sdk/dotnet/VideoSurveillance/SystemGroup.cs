// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.VideoSurveillance
{
    /// <summary>
    /// Provides a Video Surveillance System Group resource.
    /// 
    /// For information about Video Surveillance System Group and how to use it, see [What is Group](https://help.aliyun.com/product/108765.html).
    /// 
    /// &gt; **NOTE:** Available in v1.135.0+.
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
    ///         var @default = new AliCloud.VideoSurveillance.SystemGroup("default", new AliCloud.VideoSurveillance.SystemGroupArgs
    ///         {
    ///             GroupName = "your_group_name",
    ///             InProtocol = "rtmp",
    ///             OutProtocol = "flv",
    ///             PlayDomain = "your_plan_domain",
    ///             PushDomain = "your_push_domain",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Video Surveillance System Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:videosurveillance/systemGroup:SystemGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:videosurveillance/systemGroup:SystemGroup")]
    public partial class SystemGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        /// </summary>
        [Output("callback")]
        public Output<string?> Callback { get; private set; } = null!;

        /// <summary>
        /// The capture image.
        /// </summary>
        [Output("captureImage")]
        public Output<int> CaptureImage { get; private set; } = null!;

        /// <summary>
        /// The capture interval.
        /// </summary>
        [Output("captureInterval")]
        public Output<int> CaptureInterval { get; private set; } = null!;

        /// <summary>
        /// The capture oss bucket.
        /// </summary>
        [Output("captureOssBucket")]
        public Output<string> CaptureOssBucket { get; private set; } = null!;

        /// <summary>
        /// The capture oss path.
        /// </summary>
        [Output("captureOssPath")]
        public Output<string> CaptureOssPath { get; private set; } = null!;

        /// <summary>
        /// The capture video.
        /// </summary>
        [Output("captureVideo")]
        public Output<int> CaptureVideo { get; private set; } = null!;

        /// <summary>
        /// The description of Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to open Group.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The Group Name.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        /// </summary>
        [Output("inProtocol")]
        public Output<string> InProtocol { get; private set; } = null!;

        /// <summary>
        /// Whether to enable on-demand streaming. Default value:`false`.
        /// </summary>
        [Output("lazyPull")]
        public Output<bool> LazyPull { get; private set; } = null!;

        /// <summary>
        /// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        /// </summary>
        [Output("outProtocol")]
        public Output<string> OutProtocol { get; private set; } = null!;

        /// <summary>
        /// The domain name of plan streaming used by the group.
        /// </summary>
        [Output("playDomain")]
        public Output<string> PlayDomain { get; private set; } = null!;

        /// <summary>
        /// The domain name of push streaming used by the group.
        /// </summary>
        [Output("pushDomain")]
        public Output<string> PushDomain { get; private set; } = null!;

        /// <summary>
        /// Whether to open Group. Valid values: `on`,`off`.
        /// </summary>
        [Output("status")]
        public Output<bool> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SystemGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemGroup(string name, SystemGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:videosurveillance/systemGroup:SystemGroup", name, args ?? new SystemGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemGroup(string name, Input<string> id, SystemGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:videosurveillance/systemGroup:SystemGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemGroup Get(string name, Input<string> id, SystemGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemGroup(name, id, state, options);
        }
    }

    public sealed class SystemGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        /// </summary>
        [Input("callback")]
        public Input<string>? Callback { get; set; }

        /// <summary>
        /// The description of Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to open Group.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Group Name.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        /// </summary>
        [Input("inProtocol", required: true)]
        public Input<string> InProtocol { get; set; } = null!;

        /// <summary>
        /// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        /// </summary>
        [Input("outProtocol", required: true)]
        public Input<string> OutProtocol { get; set; } = null!;

        /// <summary>
        /// The domain name of plan streaming used by the group.
        /// </summary>
        [Input("playDomain", required: true)]
        public Input<string> PlayDomain { get; set; } = null!;

        /// <summary>
        /// The domain name of push streaming used by the group.
        /// </summary>
        [Input("pushDomain", required: true)]
        public Input<string> PushDomain { get; set; } = null!;

        public SystemGroupArgs()
        {
        }
    }

    public sealed class SystemGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        /// </summary>
        [Input("callback")]
        public Input<string>? Callback { get; set; }

        /// <summary>
        /// The capture image.
        /// </summary>
        [Input("captureImage")]
        public Input<int>? CaptureImage { get; set; }

        /// <summary>
        /// The capture interval.
        /// </summary>
        [Input("captureInterval")]
        public Input<int>? CaptureInterval { get; set; }

        /// <summary>
        /// The capture oss bucket.
        /// </summary>
        [Input("captureOssBucket")]
        public Input<string>? CaptureOssBucket { get; set; }

        /// <summary>
        /// The capture oss path.
        /// </summary>
        [Input("captureOssPath")]
        public Input<string>? CaptureOssPath { get; set; }

        /// <summary>
        /// The capture video.
        /// </summary>
        [Input("captureVideo")]
        public Input<int>? CaptureVideo { get; set; }

        /// <summary>
        /// The description of Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to open Group.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Group Name.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        /// </summary>
        [Input("inProtocol")]
        public Input<string>? InProtocol { get; set; }

        /// <summary>
        /// Whether to enable on-demand streaming. Default value:`false`.
        /// </summary>
        [Input("lazyPull")]
        public Input<bool>? LazyPull { get; set; }

        /// <summary>
        /// The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        /// </summary>
        [Input("outProtocol")]
        public Input<string>? OutProtocol { get; set; }

        /// <summary>
        /// The domain name of plan streaming used by the group.
        /// </summary>
        [Input("playDomain")]
        public Input<string>? PlayDomain { get; set; }

        /// <summary>
        /// The domain name of push streaming used by the group.
        /// </summary>
        [Input("pushDomain")]
        public Input<string>? PushDomain { get; set; }

        /// <summary>
        /// Whether to open Group. Valid values: `on`,`off`.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        public SystemGroupState()
        {
        }
    }
}

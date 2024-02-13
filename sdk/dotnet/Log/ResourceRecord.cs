// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// Log resource is a meta store service provided by log service, resource can be used to define meta store's table structure, record can be used for table's row data.
    /// 
    /// For information about SLS Resource and how to use it, see [Resource management](https://www.alibabacloud.com/help/en/doc-detail/207732.html)
    /// 
    /// &gt; **NOTE:** Available since v1.162.0. log resource region should be set a main region: cn-heyuan.
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
    ///     var exampleResource = new AliCloud.Log.Resource("exampleResource", new()
    ///     {
    ///         Type = "userdefine",
    ///         Description = "user tf resource desc",
    ///         ExtInfo = "{}",
    ///         Schema = @"    {
    ///       ""schema"": [
    ///         {
    ///           ""column"": ""col1"",
    ///           ""desc"": ""col1   desc"",
    ///           ""ext_info"": {
    ///           },
    ///           ""required"": true,
    ///           ""type"": ""string""
    ///         },
    ///         {
    ///           ""column"": ""col2"",
    ///           ""desc"": ""col2   desc"",
    ///           ""ext_info"": ""optional"",
    ///           ""required"": true,
    ///           ""type"": ""string""
    ///         }
    ///       ]
    ///     }
    /// ",
    ///     });
    /// 
    ///     var exampleResourceRecord = new AliCloud.Log.ResourceRecord("exampleResourceRecord", new()
    ///     {
    ///         ResourceName = exampleResource.Id,
    ///         RecordId = "user_tf_resource_1",
    ///         Tag = "resource tag",
    ///         Value = @"    {
    ///       ""col1"": ""this is col1 value"",
    ///       ""col2"": ""col2   value""
    ///     }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log resource record can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:log/resourceRecord:ResourceRecord example &lt;resource_name&gt;:&lt;record_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/resourceRecord:ResourceRecord")]
    public partial class ResourceRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The record's id, should be unique.
        /// </summary>
        [Output("recordId")]
        public Output<string> RecordId { get; private set; } = null!;

        /// <summary>
        /// The name defined in log_resource, log service have some internal resource, like sls.common.user, sls.common.user_group.
        /// </summary>
        [Output("resourceName")]
        public Output<string> ResourceName { get; private set; } = null!;

        /// <summary>
        /// The record's tag, can be used for search.
        /// </summary>
        [Output("tag")]
        public Output<string> Tag { get; private set; } = null!;

        /// <summary>
        /// The json value of record.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceRecord(string name, ResourceRecordArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/resourceRecord:ResourceRecord", name, args ?? new ResourceRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceRecord(string name, Input<string> id, ResourceRecordState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/resourceRecord:ResourceRecord", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceRecord Get(string name, Input<string> id, ResourceRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceRecord(name, id, state, options);
        }
    }

    public sealed class ResourceRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The record's id, should be unique.
        /// </summary>
        [Input("recordId", required: true)]
        public Input<string> RecordId { get; set; } = null!;

        /// <summary>
        /// The name defined in log_resource, log service have some internal resource, like sls.common.user, sls.common.user_group.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// The record's tag, can be used for search.
        /// </summary>
        [Input("tag", required: true)]
        public Input<string> Tag { get; set; } = null!;

        /// <summary>
        /// The json value of record.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ResourceRecordArgs()
        {
        }
        public static new ResourceRecordArgs Empty => new ResourceRecordArgs();
    }

    public sealed class ResourceRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The record's id, should be unique.
        /// </summary>
        [Input("recordId")]
        public Input<string>? RecordId { get; set; }

        /// <summary>
        /// The name defined in log_resource, log service have some internal resource, like sls.common.user, sls.common.user_group.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// The record's tag, can be used for search.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The json value of record.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ResourceRecordState()
        {
        }
        public static new ResourceRecordState Empty => new ResourceRecordState();
    }
}

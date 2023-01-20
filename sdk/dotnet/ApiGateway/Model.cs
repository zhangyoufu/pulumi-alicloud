// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    /// <summary>
    /// Provides a Api Gateway Model resource.
    /// 
    /// For information about Api Gateway Model and how to use it, see [What is Model](https://help.aliyun.com/document_detail/400372.html).
    /// 
    /// &gt; **NOTE:** Available in v1.187.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultGroup = new AliCloud.ApiGateway.Group("defaultGroup", new()
    ///     {
    ///         Description = "example_value",
    ///     });
    /// 
    ///     var defaultModel = new AliCloud.ApiGateway.Model("defaultModel", new()
    ///     {
    ///         GroupId = defaultGroup.Id,
    ///         ModelName = "example_value",
    ///         Schema = "{\"type\":\"object\",\"properties\":{\"id\":{\"format\":\"int64\",\"maximum\":100,\"exclusiveMaximum\":true,\"type\":\"integer\"},\"name\":{\"maxLength\":10,\"type\":\"string\"}}}",
    ///         Description = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api Gateway Model can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:apigateway/model:Model example &lt;group_id&gt;:&lt;model_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/model:Model")]
    public partial class Model : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the model.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The group of the model belongs to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the model.
        /// </summary>
        [Output("modelName")]
        public Output<string> ModelName { get; private set; } = null!;

        /// <summary>
        /// The schema of the model.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;


        /// <summary>
        /// Create a Model resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Model(string name, ModelArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/model:Model", name, args ?? new ModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Model(string name, Input<string> id, ModelState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/model:Model", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Model resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Model Get(string name, Input<string> id, ModelState? state = null, CustomResourceOptions? options = null)
        {
            return new Model(name, id, state, options);
        }
    }

    public sealed class ModelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the model.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The group of the model belongs to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The name of the model.
        /// </summary>
        [Input("modelName", required: true)]
        public Input<string> ModelName { get; set; } = null!;

        /// <summary>
        /// The schema of the model.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        public ModelArgs()
        {
        }
        public static new ModelArgs Empty => new ModelArgs();
    }

    public sealed class ModelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the model.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The group of the model belongs to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the model.
        /// </summary>
        [Input("modelName")]
        public Input<string>? ModelName { get; set; }

        /// <summary>
        /// The schema of the model.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public ModelState()
        {
        }
        public static new ModelState Empty => new ModelState();
    }
}

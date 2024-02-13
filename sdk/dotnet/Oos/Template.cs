// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    /// <summary>
    /// Provides a OOS Template resource. For information about Alicloud OOS Template and how to use it, see [What is Resource Alicloud OOS Template](https://www.alibabacloud.com/help/doc-detail/120761.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.92.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.Oos.Template("example", new()
    ///     {
    ///         Content = @"  {
    ///     ""FormatVersion"": ""OOS-2019-06-01"",
    ///     ""Description"": ""Update Describe instances of given status"",
    ///     ""Parameters"":{
    ///       ""Status"":{
    ///         ""Type"": ""String"",
    ///         ""Description"": ""(Required) The status of the Ecs instance.""
    ///       }
    ///     },
    ///     ""Tasks"": [
    ///       {
    ///         ""Properties"" :{
    ///           ""Parameters"":{
    ///             ""Status"": ""{{ Status }}""
    ///           },
    ///           ""API"": ""DescribeInstances"",
    ///           ""Service"": ""Ecs""
    ///         },
    ///         ""Name"": ""foo"",
    ///         ""Action"": ""ACS::ExecuteApi""
    ///       }]
    ///   }
    ///   
    /// ",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "acceptance Test" },
    ///         },
    ///         TemplateName = "test-name",
    ///         VersionName = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OOS Template can be imported using the id or template_name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oos/template:Template example template_name
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/template:Template")]
    public partial class Template : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When deleting a template, whether to delete its related executions. Default to `false`.
        /// </summary>
        [Output("autoDeleteExecutions")]
        public Output<bool?> AutoDeleteExecutions { get; private set; } = null!;

        /// <summary>
        /// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The creator of the template.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The time when the template is created.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The description of the template.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Is it triggered successfully.
        /// </summary>
        [Output("hasTrigger")]
        public Output<bool> HasTrigger { get; private set; } = null!;

        /// <summary>
        /// The ID of resource group which the template belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        /// </summary>
        [Output("shareType")]
        public Output<string> ShareType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        /// </summary>
        [Output("templateFormat")]
        public Output<string> TemplateFormat { get; private set; } = null!;

        /// <summary>
        /// The id of OOS Template.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        /// </summary>
        [Output("templateType")]
        public Output<string> TemplateType { get; private set; } = null!;

        /// <summary>
        /// The version of OOS Template.
        /// </summary>
        [Output("templateVersion")]
        public Output<string> TemplateVersion { get; private set; } = null!;

        /// <summary>
        /// The user who updated the template.
        /// </summary>
        [Output("updatedBy")]
        public Output<string> UpdatedBy { get; private set; } = null!;

        /// <summary>
        /// The time when the template was updated.
        /// </summary>
        [Output("updatedDate")]
        public Output<string> UpdatedDate { get; private set; } = null!;

        /// <summary>
        /// The name of template version.
        /// </summary>
        [Output("versionName")]
        public Output<string?> VersionName { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oos/template:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/template:Template", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new Template(name, id, state, options);
        }
    }

    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When deleting a template, whether to delete its related executions. Default to `false`.
        /// </summary>
        [Input("autoDeleteExecutions")]
        public Input<bool>? AutoDeleteExecutions { get; set; }

        /// <summary>
        /// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The ID of resource group which the template belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// The name of template version.
        /// </summary>
        [Input("versionName")]
        public Input<string>? VersionName { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }

    public sealed class TemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When deleting a template, whether to delete its related executions. Default to `false`.
        /// </summary>
        [Input("autoDeleteExecutions")]
        public Input<bool>? AutoDeleteExecutions { get; set; }

        /// <summary>
        /// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The creator of the template.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// The time when the template is created.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The description of the template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Is it triggered successfully.
        /// </summary>
        [Input("hasTrigger")]
        public Input<bool>? HasTrigger { get; set; }

        /// <summary>
        /// The ID of resource group which the template belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        /// </summary>
        [Input("shareType")]
        public Input<string>? ShareType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        /// </summary>
        [Input("templateFormat")]
        public Input<string>? TemplateFormat { get; set; }

        /// <summary>
        /// The id of OOS Template.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        /// </summary>
        [Input("templateType")]
        public Input<string>? TemplateType { get; set; }

        /// <summary>
        /// The version of OOS Template.
        /// </summary>
        [Input("templateVersion")]
        public Input<string>? TemplateVersion { get; set; }

        /// <summary>
        /// The user who updated the template.
        /// </summary>
        [Input("updatedBy")]
        public Input<string>? UpdatedBy { get; set; }

        /// <summary>
        /// The time when the template was updated.
        /// </summary>
        [Input("updatedDate")]
        public Input<string>? UpdatedDate { get; set; }

        /// <summary>
        /// The name of template version.
        /// </summary>
        [Input("versionName")]
        public Input<string>? VersionName { get; set; }

        public TemplateState()
        {
        }
        public static new TemplateState Empty => new TemplateState();
    }
}

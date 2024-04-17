// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MaxCompute
{
    public static class GetProjects
    {
        /// <summary>
        /// This data source provides Max Compute Project available to the user.[What is Project](https://help.aliyun.com/document_detail/473479.html)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "tf_testaccmp";
        ///     var defaultProject = new AliCloud.MaxCompute.Project("default", new()
        ///     {
        ///         DefaultQuota = "默认后付费Quota",
        ///         ProjectName = name,
        ///         Comment = name,
        ///         ProductType = "PAYASYOUGO",
        ///     });
        /// 
        ///     var @default = AliCloud.MaxCompute.GetProjects.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultProject.Id,
        ///         },
        ///         NameRegex = defaultProject.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudMaxcomputeProjectExampleId"] = @default.Apply(@default =&gt; @default.Apply(getProjectsResult =&gt; getProjectsResult.Projects[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetProjectsResult> InvokeAsync(GetProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectsResult>("alicloud:maxcompute/getProjects:getProjects", args ?? new GetProjectsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Max Compute Project available to the user.[What is Project](https://help.aliyun.com/document_detail/473479.html)
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var config = new Config();
        ///     var name = config.Get("name") ?? "tf_testaccmp";
        ///     var defaultProject = new AliCloud.MaxCompute.Project("default", new()
        ///     {
        ///         DefaultQuota = "默认后付费Quota",
        ///         ProjectName = name,
        ///         Comment = name,
        ///         ProductType = "PAYASYOUGO",
        ///     });
        /// 
        ///     var @default = AliCloud.MaxCompute.GetProjects.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultProject.Id,
        ///         },
        ///         NameRegex = defaultProject.Name,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudMaxcomputeProjectExampleId"] = @default.Apply(@default =&gt; @default.Apply(getProjectsResult =&gt; getProjectsResult.Projects[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetProjectsResult> Invoke(GetProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectsResult>("alicloud:maxcompute/getProjects:getProjects", args ?? new GetProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Project IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetProjectsArgs()
        {
        }
        public static new GetProjectsArgs Empty => new GetProjectsArgs();
    }

    public sealed class GetProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Project IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetProjectsInvokeArgs()
        {
        }
        public static new GetProjectsInvokeArgs Empty => new GetProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Project IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Projects.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of Project Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectResult> Projects;

        [OutputConstructor]
        private GetProjectsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetProjectsProjectResult> projects)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Projects = projects;
        }
    }
}

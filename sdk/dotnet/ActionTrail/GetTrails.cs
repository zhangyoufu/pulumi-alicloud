// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    public static class GetTrails
    {
        /// <summary>
        /// This data source provides a list of ActionTrail Trails in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.95.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ActionTrail.GetTrails.Invoke(new()
        ///     {
        ///         NameRegex = "tf-testacc-actiontrail",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["trailName"] = @default.Apply(getTrailsResult =&gt; getTrailsResult).Apply(@default =&gt; @default.Apply(getTrailsResult =&gt; getTrailsResult.Trails[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTrailsResult> InvokeAsync(GetTrailsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrailsResult>("alicloud:actiontrail/getTrails:getTrails", args ?? new GetTrailsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of ActionTrail Trails in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.95.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.ActionTrail.GetTrails.Invoke(new()
        ///     {
        ///         NameRegex = "tf-testacc-actiontrail",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["trailName"] = @default.Apply(getTrailsResult =&gt; getTrailsResult).Apply(@default =&gt; @default.Apply(getTrailsResult =&gt; getTrailsResult.Trails[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTrailsResult> Invoke(GetTrailsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrailsResult>("alicloud:actiontrail/getTrails:getTrails", args ?? new GetTrailsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrailsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ActionTrail Trail IDs. It is the same as trail name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Whether to show organization tracking. Default to `false`.
        /// </summary>
        [Input("includeOrganizationTrail")]
        public bool? IncludeOrganizationTrail { get; set; }

        /// <summary>
        /// Whether to show shadow tracking. Default to `false`.
        /// </summary>
        [Input("includeShadowTrails")]
        public bool? IncludeShadowTrails { get; set; }

        /// <summary>
        /// A regex string to filter results by trail name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by status of the ActionTrail Trail. Valid values: `Disable`, `Enable`, `Fresh`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetTrailsArgs()
        {
        }
        public static new GetTrailsArgs Empty => new GetTrailsArgs();
    }

    public sealed class GetTrailsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of ActionTrail Trail IDs. It is the same as trail name.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Whether to show organization tracking. Default to `false`.
        /// </summary>
        [Input("includeOrganizationTrail")]
        public Input<bool>? IncludeOrganizationTrail { get; set; }

        /// <summary>
        /// Whether to show shadow tracking. Default to `false`.
        /// </summary>
        [Input("includeShadowTrails")]
        public Input<bool>? IncludeShadowTrails { get; set; }

        /// <summary>
        /// A regex string to filter results by trail name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by status of the ActionTrail Trail. Valid values: `Disable`, `Enable`, `Fresh`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetTrailsInvokeArgs()
        {
        }
        public static new GetTrailsInvokeArgs Empty => new GetTrailsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrailsResult
    {
        /// <summary>
        /// Field `actiontrails` has been deprecated from version 1.95.0. Use `trails` instead."
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrailsActiontrailResult> Actiontrails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of ActionTrail Trail ids. It is the same as trail name.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IncludeOrganizationTrail;
        public readonly bool? IncludeShadowTrails;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of trail names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the ActionTrail Trail.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A list of ActionTrail Trails. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrailsTrailResult> Trails;

        [OutputConstructor]
        private GetTrailsResult(
            ImmutableArray<Outputs.GetTrailsActiontrailResult> actiontrails,

            string id,

            ImmutableArray<string> ids,

            bool? includeOrganizationTrail,

            bool? includeShadowTrails,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableArray<Outputs.GetTrailsTrailResult> trails)
        {
            Actiontrails = actiontrails;
            Id = id;
            Ids = ids;
            IncludeOrganizationTrail = includeOrganizationTrail;
            IncludeShadowTrails = includeShadowTrails;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Trails = trails;
        }
    }
}

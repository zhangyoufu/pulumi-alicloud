// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh
{
    public static class GetServiceMeshes
    {
        /// <summary>
        /// This data source provides the Service Mesh Service Meshes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.138.0.
        /// </summary>
        public static Task<GetServiceMeshesResult> InvokeAsync(GetServiceMeshesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceMeshesResult>("alicloud:servicemesh/getServiceMeshes:getServiceMeshes", args ?? new GetServiceMeshesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Service Mesh Service Meshes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available since v1.138.0.
        /// </summary>
        public static Output<GetServiceMeshesResult> Invoke(GetServiceMeshesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceMeshesResult>("alicloud:servicemesh/getServiceMeshes:getServiceMeshes", args ?? new GetServiceMeshesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceMeshesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to query the detailed list of resource attributes. Default value: `false`.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Service Mesh IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Service Mesh name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the Service Mesh. Valid values: `running`, `initial`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetServiceMeshesArgs()
        {
        }
        public static new GetServiceMeshesArgs Empty => new GetServiceMeshesArgs();
    }

    public sealed class GetServiceMeshesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to query the detailed list of resource attributes. Default value: `false`.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Service Mesh IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Service Mesh name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the Service Mesh. Valid values: `running`, `initial`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetServiceMeshesInvokeArgs()
        {
        }
        public static new GetServiceMeshesInvokeArgs Empty => new GetServiceMeshesInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceMeshesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of Service Mesh Service Meshes. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceMeshesMeshResult> Meshes;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Service Mesh names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the Service Mesh instance.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetServiceMeshesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetServiceMeshesMeshResult> meshes,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Meshes = meshes;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}

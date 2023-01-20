// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    public static class GetFileSystems
    {
        /// <summary>
        /// This data source provides FileSystems available to the user.
        /// 
        /// &gt; **NOTE**: Available in 1.35.0+
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
        ///     var fs = AliCloud.Nas.GetFileSystems.Invoke(new()
        ///     {
        ///         DescriptionRegex = alicloud_nas_file_system.Foo.Description,
        ///         ProtocolType = "NFS",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudNasFileSystemsId"] = fs.Apply(getFileSystemsResult =&gt; getFileSystemsResult.Systems[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFileSystemsResult> InvokeAsync(GetFileSystemsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemsResult>("alicloud:nas/getFileSystems:getFileSystems", args ?? new GetFileSystemsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides FileSystems available to the user.
        /// 
        /// &gt; **NOTE**: Available in 1.35.0+
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
        ///     var fs = AliCloud.Nas.GetFileSystems.Invoke(new()
        ///     {
        ///         DescriptionRegex = alicloud_nas_file_system.Foo.Description,
        ///         ProtocolType = "NFS",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudNasFileSystemsId"] = fs.Apply(getFileSystemsResult =&gt; getFileSystemsResult.Systems[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFileSystemsResult> Invoke(GetFileSystemsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileSystemsResult>("alicloud:nas/getFileSystems:getFileSystems", args ?? new GetFileSystemsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileSystemsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter the results by the ：FileSystem description.
        /// </summary>
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of FileSystemId.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The protocol type of the file system.
        /// Valid values:
        /// `NFS`,
        /// `SMB` (Available when the `file_system_type` is `standard`).
        /// </summary>
        [Input("protocolType")]
        public string? ProtocolType { get; set; }

        /// <summary>
        /// The storage type of the file system.
        /// * Valid values:
        /// </summary>
        [Input("storageType")]
        public string? StorageType { get; set; }

        public GetFileSystemsArgs()
        {
        }
        public static new GetFileSystemsArgs Empty => new GetFileSystemsArgs();
    }

    public sealed class GetFileSystemsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter the results by the ：FileSystem description.
        /// </summary>
        [Input("descriptionRegex")]
        public Input<string>? DescriptionRegex { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of FileSystemId.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The protocol type of the file system.
        /// Valid values:
        /// `NFS`,
        /// `SMB` (Available when the `file_system_type` is `standard`).
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        /// <summary>
        /// The storage type of the file system.
        /// * Valid values:
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        public GetFileSystemsInvokeArgs()
        {
        }
        public static new GetFileSystemsInvokeArgs Empty => new GetFileSystemsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileSystemsResult
    {
        public readonly string? DescriptionRegex;
        /// <summary>
        /// A list of FileSystem descriptions.
        /// </summary>
        public readonly ImmutableArray<string> Descriptions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of FileSystem Id.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// ProtocolType block of the FileSystem
        /// </summary>
        public readonly string? ProtocolType;
        /// <summary>
        /// StorageType block of the FileSystem.
        /// </summary>
        public readonly string? StorageType;
        /// <summary>
        /// A list of VPCs. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFileSystemsSystemResult> Systems;

        [OutputConstructor]
        private GetFileSystemsResult(
            string? descriptionRegex,

            ImmutableArray<string> descriptions,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? protocolType,

            string? storageType,

            ImmutableArray<Outputs.GetFileSystemsSystemResult> systems)
        {
            DescriptionRegex = descriptionRegex;
            Descriptions = descriptions;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ProtocolType = protocolType;
            StorageType = storageType;
            Systems = systems;
        }
    }
}

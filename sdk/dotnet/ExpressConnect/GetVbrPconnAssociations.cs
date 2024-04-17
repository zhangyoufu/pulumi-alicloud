// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ExpressConnect
{
    public static class GetVbrPconnAssociations
    {
        /// <summary>
        /// This data source provides Express Connect Vbr Pconn Association available to the user.
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
        ///     var @default = AliCloud.ExpressConnect.GetVbrPconnAssociations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         VbrId = defaultAlicloudExpressConnectVbrPconnAssociation.VbrId,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudExpressConnectVbrPconnAssociationExampleId"] = @default.Apply(@default =&gt; @default.Apply(getVbrPconnAssociationsResult =&gt; getVbrPconnAssociationsResult.Associations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVbrPconnAssociationsResult> InvokeAsync(GetVbrPconnAssociationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVbrPconnAssociationsResult>("alicloud:expressconnect/getVbrPconnAssociations:getVbrPconnAssociations", args ?? new GetVbrPconnAssociationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Express Connect Vbr Pconn Association available to the user.
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
        ///     var @default = AliCloud.ExpressConnect.GetVbrPconnAssociations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         VbrId = defaultAlicloudExpressConnectVbrPconnAssociation.VbrId,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudExpressConnectVbrPconnAssociationExampleId"] = @default.Apply(@default =&gt; @default.Apply(getVbrPconnAssociationsResult =&gt; getVbrPconnAssociationsResult.Associations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVbrPconnAssociationsResult> Invoke(GetVbrPconnAssociationsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVbrPconnAssociationsResult>("alicloud:expressconnect/getVbrPconnAssociations:getVbrPconnAssociations", args ?? new GetVbrPconnAssociationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVbrPconnAssociationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Vbr Pconn Association IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The ID of the VBR instance.
        /// </summary>
        [Input("vbrId")]
        public string? VbrId { get; set; }

        public GetVbrPconnAssociationsArgs()
        {
        }
        public static new GetVbrPconnAssociationsArgs Empty => new GetVbrPconnAssociationsArgs();
    }

    public sealed class GetVbrPconnAssociationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Vbr Pconn Association IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The ID of the VBR instance.
        /// </summary>
        [Input("vbrId")]
        public Input<string>? VbrId { get; set; }

        public GetVbrPconnAssociationsInvokeArgs()
        {
        }
        public static new GetVbrPconnAssociationsInvokeArgs Empty => new GetVbrPconnAssociationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVbrPconnAssociationsResult
    {
        /// <summary>
        /// A list of Vbr Pconn Association Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVbrPconnAssociationsAssociationResult> Associations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        /// <summary>
        /// The ID of the VBR instance.
        /// </summary>
        public readonly string? VbrId;

        [OutputConstructor]
        private GetVbrPconnAssociationsResult(
            ImmutableArray<Outputs.GetVbrPconnAssociationsAssociationResult> associations,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? vbrId)
        {
            Associations = associations;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            VbrId = vbrId;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms
{
    public static class GetUserTenants
    {
        /// <summary>
        /// This data source provides a list of DMS User Tenants in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.161.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Dms.GetUserTenants.Invoke(new()
        ///     {
        ///         Status = "ACTIVE",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["tid"] = @default.Apply(@default =&gt; @default.Apply(getUserTenantsResult =&gt; getUserTenantsResult.Ids[0])),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserTenantsResult> InvokeAsync(GetUserTenantsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserTenantsResult>("alicloud:dms/getUserTenants:getUserTenants", args ?? new GetUserTenantsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of DMS User Tenants in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.161.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Dms.GetUserTenants.Invoke(new()
        ///     {
        ///         Status = "ACTIVE",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["tid"] = @default.Apply(@default =&gt; @default.Apply(getUserTenantsResult =&gt; getUserTenantsResult.Ids[0])),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserTenantsResult> Invoke(GetUserTenantsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserTenantsResult>("alicloud:dms/getUserTenants:getUserTenants", args ?? new GetUserTenantsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserTenantsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of DMS User Tenant IDs (TID).
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

        /// <summary>
        /// The status of the user tenant.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetUserTenantsArgs()
        {
        }
        public static new GetUserTenantsArgs Empty => new GetUserTenantsArgs();
    }

    public sealed class GetUserTenantsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of DMS User Tenant IDs (TID).
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

        /// <summary>
        /// The status of the user tenant.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetUserTenantsInvokeArgs()
        {
        }
        public static new GetUserTenantsInvokeArgs Empty => new GetUserTenantsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserTenantsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of DMS User Tenant IDs (UID).
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of DMS User Tenant names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the user tenant.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A list of DMS User Tenants. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserTenantsTenantResult> Tenants;

        [OutputConstructor]
        private GetUserTenantsResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableArray<Outputs.GetUserTenantsTenantResult> tenants)
        {
            Id = id;
            Ids = ids;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Tenants = tenants;
        }
    }
}

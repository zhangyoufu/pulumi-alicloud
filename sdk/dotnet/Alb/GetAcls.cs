// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    public static class GetAcls
    {
        /// <summary>
        /// This data source provides the Application Load Balancer (ALB) Acls of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.133.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.Alb.GetAcls.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Alb.GetAcls.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Acl",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["albAclId1"] = ids.Apply(getAclsResult =&gt; getAclsResult.Acls[0]?.Id),
        ///         ["albAclId2"] = nameRegex.Apply(getAclsResult =&gt; getAclsResult.Acls[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAclsResult> InvokeAsync(GetAclsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAclsResult>("alicloud:alb/getAcls:getAcls", args ?? new GetAclsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Application Load Balancer (ALB) Acls of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.133.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.Alb.GetAcls.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Alb.GetAcls.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Acl",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["albAclId1"] = ids.Apply(getAclsResult =&gt; getAclsResult.Acls[0]?.Id),
        ///         ["albAclId2"] = nameRegex.Apply(getAclsResult =&gt; getAclsResult.Acls[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAclsResult> Invoke(GetAclsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclsResult>("alicloud:alb/getAcls:getAcls", args ?? new GetAclsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAclsArgs : global::Pulumi.InvokeArgs
    {
        [Input("aclIds")]
        private List<string>? _aclIds;

        /// <summary>
        /// The acl ids.
        /// </summary>
        public List<string> AclIds
        {
            get => _aclIds ?? (_aclIds = new List<string>());
            set => _aclIds = value;
        }

        /// <summary>
        /// The ACL Name.
        /// </summary>
        [Input("aclName")]
        public string? AclName { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Acl IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Acl name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Resource Group to Which the Number.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetAclsArgs()
        {
        }
        public static new GetAclsArgs Empty => new GetAclsArgs();
    }

    public sealed class GetAclsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("aclIds")]
        private InputList<string>? _aclIds;

        /// <summary>
        /// The acl ids.
        /// </summary>
        public InputList<string> AclIds
        {
            get => _aclIds ?? (_aclIds = new InputList<string>());
            set => _aclIds = value;
        }

        /// <summary>
        /// The ACL Name.
        /// </summary>
        [Input("aclName")]
        public Input<string>? AclName { get; set; }

        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Acl IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Acl name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Resource Group to Which the Number.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`. `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetAclsInvokeArgs()
        {
        }
        public static new GetAclsInvokeArgs Empty => new GetAclsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAclsResult
    {
        public readonly ImmutableArray<string> AclIds;
        public readonly string? AclName;
        public readonly ImmutableArray<Outputs.GetAclsAclResult> Acls;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly string? Status;

        [OutputConstructor]
        private GetAclsResult(
            ImmutableArray<string> aclIds,

            string? aclName,

            ImmutableArray<Outputs.GetAclsAclResult> acls,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            string? status)
        {
            AclIds = aclIds;
            AclName = aclName;
            Acls = acls;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Status = status;
        }
    }
}

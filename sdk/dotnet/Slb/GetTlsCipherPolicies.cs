// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static class GetTlsCipherPolicies
    {
        /// <summary>
        /// This data source provides the Slb Tls Cipher Policies of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.Slb.GetTlsCipherPolicies.InvokeAsync(new AliCloud.Slb.GetTlsCipherPoliciesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value-1",
        ///                 "example_value-2",
        ///             },
        ///         }));
        ///         this.SlbTlsCipherPolicyId1 = ids.Apply(ids =&gt; ids.Policies[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.Slb.GetTlsCipherPolicies.InvokeAsync(new AliCloud.Slb.GetTlsCipherPoliciesArgs
        ///         {
        ///             NameRegex = "^My-TlsCipherPolicy",
        ///         }));
        ///         this.SlbTlsCipherPolicyId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Policies[0].Id);
        ///     }
        /// 
        ///     [Output("slbTlsCipherPolicyId1")]
        ///     public Output&lt;string&gt; SlbTlsCipherPolicyId1 { get; set; }
        ///     [Output("slbTlsCipherPolicyId2")]
        ///     public Output&lt;string&gt; SlbTlsCipherPolicyId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTlsCipherPoliciesResult> InvokeAsync(GetTlsCipherPoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTlsCipherPoliciesResult>("alicloud:slb/getTlsCipherPolicies:getTlsCipherPolicies", args ?? new GetTlsCipherPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetTlsCipherPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Tls Cipher Policy IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The include listener.
        /// </summary>
        [Input("includeListener")]
        public bool? IncludeListener { get; set; }

        /// <summary>
        /// A regex string to filter results by Tls Cipher Policy name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// TLS policy instance state.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
        /// </summary>
        [Input("tlsCipherPolicyName")]
        public string? TlsCipherPolicyName { get; set; }

        public GetTlsCipherPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTlsCipherPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IncludeListener;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetTlsCipherPoliciesPolicyResult> Policies;
        public readonly string? Status;
        public readonly string? TlsCipherPolicyName;

        [OutputConstructor]
        private GetTlsCipherPoliciesResult(
            string id,

            ImmutableArray<string> ids,

            bool? includeListener,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetTlsCipherPoliciesPolicyResult> policies,

            string? status,

            string? tlsCipherPolicyName)
        {
            Id = id;
            Ids = ids;
            IncludeListener = includeListener;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Policies = policies;
            Status = status;
            TlsCipherPolicyName = tlsCipherPolicyName;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    public static class GetRegistryEnterpriseInstances
    {
        /// <summary>
        /// This data source provides a list Container Registry Enterprise Edition instances on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.86.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myInstances = Output.Create(AliCloud.CS.GetRegistryEnterpriseInstances.InvokeAsync(new AliCloud.CS.GetRegistryEnterpriseInstancesArgs
        ///         {
        ///             NameRegex = "my-instances",
        ///             OutputFile = "my-instances-json",
        ///         }));
        ///         this.Output = myInstances.Apply(myInstances =&gt; myInstances.Instances);
        ///     }
        /// 
        ///     [Output("output")]
        ///     public Output&lt;string&gt; Output { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryEnterpriseInstancesResult> InvokeAsync(GetRegistryEnterpriseInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryEnterpriseInstancesResult>("alicloud:cs/getRegistryEnterpriseInstances:getRegistryEnterpriseInstances", args ?? new GetRegistryEnterpriseInstancesArgs(), options.WithVersion());
    }


    public sealed class GetRegistryEnterpriseInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ids to filter results by instance id.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetRegistryEnterpriseInstancesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryEnterpriseInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegistryEnterpriseInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of instance names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetRegistryEnterpriseInstancesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetRegistryEnterpriseInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}

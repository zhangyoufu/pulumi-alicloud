// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    public static class GetFunctions
    {
        /// <summary>
        /// This data source provides the Function Compute functions of the current Alibaba Cloud user.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionsResult> InvokeAsync(GetFunctionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionsResult>("alicloud:fc/getFunctions:getFunctions", args ?? new GetFunctionsArgs(), options.WithVersion());
    }


    public sealed class GetFunctionsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// - A list of functions ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by function name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Name of the service that contains the functions to find.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetFunctionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionsResult
    {
        /// <summary>
        /// A list of functions. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionsFunctionResult> Functions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of functions ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of functions names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetFunctionsResult(
            ImmutableArray<Outputs.GetFunctionsFunctionResult> functions,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string serviceName)
        {
            Functions = functions;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ServiceName = serviceName;
        }
    }
}

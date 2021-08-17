// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    /// <summary>
    /// Provides a DCDN Accelerated Domain resource.
    /// 
    /// For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/zh/doc-detail/130632.htm)
    /// 
    /// &gt; **NOTE:** Available in v1.131.0+.
    /// 
    /// ## Import
    /// 
    /// DCDN domain config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dcdn/domainConfig:DomainConfig example &lt;domain_name&gt;:&lt;function_name&gt;:&lt;config_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/domainConfig:DomainConfig")]
    public partial class DomainConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The DCDN domain config id.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The args of the domain config.
        /// </summary>
        [Output("functionArgs")]
        public Output<ImmutableArray<Outputs.DomainConfigFunctionArg>> FunctionArgs { get; private set; } = null!;

        /// <summary>
        /// The name of the domain config.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// The status of this resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DomainConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainConfig(string name, DomainConfigArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/domainConfig:DomainConfig", name, args ?? new DomainConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainConfig(string name, Input<string> id, DomainConfigState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/domainConfig:DomainConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainConfig Get(string name, Input<string> id, DomainConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainConfig(name, id, state, options);
        }
    }

    public sealed class DomainConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("functionArgs", required: true)]
        private InputList<Inputs.DomainConfigFunctionArgArgs>? _functionArgs;

        /// <summary>
        /// The args of the domain config.
        /// </summary>
        public InputList<Inputs.DomainConfigFunctionArgArgs> FunctionArgs
        {
            get => _functionArgs ?? (_functionArgs = new InputList<Inputs.DomainConfigFunctionArgArgs>());
            set => _functionArgs = value;
        }

        /// <summary>
        /// The name of the domain config.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        public DomainConfigArgs()
        {
        }
    }

    public sealed class DomainConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DCDN domain config id.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("functionArgs")]
        private InputList<Inputs.DomainConfigFunctionArgGetArgs>? _functionArgs;

        /// <summary>
        /// The args of the domain config.
        /// </summary>
        public InputList<Inputs.DomainConfigFunctionArgGetArgs> FunctionArgs
        {
            get => _functionArgs ?? (_functionArgs = new InputList<Inputs.DomainConfigFunctionArgGetArgs>());
            set => _functionArgs = value;
        }

        /// <summary>
        /// The name of the domain config.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// The status of this resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DomainConfigState()
        {
        }
    }
}

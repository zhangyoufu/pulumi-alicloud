// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Provides a RDS Parameter Group resource.
    /// 
    /// For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.119.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var @default = new AliCloud.Rds.RdsParameterGroup("default", new()
    ///     {
    ///         Engine = "mysql",
    ///         EngineVersion = "5.7",
    ///         ParamDetails = new[]
    ///         {
    ///             new AliCloud.Rds.Inputs.RdsParameterGroupParamDetailArgs
    ///             {
    ///                 ParamName = "back_log",
    ///                 ParamValue = "4000",
    ///             },
    ///             new AliCloud.Rds.Inputs.RdsParameterGroupParamDetailArgs
    ///             {
    ///                 ParamName = "wait_timeout",
    ///                 ParamValue = "86460",
    ///             },
    ///         },
    ///         ParameterGroupDesc = name,
    ///         ParameterGroupName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS Parameter Group can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rds/rdsParameterGroup:RdsParameterGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/rdsParameterGroup:RdsParameterGroup")]
    public partial class RdsParameterGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Parameter list. See `param_detail` below.
        /// </summary>
        [Output("paramDetails")]
        public Output<ImmutableArray<Outputs.RdsParameterGroupParamDetail>> ParamDetails { get; private set; } = null!;

        /// <summary>
        /// The description of the parameter template.
        /// </summary>
        [Output("parameterGroupDesc")]
        public Output<string?> ParameterGroupDesc { get; private set; } = null!;

        /// <summary>
        /// The name of the parameter template.
        /// </summary>
        [Output("parameterGroupName")]
        public Output<string> ParameterGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a RdsParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsParameterGroup(string name, RdsParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, args ?? new RdsParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsParameterGroup(string name, Input<string> id, RdsParameterGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsParameterGroup:RdsParameterGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdsParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsParameterGroup Get(string name, Input<string> id, RdsParameterGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsParameterGroup(name, id, state, options);
        }
    }

    public sealed class RdsParameterGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        [Input("paramDetails", required: true)]
        private InputList<Inputs.RdsParameterGroupParamDetailArgs>? _paramDetails;

        /// <summary>
        /// Parameter list. See `param_detail` below.
        /// </summary>
        public InputList<Inputs.RdsParameterGroupParamDetailArgs> ParamDetails
        {
            get => _paramDetails ?? (_paramDetails = new InputList<Inputs.RdsParameterGroupParamDetailArgs>());
            set => _paramDetails = value;
        }

        /// <summary>
        /// The description of the parameter template.
        /// </summary>
        [Input("parameterGroupDesc")]
        public Input<string>? ParameterGroupDesc { get; set; }

        /// <summary>
        /// The name of the parameter template.
        /// </summary>
        [Input("parameterGroupName", required: true)]
        public Input<string> ParameterGroupName { get; set; } = null!;

        public RdsParameterGroupArgs()
        {
        }
        public static new RdsParameterGroupArgs Empty => new RdsParameterGroupArgs();
    }

    public sealed class RdsParameterGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        [Input("paramDetails")]
        private InputList<Inputs.RdsParameterGroupParamDetailGetArgs>? _paramDetails;

        /// <summary>
        /// Parameter list. See `param_detail` below.
        /// </summary>
        public InputList<Inputs.RdsParameterGroupParamDetailGetArgs> ParamDetails
        {
            get => _paramDetails ?? (_paramDetails = new InputList<Inputs.RdsParameterGroupParamDetailGetArgs>());
            set => _paramDetails = value;
        }

        /// <summary>
        /// The description of the parameter template.
        /// </summary>
        [Input("parameterGroupDesc")]
        public Input<string>? ParameterGroupDesc { get; set; }

        /// <summary>
        /// The name of the parameter template.
        /// </summary>
        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        public RdsParameterGroupState()
        {
        }
        public static new RdsParameterGroupState Empty => new RdsParameterGroupState();
    }
}

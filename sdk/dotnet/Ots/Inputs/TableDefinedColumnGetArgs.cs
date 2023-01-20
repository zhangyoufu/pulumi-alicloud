// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ots.Inputs
{

    public sealed class TableDefinedColumnGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name for defined column.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Type for defined column. `Integer`, `String`, `Binary`, `Double`, `Boolean` is allowed.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TableDefinedColumnGetArgs()
        {
        }
        public static new TableDefinedColumnGetArgs Empty => new TableDefinedColumnGetArgs();
    }
}

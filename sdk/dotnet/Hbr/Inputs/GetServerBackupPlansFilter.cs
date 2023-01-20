// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Inputs
{

    public sealed class GetServerBackupPlansFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The key of the field to filter. Valid values: `planId`, `instanceId`, `planName`.
        /// </summary>
        [Input("key")]
        public string? Key { get; set; }

        [Input("values")]
        private List<string>? _values;

        /// <summary>
        /// Set of values that are accepted for the given field.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetServerBackupPlansFilterArgs()
        {
        }
        public static new GetServerBackupPlansFilterArgs Empty => new GetServerBackupPlansFilterArgs();
    }
}

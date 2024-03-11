// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class ManagedKubernetesMaintenanceWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maintenance time, values range from 1 to 24,unit is hour. For example: "3h".
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// Whether to open the maintenance window. The following parameters take effect only `enable = true`.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// Initial maintenance time, For example:"03:00:00Z".
        /// </summary>
        [Input("maintenanceTime", required: true)]
        public Input<string> MaintenanceTime { get; set; } = null!;

        /// <summary>
        /// Maintenance cycle, you can set the values from Monday to Sunday, separated by commas when the values are multiple. The default is Thursday.
        /// 
        /// for example:
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        [Input("weeklyPeriod", required: true)]
        public Input<string> WeeklyPeriod { get; set; } = null!;

        public ManagedKubernetesMaintenanceWindowArgs()
        {
        }
        public static new ManagedKubernetesMaintenanceWindowArgs Empty => new ManagedKubernetesMaintenanceWindowArgs();
    }
}

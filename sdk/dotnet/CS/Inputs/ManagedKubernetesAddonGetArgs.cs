// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class ManagedKubernetesAddonGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If this parameter is left empty, no configurations are required.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// It specifies whether to disable automatic installation. 
        /// 
        /// It is a new field since 1.75.0. You can specific network plugin, log component,ingress component and so on.
        /// 
        /// You can get more information about addons on ACK web console. When you create a ACK cluster. You can get openapi-spec before creating the cluster on submission page.
        /// 
        /// `logtail-ds` - You can specify `IngressDashboardEnabled` and `sls_project_name` in config. If you switch on `IngressDashboardEnabled` and `sls_project_name`,then logtail-ds would use `sls_project_name` as default log store.
        /// 
        /// `nginx-ingress-controller` - You can specific `IngressSlbNetworkType` in config. Options: internet|intranet.
        /// 
        /// The `main.tf`:
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// This parameter specifies the name of the component.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// It specifies the version of the component.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ManagedKubernetesAddonGetArgs()
        {
        }
        public static new ManagedKubernetesAddonGetArgs Empty => new ManagedKubernetesAddonGetArgs();
    }
}

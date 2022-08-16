// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class NodePoolKubeletConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Same as cpuManagerPolicy. The name of the policy to use. Requires the CPUManager feature gate to be enabled. Valid value is `none` or `static`.
        /// </summary>
        [Input("cpuManagerPolicy")]
        public Input<string>? CpuManagerPolicy { get; set; }

        /// <summary>
        /// Same as eventBurst. The maximum size of a burst of event creations, temporarily allows event creations to burst to this number, while still not exceeding `event_record_qps`. It is only used when `event_record_qps` is greater than 0. Valid value is `[0-100]`.
        /// </summary>
        [Input("eventBurst")]
        public Input<string>? EventBurst { get; set; }

        /// <summary>
        /// Same as eventRecordQPS. The maximum event creations per second. If 0, there is no limit enforced. Valid value is `[0-50]`.
        /// </summary>
        [Input("eventRecordQps")]
        public Input<string>? EventRecordQps { get; set; }

        [Input("evictionHard")]
        private InputMap<object>? _evictionHard;

        /// <summary>
        /// Same as evictionHard. The map of signal names to quantities that defines hard eviction thresholds. For example: `{"memory.available" = "300Mi"}`.
        /// </summary>
        public InputMap<object> EvictionHard
        {
            get => _evictionHard ?? (_evictionHard = new InputMap<object>());
            set => _evictionHard = value;
        }

        [Input("evictionSoft")]
        private InputMap<object>? _evictionSoft;

        /// <summary>
        /// Same as evictionSoft. The map of signal names to quantities that defines soft eviction thresholds. For example: `{"memory.available" = "300Mi"}`.
        /// </summary>
        public InputMap<object> EvictionSoft
        {
            get => _evictionSoft ?? (_evictionSoft = new InputMap<object>());
            set => _evictionSoft = value;
        }

        [Input("evictionSoftGracePeriod")]
        private InputMap<object>? _evictionSoftGracePeriod;

        /// <summary>
        /// Same as evictionSoftGracePeriod. The map of signal names to quantities that defines grace periods for each soft eviction signal. For example: `{"memory.available" = "30s"}`.
        /// </summary>
        public InputMap<object> EvictionSoftGracePeriod
        {
            get => _evictionSoftGracePeriod ?? (_evictionSoftGracePeriod = new InputMap<object>());
            set => _evictionSoftGracePeriod = value;
        }

        /// <summary>
        /// Same as kubeAPIBurst. The burst to allow while talking with kubernetes api-server. Valid value is `[0-100]`.
        /// </summary>
        [Input("kubeApiBurst")]
        public Input<string>? KubeApiBurst { get; set; }

        /// <summary>
        /// Same as kubeAPIQPS. The QPS to use while talking with kubernetes api-server. Valid value is `[0-50]`.
        /// </summary>
        [Input("kubeApiQps")]
        public Input<string>? KubeApiQps { get; set; }

        [Input("kubeReserved")]
        private InputMap<object>? _kubeReserved;

        /// <summary>
        /// Same as kubeReserved. The set of ResourceName=ResourceQuantity (e.g. cpu=200m,memory=150G) pairs that describe resources reserved for kubernetes system components. Currently, cpu, memory and local storage for root file system are supported. See [compute resources](http://kubernetes.io/docs/user-guide/compute-resources) for more details.
        /// </summary>
        public InputMap<object> KubeReserved
        {
            get => _kubeReserved ?? (_kubeReserved = new InputMap<object>());
            set => _kubeReserved = value;
        }

        /// <summary>
        /// Same as registryBurst. The maximum size of burst pulls, temporarily allows pulls to burst to this number, while still not exceeding `registry_pull_qps`. Only used if `registry_pull_qps` is greater than 0. Valid value is `[0-100]`.
        /// </summary>
        [Input("registryBurst")]
        public Input<string>? RegistryBurst { get; set; }

        /// <summary>
        /// Same as registryPullQPS. The limit of registry pulls per second. Setting it to `0` means no limit. Valid value is `[0-50]`.
        /// </summary>
        [Input("registryPullQps")]
        public Input<string>? RegistryPullQps { get; set; }

        /// <summary>
        /// Same as serializeImagePulls. When enabled, it tells the Kubelet to pull images one at a time. We recommend not changing the default value on nodes that run docker daemon with version &lt; 1.9 or an Aufs storage backend. Valid value is `true` or `false`.
        /// </summary>
        [Input("serializeImagePulls")]
        public Input<string>? SerializeImagePulls { get; set; }

        [Input("systemReserved")]
        private InputMap<object>? _systemReserved;

        /// <summary>
        /// Same as systemReserved. The set of ResourceName=ResourceQuantity (e.g. cpu=200m,memory=150G) pairs that describe resources reserved for non-kubernetes components. Currently, only cpu and memory are supported. See [compute resources](http://kubernetes.io/docs/user-guide/compute-resources) for more details.
        /// </summary>
        public InputMap<object> SystemReserved
        {
            get => _systemReserved ?? (_systemReserved = new InputMap<object>());
            set => _systemReserved = value;
        }

        public NodePoolKubeletConfigurationGetArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationArgs, ApplicationState } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { AutoscalingConfigArgs, AutoscalingConfigState } from "./autoscalingConfig";
export type AutoscalingConfig = import("./autoscalingConfig").AutoscalingConfig;
export const AutoscalingConfig: typeof import("./autoscalingConfig").AutoscalingConfig = null as any;
utilities.lazyLoad(exports, ["AutoscalingConfig"], () => require("./autoscalingConfig"));

export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { EdgeKubernetesArgs, EdgeKubernetesState } from "./edgeKubernetes";
export type EdgeKubernetes = import("./edgeKubernetes").EdgeKubernetes;
export const EdgeKubernetes: typeof import("./edgeKubernetes").EdgeKubernetes = null as any;
utilities.lazyLoad(exports, ["EdgeKubernetes"], () => require("./edgeKubernetes"));

export { GetAckServiceArgs, GetAckServiceResult, GetAckServiceOutputArgs } from "./getAckService";
export const getAckService: typeof import("./getAckService").getAckService = null as any;
export const getAckServiceOutput: typeof import("./getAckService").getAckServiceOutput = null as any;
utilities.lazyLoad(exports, ["getAckService","getAckServiceOutput"], () => require("./getAckService"));

export { GetClusterCredentialArgs, GetClusterCredentialResult, GetClusterCredentialOutputArgs } from "./getClusterCredential";
export const getClusterCredential: typeof import("./getClusterCredential").getClusterCredential = null as any;
export const getClusterCredentialOutput: typeof import("./getClusterCredential").getClusterCredentialOutput = null as any;
utilities.lazyLoad(exports, ["getClusterCredential","getClusterCredentialOutput"], () => require("./getClusterCredential"));

export { GetEdgeKubernetesClustersArgs, GetEdgeKubernetesClustersResult, GetEdgeKubernetesClustersOutputArgs } from "./getEdgeKubernetesClusters";
export const getEdgeKubernetesClusters: typeof import("./getEdgeKubernetesClusters").getEdgeKubernetesClusters = null as any;
export const getEdgeKubernetesClustersOutput: typeof import("./getEdgeKubernetesClusters").getEdgeKubernetesClustersOutput = null as any;
utilities.lazyLoad(exports, ["getEdgeKubernetesClusters","getEdgeKubernetesClustersOutput"], () => require("./getEdgeKubernetesClusters"));

export { GetKubernetesAddonMetadataArgs, GetKubernetesAddonMetadataResult, GetKubernetesAddonMetadataOutputArgs } from "./getKubernetesAddonMetadata";
export const getKubernetesAddonMetadata: typeof import("./getKubernetesAddonMetadata").getKubernetesAddonMetadata = null as any;
export const getKubernetesAddonMetadataOutput: typeof import("./getKubernetesAddonMetadata").getKubernetesAddonMetadataOutput = null as any;
utilities.lazyLoad(exports, ["getKubernetesAddonMetadata","getKubernetesAddonMetadataOutput"], () => require("./getKubernetesAddonMetadata"));

export { GetKubernetesAddonsArgs, GetKubernetesAddonsResult, GetKubernetesAddonsOutputArgs } from "./getKubernetesAddons";
export const getKubernetesAddons: typeof import("./getKubernetesAddons").getKubernetesAddons = null as any;
export const getKubernetesAddonsOutput: typeof import("./getKubernetesAddons").getKubernetesAddonsOutput = null as any;
utilities.lazyLoad(exports, ["getKubernetesAddons","getKubernetesAddonsOutput"], () => require("./getKubernetesAddons"));

export { GetKubernetesClustersArgs, GetKubernetesClustersResult, GetKubernetesClustersOutputArgs } from "./getKubernetesClusters";
export const getKubernetesClusters: typeof import("./getKubernetesClusters").getKubernetesClusters = null as any;
export const getKubernetesClustersOutput: typeof import("./getKubernetesClusters").getKubernetesClustersOutput = null as any;
utilities.lazyLoad(exports, ["getKubernetesClusters","getKubernetesClustersOutput"], () => require("./getKubernetesClusters"));

export { GetKubernetesPermissionArgs, GetKubernetesPermissionResult, GetKubernetesPermissionOutputArgs } from "./getKubernetesPermission";
export const getKubernetesPermission: typeof import("./getKubernetesPermission").getKubernetesPermission = null as any;
export const getKubernetesPermissionOutput: typeof import("./getKubernetesPermission").getKubernetesPermissionOutput = null as any;
utilities.lazyLoad(exports, ["getKubernetesPermission","getKubernetesPermissionOutput"], () => require("./getKubernetesPermission"));

export { GetKubernetesVersionArgs, GetKubernetesVersionResult, GetKubernetesVersionOutputArgs } from "./getKubernetesVersion";
export const getKubernetesVersion: typeof import("./getKubernetesVersion").getKubernetesVersion = null as any;
export const getKubernetesVersionOutput: typeof import("./getKubernetesVersion").getKubernetesVersionOutput = null as any;
utilities.lazyLoad(exports, ["getKubernetesVersion","getKubernetesVersionOutput"], () => require("./getKubernetesVersion"));

export { GetManagedKubernetesClustersArgs, GetManagedKubernetesClustersResult, GetManagedKubernetesClustersOutputArgs } from "./getManagedKubernetesClusters";
export const getManagedKubernetesClusters: typeof import("./getManagedKubernetesClusters").getManagedKubernetesClusters = null as any;
export const getManagedKubernetesClustersOutput: typeof import("./getManagedKubernetesClusters").getManagedKubernetesClustersOutput = null as any;
utilities.lazyLoad(exports, ["getManagedKubernetesClusters","getManagedKubernetesClustersOutput"], () => require("./getManagedKubernetesClusters"));

export { GetRegistryEnterpriseInstancesArgs, GetRegistryEnterpriseInstancesResult, GetRegistryEnterpriseInstancesOutputArgs } from "./getRegistryEnterpriseInstances";
export const getRegistryEnterpriseInstances: typeof import("./getRegistryEnterpriseInstances").getRegistryEnterpriseInstances = null as any;
export const getRegistryEnterpriseInstancesOutput: typeof import("./getRegistryEnterpriseInstances").getRegistryEnterpriseInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryEnterpriseInstances","getRegistryEnterpriseInstancesOutput"], () => require("./getRegistryEnterpriseInstances"));

export { GetRegistryEnterpriseNamespacesArgs, GetRegistryEnterpriseNamespacesResult, GetRegistryEnterpriseNamespacesOutputArgs } from "./getRegistryEnterpriseNamespaces";
export const getRegistryEnterpriseNamespaces: typeof import("./getRegistryEnterpriseNamespaces").getRegistryEnterpriseNamespaces = null as any;
export const getRegistryEnterpriseNamespacesOutput: typeof import("./getRegistryEnterpriseNamespaces").getRegistryEnterpriseNamespacesOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryEnterpriseNamespaces","getRegistryEnterpriseNamespacesOutput"], () => require("./getRegistryEnterpriseNamespaces"));

export { GetRegistryEnterpriseReposArgs, GetRegistryEnterpriseReposResult, GetRegistryEnterpriseReposOutputArgs } from "./getRegistryEnterpriseRepos";
export const getRegistryEnterpriseRepos: typeof import("./getRegistryEnterpriseRepos").getRegistryEnterpriseRepos = null as any;
export const getRegistryEnterpriseReposOutput: typeof import("./getRegistryEnterpriseRepos").getRegistryEnterpriseReposOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryEnterpriseRepos","getRegistryEnterpriseReposOutput"], () => require("./getRegistryEnterpriseRepos"));

export { GetRegistryEnterpriseSyncRulesArgs, GetRegistryEnterpriseSyncRulesResult, GetRegistryEnterpriseSyncRulesOutputArgs } from "./getRegistryEnterpriseSyncRules";
export const getRegistryEnterpriseSyncRules: typeof import("./getRegistryEnterpriseSyncRules").getRegistryEnterpriseSyncRules = null as any;
export const getRegistryEnterpriseSyncRulesOutput: typeof import("./getRegistryEnterpriseSyncRules").getRegistryEnterpriseSyncRulesOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryEnterpriseSyncRules","getRegistryEnterpriseSyncRulesOutput"], () => require("./getRegistryEnterpriseSyncRules"));

export { GetServerlessKubernetesClustersArgs, GetServerlessKubernetesClustersResult, GetServerlessKubernetesClustersOutputArgs } from "./getServerlessKubernetesClusters";
export const getServerlessKubernetesClusters: typeof import("./getServerlessKubernetesClusters").getServerlessKubernetesClusters = null as any;
export const getServerlessKubernetesClustersOutput: typeof import("./getServerlessKubernetesClusters").getServerlessKubernetesClustersOutput = null as any;
utilities.lazyLoad(exports, ["getServerlessKubernetesClusters","getServerlessKubernetesClustersOutput"], () => require("./getServerlessKubernetesClusters"));

export { KubernetesArgs, KubernetesState } from "./kubernetes";
export type Kubernetes = import("./kubernetes").Kubernetes;
export const Kubernetes: typeof import("./kubernetes").Kubernetes = null as any;
utilities.lazyLoad(exports, ["Kubernetes"], () => require("./kubernetes"));

export { KubernetesAddonArgs, KubernetesAddonState } from "./kubernetesAddon";
export type KubernetesAddon = import("./kubernetesAddon").KubernetesAddon;
export const KubernetesAddon: typeof import("./kubernetesAddon").KubernetesAddon = null as any;
utilities.lazyLoad(exports, ["KubernetesAddon"], () => require("./kubernetesAddon"));

export { KubernetesAutoscalerArgs, KubernetesAutoscalerState } from "./kubernetesAutoscaler";
export type KubernetesAutoscaler = import("./kubernetesAutoscaler").KubernetesAutoscaler;
export const KubernetesAutoscaler: typeof import("./kubernetesAutoscaler").KubernetesAutoscaler = null as any;
utilities.lazyLoad(exports, ["KubernetesAutoscaler"], () => require("./kubernetesAutoscaler"));

export { KubernetesPermissionArgs, KubernetesPermissionState } from "./kubernetesPermission";
export type KubernetesPermission = import("./kubernetesPermission").KubernetesPermission;
export const KubernetesPermission: typeof import("./kubernetesPermission").KubernetesPermission = null as any;
utilities.lazyLoad(exports, ["KubernetesPermission"], () => require("./kubernetesPermission"));

export { ManagedKubernetesArgs, ManagedKubernetesState } from "./managedKubernetes";
export type ManagedKubernetes = import("./managedKubernetes").ManagedKubernetes;
export const ManagedKubernetes: typeof import("./managedKubernetes").ManagedKubernetes = null as any;
utilities.lazyLoad(exports, ["ManagedKubernetes"], () => require("./managedKubernetes"));

export { NodePoolArgs, NodePoolState } from "./nodePool";
export type NodePool = import("./nodePool").NodePool;
export const NodePool: typeof import("./nodePool").NodePool = null as any;
utilities.lazyLoad(exports, ["NodePool"], () => require("./nodePool"));

export { RegistryEnterpriseNamespaceArgs, RegistryEnterpriseNamespaceState } from "./registryEnterpriseNamespace";
export type RegistryEnterpriseNamespace = import("./registryEnterpriseNamespace").RegistryEnterpriseNamespace;
export const RegistryEnterpriseNamespace: typeof import("./registryEnterpriseNamespace").RegistryEnterpriseNamespace = null as any;
utilities.lazyLoad(exports, ["RegistryEnterpriseNamespace"], () => require("./registryEnterpriseNamespace"));

export { RegistryEnterpriseRepoArgs, RegistryEnterpriseRepoState } from "./registryEnterpriseRepo";
export type RegistryEnterpriseRepo = import("./registryEnterpriseRepo").RegistryEnterpriseRepo;
export const RegistryEnterpriseRepo: typeof import("./registryEnterpriseRepo").RegistryEnterpriseRepo = null as any;
utilities.lazyLoad(exports, ["RegistryEnterpriseRepo"], () => require("./registryEnterpriseRepo"));

export { RegistryEnterpriseSyncRuleArgs, RegistryEnterpriseSyncRuleState } from "./registryEnterpriseSyncRule";
export type RegistryEnterpriseSyncRule = import("./registryEnterpriseSyncRule").RegistryEnterpriseSyncRule;
export const RegistryEnterpriseSyncRule: typeof import("./registryEnterpriseSyncRule").RegistryEnterpriseSyncRule = null as any;
utilities.lazyLoad(exports, ["RegistryEnterpriseSyncRule"], () => require("./registryEnterpriseSyncRule"));

export { ServerlessKubernetesArgs, ServerlessKubernetesState } from "./serverlessKubernetes";
export type ServerlessKubernetes = import("./serverlessKubernetes").ServerlessKubernetes;
export const ServerlessKubernetes: typeof import("./serverlessKubernetes").ServerlessKubernetes = null as any;
utilities.lazyLoad(exports, ["ServerlessKubernetes"], () => require("./serverlessKubernetes"));

export { SwarmArgs, SwarmState } from "./swarm";
export type Swarm = import("./swarm").Swarm;
export const Swarm: typeof import("./swarm").Swarm = null as any;
utilities.lazyLoad(exports, ["Swarm"], () => require("./swarm"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:cs/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "alicloud:cs/autoscalingConfig:AutoscalingConfig":
                return new AutoscalingConfig(name, <any>undefined, { urn })
            case "alicloud:cs/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "alicloud:cs/edgeKubernetes:EdgeKubernetes":
                return new EdgeKubernetes(name, <any>undefined, { urn })
            case "alicloud:cs/kubernetes:Kubernetes":
                return new Kubernetes(name, <any>undefined, { urn })
            case "alicloud:cs/kubernetesAddon:KubernetesAddon":
                return new KubernetesAddon(name, <any>undefined, { urn })
            case "alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler":
                return new KubernetesAutoscaler(name, <any>undefined, { urn })
            case "alicloud:cs/kubernetesPermission:KubernetesPermission":
                return new KubernetesPermission(name, <any>undefined, { urn })
            case "alicloud:cs/managedKubernetes:ManagedKubernetes":
                return new ManagedKubernetes(name, <any>undefined, { urn })
            case "alicloud:cs/nodePool:NodePool":
                return new NodePool(name, <any>undefined, { urn })
            case "alicloud:cs/registryEnterpriseNamespace:RegistryEnterpriseNamespace":
                return new RegistryEnterpriseNamespace(name, <any>undefined, { urn })
            case "alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo":
                return new RegistryEnterpriseRepo(name, <any>undefined, { urn })
            case "alicloud:cs/registryEnterpriseSyncRule:RegistryEnterpriseSyncRule":
                return new RegistryEnterpriseSyncRule(name, <any>undefined, { urn })
            case "alicloud:cs/serverlessKubernetes:ServerlessKubernetes":
                return new ServerlessKubernetes(name, <any>undefined, { urn })
            case "alicloud:cs/swarm:Swarm":
                return new Swarm(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "cs/application", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/autoscalingConfig", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/cluster", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/edgeKubernetes", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/kubernetes", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/kubernetesAddon", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/kubernetesAutoscaler", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/kubernetesPermission", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/managedKubernetes", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/nodePool", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/registryEnterpriseNamespace", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/registryEnterpriseRepo", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/registryEnterpriseSyncRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/serverlessKubernetes", _module)
pulumi.runtime.registerResourceModule("alicloud", "cs/swarm", _module)

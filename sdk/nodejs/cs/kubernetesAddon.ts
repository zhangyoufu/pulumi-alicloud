// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource will help you to manage addon in Kubernetes Cluster.
 *
 * > **NOTE:** Available in 1.150.0+.
 * **NOTE:** From version 1.166.0, support specifying addon customizable configuration.
 *
 * ## Example Usage
 *
 * **Create a managed cluster**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-test";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 2,
 *     memorySize: 4,
 *     kubernetesNodeRole: "Worker",
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.1.0.0/21",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "10.1.1.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultKeyPair = new alicloud.ecs.KeyPair("defaultKeyPair", {keyPairName: name});
 * let defaultManagedKubernetes: alicloud.cs.ManagedKubernetes | undefined;
 * if (1 == true) {
 *     defaultManagedKubernetes = new alicloud.cs.ManagedKubernetes("defaultManagedKubernetes", {
 *         clusterSpec: "ack.pro.small",
 *         isEnterpriseSecurityGroup: true,
 *         workerNumber: 2,
 *         password: "Hello1234",
 *         podCidr: "172.20.0.0/16",
 *         serviceCidr: "172.21.0.0/20",
 *         workerVswitchIds: [defaultSwitch.id],
 *         workerInstanceTypes: [defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id)],
 *     });
 * }
 * ```
 * **Installing of addon**
 * When a cluster is created, some system addons and those specified at the time of cluster creation will be installed, so when an addon resource is applied:
 * * If the addon already exists in the cluster and its version is the same as the specified version, it will be skipped and will not be reinstalled.
 * * If the addon already exists in the cluster and its version is different from the specified version, the addon will be upgraded.
 * * If the addon does not exist in the cluster, it will be installed.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ack_node_problem_detector = new alicloud.cs.KubernetesAddon("ack-node-problem-detector", {
 *     clusterId: alicloud_cs_managed_kubernetes["default"][0].id,
 *     version: "1.2.7",
 * });
 * const nginxIngressController = new alicloud.cs.KubernetesAddon("nginxIngressController", {
 *     clusterId: _var.cluster_id,
 *     version: "v1.1.2-aliyun.2",
 *     config: JSON.stringify({
 *         CpuLimit: "",
 *         CpuRequest: "100m",
 *         EnableWebhook: true,
 *         HostNetwork: false,
 *         IngressSlbNetworkType: "internet",
 *         IngressSlbSpec: "slb.s2.small",
 *         MemoryLimit: "",
 *         MemoryRequest: "200Mi",
 *         NodeSelector: [],
 *     }),
 * });
 * ```
 *
 * **Upgrading of addon**
 * First, check the `nextVersion` field of the addon that can be upgraded to through the `.tfstate file`, then overwrite the `version` field with the value of `nextVersion` and apply.
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ack_node_problem_detector = new alicloud.cs.KubernetesAddon("ack-node-problem-detector", {
 *     clusterId: alicloud_cs_managed_kubernetes["default"][0].id,
 *     version: "1.2.8",
 * });
 * // upgrade from 1.2.7 to 1.2.8
 * ```
 *
 * ## Import
 *
 * Cluster addon can be imported by cluster id and addon name. Then write the addon.tf file according to the result of `pulumi preview`.
 *
 * ```sh
 *  $ pulumi import alicloud:cs/kubernetesAddon:KubernetesAddon my_addon <cluster_id>:<addon_name>
 * ```
 */
export class KubernetesAddon extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesAddon resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesAddonState, opts?: pulumi.CustomResourceOptions): KubernetesAddon {
        return new KubernetesAddon(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/kubernetesAddon:KubernetesAddon';

    /**
     * Returns true if the given object is an instance of KubernetesAddon.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesAddon {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesAddon.__pulumiType;
    }

    /**
     * Is the addon ready for upgrade.
     */
    public /*out*/ readonly canUpgrade!: pulumi.Output<boolean>;
    /**
     * The id of kubernetes cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The name of addon.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The version which addon can be upgraded to.
     */
    public /*out*/ readonly nextVersion!: pulumi.Output<string>;
    /**
     * Is it a mandatory addon to be installed.
     */
    public /*out*/ readonly required!: pulumi.Output<boolean>;
    /**
     * The current version of addon.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a KubernetesAddon resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesAddonArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesAddonArgs | KubernetesAddonState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KubernetesAddonState | undefined;
            resourceInputs["canUpgrade"] = state ? state.canUpgrade : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextVersion"] = state ? state.nextVersion : undefined;
            resourceInputs["required"] = state ? state.required : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as KubernetesAddonArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["canUpgrade"] = undefined /*out*/;
            resourceInputs["nextVersion"] = undefined /*out*/;
            resourceInputs["required"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubernetesAddon.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesAddon resources.
 */
export interface KubernetesAddonState {
    /**
     * Is the addon ready for upgrade.
     */
    canUpgrade?: pulumi.Input<boolean>;
    /**
     * The id of kubernetes cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
     */
    config?: pulumi.Input<string>;
    /**
     * The name of addon.
     */
    name?: pulumi.Input<string>;
    /**
     * The version which addon can be upgraded to.
     */
    nextVersion?: pulumi.Input<string>;
    /**
     * Is it a mandatory addon to be installed.
     */
    required?: pulumi.Input<boolean>;
    /**
     * The current version of addon.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KubernetesAddon resource.
 */
export interface KubernetesAddonArgs {
    /**
     * The id of kubernetes cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
     */
    config?: pulumi.Input<string>;
    /**
     * The name of addon.
     */
    name?: pulumi.Input<string>;
    /**
     * The current version of addon.
     */
    version: pulumi.Input<string>;
}

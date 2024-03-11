// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Ack One Cluster resource. Fleet Manager Cluster.
 *
 * For information about Ack One Cluster and how to use it, see [What is Cluster](https://www.alibabacloud.com/help/en/ack/distributed-cloud-container-platform-for-kubernetes/developer-reference/api-adcp-2022-01-01-createhubcluster).
 *
 * > **NOTE:** Available since v1.212.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultVpc = new alicloud.vpc.Network("defaultVpc", {
 *     cidrBlock: "172.16.0.0/12",
 *     vpcName: name,
 * });
 * const defaultyVSwitch = new alicloud.vpc.Switch("defaultyVSwitch", {
 *     vpcId: defaultVpc.id,
 *     cidrBlock: "172.16.2.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultCluster = new alicloud.ackone.Cluster("defaultCluster", {network: {
 *     vpcId: defaultVpc.id,
 *     vswitches: [defaultyVSwitch.id],
 * }});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Ack One Cluster can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ackone/cluster:Cluster example <id>
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ackone/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Cluster name.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Cluster creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Cluster network information. See `network` below.
     */
    public readonly network!: pulumi.Output<outputs.ackone.ClusterNetwork>;
    /**
     * Cluster attributes. Valid values: 'Default', 'XFlow'.
     */
    public readonly profile!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Cluster name.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Cluster creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Cluster network information. See `network` below.
     */
    network?: pulumi.Input<inputs.ackone.ClusterNetwork>;
    /**
     * Cluster attributes. Valid values: 'Default', 'XFlow'.
     */
    profile?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Cluster name.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Cluster network information. See `network` below.
     */
    network: pulumi.Input<inputs.ackone.ClusterNetwork>;
    /**
     * Cluster attributes. Valid values: 'Default', 'XFlow'.
     */
    profile?: pulumi.Input<string>;
}

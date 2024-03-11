// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Microservice Engine (MSE) Znode resource.
 *
 * For information about Microservice Engine (MSE) Znode and how to use it, see [What is Znode](https://help.aliyun.com/document_detail/393622.html).
 *
 * > **NOTE:** Available in v1.162.0+.
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
 * const exampleZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("exampleSwitch", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: exampleNetwork.id,
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 * });
 * const exampleCluster = new alicloud.mse.Cluster("exampleCluster", {
 *     clusterSpecification: "MSE_SC_1_2_60_c",
 *     clusterType: "ZooKeeper",
 *     clusterVersion: "ZooKeeper_3_8_0",
 *     instanceCount: 1,
 *     netType: "privatenet",
 *     pubNetworkFlow: "1",
 *     aclEntryLists: ["127.0.0.1/32"],
 *     clusterAliasName: "terraform-example",
 *     mseVersion: "mse_dev",
 *     vswitchId: exampleSwitch.id,
 *     vpcId: exampleNetwork.id,
 * });
 * const exampleZnode = new alicloud.mse.Znode("exampleZnode", {
 *     clusterId: exampleCluster.clusterId,
 *     data: "terraform-example",
 *     path: "/example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Microservice Engine (MSE) Znode can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:mse/znode:Znode example <cluster_id>:<path>
 * ```
 */
export class Znode extends pulumi.CustomResource {
    /**
     * Get an existing Znode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZnodeState, opts?: pulumi.CustomResourceOptions): Znode {
        return new Znode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mse/znode:Znode';

    /**
     * Returns true if the given object is an instance of Znode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Znode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Znode.__pulumiType;
    }

    /**
     * The language type of the returned information. Valid values: `zh` or `en`.
     */
    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The Node data.
     */
    public readonly data!: pulumi.Output<string | undefined>;
    /**
     * The Node path. The value must start with a forward slash (/).
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a Znode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZnodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZnodeArgs | ZnodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZnodeState | undefined;
            resourceInputs["acceptLanguage"] = state ? state.acceptLanguage : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as ZnodeArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Znode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Znode resources.
 */
export interface ZnodeState {
    /**
     * The language type of the returned information. Valid values: `zh` or `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * The ID of the Cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The Node data.
     */
    data?: pulumi.Input<string>;
    /**
     * The Node path. The value must start with a forward slash (/).
     */
    path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Znode resource.
 */
export interface ZnodeArgs {
    /**
     * The language type of the returned information. Valid values: `zh` or `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * The ID of the Cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The Node data.
     */
    data?: pulumi.Input<string>;
    /**
     * The Node path. The value must start with a forward slash (/).
     */
    path: pulumi.Input<string>;
}

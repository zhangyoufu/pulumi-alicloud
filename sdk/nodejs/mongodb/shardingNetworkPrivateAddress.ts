// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a MongoDB Sharding Network Private Address resource.
 *
 * For information about MongoDB Sharding Network Private Address and how to use it, see [What is Sharding Network Private Address](https://www.alibabacloud.com/help/en/doc-detail/141403.html).
 *
 * > **NOTE:** Available since v1.157.0.
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
 * const default = alicloud.mongodb.getZones({});
 * const index = _default.then(_default => _default.zones).length.then(length => length - 1);
 * const zoneId = _default.then(_default => _default.zones[index].id);
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vswitchName: name,
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: zoneId,
 * });
 * const defaultShardingInstance = new alicloud.mongodb.ShardingInstance("default", {
 *     zoneId: zoneId,
 *     vswitchId: defaultSwitch.id,
 *     engineVersion: "4.2",
 *     name: name,
 *     shardLists: [
 *         {
 *             nodeClass: "dds.shard.mid",
 *             nodeStorage: 10,
 *         },
 *         {
 *             nodeClass: "dds.shard.standard",
 *             nodeStorage: 20,
 *             readonlyReplicas: 1,
 *         },
 *     ],
 *     mongoLists: [
 *         {
 *             nodeClass: "dds.mongos.mid",
 *         },
 *         {
 *             nodeClass: "dds.mongos.mid",
 *         },
 *     ],
 * });
 * const defaultShardingNetworkPrivateAddress = new alicloud.mongodb.ShardingNetworkPrivateAddress("default", {
 *     dbInstanceId: defaultShardingInstance.id,
 *     nodeId: defaultShardingInstance.shardLists.apply(shardLists => shardLists[0].nodeId),
 *     zoneId: defaultShardingInstance.zoneId,
 *     accountName: "example",
 *     accountPassword: "Example_123",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * MongoDB Sharding Network Private Address can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:mongodb/shardingNetworkPrivateAddress:ShardingNetworkPrivateAddress example <db_instance_id>:<node_id>
 * ```
 */
export class ShardingNetworkPrivateAddress extends pulumi.CustomResource {
    /**
     * Get an existing ShardingNetworkPrivateAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShardingNetworkPrivateAddressState, opts?: pulumi.CustomResourceOptions): ShardingNetworkPrivateAddress {
        return new ShardingNetworkPrivateAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mongodb/shardingNetworkPrivateAddress:ShardingNetworkPrivateAddress';

    /**
     * Returns true if the given object is an instance of ShardingNetworkPrivateAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShardingNetworkPrivateAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShardingNetworkPrivateAddress.__pulumiType;
    }

    /**
     * The name of the account. 
     * - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
     * - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
     * - The permissions of this account are fixed to read-only.
     */
    public readonly accountName!: pulumi.Output<string | undefined>;
    /**
     * Account password. 
     * - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
     * - The password must be 8 to 32 characters in length.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * The db instance id.
     */
    public readonly dbInstanceId!: pulumi.Output<string>;
    /**
     * The endpoint of the instance.
     */
    public /*out*/ readonly networkAddresses!: pulumi.Output<outputs.mongodb.ShardingNetworkPrivateAddressNetworkAddress[]>;
    /**
     * The ID of the Shard node or the ConfigServer node.
     */
    public readonly nodeId!: pulumi.Output<string>;
    /**
     * The zone ID of the instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ShardingNetworkPrivateAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShardingNetworkPrivateAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShardingNetworkPrivateAddressArgs | ShardingNetworkPrivateAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShardingNetworkPrivateAddressState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
            resourceInputs["networkAddresses"] = state ? state.networkAddresses : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ShardingNetworkPrivateAddressArgs | undefined;
            if ((!args || args.dbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceId'");
            }
            if ((!args || args.nodeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
            resourceInputs["nodeId"] = args ? args.nodeId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["networkAddresses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ShardingNetworkPrivateAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShardingNetworkPrivateAddress resources.
 */
export interface ShardingNetworkPrivateAddressState {
    /**
     * The name of the account. 
     * - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
     * - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
     * - The permissions of this account are fixed to read-only.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Account password. 
     * - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
     * - The password must be 8 to 32 characters in length.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The db instance id.
     */
    dbInstanceId?: pulumi.Input<string>;
    /**
     * The endpoint of the instance.
     */
    networkAddresses?: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingNetworkPrivateAddressNetworkAddress>[]>;
    /**
     * The ID of the Shard node or the ConfigServer node.
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShardingNetworkPrivateAddress resource.
 */
export interface ShardingNetworkPrivateAddressArgs {
    /**
     * The name of the account. 
     * - The name must be 4 to 16 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a lowercase letter.
     * - You need to set the account name and password only when you apply for an endpoint for a shard or Configserver node for the first time. In this case, the account name and password are used for all shard and Configserver nodes.
     * - The permissions of this account are fixed to read-only.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Account password. 
     * - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%^&*()_+-=`.
     * - The password must be 8 to 32 characters in length.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The db instance id.
     */
    dbInstanceId: pulumi.Input<string>;
    /**
     * The ID of the Shard node or the ConfigServer node.
     */
    nodeId: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId: pulumi.Input<string>;
}

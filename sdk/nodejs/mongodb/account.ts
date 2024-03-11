// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a MongoDB Account resource.
 *
 * For information about MongoDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/doc-detail/62154.html).
 *
 * > **NOTE:** Available since v1.148.0.
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
 * const defaultZones = alicloud.mongodb.getZones({});
 * const index = defaultZones.then(defaultZones => defaultZones.zones).length.then(length => length - 1);
 * const zoneId = defaultZones.then(defaultZones => defaultZones.zones[index].id);
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: zoneId,
 * });
 * const defaultInstance = new alicloud.mongodb.Instance("defaultInstance", {
 *     engineVersion: "4.2",
 *     dbInstanceClass: "dds.mongo.mid",
 *     dbInstanceStorage: 10,
 *     vswitchId: defaultSwitch.id,
 *     securityIpLists: [
 *         "10.168.1.12",
 *         "100.69.7.112",
 *     ],
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const defaultAccount = new alicloud.mongodb.Account("defaultAccount", {
 *     accountName: "root",
 *     accountPassword: "Example_123",
 *     instanceId: defaultInstance.id,
 *     accountDescription: name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * MongoDB Account can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:mongodb/account:Account example <instance_id>:<account_name>
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mongodb/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * The description of the account.
     * * The description must start with a letter, and cannot start with `http://` or `https://`.
     * * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
     */
    public readonly accountDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the account. Valid values: `root`.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The Password of the Account.
     * * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
     * * The password must be `8` to `32` characters in length.
     */
    public readonly accountPassword!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The status of the account. Valid values: `Unavailable`, `Available`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            resourceInputs["accountDescription"] = state ? state.accountDescription : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.accountPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountPassword'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accountDescription"] = args ? args.accountDescription : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * The description of the account.
     * * The description must start with a letter, and cannot start with `http://` or `https://`.
     * * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
     */
    accountDescription?: pulumi.Input<string>;
    /**
     * The name of the account. Valid values: `root`.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The Password of the Account.
     * * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
     * * The password must be `8` to `32` characters in length.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The status of the account. Valid values: `Unavailable`, `Available`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The description of the account.
     * * The description must start with a letter, and cannot start with `http://` or `https://`.
     * * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
     */
    accountDescription?: pulumi.Input<string>;
    /**
     * The name of the account. Valid values: `root`.
     */
    accountName: pulumi.Input<string>;
    /**
     * The Password of the Account.
     * * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
     * * The password must be `8` to `32` characters in length.
     */
    accountPassword: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
}

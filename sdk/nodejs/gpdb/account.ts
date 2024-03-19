// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a GPDB Account resource.
 *
 * For information about GPDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/doc-detail/86924.htm).
 *
 * > **NOTE:** Available since v1.142.0.
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
 * const defaultZones = alicloud.gpdb.getZones({});
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?.[0],
 *     zoneId: defaultZones.ids?.[0],
 * }));
 * const defaultInstance = new alicloud.gpdb.Instance("defaultInstance", {
 *     dbInstanceCategory: "HighAvailability",
 *     dbInstanceClass: "gpdb.group.segsdx1",
 *     dbInstanceMode: "StorageElastic",
 *     description: name,
 *     engine: "gpdb",
 *     engineVersion: "6.0",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.ids?.[0]),
 *     instanceNetworkType: "VPC",
 *     instanceSpec: "2C16G",
 *     masterNodeNum: 1,
 *     paymentType: "PayAsYouGo",
 *     privateIpAddress: "1.1.1.1",
 *     segStorageType: "cloud_essd",
 *     segNodeNum: 4,
 *     storageSize: 50,
 *     vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?.[0]),
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 *     ipWhitelists: [{
 *         securityIpList: "127.0.0.1",
 *     }],
 * });
 * const defaultAccount = new alicloud.gpdb.Account("defaultAccount", {
 *     accountName: "tf_example",
 *     dbInstanceId: defaultInstance.id,
 *     accountPassword: "Example1234",
 *     accountDescription: "tf_example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GPDB Account can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:gpdb/account:Account example <db_instance_id>:<account_name>
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
    public static readonly __pulumiType = 'alicloud:gpdb/account:Account';

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
     * * Starts with a letter.
     * * Does not start with `http://` or `https://`.
     * * Contains letters, underscores (_), hyphens (-), or digits.
     * * Be 2 to 256 characters in length.
     */
    public readonly accountDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the account. The account name must be unique and meet the following requirements:
     * * Starts with a letter.
     * * Contains only lowercase letters, digits, or underscores (_).
     * * Be up to 16 characters in length.
     * * Contains no reserved keywords.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) & * ( ) _ + - =`.
     */
    public readonly accountPassword!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly dbInstanceId!: pulumi.Output<string>;
    /**
     * The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
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
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.accountPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountPassword'");
            }
            if ((!args || args.dbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceId'");
            }
            resourceInputs["accountDescription"] = args ? args.accountDescription : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args ? args.accountPassword : undefined;
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * The description of the account.
     * * Starts with a letter.
     * * Does not start with `http://` or `https://`.
     * * Contains letters, underscores (_), hyphens (-), or digits.
     * * Be 2 to 256 characters in length.
     */
    accountDescription?: pulumi.Input<string>;
    /**
     * The name of the account. The account name must be unique and meet the following requirements:
     * * Starts with a letter.
     * * Contains only lowercase letters, digits, or underscores (_).
     * * Be up to 16 characters in length.
     * * Contains no reserved keywords.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) & * ( ) _ + - =`.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    dbInstanceId?: pulumi.Input<string>;
    /**
     * The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The description of the account.
     * * Starts with a letter.
     * * Does not start with `http://` or `https://`.
     * * Contains letters, underscores (_), hyphens (-), or digits.
     * * Be 2 to 256 characters in length.
     */
    accountDescription?: pulumi.Input<string>;
    /**
     * The name of the account. The account name must be unique and meet the following requirements:
     * * Starts with a letter.
     * * Contains only lowercase letters, digits, or underscores (_).
     * * Be up to 16 characters in length.
     * * Contains no reserved keywords.
     */
    accountName: pulumi.Input<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) & * ( ) _ + - =`.
     */
    accountPassword: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    dbInstanceId: pulumi.Input<string>;
}

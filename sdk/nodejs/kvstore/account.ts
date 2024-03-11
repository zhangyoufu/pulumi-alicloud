// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a KVStore Account resource.
 *
 * For information about KVStore Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/doc-detail/95973.htm).
 *
 * > **NOTE:** Available since v1.66.0.
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
 * const name = config.get("name") || "tf-example";
 * const defaultZones = alicloud.kvstore.getZones({});
 * const defaultResourceGroups = alicloud.resourcemanager.getResourceGroups({
 *     status: "OK",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultInstance = new alicloud.kvstore.Instance("defaultInstance", {
 *     dbInstanceName: name,
 *     vswitchId: defaultSwitch.id,
 *     resourceGroupId: defaultResourceGroups.then(defaultResourceGroups => defaultResourceGroups.ids?.[0]),
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     instanceClass: "redis.master.large.default",
 *     instanceType: "Redis",
 *     engineVersion: "5.0",
 *     securityIps: ["10.23.12.24"],
 *     config: {
 *         appendonly: "yes",
 *         "lazyfree-lazy-eviction": "yes",
 *     },
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * const defaultAccount = new alicloud.kvstore.Account("defaultAccount", {
 *     accountName: "tfexamplename",
 *     accountPassword: "YourPassword_123",
 *     instanceId: defaultInstance.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * KVStore account can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:kvstore/account:Account example <instance_id>:<account_name>
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
    public static readonly __pulumiType = 'alicloud:kvstore/account:Account';

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
     * The name of the account. The name must meet the following requirements:
     * * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
     * * The name can be up to 100 characters in length.
     * * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ %!^(MISSING) & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * The privilege of account access database. Default value: `RoleReadWrite`
     */
    public readonly accountPrivilege!: pulumi.Output<string | undefined>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * Default to Normal.
     */
    public readonly accountType!: pulumi.Output<string | undefined>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The status of KVStore Account.
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
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["accountPrivilege"] = state ? state.accountPrivilege : undefined;
            resourceInputs["accountType"] = state ? state.accountType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["accountPrivilege"] = args ? args.accountPrivilege : undefined;
            resourceInputs["accountType"] = args ? args.accountType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
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
     * The name of the account. The name must meet the following requirements:
     * * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
     * * The name can be up to 100 characters in length.
     * * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ %!^(MISSING) & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The privilege of account access database. Default value: `RoleReadWrite`
     */
    accountPrivilege?: pulumi.Input<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * Default to Normal.
     */
    accountType?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
     */
    instanceId?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The status of KVStore Account.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The name of the account. The name must meet the following requirements:
     * * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
     * * The name can be up to 100 characters in length.
     * * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
     */
    accountName: pulumi.Input<string>;
    /**
     * The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ %!^(MISSING) & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The privilege of account access database. Default value: `RoleReadWrite`
     */
    accountPrivilege?: pulumi.Input<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * Default to Normal.
     */
    accountType?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
     */
    instanceId: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
}

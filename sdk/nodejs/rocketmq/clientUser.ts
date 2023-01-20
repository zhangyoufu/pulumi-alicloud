// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Sag ClientUser resource. This topic describes how to manage accounts as an administrator. After you configure the network, you can create multiple accounts and distribute them to end users so that clients can access Alibaba Cloud.
 *
 * For information about Sag ClientUser and how to use it, see [What is Sag ClientUser](https://www.alibabacloud.com/help/doc-detail/108326.htm).
 *
 * > **NOTE:** Available in 1.65.0+
 *
 * > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.rocketmq.ClientUser("default", {
 *     bandwidth: 20,
 *     clientIp: "192.1.10.0",
 *     password: "xxxxxxx",
 *     sagId: "sag-xxxxx",
 *     userMail: "tftest-xxxxx@test.com",
 *     userName: "th-username-xxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * The Sag ClientUser can be imported using the name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:rocketmq/clientUser:ClientUser example sag-abc123456:tf-username-abc123456
 * ```
 */
export class ClientUser extends pulumi.CustomResource {
    /**
     * Get an existing ClientUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientUserState, opts?: pulumi.CustomResourceOptions): ClientUser {
        return new ClientUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rocketmq/clientUser:ClientUser';

    /**
     * Returns true if the given object is an instance of ClientUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientUser.__pulumiType;
    }

    /**
     * The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
     */
    public readonly clientIp!: pulumi.Output<string | undefined>;
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The ID of the SAG instance created for the SAG APP.
     */
    public readonly sagId!: pulumi.Output<string>;
    /**
     * The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
     */
    public readonly userMail!: pulumi.Output<string>;
    /**
     * The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a ClientUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientUserArgs | ClientUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientUserState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["clientIp"] = state ? state.clientIp : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["sagId"] = state ? state.sagId : undefined;
            resourceInputs["userMail"] = state ? state.userMail : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as ClientUserArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.sagId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sagId'");
            }
            if ((!args || args.userMail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userMail'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["clientIp"] = args ? args.clientIp : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["sagId"] = args ? args.sagId : undefined;
            resourceInputs["userMail"] = args ? args.userMail : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientUser resources.
 */
export interface ClientUserState {
    /**
     * The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
     */
    clientIp?: pulumi.Input<string>;
    kmsEncryptedPassword?: pulumi.Input<string>;
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the SAG instance created for the SAG APP.
     */
    sagId?: pulumi.Input<string>;
    /**
     * The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
     */
    userMail?: pulumi.Input<string>;
    /**
     * The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientUser resource.
 */
export interface ClientUserArgs {
    /**
     * The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address.
     */
    clientIp?: pulumi.Input<string>;
    kmsEncryptedPassword?: pulumi.Input<string>;
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the SAG instance created for the SAG APP.
     */
    sagId: pulumi.Input<string>;
    /**
     * The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.
     */
    userMail: pulumi.Input<string>;
    /**
     * The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.
     */
    userName?: pulumi.Input<string>;
}

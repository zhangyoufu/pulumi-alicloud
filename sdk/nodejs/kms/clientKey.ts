// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a KMS Client Key resource. Client key (of Application Access Point).
 *
 * For information about KMS Client Key and how to use it, see [What is Client Key](https://www.alibabacloud.com/help/zh/key-management-service/latest/api-createclientkey).
 *
 * > **NOTE:** Available since v1.210.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const aAP0 = new alicloud.kms.ApplicationAccessPoint("aAP0", {
 *     policies: ["aa"],
 *     description: "aa",
 *     applicationAccessPointName: name,
 * });
 * const _default = new alicloud.kms.ClientKey("default", {
 *     aapName: aAP0.applicationAccessPointName,
 *     password: "YouPassword123!",
 *     notBefore: "2023-09-01T14:11:22Z",
 *     notAfter: "2028-09-01T14:11:22Z",
 *     privateKeyDataFile: "./private_key_data_file.txt",
 * });
 * ```
 *
 * ## Import
 *
 * KMS Client Key can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:kms/clientKey:ClientKey example <id>
 * ```
 *
 *  Resource attributes such as `password`, `private_key_data_file` are not available for imported resources as this information cannot be read from the KMS API.
 */
export class ClientKey extends pulumi.CustomResource {
    /**
     * Get an existing ClientKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientKeyState, opts?: pulumi.CustomResourceOptions): ClientKey {
        return new ClientKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kms/clientKey:ClientKey';

    /**
     * Returns true if the given object is an instance of ClientKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientKey.__pulumiType;
    }

    /**
     * ClientKey's parent Application Access Point name.
     */
    public readonly aapName!: pulumi.Output<string>;
    /**
     * Create timestamp, e.g. "2022-08-10T08:03:30Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The ClientKey expiration time. Example: "2027-08-10 T08:03:30Z".
     */
    public readonly notAfter!: pulumi.Output<string | undefined>;
    /**
     * The valid start time of the ClientKey. Example: "2022-08-10 T08:03:30Z".
     */
    public readonly notBefore!: pulumi.Output<string | undefined>;
    /**
     * To enhance security, set a password for the downloaded Client Key,When an application accesses KMS, you must use the ClientKey content and this password to initialize the SDK client.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
     */
    public readonly privateKeyDataFile!: pulumi.Output<string | undefined>;

    /**
     * Create a ClientKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientKeyArgs | ClientKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientKeyState | undefined;
            resourceInputs["aapName"] = state ? state.aapName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["notAfter"] = state ? state.notAfter : undefined;
            resourceInputs["notBefore"] = state ? state.notBefore : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateKeyDataFile"] = state ? state.privateKeyDataFile : undefined;
        } else {
            const args = argsOrState as ClientKeyArgs | undefined;
            if ((!args || args.aapName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aapName'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["aapName"] = args ? args.aapName : undefined;
            resourceInputs["notAfter"] = args ? args.notAfter : undefined;
            resourceInputs["notBefore"] = args ? args.notBefore : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["privateKeyDataFile"] = args ? args.privateKeyDataFile : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ClientKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientKey resources.
 */
export interface ClientKeyState {
    /**
     * ClientKey's parent Application Access Point name.
     */
    aapName?: pulumi.Input<string>;
    /**
     * Create timestamp, e.g. "2022-08-10T08:03:30Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ClientKey expiration time. Example: "2027-08-10 T08:03:30Z".
     */
    notAfter?: pulumi.Input<string>;
    /**
     * The valid start time of the ClientKey. Example: "2022-08-10 T08:03:30Z".
     */
    notBefore?: pulumi.Input<string>;
    /**
     * To enhance security, set a password for the downloaded Client Key,When an application accesses KMS, you must use the ClientKey content and this password to initialize the SDK client.
     */
    password?: pulumi.Input<string>;
    /**
     * The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
     */
    privateKeyDataFile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientKey resource.
 */
export interface ClientKeyArgs {
    /**
     * ClientKey's parent Application Access Point name.
     */
    aapName: pulumi.Input<string>;
    /**
     * The ClientKey expiration time. Example: "2027-08-10 T08:03:30Z".
     */
    notAfter?: pulumi.Input<string>;
    /**
     * The valid start time of the ClientKey. Example: "2022-08-10 T08:03:30Z".
     */
    notBefore?: pulumi.Input<string>;
    /**
     * To enhance security, set a password for the downloaded Client Key,When an application accesses KMS, you must use the ClientKey content and this password to initialize the SDK client.
     */
    password: pulumi.Input<string>;
    /**
     * The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
     */
    privateKeyDataFile?: pulumi.Input<string>;
}

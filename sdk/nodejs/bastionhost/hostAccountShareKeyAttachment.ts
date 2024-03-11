// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Bastion Host Account Share Key Attachment resource.
 *
 * For information about Bastion Host Host Account Share Key Attachment and how to use it, see [What is Host Account Share Key Attachment](https://www.alibabacloud.com/help/en/bastion-host/latest/attachhostaccountstohostsharekey).
 *
 * > **NOTE:** Available since v1.165.0.
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
 * const name = config.get("name") || "tf_example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetworks.ids?.[0],
 *     zoneId: defaultZones.zones?.[0]?.id,
 * }));
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?.[0])});
 * const defaultInstance = new alicloud.bastionhost.Instance("defaultInstance", {
 *     description: name,
 *     licenseCode: "bhah_ent_50_asset",
 *     planCode: "cloudbastion",
 *     storage: "5",
 *     bandwidth: "5",
 *     period: 1,
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 *     securityGroupIds: [defaultSecurityGroup.id],
 * });
 * const defaultHost = new alicloud.bastionhost.Host("defaultHost", {
 *     instanceId: defaultInstance.id,
 *     hostName: name,
 *     activeAddressType: "Private",
 *     hostPrivateAddress: "172.16.0.10",
 *     osType: "Linux",
 *     source: "Local",
 * });
 * const defaultHostAccount = new alicloud.bastionhost.HostAccount("defaultHostAccount", {
 *     hostAccountName: name,
 *     hostId: defaultHost.hostId,
 *     instanceId: defaultHost.instanceId,
 *     protocolName: "SSH",
 *     password: "YourPassword12345",
 * });
 * const privateKey = config.get("privateKey") || "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBc25oc29SSVVwVXltSG1FVHJXUGxDbkhMa3c3N0JYTm44ZHcvbDg3eG10SUhjd2syCkRybjFDZk5jZkpJV0tSdkFaYkdKMlZTS1RiRDhPTmcyT3JvUHFGUHBLOHJ5QjJRb1NYQVRsaUVHWFhNeW1tRm8KeDBmem12THFscUxpNGZnOExhcTc5UC85aGxLU1djTWhJU0pYVTNHMS9KdEFBUmEyQXc4cXEzSVQvMkZ5NktrdwowMU9MdDdLN2pGUFRPaHhtdmNoSkZ1SVo1YXI0cW5HUlFHQnpCL2hoRHVIWEMwRlhJZ2ozd0NXMDZ4R2V2WjJyCmNCWWwwN1luL2lvZk95MU0wRjZZN0JrMU95N21vYndzM1JsalUyL2FpZlhLMmNOUlk2Qjl5WXNvd1RBZmQ5OTQKQ2YxSlF3TlhsaUZCeTZueEJLQk1YbDhIU1grS1o3L29PUlIwVXdJREFRQUJBb0lCQVFDbU5JSXR5ckhSY3oxdApJMGo0L0FQc295ZE1EL0owRkJMa2FoSUxKWjFaYW1tbmx4ZHh4WHBQUndXRnVXTEw2OTFVbDI5aUoxb1ptazU1Ci9ka2EvZlhnOUN3OUxXWVN2aExLdVlaMEZOTmhxZ3VoUEVBZy9uLytlR1ZCM2ZYZkxaZVZpK0E0L1VHMG15ZFMKVXVlQ2ZRSElZeWh4VkgvWnc3WER5WmNhVFVZVVdMUWlYcVN0Y2JRbnZFOXpwOGc5TWh5UkhBcWYwbEt2UTRqdwphUnNKTnlob3lhZWcvUXlFeHVYNGdBR1lIc1lTSDRFVmtXOHl5WE1aOHpRdk1OSUNiYXhmUkRuSngybUh6a09rCnFHczVXbFp5L3VrQk5jWTQwd2Y0eTY2bEVJaVpKbiswaFhtSTF4Tk5SdHRwMjZnY3ROOXZWbmVicTdLTnpjTDgKeFQrMXZJaEpBb0dCQU9iMVM1YlE4NVRFWDBoZTRmdXc2R3ExbnhRLzJUSU03emZhK2VhZThPQlh2eVNFV3JpdwpPZzM3RFhVUDFNVU1iTEJRenE0STl1dE5YSVZadEFLR0h6ZDR6WmtQeGxORjZPN0FyWnJEWElDNEdKZHdmSEhxCjJZcDkxUDlWekJlOVhkTVdZVGFCNkMzWVdtYzQwM08vYWdyRCtNb2JnL0hqMSt0d2xZR2hjdlV2QW9HQkFNWFMKT2VnWHc5VUF3VEZabFBtZzZKeDI3TzNXUFBHd1E3QzRnYitFZzZkR1pLRnJVR1ZId2VUUG1HaGtwN1BmYU5ESwplaFVoUWFnNm9XOTF4dkE2YldZZ29SQmczUWkxc01MblRWeTExeVN1UEVFSCtMT2s1N3d2akNLSk5XZnM0SjVyCmg1NGw0QXZ6UVhyWWN0UlZkSmYrNjFacGFnTkdZMVBvWVJMTHJMSWRBb0dBTndydzErRzJtNWJ0YW04S2hwU1QKMzVLbmRnajlkM3N6cStrcE03aGZpZWYvcXZGTU9jWHVJQlRjRVRFVHNWNlRyTFdsZkQ2d3NrVitybDFCbEhSbwpqaXpoT3dCU2NOZ3hlbTA3TXE0cXBwYTViYVltVW5QNUlwTjRwdDNJeFVPaFQ4UitxS0h2TnJYZ1hjZGlSYXl4CjFoejhkeFoxckxselpTNHd3M001MVlzQ2dZRUFpUDEwTEUySXg5Q2wrTTdZWTZZU2I0Zkx1MGhKRy9XOGFuemIKSFExZlBrOTVFRytJVlJyRUl2ZS95MHNvOTE4VzdyL0lteWxVbG5ORHFEUWZkK3grSmVNaXBuenRsRUorRGZxdgprQ3c4dUtJUUI5akZXV0l4T0JpVktyVnB6bll6ZG9Gd2dRd3BneDBKazFDZzlIblpMQWpVWUJyUDEwUy9ORFFRClJUdldjK0VDZ1lBeGRIZWxQNG1RdkJaS1oxMlNKbHlLbFVLeW43aHhzSHVkMkphMVNtS3FWeHBERDNlR0w0Y3QKZXA1QTZ5NkF4eGViZkI0aDdYNEZ0QTBBRURPdkZDR0J1QlRvZ3ZBdUNDVUtqK2JIUG1SNG53UVYzcWZ2M3loRAp0TGkwU2FHVElta2wvbUNCUDhZaW9JUys2N0xjby9kbHphUTNGVDlxTnJieFdFWjJlaS9LVlE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==";
 * const defaultHostShareKey = new alicloud.bastionhost.HostShareKey("defaultHostShareKey", {
 *     hostShareKeyName: name,
 *     instanceId: defaultInstance.id,
 *     privateKey: privateKey,
 * });
 * const defaultHostAccountShareKeyAttachment = new alicloud.bastionhost.HostAccountShareKeyAttachment("defaultHostAccountShareKeyAttachment", {
 *     instanceId: defaultInstance.id,
 *     hostShareKeyId: defaultHostShareKey.hostShareKeyId,
 *     hostAccountId: defaultHostAccount.hostAccountId,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Bastion Host Account Share Key Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment example <instance_id>:<host_share_key_id>:<host_account_id>
 * ```
 */
export class HostAccountShareKeyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing HostAccountShareKeyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostAccountShareKeyAttachmentState, opts?: pulumi.CustomResourceOptions): HostAccountShareKeyAttachment {
        return new HostAccountShareKeyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment';

    /**
     * Returns true if the given object is an instance of HostAccountShareKeyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostAccountShareKeyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostAccountShareKeyAttachment.__pulumiType;
    }

    /**
     * The ID list of the host account.
     */
    public readonly hostAccountId!: pulumi.Output<string>;
    /**
     * The ID of the host shared key.
     */
    public readonly hostShareKeyId!: pulumi.Output<string>;
    /**
     * The ID of the Bastion machine instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a HostAccountShareKeyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostAccountShareKeyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostAccountShareKeyAttachmentArgs | HostAccountShareKeyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostAccountShareKeyAttachmentState | undefined;
            resourceInputs["hostAccountId"] = state ? state.hostAccountId : undefined;
            resourceInputs["hostShareKeyId"] = state ? state.hostShareKeyId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as HostAccountShareKeyAttachmentArgs | undefined;
            if ((!args || args.hostAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostAccountId'");
            }
            if ((!args || args.hostShareKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostShareKeyId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["hostAccountId"] = args ? args.hostAccountId : undefined;
            resourceInputs["hostShareKeyId"] = args ? args.hostShareKeyId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostAccountShareKeyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostAccountShareKeyAttachment resources.
 */
export interface HostAccountShareKeyAttachmentState {
    /**
     * The ID list of the host account.
     */
    hostAccountId?: pulumi.Input<string>;
    /**
     * The ID of the host shared key.
     */
    hostShareKeyId?: pulumi.Input<string>;
    /**
     * The ID of the Bastion machine instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostAccountShareKeyAttachment resource.
 */
export interface HostAccountShareKeyAttachmentArgs {
    /**
     * The ID list of the host account.
     */
    hostAccountId: pulumi.Input<string>;
    /**
     * The ID of the host shared key.
     */
    hostShareKeyId: pulumi.Input<string>;
    /**
     * The ID of the Bastion machine instance.
     */
    instanceId: pulumi.Input<string>;
}

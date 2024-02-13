// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Ad Connector Directory resource.
 *
 * For information about ECD Ad Connector Directory and how to use it, see [What is Ad Connector Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createadconnectordirectory).
 *
 * > **NOTE:** Available since v1.174.0.
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
 * const defaultZones = alicloud.eds.getZones({});
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.ids?.[0]),
 *     vswitchName: name,
 * });
 * const defaultAdConnectorDirectory = new alicloud.eds.AdConnectorDirectory("defaultAdConnectorDirectory", {
 *     directoryName: name,
 *     desktopAccessType: "INTERNET",
 *     dnsAddresses: ["127.0.0.2"],
 *     domainName: "corp.example.com",
 *     domainPassword: "Example1234",
 *     domainUserName: "sAMAccountName",
 *     enableAdminAccess: false,
 *     mfaEnabled: false,
 *     specification: 1,
 *     subDomainDnsAddresses: ["127.0.0.3"],
 *     subDomainName: "child.example.com",
 *     vswitchIds: [defaultSwitch.id],
 * });
 * ```
 *
 * ## Import
 *
 * ECD Ad Connector Directory can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/adConnectorDirectory:AdConnectorDirectory example <id>
 * ```
 */
export class AdConnectorDirectory extends pulumi.CustomResource {
    /**
     * Get an existing AdConnectorDirectory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdConnectorDirectoryState, opts?: pulumi.CustomResourceOptions): AdConnectorDirectory {
        return new AdConnectorDirectory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/adConnectorDirectory:AdConnectorDirectory';

    /**
     * Returns true if the given object is an instance of AdConnectorDirectory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdConnectorDirectory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdConnectorDirectory.__pulumiType;
    }

    /**
     * The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     */
    public readonly desktopAccessType!: pulumi.Output<string>;
    /**
     * The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     */
    public readonly directoryName!: pulumi.Output<string>;
    /**
     * The DNS address list.
     */
    public readonly dnsAddresses!: pulumi.Output<string[]>;
    /**
     * The name of the domain.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The user password of the domain administrator. The maximum number of English characters is 64.
     */
    public readonly domainPassword!: pulumi.Output<string>;
    /**
     * The username of the domain administrator. The maximum number of English characters is 64.
     */
    public readonly domainUserName!: pulumi.Output<string>;
    /**
     * Whether to grant local administrator rights to users who use cloud desktops.
     */
    public readonly enableAdminAccess!: pulumi.Output<boolean>;
    /**
     * Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
     */
    public readonly mfaEnabled!: pulumi.Output<boolean>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    public readonly specification!: pulumi.Output<number | undefined>;
    /**
     * The status of directory.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
     */
    public readonly subDomainDnsAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
     */
    public readonly subDomainName!: pulumi.Output<string | undefined>;
    /**
     * List of VSwitch IDs in the directory.
     */
    public readonly vswitchIds!: pulumi.Output<string[]>;

    /**
     * Create a AdConnectorDirectory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdConnectorDirectoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdConnectorDirectoryArgs | AdConnectorDirectoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdConnectorDirectoryState | undefined;
            resourceInputs["desktopAccessType"] = state ? state.desktopAccessType : undefined;
            resourceInputs["directoryName"] = state ? state.directoryName : undefined;
            resourceInputs["dnsAddresses"] = state ? state.dnsAddresses : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["domainPassword"] = state ? state.domainPassword : undefined;
            resourceInputs["domainUserName"] = state ? state.domainUserName : undefined;
            resourceInputs["enableAdminAccess"] = state ? state.enableAdminAccess : undefined;
            resourceInputs["mfaEnabled"] = state ? state.mfaEnabled : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subDomainDnsAddresses"] = state ? state.subDomainDnsAddresses : undefined;
            resourceInputs["subDomainName"] = state ? state.subDomainName : undefined;
            resourceInputs["vswitchIds"] = state ? state.vswitchIds : undefined;
        } else {
            const args = argsOrState as AdConnectorDirectoryArgs | undefined;
            if ((!args || args.directoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'directoryName'");
            }
            if ((!args || args.dnsAddresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsAddresses'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.domainPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainPassword'");
            }
            if ((!args || args.domainUserName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainUserName'");
            }
            if ((!args || args.vswitchIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchIds'");
            }
            resourceInputs["desktopAccessType"] = args ? args.desktopAccessType : undefined;
            resourceInputs["directoryName"] = args ? args.directoryName : undefined;
            resourceInputs["dnsAddresses"] = args ? args.dnsAddresses : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainPassword"] = args?.domainPassword ? pulumi.secret(args.domainPassword) : undefined;
            resourceInputs["domainUserName"] = args ? args.domainUserName : undefined;
            resourceInputs["enableAdminAccess"] = args ? args.enableAdminAccess : undefined;
            resourceInputs["mfaEnabled"] = args ? args.mfaEnabled : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["subDomainDnsAddresses"] = args ? args.subDomainDnsAddresses : undefined;
            resourceInputs["subDomainName"] = args ? args.subDomainName : undefined;
            resourceInputs["vswitchIds"] = args ? args.vswitchIds : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["domainPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AdConnectorDirectory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdConnectorDirectory resources.
 */
export interface AdConnectorDirectoryState {
    /**
     * The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     */
    directoryName?: pulumi.Input<string>;
    /**
     * The DNS address list.
     */
    dnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the domain.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The user password of the domain administrator. The maximum number of English characters is 64.
     */
    domainPassword?: pulumi.Input<string>;
    /**
     * The username of the domain administrator. The maximum number of English characters is 64.
     */
    domainUserName?: pulumi.Input<string>;
    /**
     * Whether to grant local administrator rights to users who use cloud desktops.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    specification?: pulumi.Input<number>;
    /**
     * The status of directory.
     */
    status?: pulumi.Input<string>;
    /**
     * The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
     */
    subDomainDnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
     */
    subDomainName?: pulumi.Input<string>;
    /**
     * List of VSwitch IDs in the directory.
     */
    vswitchIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AdConnectorDirectory resource.
 */
export interface AdConnectorDirectoryArgs {
    /**
     * The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     */
    directoryName: pulumi.Input<string>;
    /**
     * The DNS address list.
     */
    dnsAddresses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the domain.
     */
    domainName: pulumi.Input<string>;
    /**
     * The user password of the domain administrator. The maximum number of English characters is 64.
     */
    domainPassword: pulumi.Input<string>;
    /**
     * The username of the domain administrator. The maximum number of English characters is 64.
     */
    domainUserName: pulumi.Input<string>;
    /**
     * Whether to grant local administrator rights to users who use cloud desktops.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    specification?: pulumi.Input<number>;
    /**
     * The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
     */
    subDomainDnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
     */
    subDomainName?: pulumi.Input<string>;
    /**
     * List of VSwitch IDs in the directory.
     */
    vswitchIds: pulumi.Input<pulumi.Input<string>[]>;
}

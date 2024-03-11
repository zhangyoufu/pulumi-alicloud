// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Ad Connector Office Site resource.
 *
 * For information about ECD Ad Connector Office Site and how to use it, see [What is Ad Connector Office Site](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createadconnectorofficesite).
 *
 * > **NOTE:** Available since v1.176.0.
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
 * const defaultInstance = new alicloud.cen.Instance("defaultInstance", {
 *     cenInstanceName: name,
 *     protectionLevel: "REDUCED",
 * });
 * const defaultAdConnectorOfficeSite = new alicloud.eds.AdConnectorOfficeSite("defaultAdConnectorOfficeSite", {
 *     adConnectorOfficeSiteName: name,
 *     bandwidth: 100,
 *     cenId: defaultInstance.id,
 *     cidrBlock: "10.0.0.0/12",
 *     desktopAccessType: "INTERNET",
 *     dnsAddresses: ["127.0.0.2"],
 *     domainName: "corp.example.com",
 *     domainPassword: "Example1234",
 *     domainUserName: "sAMAccountName",
 *     enableAdminAccess: false,
 *     enableInternetAccess: false,
 *     mfaEnabled: false,
 *     subDomainDnsAddresses: ["127.0.0.3"],
 *     subDomainName: "child.example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ECD Ad Connector Office Site can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/adConnectorOfficeSite:AdConnectorOfficeSite example <id>
 * ```
 */
export class AdConnectorOfficeSite extends pulumi.CustomResource {
    /**
     * Get an existing AdConnectorOfficeSite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdConnectorOfficeSiteState, opts?: pulumi.CustomResourceOptions): AdConnectorOfficeSite {
        return new AdConnectorOfficeSite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/adConnectorOfficeSite:AdConnectorOfficeSite';

    /**
     * Returns true if the given object is an instance of AdConnectorOfficeSite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdConnectorOfficeSite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdConnectorOfficeSite.__pulumiType;
    }

    /**
     * The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
     */
    public readonly adConnectorOfficeSiteName!: pulumi.Output<string>;
    /**
     * The ad hostname.
     */
    public readonly adHostname!: pulumi.Output<string | undefined>;
    /**
     * The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * The ID of the CEN instance.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The cen owner id.
     */
    public readonly cenOwnerId!: pulumi.Output<string | undefined>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
     */
    public readonly desktopAccessType!: pulumi.Output<string>;
    /**
     * The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
     */
    public readonly dnsAddresses!: pulumi.Output<string[]>;
    /**
     * The domain name of the enterprise AD system. You can register each domain name only once.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The password of the domain administrator. The password can be up to 64 characters in length.
     */
    public readonly domainPassword!: pulumi.Output<string | undefined>;
    /**
     * The username of the domain administrator. The username can be up to 64 characters in length.
     */
    public readonly domainUserName!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
     */
    public readonly enableAdminAccess!: pulumi.Output<boolean>;
    /**
     * Specifies whether to enable Internet access.
     */
    public readonly enableInternetAccess!: pulumi.Output<boolean>;
    /**
     * Specifies whether to enable multi-factor authentication (MFA).
     */
    public readonly mfaEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The protocol type. Valid values: `ASP`, `HDX`.
     */
    public readonly protocolType!: pulumi.Output<string | undefined>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    public readonly specification!: pulumi.Output<number | undefined>;
    /**
     * The resource State.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
     */
    public readonly subDomainDnsAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * The domain name of the enterprise AD subdomain.
     */
    public readonly subDomainName!: pulumi.Output<string | undefined>;
    /**
     * The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
     */
    public readonly verifyCode!: pulumi.Output<string | undefined>;

    /**
     * Create a AdConnectorOfficeSite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdConnectorOfficeSiteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdConnectorOfficeSiteArgs | AdConnectorOfficeSiteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdConnectorOfficeSiteState | undefined;
            resourceInputs["adConnectorOfficeSiteName"] = state ? state.adConnectorOfficeSiteName : undefined;
            resourceInputs["adHostname"] = state ? state.adHostname : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["cenOwnerId"] = state ? state.cenOwnerId : undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["desktopAccessType"] = state ? state.desktopAccessType : undefined;
            resourceInputs["dnsAddresses"] = state ? state.dnsAddresses : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["domainPassword"] = state ? state.domainPassword : undefined;
            resourceInputs["domainUserName"] = state ? state.domainUserName : undefined;
            resourceInputs["enableAdminAccess"] = state ? state.enableAdminAccess : undefined;
            resourceInputs["enableInternetAccess"] = state ? state.enableInternetAccess : undefined;
            resourceInputs["mfaEnabled"] = state ? state.mfaEnabled : undefined;
            resourceInputs["protocolType"] = state ? state.protocolType : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subDomainDnsAddresses"] = state ? state.subDomainDnsAddresses : undefined;
            resourceInputs["subDomainName"] = state ? state.subDomainName : undefined;
            resourceInputs["verifyCode"] = state ? state.verifyCode : undefined;
        } else {
            const args = argsOrState as AdConnectorOfficeSiteArgs | undefined;
            if ((!args || args.adConnectorOfficeSiteName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adConnectorOfficeSiteName'");
            }
            if ((!args || args.cenId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if ((!args || args.dnsAddresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsAddresses'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["adConnectorOfficeSiteName"] = args ? args.adConnectorOfficeSiteName : undefined;
            resourceInputs["adHostname"] = args ? args.adHostname : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["cenOwnerId"] = args ? args.cenOwnerId : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["desktopAccessType"] = args ? args.desktopAccessType : undefined;
            resourceInputs["dnsAddresses"] = args ? args.dnsAddresses : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainPassword"] = args?.domainPassword ? pulumi.secret(args.domainPassword) : undefined;
            resourceInputs["domainUserName"] = args ? args.domainUserName : undefined;
            resourceInputs["enableAdminAccess"] = args ? args.enableAdminAccess : undefined;
            resourceInputs["enableInternetAccess"] = args ? args.enableInternetAccess : undefined;
            resourceInputs["mfaEnabled"] = args ? args.mfaEnabled : undefined;
            resourceInputs["protocolType"] = args ? args.protocolType : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["subDomainDnsAddresses"] = args ? args.subDomainDnsAddresses : undefined;
            resourceInputs["subDomainName"] = args ? args.subDomainName : undefined;
            resourceInputs["verifyCode"] = args ? args.verifyCode : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["domainPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AdConnectorOfficeSite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdConnectorOfficeSite resources.
 */
export interface AdConnectorOfficeSiteState {
    /**
     * The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
     */
    adConnectorOfficeSiteName?: pulumi.Input<string>;
    /**
     * The ad hostname.
     */
    adHostname?: pulumi.Input<string>;
    /**
     * The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The ID of the CEN instance.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The cen owner id.
     */
    cenOwnerId?: pulumi.Input<string>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
     */
    dnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain name of the enterprise AD system. You can register each domain name only once.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The password of the domain administrator. The password can be up to 64 characters in length.
     */
    domainPassword?: pulumi.Input<string>;
    /**
     * The username of the domain administrator. The username can be up to 64 characters in length.
     */
    domainUserName?: pulumi.Input<string>;
    /**
     * Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable Internet access.
     */
    enableInternetAccess?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable multi-factor authentication (MFA).
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The protocol type. Valid values: `ASP`, `HDX`.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    specification?: pulumi.Input<number>;
    /**
     * The resource State.
     */
    status?: pulumi.Input<string>;
    /**
     * The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
     */
    subDomainDnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain name of the enterprise AD subdomain.
     */
    subDomainName?: pulumi.Input<string>;
    /**
     * The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
     */
    verifyCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdConnectorOfficeSite resource.
 */
export interface AdConnectorOfficeSiteArgs {
    /**
     * The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
     */
    adConnectorOfficeSiteName: pulumi.Input<string>;
    /**
     * The ad hostname.
     */
    adHostname?: pulumi.Input<string>;
    /**
     * The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The ID of the CEN instance.
     */
    cenId: pulumi.Input<string>;
    /**
     * The cen owner id.
     */
    cenOwnerId?: pulumi.Input<string>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    cidrBlock: pulumi.Input<string>;
    /**
     * The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
     */
    dnsAddresses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain name of the enterprise AD system. You can register each domain name only once.
     */
    domainName: pulumi.Input<string>;
    /**
     * The password of the domain administrator. The password can be up to 64 characters in length.
     */
    domainPassword?: pulumi.Input<string>;
    /**
     * The username of the domain administrator. The username can be up to 64 characters in length.
     */
    domainUserName?: pulumi.Input<string>;
    /**
     * Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable Internet access.
     */
    enableInternetAccess?: pulumi.Input<boolean>;
    /**
     * Specifies whether to enable multi-factor authentication (MFA).
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The protocol type. Valid values: `ASP`, `HDX`.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     */
    specification?: pulumi.Input<number>;
    /**
     * The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
     */
    subDomainDnsAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain name of the enterprise AD subdomain.
     */
    subDomainName?: pulumi.Input<string>;
    /**
     * The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
     */
    verifyCode?: pulumi.Input<string>;
}

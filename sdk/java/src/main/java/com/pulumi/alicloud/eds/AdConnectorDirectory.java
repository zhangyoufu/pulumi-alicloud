// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.AdConnectorDirectoryArgs;
import com.pulumi.alicloud.eds.inputs.AdConnectorDirectoryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECD Ad Connector Directory resource.
 * 
 * For information about ECD Ad Connector Directory and how to use it, see [What is Ad Connector Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createadconnectordirectory).
 * 
 * &gt; **NOTE:** Available since v1.174.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.eds.EdsFunctions;
 * import com.pulumi.alicloud.eds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.eds.AdConnectorDirectory;
 * import com.pulumi.alicloud.eds.AdConnectorDirectoryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get("name").orElse("terraform-example");
 *         final var default = EdsFunctions.getZones();
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName(name)
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.0.0/24")
 *             .zoneId(default_.ids()[0])
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultAdConnectorDirectory = new AdConnectorDirectory("defaultAdConnectorDirectory", AdConnectorDirectoryArgs.builder()
 *             .directoryName(String.format("%s-%s", name,defaultInteger.result()))
 *             .desktopAccessType("INTERNET")
 *             .dnsAddresses("127.0.0.2")
 *             .domainName("corp.example.com")
 *             .domainPassword("Example1234")
 *             .domainUserName("sAMAccountName")
 *             .enableAdminAccess(false)
 *             .mfaEnabled(false)
 *             .specification(1)
 *             .subDomainDnsAddresses("127.0.0.3")
 *             .subDomainName("child.example.com")
 *             .vswitchIds(defaultSwitch.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ECD Ad Connector Directory can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/adConnectorDirectory:AdConnectorDirectory example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/adConnectorDirectory:AdConnectorDirectory")
public class AdConnectorDirectory extends com.pulumi.resources.CustomResource {
    /**
     * The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     * 
     */
    @Export(name="desktopAccessType", refs={String.class}, tree="[0]")
    private Output<String> desktopAccessType;

    /**
     * @return The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
     * 
     */
    public Output<String> desktopAccessType() {
        return this.desktopAccessType;
    }
    /**
     * The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="directoryName", refs={String.class}, tree="[0]")
    private Output<String> directoryName;

    /**
     * @return The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    public Output<String> directoryName() {
        return this.directoryName;
    }
    /**
     * The DNS address list.
     * 
     */
    @Export(name="dnsAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> dnsAddresses;

    /**
     * @return The DNS address list.
     * 
     */
    public Output<List<String>> dnsAddresses() {
        return this.dnsAddresses;
    }
    /**
     * The name of the domain.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The name of the domain.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The user password of the domain administrator. The maximum number of English characters is 64.
     * 
     */
    @Export(name="domainPassword", refs={String.class}, tree="[0]")
    private Output<String> domainPassword;

    /**
     * @return The user password of the domain administrator. The maximum number of English characters is 64.
     * 
     */
    public Output<String> domainPassword() {
        return this.domainPassword;
    }
    /**
     * The username of the domain administrator. The maximum number of English characters is 64.
     * 
     */
    @Export(name="domainUserName", refs={String.class}, tree="[0]")
    private Output<String> domainUserName;

    /**
     * @return The username of the domain administrator. The maximum number of English characters is 64.
     * 
     */
    public Output<String> domainUserName() {
        return this.domainUserName;
    }
    /**
     * Whether to grant local administrator rights to users who use cloud desktops.
     * 
     */
    @Export(name="enableAdminAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableAdminAccess;

    /**
     * @return Whether to grant local administrator rights to users who use cloud desktops.
     * 
     */
    public Output<Boolean> enableAdminAccess() {
        return this.enableAdminAccess;
    }
    /**
     * Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
     * 
     */
    @Export(name="mfaEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mfaEnabled;

    /**
     * @return Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
     * 
     */
    public Output<Boolean> mfaEnabled() {
        return this.mfaEnabled;
    }
    /**
     * The AD Connector specifications. Valid values: `1`, `2`.
     * 
     */
    @Export(name="specification", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> specification;

    /**
     * @return The AD Connector specifications. Valid values: `1`, `2`.
     * 
     */
    public Output<Optional<Integer>> specification() {
        return Codegen.optional(this.specification);
    }
    /**
     * The status of directory.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of directory.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The Enterprise already has the DNS address of the AD subdomain. If `sub_domain_name` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
     * 
     */
    @Export(name="subDomainDnsAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> subDomainDnsAddresses;

    /**
     * @return The Enterprise already has the DNS address of the AD subdomain. If `sub_domain_name` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
     * 
     */
    public Output<Optional<List<String>>> subDomainDnsAddresses() {
        return Codegen.optional(this.subDomainDnsAddresses);
    }
    /**
     * The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
     * 
     */
    @Export(name="subDomainName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subDomainName;

    /**
     * @return The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
     * 
     */
    public Output<Optional<String>> subDomainName() {
        return Codegen.optional(this.subDomainName);
    }
    /**
     * List of VSwitch IDs in the directory.
     * 
     */
    @Export(name="vswitchIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> vswitchIds;

    /**
     * @return List of VSwitch IDs in the directory.
     * 
     */
    public Output<List<String>> vswitchIds() {
        return this.vswitchIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AdConnectorDirectory(java.lang.String name) {
        this(name, AdConnectorDirectoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AdConnectorDirectory(java.lang.String name, AdConnectorDirectoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AdConnectorDirectory(java.lang.String name, AdConnectorDirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AdConnectorDirectory(java.lang.String name, Output<java.lang.String> id, @Nullable AdConnectorDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, state, makeResourceOptions(options, id), false);
    }

    private static AdConnectorDirectoryArgs makeArgs(AdConnectorDirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AdConnectorDirectoryArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "domainPassword"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AdConnectorDirectory get(java.lang.String name, Output<java.lang.String> id, @Nullable AdConnectorDirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AdConnectorDirectory(name, id, state, options);
    }
}

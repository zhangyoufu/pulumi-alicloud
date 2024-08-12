// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.SimpleOfficeSiteArgs;
import com.pulumi.alicloud.eds.inputs.SimpleOfficeSiteState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECD Simple Office Site resource.
 * 
 * For information about ECD Simple Office Site and how to use it, see [What is Simple Office Site](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createsimpleofficesite).
 * 
 * &gt; **NOTE:** Available since v1.140.0.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.eds.SimpleOfficeSite;
 * import com.pulumi.alicloud.eds.SimpleOfficeSiteArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultSimpleOfficeSite = new SimpleOfficeSite("defaultSimpleOfficeSite", SimpleOfficeSiteArgs.builder()
 *             .cidrBlock("172.16.0.0/12")
 *             .enableAdminAccess(true)
 *             .desktopAccessType("Internet")
 *             .officeSiteName(String.format("terraform-example-%s", default_.result()))
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
 * ECD Simple Office Site can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/simpleOfficeSite:SimpleOfficeSite example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/simpleOfficeSite:SimpleOfficeSite")
public class SimpleOfficeSite extends com.pulumi.resources.CustomResource {
    /**
     * The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
     * 
     * @deprecated
     * Field &#39;bandwidth&#39; has been deprecated from provider version 1.142.0.
     * 
     */
    @Deprecated /* Field 'bandwidth' has been deprecated from provider version 1.142.0. */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * Cloud Enterprise Network Instance ID.
     * 
     */
    @Export(name="cenId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cenId;

    /**
     * @return Cloud Enterprise Network Instance ID.
     * 
     */
    public Output<Optional<String>> cenId() {
        return Codegen.optional(this.cenId);
    }
    /**
     * The cen owner id.
     * 
     */
    @Export(name="cenOwnerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cenOwnerId;

    /**
     * @return The cen owner id.
     * 
     */
    public Output<Optional<String>> cenOwnerId() {
        return Codegen.optional(this.cenOwnerId);
    }
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     * 
     */
    @Export(name="cidrBlock", refs={String.class}, tree="[0]")
    private Output<String> cidrBlock;

    /**
     * @return Workspace Corresponds to the Security Office Network of IPv4 Segment.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
     * 
     */
    @Export(name="desktopAccessType", refs={String.class}, tree="[0]")
    private Output<String> desktopAccessType;

    /**
     * @return Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
     * 
     */
    public Output<String> desktopAccessType() {
        return this.desktopAccessType;
    }
    /**
     * Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
     * 
     */
    @Export(name="enableAdminAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableAdminAccess;

    /**
     * @return Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
     * 
     */
    public Output<Boolean> enableAdminAccess() {
        return this.enableAdminAccess;
    }
    /**
     * Enable Cross-Desktop Access.
     * 
     */
    @Export(name="enableCrossDesktopAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableCrossDesktopAccess;

    /**
     * @return Enable Cross-Desktop Access.
     * 
     */
    public Output<Boolean> enableCrossDesktopAccess() {
        return this.enableCrossDesktopAccess;
    }
    /**
     * Whether the Open Internet Access Function.
     * 
     * @deprecated
     * Field &#39;enable_internet_access&#39; has been deprecated from provider version 1.142.0.
     * 
     */
    @Deprecated /* Field 'enable_internet_access' has been deprecated from provider version 1.142.0. */
    @Export(name="enableInternetAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableInternetAccess;

    /**
     * @return Whether the Open Internet Access Function.
     * 
     */
    public Output<Boolean> enableInternetAccess() {
        return this.enableInternetAccess;
    }
    /**
     * Whether to Enable Multi-Factor Authentication MFA.
     * 
     */
    @Export(name="mfaEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> mfaEnabled;

    /**
     * @return Whether to Enable Multi-Factor Authentication MFA.
     * 
     */
    public Output<Boolean> mfaEnabled() {
        return this.mfaEnabled;
    }
    /**
     * The office site name.
     * 
     */
    @Export(name="officeSiteName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> officeSiteName;

    /**
     * @return The office site name.
     * 
     */
    public Output<Optional<String>> officeSiteName() {
        return Codegen.optional(this.officeSiteName);
    }
    /**
     * Whether to Enable Single Sign-on (SSO) for User-Based SSO.
     * 
     */
    @Export(name="ssoEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ssoEnabled;

    /**
     * @return Whether to Enable Single Sign-on (SSO) for User-Based SSO.
     * 
     */
    public Output<Boolean> ssoEnabled() {
        return this.ssoEnabled;
    }
    /**
     * Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SimpleOfficeSite(java.lang.String name) {
        this(name, SimpleOfficeSiteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SimpleOfficeSite(java.lang.String name, SimpleOfficeSiteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SimpleOfficeSite(java.lang.String name, SimpleOfficeSiteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/simpleOfficeSite:SimpleOfficeSite", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SimpleOfficeSite(java.lang.String name, Output<java.lang.String> id, @Nullable SimpleOfficeSiteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/simpleOfficeSite:SimpleOfficeSite", name, state, makeResourceOptions(options, id), false);
    }

    private static SimpleOfficeSiteArgs makeArgs(SimpleOfficeSiteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SimpleOfficeSiteArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static SimpleOfficeSite get(java.lang.String name, Output<java.lang.String> id, @Nullable SimpleOfficeSiteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SimpleOfficeSite(name, id, state, options);
    }
}

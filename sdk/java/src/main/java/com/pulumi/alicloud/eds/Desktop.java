// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eds.DesktopArgs;
import com.pulumi.alicloud.eds.inputs.DesktopState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECD Desktop resource.
 * 
 * For information about ECD Desktop and how to use it, see [What is Desktop](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createdesktops)
 * 
 * &gt; **NOTE:** Available since v1.144.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.eds.SimpleOfficeSite;
 * import com.pulumi.alicloud.eds.SimpleOfficeSiteArgs;
 * import com.pulumi.alicloud.eds.EcdPolicyGroup;
 * import com.pulumi.alicloud.eds.EcdPolicyGroupArgs;
 * import com.pulumi.alicloud.eds.inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs;
 * import com.pulumi.alicloud.eds.inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs;
 * import com.pulumi.alicloud.eds.EdsFunctions;
 * import com.pulumi.alicloud.eds.inputs.GetBundlesArgs;
 * import com.pulumi.alicloud.eds.Desktop;
 * import com.pulumi.alicloud.eds.DesktopArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var defaultRandomInteger = new RandomInteger(&#34;defaultRandomInteger&#34;, RandomIntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultSimpleOfficeSite = new SimpleOfficeSite(&#34;defaultSimpleOfficeSite&#34;, SimpleOfficeSiteArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .enableAdminAccess(true)
 *             .desktopAccessType(&#34;Internet&#34;)
 *             .officeSiteName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s-%s&#34;, name,result)))
 *             .build());
 * 
 *         var defaultEcdPolicyGroup = new EcdPolicyGroup(&#34;defaultEcdPolicyGroup&#34;, EcdPolicyGroupArgs.builder()        
 *             .policyGroupName(name)
 *             .clipboard(&#34;read&#34;)
 *             .localDrive(&#34;read&#34;)
 *             .usbRedirect(&#34;off&#34;)
 *             .watermark(&#34;off&#34;)
 *             .authorizeAccessPolicyRules(EcdPolicyGroupAuthorizeAccessPolicyRuleArgs.builder()
 *                 .description(name)
 *                 .cidrIp(&#34;1.2.3.45/24&#34;)
 *                 .build())
 *             .authorizeSecurityPolicyRules(EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs.builder()
 *                 .type(&#34;inflow&#34;)
 *                 .policy(&#34;accept&#34;)
 *                 .description(name)
 *                 .portRange(&#34;80/80&#34;)
 *                 .ipProtocol(&#34;TCP&#34;)
 *                 .priority(&#34;1&#34;)
 *                 .cidrIp(&#34;1.2.3.4/24&#34;)
 *                 .build())
 *             .build());
 * 
 *         final var defaultBundles = EdsFunctions.getBundles(GetBundlesArgs.builder()
 *             .bundleType(&#34;SYSTEM&#34;)
 *             .build());
 * 
 *         var defaultDesktop = new Desktop(&#34;defaultDesktop&#34;, DesktopArgs.builder()        
 *             .officeSiteId(defaultSimpleOfficeSite.id())
 *             .policyGroupId(defaultEcdPolicyGroup.id())
 *             .bundleId(defaultBundles.applyValue(getBundlesResult -&gt; getBundlesResult.bundles()[1].id()))
 *             .desktopName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ECD Desktop can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eds/desktop:Desktop example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eds/desktop:Desktop")
public class Desktop extends com.pulumi.resources.CustomResource {
    /**
     * The amount of the Desktop.
     * 
     */
    @Export(name="amount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> amount;

    /**
     * @return The amount of the Desktop.
     * 
     */
    public Output<Optional<Integer>> amount() {
        return Codegen.optional(this.amount);
    }
    /**
     * The auto-pay of the Desktop whether to pay automatically. values: `true`, `false`.
     * 
     */
    @Export(name="autoPay", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoPay;

    /**
     * @return The auto-pay of the Desktop whether to pay automatically. values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * The auto-renewal of the Desktop whether to renew automatically. It takes effect only when the parameter ChargeType is set to PrePaid. values: `true`, `false`.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return The auto-renewal of the Desktop whether to renew automatically. It takes effect only when the parameter ChargeType is set to PrePaid. values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The bundle id of the Desktop.
     * 
     */
    @Export(name="bundleId", refs={String.class}, tree="[0]")
    private Output<String> bundleId;

    /**
     * @return The bundle id of the Desktop.
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }
    /**
     * The desktop name of the Desktop.
     * 
     */
    @Export(name="desktopName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desktopName;

    /**
     * @return The desktop name of the Desktop.
     * 
     */
    public Output<Optional<String>> desktopName() {
        return Codegen.optional(this.desktopName);
    }
    /**
     * The desktop type of the Desktop.
     * 
     */
    @Export(name="desktopType", refs={String.class}, tree="[0]")
    private Output<String> desktopType;

    /**
     * @return The desktop type of the Desktop.
     * 
     */
    public Output<String> desktopType() {
        return this.desktopType;
    }
    /**
     * The desktop end user id of the Desktop.
     * 
     */
    @Export(name="endUserIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> endUserIds;

    /**
     * @return The desktop end user id of the Desktop.
     * 
     */
    public Output<Optional<List<String>>> endUserIds() {
        return Codegen.optional(this.endUserIds);
    }
    /**
     * The hostname of the Desktop.
     * 
     */
    @Export(name="hostName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostName;

    /**
     * @return The hostname of the Desktop.
     * 
     */
    public Output<Optional<String>> hostName() {
        return Codegen.optional(this.hostName);
    }
    /**
     * The ID of the Simple Office Site.
     * 
     */
    @Export(name="officeSiteId", refs={String.class}, tree="[0]")
    private Output<String> officeSiteId;

    /**
     * @return The ID of the Simple Office Site.
     * 
     */
    public Output<String> officeSiteId() {
        return this.officeSiteId;
    }
    /**
     * The payment type of the Desktop. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The payment type of the Desktop. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The period of the Desktop.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The period of the Desktop.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The period unit of the Desktop.
     * 
     */
    @Export(name="periodUnit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> periodUnit;

    /**
     * @return The period unit of the Desktop.
     * 
     */
    public Output<Optional<String>> periodUnit() {
        return Codegen.optional(this.periodUnit);
    }
    /**
     * The policy group id of the Desktop.
     * 
     */
    @Export(name="policyGroupId", refs={String.class}, tree="[0]")
    private Output<String> policyGroupId;

    /**
     * @return The policy group id of the Desktop.
     * 
     */
    public Output<String> policyGroupId() {
        return this.policyGroupId;
    }
    /**
     * The root disk size gib of the Desktop.
     * 
     */
    @Export(name="rootDiskSizeGib", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rootDiskSizeGib;

    /**
     * @return The root disk size gib of the Desktop.
     * 
     */
    public Output<Optional<Integer>> rootDiskSizeGib() {
        return Codegen.optional(this.rootDiskSizeGib);
    }
    /**
     * The status of the Desktop. Valid values: `Deleted`, `Expired`, `Pending`, `Running`, `Starting`, `Stopped`, `Stopping`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Desktop. Valid values: `Deleted`, `Expired`, `Pending`, `Running`, `Starting`, `Stopped`, `Stopping`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The stopped mode of the Desktop.
     * 
     */
    @Export(name="stoppedMode", refs={String.class}, tree="[0]")
    private Output<String> stoppedMode;

    /**
     * @return The stopped mode of the Desktop.
     * 
     */
    public Output<String> stoppedMode() {
        return this.stoppedMode;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The user assign mode of the Desktop. Valid values: `ALL`, `PER_USER`. Default to `ALL`.
     * 
     */
    @Export(name="userAssignMode", refs={String.class}, tree="[0]")
    private Output<String> userAssignMode;

    /**
     * @return The user assign mode of the Desktop. Valid values: `ALL`, `PER_USER`. Default to `ALL`.
     * 
     */
    public Output<String> userAssignMode() {
        return this.userAssignMode;
    }
    /**
     * The user disk size gib of the Desktop.
     * 
     */
    @Export(name="userDiskSizeGib", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> userDiskSizeGib;

    /**
     * @return The user disk size gib of the Desktop.
     * 
     */
    public Output<Optional<Integer>> userDiskSizeGib() {
        return Codegen.optional(this.userDiskSizeGib);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Desktop(String name) {
        this(name, DesktopArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Desktop(String name, DesktopArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Desktop(String name, DesktopArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/desktop:Desktop", name, args == null ? DesktopArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Desktop(String name, Output<String> id, @Nullable DesktopState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eds/desktop:Desktop", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Desktop get(String name, Output<String> id, @Nullable DesktopState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Desktop(name, id, state, options);
    }
}

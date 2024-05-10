// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.AcceleratorArgs;
import com.pulumi.alicloud.ga.inputs.AcceleratorState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Accelerator resource.
 * 
 * For information about Global Accelerator (GA) Accelerator and how to use it, see [What is Accelerator](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createaccelerator).
 * 
 * &gt; **NOTE:** Available since v1.111.0.
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
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
 *         var example = new Accelerator("example", AcceleratorArgs.builder()        
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec("1")
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
 * Ga Accelerator can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/accelerator:Accelerator example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/accelerator:Accelerator")
public class Accelerator extends com.pulumi.resources.CustomResource {
    /**
     * The Name of the GA instance.
     * 
     */
    @Export(name="acceleratorName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> acceleratorName;

    /**
     * @return The Name of the GA instance.
     * 
     */
    public Output<Optional<String>> acceleratorName() {
        return Codegen.optional(this.acceleratorName);
    }
    /**
     * Auto renewal period of an instance, in the unit of month. The value range is 1-12.
     * 
     */
    @Export(name="autoRenewDuration", refs={Integer.class}, tree="[0]")
    private Output<Integer> autoRenewDuration;

    /**
     * @return Auto renewal period of an instance, in the unit of month. The value range is 1-12.
     * 
     */
    public Output<Integer> autoRenewDuration() {
        return this.autoRenewDuration;
    }
    /**
     * Use coupons to pay bills automatically. Default value: `false`. Valid values:
     * 
     */
    @Export(name="autoUseCoupon", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoUseCoupon;

    /**
     * @return Use coupons to pay bills automatically. Default value: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> autoUseCoupon() {
        return Codegen.optional(this.autoUseCoupon);
    }
    /**
     * The bandwidth billing method. Default value: `BandwidthPackage`. Valid values:
     * 
     */
    @Export(name="bandwidthBillingType", refs={String.class}, tree="[0]")
    private Output<String> bandwidthBillingType;

    /**
     * @return The bandwidth billing method. Default value: `BandwidthPackage`. Valid values:
     * 
     */
    public Output<String> bandwidthBillingType() {
        return this.bandwidthBillingType;
    }
    /**
     * The type of cross-border acceleration. Default value: `bgpPro`. Valid values: `bgpPro`, `private`. **NOTE:** `cross_border_mode` is valid only when `cross_border_status` is set to `true`.
     * 
     */
    @Export(name="crossBorderMode", refs={String.class}, tree="[0]")
    private Output<String> crossBorderMode;

    /**
     * @return The type of cross-border acceleration. Default value: `bgpPro`. Valid values: `bgpPro`, `private`. **NOTE:** `cross_border_mode` is valid only when `cross_border_status` is set to `true`.
     * 
     */
    public Output<String> crossBorderMode() {
        return this.crossBorderMode;
    }
    /**
     * Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
     * 
     */
    @Export(name="crossBorderStatus", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> crossBorderStatus;

    /**
     * @return Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> crossBorderStatus() {
        return Codegen.optional(this.crossBorderStatus);
    }
    /**
     * Descriptive information of the global acceleration instance.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Descriptive information of the global acceleration instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The subscription duration.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
     * 
     */
    @Export(name="duration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> duration;

    /**
     * @return The subscription duration.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
     * 
     */
    public Output<Optional<Integer>> duration() {
        return Codegen.optional(this.duration);
    }
    /**
     * The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The billing cycle of the GA instance. Default value: `Month`. Valid values:
     * 
     */
    @Export(name="pricingCycle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pricingCycle;

    /**
     * @return The billing cycle of the GA instance. Default value: `Month`. Valid values:
     * 
     */
    public Output<Optional<String>> pricingCycle() {
        return Codegen.optional(this.pricingCycle);
    }
    /**
     * The code of the coupon. **NOTE:** The `promotion_option_no` takes effect only for accounts registered on the international site (alibabacloud.com).
     * 
     */
    @Export(name="promotionOptionNo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> promotionOptionNo;

    /**
     * @return The code of the coupon. **NOTE:** The `promotion_option_no` takes effect only for accounts registered on the international site (alibabacloud.com).
     * 
     */
    public Output<Optional<String>> promotionOptionNo() {
        return Codegen.optional(this.promotionOptionNo);
    }
    /**
     * Whether to renew an accelerator automatically or not. Default value: `Normal`. Valid values:
     * 
     */
    @Export(name="renewalStatus", refs={String.class}, tree="[0]")
    private Output<String> renewalStatus;

    /**
     * @return Whether to renew an accelerator automatically or not. Default value: `Normal`. Valid values:
     * 
     */
    public Output<String> renewalStatus() {
        return this.renewalStatus;
    }
    /**
     * The instance type of the GA instance. Specification of global acceleration instance. Valid values:
     * 
     */
    @Export(name="spec", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> spec;

    /**
     * @return The instance type of the GA instance. Specification of global acceleration instance. Valid values:
     * 
     */
    public Output<Optional<String>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * The status of the GA instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the GA instance.
     * 
     */
    public Output<String> status() {
        return this.status;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Accelerator(String name) {
        this(name, AcceleratorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Accelerator(String name, @Nullable AcceleratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Accelerator(String name, @Nullable AcceleratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/accelerator:Accelerator", name, args == null ? AcceleratorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Accelerator(String name, Output<String> id, @Nullable AcceleratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/accelerator:Accelerator", name, state, makeResourceOptions(options, id));
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
    public static Accelerator get(String name, Output<String> id, @Nullable AcceleratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Accelerator(name, id, state, options);
    }
}

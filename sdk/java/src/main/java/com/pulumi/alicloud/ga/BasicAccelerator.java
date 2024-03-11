// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.BasicAcceleratorArgs;
import com.pulumi.alicloud.ga.inputs.BasicAcceleratorState;
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
 * Provides a Global Accelerator (GA) Basic Accelerator resource.
 * 
 * For information about Global Accelerator (GA) Basic Accelerator and how to use it, see [What is Basic Accelerator](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicaccelerator).
 * 
 * &gt; **NOTE:** Available since v1.194.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ga.BasicAccelerator;
 * import com.pulumi.alicloud.ga.BasicAcceleratorArgs;
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
 *         var default_ = new BasicAccelerator(&#34;default&#34;, BasicAcceleratorArgs.builder()        
 *             .autoPay(true)
 *             .autoUseCoupon(&#34;true&#34;)
 *             .bandwidthBillingType(&#34;BandwidthPackage&#34;)
 *             .basicAcceleratorName(&#34;tf-example-value&#34;)
 *             .description(&#34;tf-example-value&#34;)
 *             .duration(1)
 *             .pricingCycle(&#34;Month&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Global Accelerator (GA) Basic Accelerator can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ga/basicAccelerator:BasicAccelerator example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/basicAccelerator:BasicAccelerator")
public class BasicAccelerator extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to enable automatic payment. Default value: `false`. Valid values:
     * 
     */
    @Export(name="autoPay", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoPay;

    /**
     * @return Specifies whether to enable automatic payment. Default value: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `auto_renew` is set to `true`.
     * 
     */
    @Export(name="autoRenewDuration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> autoRenewDuration;

    /**
     * @return The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `auto_renew` is set to `true`.
     * 
     */
    public Output<Optional<Integer>> autoRenewDuration() {
        return Codegen.optional(this.autoRenewDuration);
    }
    /**
     * Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `auto_pay` is set to `true`.
     * 
     */
    @Export(name="autoUseCoupon", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoUseCoupon;

    /**
     * @return Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `auto_pay` is set to `true`.
     * 
     */
    public Output<Optional<String>> autoUseCoupon() {
        return Codegen.optional(this.autoUseCoupon);
    }
    /**
     * The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
     * 
     */
    @Export(name="bandwidthBillingType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bandwidthBillingType;

    /**
     * @return The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
     * 
     */
    public Output<Optional<String>> bandwidthBillingType() {
        return Codegen.optional(this.bandwidthBillingType);
    }
    /**
     * The name of the Global Accelerator Basic Accelerator instance.
     * 
     */
    @Export(name="basicAcceleratorName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> basicAcceleratorName;

    /**
     * @return The name of the Global Accelerator Basic Accelerator instance.
     * 
     */
    public Output<Optional<String>> basicAcceleratorName() {
        return Codegen.optional(this.basicAcceleratorName);
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
     * The description of the Global Accelerator Basic Accelerator instance.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Global Accelerator Basic Accelerator instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The subscription duration. Default value: `1`.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
     * 
     */
    @Export(name="duration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> duration;

    /**
     * @return The subscription duration. Default value: `1`.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
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
     * The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
     * 
     */
    @Export(name="pricingCycle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pricingCycle;

    /**
     * @return The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
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
     * The status of the Basic Accelerator instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Basic Accelerator instance.
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
    public BasicAccelerator(String name) {
        this(name, BasicAcceleratorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BasicAccelerator(String name, @Nullable BasicAcceleratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BasicAccelerator(String name, @Nullable BasicAcceleratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/basicAccelerator:BasicAccelerator", name, args == null ? BasicAcceleratorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BasicAccelerator(String name, Output<String> id, @Nullable BasicAcceleratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/basicAccelerator:BasicAccelerator", name, state, makeResourceOptions(options, id));
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
    public static BasicAccelerator get(String name, Output<String> id, @Nullable BasicAcceleratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BasicAccelerator(name, id, state, options);
    }
}

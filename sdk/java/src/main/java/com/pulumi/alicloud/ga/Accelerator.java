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
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Global Accelerator (GA) Accelerator resource.
 * 
 * For information about Global Accelerator (GA) Accelerator and how to use it, see [What is Accelerator](https://help.aliyun.com/document_detail/153235.html).
 * 
 * &gt; **NOTE:** At present, The `alicloud.ga.Accelerator` cannot be deleted. you need to wait until the resource is outdated and released automatically.
 * 
 * &gt; **NOTE:** Available in v1.111.0+.
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
 *         var example = new Accelerator(&#34;example&#34;, AcceleratorArgs.builder()        
 *             .autoUseCoupon(true)
 *             .duration(1)
 *             .spec(&#34;1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ga Accelerator can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ga/accelerator:Accelerator example &lt;accelerator_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/accelerator:Accelerator")
public class Accelerator extends com.pulumi.resources.CustomResource {
    /**
     * The Name of the GA instance.
     * 
     */
    @Export(name="acceleratorName", type=String.class, parameters={})
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
    @Export(name="autoRenewDuration", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> autoRenewDuration;

    /**
     * @return Auto renewal period of an instance, in the unit of month. The value range is 1-12.
     * 
     */
    public Output<Optional<Integer>> autoRenewDuration() {
        return Codegen.optional(this.autoRenewDuration);
    }
    /**
     * Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
     * 
     */
    @Export(name="autoUseCoupon", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoUseCoupon;

    /**
     * @return Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
     * 
     */
    public Output<Optional<Boolean>> autoUseCoupon() {
        return Codegen.optional(this.autoUseCoupon);
    }
    /**
     * Descriptive information of the global acceleration instance.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Descriptive information of the global acceleration instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricing_cycle` are both required.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
     * 
     */
    @Export(name="duration", type=Integer.class, parameters={})
    private Output<Integer> duration;

    /**
     * @return The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricing_cycle` are both required.
     * * If the `pricing_cycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
     * * If the `pricing_cycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
     * 
     */
    public Output<Integer> duration() {
        return this.duration;
    }
    /**
     * The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
     * 
     */
    @Export(name="pricingCycle", type=String.class, parameters={})
    private Output<String> pricingCycle;

    /**
     * @return The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
     * 
     */
    public Output<String> pricingCycle() {
        return this.pricingCycle;
    }
    /**
     * Whether to renew an accelerator automatically or not. Default to &#34;Normal&#34;. Valid values:
     * 
     */
    @Export(name="renewalStatus", type=String.class, parameters={})
    private Output<String> renewalStatus;

    /**
     * @return Whether to renew an accelerator automatically or not. Default to &#34;Normal&#34;. Valid values:
     * 
     */
    public Output<String> renewalStatus() {
        return this.renewalStatus;
    }
    /**
     * The instance type of the GA instance. Specification of global acceleration instance, value:
     * `1`: Small 1.
     * `2`: Small 2.
     * `3`: Small 3.
     * `5`: Medium 1.
     * `8`: Medium 2.
     * `10`: Medium 3.
     * 
     */
    @Export(name="spec", type=String.class, parameters={})
    private Output<String> spec;

    /**
     * @return The instance type of the GA instance. Specification of global acceleration instance, value:
     * `1`: Small 1.
     * `2`: Small 2.
     * `3`: Small 3.
     * `5`: Medium 1.
     * `8`: Medium 2.
     * `10`: Medium 3.
     * 
     */
    public Output<String> spec() {
        return this.spec;
    }
    /**
     * The status of the GA instance.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the GA instance.
     * 
     */
    public Output<String> status() {
        return this.status;
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
    public Accelerator(String name, AcceleratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Accelerator(String name, AcceleratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
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

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.marketplace;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.marketplace.OrderArgs;
import com.pulumi.alicloud.marketplace.inputs.OrderState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.marketplace.Order;
 * import com.pulumi.alicloud.marketplace.OrderArgs;
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
 *         var order = new Order(&#34;order&#34;, OrderArgs.builder()        
 *             .couponId(&#34;&#34;)
 *             .duration(1)
 *             .packageVersion(&#34;yuncode2713600001&#34;)
 *             .payType(&#34;prepay&#34;)
 *             .pricingCycle(&#34;Month&#34;)
 *             .productCode(&#34;cmapi033136&#34;)
 *             .quantity(1)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Market order can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:marketplace/order:Order order your-order-id
 * ```
 * 
 */
@ResourceType(type="alicloud:marketplace/order:Order")
public class Order extends com.pulumi.resources.CustomResource {
    /**
     * Service providers customize additional components.
     * 
     */
    @Export(name="components", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> components;

    /**
     * @return Service providers customize additional components.
     * 
     */
    public Output<Optional<Map<String,Object>>> components() {
        return Codegen.optional(this.components);
    }
    /**
     * The coupon id of the market product.
     * 
     */
    @Export(name="couponId", type=String.class, parameters={})
    private Output</* @Nullable */ String> couponId;

    /**
     * @return The coupon id of the market product.
     * 
     */
    public Output<Optional<String>> couponId() {
        return Codegen.optional(this.couponId);
    }
    /**
     * The number of purchase cycles.
     * 
     */
    @Export(name="duration", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> duration;

    /**
     * @return The number of purchase cycles.
     * 
     */
    public Output<Optional<Integer>> duration() {
        return Codegen.optional(this.duration);
    }
    /**
     * The package version of the market product.
     * 
     */
    @Export(name="packageVersion", type=String.class, parameters={})
    private Output<String> packageVersion;

    /**
     * @return The package version of the market product.
     * 
     */
    public Output<String> packageVersion() {
        return this.packageVersion;
    }
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     * 
     */
    @Export(name="payType", type=String.class, parameters={})
    private Output</* @Nullable */ String> payType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     * 
     */
    public Output<Optional<String>> payType() {
        return Codegen.optional(this.payType);
    }
    /**
     * The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     * 
     */
    @Export(name="pricingCycle", type=String.class, parameters={})
    private Output<String> pricingCycle;

    /**
     * @return The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     * 
     */
    public Output<String> pricingCycle() {
        return this.pricingCycle;
    }
    /**
     * The product_code of market place product.
     * 
     */
    @Export(name="productCode", type=String.class, parameters={})
    private Output<String> productCode;

    /**
     * @return The product_code of market place product.
     * 
     */
    public Output<String> productCode() {
        return this.productCode;
    }
    /**
     * The quantity of the market product will be purchased.
     * 
     */
    @Export(name="quantity", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> quantity;

    /**
     * @return The quantity of the market product will be purchased.
     * 
     */
    public Output<Optional<Integer>> quantity() {
        return Codegen.optional(this.quantity);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Order(String name) {
        this(name, OrderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Order(String name, OrderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Order(String name, OrderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:marketplace/order:Order", name, args == null ? OrderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Order(String name, Output<String> id, @Nullable OrderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:marketplace/order:Order", name, state, makeResourceOptions(options, id));
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
    public static Order get(String name, Output<String> id, @Nullable OrderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Order(name, id, state, options);
    }
}

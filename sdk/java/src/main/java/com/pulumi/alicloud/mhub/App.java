// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mhub;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mhub.AppArgs;
import com.pulumi.alicloud.mhub.inputs.AppState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MHUB App resource.
 * 
 * For information about MHUB App and how to use it, see [What is App](https://help.aliyun.com/product/65109.html).
 * 
 * &gt; **NOTE:** Available since v1.138.0+.
 * 
 * &gt; **NOTE:** At present, the resource only supports cn-shanghai region.
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
 * import com.pulumi.alicloud.mhub.Product;
 * import com.pulumi.alicloud.mhub.ProductArgs;
 * import com.pulumi.alicloud.mhub.App;
 * import com.pulumi.alicloud.mhub.AppArgs;
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
 *         final var name = config.get("name").orElse("example_value");
 *         var default_ = new Product("default", ProductArgs.builder()
 *             .productName(name)
 *             .build());
 * 
 *         var defaultApp = new App("defaultApp", AppArgs.builder()
 *             .appName(name)
 *             .productId(default_.id())
 *             .packageName("com.example.android")
 *             .type("Android")
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
 * MHUB App can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mhub/app:App example &lt;product_id&gt;:&lt;app_key&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:mhub/app:App")
public class App extends com.pulumi.resources.CustomResource {
    /**
     * AppName.
     * 
     */
    @Export(name="appName", refs={String.class}, tree="[0]")
    private Output<String> appName;

    /**
     * @return AppName.
     * 
     */
    public Output<String> appName() {
        return this.appName;
    }
    /**
     * The app id of iOS. **NOTE:** Either `bundle_id` or `package_name` must be set.
     * 
     */
    @Export(name="bundleId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bundleId;

    /**
     * @return The app id of iOS. **NOTE:** Either `bundle_id` or `package_name` must be set.
     * 
     */
    public Output<Optional<String>> bundleId() {
        return Codegen.optional(this.bundleId);
    }
    /**
     * Base64 string of picture.
     * 
     */
    @Export(name="encodedIcon", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encodedIcon;

    /**
     * @return Base64 string of picture.
     * 
     */
    public Output<Optional<String>> encodedIcon() {
        return Codegen.optional(this.encodedIcon);
    }
    /**
     * The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
     * 
     */
    @Export(name="industryId", refs={String.class}, tree="[0]")
    private Output<String> industryId;

    /**
     * @return The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
     * 
     */
    public Output<String> industryId() {
        return this.industryId;
    }
    /**
     * Android App package name. **NOTE:** Either `bundle_id` or `package_name` must be set.
     * 
     */
    @Export(name="packageName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> packageName;

    /**
     * @return Android App package name. **NOTE:** Either `bundle_id` or `package_name` must be set.
     * 
     */
    public Output<Optional<String>> packageName() {
        return Codegen.optional(this.packageName);
    }
    /**
     * The ID of the Product.
     * 
     */
    @Export(name="productId", refs={String.class}, tree="[0]")
    private Output<String> productId;

    /**
     * @return The ID of the Product.
     * 
     */
    public Output<String> productId() {
        return this.productId;
    }
    /**
     * The type of the Product. Valid values: `Android` and `iOS`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the Product. Valid values: `Android` and `iOS`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public App(java.lang.String name) {
        this(name, AppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public App(java.lang.String name, AppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public App(java.lang.String name, AppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mhub/app:App", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private App(java.lang.String name, Output<java.lang.String> id, @Nullable AppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mhub/app:App", name, state, makeResourceOptions(options, id), false);
    }

    private static AppArgs makeArgs(AppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AppArgs.Empty : args;
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
    public static App get(java.lang.String name, Output<java.lang.String> id, @Nullable AppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new App(name, id, state, options);
    }
}

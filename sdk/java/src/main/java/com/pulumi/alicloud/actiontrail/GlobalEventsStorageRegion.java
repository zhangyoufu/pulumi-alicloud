// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.actiontrail.GlobalEventsStorageRegionArgs;
import com.pulumi.alicloud.actiontrail.inputs.GlobalEventsStorageRegionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Global events storage region resource.
 * 
 * For information about global events storage region and how to use it, see [What is Global Events Storage Region](https://help.aliyun.com/zh/actiontrail/developer-reference/api-actiontrail-2020-07-06-updateglobaleventsstorageregion).
 * 
 * &gt; **NOTE:** Available since v1.201.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.actiontrail.GlobalEventsStorageRegion;
 * import com.pulumi.alicloud.actiontrail.GlobalEventsStorageRegionArgs;
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
 *         var foo = new GlobalEventsStorageRegion(&#34;foo&#34;, GlobalEventsStorageRegionArgs.builder()        
 *             .storageRegion(&#34;cn-hangzhou&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Global events storage region not can be imported.
 * 
 */
@ResourceType(type="alicloud:actiontrail/globalEventsStorageRegion:GlobalEventsStorageRegion")
public class GlobalEventsStorageRegion extends com.pulumi.resources.CustomResource {
    /**
     * Global Events Storage Region.
     * 
     */
    @Export(name="storageRegion", refs={String.class}, tree="[0]")
    private Output<String> storageRegion;

    /**
     * @return Global Events Storage Region.
     * 
     */
    public Output<String> storageRegion() {
        return this.storageRegion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalEventsStorageRegion(String name) {
        this(name, GlobalEventsStorageRegionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalEventsStorageRegion(String name, @Nullable GlobalEventsStorageRegionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalEventsStorageRegion(String name, @Nullable GlobalEventsStorageRegionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/globalEventsStorageRegion:GlobalEventsStorageRegion", name, args == null ? GlobalEventsStorageRegionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalEventsStorageRegion(String name, Output<String> id, @Nullable GlobalEventsStorageRegionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/globalEventsStorageRegion:GlobalEventsStorageRegion", name, state, makeResourceOptions(options, id));
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
    public static GlobalEventsStorageRegion get(String name, Output<String> id, @Nullable GlobalEventsStorageRegionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalEventsStorageRegion(name, id, state, options);
    }
}

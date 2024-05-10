// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudstoragegateway.StorageBundleArgs;
import com.pulumi.alicloud.cloudstoragegateway.inputs.StorageBundleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Storage Gateway Storage Bundle resource.
 * 
 * For information about Cloud Storage Gateway Storage Bundle and how to use it, see [What is Storage Bundle](https://www.alibabacloud.com/help/en/cloud-storage-gateway/latest/createstoragebundle).
 * 
 * &gt; **NOTE:** Available since v1.116.0.
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
 * import com.pulumi.alicloud.cloudstoragegateway.StorageBundle;
 * import com.pulumi.alicloud.cloudstoragegateway.StorageBundleArgs;
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
 *         var example = new StorageBundle("example", StorageBundleArgs.builder()        
 *             .storageBundleName("example_value")
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
 * Cloud Storage Gateway Storage Bundle can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudstoragegateway/storageBundle:StorageBundle example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudstoragegateway/storageBundle:StorageBundle")
public class StorageBundle extends com.pulumi.resources.CustomResource {
    /**
     * The description of storage bundle.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of storage bundle.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of storage bundle.
     * 
     */
    @Export(name="storageBundleName", refs={String.class}, tree="[0]")
    private Output<String> storageBundleName;

    /**
     * @return The name of storage bundle.
     * 
     */
    public Output<String> storageBundleName() {
        return this.storageBundleName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StorageBundle(String name) {
        this(name, StorageBundleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StorageBundle(String name, StorageBundleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StorageBundle(String name, StorageBundleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/storageBundle:StorageBundle", name, args == null ? StorageBundleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private StorageBundle(String name, Output<String> id, @Nullable StorageBundleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/storageBundle:StorageBundle", name, state, makeResourceOptions(options, id));
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
    public static StorageBundle get(String name, Output<String> id, @Nullable StorageBundleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StorageBundle(name, id, state, options);
    }
}

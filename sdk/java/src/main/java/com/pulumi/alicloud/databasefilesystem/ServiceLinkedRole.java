// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.databasefilesystem;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.databasefilesystem.ServiceLinkedRoleArgs;
import com.pulumi.alicloud.databasefilesystem.inputs.ServiceLinkedRoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Using this data source can create Dbfs service-linked roles(SLR). Dbfs may need to access another Alibaba Cloud service to implement a specific feature. In this case, Dbfs must assume a specific service-linked role, which is a Resource Access Management (RAM) role, to obtain permissions to access another Alibaba Cloud service.
 * 
 * For information about Dbfs service-linked roles(SLR) and how to use it, see [What is service-linked roles](https://www.alibabacloud.com/help/doc-detail/181425.htm).
 * 
 * &gt; **NOTE:** Available since v1.157.0.
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
 * import com.pulumi.alicloud.databasefilesystem.ServiceLinkedRole;
 * import com.pulumi.alicloud.databasefilesystem.ServiceLinkedRoleArgs;
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
 *         var serviceLinkedRole = new ServiceLinkedRole(&#34;serviceLinkedRole&#34;, ServiceLinkedRoleArgs.builder()        
 *             .productName(&#34;AliyunServiceRoleForDbfs&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Dbfs service-linked roles(SLR) can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole example &lt;product_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole")
public class ServiceLinkedRole extends com.pulumi.resources.CustomResource {
    /**
     * The product name for SLR. Dbfs can automatically create the following service-linked roles: `AliyunServiceRoleForDbfs`.
     * 
     */
    @Export(name="productName", refs={String.class}, tree="[0]")
    private Output<String> productName;

    /**
     * @return The product name for SLR. Dbfs can automatically create the following service-linked roles: `AliyunServiceRoleForDbfs`.
     * 
     */
    public Output<String> productName() {
        return this.productName;
    }
    /**
     * The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     * 
     */
    @Export(name="status", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> status;

    /**
     * @return The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
     * 
     */
    public Output<Boolean> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceLinkedRole(String name) {
        this(name, ServiceLinkedRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceLinkedRole(String name, ServiceLinkedRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceLinkedRole(String name, ServiceLinkedRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole", name, args == null ? ServiceLinkedRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceLinkedRole(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:databasefilesystem/serviceLinkedRole:ServiceLinkedRole", name, state, makeResourceOptions(options, id));
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
    public static ServiceLinkedRole get(String name, Output<String> id, @Nullable ServiceLinkedRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceLinkedRole(name, id, state, options);
    }
}

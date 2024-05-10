// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.DelegatedAdministratorArgs;
import com.pulumi.alicloud.resourcemanager.inputs.DelegatedAdministratorState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager Delegated Administrator resource.
 * 
 * For information about Resource Manager Delegated Administrator and how to use it, see [What is Delegated Administrator](https://www.alibabacloud.com/help/en/resource-management/latest/registerdelegatedadministrator#doc-api-ResourceManager-RegisterDelegatedAdministrator).
 * 
 * &gt; **NOTE:** Available since v1.181.0.
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
 * import com.pulumi.alicloud.resourcemanager.Folder;
 * import com.pulumi.alicloud.resourcemanager.FolderArgs;
 * import com.pulumi.alicloud.resourcemanager.Account;
 * import com.pulumi.alicloud.resourcemanager.AccountArgs;
 * import com.pulumi.alicloud.resourcemanager.DelegatedAdministrator;
 * import com.pulumi.alicloud.resourcemanager.DelegatedAdministratorArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var displayName = config.get("displayName").orElse("EAccount");
 *         var default_ = new Integer("default", IntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var example = new Folder("example", FolderArgs.builder()        
 *             .folderName(String.format("%s-%s", name,default_.result()))
 *             .build());
 * 
 *         var exampleAccount = new Account("exampleAccount", AccountArgs.builder()        
 *             .displayName(String.format("%s-%s", displayName,default_.result()))
 *             .folderId(example.id())
 *             .build());
 * 
 *         var exampleDelegatedAdministrator = new DelegatedAdministrator("exampleDelegatedAdministrator", DelegatedAdministratorArgs.builder()        
 *             .accountId(exampleAccount.id())
 *             .servicePrincipal("cloudfw.aliyuncs.com")
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
 * Resource Manager Delegated Administrator can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:resourcemanager/delegatedAdministrator:DelegatedAdministrator example &lt;account_id&gt;:&lt;service_principal&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/delegatedAdministrator:DelegatedAdministrator")
public class DelegatedAdministrator extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the member account in the resource directory.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The ID of the member account in the resource directory.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The identification of the trusted service. **NOTE:** Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://www.alibabacloud.com/help/en/resource-management/latest/manage-trusted-services-overview).
     * 
     */
    @Export(name="servicePrincipal", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipal;

    /**
     * @return The identification of the trusted service. **NOTE:** Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://www.alibabacloud.com/help/en/resource-management/latest/manage-trusted-services-overview).
     * 
     */
    public Output<String> servicePrincipal() {
        return this.servicePrincipal;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DelegatedAdministrator(String name) {
        this(name, DelegatedAdministratorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DelegatedAdministrator(String name, DelegatedAdministratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DelegatedAdministrator(String name, DelegatedAdministratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/delegatedAdministrator:DelegatedAdministrator", name, args == null ? DelegatedAdministratorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DelegatedAdministrator(String name, Output<String> id, @Nullable DelegatedAdministratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/delegatedAdministrator:DelegatedAdministrator", name, state, makeResourceOptions(options, id));
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
    public static DelegatedAdministrator get(String name, Output<String> id, @Nullable DelegatedAdministratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DelegatedAdministrator(name, id, state, options);
    }
}

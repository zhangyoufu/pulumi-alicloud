// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudfirewall.AddressBookArgs;
import com.pulumi.alicloud.cloudfirewall.inputs.AddressBookState;
import com.pulumi.alicloud.cloudfirewall.outputs.AddressBookEcsTag;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Firewall Address Book resource.
 * 
 * For information about Cloud Firewall Address Book and how to use it, see [What is Address Book](https://www.alibabacloud.com/help/en/cloud-firewall/developer-reference/api-cloudfw-2017-12-07-addaddressbook).
 * 
 * &gt; **NOTE:** Available since v1.178.0.
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
 * import com.pulumi.alicloud.cloudfirewall.AddressBook;
 * import com.pulumi.alicloud.cloudfirewall.AddressBookArgs;
 * import com.pulumi.alicloud.cloudfirewall.inputs.AddressBookEcsTagArgs;
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
 *         var example = new AddressBook("example", AddressBookArgs.builder()        
 *             .description("example_value")
 *             .groupName("example_value")
 *             .groupType("tag")
 *             .tagRelation("and")
 *             .autoAddTagEcs(0)
 *             .ecsTags(AddressBookEcsTagArgs.builder()
 *                 .tagKey("created")
 *                 .tagValue("tfTestAcc0")
 *                 .build())
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
 * Cloud Firewall Address Book can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudfirewall/addressBook:AddressBook example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudfirewall/addressBook:AddressBook")
public class AddressBook extends com.pulumi.resources.CustomResource {
    /**
     * The list of addresses.
     * 
     */
    @Export(name="addressLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> addressLists;

    /**
     * @return The list of addresses.
     * 
     */
    public Output<Optional<List<String>>> addressLists() {
        return Codegen.optional(this.addressLists);
    }
    /**
     * Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     * 
     */
    @Export(name="autoAddTagEcs", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> autoAddTagEcs;

    /**
     * @return Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     * 
     */
    public Output<Optional<Integer>> autoAddTagEcs() {
        return Codegen.optional(this.autoAddTagEcs);
    }
    /**
     * The description of the Address Book.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the Address Book.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * A list of ECS tags. See `ecs_tags` below.
     * 
     */
    @Export(name="ecsTags", refs={List.class,AddressBookEcsTag.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AddressBookEcsTag>> ecsTags;

    /**
     * @return A list of ECS tags. See `ecs_tags` below.
     * 
     */
    public Output<Optional<List<AddressBookEcsTag>>> ecsTags() {
        return Codegen.optional(this.ecsTags);
    }
    /**
     * The name of the Address Book.
     * 
     */
    @Export(name="groupName", refs={String.class}, tree="[0]")
    private Output<String> groupName;

    /**
     * @return The name of the Address Book.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }
    /**
     * The type of the Address Book. Valid values: `ip`, `ipv6`, `domain`, `port`, `tag`.
     * **NOTE:** From version 1.213.1, `group_type` can be set to `ipv6`, `domain`, `port`.
     * 
     */
    @Export(name="groupType", refs={String.class}, tree="[0]")
    private Output<String> groupType;

    /**
     * @return The type of the Address Book. Valid values: `ip`, `ipv6`, `domain`, `port`, `tag`.
     * **NOTE:** From version 1.213.1, `group_type` can be set to `ipv6`, `domain`, `port`.
     * 
     */
    public Output<String> groupType() {
        return this.groupType;
    }
    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
     * 
     */
    @Export(name="tagRelation", refs={String.class}, tree="[0]")
    private Output<String> tagRelation;

    /**
     * @return The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
     * 
     */
    public Output<String> tagRelation() {
        return this.tagRelation;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AddressBook(String name) {
        this(name, AddressBookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AddressBook(String name, AddressBookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AddressBook(String name, AddressBookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/addressBook:AddressBook", name, args == null ? AddressBookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AddressBook(String name, Output<String> id, @Nullable AddressBookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/addressBook:AddressBook", name, state, makeResourceOptions(options, id));
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
    public static AddressBook get(String name, Output<String> id, @Nullable AddressBookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AddressBook(name, id, state, options);
    }
}

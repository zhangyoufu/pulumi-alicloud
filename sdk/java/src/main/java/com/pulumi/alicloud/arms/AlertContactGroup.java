// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.AlertContactGroupArgs;
import com.pulumi.alicloud.arms.inputs.AlertContactGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact Group resource.
 * 
 * For information about Application Real-Time Monitoring Service (ARMS) Alert Contact Group and how to use it, see [What is Alert Contact Group](https://www.alibabacloud.com/help/en/doc-detail/130677.htm).
 * 
 * &gt; **NOTE:** Available since v1.131.0.
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
 * import com.pulumi.alicloud.arms.AlertContact;
 * import com.pulumi.alicloud.arms.AlertContactArgs;
 * import com.pulumi.alicloud.arms.AlertContactGroup;
 * import com.pulumi.alicloud.arms.AlertContactGroupArgs;
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
 *         var exampleAlertContact = new AlertContact(&#34;exampleAlertContact&#34;, AlertContactArgs.builder()        
 *             .alertContactName(&#34;example_value&#34;)
 *             .dingRobotWebhookUrl(&#34;https://oapi.dingtalk.com/robot/send?access_token=91f2f6****&#34;)
 *             .email(&#34;someone@example.com&#34;)
 *             .phoneNum(&#34;1381111****&#34;)
 *             .build());
 * 
 *         var exampleAlertContactGroup = new AlertContactGroup(&#34;exampleAlertContactGroup&#34;, AlertContactGroupArgs.builder()        
 *             .alertContactGroupName(&#34;example_value&#34;)
 *             .contactIds(exampleAlertContact.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Application Real-Time Monitoring Service (ARMS) Alert Contact Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/alertContactGroup:AlertContactGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/alertContactGroup:AlertContactGroup")
public class AlertContactGroup extends com.pulumi.resources.CustomResource {
    /**
     * The name of the resource.
     * 
     */
    @Export(name="alertContactGroupName", refs={String.class}, tree="[0]")
    private Output<String> alertContactGroupName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> alertContactGroupName() {
        return this.alertContactGroupName;
    }
    /**
     * The list id of alert contact.
     * 
     */
    @Export(name="contactIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> contactIds;

    /**
     * @return The list id of alert contact.
     * 
     */
    public Output<Optional<List<String>>> contactIds() {
        return Codegen.optional(this.contactIds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlertContactGroup(String name) {
        this(name, AlertContactGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlertContactGroup(String name, AlertContactGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlertContactGroup(String name, AlertContactGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/alertContactGroup:AlertContactGroup", name, args == null ? AlertContactGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlertContactGroup(String name, Output<String> id, @Nullable AlertContactGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/alertContactGroup:AlertContactGroup", name, state, makeResourceOptions(options, id));
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
    public static AlertContactGroup get(String name, Output<String> id, @Nullable AlertContactGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlertContactGroup(name, id, state, options);
    }
}

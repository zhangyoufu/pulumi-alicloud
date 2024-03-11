// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.AlertContactArgs;
import com.pulumi.alicloud.arms.inputs.AlertContactState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Alert Contact resource.
 * 
 * For information about Application Real-Time Monitoring Service (ARMS) Alert Contact and how to use it, see [What is Alert Contact](https://www.alibabacloud.com/help/en/application-real-time-monitoring-service/latest/createalertcontact).
 * 
 * &gt; **NOTE:** Available since v1.129.0.
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
 *         var example = new AlertContact(&#34;example&#34;, AlertContactArgs.builder()        
 *             .alertContactName(&#34;example_value&#34;)
 *             .dingRobotWebhookUrl(&#34;https://oapi.dingtalk.com/robot/send?access_token=91f2f6****&#34;)
 *             .email(&#34;someone@example.com&#34;)
 *             .phoneNum(&#34;1381111****&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/alertContact:AlertContact example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/alertContact:AlertContact")
public class AlertContact extends com.pulumi.resources.CustomResource {
    /**
     * The name of the alert contact.
     * 
     */
    @Export(name="alertContactName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> alertContactName;

    /**
     * @return The name of the alert contact.
     * 
     */
    public Output<Optional<String>> alertContactName() {
        return Codegen.optional(this.alertContactName);
    }
    /**
     * The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    @Export(name="dingRobotWebhookUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dingRobotWebhookUrl;

    /**
     * @return The webhook URL of the DingTalk chatbot. For more information about how to obtain the URL, see Configure a DingTalk chatbot to send alert notifications: https://www.alibabacloud.com/help/en/doc-detail/106247.htm. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    public Output<Optional<String>> dingRobotWebhookUrl() {
        return Codegen.optional(this.dingRobotWebhookUrl);
    }
    /**
     * The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> email;

    /**
     * @return The email address of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    public Output<Optional<String>> email() {
        return Codegen.optional(this.email);
    }
    /**
     * The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    @Export(name="phoneNum", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> phoneNum;

    /**
     * @return The mobile number of the alert contact. You must specify at least one of the following parameters: PhoneNum, Email, and DingRobotWebhookUrl.
     * 
     */
    public Output<Optional<String>> phoneNum() {
        return Codegen.optional(this.phoneNum);
    }
    /**
     * Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
     * 
     */
    @Export(name="systemNoc", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> systemNoc;

    /**
     * @return Specifies whether the alert contact receives system notifications. Valid values:  true: receives system notifications. false: does not receive system notifications.
     * 
     */
    public Output<Optional<Boolean>> systemNoc() {
        return Codegen.optional(this.systemNoc);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlertContact(String name) {
        this(name, AlertContactArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlertContact(String name, @Nullable AlertContactArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlertContact(String name, @Nullable AlertContactArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/alertContact:AlertContact", name, args == null ? AlertContactArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlertContact(String name, Output<String> id, @Nullable AlertContactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/alertContact:AlertContact", name, state, makeResourceOptions(options, id));
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
    public static AlertContact get(String name, Output<String> id, @Nullable AlertContactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlertContact(name, id, state, options);
    }
}

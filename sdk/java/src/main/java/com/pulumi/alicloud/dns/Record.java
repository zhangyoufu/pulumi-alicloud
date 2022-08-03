// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.RecordArgs;
import com.pulumi.alicloud.dns.inputs.RecordState;
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
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.dns.Record;
 * import com.pulumi.alicloud.dns.RecordArgs;
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
 *         var record = new Record(&#34;record&#34;, RecordArgs.builder()        
 *             .hostRecord(&#34;@&#34;)
 *             .type(&#34;A&#34;)
 *             .value(&#34;192.168.99.99&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDS record can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dns/record:Record example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/record:Record")
public class Record extends com.pulumi.resources.CustomResource {
    /**
     * Host record for the domain record. This host_record can have at most 253 characters, and each part split with &#34;.&#34; can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;*&#34;,&#34;@&#34;,  and must not begin or end with &#34;-&#34;.
     * 
     */
    @Export(name="hostRecord", type=String.class, parameters={})
    private Output<String> hostRecord;

    /**
     * @return Host record for the domain record. This host_record can have at most 253 characters, and each part split with &#34;.&#34; can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;*&#34;,&#34;@&#34;,  and must not begin or end with &#34;-&#34;.
     * 
     */
    public Output<String> hostRecord() {
        return this.hostRecord;
    }
    @Export(name="locked", type=Boolean.class, parameters={})
    private Output<Boolean> locked;

    public Output<Boolean> locked() {
        return this.locked;
    }
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The priority of domain record. Valid values are `[1-10]`. When the `type` is `MX`, this parameter is required.
     * 
     */
    @Export(name="priority", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return The priority of domain record. Valid values are `[1-10]`. When the `type` is `MX`, this parameter is required.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The resolution line of domain record. Valid values are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. When the `type` is `FORWORD_URL`, this parameter must be `default`. Default value is `default`. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud.dns.getResolutionLines in data source to get the value.
     * 
     */
    @Export(name="routing", type=String.class, parameters={})
    private Output</* @Nullable */ String> routing;

    /**
     * @return The resolution line of domain record. Valid values are `default`, `telecom`, `unicom`, `mobile`, `oversea`, `edu`, `drpeng`, `btvn`, .etc. When the `type` is `FORWORD_URL`, this parameter must be `default`. Default value is `default`. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud.dns.getResolutionLines in data source to get the value.
     * 
     */
    public Output<Optional<String>> routing() {
        return Codegen.optional(this.routing);
    }
    /**
     * The record status. `Enable` or `Disable`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The record status. `Enable` or `Disable`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is `[600, 86400]`, Basic is `[120, 86400]`, Standard is `[60, 86400]`, Ultimate is `[10, 86400]`, Exclusive is `[1, 86400]`. Default value is `600`.
     * 
     */
    @Export(name="ttl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is `[600, 86400]`, Basic is `[120, 86400]`, Standard is `[60, 86400]`, Ultimate is `[10, 86400]`, Exclusive is `[1, 86400]`. Default value is `600`.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * The type of domain record. Valid values are `A`,`NS`,`MX`,`TXT`,`CNAME`,`SRV`,`AAAA`,`CAA`, `REDIRECT_URL` and `FORWORD_URL`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of domain record. Valid values are `A`,`NS`,`MX`,`TXT`,`CNAME`,`SRV`,`AAAA`,`CAA`, `REDIRECT_URL` and `FORWORD_URL`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The value of domain record, When the `type` is `MX`,`NS`,`CNAME`,`SRV`, the server will treat the `value` as a fully qualified domain name, so it&#39;s no need to add a `.` at the end.
     * 
     */
    @Export(name="value", type=String.class, parameters={})
    private Output<String> value;

    /**
     * @return The value of domain record, When the `type` is `MX`,`NS`,`CNAME`,`SRV`, the server will treat the `value` as a fully qualified domain name, so it&#39;s no need to add a `.` at the end.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Record(String name) {
        this(name, RecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Record(String name, RecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Record(String name, RecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/record:Record", name, args == null ? RecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Record(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/record:Record", name, state, makeResourceOptions(options, id));
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
    public static Record get(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Record(name, id, state, options);
    }
}

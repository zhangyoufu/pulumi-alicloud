// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleRuleActionRewriteConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleRuleActionRewriteConfigArgs Empty = new RuleRuleActionRewriteConfigArgs();

    /**
     * The host name of the destination to which requests are redirected within ALB. Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The host name of the destination to which requests are redirected within ALB. Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * The path to which requests are to be redirected within ALB. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; %!;(MISSING) ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. This value can be used only once. You can use it with a valid string.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to which requests are to be redirected within ALB. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; %!;(MISSING) ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. This value can be used only once. You can use it with a valid string.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The query string of the request to be redirected within ALB. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. This value can be used only once. You can use it with a valid string.
     * 
     */
    @Import(name="query")
    private @Nullable Output<String> query;

    /**
     * @return The query string of the request to be redirected within ALB. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. This value can be used only once. You can use it with a valid string.
     * 
     */
    public Optional<Output<String>> query() {
        return Optional.ofNullable(this.query);
    }

    private RuleRuleActionRewriteConfigArgs() {}

    private RuleRuleActionRewriteConfigArgs(RuleRuleActionRewriteConfigArgs $) {
        this.host = $.host;
        this.path = $.path;
        this.query = $.query;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleRuleActionRewriteConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleRuleActionRewriteConfigArgs $;

        public Builder() {
            $ = new RuleRuleActionRewriteConfigArgs();
        }

        public Builder(RuleRuleActionRewriteConfigArgs defaults) {
            $ = new RuleRuleActionRewriteConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param host The host name of the destination to which requests are redirected within ALB. Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host name of the destination to which requests are redirected within ALB. Valid values:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param path The path to which requests are to be redirected within ALB. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; %!;(MISSING) ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to which requests are to be redirected within ALB. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?)and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; %!;(MISSING) ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param query The query string of the request to be redirected within ALB. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder query(@Nullable Output<String> query) {
            $.query = query;
            return this;
        }

        /**
         * @param query The query string of the request to be redirected within ALB. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. This value can be used only once. You can use it with a valid string.
         * 
         * @return builder
         * 
         */
        public Builder query(String query) {
            return query(Output.of(query));
        }

        public RuleRuleActionRewriteConfigArgs build() {
            return $;
        }
    }

}

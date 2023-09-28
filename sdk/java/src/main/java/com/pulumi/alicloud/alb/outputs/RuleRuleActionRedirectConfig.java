// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuleRuleActionRedirectConfig {
    /**
     * @return The host name of the destination to which requests are directed. The host name must meet the following rules:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    private @Nullable String host;
    /**
     * @return The HTTP status code of the response. The code must be an `HTTP_2xx`, `HTTP_4xx` or `HTTP_5xx.x` is a digit.
     * 
     */
    private @Nullable String httpCode;
    /**
     * @return The path of the destination to which requests are directed. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?) and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them with a valid string.
     * 
     */
    private @Nullable String path;
    /**
     * @return The port of the destination to which requests are redirected. Valid values: 1 to 63335. Default value: ${port}. You cannot use this value together with other characters at the same time.
     * 
     */
    private @Nullable String port;
    /**
     * @return The protocol of the requests to be redirected. Valid values: `HTTP` and `HTTPS`. Default value: `${protocol}`. You cannot use this value together with other characters at the same time. Note HTTPS listeners can redirect only HTTPS requests.
     * 
     */
    private @Nullable String protocol;
    /**
     * @return The query string of the request to be redirected. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them together with a valid string.
     * 
     */
    private @Nullable String query;

    private RuleRuleActionRedirectConfig() {}
    /**
     * @return The host name of the destination to which requests are directed. The host name must meet the following rules:  The host name must be 3 to 128 characters in length, and can contain letters, digits, hyphens (-), periods (.), asterisks (*), and question marks (?). The host name must contain at least one period (.), and cannot start or end with a period (.). The rightmost domain label can contain only letters, asterisks (*) and question marks (?) and cannot contain digits or hyphens (-). Other domain labels cannot start or end with a hyphen (-). You can include asterisks (*) and question marks (?) anywhere in a domain label. Default value: ${host}. You cannot use this value with other characters at the same time.
     * 
     */
    public Optional<String> host() {
        return Optional.ofNullable(this.host);
    }
    /**
     * @return The HTTP status code of the response. The code must be an `HTTP_2xx`, `HTTP_4xx` or `HTTP_5xx.x` is a digit.
     * 
     */
    public Optional<String> httpCode() {
        return Optional.ofNullable(this.httpCode);
    }
    /**
     * @return The path of the destination to which requests are directed. Valid values: The path must be 1 to 128 characters in length, and start with a forward slash (/). The path can contain letters, digits, asterisks (*), question marks (?) and the following special characters: $ - _ . + / &amp; ~ @ :. It cannot contain the following special characters: &#34; % # ; ! ( ) [ ] ^ , ”. The path is case-sensitive. Default value: ${path}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them with a valid string.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    /**
     * @return The port of the destination to which requests are redirected. Valid values: 1 to 63335. Default value: ${port}. You cannot use this value together with other characters at the same time.
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The protocol of the requests to be redirected. Valid values: `HTTP` and `HTTPS`. Default value: `${protocol}`. You cannot use this value together with other characters at the same time. Note HTTPS listeners can redirect only HTTPS requests.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }
    /**
     * @return The query string of the request to be redirected. The query string must be 1 to 128 characters in length, can contain letters and printable characters. It cannot contain the following special characters: # [ ] { } \ | &lt; &gt; &amp;. Default value: ${query}. You can also reference ${host}, ${protocol}, and ${port}. Each variable can appear at most once. You can use the preceding variables at the same time, or use them together with a valid string.
     * 
     */
    public Optional<String> query() {
        return Optional.ofNullable(this.query);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleRuleActionRedirectConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String host;
        private @Nullable String httpCode;
        private @Nullable String path;
        private @Nullable String port;
        private @Nullable String protocol;
        private @Nullable String query;
        public Builder() {}
        public Builder(RuleRuleActionRedirectConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.host = defaults.host;
    	      this.httpCode = defaults.httpCode;
    	      this.path = defaults.path;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
    	      this.query = defaults.query;
        }

        @CustomType.Setter
        public Builder host(@Nullable String host) {
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder httpCode(@Nullable String httpCode) {
            this.httpCode = httpCode;
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable String port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder query(@Nullable String query) {
            this.query = query;
            return this;
        }
        public RuleRuleActionRedirectConfig build() {
            final var o = new RuleRuleActionRedirectConfig();
            o.host = host;
            o.httpCode = httpCode;
            o.path = path;
            o.port = port;
            o.protocol = protocol;
            o.query = query;
            return o;
        }
    }
}

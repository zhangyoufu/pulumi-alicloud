// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSystemSecurityPoliciesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSystemSecurityPoliciesPlainArgs Empty = new GetSystemSecurityPoliciesPlainArgs();

    /**
     * A list of System Security Policy IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of System Security Policy IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetSystemSecurityPoliciesPlainArgs() {}

    private GetSystemSecurityPoliciesPlainArgs(GetSystemSecurityPoliciesPlainArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSystemSecurityPoliciesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSystemSecurityPoliciesPlainArgs $;

        public Builder() {
            $ = new GetSystemSecurityPoliciesPlainArgs();
        }

        public Builder(GetSystemSecurityPoliciesPlainArgs defaults) {
            $ = new GetSystemSecurityPoliciesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of System Security Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of System Security Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        public GetSystemSecurityPoliciesPlainArgs build() {
            return $;
        }
    }

}

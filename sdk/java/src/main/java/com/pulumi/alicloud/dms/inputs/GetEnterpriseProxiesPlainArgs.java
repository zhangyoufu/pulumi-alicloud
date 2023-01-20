// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEnterpriseProxiesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnterpriseProxiesPlainArgs Empty = new GetEnterpriseProxiesPlainArgs();

    /**
     * A list of Proxy IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Proxy IDs.
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

    /**
     * The ID of the tenant.
     * 
     */
    @Import(name="tid")
    private @Nullable String tid;

    /**
     * @return The ID of the tenant.
     * 
     */
    public Optional<String> tid() {
        return Optional.ofNullable(this.tid);
    }

    private GetEnterpriseProxiesPlainArgs() {}

    private GetEnterpriseProxiesPlainArgs(GetEnterpriseProxiesPlainArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.tid = $.tid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnterpriseProxiesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnterpriseProxiesPlainArgs $;

        public Builder() {
            $ = new GetEnterpriseProxiesPlainArgs();
        }

        public Builder(GetEnterpriseProxiesPlainArgs defaults) {
            $ = new GetEnterpriseProxiesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Proxy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Proxy IDs.
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

        /**
         * @param tid The ID of the tenant.
         * 
         * @return builder
         * 
         */
        public Builder tid(@Nullable String tid) {
            $.tid = tid;
            return this;
        }

        public GetEnterpriseProxiesPlainArgs build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.alicloud.hbr.inputs.GetServerBackupPlansFilter;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServerBackupPlansPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServerBackupPlansPlainArgs Empty = new GetServerBackupPlansPlainArgs();

    /**
     * The filters.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetServerBackupPlansFilter> filters;

    /**
     * @return The filters.
     * 
     */
    public Optional<List<GetServerBackupPlansFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * A list of Server Backup Plan IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Server Backup Plan IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetServerBackupPlansPlainArgs() {}

    private GetServerBackupPlansPlainArgs(GetServerBackupPlansPlainArgs $) {
        this.filters = $.filters;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerBackupPlansPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerBackupPlansPlainArgs $;

        public Builder() {
            $ = new GetServerBackupPlansPlainArgs();
        }

        public Builder(GetServerBackupPlansPlainArgs defaults) {
            $ = new GetServerBackupPlansPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters The filters.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetServerBackupPlansFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters The filters.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetServerBackupPlansFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param ids A list of Server Backup Plan IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Server Backup Plan IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetServerBackupPlansPlainArgs build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFilesetsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFilesetsPlainArgs Empty = new GetFilesetsPlainArgs();

    /**
     * The ID of the file system.
     * 
     */
    @Import(name="fileSystemId", required=true)
    private String fileSystemId;

    /**
     * @return The ID of the file system.
     * 
     */
    public String fileSystemId() {
        return this.fileSystemId;
    }

    /**
     * A list of Fileset IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Fileset IDs.
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

    /**
     * The status of the fileset.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the fileset.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetFilesetsPlainArgs() {}

    private GetFilesetsPlainArgs(GetFilesetsPlainArgs $) {
        this.fileSystemId = $.fileSystemId;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFilesetsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFilesetsPlainArgs $;

        public Builder() {
            $ = new GetFilesetsPlainArgs();
        }

        public Builder(GetFilesetsPlainArgs defaults) {
            $ = new GetFilesetsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param fileSystemId The ID of the file system.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(String fileSystemId) {
            $.fileSystemId = fileSystemId;
            return this;
        }

        /**
         * @param ids A list of Fileset IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Fileset IDs.
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

        /**
         * @param status The status of the fileset.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetFilesetsPlainArgs build() {
            $.fileSystemId = Objects.requireNonNull($.fileSystemId, "expected parameter 'fileSystemId' to be non-null");
            return $;
        }
    }

}

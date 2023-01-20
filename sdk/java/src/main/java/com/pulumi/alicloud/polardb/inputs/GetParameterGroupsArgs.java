// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetParameterGroupsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetParameterGroupsArgs Empty = new GetParameterGroupsArgs();

    /**
     * The type of the database engine.
     * 
     */
    @Import(name="dbType")
    private @Nullable Output<String> dbType;

    /**
     * @return The type of the database engine.
     * 
     */
    public Optional<Output<String>> dbType() {
        return Optional.ofNullable(this.dbType);
    }

    /**
     * The version number of the database engine.
     * 
     */
    @Import(name="dbVersion")
    private @Nullable Output<String> dbVersion;

    /**
     * @return The version number of the database engine.
     * 
     */
    public Optional<Output<String>> dbVersion() {
        return Optional.ofNullable(this.dbVersion);
    }

    /**
     * A list of Parameter Group IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Parameter Group IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Parameter Group name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Parameter Group name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetParameterGroupsArgs() {}

    private GetParameterGroupsArgs(GetParameterGroupsArgs $) {
        this.dbType = $.dbType;
        this.dbVersion = $.dbVersion;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetParameterGroupsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetParameterGroupsArgs $;

        public Builder() {
            $ = new GetParameterGroupsArgs();
        }

        public Builder(GetParameterGroupsArgs defaults) {
            $ = new GetParameterGroupsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbType The type of the database engine.
         * 
         * @return builder
         * 
         */
        public Builder dbType(@Nullable Output<String> dbType) {
            $.dbType = dbType;
            return this;
        }

        /**
         * @param dbType The type of the database engine.
         * 
         * @return builder
         * 
         */
        public Builder dbType(String dbType) {
            return dbType(Output.of(dbType));
        }

        /**
         * @param dbVersion The version number of the database engine.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(@Nullable Output<String> dbVersion) {
            $.dbVersion = dbVersion;
            return this;
        }

        /**
         * @param dbVersion The version number of the database engine.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(String dbVersion) {
            return dbVersion(Output.of(dbVersion));
        }

        /**
         * @param ids A list of Parameter Group IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Parameter Group IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Parameter Group IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Parameter Group name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Parameter Group name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public GetParameterGroupsArgs build() {
            return $;
        }
    }

}

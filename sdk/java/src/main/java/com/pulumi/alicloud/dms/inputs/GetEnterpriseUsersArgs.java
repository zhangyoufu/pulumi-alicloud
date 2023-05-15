// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEnterpriseUsersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnterpriseUsersArgs Empty = new GetEnterpriseUsersArgs();

    /**
     * A list of DMS Enterprise User IDs (UID).
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of DMS Enterprise User IDs (UID).
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter the results by the DMS Enterprise User nick_name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter the results by the DMS Enterprise User nick_name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The role of the user to query.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The role of the user to query.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * The keyword used to query users.
     * 
     */
    @Import(name="searchKey")
    private @Nullable Output<String> searchKey;

    /**
     * @return The keyword used to query users.
     * 
     */
    public Optional<Output<String>> searchKey() {
        return Optional.ofNullable(this.searchKey);
    }

    /**
     * The status of the user.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the user.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the tenant in DMS Enterprise.
     * 
     */
    @Import(name="tid")
    private @Nullable Output<Integer> tid;

    /**
     * @return The ID of the tenant in DMS Enterprise.
     * 
     */
    public Optional<Output<Integer>> tid() {
        return Optional.ofNullable(this.tid);
    }

    private GetEnterpriseUsersArgs() {}

    private GetEnterpriseUsersArgs(GetEnterpriseUsersArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.role = $.role;
        this.searchKey = $.searchKey;
        this.status = $.status;
        this.tid = $.tid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnterpriseUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnterpriseUsersArgs $;

        public Builder() {
            $ = new GetEnterpriseUsersArgs();
        }

        public Builder(GetEnterpriseUsersArgs defaults) {
            $ = new GetEnterpriseUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of DMS Enterprise User IDs (UID).
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of DMS Enterprise User IDs (UID).
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of DMS Enterprise User IDs (UID).
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter the results by the DMS Enterprise User nick_name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter the results by the DMS Enterprise User nick_name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param role The role of the user to query.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role of the user to query.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param searchKey The keyword used to query users.
         * 
         * @return builder
         * 
         */
        public Builder searchKey(@Nullable Output<String> searchKey) {
            $.searchKey = searchKey;
            return this;
        }

        /**
         * @param searchKey The keyword used to query users.
         * 
         * @return builder
         * 
         */
        public Builder searchKey(String searchKey) {
            return searchKey(Output.of(searchKey));
        }

        /**
         * @param status The status of the user.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the user.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tid The ID of the tenant in DMS Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder tid(@Nullable Output<Integer> tid) {
            $.tid = tid;
            return this;
        }

        /**
         * @param tid The ID of the tenant in DMS Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder tid(Integer tid) {
            return tid(Output.of(tid));
        }

        public GetEnterpriseUsersArgs build() {
            return $;
        }
    }

}

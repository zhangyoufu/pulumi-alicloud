// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceQueuesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceQueuesPlainArgs Empty = new GetServiceQueuesPlainArgs();

    /**
     * A list of Queue IDs. Its element value is same as Queue Name.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Queue IDs. Its element value is same as Queue Name.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Queue name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Queue name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The name of the queue.
     * 
     */
    @Import(name="queueName")
    private @Nullable String queueName;

    /**
     * @return The name of the queue.
     * 
     */
    public Optional<String> queueName() {
        return Optional.ofNullable(this.queueName);
    }

    private GetServiceQueuesPlainArgs() {}

    private GetServiceQueuesPlainArgs(GetServiceQueuesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.queueName = $.queueName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceQueuesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceQueuesPlainArgs $;

        public Builder() {
            $ = new GetServiceQueuesPlainArgs();
        }

        public Builder(GetServiceQueuesPlainArgs defaults) {
            $ = new GetServiceQueuesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Queue IDs. Its element value is same as Queue Name.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Queue IDs. Its element value is same as Queue Name.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Queue name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param queueName The name of the queue.
         * 
         * @return builder
         * 
         */
        public Builder queueName(@Nullable String queueName) {
            $.queueName = queueName;
            return this;
        }

        public GetServiceQueuesPlainArgs build() {
            return $;
        }
    }

}

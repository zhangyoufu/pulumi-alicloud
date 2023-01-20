// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ShardingInstanceMongoListArgs extends com.pulumi.resources.ResourceArgs {

    public static final ShardingInstanceMongoListArgs Empty = new ShardingInstanceMongoListArgs();

    /**
     * The connection address of the Config Server node.
     * 
     */
    @Import(name="connectString")
    private @Nullable Output<String> connectString;

    /**
     * @return The connection address of the Config Server node.
     * 
     */
    public Optional<Output<String>> connectString() {
        return Optional.ofNullable(this.connectString);
    }

    /**
     * Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    @Import(name="nodeClass", required=true)
    private Output<String> nodeClass;

    /**
     * @return Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    public Output<String> nodeClass() {
        return this.nodeClass;
    }

    /**
     * The ID of the Config Server node.
     * 
     */
    @Import(name="nodeId")
    private @Nullable Output<String> nodeId;

    /**
     * @return The ID of the Config Server node.
     * 
     */
    public Optional<Output<String>> nodeId() {
        return Optional.ofNullable(this.nodeId);
    }

    /**
     * The connection port of the Config Server node.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The connection port of the Config Server node.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private ShardingInstanceMongoListArgs() {}

    private ShardingInstanceMongoListArgs(ShardingInstanceMongoListArgs $) {
        this.connectString = $.connectString;
        this.nodeClass = $.nodeClass;
        this.nodeId = $.nodeId;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShardingInstanceMongoListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShardingInstanceMongoListArgs $;

        public Builder() {
            $ = new ShardingInstanceMongoListArgs();
        }

        public Builder(ShardingInstanceMongoListArgs defaults) {
            $ = new ShardingInstanceMongoListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectString The connection address of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder connectString(@Nullable Output<String> connectString) {
            $.connectString = connectString;
            return this;
        }

        /**
         * @param connectString The connection address of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder connectString(String connectString) {
            return connectString(Output.of(connectString));
        }

        /**
         * @param nodeClass Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
         * 
         * @return builder
         * 
         */
        public Builder nodeClass(Output<String> nodeClass) {
            $.nodeClass = nodeClass;
            return this;
        }

        /**
         * @param nodeClass Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
         * 
         * @return builder
         * 
         */
        public Builder nodeClass(String nodeClass) {
            return nodeClass(Output.of(nodeClass));
        }

        /**
         * @param nodeId The ID of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder nodeId(@Nullable Output<String> nodeId) {
            $.nodeId = nodeId;
            return this;
        }

        /**
         * @param nodeId The ID of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder nodeId(String nodeId) {
            return nodeId(Output.of(nodeId));
        }

        /**
         * @param port The connection port of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The connection port of the Config Server node.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public ShardingInstanceMongoListArgs build() {
            $.nodeClass = Objects.requireNonNull($.nodeClass, "expected parameter 'nodeClass' to be non-null");
            return $;
        }
    }

}

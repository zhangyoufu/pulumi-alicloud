// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AclAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final AclAttachmentState Empty = new AclAttachmentState();

    /**
     * The ID of an ACL.
     * 
     */
    @Import(name="aclId")
    private @Nullable Output<String> aclId;

    /**
     * @return The ID of an ACL.
     * 
     */
    public Optional<Output<String>> aclId() {
        return Optional.ofNullable(this.aclId);
    }

    /**
     * The type of the ACL. Valid values:
     * 
     */
    @Import(name="aclType")
    private @Nullable Output<String> aclType;

    /**
     * @return The type of the ACL. Valid values:
     * 
     */
    public Optional<Output<String>> aclType() {
        return Optional.ofNullable(this.aclType);
    }

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The ID of the listener.
     * 
     */
    @Import(name="listenerId")
    private @Nullable Output<String> listenerId;

    /**
     * @return The ID of the listener.
     * 
     */
    public Optional<Output<String>> listenerId() {
        return Optional.ofNullable(this.listenerId);
    }

    /**
     * The status of the Acl Attachment.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Acl Attachment.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private AclAttachmentState() {}

    private AclAttachmentState(AclAttachmentState $) {
        this.aclId = $.aclId;
        this.aclType = $.aclType;
        this.dryRun = $.dryRun;
        this.listenerId = $.listenerId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AclAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AclAttachmentState $;

        public Builder() {
            $ = new AclAttachmentState();
        }

        public Builder(AclAttachmentState defaults) {
            $ = new AclAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclId The ID of an ACL.
         * 
         * @return builder
         * 
         */
        public Builder aclId(@Nullable Output<String> aclId) {
            $.aclId = aclId;
            return this;
        }

        /**
         * @param aclId The ID of an ACL.
         * 
         * @return builder
         * 
         */
        public Builder aclId(String aclId) {
            return aclId(Output.of(aclId));
        }

        /**
         * @param aclType The type of the ACL. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder aclType(@Nullable Output<String> aclType) {
            $.aclType = aclType;
            return this;
        }

        /**
         * @param aclType The type of the ACL. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder aclType(String aclType) {
            return aclType(Output.of(aclType));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param listenerId The ID of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(@Nullable Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The ID of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param status The status of the Acl Attachment.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Acl Attachment.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public AclAttachmentState build() {
            return $;
        }
    }

}

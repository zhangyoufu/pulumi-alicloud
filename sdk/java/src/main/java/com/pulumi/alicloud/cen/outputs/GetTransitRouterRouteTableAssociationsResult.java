// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.alicloud.cen.outputs.GetTransitRouterRouteTableAssociationsAssociation;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTransitRouterRouteTableAssociationsResult {
    /**
     * @return A list of CEN Transit Router Route Table Associations. Each element contains the following attributes:
     * 
     */
    private List<GetTransitRouterRouteTableAssociationsAssociation> associations;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of CEN Transit Router Route Table Association IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return The status of the route table.
     * 
     */
    private @Nullable String status;
    /**
     * @return ID of the transit router attachment.
     * 
     */
    private @Nullable String transitRouterAttachmentId;
    private @Nullable String transitRouterAttachmentResourceId;
    private @Nullable String transitRouterAttachmentResourceType;
    /**
     * @return ID of the transit router route table.
     * 
     */
    private @Nullable String transitRouterRouteTableId;

    private GetTransitRouterRouteTableAssociationsResult() {}
    /**
     * @return A list of CEN Transit Router Route Table Associations. Each element contains the following attributes:
     * 
     */
    public List<GetTransitRouterRouteTableAssociationsAssociation> associations() {
        return this.associations;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of CEN Transit Router Route Table Association IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The status of the route table.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return ID of the transit router attachment.
     * 
     */
    public Optional<String> transitRouterAttachmentId() {
        return Optional.ofNullable(this.transitRouterAttachmentId);
    }
    public Optional<String> transitRouterAttachmentResourceId() {
        return Optional.ofNullable(this.transitRouterAttachmentResourceId);
    }
    public Optional<String> transitRouterAttachmentResourceType() {
        return Optional.ofNullable(this.transitRouterAttachmentResourceType);
    }
    /**
     * @return ID of the transit router route table.
     * 
     */
    public Optional<String> transitRouterRouteTableId() {
        return Optional.ofNullable(this.transitRouterRouteTableId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTransitRouterRouteTableAssociationsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetTransitRouterRouteTableAssociationsAssociation> associations;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable String transitRouterAttachmentId;
        private @Nullable String transitRouterAttachmentResourceId;
        private @Nullable String transitRouterAttachmentResourceType;
        private @Nullable String transitRouterRouteTableId;
        public Builder() {}
        public Builder(GetTransitRouterRouteTableAssociationsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.associations = defaults.associations;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.transitRouterAttachmentId = defaults.transitRouterAttachmentId;
    	      this.transitRouterAttachmentResourceId = defaults.transitRouterAttachmentResourceId;
    	      this.transitRouterAttachmentResourceType = defaults.transitRouterAttachmentResourceType;
    	      this.transitRouterRouteTableId = defaults.transitRouterRouteTableId;
        }

        @CustomType.Setter
        public Builder associations(List<GetTransitRouterRouteTableAssociationsAssociation> associations) {
            this.associations = Objects.requireNonNull(associations);
            return this;
        }
        public Builder associations(GetTransitRouterRouteTableAssociationsAssociation... associations) {
            return associations(List.of(associations));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterAttachmentId(@Nullable String transitRouterAttachmentId) {
            this.transitRouterAttachmentId = transitRouterAttachmentId;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterAttachmentResourceId(@Nullable String transitRouterAttachmentResourceId) {
            this.transitRouterAttachmentResourceId = transitRouterAttachmentResourceId;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterAttachmentResourceType(@Nullable String transitRouterAttachmentResourceType) {
            this.transitRouterAttachmentResourceType = transitRouterAttachmentResourceType;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterRouteTableId(@Nullable String transitRouterRouteTableId) {
            this.transitRouterRouteTableId = transitRouterRouteTableId;
            return this;
        }
        public GetTransitRouterRouteTableAssociationsResult build() {
            final var o = new GetTransitRouterRouteTableAssociationsResult();
            o.associations = associations;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.status = status;
            o.transitRouterAttachmentId = transitRouterAttachmentId;
            o.transitRouterAttachmentResourceId = transitRouterAttachmentResourceId;
            o.transitRouterAttachmentResourceType = transitRouterAttachmentResourceType;
            o.transitRouterRouteTableId = transitRouterRouteTableId;
            return o;
        }
    }
}

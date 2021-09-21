// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetListenersListenerAclConfigAclRelationResult
    {
        /// <summary>
        /// Snooping Binding of the Access Policy Group ID List.
        /// </summary>
        public readonly string AclId;
        /// <summary>
        /// The association status between the ACL and the listener.  Valid values: `Associating`, `Associated` Or `Dissociating`. `Associating`: The ACL is being associated with the listener. `Associated`: The ACL is associated with the listener. `Dissociating`: The ACL is being disassociated from the listener.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetListenersListenerAclConfigAclRelationResult(
            string aclId,

            string status)
        {
            AclId = aclId;
            Status = status;
        }
    }
}

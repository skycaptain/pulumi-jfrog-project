// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.JfrogProject.Outputs
{

    [OutputType]
    public sealed class ProjectAdminPrivilege
    {
        /// <summary>
        /// Enables a project admin to define the resources to be indexed by Xray
        /// </summary>
        public readonly bool IndexResources;
        /// <summary>
        /// Allows the Project Admin to manage Platform users/groups as project members with different roles.
        /// </summary>
        public readonly bool ManageMembers;
        /// <summary>
        /// Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
        /// </summary>
        public readonly bool ManageResources;

        [OutputConstructor]
        private ProjectAdminPrivilege(
            bool indexResources,

            bool manageMembers,

            bool manageResources)
        {
            IndexResources = indexResources;
            ManageMembers = manageMembers;
            ManageResources = manageResources;
        }
    }
}
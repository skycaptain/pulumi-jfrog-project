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
    public sealed class ProjectMember
    {
        /// <summary>
        /// Must be existing Artifactory user
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of pre-defined Project or custom roles
        /// </summary>
        public readonly ImmutableArray<string> Roles;

        [OutputConstructor]
        private ProjectMember(
            string name,

            ImmutableArray<string> roles)
        {
            Name = name;
            Roles = roles;
        }
    }
}
# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .environment import *
from .group import *
from .project import *
from .provider import *
from .repository import *
from .role import *
from .share_repository import *
from .share_repository_with_all import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_jfrog_project.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_jfrog_project.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "jfrog-project",
  "mod": "index/environment",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/environment:Environment": "Environment"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/group",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/group:Group": "Group"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/project",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/project:Project": "Project"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/repository",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/repository:Repository": "Repository"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/role",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/role:Role": "Role"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/shareRepository",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/shareRepository:ShareRepository": "ShareRepository"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/shareRepositoryWithAll",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/shareRepositoryWithAll:ShareRepositoryWithAll": "ShareRepositoryWithAll"
  }
 },
 {
  "pkg": "jfrog-project",
  "mod": "index/user",
  "fqn": "pulumi_jfrog_project",
  "classes": {
   "jfrog-project:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "jfrog-project",
  "token": "pulumi:providers:jfrog-project",
  "fqn": "pulumi_jfrog_project",
  "class": "Provider"
 }
]
"""
)

<!--
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
-->

# Releasing Apache Casbin (Incubating) Govaluate

Apache Casbin (Incubating) is an effort undergoing incubation at the Apache
Software Foundation (ASF), sponsored by the Apache Incubator PMC.
Incubation is required of all newly accepted projects until a further review
indicates that the infrastructure, communications, and decision making process
have stabilized in a manner consistent with other successful ASF projects.
While incubation status is not necessarily a reflection of the completeness or
stability of the code, it does indicate that the project has yet to be fully
endorsed by the ASF.

## Source candidate

1. Select one reviewed commit with passing CI. Review inherited MIT provenance,
   LICENSE, NOTICE, DISCLAIMER, dependencies and the actual archive inventory.
   Keep the module/import path `github.com/casbin/govaluate`. An Apache repository
   transfer does not require changing Go imports.
2. Create the next candidate tag, for example `v1.11.0-rc1`, at that commit.
   The source workflow uploads an Actions artifact only. It does not create a
   GitHub Release, push tags or publish a final Go module. A final tag does not
   run this workflow. Do not create the final tag before release approval.
3. Download the candidate tarball and verify its SHA-512. Its basename is
   `apache-casbin-govaluate-1.11.0-incubating-src.tar.gz`; the staging directory
   identifies the RC. Independently extract, build and test the exact bytes.
   Test both default and `-tags cache` builds on Go 1.24 or later. Run RAT and
   review its unknown/binary/excluded files individually; inherited MIT files
   must retain their license and attribution, not receive fabricated ASF grants.
   The opt-in torture test uses `GOVALUATE_TORTURE_TEST=true`.
4. The release manager signs that exact tarball locally with a key in the Apache
   Casbin KEYS file and stages the archive, `.asc` and `.sha512` under the
   [ASF development distribution area](https://dist.apache.org/repos/dist/dev/incubator/casbin/).
   Verify fresh public downloads, signatures, checksums and extracted tests.
5. Request PPMC endorsement on dev@casbin.apache.org and binding IPMC approval
   on general@incubator.apache.org. Keep both votes and exact candidate hashes
   recorded. Changed contents require a new RC and new applicable votes.

## Approved publication

1. After both approvals, promote the voted source tarball and its original
   signature/checksum to the ASF release distribution area. Do not rebuild it.
   Verify the public downloads against the voted originals.
2. Create the final Go tag (for example `v1.11.0`) at the approved RC commit.
   Verify that an existing tag resolves to that commit; stop on any conflict,
   and never force-update a published version. Go proxies resolve versions from
   repository tags independently of GitHub Releases. Verify that the legacy
   module path resolves through the GitHub redirect to this Apache repository,
   then download the final module in a clean consumer and run its tests.
   A changed go.mod needs separate candidate review; do not silently switch a
   downstream voted source from a pseudo-version to the final tag afterward.
3. Create the GitHub Release manually, attaching the promoted ASF originals
   downloaded and verified in step 1. Do not use a tag-triggered rebuild or
   `--clobber`. For a retry, compare existing assets byte for byte and stop on
   a mismatch. GitHub-generated source downloads are separate from the voted
   ASF source distribution. Include the full DISCLAIMER text, Apache source
   download and verification links, exact commit and vote-result links in the
   release notes. Mark any GitHub candidate as a prerelease.
4. Verify the download page, release display and Go consumer path; announce and
   preserve the vote, distribution, module and announcement evidence. Distinguish
   releases predating incubation from Apache releases on their display pages.

GitHub/tag access, signing, ASF dev/release SVN permissions and platform
administration are distinct. Confirm access without printing credentials.
Use the [ASF release policy](https://www.apache.org/legal/release-policy.html),
[Incubation Policy](https://incubator.apache.org/policy/incubation.html#releases)
and [distribution guidelines](https://incubator.apache.org/guides/distribution.html).

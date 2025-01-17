## Changelog (Current version: 1.3.0)

-----------------

## 1.3.0 (2016 Apr 12)

### Release Notes

* __BREAKING__ : Now you can delete/reset environment variables by setting the value to empty string ("").
  Previously an empty value (e.g. `- an_input: ""`) was just ignored,
  now it actually sets the value to an empty value.
* __NEW__ : Plugins ("beta"), to extend the `bitrise` functionality without modifying the "core"
  * Install plugin: `bitrise plugin install [PLUGIN_NAME]`
  * Delete plugin: `bitrise plugin delete [PLUGIN_NAME]`
  * List installed plugins: `bitrise plugin list`
  * Run plugin: `bitrise :[PLUGIN_NAME]`
  * bitrise cli now installs default plugins at `bitrise setup`.
* __NEW__ docs & tutorials:
  * Step Development Guideline: `_docs/step-development-guideline.md`
  * React Native Tutorial: `_examples/tutorials/react-native`
* Step Template revision:
  * Generic format update
  * Using `change-workdir` instead of a custom script
  * Added a `share-this-step` workflow for quick & easy step sharing
* New `--format=json` & `--format=yml` output modes (beta, only a few command supports this flag right now)
  * Added a new `version` command which now supports the `--format=json`
* README.md updates
  * Tooling and `--format=json`
  * Share your own Step section
* `bitrise workflows`, `bitrise step-info [STEP_ID]`, `bitrise step-list` cmd output improvements
* `bitrise validate` cmd updates:
  * workflow id validation
  * check for duplicated inputs
* bitrise output log improvements
  * Now build log contains deprecation infos about deprecated steps
* typo fixes
* Requires new `envman` (`1.1.0`) and `stepman` (`0.9.18`) versions - it'll
  auto-install these at first run if the required new versions are not found.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.3.0/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.2.4 -> 1.3.0

* [5181b50] Krisztián Gödrei - Merge pull request #339 from godrei/cache_dir_env (2016 Apr 11)
* [ce86551] godrei - cache dir env (2016 Apr 11)
* [da1edb6] Krisztián Gödrei - Merge pull request #337 from godrei/plugin_update (2016 Apr 11)
* [3f28676] godrei - plugin update fix & analytics 0.9.4 (2016 Apr 11)
* [590015d] Krisztián Gödrei - Merge pull request #336 from godrei/version_cmd (2016 Apr 07)
* [3c2bfe8] godrei - cleanup (2016 Apr 07)
* [d784f98] godrei - include commit in full version (2016 Apr 07)
* [f33738b] godrei - outputFormat moved to output package (2016 Apr 07)
* [e868e78] Krisztián Gödrei - Merge pull request #334 from bitrise-io/update-react-example (2016 Apr 07)
* [534119c] Krisztián Gödrei - Merge pull request #335 from godrei/build_number (2016 Apr 06)
* [bc2bbe6] godrei - move binaries to deploy dir (2016 Apr 06)
* [da00fb2] godrei - PR fix (2016 Apr 06)
* [b999737] Agnes Vasarhelyi - Update bitrise.yml (2016 Apr 06)
* [6feb388] godrei - build number (2016 Apr 06)
* [0bc68e7] vasarhelyia - Remove local path (2016 Apr 06)
* [19cde67] vasarhelyia - Use dedicated steps (2016 Apr 06)
* [a02d148] Krisztián Gödrei - Merge pull request #333 from godrei/master (2016 Apr 06)
* [e0e98ec] godrei - release notes (2016 Apr 06)
* [8d4e86a] godrei - v1.3.0 (2016 Apr 06)
* [471b2ab] godrei - bitrise.yml typo fix (2016 Apr 06)
* [19e845a] Krisztián Gödrei - Merge pull request #332 from godrei/prepare_for_relelase (2016 Apr 06)
* [e5187ac] godrei - removed old changelogs (2016 Apr 06)
* [77ee955] godrei - prepare for release (2016 Apr 06)
* [f34df79] Krisztián Gödrei - Merge pull request #331 from godrei/feature/default_plugins (2016 Apr 06)
* [68c93fa] godrei - default analytics plugin min version update (2016 Apr 06)
* [e29d733] godrei - log installed plugin (2016 Apr 05)
* [d453d8b] godrei - default plugins (2016 Apr 05)
* [2b398f1] Krisztián Gödrei - Merge pull request #330 from godrei/duplicated_inputs (2016 Apr 05)
* [5d8bc29] godrei - test fixes (2016 Apr 05)
* [5d142cb] godrei - check for duplicated inputs (2016 Apr 05)
* [e0402d0] Krisztián Gödrei - Merge pull request #329 from godrei/godep-update (2016 Apr 05)
* [fb51075] godrei - godep update (2016 Apr 05)
* [4d78370] Krisztián Gödrei - Merge pull request #328 from godrei/separate_packages (2016 Apr 05)
* [8ff9a6e] Krisztián Gödrei - Merge pull request #327 from godrei/ci_updates (2016 Apr 05)
* [1e57cb7] godrei - cleanup (2016 Apr 05)
* [ace4e2b] godrei - separate bitrise packages (2016 Apr 05)
* [2ec5090] godrei - bitrise.yml updates (2016 Apr 04)
* [4c2ca4a] Viktor Benei - Merge pull request #326 from godrei/deprecate_wildcard_workflow_id (2016 Apr 01)
* [5477a8e] godrei - PR fix (2016 Apr 01)
* [722cc71] godrei - [b4475fe] [204cd7c] typo fix [b672304] workflow id validation (2016 Apr 01)
* [dbc2e95] Krisztián Gödrei - Merge pull request #325 from godrei/global_step_info (2016 Mar 18)
* [95448a1] Krisztián Gödrei - deprecate infos (2016 Mar 18)
* [85ac8ce] Krisztián Gödrei - Merge pull request #324 from godrei/skip_if_empty (2016 Mar 18)
* [866d86e] Krisztián Gödrei - instal bitrise tool in _prepare_and_setup workflow (2016 Mar 18)
* [797a42a] Krisztián Gödrei - bitrise.yml updates (2016 Mar 18)
* [463e9ae] Krisztián Gödrei - plugin update for new envman version, release configs, bitrise.yml updates (2016 Mar 18)
* [f74bee9] Krisztián Gödrei - removed local reference in create_changelog workflow & skip_if_empty unit test new environment variable (skip_if_empty) handling (2016 Mar 18)
* [a92c659] Krisztián Gödrei - Merge pull request #323 from godrei/test_fix (2016 Mar 18)
* [384ee68] Krisztián Gödrei - plugin version check fix (2016 Mar 17)
* [ed5b5ca] Krisztián Gödrei - use bitrise-core test repos (2016 Mar 17)
* [8542528] Krisztián Gödrei - removed download test (2016 Mar 17)
* [18a1ac9] Viktor Benei - Merge pull request #321 from godrei/events (2016 Mar 17)
* [c8f88ba] Krisztián Gödrei - envman test fix, typo, error log fix (2016 Mar 16)
* [b86b4c2] Viktor Benei - Merge pull request #322 from anas10/patch-1 (2016 Mar 16)
* [64e9f20] Anas AIT ALI - Update README.md (2016 Mar 16)
* [9c4fe9b] Krisztián Gödrei - fixed TestExpandEnvs (2016 Mar 11)
* [e024a42] Krisztián Gödrei - check for updatest, before using the plugin, but only if not CI mode (2016 Mar 11)
* [922e0e8] Krisztián Gödrei - install binary by platforms (2016 Mar 05)
* [2e397fc] Krisztián Gödrei - create plugin data dir at install, check for plugin new version fix (2016 Mar 05)
* [1e61ab7] Krisztián Gödrei - log fixes, run_test update (2016 Mar 05)
* [8f6a350] Krisztián Gödrei - create plugin data dir (2016 Mar 03)
* [242b493] Krisztián Gödrei - trigger event DidFinishRun dont print any logs after workflow summary (2016 Mar 03)
* [814a2b9] Viktor Benei - Merge pull request #320 from viktorbenei/master (2016 Mar 03)
* [49f4234] Viktor Benei - experimental/upload-download-bitrise-yml : updated for Bitrise CLI 1.3 & made it better for quick fixing (download, fix, upload) (2016 Mar 03)
* [a901456] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2016 Mar 03)
* [bba88bd] Viktor Benei - Dockerfile: use go 1.6 (2016 Mar 03)
* [90d32ed] Viktor Benei - yml format update for new Bitrise CLI compatibility (2016 Mar 03)
* [7dad162] Krisztián Gödrei - Merge pull request #319 from godrei/plugin (2016 Mar 01)
* [844ea97] Krisztián Gödrei - NewEnvJSONList instead of CreateFromJSON (2016 Mar 01)
* [53cef9e] Krisztián Gödrei - test updates (2016 Mar 01)
* [8329195] Krisztián Gödrei - version fix (2016 Mar 01)
* [55e6df5] Krisztián Gödrei - plugin requirement's required min version is required, minor fixes (2016 Mar 01)
* [9cc7d95] Krisztián Gödrei - version package instead of hard coded version (2016 Mar 01)
* [3cc0bf9] Krisztián Gödrei - base plugin handling (2016 Mar 01)
* [3acccd3] Viktor Benei - Merge pull request #318 from tomgilder/patch-1 (2016 Feb 28)
* [dc423d8] Tom Gilder - Fix spelling mistake (2016 Feb 28)
* [6eed6a7] Viktor Benei - script content fix (multiline) (2016 Feb 22)
* [59696c6] Viktor Benei - Merge pull request #317 from dag-io/master (2016 Feb 17)
* [52322cf] Damien Gavard - Fix typo (2016 Feb 17)
* [de931be] Viktor Benei - Merge pull request #314 from bitrise-io/update-install-guide (2016 Feb 09)
* [6220657] vasarhelyia - Update install info (2016 Feb 09)
* [c452548] Viktor Benei - Merge pull request #313 from bitrise-io/update-react-native-example (2016 Feb 07)
* [0ccbac7] vasarhelyia - Update workflow name (2016 Feb 06)
* [ae202b1] vasarhelyia - Add sample app yml (2016 Feb 06)
* [3ecbd13] Viktor Benei - Merge pull request #312 from bitrise-io/slack-channel-badge (2016 Feb 04)
* [7c16467] vasarhelyia - Add Slack channel badge (2016 Feb 04)
* [3b64e94] Viktor Benei - Merge pull request #311 from birmacher/typo (2016 Jan 26)
* [2e589c7] birmacher - typo fix (2016 Jan 26)
* [32a52d5] Viktor Benei - Merge pull request #310 from viktorbenei/master (2015 Dec 22)
* [96037ac] Viktor Benei - godeps-update fix (2015 Dec 22)
* [1703e25] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Dec 22)
* [d5d2a66] Viktor Benei - godeps-update (2015 Dec 22)
* [956bee4] Viktor Benei - bumped required envman (to 1.1.0) & stepman (to 0.9.18) versions (2015 Dec 22)
* [60bd807] Viktor Benei - README: intro one-liner text revision (2015 Dec 17)
* [d94054e] Viktor Benei - Merge pull request #309 from viktorbenei/master (2015 Dec 17)
* [b808e12] Viktor Benei - LOG : if config (bitrise.yml) is not valid include the path of the file (2015 Dec 17)
* [35614c0] Viktor Benei - LOG : if local step info print fails it'll print the path of the YML in the logs (2015 Dec 17)
* [b4ba81c] Viktor Benei - FIX : typo: "cofing" -> "config" & "faild" -> "failed" (2015 Dec 17)
* [f98a3da] Viktor Benei - Merge pull request #306 from godrei/changelog (2015 Dec 16)
* [866127d] Krisztián Gödrei - create_changelog workflow for automatic changelog generation based on commits from last tag on master (2015 Dec 16)
* [631a097] Viktor Benei - point highlights in Development Guideline (2015 Dec 16)
* [a47606f] Viktor Benei - Development Guideline section revision in README (2015 Dec 16)
* [469cc5f] Viktor Benei - Merge pull request #305 from godrei/format_version (2015 Dec 15)
* [625dc38] Krisztián Gödrei - format version (2015 Dec 15)
* [98d74a5] Viktor Benei - Merge pull request #304 from godrei/godeps-update (2015 Dec 15)
* [92be446] Krisztián Gödrei - godeps update (2015 Dec 15)
* [a715ba7] Viktor Benei - Merge pull request #303 from godrei/plugin_compatibility (2015 Dec 15)
* [0a246c2] Krisztián Gödrei - godeps update (2015 Dec 14)
* [c630110] Krisztián Gödrei - plugin fixes (2015 Dec 14)
* [5f14af7] Viktor Benei - Merge pull request #302 from viktorbenei/master (2015 Dec 12)
* [3a31596] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Dec 12)
* [a5243bc] Viktor Benei - Merge pull request #301 from godrei/plugin (2015 Dec 12)
* [86573cb] Krisztián Gödrei - PR fix (2015 Dec 12)
* [79dc0d4] Krisztián Gödrei - PR fix (2015 Dec 12)
* [645b2bf] Krisztián Gödrei - PR fixes (2015 Dec 12)
* [d062952] Krisztián Gödrei - plugin install, delete, list (2015 Dec 12)
* [a4a3511] Viktor Benei - godeps-update (2015 Dec 12)
* [90653cb] Viktor Benei - 1.3.0-pre (2015 Dec 12)
* [e7a6dfa] Krisztián Gödrei - base plugin handling (2015 Dec 12)
* [50c8c83] Viktor Benei - Merge pull request #300 from godrei/workflows (2015 Dec 08)
* [324b08e] Krisztián Gödrei - improvements (2015 Dec 08)
* [03668e9] Viktor Benei - Merge pull request #299 from viktorbenei/master (2015 Dec 07)
* [9db86b4] Viktor Benei - typo in bitrise.yml workflow description (2015 Dec 07)
* [a2c051a] Viktor Benei - Merge pull request #297 from godrei/step_info_fix (2015 Dec 07)
* [17ca336] Viktor Benei - Merge pull request #298 from godrei/workflows_fix (2015 Dec 07)
* [7da0a6b] Krisztián Gödrei - yellow no summary/description (2015 Dec 07)
* [0a5f69f] Krisztián Gödrei - step-info, step-list fixes (2015 Dec 07)
* [b8b385f] Viktor Benei - Merge pull request #296 from godrei/delete_envs (2015 Dec 07)
* [74cdd7a] Krisztián Gödrei - change log fix (2015 Dec 07)
* [1e2978a] Krisztián Gödrei - delete env + test (2015 Dec 07)
* [7ddbd10] Viktor Benei - Merge pull request #295 from viktorbenei/master (2015 Dec 07)
* [a87bf87] Viktor Benei - godeps-update (2015 Dec 07)
* [a34bec8] Viktor Benei - Merge pull request #294 from godrei/workflow_list (2015 Dec 07)
* [780de41] Krisztián Gödrei - workflow list (2015 Dec 07)
* [19547ac] Viktor Benei - Update step-development-guideline.md (2015 Dec 05)
* [e7e18f9] Viktor Benei - Do not use submodules, or require any other resource, downloaded on-demand (2015 Dec 05)
* [1d61661] Viktor Benei - clarification (2015 Dec 05)
* [3b3adb0] Viktor Benei - Update README.md (2015 Dec 05)
* [8873c57] Viktor Benei - Create step-development-guideline.md (2015 Dec 05)
* [01e3f36] Viktor Benei - Share your own Step section (2015 Nov 19)
* [cfb5e5b] Viktor Benei - Merge pull request #293 from viktorbenei/master (2015 Nov 09)
* [a451eb3] Viktor Benei - readme : tooling and `--format=json` (2015 Nov 09)
* [acd8356] Viktor Benei - godeps update (2015 Nov 09)
* [9865ba1] Viktor Benei - Merge pull request #292 from viktorbenei/master (2015 Nov 09)
* [cfb4319] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Nov 09)
* [0c024c9] Viktor Benei - `yml` option added/enabled for Output Format (2015 Nov 09)
* [4b64cb7] Viktor Benei - Merge pull request #291 from viktorbenei/master (2015 Nov 07)
* [9ea7723] Viktor Benei - test fix (2015 Nov 07)
* [69b5b9a] Viktor Benei - new packages : configs and output - to help with the new `--format=json` output mode ; added a new `version` command which now supports the `--format=json` flag (2015 Nov 07)
* [0404f7c] Viktor Benei - Merge pull request #290 from viktorbenei/master (2015 Nov 06)
* [2b2999b] Viktor Benei - step template revision : generic format update, using `change-workdir` instead of a custom script, and added a `share-this-step` workflow for quick & easy step sharing (2015 Nov 06)
* [a7ac606] Viktor Benei - Create .gitignore (2015 Nov 04)
* [00a0ab3] Viktor Benei - bitrise.io/cli (2015 Nov 04)
* [921d903] Viktor Benei - Update README.md (2015 Nov 04)


## 1.2.4 (2015 Nov 02)

### Release Notes

* __envman__ updated to `1.0.0`, which also includes a new ENV size limit feature. You can read more about the release at: https://github.com/bitrise-io/envman/releases/tag/1.0.0

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.2.4/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.2.3 -> 1.2.4

* [044a87c] Viktor Benei - Merge pull request #288 from viktorbenei/master (2015 Nov 02)
* [9d0ba2d] Viktor Benei - 1.2.4 changelog (2015 Oct 31)
* [7e1a93a] Viktor Benei - Merge pull request #287 from viktorbenei/master (2015 Oct 31)
* [8e642cb] Viktor Benei - CI fix (2015 Oct 31)
* [ce1903e] Viktor Benei - 1.2.4 (2015 Oct 31)
* [57c0f18] Viktor Benei - Merge pull request #286 from viktorbenei/master (2015 Oct 31)
* [69849d1] Viktor Benei - just one more.. (2015 Oct 31)
* [4d05b0f] Viktor Benei - hopefully last CI fix :) (2015 Oct 31)
* [87f05be] Viktor Benei - one more CI workflow fix (2015 Oct 31)
* [f496b5b] Viktor Benei - bitrise.yml fix (2015 Oct 31)
* [3c44e47] Viktor Benei - godeps-update (2015 Oct 31)
* [733ae3d] Viktor Benei - envman : 1.0.0 (2015 Oct 31)
* [1537968] Viktor Benei - bitrise.yml test_and_install fix (2015 Oct 31)


## 1.2.3 (2015 Oct 19)

### Release Notes

* __FIX__ : `bitrise share create` had a parameter issue, calling `stepman share create` with wrong `--stepid` param. Fixed.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.2.3/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.2.2 -> 1.2.3

* [89c8ebf] Viktor Benei - Merge pull request #285 from viktorbenei/master (2015 Oct 19)
* [337cf23] Viktor Benei - golint : removed `set -e` (2015 Oct 19)
* [3e89237] Viktor Benei - changelog & godeps-update (2015 Oct 19)
* [4c09ffe] Viktor Benei - Merge pull request #284 from viktorbenei/master (2015 Oct 19)
* [a9e3f59] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Oct 19)
* [fd68787] Viktor Benei - 1.2.3 (2015 Oct 19)
* [8f65724] Viktor Benei - Merge pull request #283 from viktorbenei/master (2015 Oct 19)
* [3424a7c] Viktor Benei - next version changelog (2015 Oct 13)
* [aa15018] Viktor Benei - Merge pull request #282 from gkiki90/bitrise_share_fix (2015 Oct 12)
* [cd0ff8a] Krisztian Godrei - share fix (2015 Oct 12)


## 1.2.2 (2015 Oct 12)

### Release Notes

* __Fixed__ step log, at build failed mode (at step log footer section Issue tracker and Source row trimming fixed).
* __Fixed__ `bitrise validate` if called with `--format=json` : in case the validation failed it printed two JSON responses instead of just one. Fixed.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.2.2/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.2.1 -> 1.2.2

* [c0abd8c] Viktor Benei - Merge pull request #281 from viktorbenei/master (2015 Oct 12)
* [8d36ef3] Viktor Benei - 1.2.2 (2015 Oct 12)
* [7e06232] Viktor Benei - Merge pull request #280 from viktorbenei/1.2.2-pre (2015 Oct 09)
* [c65bf52] Viktor Benei - 1.2.2-pre (2015 Oct 09)
* [1ef1098] Viktor Benei - Merge pull request #279 from gkiki90/typo (2015 Oct 09)
* [08e6e24] Krisztian Godrei - log fix (2015 Oct 09)
* [74e697d] Krisztian Godrei - typo (2015 Oct 09)
* [44d2b40] Viktor Benei - Merge pull request #278 from gkiki90/validate_fix (2015 Oct 09)
* [e3b0d1c] Krisztian Godrei - validate fix (2015 Oct 09)
* [fec3772] Viktor Benei - Merge pull request #277 from gkiki90/trimming_issue_fix (2015 Oct 07)
* [5ca49d0] Krisztian Godrei - changelog (2015 Oct 07)
* [fe34675] Krisztian Godrei - fixed step log footer trimming issue (2015 Oct 07)


## 1.2.1 (2015 Oct 07)

### Release Notes

* __FIX__ : `trigger_map` handling in Pull Request mode: if the pattern does match an item which has `is_pull_request_allowed=false` it won't fail now, it'll just skip the item and the next one will be tried.
* __new command__ : `bitrise share`, to share your step through `bitrise` (this is just a wrapper around `stepman share`, does exactly the same, but hopefully it's a bit more convenient if you never used `stepman` directly before).
* __new flag__ : similar to the `--ci` flag the Pull Request mode can now be allowed by calling any command with `bitrise --pr [command]`.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.2.1/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.2.0 -> 1.2.1

* [7e3aa83] Viktor Benei - Merge pull request #276 from viktorbenei/master (2015 Oct 07)
* [366b459] Viktor Benei - v1.2.1 (2015 Oct 07)
* [dd1037d] Viktor Benei - Merge pull request #275 from gkiki90/changelog (2015 Oct 06)
* [24004c1] Krisztian Godrei - changelog (2015 Oct 06)
* [cdd4662] Viktor Benei - removed Gitter (2015 Oct 06)
* [62026fd] Viktor Benei - Merge pull request #274 from gkiki90/share (2015 Oct 05)
* [2648f37] Krisztian Godrei - share (2015 Oct 05)
* [b074d00] Viktor Benei - Merge pull request #273 from gkiki90/PR_mode (2015 Oct 05)
* [d7e28db] Krisztian Godrei - PR fix (2015 Oct 05)
* [0136c97] Viktor Benei - Merge pull request #272 from gkiki90/PR_mode (2015 Oct 05)
* [f539e6f] Viktor Benei - Merge pull request #271 from viktorbenei/master (2015 Oct 05)
* [f3be793] Krisztian Godrei - test fixes (2015 Oct 05)
* [aeffd47] Krisztian Godrei - ci iml fix, pr mode fix (2015 Oct 05)
* [2d5296f] Krisztian Godrei - removed unnecessary descriptions (2015 Oct 05)
* [3309132] Krisztian Godrei - Merge branch 'ci_fix' into PR_mode (2015 Oct 05)
* [bfb6b27] Krisztian Godrei - typo fix, pr mode trigger fix (2015 Oct 05)
* [20e98bf] Krisztian Goedrei - ci fix (2015 Oct 05)
* [75a188f] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Oct 03)
* [94c9187] Viktor Benei - install envman & stepman with `curl -fL` as it's the new recommended way (2015 Oct 02)


## 1.2.0 (2015 Oct 02)

### Release Notes

* __BREAKING__ / __FIX__ : If `bitrise trigger` is called with a trigger pattern that doesn't match any expression in `trigger_map` the workflow with the same name as the trigger pattern name **will no longer be selected**. This feature proved to have more issues than pros.
* __DEPRECATION__ : the previous `dependencies` property is now deprecated. From now on, dependencies should be declared in the `deps` property which has a new syntax, grouped by dependency managers. The syntax will be extended in the near future but in a backward compatible way.
  Supported dependency managers: `brew`, `apt_get` and `check_only` (for checking dependencies which can't be installed automatically, ex the `Xcode.app`).
  Example:

  ```
  - step:
      deps:
        brew:
        - name: cmake
        - name: git
        - name: node
        apt_get:
        - name: cmake
        check_only:
        - name: xcode
  ```
* Improved validate command output.
* __BREAKING__ : if you don't specify the version of a step `bitrise` will now try to update the local Step Lib cache before using the Step. Previously the latest version **available in the local cache** was used, but this caused more confusion. The local cache is still used in this case if the Step Lib can't be updated, it'll still work in case of a network issue.
* __BREAKING__ : From `bitrise step-info` the `--id` flag was removed, the first cli param used as step id, no need to write the `--id` flag anymore. Example: `bitrise step-info script` instead of `bitrise step-info --id script`.
* __IMPORTANT__ : `format_version` bumped to `1.1.0`, which means that the `bitrise.yml` generate with this `bitrise` version won't be compatible with previous `bitrise` versions. Previous `bitrise.yml`s of course still work with this new `bitrise` version.
* Now you can use all your environment variables (secrets, app, workflow envs and step outputs) in `run_if` fields and in step inputs.
* `bitrise step-info` got a new option `--step-yml` flag, which allows printing step info from the specified `step.yml` directly (useful for local Step development).
* Step inputs got a new field: `IsTemplate` / `is_template`. This field indicates whether the value contains template expressions which should be evaluated before using the value, just like in case of `is_expand`. The template expression have to be written in Go's template language, and can use the same properties as `run_if` templates can. Example:

  ```
  - script:
    title: Template example
    inputs:
    - content: |-
        {{if .IsCI}}
        echo "CI mode"
        {{else}}
        echo "not CI mode"
        {{end}}
      opts:
        is_template: true
  ```
* Improved environment and input value and options casting:
    * Now you can use `"NO"`, `"No"`, `"YES"`, `"Yes"`, `true`, `false`, `"true"`, `"false"` in every place `bitrise` expects a bool value (ex: `is_expand`).
    * Every field where `bitrise` expects a string in now casted into a string. This means that you can now use `true` and `false` instead of `"true"` and `"false"` in `value_options`. Same is true for the input and environments value itself, so you can now write `true` instead of `"true"` and it'll still be casted to string.
* Pull Request and CI mode handling extension: related flag environment variables can now be defined in `secrets` / `inventory` as well.
* `bitrise` now prints if it runs in "Pull Request mode", just like it did for "CI" mode before.
* Step info logging got a complete revision, to make it more helpful, especially in case the Step fails. It now included the Step's issue tracker and source code URL infos in the log directly.
* __FIX__ : `--log-level` handling fix, the previous version had issues if the log level was set to `debug`.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.2.0/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.1.2 -> 1.2.0

* [9c785d1] Viktor Benei - Merge pull request #270 from viktorbenei/master (2015 Oct 02)
* [67ffc05] Viktor Benei - changelog for 1.2.0 (2015 Oct 02)
* [81dad75] Viktor Benei - v1.2.0 (2015 Oct 02)
* [b0ae673] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Oct 02)
* [59f7cee] Viktor Benei - updated next-version changelog (2015 Oct 02)
* [260a2f4] Viktor Benei - Merge pull request #269 from viktorbenei/master (2015 Oct 02)
* [56ac8b8] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Oct 02)
* [e33aaae] Viktor Benei - Merge pull request #268 from gkiki90/fake_home_fix (2015 Oct 02)
* [f328843] Krisztian Goedrei - PR fix (2015 Oct 02)
* [e80e2a3] Viktor Benei - Merge pull request #267 from gkiki90/fake_home_fix (2015 Oct 02)
* [5e6c187] Viktor Benei - Merge pull request #266 from gkiki90/dep_logs (2015 Oct 02)
* [e65ee13] Krisztian Goedrei - fake home fix (2015 Oct 02)
* [43b31c2] Viktor Benei - `DefaultIsTemplate` typo fix (2015 Oct 02)
* [be50cfe] Viktor Benei - godeps-update (2015 Oct 02)
* [99e7c63] Viktor Benei - required envman and stepman version bumps (2015 Oct 02)
* [306459d] Krisztian Goedrei - check only log (2015 Oct 02)
* [2b13244] Krisztian Goedrei - dep logs (2015 Oct 02)
* [3997ca2] Viktor Benei - Merge pull request #265 from gkiki90/envman_init_fix (2015 Oct 02)
* [4ad7773] Krisztian Goedrei - no internet connection (2015 Oct 02)
* [42d326c] Krisztian Goedrei - step info version fix (2015 Oct 02)
* [591e145] Krisztian Goedrei - fail test (2015 Oct 02)
* [1a8f411] Krisztian Goedrei - fixes (2015 Oct 02)
* [3c74d22] Krisztian Goedrei - fixes (2015 Oct 02)
* [7185c74] Krisztian Goedrei - fixed envman init (2015 Oct 02)
* [f9766e9] Viktor Benei - Merge pull request #264 from viktorbenei/master (2015 Oct 02)
* [49bc8d2] Viktor Benei - base codeclimate config (2015 Oct 02)
* [fe3f460] Viktor Benei - Merge pull request #263 from gkiki90/changelog (2015 Oct 01)
* [08176a8] Krisztian Goedrei - changelog (2015 Oct 01)
* [b14bf8b] Viktor Benei - Merge pull request #261 from gkiki90/latest (2015 Oct 01)
* [87fd6f8] Krisztian Goedrei - godeps (2015 Oct 01)
* [40782a1] Krisztian Goedrei - godep, fixes (2015 Oct 01)
* [4abf72e] Krisztian Goedrei - godep save (2015 Oct 01)
* [8ea8560] Krisztian Goedrei - merge (2015 Oct 01)
* [72f0f75] Viktor Benei - Merge pull request #262 from gkiki90/PR_mode (2015 Oct 01)
* [59976e4] Krisztian Goedrei - fixes (2015 Oct 01)
* [381835f] Krisztian Goedrei - PR & CI mode fix (2015 Oct 01)
* [49b128f] Viktor Benei - Merge pull request #260 from gkiki90/trigger_fix (2015 Sep 29)
* [d4ed963] Krisztian Goedrei - test fix (2015 Sep 28)
* [4b410a8] Krisztian Goedrei - start (2015 Sep 28)
* [8226ad3] Viktor Benei - Merge pull request #258 from gkiki90/is_template (2015 Sep 28)
* [f8ff04b] Krisztian Goedrei - envlist as template input (2015 Sep 27)
* [3d5a371] Viktor Benei - Merge pull request #257 from gkiki90/is_template (2015 Sep 25)
* [ca6f87e] Krisztian Goedrei - template run test (2015 Sep 25)
* [3fd9382] Krisztian Goedrei - template handling, godep (2015 Sep 25)
* [1a462bd] Krisztian Goedrei - template in tests (2015 Sep 25)
* [e394774] Krisztian Goedrei - IsTemplate in model methods (2015 Sep 25)
* [98b68d6] Krisztian Goedrei - require in test (2015 Sep 25)
* [e11fc8a] Viktor Benei - Merge pull request #256 from viktorbenei/master (2015 Sep 24)
* [1b4337c] Viktor Benei - `_tmp` added to .gitignore (2015 Sep 24)
* [53b798a] Viktor Benei - step template README update (2015 Sep 24)
* [15bf241] Viktor Benei - updated `_step_template` (2015 Sep 24)
* [f787c38] Viktor Benei - Merge pull request #255 from viktorbenei/master (2015 Sep 24)
* [b037300] Viktor Benei - format_version bumped in bitrise.yml (2015 Sep 24)
* [5f83c5c] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 24)
* [4d3fcd9] Viktor Benei - v1.1.3-pre (2015 Sep 24)
* [fda85d8] Viktor Benei - Merge pull request #254 from gkiki90/deps (2015 Sep 24)
* [885bbf0] Krisztian Goedrei - docker installs sudo, dependencies bitrise yml for linux and osx (2015 Sep 24)
* [23d41df] Krisztian Goedrei - godeps (2015 Sep 24)
* [7cb5b1f] Krisztian Goedrei - new deps (2015 Sep 24)
* [ef2dc12] Krisztian Goedrei - new deps in progress (2015 Sep 23)
* [24e45a4] Viktor Benei - Merge pull request #253 from gkiki90/triggered_workflow (2015 Sep 22)
* [3ba8dbc] Krisztian Goedrei - trigger check, output format fixes (2015 Sep 22)
* [6000035] Viktor Benei - create release with docker-compose & trigger patterns for release operations (2015 Sep 22)
* [f8a47fd] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 22)
* [7646937] Viktor Benei - updated _tests/brew_publish.yml with more info/description (2015 Sep 22)
* [5c8f602] Viktor Benei - Merge pull request #251 from gkiki90/template (2015 Sep 22)
* [4d60973] Krisztian Goedrei - fix (2015 Sep 22)
* [ccbdf33] Viktor Benei - Merge pull request #252 from gkiki90/step-info (2015 Sep 22)
* [5e6fb13] Krisztian Goedrei - local step info (2015 Sep 22)
* [993d240] Krisztian Goedrei - fix (2015 Sep 22)
* [6aa0e29] Krisztian Goedrei - template fixes (2015 Sep 22)
* [38732b7] Krisztian Goedrei - new envman version (2015 Sep 22)
* [dbe314e] Krisztian Goedrei - template (2015 Sep 22)
* [3bdecde] Viktor Benei - Merge pull request #250 from viktorbenei/master (2015 Sep 22)
* [1918bbb] Viktor Benei - deps.go comment (2015 Sep 22)
* [0af2247] Viktor Benei - Godeps update, with a new `deps.go` to include other packages required only for running the `go test`s (2015 Sep 22)
* [7002f92] Viktor Benei - Merge pull request #249 from gkiki90/step_list_fix (2015 Sep 21)
* [bb52a08] Viktor Benei - Merge pull request #248 from gkiki90/ci_fix (2015 Sep 21)
* [ef055b8] Krisztian Goedrei - step-list fix (2015 Sep 21)
* [7784c07] Krisztian Goedrei - ci fix (2015 Sep 21)
* [362f972] Viktor Benei - Merge pull request #242 from gkiki90/validation_fix (2015 Sep 21)
* [3aa743f] Viktor Benei - Merge pull request #244 from gkiki90/step_info_fix (2015 Sep 21)
* [27a7dc0] Viktor Benei - Merge pull request #245 from gkiki90/ci_fix (2015 Sep 21)
* [e890448] Viktor Benei - Merge pull request #247 from gkiki90/typo_fix (2015 Sep 21)
* [97cccd9] Krisztian Goedrei - typo (2015 Sep 21)
* [817edc7] Krisztian Goedrei - step_info fix (2015 Sep 21)
* [42af9e3] Krisztian Goedrei - bitrise.yml updates (2015 Sep 21)


## 1.1.2 (2015 Sep 21)

### Release Notes

* __FIX__ : Step outputs are now exposed (available for subsequent steps) even if the Step fails.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.1.2/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.1.1 -> 1.1.2

* [ca8f796] Viktor Benei - Merge pull request #243 from viktorbenei/master (2015 Sep 21)
* [c4b00ad] Krisztian Goedrei - in progress (2015 Sep 21)
* [b1a6a4c] Viktor Benei - v1.1.2 (2015 Sep 21)
* [829391a] Krisztian Goedrei - ci fix start (2015 Sep 21)
* [f44e4b9] Krisztian Goedrei - validation fix (2015 Sep 21)
* [5e49994] Krisztian Goedrei - validation fix (2015 Sep 21)
* [9478b7b] Viktor Benei - Merge pull request #241 from gkiki90/output_env_list_fix (2015 Sep 21)
* [26d0d1d] Krisztian Goedrei - test fixes (2015 Sep 21)
* [571cdfe] Krisztian Goedrei - step output fix (2015 Sep 21)
* [1795cb1] Krisztian Goedrei - validation exit codes (2015 Sep 21)


## 1.1.1 (2015 Sep 18)

### Release Notes

* __FIX__ : If `$BITRISE_SOURCE_DIR` is defined in an environment with an empty value `bitrise` now skips the definition. Practically this means that if you have an empty `BITRISE_SOURCE_DIR` item in your Workflow or App Environment but you define a real value in your `.bitrise.secrets.yml` `bitrise` will now use the (real) value defined in `.bitrise.secrets.yml`, instead of going with the empty value defined in the Workflow environments.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.1.1/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.1.0 -> 1.1.1

* [35c9a74] Viktor Benei - Merge pull request #239 from viktorbenei/master (2015 Sep 18)
* [50e09fa] Viktor Benei - v1.1.1 (2015 Sep 18)
* [e364fb3] Viktor Benei - Merge pull request #238 from viktorbenei/master (2015 Sep 18)
* [4db999d] Viktor Benei - `BITRISE_SOURCE_DIR` handling fix: skip empty values (2015 Sep 18)
* [956abad] Viktor Benei - slack ENVs revision (2015 Sep 18)


## 1.1.0 (2015 Sep 18)

### Release Notes

* BITRISE build result log improvements:
    * step run summary contains step version, and update note, if new version available
    * build run summary step sections contains step version, and update note, if new version available
* __BREAKING/FIX__ : `bitrise trigger` will **NOT** select any workflow in Pull Request mode if the pattern does not match any of the `trigger_map` definition.
* unified `config` and `inventory` flag handling: you can specify paths with `--config` and `--inventory`, and base64 encoded direct input with `--config-base64` and `--inventory-base64`. Can be used by tools, to skip the need to write into files.
* __FIX/BREAKING__ : environment handling order : App Envs can now overwrite the values defined in inventory/secrets (in the last version the secrets/inventory could overwrite the App Envs).
* `validate` command accepts `--format` flag: `--format=[json/raw]` (default is `raw`)
* new command: `step-list` (lis of available steps in Step Lib) `bitrise step-list`
* new command: `step-info` (infos about defined step) `bitrise step-info --id script --version 0.9.0`
* revision of `normalize`, to generate a better list/shorter output list
* __NEW__ : `$BITRISE_SOURCE_DIR` now updated per step, and can be changed by the steps. `$BITRISE_SOURCE_DIR` can be use for defining a new working directory. Example: if you want to create CI workflow for your Go project you have to switch your working directory to the proper one, located inside the `$GOPATH` (this is a Go requirement). You can find an example below. This feature is still a bit in "experimental" stage, and we might add new capabilities in the future. Right now, if you want to re-define the `$BITRISE_SOURCE_DIR` you have to set an **absolute** path, no expansion will be performed on the specified value! So, you should **NOT** store a reference like `$GOPATH/src/your/project/path` as it's value, but the actual, absolute path!

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.1.0/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 1.0.0 -> 1.1.0

* [3a37cb9] Viktor Benei - Merge pull request #237 from viktorbenei/master (2015 Sep 18)
* [6abe52a] Viktor Benei - v1.1.0 - changelog (2015 Sep 18)
* [8658bb0] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 18)
* [d955103] Viktor Benei - version 1.1.0 (2015 Sep 18)
* [f1493ab] Viktor Benei - run.go : param name revision, for clarity (2015 Sep 18)
* [6f6b358] Viktor Benei - Next version: note about BITRISE_SOURCE_DIR (2015 Sep 18)
* [a13dda7] Viktor Benei - Dockerfile : pre-install required tools (2015 Sep 18)
* [3b75762] Viktor Benei - Merge pull request #235 from gkiki90/breaking_source_dir (2015 Sep 18)
* [1a42c32] Krisztian Goedrei - test fix (2015 Sep 18)
* [c37cf93] Krisztian Goedrei - code cleaning (2015 Sep 18)
* [93db802] Krisztian Goedrei - removed expand (2015 Sep 18)
* [8ddef2d] Krisztian Goedrei - tmp path (2015 Sep 18)
* [3f74b81] Krisztian Goedrei - source dir updated per step (2015 Sep 18)
* [828a59d] Viktor Benei - next version changelog (2015 Sep 18)
* [cec4252] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 18)
* [34e87da] Viktor Benei - Merge pull request #234 from gkiki90/ci_breaking_change (2015 Sep 18)
* [2f118ae] Krisztian Goedrei - comments moved to description (2015 Sep 18)
* [90936c3] Krisztian Goedrei - new yml (2015 Sep 17)
* [0b72e80] Krisztian Goedrei - godeps-update, min stepman version (2015 Sep 17)
* [571d9f0] Krisztian Goedrei - bitrise.yml, stepman update (2015 Sep 17)
* [9689dd6] Krisztian Goedrei - fix (2015 Sep 17)
* [e192671] Krisztian Goedrei - godeps-update (2015 Sep 17)
* [7ee4938] Krisztian Goedrei - no message (2015 Sep 17)
* [736067e] Krisztian Goedrei - godeps-update (2015 Sep 17)
* [36a2307] Krisztian Goedrei - godeps-update (2015 Sep 17)
* [f2a1a9a] Krisztian Goedrei - godeps-update (2015 Sep 17)
* [6ae1c20] Krisztian Goedrei - change log (2015 Sep 17)
* [1e30001] Krisztian Goedrei - validation formats (2015 Sep 17)
* [b5fc3c6] Krisztian Goedrei - step-info, step-list (2015 Sep 17)
* [44aec85] Krisztian Goedrei - test fix (2015 Sep 17)
* [56f271a] Krisztian Goedrei - code style (2015 Sep 17)
* [bddd467] Krisztian Goedrei - env order fix, test (2015 Sep 17)
* [716ad8a] Krisztian Goedrei - in progress (2015 Sep 17)
* [fa1dd91] Viktor Benei - Merge pull request #229 from gkiki90/step_version (2015 Sep 16)
* [46c762f] Viktor Benei - Merge pull request #230 from bazscsa/master (2015 Sep 16)
* [67f7e0b] Krisztian Goedrei - PR fix (2015 Sep 16)
* [d65f7d8] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Sep 16)
* [ffb84f7] Tamás Bazsonyi - Updated lesson links (2015 Sep 16)
* [8fa1397] Tamás Bazsonyi - Lesson 5 and lesson 6 update (2015 Sep 16)
* [8efcf05] Krisztian Goedrei - PR fix (2015 Sep 16)
* [4fbdbc7] Viktor Benei - changelog template: curl call "fix" (2015 Sep 16)
* [1a9c807] Krisztian Goedrei - step version logs (2015 Sep 16)
* [e823193] Krisztian Goedrei - version print (2015 Sep 16)
* [cc37c11] Krisztian Goedrei - print fix (2015 Sep 16)
* [5592174] Krisztian Goedrei - start using stepinfo model for print (2015 Sep 16)
* [831fd6f] Krisztian Goedrei - step info model (2015 Sep 16)
* [1845bd5] Tamás Bazsonyi - Lesson 5 WF (2015 Sep 16)
* [0830d33] Viktor Benei - Merge pull request #226 from gkiki90/validate_config (2015 Sep 15)
* [56e8540] Krisztian Goedrei - test fix (2015 Sep 15)
* [a4d0547] Viktor Benei - Merge pull request #228 from gkiki90/trigger_fix (2015 Sep 15)
* [c8a26eb] Krisztian Goedrei - ci fix (2015 Sep 15)
* [ce5b886] Tamás Bazsonyi - Added trigger lesson (2015 Sep 15)
* [93a8e98] Krisztian Goedrei - trigger fix (2015 Sep 15)
* [50cd432] Viktor Benei - Merge pull request #227 from bazscsa/master (2015 Sep 15)
* [9a1da09] Tamás Bazsonyi - links in new line (2015 Sep 15)
* [e6c3fbb] Tamás Bazsonyi - updated links (2015 Sep 15)
* [5e11086] Tamás Bazsonyi - lesson 1 links (2015 Sep 15)
* [030d233] Tamás Bazsonyi - lessons links (2015 Sep 15)
* [dac981f] Tamás Bazsonyi - updated yml (2015 Sep 15)
* [5a8d115] Tamás Bazsonyi - removed readme (2015 Sep 15)
* [21e0df6] Tamás Bazsonyi - removed yml (2015 Sep 15)
* [32d98d3] Tamás Bazsonyi - Added Lessons (2015 Sep 15)
* [ca37f08] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Sep 15)
* [da08152] Viktor Benei - Merge pull request #225 from gkiki90/abc (2015 Sep 14)
* [1b4795d] Krisztian Goedrei - validate inventory (2015 Sep 14)
* [4388a16] Krisztian Goedrei - ci fix (2015 Sep 14)
* [037a0f6] Krisztian Goedrei - ci fix (2015 Sep 14)
* [4c5ce86] Krisztian Goedrei - bitrise.yml fix, test fix (2015 Sep 14)
* [028dc4d] Krisztian Goedrei - fixes (2015 Sep 12)
* [d14a781] Krisztian Goedrei - PR fix (2015 Sep 12)
* [e8fdff8] Krisztian Goedrei - sort (2015 Sep 12)
* [ecefc8a] Tamás Bazsonyi - paths updated (2015 Sep 12)
* [9840a57] Tamás Bazsonyi - path (2015 Sep 12)
* [d9552fa] Tamás Bazsonyi - Lessons README (2015 Sep 12)
* [772bf3e] Viktor Benei - Merge pull request #223 from gkiki90/custom_step (2015 Sep 11)
* [9a0b9b1] Viktor Benei - Merge pull request #224 from viktorbenei/master (2015 Sep 11)
* [168d496] Krisztian Goedrei - normalize fix (2015 Sep 11)
* [90245bd] Viktor Benei - start of v1.0.1 (2015 Sep 11)


## 1.0.0 (2015 Sep 11)

### Release Notes

* __Linux support__ : first official Linux release. No dependency manager support is available for Linux yet, but everything else should work the same as on OS X.
* Improved `bitrise init`, with better guides, `trigger_map` and more!
* Total runtime summary at the end of a build.
* Lots of internal code revision, improved `bitrise normalize`.
* __New command__ : `bitrise validate` to quick-check your `bitrise.yml`.
* Configurations (`bitrise.yml` and `.bitrise.secrets.yml`) can now be specified in `base64` format as well - useful for tools.
* __DEPRECATED__ : the old `--path` flag is now deprecated, in favor of `--config`, which has it's `base64` format (`--config-base64`)
* Logs now include the `step`'s version if it's referenced from a Step Collection. Prints the version even if no version constraint is defined (mainly for debug purposes).
* __NEW__ : sets `BITRISE_SOURCE_DIR` (to current dir) and `BITRISE_DEPLOY_DIR` (to a temp dir) environments if the env is not defined
* Only do a `stepman update` once for a collection if can't find a specified step (version).
* __FIX__ : Custom steps (where the collection is `_`) don't crash anymore because of missing required fields.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/1.0.0/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.11 -> 1.0.0

* [c2c7c04] Viktor Benei - Merge pull request #222 from viktorbenei/master (2015 Sep 11)
* [7ee03c8] Viktor Benei - updated 1.0.0 changelog (2015 Sep 11)
* [a4a1ad4] Viktor Benei - Merge pull request #221 from gkiki90/custom_step (2015 Sep 11)
* [2381bb4] Krisztian Goedrei - fix (2015 Sep 11)
* [bf3f9e2] Krisztian Goedrei - custom step defaults (2015 Sep 11)
* [53270bb] Tamás Bazsonyi - Workflows (2015 Sep 10)
* [164877f] Tamás Bazsonyi - Added step yml (2015 Sep 10)
* [0152403] Viktor Benei - Merge pull request #220 from gkiki90/pointers (2015 Sep 10)
* [3471df4] Krisztian Goedrei - pointer fixes (2015 Sep 10)
* [ead8159] Tamás Bazsonyi - removed <> (2015 Sep 10)
* [5ac8f5a] Tamás Bazsonyi - corrected format (2015 Sep 10)
* [88c0d5d] Tamás Bazsonyi - Added lesson 1 (2015 Sep 10)
* [5739093] Viktor Benei - Merge pull request #219 from gkiki90/deploy_dir (2015 Sep 10)
* [c52bd9a] Krisztian Goedrei - PR fix (2015 Sep 10)
* [8410daf] Krisztian Goedrei - PR fix (2015 Sep 10)
* [ba2f6b8] Krisztian Goedrei - deploy dir (2015 Sep 10)
* [4e71a04] Viktor Benei - Merge pull request #218 from viktorbenei/master (2015 Sep 09)
* [58b739d] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 09)
* [464f0e7] Viktor Benei - Merge pull request #217 from gkiki90/bitrise_src_dir (2015 Sep 09)
* [3912ffd] Krisztian Goedrei - BITRISE_SOURCE_DIR handling & test (2015 Sep 09)
* [ab20912] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 09)
* [118c543] Viktor Benei - BITRISE_PROJECT_TITLE renamed in `init` to BITRISE_APP_TITLE - to match the bitrise.io one (2015 Sep 09)
* [8869945] Viktor Benei - Merge pull request #216 from gkiki90/print_fix (2015 Sep 09)
* [f449dfa] Krisztian Goedrei - print tests (2015 Sep 09)
* [210ef4f] Tamás Bazsonyi - added initial readmes (2015 Sep 08)
* [7452e53] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 08)
* [97bdeaa] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 08)
* [8920f8a] Viktor Benei - Merge pull request #215 from viktorbenei/master (2015 Sep 08)
* [5b9b36c] Viktor Benei - godeps-update (2015 Sep 08)
* [9f034b9] Viktor Benei - step log version printing fix - trimming version string. Mainly affects the steps which are not used from a steplib (2015 Sep 08)
* [caac385] Viktor Benei - Merge pull request #213 from gkiki90/step_version_log_fix (2015 Sep 08)
* [cb9d4d2] Krisztian Goedrei - fix (2015 Sep 08)
* [cddeda9] Krisztian Goedrei - test (2015 Sep 08)
* [28673aa] Viktor Benei - Merge pull request #212 from viktorbenei/master (2015 Sep 08)
* [19da3cc] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 08)
* [a621c28] Viktor Benei - Merge pull request #211 from gkiki90/trigger_fix (2015 Sep 08)
* [2b81fda] Viktor Benei - godeps-update (2015 Sep 08)
* [2a4e93b] Viktor Benei - required stepman version bump (2015 Sep 08)
* [d3e9f76] Viktor Benei - full godeps-update (2015 Sep 08)
* [9b14e35] Krisztian Goedrei - PR fix (2015 Sep 08)
* [eabab55] Krisztian Goedrei - fix (2015 Sep 08)
* [bfde226] Viktor Benei - Merge pull request #210 from viktorbenei/master (2015 Sep 08)
* [d20d08a] Viktor Benei - godeps-update : CopyDir fix & stepman model property order change (2015 Sep 08)
* [ebcba2b] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 08)
* [d0763cb] Viktor Benei - _test/bitrise.yml step title fix (2015 Sep 08)
* [a378ab9] Viktor Benei - Merge pull request #209 from viktorbenei/master (2015 Sep 07)
* [ec5c61c] Viktor Benei - base trigger_map added to bitrise.yml, for CI (2015 Sep 07)
* [7773df7] Viktor Benei - step version printing note added to changelog (2015 Sep 07)
* [909ef5b] Viktor Benei - changelog - version fix (1.0.0) (2015 Sep 07)
* [ab1cefd] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 07)
* [dddf2e5] Viktor Benei - bit more explanation for setup --minimal in CI (2015 Sep 07)
* [f836cfa] Viktor Benei - Merge pull request #208 from viktorbenei/master (2015 Sep 07)
* [8f62ca5] Viktor Benei - minimal refactoring for CI (2015 Sep 07)
* [6fb55a4] Viktor Benei - tmp build fix for CI (2015 Sep 07)
* [359b8af] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 07)
* [44db09f] Viktor Benei - run a minimal setup at start of CI (2015 Sep 07)
* [8f42f56] Viktor Benei - Merge pull request #206 from gkiki90/step_version (2015 Sep 07)
* [017f05c] Krisztian Goedrei - PR fix (+2 squashed commits) Squashed commits: [9e44a47] fix [6ca52f8] step version (2015 Sep 07)
* [c525a42] Viktor Benei - Merge pull request #207 from viktorbenei/master (2015 Sep 07)
* [978ca1d] Viktor Benei - Linux ready release configuration; v1.0.0 changelog (2015 Sep 07)
* [ea9a59f] Viktor Benei - skip `brew` dependencies if platform is Linux (2015 Sep 07)
* [318119f] Viktor Benei - step-template update (2015 Sep 07)
* [9e55f89] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 07)
* [2017941] Viktor Benei - Merge pull request #205 from gkiki90/step_template (2015 Sep 07)
* [3d83f59] Viktor Benei - `bitrise init` now embeds the models.Version instead of a fixed 1.0.0; init now uses the new config 'title, summary, description' instead of YML comments (2015 Sep 07)
* [35d17af] Krisztian Goedrei - readme (2015 Sep 07)
* [f308373] Viktor Benei - Title, Summary, Description added to AppModel (main config model) & reordered the three, to be in this order in every model. (2015 Sep 07)
* [8a63a15] Viktor Benei - Merge pull request #202 from gkiki90/normalize_fix (2015 Sep 07)
* [8ed3179] Viktor Benei - Merge pull request #201 from gkiki90/step_template (2015 Sep 07)
* [bdd942b] Viktor Benei - Merge pull request #204 from gkiki90/util_test (2015 Sep 07)
* [6ebe2c2] Krisztian Goedrei - PR fix (2015 Sep 07)
* [b7d53d1] Viktor Benei - Merge pull request #203 from gkiki90/total_runtime (2015 Sep 07)
* [76fafc1] Krisztian Goedrei - PR fix (2015 Sep 07)
* [e4a39eb] Krisztian Goedrei - test (2015 Sep 07)
* [2ef20a8] Krisztian Goedrei - slice tests (2015 Sep 07)
* [eb7ffe7] Krisztian Goedrei - total runtime (2015 Sep 06)
* [4f1705b] Krisztian Goedrei - normalize fix (2015 Sep 05)
* [364df10] Krisztian Goedrei - missing fields (2015 Sep 05)
* [e2c2652] Viktor Benei - Merge pull request #200 from viktorbenei/master (2015 Sep 05)
* [31a7be0] Viktor Benei - updated Dockerfile & bitrise.yml for building `bitrise` in Docker, using the `bitrise.yml` (2015 Sep 04)
* [397838a] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 04)
* [40c3996] Viktor Benei - Merge pull request #199 from gkiki90/normalize_fix (2015 Sep 04)
* [432e3da] Viktor Benei - Merge pull request #198 from gkiki90/base64 (2015 Sep 04)
* [f22ff6b] Krisztian Goedrei - fix (2015 Sep 04)
* [9082415] Krisztian Goedrei - PR fix (2015 Sep 04)
* [d6ec4a1] Krisztian Goedrei - fix (2015 Sep 04)
* [2046c36] Viktor Benei - upload&download bitrise.yml : ensure-clean-git & create backup (2015 Sep 04)
* [e7cd7c5] Viktor Benei - experimental : upload & download bitrise.yml to/from bitrise.io (2015 Sep 04)
* [e9fe439] Krisztian Goedrei - PR fix (2015 Sep 04)
* [d4b9359] Krisztian Goedrei - PR fix (2015 Sep 04)
* [f9bd982] Krisztian Goedrei - test (2015 Sep 04)
* [937de7d] Krisztian Goedrei - base64 (2015 Sep 04)
* [3b6b4e6] Viktor Benei - Merge pull request #197 from viktorbenei/master (2015 Sep 04)
* [aba4989] Viktor Benei - bitrise.yml cleanup & format version update (2015 Sep 04)
* [2e42496] Viktor Benei - more thorough template expression simple "true/false" tests (2015 Sep 04)
* [7b472fa] Viktor Benei - Merge pull request #195 from viktorbenei/feature/trigger-map-in-init (2015 Sep 03)
* [c70451f] Viktor Benei - just a little bit more test for the `init` content (2015 Sep 03)
* [ea4c0af] Viktor Benei - godeps-update + a fix for recursive `godep save` (2015 Sep 03)
* [a43074b] Viktor Benei - annotated, and formatted `bitrise.yml` after init, with `trigger_map`, test, and a bit of info about `bitrise trigger` (2015 Sep 03)
* [128ead0] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Sep 03)
* [631bf0f] Viktor Benei - Merge pull request #194 from gkiki90/validate_config (2015 Sep 03)
* [90ade1d] Krisztian Goedrei - removed alias (2015 Sep 03)
* [1da63ba] Krisztian Goedrei - validate (2015 Sep 03)
* [b9f9593] Viktor Benei - base Linux setup/support (2015 Sep 03)
* [cf0a83c] Viktor Benei - Merge pull request #192 from gkiki90/normalize_fix (2015 Sep 03)
* [c06de4e] Krisztian Goedrei - PR fix (2015 Sep 03)
* [990f72d] Krisztian Goedrei - code cleaning (2015 Sep 03)
* [876c89c] Krisztian Goedrei - PR fix (2015 Sep 03)
* [77a3797] Viktor Benei - Merge pull request #193 from viktorbenei/master (2015 Sep 03)
* [c566b29] Krisztian Goedrei - pointer fix (2015 Sep 03)
* [96bfe66] Krisztian Goedrei - normalize fix (2015 Sep 03)
* [30e60d7] Viktor Benei - test workflow extended with a fail test & trigger_map (2015 Sep 02)
* [78b2c2a] Viktor Benei - Merge pull request #191 from viktorbenei/master (2015 Sep 02)
* [db3f4df] Viktor Benei - fail if `golint` finds any issue (2015 Sep 02)
* [d5ddee7] Viktor Benei - extended _tests/bitrise.yml (2015 Sep 02)
* [ddc8666] Viktor Benei - Merge pull request #190 from viktorbenei/master (2015 Aug 31)
* [58da26f] Viktor Benei - updated bitrise-cli install (2015 Aug 31)
* [22b32bc] Viktor Benei - start of v0.9.12 (2015 Aug 31)


## 0.9.11 (2015 Aug 31)

### Release Notes

* __NEW__ : `bitrise.yml` can now be exported into JSON (with `bitrise export`), and `.json` configuration is also acceptable now for a `bitrise run`.
* __NEW__ / __BREAKING__ : workflow names which start with an underscore (ex: _my_wf) are now treated as "utility" workflow, which can only be triggered by another workflow (as a `before_run` or `after_run` workflow). These "utility" workflows will only be listed by a `bitrise run` call as another section (utility workflows), to provide a better way to organize workflows which are not intended to be called directly.
* __FIX__ : Input environments handling fix: Step inputs are now isolated, one step's input won't affect another's with the same environment key
* __NEW__ : The workflow which was triggered by `bitrise run WORKFLOW-NAME` is now available as an environment variable
    * `BITRISE_TRIGGERED_WORKFLOW_ID` : contains the ID of the workflow
    * `BITRISE_TRIGGERED_WORKFLOW_TITLE` : contains the `title` of the workflow, if specified
* __NEW__ : `BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH` is now also defined, as a temporary file path.
* __NEW__ : `bitrise normalize` command, to help you "normalize" your `bitrise.yml`.
* __NEW__ : `trigger_map` definition and `bitrise trigger` action : with this you can map expressions to workflows. A common use case for this is to map branch names (ex: `feature/xyz`) to workflows, simply by defining the mapping in the `bitrise.yml`.
* Log format revision, to make it more obvious where a Step starts and ends, and at the end of the build it provides a much improved summary.
* A new "StepLib" source type (`_`), to provide compatibility with Steps which don't have an up-to-date `step.yml` in the Step's repository. Effectively the same as `git::http://step/url.git@version`, but it won't check for a `step.yml` at all - which means that every information have to be included in the `bitrise.yml`.
* Every configuration level (environments, step, step inputs, ...) which had at least a `title` or a `description` or `summary` now has all three: `title`, `summary` and `description`.
* Other internal revisions and minor fixes, and __lots__ of test added.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.11/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.10 -> 0.9.11

* [217e649] Viktor Benei - Merge pull request #189 from viktorbenei/master (2015 Aug 31)
* [53a2eb6] Viktor Benei - bitrise.yml revision : updated `test` handling (2015 Aug 31)
* [46490e7] Viktor Benei - changelog addition (2015 Aug 31)
* [c5ea28f] Viktor Benei - godeps-update (2015 Aug 31)
* [f5bd2ed] Viktor Benei - Merge pull request #188 from gkiki90/model_version (2015 Aug 31)
* [cb62ebb] Krisztian Goedrei - models version (2015 Aug 31)
* [badbdf4] Krisztian Goedrei - model version (2015 Aug 31)
* [870309c] Viktor Benei - Merge pull request #187 from gkiki90/published_at_fix (2015 Aug 31)
* [6244d9f] Krisztian Goedrei - PR fix (2015 Aug 31)
* [8493427] Krisztian Goedrei - PR fix (2015 Aug 31)
* [2bab801] Krisztian Goedrei - godeps-update (2015 Aug 31)
* [3b13c39] Krisztian Goedrei - published_at type fix (2015 Aug 31)
* [742321a] Viktor Benei - Merge pull request #186 from gkiki90/1_0_0_models (2015 Aug 31)
* [f10a1bb] Krisztian Goedrei - merge (2015 Aug 31)
* [dd63058] Krisztian Goedrei - # This is a combination of 2 commits. # The first commit's message is: (2015 Aug 31)
* [fae8b05] Viktor Benei - Merge pull request #185 from gkiki90/last_step_fix (2015 Aug 31)
* [8555f5d] Krisztian Goedrei - PR fix (2015 Aug 31)
* [64f78a9] Krisztian Goedrei - trigger workflow & last step fix (2015 Aug 31)
* [c2c8d94] Viktor Benei - Merge pull request #183 from gkiki90/run_summary (2015 Aug 28)
* [e2c65bf] Krisztian Goedrei - removed log (2015 Aug 28)
* [6197e11] Krisztian Goedrei - print run summary (2015 Aug 28)
* [5766950] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Aug 27)
* [1195f37] Viktor Benei - Merge pull request #182 from gkiki90/ci (2015 Aug 27)
* [0506d05] Krisztian Goedrei - ci fix (2015 Aug 27)
* [38dbd44] Krisztian Goedrei - ci (2015 Aug 27)
* [f54b01c] Viktor Benei - Merge pull request #181 from gkiki90/master (2015 Aug 27)
* [bf0cf88] Krisztian Goedrei - bypass checking for a TTY before outputting colors (2015 Aug 27)
* [9bac0a9] Viktor Benei - Merge pull request #180 from viktorbenei/master (2015 Aug 26)
* [d526963] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 26)
* [b85fbc5] Viktor Benei - changelog update (2015 Aug 26)
* [a129cbb] Viktor Benei - Merge pull request #179 from viktorbenei/master (2015 Aug 26)
* [6ad9602] Viktor Benei - step run summary log box revision & a bit longer fail-test in bitrise.yml (2015 Aug 26)
* [25b0e67] Viktor Benei - Merge pull request #178 from gkiki90/remove_defaults (2015 Aug 26)
* [a9a811e] Krisztian Goedrei - code style, tests (2015 Aug 26)
* [a1839c5] Krisztian Goedrei - remove defaults, fill step outputs (2015 Aug 26)
* [f10ada6] Viktor Benei - Merge pull request #177 from gkiki90/init_fix (2015 Aug 26)
* [93db1a8] Tamás Bazsonyi - Added local app install sample (2015 Aug 25)
* [44571e7] Krisztian Goedrei - init fix (2015 Aug 25)
* [90a9d83] Viktor Benei - Merge pull request #176 from viktorbenei/master (2015 Aug 24)
* [dc7f135] Viktor Benei - Merge pull request #175 from bazscsa/master (2015 Aug 24)
* [a8cad55] Viktor Benei - doRun_test major revision: most of the tests now use `runWorkflowWithConfiguration` to run the test, which is much closer to how a full `bitrise run` happens (2015 Aug 24)
* [69562d7] Tamás Bazsonyi - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Aug 24)
* [9eb00ae] Tamás Bazsonyi - Some rephrasing (2015 Aug 24)
* [589c1b3] Tamás Bazsonyi - Listified the Documentation overview (2015 Aug 24)
* [3cacdfd] Tamás Bazsonyi - Added _docs to README.md (2015 Aug 24)
* [84ce3ef] Tamás Bazsonyi - Added README.md to the docs folder (2015 Aug 24)
* [9b83f6e] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 24)
* [7c0af85] Viktor Benei - output env test & some logging text fix (workflow 'ID' instead of 'title') (2015 Aug 24)
* [39b1b12] Viktor Benei - Merge pull request #174 from gkiki90/separate_run (2015 Aug 24)
* [a5798c2] Krisztian Goedrei - separated run (2015 Aug 24)
* [e6536c5] Viktor Benei - removed unnecessary 'unload' from react-native example (2015 Aug 24)
* [c802db5] Viktor Benei - changelog v0.9.11 (2015 Aug 24)
* [f995a49] Viktor Benei - Merge pull request #172 from gkiki90/run_old_steps (2015 Aug 24)
* [eea413e] Viktor Benei - Merge pull request #173 from viktorbenei/master (2015 Aug 24)
* [090005b] Krisztian Goedrei - typo (2015 Aug 24)
* [250b197] Viktor Benei - AppendEnvironmentSlice replaced with the built in "append" method (2015 Aug 24)
* [433a48a] Viktor Benei - godeps-update & envman and stepman min version bump (2015 Aug 24)
* [4b6aeb5] Krisztian Goedrei - code style (2015 Aug 24)
* [231a151] Krisztian Goedrei - code style (2015 Aug 24)
* [9349b1d] Krisztian Goedrei - code style (2015 Aug 24)
* [5a5be3d] Krisztian Goedrei - run old steps (2015 Aug 24)
* [757f4f1] Viktor Benei - Merge pull request #171 from gkiki90/title_summary_desc (2015 Aug 19)
* [989ff21] Krisztian Goedrei - title, summary, description (2015 Aug 19)
* [c4da2c3] Krisztian Goedrei - Merge branch 'master' of github.com:bitrise-io/bitrise-cli (2015 Aug 19)
* [237eea0] Viktor Benei - Merge pull request #170 from gkiki90/step_working_dir (2015 Aug 19)
* [378a1c6] Krisztian Goedrei - Merge branch 'step_working_dir' (2015 Aug 19)
* [41c3069] Krisztian Goedrei - step working dir (2015 Aug 19)
* [a799972] Viktor Benei - Merge pull request #169 from viktorbenei/master (2015 Aug 18)
* [6670815] Viktor Benei - Slack examples version update: from 2.0.0 to 2.1.0 (2015 Aug 18)
* [d5faf3d] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 18)
* [91c9185] Viktor Benei - added RunIf ENV template tests (2015 Aug 18)
* [46446f2] Viktor Benei - Merge pull request #168 from gkiki90/go-utils (2015 Aug 18)
* [6cc2f51] Viktor Benei - Merge pull request #167 from gkiki90/workflow_title (2015 Aug 18)
* [8ef0281] Krisztian Goedrei - missing go-utils methods (2015 Aug 18)
* [3c8f325] Krisztian Goedrei - workflow title (2015 Aug 18)
* [77f14e4] Viktor Benei - Merge pull request #166 from gkiki90/step_input_fix (2015 Aug 18)
* [fc62a83] Krisztian Goedrei - environment handling (2015 Aug 18)
* [4272084] Viktor Benei - Merge pull request #165 from viktorbenei/master (2015 Aug 17)
* [17ef411] Viktor Benei - minor text change (2015 Aug 17)
* [df5ba72] Viktor Benei - dependency manager "OK" message - unified (2015 Aug 17)
* [f947c2e] Viktor Benei - Merge pull request #164 from viktorbenei/master (2015 Aug 17)
* [c4da8b7] Viktor Benei - slack step update to the new v2.0.0 version (2015 Aug 17)
* [5e1a2c2] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 17)
* [4845209] Viktor Benei - added "dependencies" to step-template (2015 Aug 17)
* [bf9533e] Viktor Benei - Merge pull request #163 from bazscsa/master (2015 Aug 17)
* [229c201] Tamás Bazsonyi - Added brew update (2015 Aug 17)
* [1674a3e] Tamás Bazsonyi - Some grammar corrections (2015 Aug 17)
* [278d29f] Tamás Bazsonyi - revisions (2015 Aug 17)
* [c037187] Tamás Bazsonyi - Added Share Guide (2015 Aug 17)
* [22ae699] Tamás Bazsonyi - Added React Native (2015 Aug 17)
* [17a7b1c] Tamás Bazsonyi - Added CLI share guide (2015 Aug 17)
* [796e8af] Tamás Bazsonyi - Added CLI introduction (2015 Aug 17)
* [02a0de9] Tamás Bazsonyi - How to guide (2015 Aug 17)
* [f73a40b] Viktor Benei - Merge pull request #162 from viktorbenei/feature/export_command (2015 Aug 17)
* [03eba50] Viktor Benei - godep-update for a required envman model fix (2015 Aug 17)
* [a00fd3e] Viktor Benei - export command : export a bitrise config file in either YAML or JSON format, with optional pretty printed JSON (2015 Aug 17)
* [9e52a55] Viktor Benei - Merge pull request #161 from viktorbenei/feature/predefined_envs (2015 Aug 17)
* [f633487] Viktor Benei - set predefined ENVs, so far only one: BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH (2015 Aug 17)
* [686b5a5] Viktor Benei - Update README.md (2015 Aug 14)
* [c02a212] Viktor Benei - Merge pull request #160 from viktorbenei/master (2015 Aug 14)
* [43d7c05] Viktor Benei - README revision (2015 Aug 14)
* [2fffdb7] Viktor Benei - Merge pull request #159 from viktorbenei/master (2015 Aug 14)
* [8039c00] Viktor Benei - removed the now obsolete reference to `brew update` in setup's `--minimal` flag (2015 Aug 14)
* [775a146] Viktor Benei - Merge pull request #158 from viktorbenei/master (2015 Aug 14)
* [6d396da] Viktor Benei - switch to bitrise 0.9.10 for CI (2015 Aug 14)
* [77da03c] Viktor Benei - start of v0.9.11 (2015 Aug 14)
* [0d2e718] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 14)


## 0.9.10 (2015 Aug 14)

### Release Notes

* Improved `setup` : it has a new `--minimal` flag to skip more advanced setup checks, like the `brew doctor` call.
* Removed `brew update` completely from the `setup`.
* Step dependencies: before installing a dependency with `brew` bitrise now specifically asks for permission to do so, except in `--ci` mode.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.10/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.9 -> 0.9.10

* [59bd349] Viktor Benei - Merge pull request #157 from viktorbenei/master (2015 Aug 14)
* [37f36fb] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 14)
* [08f1d6a] Viktor Benei - v0.9.10 changelog (2015 Aug 14)
* [59ffe92] Viktor Benei - Merge pull request #156 from viktorbenei/master (2015 Aug 14)
* [ba886cb] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 14)
* [18f2bfd] Viktor Benei - Merge pull request #155 from viktorbenei/master (2015 Aug 14)
* [29abf82] Viktor Benei - godeps-update (2015 Aug 14)
* [6f2fa64] Viktor Benei - stepman min version bump (2015 Aug 14)
* [4cdde7f] Viktor Benei - Merge pull request #154 from viktorbenei/master (2015 Aug 14)
* [3745167] Viktor Benei - _step_template and it's CI test moved from the stepman project to this repo (2015 Aug 14)
* [64f6f4a] Viktor Benei - Merge pull request #153 from viktorbenei/master (2015 Aug 14)
* [a8ba919] Viktor Benei - prepare for bitrise setup --minimal (2015 Aug 14)
* [a940056] Viktor Benei - Merge pull request #152 from viktorbenei/master (2015 Aug 14)
* [7f0c2e4] Viktor Benei - install bitrise cli for ci script (2015 Aug 14)
* [99b365b] Viktor Benei - every `brew doctor` issue counts - use the `--minimal` flag to skip `brew doctor` (2015 Aug 14)
* [2663282] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 14)
* [9dc16b3] Viktor Benei - don't do `brew update` in setup (2015 Aug 14)
* [c1113d4] Viktor Benei - Merge pull request #151 from gkiki90/dependency_fix (2015 Aug 14)
* [7932300] Krisztian Goedrei - dependency fixes (2015 Aug 14)
* [a6f7827] Viktor Benei - Merge pull request #150 from viktorbenei/master (2015 Aug 14)
* [f3a91dd] Viktor Benei - minimum envman version bump (2015 Aug 14)
* [f2e4caf] Viktor Benei - Merge pull request #149 from viktorbenei/master (2015 Aug 14)
* [0bb4e1d] Viktor Benei - minimal setup mode : skips brew update and brew doctor (2015 Aug 14)
* [6ee580f] Viktor Benei - Merge pull request #148 from viktorbenei/master (2015 Aug 14)
* [da000b2] Viktor Benei - godeps-update (2015 Aug 14)
* [b840881] Viktor Benei - Merge pull request #146 from gkiki90/master (2015 Aug 13)
* [f75b685] Viktor Benei - Merge pull request #147 from viktorbenei/master (2015 Aug 13)
* [6fb4c64] Viktor Benei - stepman dependency bump (2015 Aug 13)
* [63c6059] Krisztian Goedrei - cli fixes (2015 Aug 13)
* [fed4916] Krisztian Goedrei - godep-update (2015 Aug 13)
* [148f29a] Krisztian Goedrei - go-util update (2015 Aug 13)
* [b8de012] Viktor Benei - Merge pull request #145 from gkiki90/workflow_fixes (2015 Aug 13)
* [cb7e3c5] Viktor Benei - Merge pull request #144 from viktorbenei/master (2015 Aug 13)
* [6689c24] Viktor Benei - start of v0.9.10 (2015 Aug 13)


## 0.9.9 (2015 Aug 13)

### Release Notes

* `bitrise setup` revision : better `brew` checking (calls `brew update` and `brew doctor` too) but no direct Command Line Tools checking.
    * The previous solution was incompatible with OS X Mountain Lion and earlier versions, this version solves this incompatibility.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.9/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.8 -> 0.9.9

* [a832bd5] Viktor Benei - Merge pull request #143 from viktorbenei/master (2015 Aug 13)
* [22b4e71] Viktor Benei - 0.9.9 changelog (2015 Aug 13)
* [2ae0825] Viktor Benei - Merge pull request #142 from viktorbenei/master (2015 Aug 13)
* [fc44dc0] Viktor Benei - Xcode CLT is not an explicit dependency anymore, only brew; but brew check extended with brew update and brew doctor (2015 Aug 13)
* [8bd8098] Viktor Benei - Merge pull request #141 from mistydemeo/xcode-select (2015 Aug 13)
* [6dd0fe9] Misty De Meo - Dependencies: fix xcode-select argument (2015 Aug 12)
* [2f72f83] Viktor Benei - Merge pull request #140 from viktorbenei/master (2015 Aug 12)
* [3cefbbb] Viktor Benei - start of v0.9.9 (2015 Aug 12)
* [662012a] Viktor Benei - changelog (2015 Aug 12)


## 0.9.8 (2015 Aug 12)

### Release Notes

* __BREAKING__ : `step.yml` shared in Step Libraries / Step Collections now have to include a `commit` (hash) property inside the `source` property, for better version validation (version tag have to match this commit hash)!
    * You should switch to the new, final default StepLib, hosted on GitHub, which contains these commit hashes and works with stepman 0.9.8! URL: https://github.com/bitrise-io/bitrise-steplib
    * We'll soon (in about 1 day) start to accept Step contributions to this new StepLib!
    * You should replace the previous `https://bitbucket.org/bitrise-team/bitrise-new-steps-spec` `default_step_lib_source` and every other reference to this old (now deprecated) StepLib, and **replace it** with `https://github.com/bitrise-io/bitrise-steplib.git`!
* __BUGFIX__ : the `$STEPLIB_BUILD_STATUS` and `$BITRISE_BUILD_STATUS` environments were not set correctly in the previous version for a couple of multi-workflow setups.
* __NEW__ : `bitrise init` now automatically adds `.bitrise*` to the `.gitignore` file in the current folder, to prevent accidentally sharing your `.bitrise.secrets.yml` or other bitrise generated temporary files/folders.
* __NEW__ : built in commands to `share` a new step into a StepLib - through `stepman`.
* __NEW__ : `run_if` expressions can now use the new `.IsPR` check, to declare whether a given step should run in case of a Pull Request build.
* __NEW__ : Step dependencies : `Xcode` can now be specified as a dependency for steps. Unfortunately it can't be installed automatically, but you'll get proper error message about the missing full Xcode in this case, rather than a generic error message during running the step.
* __NEW__ : bitrise now checks the `format_version` of the `bitrise.yml` file and doesn't run it if it was created for a newer version.
* You no longer have to call `setup` after the installation or upgrade of `bitrise`, it'll automatically check whether `setup` was called (and succeeded) when you call `run`.
* Bitrise now creates it's temporary working cache dir in a System temp folder, instead of spamming the current directory with a `.bitrise` folder at every `bitrise run`.
* Improved `bitrise run` logs.
* LOTS of code revision

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.8/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.7 -> 0.9.8

* [e03719c] Viktor Benei - Merge pull request #139 from viktorbenei/master (2015 Aug 12)
* [f1e0a95] Viktor Benei - minimum envman and stepman version bump (2015 Aug 12)
* [f88e7da] Krisztian Goedrei - workflow fixes (2015 Aug 12)
* [e819d32] Viktor Benei - Merge pull request #138 from viktorbenei/master (2015 Aug 12)
* [84d15d2] Viktor Benei - godeps-update (2015 Aug 12)
* [9e72f25] Viktor Benei - Merge pull request #137 from viktorbenei/feature/debug-env (2015 Aug 12)
* [015089d] Viktor Benei - bit of ENV revision in general, and a new "--debug" flag (or DEBUG=true ENV) to run in Debug Mode (2015 Aug 12)
* [91c1b5c] Viktor Benei - Merge pull request #136 from viktorbenei/feature/pr-env (2015 Aug 12)
* [4fe4dba] Viktor Benei - PULL_REQUEST_ID env related template-expressions: .IsPR can now be used as Run-If expression (2015 Aug 12)
* [73e69cf] Viktor Benei - Merge pull request #135 from viktorbenei/feature/xcode-dependency (2015 Aug 11)
* [a91e67f] Viktor Benei - special "try check" dependency type, which can't be installed - a special one is 'xcode', which has it's own error msg (2015 Aug 11)
* [b33450e] Viktor Benei - Merge pull request #134 from gkiki90/master (2015 Aug 11)
* [d1742dd] Krisztian Goedrei - step title check (2015 Aug 11)
* [76a7d7a] Viktor Benei - typo fix (2015 Aug 11)
* [f8be2a4] Viktor Benei - Merge pull request #133 from gkiki90/go-utils (2015 Aug 11)
* [f98128d] Krisztian Goedrei - updated to go-utils (2015 Aug 11)
* [1e14821] Viktor Benei - Merge pull request #132 from viktorbenei/feature/check-workflow-version (2015 Aug 11)
* [faafea3] Viktor Benei - bitrise.yml format version check (2015 Aug 11)
* [55848a3] Krisztian Goedrei - RemoveFile instead of RemoveDir (2015 Aug 11)
* [946033e] Viktor Benei - Merge pull request #130 from gkiki90/working_directory (2015 Aug 11)
* [83aabab] Krisztian Goedrei - os temp workdir (2015 Aug 11)
* [8a5c48b] Viktor Benei - Merge pull request #127 from gkiki90/define_configs_as_string (2015 Aug 11)
* [e51c416] Krisztian Goedrei - removed test.yml (2015 Aug 11)
* [a36c4ab] Krisztian Goedrei - tests (2015 Aug 11)
* [7ab56cd] Viktor Benei - Merge pull request #128 from viktorbenei/master (2015 Aug 11)
* [fee03a8] Viktor Benei - godep-update (2015 Aug 11)
* [babb0c7] Viktor Benei - go-utils migration (2015 Aug 11)
* [6c3778d] Viktor Benei - Merge pull request #126 from gkiki90/refactor (2015 Aug 11)
* [d9d329a] Krisztian Goedrei - failed not important -> skippable (2015 Aug 11)
* [ec81f42] Viktor Benei - Merge pull request #125 from gkiki90/build_failed_fix (2015 Aug 10)
* [1f4b840] Krisztian Goedrei - build status env test fix (+3 squashed commits) Squashed commits: [3d3392c] build status [e77352e] build status env tests [9d9f504] var to const (+3 squashed commits) Squashed commits: [a60406a] godep-update [0f48b5f] run tests [8223985] do run step status fix (2015 Aug 10)
* [7df8b74] Viktor Benei - Merge pull request #123 from viktorbenei/feature/print-version-under-ascii-header (2015 Aug 10)
* [f21571f] Viktor Benei - print the version number under the ASCII header (2015 Aug 10)
* [d9ceb9e] Viktor Benei - Merge pull request #124 from gkiki90/build_failed_fix (2015 Aug 10)
* [9d9f504] Krisztian Goedrei - var to const (+3 squashed commits) Squashed commits: [a60406a] godep-update [0f48b5f] run tests [8223985] do run step status fix (2015 Aug 10)
* [ce2c9f1] Viktor Benei - Merge pull request #122 from viktorbenei/master (2015 Aug 09)
* [6cc9dbd] Viktor Benei - separate slack from-name for CI OK and error (2015 Aug 09)
* [c2cc3e7] Viktor Benei - Merge pull request #120 from viktorbenei/feature/init_add_items_to_gitignore (2015 Aug 09)
* [fed202a] Viktor Benei - Merge pull request #121 from viktorbenei/feature/build_status_env_fix (2015 Aug 09)
* [44fa69f] Viktor Benei - PR fix (2015 Aug 09)
* [e691054] Viktor Benei - fixed (2015 Aug 09)
* [cd81303] Viktor Benei - doInit : append the '.bitrise*' pattern to the .gitignore file in the current dir + example workflow renamed to 'test' (2015 Aug 09)
* [bde1c0a] Viktor Benei - Merge pull request #119 from viktorbenei/feature/include_build_url_in_ci_test_slack_msg (2015 Aug 09)
* [b6048b8] Viktor Benei - Merge pull request #118 from viktorbenei/master (2015 Aug 09)
* [7dfd358] Viktor Benei - Build URL added to CI slack msgs (2015 Aug 09)
* [ea67acf] Viktor Benei - include branch name in CI msg (2015 Aug 09)
* [1806749] Viktor Benei - Merge pull request #117 from viktorbenei/master (2015 Aug 09)
* [8663044] Viktor Benei - test added back + error text typo (2015 Aug 09)
* [8184360] Viktor Benei - bitrise.yml updated - 'ci' workflow added (2015 Aug 09)
* [4e9ff9c] Viktor Benei - color strings updated - white color was removed because it was invisible on white background (2015 Aug 09)
* [bee0e0f] Viktor Benei - Merge pull request #116 from viktorbenei/master (2015 Aug 09)
* [5a21b33] Viktor Benei - moved a couple of things into the new go-utils repo (2015 Aug 08)
* [2f52692] Viktor Benei - colorstring package moved to go-utils repo (2015 Aug 08)
* [4cc0d05] Viktor Benei - Merge pull request #115 from viktorbenei/master (2015 Aug 08)
* [f2bb4a6] Viktor Benei - godeps update (2015 Aug 08)
* [f2dd7ed] Viktor Benei - godeps update (2015 Aug 08)
* [bfa49f2] Viktor Benei - timestamp ten step revision / ref fix (2015 Aug 08)
* [d9985ad] Viktor Benei - added '.git' to the end of the default step lib source (https://github.com/bitrise-io/bitrise-steplib.git), for clarity (2015 Aug 08)
* [74dedb1] Viktor Benei - Merge pull request #114 from gkiki90/master (2015 Aug 08)
* [d4c4449] Krisztian Goedrei - Merge branch 'master' of github.com:bitrise-io/bitrise-cli (2015 Aug 08)
* [576d095] Viktor Benei - Merge pull request #113 from viktorbenei/master (2015 Aug 08)
* [0cd5690] Krisztian Goedrei - godep-update (2015 Aug 08)
* [f3ba246] Viktor Benei - godeps update (+1 squashed commit) Squashed commits: [e9ed7e8] util fix (2015 Aug 08)
* [7f3dfee] Krisztian Goedrei - run print fix (2015 Aug 08)
* [151f879] Viktor Benei - Merge pull request #112 from viktorbenei/master (2015 Aug 08)
* [e4f2052] Viktor Benei - print fix (2015 Aug 08)
* [cb5ecf8] Viktor Benei - Merge pull request #111 from gkiki90/master (2015 Aug 08)
* [0c048a4] Krisztian Goedrei - math fix (2015 Aug 08)
* [9e9fefb] Viktor Benei - Merge pull request #110 from viktorbenei/master (2015 Aug 08)
* [c877ca8] Viktor Benei - godep update (2015 Aug 08)
* [95cf4e5] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 08)
* [b2b3c48] Viktor Benei - Merge pull request #109 from gkiki90/master (2015 Aug 08)
* [a4d1e07] Krisztian Goedrei - godep update (2015 Aug 08)
* [1bf5e85] Krisztian Goedrei - stepman migration fix (2015 Aug 08)
* [94f4de3] Krisztian Goedrei - Merge branch 'master' of github.com:bitrise-io/bitrise-cli (2015 Aug 08)
* [12be6ca] Viktor Benei - godep update (2015 Aug 08)
* [51389eb] Viktor Benei - brew_test mL removed but a new brew_publish one was added (2015 Aug 08)
* [748ed98] Viktor Benei - updated main / default step lib spec repo url (2015 Aug 08)
* [8f327d9] Krisztian Goedrei - stepman migration (2015 Aug 08)
* [a0f4c43] Viktor Benei - Merge pull request #108 from gkiki90/master (2015 Aug 06)
* [a5182d1] Krisztian Goedrei - reference cycle test config moved to _tests (2015 Aug 06)
* [605fc51] Viktor Benei - Merge pull request #107 from viktorbenei/master (2015 Aug 06)
* [455d9c9] Viktor Benei - Merge pull request #106 from gkiki90/run_tests (2015 Aug 06)
* [7b86e33] Viktor Benei - _tests for test bitrise.ymls (2015 Aug 06)
* [5844849] Krisztian Goedrei - reference cycle test (2015 Aug 06)
* [ec2e74b] Viktor Benei - Merge pull request #105 from viktorbenei/master (2015 Aug 06)
* [fd98b01] Viktor Benei - Godeps update - pathutil (2015 Aug 06)
* [5bb93ed] Viktor Benei - 'setup' is now called automatically if it was not called for the current version of bitrise when 'run' is called + code revisions (2015 Aug 06)
* [4226622] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 06)
* [a561b0d] Viktor Benei - store `bitrise setup` for the given bitrise version, so that it can be checked whether a setup was done for the current version (2015 Aug 06)
* [ef60060] Viktor Benei - Merge pull request #104 from gkiki90/build_failed_test (2015 Aug 06)
* [407df1e] Viktor Benei - Merge pull request #103 from viktorbenei/master (2015 Aug 06)
* [4de3fe1] Krisztian Goedrei - ci.sh fix (2015 Aug 06)
* [710f90f] Krisztian Goedrei - fixed ci.sh (2015 Aug 06)
* [884006c] Krisztian Goedrei - fixed build failed (2015 Aug 06)
* [5f9e786] Krisztian Goedrei - doRun fixes, doRun_test (2015 Aug 06)
* [4942a3a] Krisztian Goedrei - reorganized code (2015 Aug 06)
* [f4e1ed3] Viktor Benei - switching to the new StepLib, hosted on GitHub (2015 Aug 06)
* [a683856] Viktor Benei - Merge pull request #102 from gkiki90/master (2015 Aug 05)
* [3aee65e] Krisztian Goedrei - removed bitrise from gitignore (2015 Aug 05)
* [b33ff0c] Viktor Benei - Merge pull request #101 from viktorbenei/master (2015 Aug 05)
* [6bd0268] Viktor Benei - start of v0.9.8 (2015 Aug 05)
* [17e1c54] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 05)


## 0.9.7 (2015 Aug 05)

### Release Notes

* __IMPORTANT__ : The project was renamed from `bitrise-cli` to just `bitrise`, which means that from now on you have to call your commands with `bitrise [command]`, instead of the previous, longer `bitrise-cli [command]`.
* Improved step dependency management with `brew`.
* Log improvements.

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.7/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.6 -> 0.9.7

* [c1759d3] Viktor Benei - Merge pull request #100 from gkiki90/master (2015 Aug 05)
* [a35b632] Viktor Benei - slack step titles (2015 Aug 05)
* [23a5966] Krisztian Goedrei - tool dependecies (2015 Aug 05)
* [c77a3c8] Krisztian Goedrei - Merge branch 'master' of github.com:bitrise-io/bitrise-cli (2015 Aug 05)
* [0f6629f] Krisztian Goedrei - changelog (2015 Aug 05)
* [c484396] Viktor Benei - Merge pull request #99 from viktorbenei/master (2015 Aug 05)
* [4bfdfc8] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 05)
* [4b9d9ca] Viktor Benei - added another Slack msg to announce (2015 Aug 05)
* [4e47b96] Viktor Benei - Merge pull request #98 from gkiki90/master (2015 Aug 05)
* [6d0a9ef] Viktor Benei - Godep update (2015 Aug 05)
* [98a0a76] Krisztian Goedrei - flag fixes (2015 Aug 05)
* [abf0257] Krisztian Goedrei - init highligth (2015 Aug 05)
* [82daa8f] Viktor Benei - Merge pull request #97 from viktorbenei/master (2015 Aug 05)
* [88ab022] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise (2015 Aug 05)
* [a06ac0d] Viktor Benei - Merge pull request #96 from gkiki90/skipped_handling (2015 Aug 05)
* [a9de033] Viktor Benei - Merge pull request #95 from gkiki90/runtime (2015 Aug 05)
* [2df0d40] Viktor Benei - renames, from the old `bitrise-cli` to the new, official `bitrise` tool name (2015 Aug 05)
* [4b16a00] Krisztian Goedrei - run if (2015 Aug 05)
* [2d93845] Krisztian Goedrei - runtime, dependency fixes (2015 Aug 05)
* [c5cae54] Viktor Benei - Merge pull request #94 from gkiki90/master (2015 Aug 05)
* [bc46501] Krisztian Goedrei - CI flag fixes (2015 Aug 05)
* [1a9a1ed] Krisztian Goedrei - chek with brew if installed (2015 Aug 05)
* [0fc0102] Viktor Benei - Merge pull request #93 from viktorbenei/master (2015 Aug 04)
* [93eb603] Viktor Benei - start of v0.9.7 (2015 Aug 04)
* [a205dec] Viktor Benei - changeling template + changeling generator added to create-release workflow, similar to the one in stepman&envman (2015 Aug 04)


## 0.9.6 (2015 Aug 04)

### Release Notes

* __BREAKING__ : `.bitrise.secrets.yml` 's syntax changed, to match the environments syntax used everywhere else. This means that instead of directly specifying `is_expand` at the same level as the key and value you should now move this into an `opts:` section, just like in every other `envs` list in `bitrise.yml`.
* __NEW__ : dependency management built into `bitrise.yml` syntax. Right now only `brew` is supported, on OS X, but this will be expanded.
* if a step or a version can't be found in the local cache from a step library `bitrise` will now update the local cache before failing with "step not found"
* greatly improved logs, colored step sections and step run summaries. It starts to look decent and is much more helpful than the previous log outputs.
* updated `setup` - only the Xcode Command Line tools are required now, if no full Xcode found it'll print a warning message about it but you can still use `bitrise`.
* quite a lot of minor bug fixes

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.6/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.5 -> 0.9.6

* [1dd2a00] Viktor Benei - Merge pull request #92 from gkiki90/master (2015 Aug 04)
* [96cab0c] Krisztian Goedrei - PR fixes (2015 Aug 04)
* [58c4154] Viktor Benei - Merge pull request #91 from gkiki90/envman_run_fix (2015 Aug 04)
* [9964b04] Krisztian Goedrei - PR fix (2015 Aug 04)
* [dea2cdf] Krisztian Goedrei - PR fixes (2015 Aug 04)
* [6c1e275] Krisztian Goedrei - running step log fix (2015 Aug 04)
* [9acbefb] Krisztian Goedrei - log fixes (2015 Aug 04)
* [e136db0] Krisztian Goedrei - run results fix (2015 Aug 04)
* [26d0296] Krisztian Goedrei - envman run with exit code, log fixes in progress (2015 Aug 04)
* [d2f25ff] Viktor Benei - Merge pull request #90 from viktorbenei/master (2015 Aug 04)
* [4b37119] Viktor Benei - Godep update : goinp bool parse improvement (2015 Aug 04)
* [4806c04] Viktor Benei - Merge pull request #89 from viktorbenei/master (2015 Aug 04)
* [48ec5b2] Viktor Benei - at setup the "[OK]" strings are now highlighted with green color; Xcode CLT setup/check revisions: warning if only CLT is available but not a full Xcode, with highlight, and as version it prints the info text (2015 Aug 04)
* [f643144] Viktor Benei - dependencies: no full Xcode required, only Command line tools (2015 Aug 03)
* [4795201] Viktor Benei - setup: envman v0.9.2 and stepman v0.9.6 required (2015 Aug 03)
* [71ce059] Viktor Benei - init: generate new style .secrets (2015 Aug 03)
* [1b5112e] Viktor Benei - Merge pull request #88 from gkiki90/dependencies (2015 Aug 03)
* [3635cfe] Krisztian Goedrei - print header fix (2015 Aug 03)
* [16946a4] Krisztian Goedrei - godep update (2015 Aug 03)
* [95d7529] Krisztian Goedrei - refactor fixes (2015 Aug 03)
* [c60cd93] Krisztian Goedrei - godep update (2015 Aug 03)
* [28943df] Krisztian Goedrei - depman update (2015 Aug 03)
* [9fd991f] Krisztian Goedrei - log CI mode, printASCIIHeader moved to ci.go (2015 Jul 31)
* [3c118a8] Krisztian Goedrei - dependencies (2015 Jul 31)
* [5ad6512] Krisztian Goedrei - dependency handling (2015 Jul 31)
* [575be1e] Krisztian Goedrei - dependencies in progress (2015 Jul 31)
* [1e707d7] Krisztian Goedrei - fixed env merge, models_methods_tests (2015 Jul 31)
* [3662056] Krisztian Goedrei - test start (2015 Jul 30)
* [b08a92d] Krisztian Goedrei - godep update (2015 Jul 30)
* [3e86ee6] Krisztian Goedrei - fixed stepman update flag, typo fix (2015 Jul 30)
* [630cc2a] Krisztian Goedrei - refactor, code style (2015 Jul 30)
* [9af552b] Viktor Benei - Merge pull request #86 from gkiki90/new_envman_models (2015 Jul 29)
* [1272839] Krisztian Goedrei - test fix (2015 Jul 29)
* [1fe5d13] Krisztian Goedrei - PR fix (2015 Jul 29)
* [d442821] Krisztian Goedrei - godep update (2015 Jul 29)
* [91324a2] Krisztian Goedrei - godep (2015 Jul 29)
* [57c2476] Krisztian Goedrei - godep (2015 Jul 29)
* [0301e85] Krisztian Goedrei - use envman models (2015 Jul 29)
* [9d52cfc] Viktor Benei - Merge pull request #85 from viktorbenei/master (2015 Jul 28)
* [654aa47] Viktor Benei - start of v0.9.6 (2015 Jul 28)


## 0.9.5 (2015 Jul 28)

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.5/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.4 -> 0.9.5

* [017e896] Viktor Benei - Merge pull request #84 from viktorbenei/master (2015 Jul 28)
* [87dc5a9] Viktor Benei - require Stepman 0.9.5 (2015 Jul 28)
* [4a20cb1] Viktor Benei - Merge pull request #83 from viktorbenei/master (2015 Jul 28)
* [04910f2] Viktor Benei - Godeps update (2015 Jul 28)
* [1ae7e36] Viktor Benei - Merge pull request #82 from gkiki90/log_improvements (2015 Jul 28)
* [a2975fa] Krisztian Goedrei - test (2015 Jul 28)
* [63ca923] Krisztian Goedrei - Merge branch 'master' into log_improvements (2015 Jul 28)
* [eefa57e] Krisztian Goedrei - build failed fix (2015 Jul 28)
* [dabdb97] Viktor Benei - Merge pull request #81 from viktorbenei/master (2015 Jul 28)
* [4ffbfa6] Viktor Benei - BITRISE_BUILD_STATUS and STEPLIB_BUILD_STATUS printing in before-after test in bitrise.yml (2015 Jul 28)
* [c698340] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 28)
* [cc9030a] Viktor Benei - Merge pull request #80 from gkiki90/log_improvements (2015 Jul 28)
* [9ca35a0] Krisztian Goedrei - template_utils_test BuildRunResultsModel fix (2015 Jul 28)
* [3e7f3ae] Krisztian Goedrei - fixed build failed mode (2015 Jul 28)
* [c5c94eb] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 28)
* [3cc6fad] Viktor Benei - Merge pull request #79 from gkiki90/log_improvements (2015 Jul 28)
* [c8a6ed0] Viktor Benei - before-after ENV accessibility test (2015 Jul 28)
* [ec7c9d2] Viktor Benei - Merge pull request #78 from viktorbenei/master (2015 Jul 28)
* [0929268] Krisztian Goedrei - log fixes, env handling fixes (2015 Jul 28)
* [68ff822] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 28)
* [6fbce1b] Viktor Benei - Merge pull request #77 from gkiki90/master (2015 Jul 28)
* [595f6fd] Krisztian Goedrei - validate fix (2015 Jul 28)
* [061bd82] Krisztian Goedrei - log fix in progress (2015 Jul 28)
* [6e8ce72] Krisztian Goedrei - comments, env imports (2015 Jul 28)
* [2eef2e8] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 28)
* [d2e8ca2] Viktor Benei - before_run for install (2015 Jul 28)
* [5bd2547] Viktor Benei - init : doesn't print the content anymore (2015 Jul 28)
* [59e2961] Viktor Benei - using rsync instead of cp to copy local path:: step source - it can handle the case if you want to run it from the step's dir directly, for example while developing the step (2015 Jul 28)
* [179ed5f] Viktor Benei - Merge pull request #76 from gkiki90/after_before (2015 Jul 28)
* [811f7e2] Krisztian Goedrei - godep (2015 Jul 28)
* [b4e3c9d] Krisztian Goedrei - PR fix (2015 Jul 28)
* [5d5fdd1] Krisztian Goedrei - Merge branch 'master' into after_before (2015 Jul 28)
* [39870f5] Krisztian Goedrei - log fixes (2015 Jul 28)
* [1a34984] Krisztian Goedrei - validating bitrisedata, workflow logs, (2015 Jul 28)
* [1b8747d] Viktor Benei - Merge pull request #75 from gkiki90/log_fix (2015 Jul 28)
* [4961f6c] Krisztian Goedrei - test (2015 Jul 27)
* [47e65d6] Krisztian Goedrei - run in progress (2015 Jul 27)
* [94d9574] Krisztian Goedrei - colorstring package and usage (2015 Jul 27)
* [821a6e6] Krisztian Goedrei - Merge branch 'master' into log_fix (2015 Jul 27)
* [388b536] Krisztian Goedrei - color log in progress (2015 Jul 27)
* [f4482de] Viktor Benei - Merge pull request #74 from gkiki90/step_source (2015 Jul 27)
* [4b59182] Krisztian Goedrei - bitrise.yml fix (2015 Jul 27)
* [189453c] Krisztian Goedrei - help messages fix, code style (2015 Jul 27)
* [70c1719] Krisztian Goedrei - timestamp-gen workflow fixes, step source log (2015 Jul 27)
* [60c5475] Viktor Benei - Merge pull request #73 from viktorbenei/master (2015 Jul 25)
* [b394e22] Viktor Benei - just a bit of template expression doc note (2015 Jul 25)
* [0af570e] Viktor Benei - Merge pull request #72 from viktorbenei/master (2015 Jul 25)
* [e178a1d] Viktor Benei - test for "$.Prop" style referencing, and annotated template examples (2015 Jul 25)
* [e99c312] Viktor Benei - enveq function, for easier ENV testing (2015 Jul 25)
* [1b482b8] Viktor Benei - A TemplateDataModel is now available for step property expressions, for easier "IsCI" detection. You can also just write ".IsCI", instead of the longer "{{.IsCI}}" the "CI=true" env is set at the start to force every tool to work in CI mode (even if the CI mode was just a command line param) (2015 Jul 25)
* [6dfd682] Viktor Benei - Merge pull request #71 from viktorbenei/run_if_and_templates (2015 Jul 24)
* [d232bad] Viktor Benei - first version of Run-If template handling & a couple of revisions (2015 Jul 24)
* [8ca8b9f] Viktor Benei - MergeStepWith #fix (2015 Jul 24)
* [026a752] Viktor Benei - Examples & tutorials section (2015 Jul 24)
* [d81a27b] Viktor Benei - Merge pull request #70 from viktorbenei/master (2015 Jul 24)
* [0a5a60b] Viktor Benei - the failing steps examples are also moved into examples/tutorials (2015 Jul 24)
* [e3bd023] Viktor Benei - a couple of experimentals (2015 Jul 24)
* [5cfd72d] Viktor Benei - examples/tutorials (2015 Jul 24)
* [afaf8e4] Viktor Benei - _examples folder to include a couple of example bitrise cli configs and workflows (2015 Jul 24)
* [c3d9605] Viktor Benei - Merge pull request #69 from viktorbenei/master (2015 Jul 24)
* [ccf3c89] Viktor Benei - Install instructions now points to /releases (2015 Jul 24)
* [052fb87] Viktor Benei - start of v0.9.5 (2015 Jul 24)


## 0.9.4 (2015 Jul 24)

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.4/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.3 -> 0.9.4

* [017b840] Viktor Benei - Merge pull request #68 from viktorbenei/master (2015 Jul 24)
* [5f5be0f] Viktor Benei - Godeps update (2015 Jul 24)
* [8db50a3] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 24)
* [bcb3ec3] Viktor Benei - Stepman update - related: IsNotImportant is now IsSkippable (2015 Jul 24)
* [80a1348] Viktor Benei - Merge pull request #67 from viktorbenei/master (2015 Jul 24)
* [1268b06] Viktor Benei - Godep-update workflow (2015 Jul 24)
* [6e1e1bf] Viktor Benei - ssh style remote git step in test workflows (2015 Jul 24)
* [0238347] Viktor Benei - fix: in case of direct git uri which contains @ as part of the url it should still work correctly (ex: if git url is: git@github.com:bitrise-io/steps-timestamp.git); no path to absolute-path conversion should happen in CreateStepIDDataFromString; unit tests for the new "path::" and "git::" style step IDs (2015 Jul 24)
* [760c42d] Viktor Benei - StepIDData : now that it supports local and direct-git-url options the previous ID was renamed to IDorURI and some documentation is provided for relevant places (2015 Jul 24)
* [3fc9807] Viktor Benei - in case the step source is defined as a direct git uri a version (branch or tag) is also required (2015 Jul 24)
* [ce65526] Viktor Benei - a bit more, and more thorough test workflows (2015 Jul 24)
* [ed71074] Viktor Benei - Merge pull request #66 from gkiki90/master (2015 Jul 24)
* [7e8c95f] Krisztian Goedrei - git src (2015 Jul 24)
* [ee8ce3a] Krisztian Goedrei - err check fixes (2015 Jul 24)
* [23a365a] Krisztian Goedrei - move local step to .bitrise work dir (2015 Jul 24)
* [4517333] Krisztian Goedrei - Merge branch 'master' of github.com:bitrise-io/bitrise-cli (2015 Jul 24)
* [c20bc6a] Krisztian Goedrei - PR fixes (2015 Jul 24)
* [0dd36ff] Krisztian Goedrei - support for ~/your/path (2015 Jul 24)
* [973b4fc] Krisztian Goedrei - local steps (2015 Jul 24)
* [b60f4a3] Krisztian Goedrei - local path in models and model_methods (2015 Jul 24)
* [f620295] Viktor Benei - Merge pull request #65 from viktorbenei/master (2015 Jul 24)
* [b5cc18b] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 24)
* [ac1564b] Viktor Benei - minor run command log formatting (2015 Jul 23)
* [da057fc] Viktor Benei - start of 0.9.4 (2015 Jul 23)
* [b700273] Viktor Benei - install - 0.9.3 (2015 Jul 23)


## 0.9.3 (2015 Jul 23)

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.3/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.2 -> 0.9.3

* [469750c] Viktor Benei - Merge pull request #64 from viktorbenei/master (2015 Jul 23)
* [8628a5a] Viktor Benei - requires stepman 0.9.3 (2015 Jul 23)
* [121afa1] Viktor Benei - Godeps update (2015 Jul 23)
* [ee301cd] Viktor Benei - Merge pull request #62 from viktorbenei/setup_improvements (2015 Jul 23)
* [269aec9] Viktor Benei - temp switch back to previous stepman min ver (2015 Jul 23)
* [1874e0a] Viktor Benei - Xcode CLT version check and better Brew warn (2015 Jul 23)
* [61d42cc] Viktor Benei - Merge branch 'master' into setup_improvements (2015 Jul 23)
* [41bdc3a] Viktor Benei - Setup can now update the required Bitrise Tools if an older version found (2015 Jul 23)
* [9fb88c3] Viktor Benei - Merge pull request #63 from gkiki90/build_status_env (2015 Jul 23)
* [bb503ca] Krisztian Goedrei - set build status env fix (2015 Jul 23)
* [db8e469] Krisztian Goedrei - build failed envs (2015 Jul 23)
* [b7b34c0] Viktor Benei - Godeps update: stepman (2015 Jul 23)
* [7c7878c] Viktor Benei - Merge pull request #61 from gkiki90/build_time (2015 Jul 23)
* [6b7de38] Krisztian Goedrei - total count fix (2015 Jul 23)
* [341b560] Krisztian Goedrei - log success count (2015 Jul 23)
* [13f4df0] Krisztian Goedrei - log fixes (2015 Jul 23)
* [8901ab9] Krisztian Goedrei - skipped steps (2015 Jul 23)
* [6066019] Krisztian Goedrei - log fixes (2015 Jul 23)
* [59b895b] Krisztian Goedrei - typo (2015 Jul 23)
* [075835b] Krisztian Goedrei - summary log (2015 Jul 23)
* [a0f3e5d] Krisztian Goedrei - build finish fixes (2015 Jul 23)
* [35d2390] Krisztian Goedrei - code style (2015 Jul 23)
* [afe217f] Krisztian Goedrei - failed step fix (2015 Jul 23)
* [6fcf9e4] Krisztian Goedrei - revision (2015 Jul 23)
* [b18d334] Krisztian Goedrei - revision (2015 Jul 23)
* [a241271] Krisztian Goedrei - fixed merge (2015 Jul 23)
* [ea56eef] Krisztian Goedrei - merge fix (2015 Jul 23)
* [f6136c5] Krisztian Goedrei - Merge branch 'master' into build_time (2015 Jul 23)
* [5787288] Krisztian Goedrei - register build status methods (2015 Jul 23)
* [b9d861a] Viktor Benei - Merge pull request #60 from viktorbenei/master (2015 Jul 22)
* [da46aba] Viktor Benei - start of v0.9.3 (2015 Jul 22)
* [0650b85] Viktor Benei - install - v0.9.2 (2015 Jul 22)


## 0.9.2 (2015 Jul 22)

### Install or upgrade

To install this version, run the following commands (in a bash shell):

```
curl -fL https://github.com/bitrise-io/bitrise/releases/download/0.9.2/bitrise-$(uname -s)-$(uname -m) > /usr/local/bin/bitrise
```

Then:

```
chmod +x /usr/local/bin/bitrise
```

That's all, you're ready to go!

Optionally, you can call `bitrise setup` to verify that everything what's required for bitrise to run
is installed and available, but if you forget to do this it'll be performed the first
time you call bitrise run.

### Release Commits - 0.9.1 -> 0.9.2

* [45e51d5] Viktor Benei - Merge pull request #59 from viktorbenei/master (2015 Jul 22)
* [7300a42] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 22)
* [8c75a3b] Viktor Benei - Goddess update (2015 Jul 22)
* [8f0da08] Viktor Benei - doSetup: install stepman v0.9.2 (2015 Jul 22)
* [a060d39] Viktor Benei - Merge pull request #58 from viktorbenei/master (2015 Jul 22)
* [86790aa] Viktor Benei - Merge pull request #57 from bazscsa/master (2015 Jul 22)
* [aa3c01d] Viktor Benei - ci: now does a build & calls setup on it (2015 Jul 22)
* [b641a4f] Viktor Benei - fixed possible infinite recursion in Setup (2015 Jul 22)
* [a1f6f45] Viktor Benei - create-release : now sends a Slack msg as well (2015 Jul 22)
* [2c31fc3] Viktor Benei - envman call fix: do adds with --append (2015 Jul 22)
* [f1d612f] Tamás Bazsonyi - Setup (2015 Jul 22)
* [4829d33] Viktor Benei - Merge pull request #56 from gkiki90/isNotImportant_handling (2015 Jul 22)
* [34a9326] Krisztian Goedrei - refactor, typo (2015 Jul 22)
* [6d66525] Krisztian Goedrei - build time in progress (2015 Jul 22)
* [15f0d63] Tamás Bazsonyi - setup description revision (2015 Jul 22)
* [f0ce128] Krisztian Goedrei - isNotImportent handling, buildFailedMode fixes (2015 Jul 22)
* [84a9d8d] Viktor Benei - Merge pull request #55 from viktorbenei/master (2015 Jul 22)
* [5b11e83] Viktor Benei - "environments" is now simply "envs" (2015 Jul 22)
* [26a3eae] Viktor Benei - doSetup : don't ask for permission to install required dependencies (envman & stepman) (2015 Jul 22)
* [0a9955d] Viktor Benei - Install command syntax change, for clarity (2015 Jul 22)
* [1b59895] Viktor Benei - Merge branch 'master' of https://github.com/bitrise-io/bitrise-cli (2015 Jul 22)
* [ba9a6e3] Viktor Benei - start of v0.9.2 - bitrise.yml now contains a 'create-release' workflow (2015 Jul 22)
* [b4cf5e8] Viktor Benei - just a minor format change for setup (2015 Jul 22)
* [3ae6a0e] Viktor Benei - Install and setup instructions (2015 Jul 22)
* [e38911e] Viktor Benei - setup: now can install stepman as well as envman (2015 Jul 22)
* [40cfd3a] Viktor Benei - better 'init' command : it now adds a default step lib source & a simple 'script' step with hello (2015 Jul 22)


-----------------

Updated: 2016 Apr 12
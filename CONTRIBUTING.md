# Contributing guidelines

## Before contributing

Welcome to nairobi-traffic-pcb! Before sending your pull requests, make sure that you **read the whole guidelines**. If you have any doubt on the contributing guide, please feel free to [state it clearly in an issue](https://github.com/0x6flab/nairobi-traffic-pcb/issues/new).

## Contributing

### Contributor

We are very happy that you are considering contributing to this project! Being one of our contributors, you agree and confirm that:

- Your work will be distributed under [GNU GENERAL PUBLIC LICENSE](LICENSE.md) once your pull request is merged
- Your submitted work fulfils or mostly fulfils our styles and standards

**Improving comments** and **writing proper tests** are also highly welcome.

### Contribution

We appreciate any contribution, from fixing a grammar mistake in a comment to implementing complex algorithms. Please read this section if you are contributing your work.

Your contribution will be tested by our [automated testing on GitHub Actions](https://github.com/0x6flab/nairobi-traffic-pcb/actions) to save time and mental energy. After you have submitted your pull request, you should see the GitHub Actions tests start to run at the bottom of your submission page. If those tests fail, then click on the **_details_** button try to read through the GitHub Actions output to understand the failure. If you do not understand, please leave a comment on your submission page.

If you are interested in resolving an [open issue](https://github.com/0x6flab/nairobi-traffic-pcb/issues), simply make a pull request with your proposed fix. **We do not assign issues in this repo** so please do not ask for permission to work on an issue.

Please help us keep our issue list small by adding `Fixes #{$ISSUE_NUMBER}` to the description of pull requests that resolve open issues.
For example, if your pull request fixes issue #10, then please add the following to its description:

```git
Fixes #10
```

GitHub will use this tag to [auto-close the issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue) if and when the PR is merged.

#### What is an nairobi-traffic-pcb?

nairobi-traffic-pcb is a PCB that shows traffic artwork of Nairobi Public Transport.

#### Coding Style

We want our work to be readable by others; therefore, we encourage you to note the following:

- Please focus hard on the naming of functions, classes, and variables. Help your reader by using **descriptive names** that can help you to remove redundant comments.
- Single letter variable names are _old school_ so please avoid them unless their life only spans a few lines.
- Expand acronyms because `gcd()` is hard to understand but `greatest_common_divisor()` is not.

#### Other Requirements for Submissions

- Strictly use snake_case (underscore_separated) in your file_name, as it will be easy to parse in future using scripts.
- Please avoid creating new directories if at all possible. Try to fit your work into the existing directory structure.
- If possible, follow the standard _within_ the folder you are submitting to.
- If you have modified/added code work, make sure the code compiles before submitting.
- If you have modified/added documentation work, ensure your language is concise and contains no grammar errors.

- Most importantly,
  - **Be consistent in the use of these guidelines when submitting.**
  - Happy coding!
